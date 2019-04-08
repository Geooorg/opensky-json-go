package main

import (
	"bytes"
	"encoding/json"
	"errors"
	openskyDataTypes "github.com/Geooorg/opensky-json-go/datatypes"
	"github.com/Geooorg/opensky-json-go/parser"
	"io"
	"log"
	"net/http"
)

const OpenskyNetworkUrl = "https://openskyDataTypes-network.org/api/states/all"

func main() {

	jsonStr, e := readJsonFromOpenSky()
	if e != nil {
		log.Fatal("Could not read from " + OpenskyNetworkUrl)
	}

	var states openskyDataTypes.OpenSkyJsonStruct

	json.Unmarshal(jsonStr, &states)

	flightData := parser.ConvertToFlightData(states)
	log.Printf("Found %d flights", len(flightData))
}

func readJsonFromOpenSky() ([]byte, error) {
	log.Printf("Reading flight JSON from %s", OpenskyNetworkUrl)
	response, err := http.Get(OpenskyNetworkUrl)
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
