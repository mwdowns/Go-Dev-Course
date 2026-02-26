package models

import (
	"encoding/json"
	"errors"
	db "mwdowns/rest-api/DB"
	"mwdowns/rest-api/utils"

	"github.com/supabase-community/supabase-go"
)

type User struct {
	id        int8
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	FirstName string
	LastName  string
	Uuid      string
}

const usersTableName = "users"

func (u User) Save() (string, error) {
	client, userInputs, err := initializeClientAndGetInputs(u)
	if err != nil {
		return "", err
	}
	r, err := saveToDb(client, userInputs)
	if err != nil {
		return "", err
	}
	return r.buildUser(r[0].(map[string]interface{})).Uuid, err
}

func (u User) ValidateUser() (bool, User, error) {
	client, _, err := initializeClientAndGetInputs(u)
	if err != nil {
		return false, u, err
	}

	data, _, err := client.From(usersTableName).
		Select("*", "1", false).
		Eq("email", u.Email).
		Execute()
	if err != nil {
		return false, u, errors.New("wrong email")
	}
	var r result
	err = json.Unmarshal(data, &r)
	if err != nil {
		return false, u, err
	}
	retunedUser := r.buildUser(r[0].(map[string]interface{}))
	if utils.CheckPasswordHash(u.Password, retunedUser.Password) {
		return true, retunedUser, nil
	}
	return false, u, errors.New("wrong password")
}

func (u User) inputs() (map[string]interface{}, error) {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"email":      u.Email,
		"password":   hashedPassword,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
	}, nil
}

func (r result) buildUser(m map[string]interface{}) User {
	u := User{
		Email:     m["email"].(string),
		Password:  m["password"].(string),
		FirstName: m["first_name"].(string),
		LastName:  m["last_name"].(string),
		Uuid:      m["uuid"].(string),
	}
	return u
}

func initializeClientAndGetInputs(u User) (*supabase.Client, map[string]interface{}, error) {
	client, err := db.Client()
	if err != nil {
		return nil, nil, err
	}
	userInputs, err := u.inputs()
	if err != nil {
		return nil, nil, err
	}
	return client, userInputs, nil
}

func saveToDb(client *supabase.Client, userInputs map[string]interface{}) (result, error) {
	data, _, err := client.From(usersTableName).
		Insert(userInputs, false, "", "", "exact").
		Execute()
	if err != nil {
		return nil, err
	}
	var r result
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
