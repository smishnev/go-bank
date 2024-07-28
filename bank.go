package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

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

		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
		} else if choice == 2 {
			fmt.Print("How much do you want to deposit?: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")

				return
			}

			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			fmt.Println("Your balance is: ", accountBalance)
		} else if choice == 3 {
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
		} else {
			fmt.Println("Bye!")

			break
		}
	}

	fmt.Println("Thank you for using our Bank")
}
