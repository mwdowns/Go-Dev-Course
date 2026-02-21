package menu

import (
	"errors"
	"fmt"

	"mwdowns.me/go-notes-and-todos/interfaces"
	"mwdowns.me/go-notes-and-todos/note"
	"mwdowns.me/go-notes-and-todos/todo"
)

const welcomeMessage = "Welcome to Go Notes and Todos!"
const menuPrompt = "Enter your choice: "
const menuInputError = "you must choose a valid input option"

func DisplayMenu() {
	// welcome
	fmt.Println(welcomeMessage)

	for {
		// show menu
		showMenu()
		input, err := getMenuInput(menuPrompt)

		if err != nil {
			fmt.Println(err)
		}
		// ..exit
		if input == 6 {
			fmt.Println("Goodbye!")
			return
		}
		executeInput(input)
	}
}

func showMenu() {
	fmt.Println("1. Show Notes")
	fmt.Println("2. Create New note")
	fmt.Println("3. Remove note")
	fmt.Println("4. Show Todos")
	fmt.Println("5. Create Todo")
	fmt.Println("6. Exit")
}

func getMenuInput(prompt string) (input int, err error) {
	fmt.Print(prompt)
	fmt.Scanln(&input)
	if input < 1 || input > 6 {
		return 0, errors.New(menuInputError)
	}
	return input, err
}

func executeInput(input int) {
	switch input {
	case 1:
		note.DisplayNotes()
	case 2:
		n, err := note.CreateNewNote()
		saveObject(n, err)
	case 3:
		fmt.Println("...coming soon...")
	case 4:
		todo.DisplayTodos()
	case 5:
		td, err := todo.CreateNewTodo()
		saveObject(td, err)
	default:
		fmt.Println("wut?")
	}
}

func saveObject(object interfaces.SaveAndDisplay, err error) {
	if err != nil {
		fmt.Println(err)
	}
	err = object.Save()
	if err != nil {
		fmt.Println(err)
	}
	object.SuccessMessage()
}
