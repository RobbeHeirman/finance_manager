package peristence

import (
	"context"
	_ "embed"
	"finance_manager/src/transactions/domain"
	"github.com/jackc/pgx/v5"
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

func (repo *TransactionPostgresRepository) UpsertAccounts(transactions []*domain.TransactionalAccount) error {
	ctx := context.Background()
	tx, err := repo.pool.Begin(ctx)
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			slog.Error("Couldnt rollback", "error", err)
		}
	}(tx, ctx)
	if err != nil {
		slog.Error("Error starting transaction", "error", err)
		return err
	}

	_, err = tx.Exec(context.Background(),
		`CREATE TEMP TABLE staging_accounts 
			(
    			user_id UUID,
    			account_no UUID
            )
    		ON COMMIT DROP;`,
	)
	if err != nil {
		slog.Error("Error upserting accounts", "error", err)
		return err
	}
	rows := make([][]interface{}, len(transactions))
	for i, transaction := range transactions {
		rows[i] = []interface{}{transaction.UserId, transaction.AccountNo}
	}
	_, err = tx.CopyFrom(context.Background(), pgx.Identifier{"staging_accounts"}, []string{"user_id", "account_no"}, pgx.CopyFromRows(rows))
	if err != nil {
		slog.Error("Error upserting accounts", "error", err)
		return err
	}

	_, err = tx.Exec(context.Background(), `
		INSERT INTO transactional_account(user_id, account_no)
		SELECT (user_id, account_no) FROM staging_accounts
		ON CONFLICT (user_id, account_no)
		DO NOTHING;`,
	)
	if err != nil {
		slog.Error("Error upserting accounts", "error", err)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		slog.Error("Error committing transaction", "error", err)
		return err
	}
	return nil
}

func (repo *TransactionPostgresRepository) UpsertRecipients(transactions []*domain.Recipient) error {
	//TODO implement me
	panic("implement me")
}

func (repo *TransactionPostgresRepository) UpsertTransactions(transactions []*domain.Transaction) error {
	//TODO implement me
	panic("implement me")
}
