package core

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"os"
)

func CreateConnectionPoolFromEnv() (*pgxpool.Pool, error) {
	user := getEnv("DB_USER", "admin")
	password := getEnv("DB_PASSWORD", "admin")
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	database := getEnv("DB_DATABASE", "finance_manager")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		slog.Error("Could not connect to database", "connectionString", connectionString, "err", err)
		return nil, err
	}
	return pool, nil
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		slog.Warn("Environment variable not set. Defaulting to", key, fallback)
		return fallback
	}
	return value
}
