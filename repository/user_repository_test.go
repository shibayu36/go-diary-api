package repository

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/Songmu/flextime"
	"github.com/shibayu36/go-diary-api/model"
	"github.com/shibayu36/go-diary-api/testutil"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreate(t *testing.T) {
	t.Run("with valid parameters", func(t *testing.T) {
		email := testutil.RandomEmail()
		name := testutil.RandomString(10)

		user, err := repo.CreateUser(
			email, name,
		)
		assert.Nil(t, err)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, name, user.Name)
		assert.Equal(t, flextime.Now(), user.CreatedAt)
		assert.Equal(t, flextime.Now(), user.UpdatedAt)

		foundUser, _ := repo.FindUserByID(user.UserID)
		assert.Equal(t, user.UserID, foundUser.UserID, "user is created correctly")
	})

	t.Run("with invalid parameters", func(t *testing.T) {
		email := "invalidemail"
		name := testutil.RandomString(10)

		_, err := repo.CreateUser(
			email, name,
		)

		assert.Equal(t, model.NewValidationError("email is invalid"), err)
	})

	t.Run("with duplicated email", func(t *testing.T) {
		email := testutil.RandomEmail()
		name := testutil.RandomString(10)

		_, _ = repo.CreateUser(
			email, name,
		)

		_, err := repo.CreateUser(
			email, name,
		)

		var dupErr *DuplicationError
		assert.True(t, errors.As(err, &dupErr))
	})
}

func TestUserRepositoryFindByID(t *testing.T) {
	t.Run("user found", func(t *testing.T) {
		email := testutil.RandomEmail()
		name := testutil.RandomString(10)
		user, _ := repo.CreateUser(
			email, name,
		)

		foundUser, err := repo.FindUserByID(user.UserID)
		assert.Nil(t, err)
		assert.Equal(t, user.Email, foundUser.Email)
		assert.Equal(t, user.Name, foundUser.Name)
	})

	t.Run("user not found", func(t *testing.T) {
		foundUser, err := repo.FindUserByID(rand.Int63())
		assert.Nil(t, foundUser)
		assert.True(t, IsNotFound(err))
	})
}

func TestUserRepositoryFindByEmail(t *testing.T) {
	t.Run("user found", func(t *testing.T) {
		email := testutil.RandomEmail()
		name := testutil.RandomString(10)
		user, _ := repo.CreateUser(
			email, name,
		)

		foundUser, err := repo.FindUserByEmail(email)
		assert.Nil(t, err)
		assert.Equal(t, user.Email, foundUser.Email)
	})

	t.Run("user not found", func(t *testing.T) {
		foundUser, err := repo.FindUserByEmail(testutil.RandomEmail())
		assert.Nil(t, foundUser)
		assert.True(t, IsNotFound(err))
	})
}
