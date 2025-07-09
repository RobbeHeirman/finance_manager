package persistence

import (
	"context"
	_ "embed"
	"finance_manager/src/auth/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

//go:embed sql_queries/schema.sql
var iniUserTableQuery string

func NewUserRepo(pool *pgxpool.Pool) UserRepo {
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

func (repo *UserRepo) CreateUpdateUser(user *domain.User) (*uuid.UUID, error) {
	var id uuid.UUID
	err := repo.pool.QueryRow(
		context.Background(),
		`
			INSERT INTO "user" (email)
			VALUES ($1)
			ON CONFLICT (email) DO UPDATE SET email = EXCLUDED.email
			RETURNING id
			`,
		*user.GetEmail().ToString(),
	).Scan(&id)

	if err != nil {
		return nil, err
	}
	return &id, nil
}
