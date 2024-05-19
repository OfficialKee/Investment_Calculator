package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate = 5.5
	var years float64

	fmt.Println("What is your investment amount? Enter here: ")
	fmt.Scan(&investmentAmount)
	fmt.Println(("How many years will you be investing? Enter here: "))
	fmt.Scan(&years)
	fmt.Println("What is expected return rate? Enter here: ")
	fmt.Scan(&expectedReturnRate)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
