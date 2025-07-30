package domain

type TransactionRepository interface {
	UpsertTransactions(transactions *[]Transaction)
}
