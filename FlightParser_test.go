package main

import (
	"encoding/json"
	"fmt"
	opensky "github.com/Geooorg/opensky-json-go/datatypes"
	"io/ioutil"
	"log"

	"os"
)

func testDataCanBeConverted() {
	jsonFile, err := os.Open("data/test.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	var states opensky.OpenSkyJsonStruct

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &states)

	flightData := convertToFlightData(states)
	if len(flightData) != 2039 {
		log.Fatal("Size of parsed objects does not match")
	}
}
