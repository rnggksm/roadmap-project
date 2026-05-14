package service

import (
	"encoding/json"
	"expense-tracker/model"
	"fmt"
	"os"
	"strings"
	"time"
)

type TrackerService struct {
	Expenses []model.Expense
}

func NewTrackerService() (*TrackerService, error) {
	// checking data.json is exist or not
	if _, err := os.Stat("data.json"); os.IsNotExist(err) {
		// create data.json
		file, err := os.Create("data.json")
		if err != nil {
			fmt.Println("Error creating data.json:", err)
			return nil, err
		}
		file.WriteString("[]")
		defer file.Close()
	}

	// read data from data.json
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading data.json:", err)
		return nil, err
	}

	var expenses []model.Expense
	err = json.Unmarshal(data, &expenses)

	return &TrackerService{
		Expenses: expenses,
	}, nil
}

func (s *TrackerService) AddExpense(expense model.Expense) {
	expense.ID = s.getNewID()
	expense.CreatedAt = time.Now().Format(time.RFC3339)
	expense.UpdatedAt = time.Now().Format(time.RFC3339)

	s.Expenses = append(s.Expenses, expense)

	err := s.save()
	if err != nil {
		fmt.Println("Error saving expense:", err)
	}

	fmt.Printf("Expense added successfully (ID: %d)\n", expense.ID)
}

func (s *TrackerService) ListExpenses() {
	if len(s.Expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	fmt.Printf("ID | Date | Description | Amount\n")
	fmt.Println("---|------|-------------|-------")
	for _, expense := range s.Expenses {
		date := strings.Split(expense.CreatedAt, "T")[0]
		fmt.Printf("%d | %s | %s | %.2f\n", expense.ID, date, expense.Description, expense.Amount)
	}
}

func (s *TrackerService) getNewID() int {
	// get latest ID and increment by 1
	if len(s.Expenses) == 0 {
		return 1
	}
	return s.Expenses[len(s.Expenses)-1].ID + 1
}

func (s *TrackerService) save() error {
	// save expenses to data.json

	data, err := json.MarshalIndent(s.Expenses, "", "  ")
	if err != nil {
		fmt.Println("Error when marshaling data")
		return err
	}

	err = os.WriteFile("data.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing to data.json:", err)
		return err
	}

	return nil
}

func (s *TrackerService) DeleteExpense(id int) {
	for i, expense := range s.Expenses {
		if expense.ID == id {
			s.Expenses = append(s.Expenses[:i], s.Expenses[i+1:]...)
			break
		}
	}

	err := s.save()
	if err != nil {
		fmt.Println("Error saving expense:", err)
	}

	fmt.Println("Expense deleted successfully")
}

func (s *TrackerService) SummaryExpenses() {
	total := 0.0
	for _, expense := range s.Expenses {
		total += expense.Amount
	}

	fmt.Printf("Total Expenses: $%.2f\n", total)
}

func (s *TrackerService) UpdateExpense(id int, description string, amount float64) {
	for i, expense := range s.Expenses {
		if expense.ID == id {
			s.Expenses[i].Description = description
			s.Expenses[i].Amount = amount
			s.Expenses[i].UpdatedAt = time.Now().Format(time.RFC3339)
			break
		}
	}

	err := s.save()
	if err != nil {
		fmt.Println("Error saving expense:", err)
	}

	fmt.Println("Expense updated successfully")
}
