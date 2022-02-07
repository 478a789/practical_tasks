package discriminant

import (
	"fmt"
	"math"
)

func Discriminant(a, b, c float64) {
	var x1, x2 float64
	d := b*b - 4*a*c
	if d < 0 {
		fmt.Println("No design")
	} else if d == 0 {
		x1 = -b / (2 * a)
		fmt.Println(x1)
	} else if d > 0 {
		x1 = (-b + math.Sqrt(d)) / (2 * a)
		x2 = (-b - math.Sqrt(d)) / (2 * a)
		fmt.Println(x1, x2)
	}
}
