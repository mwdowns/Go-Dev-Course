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
const todoCreated = "Todo Created!"

type Todo struct {
	Id   string `json:"uuid"`
	Text string `json:"text"`
}

type Todos []Todo

func (t Todo) info() {
	fmt.Printf("Todo ID: %s\nText: %s\n", t.Id, t.Text)
}

func (t Todo) Save() error {
	// this should get the JSON file
	// if no JSON file exists, create one
	// append note to file
	jsonFile, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(todoFileName, jsonFile, 0644)
}

func (t Todo) SuccessMessage() {
	fmt.Println(todoCreated)
	fmt.Printf("Todo text: %v\n", t.Text)
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

func DisplayTodos() {
	fmt.Println("display todos")
}

func CreateNewTodo() (Todo, error) {
	return Todo{}, nil
}
