package repository

import "github.com/jmoiron/sqlx"

type ApiKeyRepository struct {
	db *sqlx.DB
}

func NewApiKeyRepository(db *sqlx.DB) *ApiKeyRepository {
	return &ApiKeyRepository{db: db}
}
