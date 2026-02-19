package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	var u user
	u.firstName = getUserData("Please enter your first name: ")
	u.lastName = getUserData("Please enter your last name: ")
	u.birthDate = getUserData("Please enter your birth date (MM/DD/YYYY): ")
	u.createdAt = time.Now()

	showUserData(u)
}

func showUserData(user user) {
	fmt.Printf("First Name: %s\n", user.firstName)
	fmt.Printf("Last Name: %s\n", user.lastName)
	fmt.Printf("Birth Date: %s\n", user.birthDate)
	fmt.Printf("Created: %s\n", user.createdAt.Format(time.RFC3339))
}

func getUserData(promptText string) (input string) {
	fmt.Println(promptText)
	fmt.Scan(&input)
	return input
}
