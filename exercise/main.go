package main

import (
	"exercise/alphabet"
	"exercise/movefirst"
	"exercise/pairodd"
	"exercise/sorter"
	"exercise/tree"
	"exercise/uplowcase"
	"fmt"
)

var names = []string{"Dan", "Rigo", "Shawn", "Jenny"}

func main() {
	//sort array of strings
	sorter.Sort(names)
	sorter.Inverse(names)

	//print tree
	tree.PrintTree(7)

	//find upper and lower cases
	wordexample := "HelloWorld12"
	upper, lower, digits := uplowcase.FindUpperAndLowerCases(wordexample)

	fmt.Println("The number of upper cases is", upper)
	fmt.Println("The number of lower cases is", lower)
	fmt.Println("The number of digits is", digits)

	//find pair and odd
	pair, odd := pairodd.Count([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println("The number of pair is", pair)
	fmt.Println("The number of odd is", odd)

	//move first letter to the end
	fmt.Println(movefirst.ToEnd("Hello"))

	//identify position of a character in the alphabet
	char := 'W'
	fmt.Printf("The position of %c in the alphabet is %d\n", char, alphabet.IdentifyPosition(char))
}
