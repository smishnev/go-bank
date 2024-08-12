package main

import (
	"fmt"
	"os"
	"strings"

	"bufio"

	"example.com/main/note"
	"example.com/main/todo"
	// "example.com/main/fileops"
	// "github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	// var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile, 1000)

	// pointer()

	// structs()

	// fmt.Println("--------------------")

	// if err != nil {
	// 	fmt.Println("Error")
	// 	fmt.Println(err)
	// 	fmt.Println("--------------------")

	// 	// panic("Can't continue, Sorry")
	// }

	// fmt.Println("Welcome to Go Bank")
	// fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	// for {
	// 	presentOptions()

	// 	var choice int
	// 	fmt.Print("Your choice: ")
	// 	fmt.Scan(&choice)

	// 	switch choice {
	// 	case 1:
	// 		fmt.Println("Your balance is: ", accountBalance)
	// 	case 2:
	// 		fmt.Print("How much do you want to deposit?: ")
	// 		var depositAmount float64
	// 		fmt.Scan(&depositAmount)

	// 		if depositAmount <= 0 {
	// 			fmt.Println("Invalid amount. Must be greater than 0")

	// 			continue
	// 		}

	// 		accountBalance += depositAmount
	// 		fmt.Println("Your balance is: ", accountBalance)
	// 		fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
	// 	case 3:
	// 		fmt.Print("How much do you want to withdraw?: ")
	// 		var withdrawAmount float64
	// 		fmt.Scan(&withdrawAmount)

	// 		if withdrawAmount <= 0 {
	// 			fmt.Println("Invalid amount. Must be greater than 0")
	// 			continue
	// 		}

	// 		if withdrawAmount > accountBalance {
	// 			fmt.Println("Invalid amount. You can't withdraw more than your balance")
	// 			continue
	// 		}

	// 		accountBalance -= withdrawAmount
	// 		fmt.Println("Your balance is: ", accountBalance)
	// 		fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
	// 	default:
	// 		fmt.Println("Bye!")
	// 		fmt.Println("Thank you for using our Bank")

	// 		return
	// 	}

	// }

	printSomething(1)
	printSomething(1.5)
	printSomething("Text")

	result := genericFunction(1, 2)

	fmt.Println(result)

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

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

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Error saving note")
		return err
	}

	fmt.Println("Note saved successfully")

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content

}

func getUserInput(promptText string) string {
	fmt.Printf("%v ", promptText)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("integer", intVal+1)
	}

	intFloat, ok := value.(float64)

	if ok {
		fmt.Println("intFloat", intFloat+1.0)
	}

	strVal, ok := value.(string)

	if ok {
		fmt.Println("string", strVal)
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("int", value)
	// case float64:
	// 	fmt.Println("float64", value)
	// case string:
	// 	fmt.Println("string", value)
	// }
}

func genericFunction[T int | float64 | string](a, b T) T {
	return a + b
}
