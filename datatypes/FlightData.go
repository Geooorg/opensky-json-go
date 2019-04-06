package datatypes

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
	Degree      uint8
}
