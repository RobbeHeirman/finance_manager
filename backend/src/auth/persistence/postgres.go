package persistence

import (
	"context"
	_ "embed"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

//go:embed sql_queries/schema.sql
var iniUserTableQuery string

func CreateUserRepo(pool *pgxpool.Pool) UserRepo {
	return UserRepo{pool: pool}
}

type UserRepo struct {
	pool *pgxpool.Pool
}

func (repo *UserRepo) Init() error {
	_, err := repo.pool.Exec(context.Background(), iniUserTableQuery)
	if err != nil {
		slog.Error("Failed to init User table", iniUserTableQuery, err.Error())
	}
	return nil
}
