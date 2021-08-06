package smap

import (
	"encoding/xml"
	"io/ioutil"
	"links"
	"log"
	"net/http"
	"strings"
)

type url struct {
	Loc string `xml:"loc"`
}

type urlSet struct {
	Xmlns string `xml:"xmlns,attr"`
	Urls  []url  `xml:"url"`
}

var baseUrl string
var maxDepth int

//DoURLS does the link collection for the initial URL
func DoURLs(url string, depth int) map[string]bool {
	baseUrl = url
	maxDepth = depth
	var result = make(map[string]bool)
	return CollectLinks(url, result, 1)
}

//CollectLinks gets the current page and collect the links in it recursively
func CollectLinks(url string, result map[string]bool, depth int) map[string]bool {
	if depth > maxDepth {
		return result
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	initialLinks := links.ParseContent(resp.Body)
	defer resp.Body.Close()

	filteredLinks := filterLinks(initialLinks)
	for _, link := range filteredLinks {
		if strings.HasPrefix(link.Href, "/") {
			continue
		}

		result[link.Href] = true

		result = CollectLinks(link.Href, result, depth+1)
	}

	return result

}

//WriteXML converts a slice of links into an XML file
func WriteXML(domain string, results map[string]bool) {
	urls := new(urlSet)
	urls.Xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
	for link := range results {
		urls.Urls = append(urls.Urls, url{Loc: link})
	}

	output := xml.Header

	xmlResult, err := xml.MarshalIndent(urls, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	output += string(xmlResult)

	err = ioutil.WriteFile("output.xml", []byte(output), 0777)
	if err != nil {
		log.Fatal(err)
	}

}

//filterLinks returns the links in an array that contain the given url or are local liks
func filterLinks(values []links.Link) []links.Link {
	var results []links.Link
	for _, val := range values {
		if val.Href == "/" {
			continue
		}

		if strings.HasPrefix(val.Href, "/") {
			val.Href = baseUrl + val.Href
			results = append(results, val)
		} else if strings.Contains(val.Href, baseUrl) {
			results = append(results, val)
		}
	}
	return results

}
