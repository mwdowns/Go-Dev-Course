package main

import (
	"errors"
	"fmt"

	"mwdowns.me/bank/fileops"
)

func getUserInput() (input int) {
	fmt.Scanf("%d", &input)
	return input
}

func executeInput(input int, balance float64) float64 {
	switch input {
	case 1:
		fmt.Printf("Your balance is %.2f\n", balance)
	case 2:
		depAm, err := getDepositAmount()
		if err != nil {
			fmt.Println(err)
			break
		}
		balance += depAm
		fileops.WriteFloatToFile(accountBalanceFile, balance)
		fmt.Printf("Your new balance is %.2f\n", balance)
	case 3:
		withAm, err := getWithdrawalAmount(balance)
		if err != nil {
			fmt.Println(err)
			break
		}
		balance -= withAm
		fileops.WriteFloatToFile(accountBalanceFile, balance)
		fmt.Printf("Your new balance is %.2f\n", balance)
	default:
		fmt.Println(menuInputErrorMessage)
	}
	return balance
}

func getDepositAmount() (depAm float64, error error) {
	fmt.Print(depositPrompt)
	fmt.Scan(&depAm)
	if depAm <= 0 {
		return 0.0, errors.New(depositErrorMessage)
	}
	return depAm, nil
}

func getWithdrawalAmount(balance float64) (withAm float64, error error) {
	fmt.Print(withdrawalPrompt)
	fmt.Scan(&withAm)
	if withAm <= 0 || withAm > balance {
		return 0.0, errors.New(withdrawalErrorMessage)
	}
	return withAm, nil
}
