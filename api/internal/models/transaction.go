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

// TODO: como posso alterar esta funcao para validar somente o campo Amount quando for chamado na atualizacao de uma transacao? 
// atraves de um parametro talvez?
// mas como fazer isso sem quebrar as chamadas atuais?
// ok, mas preciso colocar um valor padrao para o parametro, senao todas as chamadas atuais iram quebrar
// como colocar um valor padrao em go?

func (transaction *Transaction) Validate(isUpdate bool) error {
	if transaction.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if !isUpdate {
		if transaction.Description == "" {
			return errors.New("description is required")
		}

		if transaction.Category == "" {
			return errors.New("category is required")
		}

		if transaction.TransactionDate.IsZero() {
			return errors.New("transaction date is required")
		}
	}
	

	return nil
}
