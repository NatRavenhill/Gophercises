package camelcase

import (
	"testing"
)

func TestCamelCase(t *testing.T) {
	var tests = []struct {
		input          string
		expectedOutput int
	}{
		{"", 0},
		{"saveChangesInTheEditor", 5},
		{"SaveChangesInTheEditor", 0},
		{"saveChangesInTheEditor5753753", 0},
		{"save Changes In The Editor", 0},
	}

	for _, test := range tests {
		result := Camelcase(test.input)
		if result != int32(test.expectedOutput) {
			t.Fatalf("Got %d, expected %d", result, test.expectedOutput)
		}
	}
}
