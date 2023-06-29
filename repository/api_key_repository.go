package repository

import (
	"database/sql"

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

func (r *ApiKeyRepository) FindByApiKey(apiKey string) (*model.ApiKey, error) {
	var key model.ApiKey
	err := r.db.Get(&key, "SELECT * FROM api_keys WHERE api_key = ?", apiKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewNotFoundError("api key")
		}
		return nil, err
	}
	return &key, nil
}
