package main

import (
	"fmt"
)

const welcomeMessage = "Welcome to Go Notes!"
const menuPrompt = "Enter your choice: "
const menuInputError = "you must choose a valid input option"

func main() {
	// ..welcome
	fmt.Println(welcomeMessage)

	for {
		// show menu
		showMenu()
		input, err := getMenuInput(menuPrompt)

		if err != nil {
			fmt.Println(err)
		}

		// ..exit
		if input == 4 {
			fmt.Println("Goodbye!")
			return
		}

		executeInput(input)
		// ..show note
		// ..add note
		// ..remove note (stretch goal)

	}
}

func showMenu() {
	fmt.Println("1. Show Notes")
	fmt.Println("2. Create New Note")
	fmt.Println("3. Remove Note")
	fmt.Println("4. Exit")
}
