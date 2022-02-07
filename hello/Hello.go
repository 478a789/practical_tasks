package hello

import "fmt"

func Hello() int {
	var choice int
	type Description struct {
		Names string
	}

	var listOfPrograms = map[int]Description{
		0: {"Solving the quadratic equation"},
		1: {"Calculation of interest on the deposit"},
	}

	fmt.Printf("Good day! \nThis is a collection of practical tasks in Golang \n")

	for key, name := range listOfPrograms {
		fmt.Println(key, name)
	}
	fmt.Printf("Specify program number and press Enter: ")
	fmt.Scan(&choice)
	fmt.Println(listOfPrograms[choice])
	return choice
}
