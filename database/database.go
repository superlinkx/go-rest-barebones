package database

import (
	"database/sql"
	"fmt"

	"github.com/superlinkx/go-rest-barebones/config"
)

func NewDatabase(config config.Config) (*sql.DB, error) {
	if db, err := sql.Open("postgres", config.GetPsqlHostString()); err != nil {
		return nil, fmt.Errorf("error opening database: %s", err)
	} else if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	} else {
		return db, nil
	}
}
