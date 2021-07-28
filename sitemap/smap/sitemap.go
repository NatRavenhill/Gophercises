package smap

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"strings"

	"links"
)

type url struct {
	Loc string `xml:"loc"`
}

type urlSet struct {
	Xmlns string `xml:"xmlns,attr"`
	Urls  []url  `xml:"url"`
}

//WriteXML converts a slice of links into an XML file
func WriteXML(domain string, results []links.Link) {
	urlSet := new(urlSet)
	urlSet.Xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
	for _, link := range getSameDomainLinks(domain, results) {
		urlSet.Urls = append(urlSet.Urls, url{Loc: link.Href})
	}

	output := xml.Header

	xmlResult, err := xml.MarshalIndent(urlSet, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	output += string(xmlResult)

	err = ioutil.WriteFile("output.xml", []byte(output), 0777)
	if err != nil {
		log.Fatal(err)
	}
}

//GetSameDomainLinks returns the links in an array that contain the given url
func getSameDomainLinks(url string, values []links.Link) []links.Link {
	var results []links.Link
	for _, val := range values {
		if val.Href == "/" {
			continue
		}

		if strings.HasPrefix(val.Href, "/") || strings.Contains(val.Href, url) {
			results = append(results, val)
		}
	}
	return results

}
