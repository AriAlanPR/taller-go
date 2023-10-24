package main

import (
	b "salarycalc/basic"
	"salarycalc/basic/gross"

	"fmt"
	"strconv"
)

type float float64
func (fl float) toString() string {
	return strconv.FormatFloat(float64(fl), 'f', -1, 64)
}

func main() {
	b.Basic = 10000
	f := float(b.Basic)
	str := f.toString()
	fmt.Println(gross.GrossSalary())
	fmt.Println("Basic salary: " + str)
}
