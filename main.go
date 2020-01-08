package main

import (
	"./parser"
	"log"
)

func main() {
	flightData := parser.Api{}.ReadFromWebserviceAndConvertJsonToFlightData()
	log.Printf("Found %d flights", len(flightData))
}
