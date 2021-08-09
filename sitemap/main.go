package main

import (
	"flag"

	"smap"
)

func main() {
	url := flag.String("url", "https://calhoun.io", "Url to get sitemap for")
	maxDepth := flag.Int("maxDepth", 3, "Maximum depth of the search")
	flag.Parse()

	result := smap.DoURLs(*url, *maxDepth)
	smap.WriteXML(result)
}
