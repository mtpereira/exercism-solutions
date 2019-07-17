// Package scale implements musical scale generators.
package scale

var sharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

// Scale returns a musical scale starting at the tonic with the given interval.
func Scale(tonic string, interval string) []string {
	var scale []string
	switch tonic {
	case "C":
		scale = sharps
	case "F":
		scale = flats
	}

	for i, t := range scale {
		if t == tonic {
			return append(scale[i:len(sharps)], scale[0:i]...)
		}
	}
	return []string{}
}
