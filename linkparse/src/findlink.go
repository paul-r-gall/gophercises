package findlink

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// Link contains an Href and Content field.
type Link struct {
	Href    string
	Content string
}

func cleanWhitespace(s string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(s, " ")
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
	parseNode(n.NextSibling, b, top)
}

func checkFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func isLink(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

func nodeDFS(root *html.Node, linkList *[]Link) {
	if root == nil {
		return
	}
	if isLink(root) {
		var href string
		for _, att := range root.Attr {
			if att.Key == "href" {
				href = att.Val
			}
		}
		textBuffer := new(bytes.Buffer)
		parseNode(root, textBuffer, root)
		*linkList = append(*linkList, Link{
			Href:    href,
			Content: cleanWhitespace(textBuffer.String()),
		})

	} else {
		nodeDFS(root.FirstChild, linkList)
	}
	nodeDFS(root.NextSibling, linkList)
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

	var linkArr []Link
	nodeDFS(doc, &linkArr)

	return linkArr
}
