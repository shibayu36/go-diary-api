package model

import "time"

// TODO: Commonize CreatedAt and UpdatedAt
type ApiKey struct {
	ApiKeyID  int64     `db:"api_key_id"`
	UserID    int64     `db:"user_id"`
	ApiKey    string    `db:"api_key"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
