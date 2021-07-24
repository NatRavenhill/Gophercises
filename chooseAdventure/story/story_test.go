package story

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestDecodeJSON tests that the file is parsed into the expected number of chapters
func TestDecodeJSON(t *testing.T) {
	var tests = []struct {
		filename      string
		shouldError   bool
		totalChapters int
	}{
		{"./../gopher.json", false, 7},
		{"./../badFile.json", true, 0},
	}

	for _, test := range tests {
		chapters, err := DecodeJSON(test.filename)

		if (err != nil) != test.shouldError {
			t.Fatalf("Got %v, expected %v", err != nil, test.shouldError)
		}

		if len(chapters) != test.totalChapters {
			t.Fatalf("Got %d, expected %d", len(chapters), test.totalChapters)
		}
	}
}

//TestBuildTemplate tests that the built template gives the correct output for each path
func TestBuildTemplate(t *testing.T) {
	chapters, err := DecodeJSON("./../gopher.json")
	if err != nil {
		t.Fatal(err)
	}

	var tests = []struct {
		path            string
		expectedContent string
	}{
		{"/", "The Little Blue Gopher"},
		{"/new-york", "Visiting New York"},
		{"/ffgf", "Chapter not found!"},
	}

	for _, test := range tests {
		req, _ := http.NewRequest("GET", test.path, nil)
		rr := httptest.NewRecorder()
		template := buildTemplate(chapters, "./../story.html")
		handler := http.HandlerFunc(template)
		handler.ServeHTTP(rr, req)

		if !strings.Contains(rr.Body.String(), test.expectedContent) {
			t.Fail()
		}
	}

}
