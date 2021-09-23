package repository

import (
	"fmt"

	"github.com/Toolnado/SwalloW/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(c *config.Config) (*sqlx.DB, error) {
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", c.DatabaseUser, c.DatabasePassword, c.DatabaseHost, c.DatabasePort, c.DatabaseName, c.DatabaseSSLMode)
	db, err := sqlx.Connect("postgres", databaseURL)

	if err != nil {
		return nil, err
	}

	return db, nil
}
