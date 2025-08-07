package adapters

import (
	"errors"
	"finance_manager/src/core/data_structures"
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
	accounts     map[string]*domain.TransactionalAccount
	recipients   map[string]*domain.Recipient
	transactions []*domain.Transaction
}

func NewParserManager(ExpectedCapacity int) *KbcParserManager {
	return &KbcParserManager{
		accounts:     make(map[string]*domain.TransactionalAccount, 1),
		recipients:   make(map[string]*domain.Recipient, ExpectedCapacity/2),
		transactions: make([]*domain.Transaction, ExpectedCapacity),
	}
}

func (p *KbcParserManager) ParseLine(inp *[]string) error {
	err := ParseAndUpdateMap(inp, kbcCSVLineToTransactionalAccount, p.accounts, 0)
	if err != nil {
		return err
	}

	err = ParseAndUpdateMap(inp, kbcCSVLineToRecipientAccount, p.recipients, 0)
	if err != nil {
		return err
	}

	transaction, err := kbcCSVLineToTransactionLine(inp)
	if err != nil {
		slog.Error("Could not parse line", "line", inp, "error", err)
		return err
	}
	p.transactions = append(p.transactions, transaction)
	return nil
}

func (p *KbcParserManager) GetAccounts() []*domain.TransactionalAccount {
	return data_structures.GetMapValues(p.accounts)
}

func (p *KbcParserManager) GetRecipients() []*domain.Recipient {
	return data_structures.GetMapValues(p.recipients)
}

func (p *KbcParserManager) GetTransactions() []*domain.Transaction {
	return p.transactions
}
