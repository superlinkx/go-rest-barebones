package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppHostName  string
	AppPort      string
	PsqlHost     string
	PsqlPort     string
	PsqlUser     string
	PsqlPassword string
	PsqlDb       string
	PsqlSslMode  string
}

func NewConfig() Config {
	config := Config{
		AppHostName:  "localhost",
		AppPort:      "8080",
		PsqlHost:     "localhost",
		PsqlPort:     "5432",
		PsqlUser:     "postgres",
		PsqlPassword: "postgres",
		PsqlDb:       "postgres",
	}

	if appHostName := os.Getenv("APP_HOSTNAME"); appHostName != "" {
		config.AppHostName = appHostName
	}

	if appPort := os.Getenv("APP_PORT"); appPort != "" {
		config.AppPort = appPort
	}

	if psqlHost := os.Getenv("PSQL_HOST"); psqlHost != "" {
		config.PsqlHost = psqlHost
	}

	if psqlPort := os.Getenv("PSQL_PORT"); psqlPort != "" {
		config.PsqlPort = psqlPort
	}

	if psqlUser := os.Getenv("PSQL_USER"); psqlUser != "" {
		config.PsqlUser = psqlUser
	}

	if psqlPassword := os.Getenv("PSQL_PASSWORD"); psqlPassword != "" {
		config.PsqlPassword = psqlPassword
	}

	if psqlDb := os.Getenv("PSQL_DB"); psqlDb != "" {
		config.PsqlDb = psqlDb
	}

	if psqlSslMode := os.Getenv("PSQL_SSLMODE"); psqlSslMode != "" {
		config.PsqlSslMode = psqlSslMode
	}

	return config
}

func (s Config) GetAppHostString() string {
	return fmt.Sprintf("%s:%s", s.AppHostName, s.AppPort)
}

func (s Config) GetPsqlHostString() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		s.PsqlUser, s.PsqlPassword, s.PsqlHost, s.PsqlPort, s.PsqlDb, s.PsqlSslMode)
}
