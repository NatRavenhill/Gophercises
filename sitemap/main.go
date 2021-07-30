package main

import (
	"flag"

	"smap"
)

func main() {
	url := flag.String("url", "http://calhoun.io", "Url to get sitemap for")
	result := smap.CollectLinks(*url)
	smap.WriteXML(*url, result)
}
