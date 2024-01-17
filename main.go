package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/HernanAlbertoLaborde/note-json/note"
	"github.com/HernanAlbertoLaborde/note-json/todo"
)

type saver interface {
	Save() error // Have a method Save and returns an error.
}

func main() {

	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()

	err = saveData(todo)
	if err != nil {

	}

	userNote.Display()

	err = saveData(userNote)
	if err != nil {
		return
	}

}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving fail")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (title, content string) {
	title = getUserInput("Note title:")

	content = getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) (text string) {

	fmt.Printf("%v ", prompt)
	//fmt.Scanln(&value)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n') // Here is important single quotes because it's a rune.

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
