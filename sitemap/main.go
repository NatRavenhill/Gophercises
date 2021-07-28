package main

import (
	"fmt"
	"links"
	"log"
	"net/http"
)

func main() {
	//get the webpage
	resp, err := http.Get("http://calhoun.io")
	if err != nil {
		log.Fatal(err)
	}

	links := links.ParseContent(resp.Body)
	for _, link := range links {
		fmt.Println(link)
	}

}
