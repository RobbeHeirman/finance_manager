package rest

import (
	"errors"
	"finance_manager/src/transactions/domain"
	"log/slog"
)

func kbcCSVLineToTransactionalAccount(inp *[]string) (*domain.TransactionalAccount, error) {
	return nil, errors.New("Not yet implemented")
}

func kbcCSVLineToRecipientAccount(inp *[]string) (*domain.Recipient, error) {
	return nil, errors.New("Not yet implemented")
}

func kbcCSVLineToTransactionLine(inp *[]string) (*domain.Transaction, error) {
	return nil, errors.New("Not yet implemented")
}

type KbcParserManager struct {
	Accounts     map[string]*domain.TransactionalAccount
	Recipients   map[string]*domain.Recipient
	Transactions []*domain.Transaction
}

func NewParserManager(ExpectedCapacity int) *KbcParserManager {
	return &KbcParserManager{
		Accounts:     make(map[string]*domain.TransactionalAccount, 1),
		Recipients:   make(map[string]*domain.Recipient, ExpectedCapacity/2),
		Transactions: make([]*domain.Transaction, ExpectedCapacity),
	}
}

func (p *KbcParserManager) ParseLine(inp *[]string) bool {
	transactionalAccountNo, err := kbcCSVLineToTransactionalAccount(inp)
	if err != nil {
		slog.Error("Could not parse line", "line", inp, "error", err)
		return false
	}
	p.Accounts[transactionalAccountNo.AccountNo] = transactionalAccountNo

	recipientAccountNo, err := kbcCSVLineToRecipientAccount(inp)
	if err != nil {
		slog.Error("Could not parse line", "line", inp, "error", err)
		return false
	}
	p.Recipients[recipientAccountNo.AccountNo] = recipientAccountNo
	transaction, err := kbcCSVLineToTransactionLine(inp)
	if err != nil {
		slog.Error("Could not parse line", "line", inp, "error", err)
		return false
	}
	p.Transactions = append(p.Transactions, transaction)
	return true
}
