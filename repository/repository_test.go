package repository

import (
	"os"
	"testing"

	"github.com/Songmu/flextime"
	"github.com/shibayu36/go-diary-api/config"
	"github.com/stretchr/testify/assert"
)

var (
	repo *Repository
)

func TestMain(m *testing.M) {
	restore := flextime.Fix(flextime.Now())
	defer restore()

	c, err := config.Load()
	if err != nil {
		panic(err)
	}
	repo, err = NewRepository(c.DbDsn)
	if err != nil {
		panic(err)
	}
	defer repo.Close()

	os.Exit(m.Run())
}

func TestNewRepository(t *testing.T) {
	c, _ := config.Load()
	repo, err := NewRepository(c.DbDsn)

	assert.Nil(t, err)
	assert.Nil(t, repo.db.Ping(), "db should be connected")
}

func TestClose(t *testing.T) {
	c, _ := config.Load()
	repo, _ := NewRepository(c.DbDsn)

	assert.Nil(t, repo.Close())
	assert.NotNil(t, repo.db.Ping(), "db should be disconnected")
}
