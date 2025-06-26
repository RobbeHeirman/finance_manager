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

func (repo *UserRepo) CreateUpdateUser(user *domain.User) (*domain.User, error) {
	var id uuid.UUID
	err := repo.pool.QueryRow(
		context.Background(),
		`
			INSERT INTO "user" (email, first_name, last_name, picture_url)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT	(email) DO UPDATE 
			SET
			    first_name = EXCLUDED.first_name,
			    last_name = EXCLUDED.last_name,
			    picture_url = EXCLUDED.picture_url
			RETURNING id
			`,
		user.GetEmail().ToString(),
		user.GetFirstName(),
		user.GeTLastName(),
		user.GeTImageURL()).Scan(&id)

	if err != nil {
		return nil, err
	}
	return domain.NewUser(
		&id,
		user.GetEmail(),
		user.GetFirstName().GetUnchecked(),
		user.GeTLastName().GetUnchecked(),
		user.GeTImageURL().GetUnchecked(),
	)
}
