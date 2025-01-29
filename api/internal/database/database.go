package database

import (
	"api/internal/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

const (
	driver = "mysql"
)

// Connect opens the connection with the database
func Connect() (*sql.DB, error) {
	db, err := sql.Open(driver, config.ConnectionString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
