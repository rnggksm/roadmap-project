package handler

import (
	"expense-tracker/internal/service"
	"expense-tracker/model"
)

type ExpenseHandler struct {
	svc *service.TrackerService
}

func NewExpenseHandler(svc *service.TrackerService) *ExpenseHandler {
	return &ExpenseHandler{svc: svc}
}

func (h *ExpenseHandler) AddExpense(description string, amount float64) {
	h.svc.AddExpense(model.Expense{
		Description: description,
		Amount:      amount,
	})
}

func (h *ExpenseHandler) ListExpenses() {
	h.svc.ListExpenses()
}

func (h *ExpenseHandler) DeleteExpense(id int) {
	h.svc.DeleteExpense(id)
}

func (h *ExpenseHandler) SummaryExpenses() {
	h.svc.SummaryExpenses()
}

func (h *ExpenseHandler) UpdateExpense(id int, description string, amount float64) {
	h.svc.UpdateExpense(id, description, amount)
}
