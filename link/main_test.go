package main

import (
	"testing"
)

//Tests extracting links works correctly for the given files
func TestExtractLinks(t *testing.T) {
	var tests = []struct {
		filename       string
		expectedResult []Link
	}{
		{"ex1.html", []Link{{"/other-page", "A link to another page"}}},
		{"ex2.html", []Link{{"https://www.twitter.com/joncalhoun", "Check me out on twitter"}, {"https://github.com/gophercises", "Gophercises is on Github!"}}},
		{"ex3.html", []Link{{"#", "Login"}, {"/lost", "Lost? Need help?"}, {"https://twitter.com/marcusolsson", "@marcusolsson"}}},
		{"ex4.html", []Link{{"/dog-cat", "dog cat"}}},
	}

	for _, test := range tests {
		Links = []Link{}
		tree := parseFile(test.filename)
		extractLinks(tree)

		for i := 0; i < len(Links); i++ {
			if Links[i] != test.expectedResult[i] {
				t.Fatalf("Got %v, expected %v", Links[i], test.expectedResult[i])
			}
		}

	}
}
