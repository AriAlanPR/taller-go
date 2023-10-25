package uplowcase

import (
	"regexp"
)

// finds and returns the number of uppercase, lowercase and digits in a string.
func FindUpperAndLowerCases(name string) (upper, lower, digits int) {
	upper = 0
	lower = 0
	digits = 0

	//another way of regexp is regexp.MatchString(pattern, string)
	// if matched, _ := regexp.MatchString(uppercase, string(char)); matched {
	// 	upper++
	// } else if matched, _ := regexp.MatchString(lowercase, string(char)); matched {
	// 	lower++
	// } else if matched, _ := regexp.MatchString(digit, string(char)); matched {
	// 	digits++
	// }
	//other way is using package unicode and function IsLower and isUpper
	//example: unicode.IsLower(char)

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
		}

	}
	return upper, lower, digits
}
