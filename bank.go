package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	balance, _ := os.ReadFile(accountBalanceFile)

	balanceText, _ := strconv.ParseFloat(string(balance), 64)

	return balanceText
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprintf("%.2f", balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance float64 = getBalanceFromFile()

	fmt.Println("Welcome to Go Bank")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Bye!")
			fmt.Println("Thank you for using our Bank")

			return
		}

	}

}
