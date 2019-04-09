package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	findlink "../linkparse/src"
)

//XMLURL ...
type XMLURL struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

//URLSet ...
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
	linkSet[urlAct.String()] = true
	xmlurlset := URLSet{XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9"}

	var DFSLink func(string)
	DFSLink = func(s string) {
		if s == "https:void(0)" {
			return
		}
		fmt.Println(s)
		linkList := findlink.FindLinks(s)
		if len(linkList) == 0 {
			return
		}
		for _, link := range linkList {
			urlAct, err := url.Parse(link.Href)
			if err != nil {
				fmt.Println("fatal link parse error")
				fmt.Printf("%T\n", link.Href)
				fmt.Println(link.Href)
				log.Fatal(err)
			}
			if urlAct.String() == "void(0)" {
				continue
			}
			if urlAct.Scheme == "mailto" {
				//fmt.Println(urlAct.Path)
				continue
			}

			if urlAct.Host == "" {
				urlAct.Host = host
				urlAct.Scheme = "https"
				//fmt.Println(urlAct.String())
			}
			urlAct.Scheme = "https"
			if (urlAct.Host == host) && !linkSet[urlAct.String()] {
				//fmt.Println(urlAct)
				linkSet[urlAct.String()] = true
				xmlurlset.URLs = append(xmlurlset.URLs, XMLURL{Loc: link.Href})
				DFSLink(urlAct.String())
			}

		}
		return
	}

	DFSLink(urlAct.String())
	output, err := xml.MarshalIndent(xmlurlset, "", "	")
	os.Stdout.Write(output)
}
