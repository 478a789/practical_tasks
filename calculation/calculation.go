package calculation

import (
	"fmt"
	"math"
)

func Calculation() {
	var amount, percent, years, total, days, period float64
	days = 365
	fmt.Println("Enter 'initial amount', 'percent', 'the number of years' " +
		"and the 'bank period in days for interest accrual' separated by a space")
	fmt.Scan(&amount, &percent, &years, &period)
	total = amount * math.Pow(1+((percent*period)/(100*days)), years*days/period)
	fmt.Println("Total amount: ", total)
}
