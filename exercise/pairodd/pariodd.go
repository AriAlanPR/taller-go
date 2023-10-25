package pairodd

// counts and returns the number of pair and odd numbers.
func Count(numbers []int) (pair, odd int) {
	pair = 0
	odd = 0

	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			pair++
		} else {
			odd++
		}
	}

	return pair, odd
}
