// Package link parses an HTML file and extracts all of its links into a data structure
package main

import (
	"fmt"
	"log"
	"os"

	"links"
)

func main() {
	content, err := os.Open("ex1.html")
	if err != nil {
		log.Fatal(err)
	}
	links := links.ParseContent(content)

	for _, val := range links {
		fmt.Println(val)
	}
}
