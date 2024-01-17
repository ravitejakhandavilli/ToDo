package main

import (
	"bufio"
	"fmt"
	"os"
)

// create a func for displaying the menu
// User must be able to select options from menu
// func must be able take user input and save it
// func must be show all saved user notes
// must be able to delete notes
// program should exit when user select to quit option

type Notes struct {
	id    int
	title string
	text  string
}

type Note []Notes

var Notelist Note

func main() {

	toDoList()

}
func Intro() {
	fmt.Println("Welcome to Note Taker!")
	fmt.Println("1.Write a New Note to List")
	fmt.Println("2.See Your ToDo List")
	fmt.Println("3.Delete a Note From List")
	fmt.Println("4.Exit")
	fmt.Println("Please Select One Option")
	fmt.Println()
}
func scanner() string {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	return reader.Text()

}

func toDoList() {
	for {
		Intro()
		fmt.Print("-->")

		UserInput := scanner()
		// reader := bufio.NewReader(os.Stdin)
		// userInput, _ := reader.ReadString('\n')
		// userInput = strings.Replace(userInput, "\n", "", -1)
		//Selected, _ := strconv.Atoi(userInput)

		switch UserInput {
		case "1":
			fmt.Println("You selected to write a New Note to List")
			addNote()
		case "2":
			fmt.Println("You selected to See Your ToDo List")
			showNotes()
		case "3":
			fmt.Println("You selected to Delete a Note From List")
		case "4":
			fmt.Println("Exiting the Program...")
			return
		default:
			fmt.Println("Invalid choice. Please choose a valid option.")

		}

	}

}

func addNote() {

	fmt.Println("Enter Title")
	fmt.Print("--> ")
	Title := scanner()
	fmt.Println("Enter Text")
	fmt.Print("--> ")
	Text := scanner()

	new := Notes{
		id:    len(Notelist) + 1,
		title: Title,
		text:  Text,
	}

	Notelist = append(Notelist, new)
	fmt.Println("User Notes saved to toDoList...")
}

func showNotes() {
	fmt.Println("All Saved Notes:")
	for _, note := range Notelist {
		fmt.Printf("id:%d\nTitle:%s.\nNote:%s.\n\n", note.id, note.title, note.text)
	}
	fmt.Println()
}
