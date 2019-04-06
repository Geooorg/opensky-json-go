package datatypes

type OpenSkyJsonStruct struct {
	Time              float64         `json:"time"`
	StatesListOfLists [][]interface{} `json:"states"`
}
