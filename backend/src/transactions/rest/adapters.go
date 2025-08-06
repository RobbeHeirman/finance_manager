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

func ParseAndUpdateMap[T any](inp *[]string, op func(inpt *[]string) (T, error), m map[string]T, idIndex int) error {
	key := (*inp)[idIndex]
	if _, ok := m[key]; ok {
		return nil
	}
	result, err := op(inp)
	if err != nil {
		slog.Error("Could not parse line", "line", inp, "error", err)
		return err
	}
	m[key] = result
	return nil
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

func (p *KbcParserManager) ParseLine(inp *[]string) error {
	err := ParseAndUpdateMap(inp, kbcCSVLineToTransactionalAccount, p.Accounts, 0)
	if err != nil {
		return err
	}

	err = ParseAndUpdateMap(inp, kbcCSVLineToRecipientAccount, p.Recipients, 0)
	if err != nil {
		return err
	}

	transaction, err := kbcCSVLineToTransactionLine(inp)
	if err != nil {
		slog.Error("Could not parse line", "line", inp, "error", err)
		return err
	}
	p.Transactions = append(p.Transactions, transaction)
	return nil
}
