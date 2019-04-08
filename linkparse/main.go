package main

import (
	"flag"
	"fmt"

	findlink "./src"
)

func main() {
	var siteURL string
	flag.StringVar(&siteURL, "site", "https://google.com", "site to parse")

	m := findlink.FindLinks(siteURL)

	fmt.Println(m)
}
