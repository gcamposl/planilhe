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

func (repo *Transactions) Create(trs models.Transactions) (uint64, error) {
	sql := `
		insert into transactions 
			(user_id, type, amount, description, category, transaction_date)
		values
			(?, ?, ?, ?, ?, ?)`

	stmt, err := repo.db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		trs.UserID,
		trs.Type,
		trs.Amount,
		trs.Description,
		trs.Category,
		trs.TransactionDate,
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
