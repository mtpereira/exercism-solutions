// Package bob implements the behaviour of a lackadaisical teenager.
package bob

import "strings"

// Hey returns the appropriate answer to a conversational partner from a lackadaisical teenager.
// The conversational partner is expected to follow normal sentence punctuation rules.
func Hey(remark string) string {
	r := strings.TrimSpace(remark)
	switch {
	case r == "":
		return "Fine. Be that way!"
	case isQuestion(r) && isScream(r):
		return "Calm down, I know what I'm doing!"
	case isScream(r):
		return "Whoa, chill out!"
	case isQuestion(r):
		return "Sure."
	}
	return "Whatever."
}

func isScream(remark string) bool {
	return strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark
}

func isQuestion(remark string) bool {
	return string(remark[len(remark)-1]) == "?"
}
