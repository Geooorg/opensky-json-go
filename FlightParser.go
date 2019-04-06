package main

import (
	"bytes"
	"encoding/json"
	"errors"
	opensky "github.com/Geooorg/opensky-json-go/datatypes"
	"io"
	"log"
	//"log"
	"net/http"
)

const OPENSKY_NETWORK_URL = "https://opensky-network.org/api/states/all"

func main() {

	jsonStr, e := readJsonFromOpenSky()
	if e != nil {
		log.Fatal("Could not read from " + OPENSKY_NETWORK_URL)
	}

	var states opensky.OpenSkyJsonStruct

	json.Unmarshal(jsonStr, &states)

	flightData := convertToFlightData(states)
	log.Printf("Found %d flights", len(flightData))
}

func convertToFlightData(states opensky.OpenSkyJsonStruct) []opensky.FlightData {
	var result = make([]opensky.FlightData, len(states.StatesListOfLists))

	for i := 0; i < len(states.StatesListOfLists); i++ {

		state := states.StatesListOfLists[i]

		f := opensky.FlightData{}

		f.Id = state[0].(string)
		f.Callsign = state[1].(string)
		f.Country = state[2].(string)

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
	log.Printf("Reading flight JSON from %s", OPENSKY_NETWORK_URL)
	response, err := http.Get(OPENSKY_NETWORK_URL)
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
