package smap

import (
	"io/ioutil"
	"links"
	"reflect"
	"strings"
	"testing"
)

var (
	empty            = links.Link{"/", "default"}
	internal         = links.Link{"/courses", "courses"}
	expandedInternal = links.Link{"http://example.com/courses", "courses"}
	example          = links.Link{"http://example.com", "example"}
)

//Tests that DoURLs behaves as expected
func TestDoUrls(t *testing.T) {
	results := DoURLs("http://example.com", 3)

	if len(results) != 0 {
		t.Fatal("Exepected no links to be found in example.com")
	}
}

//Tests that CollectLinks behaves as expected
func TestCollectLinks(t *testing.T) {
	baseUrl = "http://example.com"
	maxDepth = 3
	results := make(map[string]bool)
	CollectLinks("http://example.com", results, 3)

	if len(results) != 0 {
		t.Fatal("Exepected no links to be found in example.com")
	}
}

//TestWriteXML tests that the links in the results map are written to an XML file
func TestWriteXML(t *testing.T) {
	baseUrl = "http://example.com"
	results := make(map[string]bool)
	results["example.com"] = true
	WriteXML(results)

	//check file exists
	contents, err := ioutil.ReadFile("output.xml")

	if err != nil {
		t.Fatal(err)
	}

	if len(contents) == 0 {
		t.Fatal("Expected a file to be created")
	}

	strCont := string(contents)
	if !strings.Contains(strCont, "example.com") {
		t.Fatal("Expected example.com to be written to XML")
	}
}

//TestFilterLinks tests that empty links are removed and internal links are fully quantified
func TestFilterLinks(t *testing.T) {
	baseUrl = "http://example.com"

	output := filterLinks([]links.Link{example, empty, internal})

	expectedOutput := []links.Link{example, expandedInternal}

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Fatalf("Expected %v but got %v", output, expectedOutput)
	}

}
