package main

import (
	"fmt"
	"net/http"
	"./handler"
	"flag"
	"io/ioutil"
	"log"
)

func main() {

	filePtr := flag.String("yaml", "NOFILE", "file where yaml config is stored")
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// retrieve the yaml data from the file given by the flag.
	// if there is no flag file, or 
	var yaml []byte
	flag.Parse()
	file := *filePtr
	if file == "NOFILE" {
		yaml = []byte(`
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`)
	} else {
		var err error
		yaml, err = ioutil.ReadFile(file)
		if err != nil {
			log.Println(err)
			fmt.Println("File Read Invalid -- Using default YAML")
			yaml = []byte(`
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`)
		}
	}

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}