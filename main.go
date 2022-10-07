package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/superlinkx/go-rest-barebones/app"
	"github.com/superlinkx/go-rest-barebones/config"
	"github.com/superlinkx/go-rest-barebones/database"
	"github.com/superlinkx/go-rest-barebones/router"
	"github.com/superlinkx/go-rest-barebones/sqlc"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	var (
		db                   *sql.DB
		appConfig            config.Config
		maxReconnectAttempts int
		attempt              = 0
	)

	if err := godotenv.Load(); err != nil {
		log.Print("WARNING: No .env file found. Using environment variables.")
	}
	appConfig = config.NewConfig()

	if mra, err := strconv.Atoi(appConfig.MaxReconnectAttempts); err != nil {
		log.Printf("WARNING: MaxReconnectAttempts not a valid number: %v", err)
		maxReconnectAttempts = 10
	} else {
		maxReconnectAttempts = mra
	}

	for {
		if attempt < maxReconnectAttempts {
			if ndb, err := database.NewDatabase(appConfig); err != nil {
				log.Printf("Database connnection failure: %v", err)
				attempt++
				time.Sleep(time.Second * 10)
				continue
			} else {
				log.Print("Database connection success. Continuing...")
				db = ndb
				break
			}
		} else {
			log.Fatal("Exhausted reconnection attempts")
		}
	}

	defer db.Close()
	ctx := context.Background()
	queries := sqlc.New(db)
	app := app.NewApp(db, queries, &appConfig, ctx)
	log.Printf("Starting server on %s", appConfig.GetAppHostString())
	log.Fatal(http.ListenAndServe(appConfig.GetAppHostString(), router.NewRouter(app)))
}
