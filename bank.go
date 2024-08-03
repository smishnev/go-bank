package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile, 1000)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------------------")

		// panic("Can't continue, Sorry")
	}

	fmt.Println("Welcome to Go Bank")

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("How much do you want to deposit?: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")

				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your balance is: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("How much do you want to withdraw?: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than your balance")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Your balance is: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Bye!")
			fmt.Println("Thank you for using our Bank")

			return
		}

	}

}
