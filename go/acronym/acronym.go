// Package acronym implements methods for converting phrases to acronyms.
package acronym

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

// Abbreviate returns a three letter acronym from a given three word name.
func Abbreviate(s string) string {
	lex := &parser{reader: strings.NewReader(s)}
	result, err := read(lex)
	if err != nil {
		fmt.Printf("Error on read: %v\n", err)
		return ""
	}
	return result
}

type parser struct {
	reader *strings.Reader
	token  rune
	added  bool
}

func (lex *parser) next() error {
	var err error
	lex.token, _, err = lex.reader.ReadRune()
	return err
}

func (lex *parser) text() string {
	return string(lex.token)
}

func read(lex *parser) (string, error) {
	var builder strings.Builder
	for err := lex.next(); err != io.EOF; err = lex.next() {
		if err != nil {
			return "", err
		}
		switch {
		case isSeparator(lex.token):
			lex.added = false
		case unicode.IsLetter(lex.token):
			if !lex.added {
				builder.WriteString(strings.ToUpper(lex.text()))
				lex.added = true
			}
		}
	}
	return builder.String(), nil
}

func isSeparator(token rune) bool {
	return unicode.IsSpace(token) || (unicode.IsPunct(token) && token != '\'')
}
