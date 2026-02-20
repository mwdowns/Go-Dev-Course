package note

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

const noteFileName = "notes.json"
const noteError = "could not create note"
const noteCreated = "Note Created!"
const invalidNoteError = "you must have something in your note"
const titlePrompt = "What is the title of your note?"
const contentPrompt = "What is your note?"
const noteSaveError = "could not save note"
const createYourFirstNoteMessage = "No notes found. You should create a note!"

type Note struct {
	Id        string    `json:"uuid"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n Note) info() {
	fmt.Printf("Note ID: %s\nTitle: %s\nContent: %s\n", n.Id, n.Title, n.Content)
}

func (n Note) Save() error {
	// this should get the JSON file
	// if no JSON file exists, create one
	// append note to file
	jsonFile, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(noteFileName, jsonFile, 0644)
}

func (n Note) SuccessMessage() {
	fmt.Println(noteCreated)
	fmt.Printf("note title: %v\nnote content: %v\n", n.Title, n.Content)
}

type Notes []Note

func (n Notes) showNotesInfo() {
	for _, note := range n {
		note.info()
	}
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New(invalidNoteError)
	}

	return &Note{
		Id:        uuid.New().String(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func DisplayNotes() {
	notes := getNotes()
	//note1, _ := New("first note", "first content")
	//note2, _ := New("second note", "second content")
	//notes := Notes{*note1, *note2}
	if len(notes) != 0 {
		notes.showNotesInfo()
	}
	fmt.Println(createYourFirstNoteMessage)
}

func CreateNewNote() (Note, error) {
	title, content := getNoteData()
	// create note using New
	note, err := New(title, content)
	// if err show error message
	if err != nil {
		fmt.Println(noteError)
		fmt.Println(err)
		return Note{}, err
	}
	return *note, nil
}

func getNotes() Notes {
	return make(Notes, 0)
	// look for JSON file of name "notes.json"
	// if no JSON info exists, return empty slice
	// else take JSON info and turn them in to a Notes slice
}

func getNoteData() (title string, content string) {
	// get title
	title = getNoteInput(titlePrompt)
	// get content
	content = getNoteInput(contentPrompt)
	return title, content
}

func getNoteInput(prompt string) (input string) {
	fmt.Println(prompt)
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return ""
	}
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	return input
}
