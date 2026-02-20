package main

import (
	"errors"
	"fmt"

	"mwdowns.me/go-notes/note"
)

func getMenuInput(prompt string) (input int, err error) {
	fmt.Print(prompt)
	fmt.Scanln(&input)
	if input < 1 || input > 4 {
		return 0, errors.New(menuInputError)
	}
	return input, err
}

func executeInput(input int) {
	switch input {
	case 1:
		note.ShowNotes()
	case 2:
		note.CreateNewNote()
	default:
		fmt.Println("...coming soon...")
		return
	}
}
