package main

import (
	"log"
	"story"
)

func main() {
	chapters, err := story.DecodeJSON("gopher.json")
	if err != nil {
		log.Fatal(err)
	}

	story.StartGame(chapters)
}
