// Package acronym implements methods for converting phrases to acronyms.
package acronym

import "strings"
import "regexp"

// Abbreviate returns a three letter acronym from a given three word name.
func Abbreviate(s string) string {
	reg, err := regexp.Compile("[_-]")
	if err != nil {
		panic(err)
	}

	s = reg.ReplaceAllString(s, " ")
	words := strings.Fields(s)
	var result string
	for _, w := range words {
		result += w[0:1]
	}
	return strings.ToUpper(result)
}
