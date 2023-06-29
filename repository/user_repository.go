package repository

import (
	"database/sql"

	"github.com/Songmu/flextime"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shibayu36/go-diary-api/model"
)

func (r *Repository) CreateUser(email string, name string) (*model.User, error) {
	if err := model.ValidateUser(email, name); err != nil {
		return nil, err
	}

	now := flextime.Now()
	res, err := r.db.Exec(
		`INSERT INTO users (email, name, created_at, updated_at)
			VALUES (?, ?, ?, ?)`,
		email, name, now, now,
	)
	if err != nil {
		if IsDuplicationError(err) {
			return nil, NewDuplicationError(err.Error())
		}

		return nil, err
	}

	id, _ := res.LastInsertId()
	user := &model.User{
		UserID:    id,
		Email:     email,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return user, nil
}

func (r *Repository) FindUserByID(id int64) (*model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE user_id = ?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewNotFoundError("user")
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewNotFoundError("user")
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) FindUserByName(name string) (*model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE name = ?", name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NewNotFoundError("user")
		}
		return nil, err
	}
	return &user, nil
}
