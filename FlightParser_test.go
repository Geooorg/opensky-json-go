package main

import (
	"./parser"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestUrlIsParameterizedCorrectly(t *testing.T) {
	os.Setenv("OPENSKY_LATITUDE_MIN", "53.477820")
	os.Setenv("OPENSKY_LONGITUDE_MIN", "9.760569")
	os.Setenv("OPENSKY_LATITUDE_MAX", "53.730380")
	os.Setenv("OPENSKY_LONGITUDE_MAX", "10.326908")

	url := parser.GetParameterizedUrl()
	assert.Equal(t, url, "https://opensky-network.org/api/states/all?lamin=53.477820&lomin=9.760569&lamax=53.730380&lomax=10.326908", "URL without authentication matches")

	os.Setenv("OPENSKY_USER", "sampleUser")
	os.Setenv("OPENSKY_PASSWORD", "verysecret!")

	urlWithAuthentication := parser.GetParameterizedUrl()
	assert.Equal(t, urlWithAuthentication, "https://sampleUser:verysecret!@opensky-network.org/api/states/all?lamin=53.477820&lomin=9.760569&lamax=53.730380&lomax=10.326908", "URL with authentication matches")
}

func TestDataCanBeConverted(t *testing.T) {
	jsonFile, err := os.Open("data/test.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	var states OpenSkyJsonStruct

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &states)

	flightData := parser.ConvertToFlightData(states)

	if len(flightData) != 2039 {
		t.Errorf("Size of parsed objects is %d, expected are 2039", len(flightData))
	}

	for i := 0; i < len(flightData); i++ {
		log.Printf("Flight %d is: %s", i, (flightData[i]))
	}

	flight_EZY64KP := flightData[1]
	assert.Equal(t, flight_EZY64KP.Callsign, "EZY64KP")
	assert.Equal(t, flight_EZY64KP.Id, "406b90", "Icao is 406b90")
	assert.Equal(t, flight_EZY64KP.DateAndTime, "1483905638", "Unix timestamp matches")
	assert.Equal(t, flight_EZY64KP.Altitude, 3505.2)
	assert.Equal(t, flight_EZY64KP.Latitude, 49.2815)
	assert.Equal(t, flight_EZY64KP.Longitude, 1.9863)
	assert.Equal(t, flight_EZY64KP.Country, "United Kingdom")
	assert.Equal(t, flight_EZY64KP.Landing, true)
	assert.Equal(t, flight_EZY64KP.Degree, 94.25)
}
