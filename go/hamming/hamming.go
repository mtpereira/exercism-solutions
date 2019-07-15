// Package hamming implements the Hamming difference calculation between two DNA strands.
package hamming

import "errors"

// Distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("sequences have different lengths")
	}

	count := 0
	for i, n := range a {
		if n != rune(b[i]) {
			count++
		}
	}
	return count, nil
}
