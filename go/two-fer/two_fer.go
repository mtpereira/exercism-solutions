// Package twofer implements a library for printing messages about sharing things with someone.
package twofer

import "fmt"

// ShareWith returns a message for politely stating that you're sharing something with someone.
// It returns either "One for you, one for me" or "One for `name`, one for me", depending on the contents of `name`.
// If name is an empty string, it returns the first form, otherwise returns the second form.
func ShareWith(name string) string {
	person := name
	if person == "" {
		person = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", person)
}
