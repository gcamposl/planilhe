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
func (userRepo user) Create(user models.User) (uint64, error) {
	stmt, err := userRepo.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}
