// Package raindrops implements ways to reproduce raindrop sounds.
package raindrops

import (
	"strconv"
)

// Convert a number into a string representing a raindrop sound.
func Convert(number int) string {
	var result string

	for _, factor := range []int{3, 5, 7} {
		if number%factor == 0 {
			switch factor {
			case 3:
				result += "Pling"
			case 5:
				result += "Plang"
			case 7:
				result += "Plong"
			}
		}
	}

	if result == "" {
		result = strconv.Itoa(number)
	}

	return result
}
