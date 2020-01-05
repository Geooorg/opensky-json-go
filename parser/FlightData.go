package parser

import (
	"bytes"
	"fmt"
)

type FlightData struct {
	Country   string
	Timestamp int
	Altitude  float64
	Icao      string
	Callsign  string
	Latitude  float64
	Longitude float64
	Landing   bool
	Degree    float64
}

func (it FlightData) String() string {
	var buffer bytes.Buffer

	isLanding := "false"
	if it.Landing {
		isLanding = "true"
	}
	buffer.WriteString(", Country " + it.Country)
	buffer.WriteString(", Timestamp " + fmt.Sprint(it.Timestamp))
	buffer.WriteString(", Altitude " + fmt.Sprintf("%f", it.Altitude))
	buffer.WriteString(", Icao " + it.Icao)
	buffer.WriteString(", Callsign " + it.Callsign)
	buffer.WriteString(", Latitude " + fmt.Sprintf("%f", it.Latitude))
	buffer.WriteString(", Longitude " + fmt.Sprintf("%f", it.Longitude))
	buffer.WriteString(", Landing " + isLanding)
	buffer.WriteString(", Degree " + fmt.Sprintf("%f", it.Degree))

	return buffer.String()
}
