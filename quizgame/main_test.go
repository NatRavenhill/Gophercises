package main

import (
	"testing"
)

// Tests that parseQuestions gets all questions in the file
func TestParseQuestions(t *testing.T) {
	questions := parseQuestions("test")

	if len(questions) != 2 {
		t.Fatal("Expected a single question in the file")

	}
}

//Tests that DoQuestions completes after the time limit
func TestDoQuestions(t *testing.T) {
	questions := parseQuestions("test")
	doQuestions(questions, 3)
}
