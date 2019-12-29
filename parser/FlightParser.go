package parser

import (
	//opensky  "github.com/Geooorg/opensky-json-go/datatypes"
	opensky "../datatypes" //  "github.com/Geooorg/opensky-json-go/datatypes"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const OPENSKY_URL_TEMPLATE = "https://opensky-network.org/api/states/all?lamin=%s&lomin=%s&lamax=%s&lomax=%s"
const HTTP_TIMEOUT = 10

func ConvertToFlightData(states opensky.OpenSkyJsonStruct) []opensky.FlightData {
	var result = make([]opensky.FlightData, len(states.StatesListOfLists))

	for i := 0; i < len(states.StatesListOfLists); i++ {

		state := states.StatesListOfLists[i]

		f := opensky.FlightData{}

		f.Id = strings.TrimSpace(state[0].(string))
		f.Callsign = strings.TrimSpace(state[1].(string))
		f.Country = strings.TrimSpace(state[2].(string))

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

func ReadJsonFromOpenSky() ([]byte, error) {

	paramaterizedUrl := GetParameterizedUrl()

	log.Printf("INFO: Reading flight JSON from %s", paramaterizedUrl)

	client := http.Client{
		Timeout: HTTP_TIMEOUT * time.Second,
	}
	response, err := client.Get(paramaterizedUrl)
	if err != nil {
		log.Printf("WARN: Reading failed! :-( ... %s", err.Error())
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

	lamin := os.Getenv("OPENSKY_LATITUDE_MIN")
	lomin := os.Getenv("OPENSKY_LONGITUDE_MIN")
	lamax := os.Getenv("OPENSKY_LATITUDE_MAX")
	lomax := os.Getenv("OPENSKY_LONGITUDE_MAX")

	return fmt.Sprintf(OPENSKY_URL_TEMPLATE, lamin, lomin, lamax, lomax)
}