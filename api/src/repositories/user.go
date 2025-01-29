package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type user struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sql.DB) *user {
	return &user{db}
}

// Create creates a new user
func (repo user) Create(user models.User) (uint64, error) {
	stmt, err := repo.db.Prepare(
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

// Find returns all users that match the nickname
func (repo user) Find(nick string) ([]models.User, error) {
	query := "select id, name, nick, email, created_at from users where nick like ?"
	nick = fmt.Sprintf("%%%s%%", nick)
	lines, err := repo.db.Query(query, nick)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User
		if err := lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// FindByID returns a user by id
func (repo user) FindById(id uint64) (models.User, error) {
	query := "select id, name, nick, email, created_at from users where id = ?"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	line := stmt.QueryRow(id)
	var user models.User
	if err := line.Scan(
		&user.ID,
		&user.Name,
		&user.Nick,
		&user.Email,
		&user.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, err
		}

		return models.User{}, err
	}

	return user, nil
}
