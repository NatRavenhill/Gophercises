package links

import (
	"log"
	"os"
	"testing"

	"golang.org/x/net/html/atom"
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
		content, err := os.Open("../" + test.filename)
		if err != nil {
			log.Fatal(err)
		}

		links := ParseContent(content)

		for i := 0; i < len(links); i++ {
			if links[i] != test.expectedResult[i] {
				t.Fatalf("Got %v, expected %v", links[i], test.expectedResult[i])
			}
		}

	}
}

//TestContainsTag tests that containsTag only returns true for valid tags
func TestContainsTag(t *testing.T) {
	var tests = []struct {
		tag            atom.Atom
		expectedResult bool
	}{
		{atom.Span, true},
		{atom.Class, false},
	}

	for _, test := range tests {
		result := containsTag(test.tag)

		if result != test.expectedResult {
			t.Fatalf("Got %v, expected %v", result, test.expectedResult)
		}
	}
}
