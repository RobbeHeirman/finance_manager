package persistence

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

func CreateConnectionPool(config PostgresConfigRepository) (*pgxpool.Pool, error) {
	user := config.GetDatabaseUser()
	password := config.GetDatabasePassword()
	host := config.GetDatabaseHost()
	port := config.GetDatabasePort()
	database := config.GetDatabaseName()

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		slog.Error("Could not connect to database", "connectionString", connectionString, "err", err)
		return nil, err
	}
	return pool, nil
}
