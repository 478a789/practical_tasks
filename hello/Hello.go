package hello

import (
	"fmt"
	"log"
)

func Hello() int {
	var choice int
	type Description struct {
		Names string
	}

	var listOfPrograms = map[int]Description{
		0: {"Solving the quadratic equation"},
		1: {"Calculation of interest on the deposit"},
		2: {"Calculator, variant 1"},
		3: {"Calculator, variant 2"},
	}

	fmt.Printf("Good day! \nThis is a collection of practical tasks in Golang \n")

	for i := 0; i < len(listOfPrograms); i++ {
		fmt.Println(i, " ", listOfPrograms[i])
	}
	fmt.Printf("Specify program number and press Enter: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println("Try again")
		Hello()
	} else {
		fmt.Println(listOfPrograms[choice])
		log.Println("Choice is: ", listOfPrograms[choice])
		return choice
	}
	return 999
}
