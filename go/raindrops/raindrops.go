// Package raindrops implements ways to reproduce raindrop sounds.
package raindrops

import (
	"strconv"
	"strings"
)

// Convert a number into a string representing a raindrop sound.
func Convert(number int) string {
	var result strings.Builder

	for factor := 1; factor <= number; factor++ {
		if number%factor == 0 {
			switch factor {
			case 3:
				result.WriteString("Pling")
			case 5:
				result.WriteString("Plang")
			case 7:
				result.WriteString("Plong")
			}
		}
	}

	if result.Len() == 0 {
		result.WriteString(strconv.Itoa(number))
	}

	return result.String()
}
