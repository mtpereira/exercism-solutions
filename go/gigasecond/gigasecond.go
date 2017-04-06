package gigasecond

import "time"

const (
	testVersion = 4
	gigasecond  = time.Duration(1000000000 * time.Second)
)

// AddGigasecond  Calculate the moment when someone has lived for 10^9 seconds.
// A gigasecond is 10^9 (1,000,000,000) seconds.
func AddGigasecond(initial time.Time) time.Time {
	return initial.Add(gigasecond)
}
