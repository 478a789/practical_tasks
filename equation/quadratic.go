package equation

import (
	"fmt"
	"practical_tasks/discriminant"
)

func Quadratic() {
	var a, b, c float64
	fmt.Println("Enter a, b, c separated by a space")
	fmt.Scan(&a, &b, &c)
	discriminant.Discriminant(a, b, c)
}
