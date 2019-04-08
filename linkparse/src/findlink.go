package findlink

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Link contains an Href and Content field.
type Link struct {
	Href    string
	Content string
}

func parseNode(n *html.Node, b *bytes.Buffer, top *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.TextNode {
		b.WriteString(string(n.Data))
	}
	parseNode(n.FirstChild, b, top)
	if n == top {
		return
	}
}

func checkFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func isLink(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

func nodeDFS(root *html.Node, aList []*html.Node) {
	if root == nil {
		return
	}
	if isLink(root) {
		aList = append(aList, root)
	} else {
		nodeDFS(root.FirstChild, aList)
	}

	if root.NextSibling != nil {
		nodeDFS(root.NextSibling, aList)
	}
	return
}

// FindLinks converts a siteURL to a list of all its contained hyperlinks.
func FindLinks(siteURL string) []Link {
	resp, err := http.Get(siteURL)
	for err != nil {
		fmt.Println("Enter a valid URL")
		fmt.Scanln(&siteURL)
		resp, err = http.Get(siteURL)
	}

	defer resp.Body.Close()
	src, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(strings.NewReader(string(src)))
	checkFatal(err)

	var aList []*html.Node
	nodeDFS(doc, aList)

	var linkArr []Link

	for _, link := range aList {
		lBuf := new(bytes.Buffer)
		parseNode(link, lBuf, link)
		var href string
		for _, att := range link.Attr {
			if att.Key == "href" {
				href = att.Val
			}
		}
		linkArr = append(linkArr, Link{
			Href:    href,
			Content: lBuf.String()})
	}
	return linkArr
}
