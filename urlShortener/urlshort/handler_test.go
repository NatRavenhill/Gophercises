package urlshort

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var pathsToUrls = make(map[string]string)

func init() {
	pathsToUrls["/google"] = "google.com"

}

func makeRequest(path string, testFunc http.HandlerFunc) int {
	req, _ := http.NewRequest("GET", path, nil)
	req.RequestURI = path
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testFunc)
	handler.ServeHTTP(rr, req)

	return rr.Code

}

//TestMap Handler tests the map handler method redirects
func TestMapHandler(t *testing.T) {
	var tests = []struct {
		path   string
		status int
	}{
		{"/google", http.StatusFound},
		{"/abc", http.StatusOK},
	}

	for _, test := range tests {
		testFunc := MapHandler(pathsToUrls, nil)

		code := makeRequest(test.path, testFunc)
		if code != test.status {
			t.Fatalf("Got %v, expected %v", code, test.status)
		}
	}

}

//TestYAMLHandler tests the YAML handler method redirects
func TestYAMLHandler(t *testing.T) {
	var tests = []struct {
		path   string
		status int
	}{
		{"/urlshort", http.StatusFound},
		{"/abc", http.StatusOK},
	}

	for _, test := range tests {
		content, err := ioutil.ReadFile("./../paths.yaml")
		if err != nil {
			t.Fatal(err)
		}

		testfunc, err := YAMLHandler(content, nil)
		if err != nil {
			t.Fatal(err)
		}

		code := makeRequest(test.path, testfunc)
		if code != test.status {
			t.Fatalf("Got %v, expected %v", code, test.status)
		}

	}

}

//TestJSONHandler tests the JSON handler method redirects
func TestJSONHandler(t *testing.T) {
	var tests = []struct {
		path   string
		status int
	}{
		{"/quiz", http.StatusFound},
		{"/abc", http.StatusOK},
	}
	for _, test := range tests {
		content, err := ioutil.ReadFile("./../paths.json")
		if err != nil {
			t.Fatal(err)
		}

		testfunc, err := JSONHandler(content, nil)
		if err != nil {
			t.Fatal(err)
		}

		code := makeRequest(test.path, testfunc)
		if code != test.status {
			t.Fatalf("Got %v, expected %v", code, test.status)
		}

	}

}
