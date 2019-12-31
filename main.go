package main

import (
	"./parser"
	"encoding/json"
	"log"
)

func main() {

	jsonStr, e := parser.ReadJsonFromOpenSky()
	if e != nil {
		log.Fatal("Could not read data from OpenSky")
	}

	var states parser.OpenSkyJsonStruct

	json.Unmarshal(jsonStr, &states)

	flightData := parser.ConvertToFlightData(states)
	log.Printf("Found %d flights", len(flightData))
}
