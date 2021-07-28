package main

import (
	"flag"
	"links"
	"log"
	"net/http"

	"smap"
)

func main() {
	url := flag.String("url", "calhoun.io", "Url to get sitemap for")

	resp, err := http.Get("http://" + *url)
	if err != nil {
		log.Fatal(err)
	}

	result := links.ParseContent(resp.Body)
	smap.WriteXML(*url, result)
}
