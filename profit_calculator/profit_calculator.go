package main

import (
	"fmt"
)

func main() {

	var revenue, expenses float64
	var taxRate float64

	fmt.Print("WHat is the total revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("What is the Tax Rate: ")
	fmt.Scan(&taxRate)

	fmt.Print("What are the total expenses: ")
	fmt.Scan(&expenses)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}
