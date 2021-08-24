package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"caesarCipher"
	"camelcase"
)

func main() {
	problem := flag.String("problem", "camelCase", "Name of problem to test")
	flag.Parse()

	switch *problem {
	case "camelCase":
		if len(os.Args) < 2 {
			log.Fatal("Expected a string argument")
		}
		s := strings.TrimRight(os.Args[len(os.Args)-1], "\r\n")

		result := camelcase.Camelcase(s)

		fmt.Printf("%d\n", result)
	case "caesarCipher":
		caesarCipher.ReadCaesarFile()
	}

}
