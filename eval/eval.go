package eval

import (
	"fmt"
	evalGit "github.com/apaxa-go/eval"
	"log"
)

func Eval() {
	var src string
	var err error
	fmt.Println("Enter expression without use space: 2+5")
	_, err = fmt.Scan(&src)
	log.Println("Expression: ", src)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		expr, err := evalGit.ParseString(src, "")
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			result, err := expr.EvalToInterface(nil)
			if err != nil {
				fmt.Println("Error: ", err)
				fmt.Println("Try again")
				Eval()
			} else {
				fmt.Printf("Result (type) \n    %v (%T)", result, result)
				log.Printf("Result (type) \n    %v (%T)", result, result)
			}
		}
	}
}
