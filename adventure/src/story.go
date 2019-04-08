package story

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"
)

type ArcOption struct {
	Text string `json:"text"`
	ArcLabel string `json:"arc"`
}

type Arc struct {
	Title string `json:"title"`
	Story []string `json:"story"`
	Options []ArcOption `json:"options"`
}

// parses JSON to a map from arc-names to Arcs and returns that
func ParseJSONtoStory(jsonFile string) (map[string]Arc, error) {
	jsonText, err := ioutil.ReadFile(jsonFile)
	if err!=nil {
		log.Println(err)
	}
	fs := make(map[string]Arc)
	err = json.Unmarshal(jsonText, &fs)
	return fs,err
}


