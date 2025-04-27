package models

import (
	"errors"
	"time"
)

type Transaction struct {
	ID              uint64    `json:"id,omitempty"`
	UserID          uint64    `json:"user_id,omitempty"`
	Type            string    `json:"type,omitempty"` // 'income' or 'expensive'
	Amount          float64   `json:"amount,omitempty"`
	Description     string    `json:"description,omitempty"`
	Category        string    `json:"category,omitempty"`
	TransactionDate time.Time `json:"transaction_date,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	DeletedAt       time.Time `json:"deleted_at,omitempty"`
}

func (transaction *Transaction) Validate() error {
	if transaction.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if transaction.Description == "" {
		return errors.New("description is required")
	}

	if transaction.Category == "" {
		return errors.New("category is required")
	}

	if transaction.TransactionDate.IsZero() {
		return errors.New("transaction date is required")
	}

	return nil
}
