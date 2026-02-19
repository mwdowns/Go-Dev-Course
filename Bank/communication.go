package main

import "fmt"

const welcomeMessage = "Welcome to Go Bank!"
const goodbyeMessage = "Goodbye"
const menuInputErrorMessage = "Please make sure to use and input between 1 and 4"
const depositPrompt = "Enter deposit amount: "
const depositErrorMessage = "the deposit amount must be greater than zero"
const withdrawalPrompt = "Enter withdrawal amount: "
const withdrawalErrorMessage = "the withdrawal amount must be greater than zero or less than your balance"

func showMenu() {
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Print("Your choice: ")
}
