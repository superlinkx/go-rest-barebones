package app

import (
	"context"
	"database/sql"

	"github.com/superlinkx/go-rest-barebones/config"
	"github.com/superlinkx/go-rest-barebones/sqlc"
)

type App struct {
	DB      *sql.DB
	Queries *sqlc.Queries
	Config  *config.Config
	Ctx     context.Context
}

func NewApp(db *sql.DB, queries *sqlc.Queries, config *config.Config, ctx context.Context) App {
	return App{
		DB:      db,
		Queries: queries,
		Config:  config,
		Ctx:     ctx,
	}
}
