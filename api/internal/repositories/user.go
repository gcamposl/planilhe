package repositories

import (
	"api/internal/models"
	"database/sql"
	"fmt"
)

type User struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sql.DB) *User {
	return &User{db}
}

// Create creates a new user
func (repo User) Create(user models.User) (uint64, error) {
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
func (repo User) Find(nick string) ([]models.User, error) {
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
func (repo User) FindById(id uint64) (models.User, error) {
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

// FindByEmail returns a user by email
func (repo User) FindByEmail(email string) (models.User, error) {
	query := "select id, password from users where email = ?"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	line := stmt.QueryRow(email)
	var user models.User
	if err := line.Scan(&user.ID, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, err
		}
	}

	return user, nil
}

// Update updates a user
func (repo User) Update(user models.User) error {
	stmt, err := repo.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)

	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Name, user.Nick, user.Email, user.ID); err != nil {
		return err
	}

	return nil
}

// Delete deletes a user
func (repo User) Delete(id uint64) error {
	stmt, err := repo.db.Prepare(
		"delete from users where id = ?",
	)

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

// FindPassword returns the password of a user by id
func (repo User) FindPassword(id uint64) (string, error) {
	line, err := repo.db.Query("select password from users where id = ?", id)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

// UpdatePassword updates the password of a user
func (repo User) UpdatePassword(id uint64, password string) error {
	stmt, err := repo.db.Prepare("update users set password = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(password, id); err != nil {
		return err
	}

	return nil
}
