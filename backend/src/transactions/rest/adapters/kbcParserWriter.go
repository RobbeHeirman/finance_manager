package adapters

import (
	"finance_manager/src/core/data_structures"
	"finance_manager/src/transactions/domain"
)

func KbcParserWriter(domain domain.TransactionService, manager *KbcParserManager, ret chan error) {
	defer close(ret)

	accountPromise := data_structures.NewPromise(func() error {
		return domain.UpsertTransactionalAccounts(manager.GetAccounts())
	},
	)
	recipientPromise := data_structures.NewPromise(func() error {
		return domain.UpsertRecipient(manager.GetRecipients())
	})

	accountErr := accountPromise.AwaitPromise()
	if accountErr != nil {
		ret <- accountErr
	}
	recepErr := recipientPromise.AwaitPromise()
	if recepErr != nil {
		ret <- recepErr
	}
	if accountErr != nil || recepErr != nil {
		return
	}

	transactionErr := data_structures.NewPromise(func() error {
		return domain.UpsertTransactions(manager.GetTransactions())
	}).AwaitPromise()
	if transactionErr != nil {
		ret <- transactionErr
	}
}
