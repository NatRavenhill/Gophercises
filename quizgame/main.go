//package main reads in a quiz provided via a CSV file and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var score, total int

func main() {
	filename := flag.String("filename", "problems", "The name of the file where the questions are located")
	flag.Parse()

	file, err := os.Open(*filename + ".csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			fmt.Printf("Game over! Your score was %d/%d", score, total)
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		doQuestion(record)
	}

}

//doQuestion asks the user a question using the given record
func doQuestion(record []string) {

	fmt.Println(record[0])

	//wait for user input
	var input string
	fmt.Scanln(&input)

	//check against answer
	if input == record[1] {
		score++
	}

	total++
}
