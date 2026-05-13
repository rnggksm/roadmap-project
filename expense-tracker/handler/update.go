package handler

import "fmt"

func UpdateExpense(id int, description string, amount float64) {
	fmt.Println("Updating expense")
	fmt.Println("ID:", id)
	fmt.Println("Description:", description)
	fmt.Println("Amount:", amount)
}
