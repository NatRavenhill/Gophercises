package camelcase

import (
	"regexp"
)

//Camelcase determine the number of words in a camel case string, or returns 0 if not in camel case.
func Camelcase(s string) int32 {
	camelCaseRegex := regexp.MustCompile("^[a-z][a-zA-Z]+$")
	match := camelCaseRegex.MatchString(s)
	if match {
		upperCaseRegex := regexp.MustCompile("[A-Z]")
		matchCount := len(upperCaseRegex.FindAllString(s, -1)) + 1
		return int32(matchCount)

	} else {
		return 0
	}
}
