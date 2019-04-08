package main

import (
	"./src"
	"fmt"
	"net/http"
	"html/template"
	"log"
)

type StoryHandler struct {
	fullStory map[string]story.Arc
}

func (sh StoryHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	templ, err := template.ParseFiles("hFiles/arcTemplate.html")
	if err!=nil{
		log.Fatal(err)
	}
	// p will be of the form "/label"
	label := r.URL.Path[1:]

	var dArc story.Arc
	if label == "" {
		dArc = sh.fullStory["intro"]
	} else {
		dArc = sh.fullStory[label]
	}
	
	templ.Execute(w, dArc)

	return
}
func main() {
	fs, err := story.ParseJSONtoStory("story.json")
	if err!=nil {
		fmt.Println(err)
	}
	
	http.ListenAndServe(":5000", StoryHandler{fs})

	//fmt.Println(fs)
}