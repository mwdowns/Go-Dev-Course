package models

import (
	"encoding/json"
	db "mwdowns/rest-api/DB"
)

type User struct {
	id        int8
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	FirstName string `binding:"required"`
	LastName  string `binding:"required"`
	Uuid      string
}

const usersTableName = "users"

func (u User) Save() (string, error) {
	client, err := db.Client()
	if err != nil {
		return "", err
	}
	data, _, err := client.From(usersTableName).
		Insert(u.inputs(), false, "", "", "exact").
		Execute()
	if err != nil {
		return "", err
	}
	var r result
	err = json.Unmarshal(data, &r)
	if err != nil {
		return "", err
	}
	return r.buildUser(r[0].(map[string]interface{})).Uuid, err
}

func (u User) inputs() map[string]interface{} {
	return map[string]interface{}{
		"email":      u.Email,
		"password":   u.Password,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
	}
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
