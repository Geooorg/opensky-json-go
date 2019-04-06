package datatypes

type OpenSkyJsonStruct struct {
	Time              int             `json:"time"`
	StatesListOfLists [][]interface{} `json:"states"`
}
