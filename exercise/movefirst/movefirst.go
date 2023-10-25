package movefirst

//moves the first leter of a string to the end
func ToEnd(word string) string {
	var first, changed string
	for i, char := range word {
		if i == 0 {
			first = string(char)
		} else {
			changed += string(char)
		}
	}

	return changed + first
}
