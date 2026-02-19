package main

import (
	"errors"
	"fmt"
	"os"
)

const revenuePrompt = "How much did you make this year? "
const expensesPrompt = "How much did you spend this year? "
const taxRatePrompt = "What is the tax rate? "

// Goals
// 1. Validate user input
//    - show error and exit if invalid input
//    - no negative numbers
//    . no 0
// 2. store calculated results into file

func main() {
	revenue, rErr := getUserInput(revenuePrompt)
	expenses, eErr := getUserInput(expensesPrompt)
	taxRate, tErr := getUserInput(taxRatePrompt)

	if rErr != nil || eErr != nil || tErr != nil {
		fmt.Println("Only positive numbers allowed")
		return
	}

	printOutput(calcFinancials(revenue, expenses, taxRate))
}

func getUserInput(prompt string) (input float64, err error) {
	fmt.Print(prompt)
	fmt.Scan(&input)
	if input > 0 {
		return input, nil
	}
	return 0.0, errors.New("invalid input")
}

func calcFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}

func printOutput(ebt, profit, ratio float64) {
	fmt.Println("...")
	fmt.Printf("Earnings before tax: %.2f\n", ebt)
	fmt.Printf("Earnings after tax: %.2f\n", profit)
	fmt.Printf("Earnings ratio: %.4f\n", ratio)
	textString := []byte(fmt.Sprintf("%.2f, %.2f, %.4f", ebt, profit, ratio))
	os.WriteFile("profits.txt", textString, 0644)
}
