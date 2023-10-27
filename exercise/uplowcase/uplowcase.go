package uplowcase

import (
	"regexp"
)

// finds and returns the number of uppercase, lowercase and digits in a string.
func FindUpperAndLowerCases(name string) (upper, lower, digits, other int) {
	upper = 0
	lower = 0
	digits = 0
	other = 0

	uppercase := regexp.MustCompile(`^[A-Z]$`)
	lowercase := regexp.MustCompile(`^[a-z]$`)
	digit := regexp.MustCompile(`^[0-9]+$`)

	for _, char := range name {
		switch {
		case uppercase.MatchString(string(char)):
			upper++
		case lowercase.MatchString(string(char)):
			lower++
		case digit.MatchString(string(char)):
			digits++
		default:
			other++
		}

	}
	return upper, lower, digits, other
}

func IsValidPassword(name string) bool {
	upper, lower, digits, other := FindUpperAndLowerCases(name)

	return upper >= 1 && lower >= 1 && digits >= 1 && other == 0
}
