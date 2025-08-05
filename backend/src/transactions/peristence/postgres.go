package peristence

import (
	"context"
	_ "embed"
	"finance_manager/src/transactions/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type TransactionPostgresRepository struct {
	pool *pgxpool.Pool
}

func CreateNewTransactionRepository(pool *pgxpool.Pool) *TransactionPostgresRepository {
	return &TransactionPostgresRepository{
		pool: pool,
	}
}

//go:embed init_transaction.sql
var initTableQuery string

func (repo *TransactionPostgresRepository) Init() error {
	_, err := repo.pool.Exec(context.Background(), initTableQuery)
	if err != nil {
		slog.Error("Error initializing Postgres table", "error", err)
		return err
	}
	return nil
}

func (repo *TransactionPostgresRepository) UpsertTransactions(transactions *[]domain.Transaction) {
	//TODO implement me
	panic("implement me")
}
