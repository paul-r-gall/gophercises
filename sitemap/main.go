package main

import (
	"encoding/xml"
	"flag"
	"log"
	"net/url"
	"os"

	findlink "github.com/paul-r-gall/gophercises/linkparse/src"
)

type XMLURL struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	URLs    []XMLURL
}

func main() {
	var urlStr string
	flag.StringVar(&urlStr, "url", "https://google.com", "build the sitemap of the given url")

	flag.Parse()

	urlAct, err := url.Parse(urlStr)
	if err != nil {
		log.Fatal(err)
	}
	host := urlAct.Host
	linkSet := make(map[string]bool)
	linkSet[urlStr] = true
	xmlurlset := URLSet{XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	var DFSLink func(string)
	DFSLink = func(s string) {
		linkList := findlink.FindLinks(s)
		for _, link := range linkList {
			urlAct, err := url.Parse(link.Href)
			if err != nil {
				log.Fatal(err)
			}
			if (urlAct.Host == host || urlAct.Host == "") && !linkSet[link.Href] {
				linkSet[link.Href] = true
				xmlurlset.URLs = append(xmlurlset.URLs, XMLURL{Loc: link.Href})
				DFSLink(link.Href)
			}

		}
		return
	}

	DFSLink(urlStr)
	output, err := xml.MarshalIndent(xmlurlset, "", "	")
	os.Stdout.Write(output)
}
