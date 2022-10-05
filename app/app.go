package app

import (
	"context"
	"database/sql"

	"github.com/superlinkx/go-rest-barebones/config"
	"github.com/superlinkx/go-rest-barebones/entity"
)

type App struct {
	DB      *sql.DB
	Queries *entity.Queries
	Config  *config.Config
	Ctx     context.Context
}

func NewApp(db *sql.DB, queries *entity.Queries, config *config.Config, ctx context.Context) App {
	return App{
		DB:      db,
		Queries: queries,
		Config:  config,
		Ctx:     ctx,
	}
}
