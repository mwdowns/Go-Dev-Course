package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
)

const todoFileName = "todo.json"
const invalidTodoError = "cannot be empty"

type Todo struct {
	Id   string `json:"uuid"`
	Text string `json:"text"`
}

func (t Todo) info() {
	fmt.Printf("Todo ID: %s\nText: %s\n", t.Id, t.Text)
}

func (t Todo) save() error {
	// this should get the JSON file
	// if no JSON file exists, create one
	// append note to file
	jsonFile, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(todoFileName, jsonFile, 0644)
}

func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New(invalidTodoError)
	}

	return &Todo{
		Id:   uuid.New().String(),
		Text: text,
	}, nil
}

func Display() {
	fmt.Println("display todos")
}

func CreateNewTodo() {
	fmt.Println("create todos")
}
