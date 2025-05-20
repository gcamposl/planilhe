package repositories

import (
	"api/internal/models"
	"database/sql"
)

type Transactions struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *Transactions {
	return &Transactions{db}
}

func (repo *Transactions) Create(transaction models.Transaction) (uint64, error) {
	query := `
		insert into transactions 
			(user_id, type, amount, description, category, transaction_date)
		values
			(?, ?, ?, ?, ?, ?)`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		transaction.UserID,
		transaction.Type,
		transaction.Amount,
		transaction.Description,
		transaction.Category,
		transaction.TransactionDate,
	)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}

func (repo *Transactions) Find(userID uint64) ([]models.Transaction, error) {
	query := `
		select 
			id,
			user_id,
			type,
			amount,
			description,
			category,
			transaction_date,
			created_at,
			updated_at,
			deleted_at
		from transactions
		where user_id = ?
		order by transaction_date desc
	`

	stmt, err := repo.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var transactions []models.Transaction
	for stmt.Next() {
		var transaction models.Transaction
		if err := stmt.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.Type,
			&transaction.Amount,
			&transaction.Description,
			&transaction.Category,
			&transaction.TransactionDate,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&transaction.DeletedAt,
		); err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (repo *Transactions) Delete(transactionID uint64) error {
	query := "delete transactions where id = ?"

	stmt, err := repo.db.Query(query, transactionID)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return nil
}
