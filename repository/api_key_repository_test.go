package repository

import (
	"testing"

	"github.com/shibayu36/go-diary-api/testutil"
	"github.com/stretchr/testify/assert"
)

func TestApiKeyRepositoryCreateByUser(t *testing.T) {
	email := testutil.RandomEmail()
	name := testutil.RandomString(10)

	user, _ := repo.CreateUser(
		email, name,
	)

	apiKey1, err := repo.CreateApiKeyByUser(user)
	assert.Nil(t, err)
	assert.Equal(t, user.UserID, apiKey1.UserID)
	assert.Equal(t, 64, len(apiKey1.ApiKey))

	apiKey2, err := repo.CreateApiKeyByUser(user)
	assert.Nil(t, err)
	assert.Equal(t, user.UserID, apiKey2.UserID)
	assert.Equal(t, 64, len(apiKey2.ApiKey))
}
