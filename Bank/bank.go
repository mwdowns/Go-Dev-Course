package main

import (
	"fmt"

	"mwdowns.me/bank/fileops"
)

const accountBalanceFile = "balances.txt"

func main() {
	balance, err := fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil && err.Error() == "could not parse balance" {
		fmt.Println(err)
		return
	}

	fmt.Println(welcomeMessage)
	for {
		showMenu()
		input := getUserInput()
		if input != 4 {
			newBalance := executeInput(input, balance)
			balance = newBalance
		} else {
			fmt.Println(goodbyeMessage)
			break
		}
	}
}
