package datatypes

type FlightData struct {
	Id          uint32
	DateAndTime uint32
	Altitude    float32
	Icao        string
	Callsign    string
	Latitude    float32
	Longitude   float32
	Distance    float32
	After22     bool
	After24     bool
	Landing     bool
	Degree      uint8
}
