package main

import (
	"fmt"
	"log"
	"os"
	"practical_tasks/calculation"
	"practical_tasks/equation"
	"practical_tasks/eval"
	"practical_tasks/hello"
)

/*
	This is a collection of practical tasks in Golang
*/
func main() {
	file, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(": ", err)
	}
	defer func() {
		fmt.Printf("Close log file \n")
		err := file.Close()
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}()
	log.SetOutput(file)
	log.Println("Start program")
	Choice()

}

func Choice() {
	switch hello.Hello() {
	case 0:
		equation.Quadratic()
	case 1:
		calculation.Calculation()
	case 2:
		eval.Eval()
	case 3:
		fmt.Println("В разработке, выберите другой вариант")
		Choice()
	default:
		fmt.Println("Retype, please")
		Choice()
	}
}
