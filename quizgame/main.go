//package main reads in a quiz provided via a CSV file and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var score int

type Question = struct {
	q string
	a string
}

func main() {
	filename := flag.String("filename", "problems", "The name of the file where the questions are located")
	timelimit := flag.Int("timelimit", 30, "Time limit to complete the quiz in")
	shuffle := flag.Bool("shuffle", false, "Should we shuffle the questions?")
	flag.Parse()

	questions := parseQuestions(*filename)

	if *shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })
	}

	fmt.Println("Press enter to start the quiz")
	fmt.Scanln()

	doQuestions(questions, *timelimit)

	fmt.Printf("Game over! Your score was %d/%d", score, len(questions))

}

// parseQuestions parses a csv file into a questions array
func parseQuestions(filename string) []Question {
	file, err := os.Open(filename + ".csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)
	var questions []Question

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		questions = append(questions, Question{record[0], record[1]})

	}

	return questions
}

//doQuestion asks the user the quiz questions
func doQuestions(questions []Question, timelimit int) {
	timer := time.NewTimer(time.Duration(timelimit) * time.Second)
	answerChannel := make(chan string)

	for _, question := range questions {
		go func() {
			fmt.Println(question.q)
			var input string
			fmt.Scanln(&input)
			answerChannel <- input
		}()

		select {
		case <-timer.C:
			return
		case input := <-answerChannel:
			//check against answer
			if strings.TrimSpace(input) == question.a {
				score++
			}
		}
	}
}
