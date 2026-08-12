package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Enter your Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Enter your TaxRate: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Enter your Expenses: ")
	if err != nil {
		//		fmt.Println(err)
		panic(err)
	}

	ebt, profit, ratio := calculate_financials(revenue, expenses, taxRate)
	fmt.Printf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	storeResults(ebt, profit, ratio)
}

// Performing calculations
func calculate_financials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

// Taking user input
func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Value must be a postive number")
	}
	return userInput, nil

}

// storing output to a file
func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
