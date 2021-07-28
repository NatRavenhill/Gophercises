package links

import (
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Link = struct {
	Href string
	Text string
}

var validTags = []atom.Atom{atom.Strong, atom.Span, atom.B, atom.I}

//ExtractLinks traverses the tree and extracts links
func extractLinks(node *html.Node) []Link {
	var links []Link
	if node.Type == html.ElementNode && node.DataAtom == atom.A {
		for _, a := range node.Attr {
			if a.Key == "href" && node.FirstChild != nil {
				link := Link{a.Val, buildText(node)}
				links = append(links, link)
				break
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, extractLinks(c)...)
	}

	return links
}

//ParseContent parses a given input into a html tree
func ParseContent(content io.Reader) []Link {
	tree, err := html.Parse(content)
	if err != nil {
		log.Fatal(err)
	}

	return extractLinks(tree)
}

//BuildText builds the text value from the a tag's content
func buildText(node *html.Node) string {
	var result string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			result += strings.TrimSpace(c.Data)
		}
		if c.Type == html.ElementNode && containsTag(c.DataAtom) {
			toAdd := buildText(c)
			if len(toAdd) > 0 {
				result += " " + toAdd
			}
		}
	}
	return result
}

//ContainsTag checks if the given tag is in the valid tags
func containsTag(tag atom.Atom) bool {
	for _, val := range validTags {
		if val == tag {
			return true
		}
	}

	return false
}
