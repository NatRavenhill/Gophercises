// Package link parses an HTML file and extracts all of its links into a data structure
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Link = struct {
	Href string
	Text string
}

var Links []Link

func main() {
	tree := parseFile("ex1.html")
	extractLinks(tree)

	for _, val := range Links {
		fmt.Println(val)
	}
}

//ParseFile parses a given file into a html tree
func parseFile(filename string) *html.Node {
	content, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	tree, err := html.Parse(content)
	if err != nil {
		log.Fatal(err)
	}

	return tree
}

//ExtractLinks traverses the tree and extracts links
func extractLinks(node *html.Node) {
	if node.Type == html.ElementNode && node.DataAtom == atom.A {
		for _, a := range node.Attr {
			if a.Key == "href" && node.FirstChild != nil {
				link := Link{a.Val, buildText(node)}
				Links = append(Links, link)
				break
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractLinks(c)
	}
}

//BuildText builds the text value from the a tag's content
func buildText(node *html.Node) string {
	var result string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			result += strings.TrimSpace(c.Data)
		}
		if c.Type == html.ElementNode && c.DataAtom == atom.Strong {
			toAdd := buildText(c)
			if len(toAdd) > 0 {
				result += " " + toAdd
			}
		}
	}
	return result
}
