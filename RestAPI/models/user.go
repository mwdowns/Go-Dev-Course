package models

import (
	"errors"
	db "mwdowns/rest-api/DB"
	"mwdowns/rest-api/utils"

	"github.com/google/uuid"
)

type User struct {
	Id        int64     `json:"id"`
	Email     string    `binding:"required" json:"email"`
	Password  string    `binding:"required" json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Uuid      uuid.UUID `json:"uuid"`
}

const saveUserQuery = "INSERT INTO users(email, password, first_name, last_name, uuid) VALUES (?, ?, ?, ?, ?)"
const getUserQuery = "SELECT id, password FROM users WHERE email = ?"
const loginErrorMessage = "credentials invalid"

func (u User) Save() (string, error) {
	userUuid := uuid.New()
	u.Uuid = userUuid
	stmt, err := db.DB.Prepare(saveUserQuery)
	if err != nil {
		return "", err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return "", err
	}

	_, err = stmt.Exec(u.Email, hashedPassword, u.FirstName, u.LastName, u.Uuid)
	if err != nil {
		return "", err
	}

	return u.Uuid.String(), nil
}

func (u *User) ValidateUser() error {
	row := db.DB.QueryRow(getUserQuery, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.Id, &retrievedPassword)
	if err != nil {
		return errors.New(loginErrorMessage)
	}

	if !utils.CheckPasswordHash(u.Password, retrievedPassword) {
		return errors.New(loginErrorMessage)
	}

	return nil
}
