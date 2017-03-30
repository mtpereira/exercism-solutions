package clock

import "fmt"

const testVersion = 4
const minutesInDay = 1440

// Clock Struct that stores hours and minutes as int.
type Clock struct {
	minutes int
}

// New Create a new clock with a starting time defined by hour and minute.
func New(hour, minute int) Clock {
	totalMinutes := minute + hour*60
	normalizeMinutes(&totalMinutes)
	return Clock{minutes: totalMinutes}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", (c.minutes/60)%24, c.minutes%60)
}

// Add Advances clock by the amount defined in minutes.
func (c Clock) Add(minutes int) Clock {
	c.minutes += minutes
	normalizeMinutes(&c.minutes)
	return c
}

// normalizeMinutes Add the total amount of minutes in a
// day if the total is a negative number. This makes it easier to print and
// compare dates.
func normalizeMinutes(minutes *int) {
	for *minutes < 0 {
		*minutes += minutesInDay
	}
	for *minutes >= minutesInDay {
		*minutes -= minutesInDay
	}
}
