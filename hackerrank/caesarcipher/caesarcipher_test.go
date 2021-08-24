package caesarcipher

import "testing"

func TestCaesarCipher(t *testing.T) {

	tests := []struct {
		input          string
		k              int32
		expectedOutput string
	}{
		{"Always-Look-on-the-Bright-Side-of-Life", 5, "Fqbfdx-Qttp-ts-ymj-Gwnlmy-Xnij-tk-Qnkj"},
		{"middle-Outz", 2, "okffng-Qwvb"},
		{"Hello_World!", 4, "Lipps_Asvph!"},
		{"aabc", 4, "eefg"},
		{"www.abc.xy", 87, "fff.jkl.gh"},
	}

	for _, test := range tests {
		result := caesarCipher(test.input, test.k)
		if result != test.expectedOutput {
			t.Fatalf("Got %s, expected %s", result, test.expectedOutput)
		}
	}

}
