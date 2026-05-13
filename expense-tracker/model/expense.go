package model

type Expense struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
