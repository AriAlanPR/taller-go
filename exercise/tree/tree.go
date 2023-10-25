package tree

import "fmt"

// make a fancy tree with the inputted level and print it.
func PrintTree(level int) {
	for i := 0; i < level; i++ {
		for j := 0; j < level-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
