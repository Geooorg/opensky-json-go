package parser

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const OPENSKY_URL_PROTOCOL = "https://"
const OPENSKY_URL_TEMPLATE = "opensky-network.org/api/states/all?lamin=%s&lomin=%s&lamax=%s&lomax=%s"
const OPENSKY_AUTHENTICATON_PREFIX_TEMPLATE = "%s:%s@" + OPENSKY_URL_TEMPLATE
const HTTP_TIMEOUT = 10

type Api struct {
}

type PublicApi interface {
	ReadFromWebserviceAndConvertJsonToFlightData() []FlightData
}

type OpenSkyJsonStruct struct {
	Time              int             `json:"time"`
	StatesListOfLists [][]interface{} `json:"states"`
}

func (api Api) ReadFromWebserviceAndConvertJsonToFlightData() []FlightData {
	jsonStr, e := readJsonFromOpenSky()
	if e != nil {
		log.Print("Could not read data from OpenSky")
		return nil
	}

	var states OpenSkyJsonStruct
	json.Unmarshal(jsonStr, &states)
	return ConvertToFlightData(states)
}

func ConvertToFlightData(states OpenSkyJsonStruct) []FlightData {
	var result = make([]FlightData, len(states.StatesListOfLists))

	for i := 0; i < len(states.StatesListOfLists); i++ {

		state := states.StatesListOfLists[i]

		f := FlightData{}

		f.Icao = strings.TrimSpace(state[0].(string))
		f.Callsign = strings.TrimSpace(state[1].(string))
		f.Country = strings.TrimSpace(state[2].(string))

		if state[3] != nil {
			f.Timestamp = int(state[3].(float64))
		}

		if state[6] != nil {
			f.Latitude = state[6].(float64)
		}
		if state[5] != nil {
			f.Longitude = state[5].(float64)
		}
		if state[7] != nil {
			f.Altitude = state[7].(float64)
		}
		if state[10] != nil {
			f.Degree = state[10].(float64)
		}
		if state[11] != nil {
			f.Landing = state[11].(float64) < 0
		}

		result[i] = f
	}

	return result
}

func readJsonFromOpenSky() ([]byte, error) {

	paramaterizedUrl := GetParameterizedUrl()

	log.Printf("DEBUG: Reading flight JSON from %s", paramaterizedUrl)

	client := http.Client{
		Timeout: HTTP_TIMEOUT * time.Second,
	}
	response, err := client.Get(paramaterizedUrl)
	if err != nil {
		log.Printf("WARN: Reading OpenSky data failed! :-( ... %s", err.Error())
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	var data bytes.Buffer
	_, err = io.Copy(&data, response.Body)
	if err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}

func GetParameterizedUrl() string {

	user, withAuthentication := os.LookupEnv("OPENSKY_USER")
	password := os.Getenv("OPENSKY_PASSWORD")

	lamin := os.Getenv("OPENSKY_LATITUDE_MIN")
	lomin := os.Getenv("OPENSKY_LONGITUDE_MIN")
	lamax := os.Getenv("OPENSKY_LATITUDE_MAX")
	lomax := os.Getenv("OPENSKY_LONGITUDE_MAX")

	if !withAuthentication {
		return fmt.Sprintf(OPENSKY_URL_PROTOCOL+OPENSKY_URL_TEMPLATE, lamin, lomin, lamax, lomax)
	}

	return fmt.Sprintf(OPENSKY_URL_PROTOCOL+OPENSKY_AUTHENTICATON_PREFIX_TEMPLATE, user, password, lamin, lomin, lamax, lomax)
}
