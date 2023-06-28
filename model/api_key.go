package model

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

// TODO: Commonize CreatedAt and UpdatedAt
type ApiKey struct {
	ApiKeyID  int64     `db:"api_key_id"`
	UserID    int64     `db:"user_id"`
	ApiKey    string    `db:"api_key"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// GenerateApiKey generates a random string of 32 bytes and returns it as a hex string.
func GenerateApiKey() (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(randomBytes), nil
}
