package alphabet

import "unicode"

func IdentifyPosition(char rune) int {
	//convert to lower case for comparison
	char = unicode.ToLower(char)

	return int(char) - int('a') + 1
}
