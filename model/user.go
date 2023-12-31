package model

import (
	"net/mail"
	"time"
)

type User struct {
	UserID    int64     `db:"user_id"`
	Email     string    `db:"email"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func ValidateUser(email string, name string) error {
	if len(email) > 255 {
		return &ValidationError{"email is too long"}
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return &ValidationError{"email is invalid"}
	}

	if len(name) < 3 {
		return &ValidationError{"name is too short"}
	}
	if len(name) > 255 {
		return &ValidationError{"name is too long"}
	}

	return nil
}
