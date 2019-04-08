package main

import (
	"encoding/json"
	"fmt"
	"github.com/Geooorg/opensky-json-go/datatypes"
	"github.com/Geooorg/opensky-json-go/parser"
	"io/ioutil"
	"log"
	"testing"

	"os"
)

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
