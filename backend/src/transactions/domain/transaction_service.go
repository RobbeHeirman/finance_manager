package domain

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/text/currency"
	"time"
)

type TransactionalAccount struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	AccountNo string
}

type Recipient struct {
	Id        uuid.UUID
	Name      string
	AccountNo string
}

type Amount struct {
	amount float32
	unit   currency.Unit
}

type Transaction struct {
	Id                     uuid.UUID
	OriginalTransactionId  string
	TransactionalAccountId uuid.UUID
	RecipientId            uuid.UUID
	Amount                 Amount
	TransactionDateTime    time.Time
}

type TransactionService interface {
	UpsertTransactions(transaction []*Transaction)
}

type TransactionServiceImpl struct {
	repository TransactionRepository
}

func CreateNewTransactionService(repo TransactionRepository) TransactionService {
	return &TransactionServiceImpl{repository: repo}
}

func (service *TransactionServiceImpl) UpsertTransactions(transactions []*Transaction) {
	for _, transaction := range transactions {
		fmt.Printf("%#v\n", transaction)
	}
}
