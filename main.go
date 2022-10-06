package main

import (
	"context"
	"log"
	"net/http"

	"github.com/superlinkx/go-rest-barebones/app"
	"github.com/superlinkx/go-rest-barebones/config"
	"github.com/superlinkx/go-rest-barebones/database"
	"github.com/superlinkx/go-rest-barebones/router"
	"github.com/superlinkx/go-rest-barebones/sqlc"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	var appConfig config.Config

	if err := godotenv.Load(); err != nil {
		log.Print("WARNING: No .env file found. Using environment variables.")
	} else {
		appConfig = config.NewConfig()
	}

	if db, err := database.NewDatabase(appConfig); err != nil {
		log.Fatal(err)
	} else {
		defer db.Close()
		ctx := context.Background()
		queries := sqlc.New(db)
		app := app.NewApp(db, queries, &appConfig, ctx)
		log.Fatal(http.ListenAndServe(appConfig.GetAppHostString(), router.NewRouter(app)))
	}
}
