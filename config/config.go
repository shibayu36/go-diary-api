package config

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	DbDsn string
}

func Load() (*Config, error) {
	config := &Config{}

	mConf := mysql.Config{
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Net:       "tcp",
		Addr:      fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName:    os.Getenv("DB_NAME"),
		ParseTime: true,
	}
	config.DbDsn = mConf.FormatDSN()

	return config, nil
}
