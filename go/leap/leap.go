package leap

const testVersion = 3

// IsLeapYear Check if a given year is a leap year.
// A leap year is on every year that is evenly divisible by 4,
// except every year that is evenly divisible by 100
// unless the year is also evenly divisible by 400.
func IsLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
