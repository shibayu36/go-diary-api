package repository

import (
	"os"
	"testing"

	"github.com/Songmu/flextime"
	"github.com/shibayu36/go-diary-api/config"
)

var (
	repos *Repositories
)

func TestMain(m *testing.M) {
	restore := flextime.Fix(flextime.Now())
	defer restore()

	c, err := config.Load()
	if err != nil {
		panic(err)
	}
	repos, err = NewRepositories(c.DbDsn)
	if err != nil {
		panic(err)
	}
	defer repos.Close()

	os.Exit(m.Run())
}
