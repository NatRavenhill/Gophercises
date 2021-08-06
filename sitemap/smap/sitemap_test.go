package smap

import (
	"links"
	"reflect"
	"testing"
)

var (
	empty            = links.Link{"/", "default"}
	internal         = links.Link{"/courses", "courses"}
	expandedInternal = links.Link{"http://example.com/courses", "courses"}
	example          = links.Link{"http://example.com", "example"}
)

//TestFilterLinks tests that empty links are removed and internal links are fully quantified
func TestFilterLinks(t *testing.T) {
	baseUrl = "http://example.com"

	output := filterLinks([]links.Link{example, empty, internal})

	expectedOutput := []links.Link{example, expandedInternal}

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Fatalf("Expected %v but got %v", output, expectedOutput)
	}

}
