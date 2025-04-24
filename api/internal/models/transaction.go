package models

import "time"

type Transactions struct {
	ID              int       `json:"id,omitempty"`
	UserID          int       `json:"user_id,omitempty"`
	Type            string    `json:"type,omitempty"` // 'income' or 'expensive'
	Amount          float64   `json:"amount,omitempty"`
	Description     string    `json:"description,omitempty"`
	Category        string    `json:"category,omitempty"`
	TransactionDate time.Time `json:"transaction_date,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	DeletedAt       time.Time `json:"deleted_at,omitempty"`
}
