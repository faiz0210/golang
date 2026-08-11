package main

import (
	"fmt"
)

func main() {

	var accountBalance float64 = 1000
	fmt.Println("Welcome to Go Bank")
	for {
		fmt.Println("How can I help you today")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Please select an option: ")
		fmt.Scan(&choice)
		fmt.Println("You have choosen: ", choice)

		//	wantsCheckBalance := choice == 1

		if choice == 1 {
			fmt.Println("Your Balance is: ", accountBalance)

		} else if choice == 2 {
			fmt.Println("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount, please enter a valid amount")
				continue
			}
			accountBalance += depositAmount
			fmt.Printf("Balance updated, New amount: %.1f\n", accountBalance)
		} else if choice == 3 {
			fmt.Print("How much do you want to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance || withdrawAmount <= 0 {
				fmt.Printf("Your current balance is %v, pleae enter a value lower or equal to %v.\n", accountBalance, accountBalance)
				continue
			} else {
				accountBalance -= withdrawAmount
				fmt.Printf("Your update balance is: %.2f\n", accountBalance)

			}
		} else {
			fmt.Println("Exiting!!!")
			break

		}
	}

	fmt.Println("Thank you for banking with us.")

}
