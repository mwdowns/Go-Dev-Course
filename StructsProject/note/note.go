package note

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const noteError = "could not create note"
const noteCreated = "Note created!"
const invalidNoteError = "you must have something in your note"
const titlePrompt = "What is the title of your note?"
const contentPrompt = "What is your note?"

type Note struct {
	id      string
	title   string
	content string
	created time.Time
}

func (n Note) noteInfo() {
	fmt.Printf("Note ID: %s\nTitle: %s\nContent: %s\n", n.id, n.title, n.content)
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
		id:      uuid.New().String(),
		title:   title,
		content: content,
		created: time.Now(),
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

	// show success message
	fmt.Println(noteCreated)
	fmt.Printf("Note title: %v\nNote content: %v\n", note.title, note.content)
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
	fmt.Scanln(&input)
	return input
}
