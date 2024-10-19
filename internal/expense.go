package internal

import "time"

type Expense struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	CreatedAt   time.Time `json:"createdAt"`
}
