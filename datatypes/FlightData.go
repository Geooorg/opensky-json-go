package datatypes

type FlightData struct {
	id          uint32
	dateAndTime uint32
	altitude    float32
	icao        string
	callsign    string
	latitude    float32
	longitude   float32
	distance    float32
	after22     bool
	after24     bool
	landing     bool
	degree      uint8
}
