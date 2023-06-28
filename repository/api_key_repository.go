package repository

import (
	"github.com/Songmu/flextime"
	"github.com/jmoiron/sqlx"
	"github.com/shibayu36/go-diary-api/model"
)

type ApiKeyRepository struct {
	db *sqlx.DB
}

func NewApiKeyRepository(db *sqlx.DB) *ApiKeyRepository {
	return &ApiKeyRepository{db: db}
}

func (r *ApiKeyRepository) CreateByUser(user *model.User) (*model.ApiKey, error) {
	now := flextime.Now()

	key, err := model.GenerateApiKey()
	if err != nil {
		return nil, err
	}

	res, err := r.db.Exec(
		`INSERT INTO api_keys (user_id, api_key, created_at, updated_at)
			VALUES (?, ?, ?, ?)`,
		user.UserID, key, now, now,
	)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	apiKey := &model.ApiKey{
		ApiKeyID:  id,
		UserID:    user.UserID,
		ApiKey:    key,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return apiKey, nil
}
