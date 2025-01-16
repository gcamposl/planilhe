package repositories

import (
	"api/src/models"
	"database/sql"
)

type user struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sql.DB) *user {
	return &user{db}
}

// Create creates a new user
func (u *user) Create(user models.User) (uint64, error) {
	return 0, nil
}
