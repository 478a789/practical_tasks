package equation

import (
	"fmt"
	"log"
	"practical_tasks/discriminant"
)

func Quadratic() {
	var a, b, c float64
	fmt.Println("Enter a, b, c separated by a space")
	_, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println("Try again")
		Quadratic()
	} else {
		fmt.Printf("Equation %.2fx^2+%.2fx+%.2f\n", a, b, c)
		log.Printf("Equation %.2fx^2+%.2fx+%.2f\n", a, b, c)
		discriminant.Discriminant(a, b, c)
	}
}
