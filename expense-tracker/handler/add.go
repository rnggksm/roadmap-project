package handler

import "fmt"

func AddExpense(description string, amount float64) {
	fmt.Println("Adding expense")
	fmt.Println("Description:", description)
	fmt.Println("Amount:", amount)
}
