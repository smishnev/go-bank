package main

import "fmt"

// "example.com/main/fileops"
// "github.com/Pallinder/go-randomdata"

const accountBalanceFile = "balance.txt"

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

// func main() {
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

// printSomething(1)
// printSomething(1.5)
// printSomething("Text")

// result := genericFunction(1, 2)

// fmt.Println(result)

// title, content := getNoteData()
// todoText := getUserInput("Todo text: ")

// todo, err := todo.New(todoText)

// if err != nil {
// 	fmt.Println(err)
// 	return
// }

// userNote, err := note.New(title, content)

// if err != nil {
// 	fmt.Println(err)
// 	return
// }

// err = outputData(todo)

// if err != nil {
// 	return
// }

// outputData(userNote)

// }

// func outputData(data outputtable) error {
// 	data.Display()
// 	return saveData(data)

// }

// func saveData(data saver) error {
// 	err := data.Save()

// 	if err != nil {
// 		fmt.Println("Error saving note")
// 		return err
// 	}

// 	fmt.Println("Note saved successfully")

// 	return nil
// }

// func getNoteData() (string, string) {
// 	title := getUserInput("Note title: ")

// 	content := getUserInput("Note content: ")

// 	return title, content

// }

// func getUserInput(promptText string) string {
// 	fmt.Printf("%v ", promptText)

// 	reader := bufio.NewReader(os.Stdin)

// 	text, err := reader.ReadString('\n')

// 	if err != nil {
// 		return ""
// 	}

// 	text = strings.TrimSuffix(text, "\n")
// 	text = strings.TrimSuffix(text, "\r")

// 	return text
// }

// func printSomething(value interface{}) {
// 	intVal, ok := value.(int)

// 	if ok {
// 		fmt.Println("integer", intVal+1)
// 	}

// 	intFloat, ok := value.(float64)

// 	if ok {
// 		fmt.Println("intFloat", intFloat+1.0)
// 	}

// 	strVal, ok := value.(string)

// 	if ok {
// 		fmt.Println("string", strVal)
// 	}

// 	// switch value.(type) {
// 	// case int:
// 	// 	fmt.Println("int", value)
// 	// case float64:
// 	// 	fmt.Println("float64", value)
// 	// case string:
// 	// 	fmt.Println("string", value)
// 	// }
// }

// func genericFunction[T int | float64 | string](a, b T) T {
// 	return a + b
// }

// Arrays
// func main() {
// 	price := [4]float64{1.1, 2.2, 3.3, 4.4}
// 	fmt.Println(price)

// 	// Slice
// 	featurePrice := price[1:3]
// 	fmt.Println(featurePrice)

// 	//Dynamic array

// 	dynamicArray := []float64{1.1, 2.2}
// 	dynamicArray[1] = 3.3

// 	updatedArray := append(dynamicArray, 4.4)

// 	//to re-assigning
// 	dynamicArray = append(dynamicArray, 4.4)
// 	fmt.Println(updatedArray, dynamicArray)

// }

// Maps

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// websites := map[string]string{
	// 	"Google":  "https://google.com",
	// 	"Youtube": "https://youtube.com",
	// }

	// fmt.Println((websites))
	// fmt.Println(websites["Google"])

	// websites["Facebook"] = "https://facebook.com"
	// fmt.Println(websites)

	// delete(websites, "Facebook")

	// fmt.Println(websites)

	userName := make([]string, 2, 5)

	userName = append(userName, "John")
	userName = append(userName, "Jane")

	fmt.Println(userName)

	courseRating := make(floatMap, 3)

	courseRating["JS"] = 4.7
	courseRating["Python"] = 4.8
	courseRating["Golang"] = 4.9

	courseRating.output()

	for index, value := range userName {
		fmt.Println(index, ":", value)
	}

	fmt.Println("Rating: ")
	for key, value := range courseRating {
		fmt.Println(key, ":", value)
	}
}
