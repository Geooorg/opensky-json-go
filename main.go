package main

import (
	"encoding/json"
	openskyDataTypes "github.com/Geooorg/opensky-json-go/datatypes"
	"github.com/Geooorg/opensky-json-go/parser"
	"log"
)

func main() {

	jsonStr, e := parser.ReadJsonFromOpenSky()
	if e != nil {
		log.Fatal("Could not read from " + parser.OpenskyNetworkUrl)
	}

	var states openskyDataTypes.OpenSkyJsonStruct

	json.Unmarshal(jsonStr, &states)

	flightData := parser.ConvertToFlightData(states)
	log.Printf("Found %d flights", len(flightData))
}
