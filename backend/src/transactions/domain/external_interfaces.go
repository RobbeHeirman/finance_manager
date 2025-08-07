package domain

type TransactionRepository interface {
	UpsertAccounts(transactions []*TransactionalAccount) error
	UpsertRecipients(transactions []*Recipient) error
	UpsertTransactions(transactions []*Transaction) error
}
