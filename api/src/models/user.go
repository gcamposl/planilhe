package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	Active    bool      `json:"active,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// Prepare call the methods for validation and formatting the user data
func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}
	user.format()

	return nil
}
func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Nick == "" {
		return errors.New("nick is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
