package main

import (
	"fmt"
)

func main() {
	revenue, taxRate, expenses := getUserInput("Enter your Revenue, TaxRate and Expenses in this order space separated: ")
	ebt, profit, ratio := calculate_financials(revenue, expenses, taxRate)
	fmt.Printf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
}

// Performing calculations
func calculate_financials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

// Taking user input
func getUserInput(infoText string) (float64, float64, float64) {
	var userInputRevenue, userInputTaxRate, userInputexpenses float64
	fmt.Print(infoText)
	fmt.Scan(&userInputRevenue, &userInputTaxRate, &userInputexpenses)
	return userInputRevenue, userInputTaxRate, userInputexpenses

}
