package datatypes

import (
	"bytes"
	"fmt"
)

type FlightData struct {
	Id          string
	Country     string
	DateAndTime uint32
	Altitude    float64
	Icao        string
	Callsign    string
	Latitude    float64
	Longitude   float64
	Distance    float64
	After22     bool
	After24     bool
	Landing     bool
	Degree      float64
}

func (it FlightData) String() string {
	var buffer bytes.Buffer
	isAfter22 := "false"
	if it.After22 {
		isAfter22 = "true"
	}
	isAfter24 := "false"
	if it.After24 {
		isAfter24 = "true"
	}
	isLanding := "false"
	if it.Landing {
		isLanding = "true"
	}
	buffer.WriteString("Id " + it.Id)
	buffer.WriteString(", Country " + it.Country)
	buffer.WriteString(", DateAndTime " + fmt.Sprint(it.DateAndTime))
	buffer.WriteString(", Altitude " + fmt.Sprintf("%f", it.Altitude))
	buffer.WriteString(", Icao " + it.Icao)
	buffer.WriteString(", Callsign " + it.Callsign)
	buffer.WriteString(", Latitude " + fmt.Sprintf("%f", it.Latitude))
	buffer.WriteString(", Longitude " + fmt.Sprintf("%f", it.Longitude))
	buffer.WriteString(", After22 " + isAfter22)
	buffer.WriteString(", After24 " + isAfter24)
	buffer.WriteString(", Landing " + isLanding)
	buffer.WriteString(", Degree " + fmt.Sprintf("%f", it.Degree))

	return buffer.String()
}
