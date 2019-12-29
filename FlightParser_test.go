package main

import (
	"./datatypes"
	"./parser"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
	"os"
)

func TestDataCanBeRetrieved(t *testing.T) {
	os.Setenv("OPENSKY_LATITUDE_MIN", "53.477820")
	os.Setenv("OPENSKY_LONGITUDE_MIN", "9.760569")
	os.Setenv("OPENSKY_LATITUDE_MAX", "53.730380")
	os.Setenv("OPENSKY_LONGITUDE_MAX", "10.326908")

	url := parser.GetParameterizedUrl()
	log.Printf(url)
}

func TestDataCanBeConverted(t *testing.T) {
	jsonFile, err := os.Open("data/test.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	var states datatypes.OpenSkyJsonStruct

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &states)

	flightData := parser.ConvertToFlightData(states)

	if len(flightData) != 2039 {
		t.Errorf("Size of parsed objects is %d, expected are 2039", len(flightData))
	}

	for i := 0; i < len(flightData); i++ {
		log.Printf("Flight %d is: %s", i, (flightData[i]))
	}
}
