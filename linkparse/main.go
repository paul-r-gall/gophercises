package main

import (
	"flag"
	"fmt"

	findlink "./src"
)

func main() {
	var siteURL string
	flag.StringVar(&siteURL, "site", "https://google.com", "site to parse")

	flag.Parse()

	fmt.Println(siteURL)
	m := findlink.FindLinks(siteURL)

	fmt.Println(m)
}
