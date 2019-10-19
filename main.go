package main

import (
	"./datatypes" //"github.com/Geooorg/opensky-json-go/datatypes"
	"./parser"
	"encoding/json"
	"log"
)

func main() {

	jsonStr, e := parser.ReadJsonFromOpenSky()
	if e != nil {
		log.Fatal("Could not read from " + parser.OPENSKY_URL)
	}

	var states datatypes.OpenSkyJsonStruct

	json.Unmarshal(jsonStr, &states)

	flightData := parser.ConvertToFlightData(states)
	log.Printf("Found %d flights", len(flightData))
}
