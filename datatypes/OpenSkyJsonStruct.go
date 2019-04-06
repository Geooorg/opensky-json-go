package datatypes

type OpenSkyJsonStruct struct {
	Time              int        `json:"time"`
	StatesListOfLists [][]string `json:"states"`
}
