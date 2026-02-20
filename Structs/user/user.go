package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) (*Admin, error) {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "----",
			createdAt: time.Now(),
		},
	}, nil
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("must enter information for values")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func (u *User) ShowUserData() {
	fmt.Printf("First Name: %s\n", u.firstName)
	fmt.Printf("Last Name: %s\n", u.lastName)
	fmt.Printf("Birth Date: %s\n", u.birthDate)
	fmt.Printf("Created: %s\n", u.createdAt.Format(time.RFC3339))
}

func (a *Admin) ShowAdminData() {
	fmt.Printf("First Name: %s\n", a.firstName)
	fmt.Printf("Last Name: %s\n", a.lastName)
	fmt.Printf("Birth Date: %s\n", a.birthDate)
	fmt.Printf("Created: %s\n", a.createdAt.Format(time.RFC3339))
	fmt.Printf("First Name: %s\n", a.password)
	fmt.Printf("Last Name: %s\n", a.email)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
