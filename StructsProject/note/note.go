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

const noteError = "could not create note"
const noteCreated = "Note CreatedAt!"
const invalidNoteError = "you must have something in your note"
const titlePrompt = "What is the title of your note?"
const contentPrompt = "What is your note?"
const noteSaveError = "could not save note"

type Note struct {
	Id        string    `json:"uuid"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n Note) noteInfo() {
	fmt.Printf("Note ID: %s\nTitle: %s\nContent: %s\n", n.Id, n.Title, n.Content)
}

func (n Note) save() error {
	fileName := n.Id + ".json"
	jsonFile, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonFile, 0644)
}

type Notes []Note

func (n Notes) showNotes() {
	for _, note := range n {
		note.noteInfo()
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

func ShowNotes() {
	note1, _ := New("first note", "first content")
	note2, _ := New("second note", "second content")
	notes := Notes{*note1, *note2}
	notes.showNotes()
}

func CreateNewNote() {
	title, content := getNoteData()
	// create note using New
	note, err := New(title, content)
	// if err show error message
	if err != nil {
		fmt.Println(noteError)
		fmt.Println(err)
		return
	}
	// write to JSON file
	err = note.save()
	if err != nil {
		fmt.Println(noteSaveError)
	}
	// show success message
	fmt.Println(noteCreated)
	fmt.Printf("Note title: %v\nNote content: %v\n", note.Title, note.Content)
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
