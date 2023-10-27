package main

import (
	"exercise/uplowcase"
	"fmt"
)

func main() {
	//find upper and lower cases
	wordexample := "HelloWorld12133bbb"
	isvalid := uplowcase.IsValidPassword(wordexample)
	a := "HelloWorld" == ""

	if isvalid {
		fmt.Printf("The password is valid %t\n", a)
	} else {
		fmt.Println("The password is not valid")
	}

	// //sort array of strings
	// sorter.Sort(names)
	// sorter.Inverse(names)

	// //print tree
	// tree.PrintTree(7)

	// //find pair and odd
	// pair, odd := pairodd.Count([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// fmt.Println("The number of pair is", pair)
	// fmt.Println("The number of odd is", odd)

	// //move first letter to the end
	// fmt.Println(movefirst.ToEnd("Hello"))

	// //identify position of a character in the alphabet
	// char := 'W'
	// fmt.Printf("The position of %c in the alphabet is %d\n", char, alphabet.IdentifyPosition(char))
}
