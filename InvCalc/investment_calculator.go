package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Investment Rate: ")
	fmt.Scan(&expectedReturnRate)

	printLine(calcFutureValues(investmentAmount, expectedReturnRate, years))
}

func calcFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, iv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	iv = fv / math.Pow(1+inflationRate/100, years)

	return fv, iv
}

func printLine(futureValue float64, inflationValue float64) {
	answerString := fmt.Sprintf(
		"Future Value: %.2f\nFuture Value (adjusted for inflation): %.2f",
		futureValue, inflationValue)

	fmt.Println("...")
	fmt.Println(answerString)
}
