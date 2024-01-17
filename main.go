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

// type displayer interface {
// 	Display()
// }

// type outputtable interface {
// 	Save() error
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

func main() {

	printSomething(100)
	printSomething(1.618)
	printSomething("Oli")

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

	err = outputData(todo)
	if err != nil {
		return
	}

	outputData(userNote)

}

func printSomething(value interface{}) {
	// using interface{} allows this function to receive any value of any type.

	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer: ", intVal)
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float: ", floatVal)
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String: ", stringVal)
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float: ", value)
	// case string:
	// 	fmt.Println("String: ", value)
	// }

}

func outputData(data outputtable) error {
	//This function Display the data in shell and also writes it in file.
	data.Display()
	return saveData(data)
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
