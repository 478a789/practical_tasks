package calculation

import (
	"fmt"
	"log"
	"math"
)

func Calculation() {
	var amount, percent, years, total, days, period float64
	days = 365
	fmt.Printf("        Enter 'initial amount', 'percent', 'the number of years' " +
		"and \n   the 'bank period in days for interest accrual' separated by a space." +
		"\n   for example: 50000 8 15 30.5 \n")

	_, err := fmt.Scan(&amount, &percent, &years, &period)
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println("Try again")
		Calculation()
	} else {
		total = amount * math.Pow(1+((percent*period)/(100*days)), years*days/period)
		fmt.Println("Total amount: ", total)
		log.Println("Total amount: ", total)
	}
}
