package sorter

import (
	"fmt"
	"sort"
)

// sort an array of strings alphabetically and print it.
func Sort(names []string) {
	sort.Strings(names)
	fmt.Println(names)
}

func Inverse(names []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(names)))
	fmt.Println(names)
}
