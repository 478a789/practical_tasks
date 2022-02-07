package main

import (
	"fmt"
	"practical_tasks/calculation"
	"practical_tasks/equation"
	"practical_tasks/hello"
)

/*
	This is a collection of practical tasks in Golang
*/
func main() {

	switch hello.Hello() {
	case 0:
		equation.Quadratic()
	case 1:
		calculation.Calculation()
	default:
		fmt.Println(00)
	}

}
