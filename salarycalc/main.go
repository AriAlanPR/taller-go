package main

import (
	b "salarycalc/basic"
	"salarycalc/basic/gross"

	"fmt"
)

func main() {
	b.Basic = 10000
	fmt.Println(gross.GrossSalary())
}
