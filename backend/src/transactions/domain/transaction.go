package domain

import "fmt"

type Transaction struct {
}

type TransactionService interface {
	UpsertTransactions(transaction []*Transaction)
}

type TransactionServiceImpl struct {
}

func (service *TransactionServiceImpl) UpsertTransactions(transactions []*Transaction) {
	for _, transaction := range transactions {
		fmt.Printf("%#v\n", transaction)
	}
}
