package gross

import (
	b "salarycalc/basic"
)

func GrossSalary() int {
	a, t := b.Calculation()
	return ((b.Basic + a) - t)
}
