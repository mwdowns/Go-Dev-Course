package main

import (
	"fmt"

	"mwdowns.me/structs/user"
)

func main() {
	u, err := user.New(
		getUserData("Please enter your first name: "),
		getUserData("Please enter your last name: "),
		getUserData("Please enter your birth date (MM/DD/YYYY): "),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	a, aErr := user.NewAdmin("test@test.com", "test")

	if aErr != nil {
		fmt.Println(aErr)
		return
	}

	u.ShowUserData()
	a.ShowAdminData()
	a.ClearUserName()
	a.ShowAdminData()
}

func getUserData(promptText string) (input string) {
	fmt.Println(promptText)
	fmt.Scanln(&input)
	return input
}
