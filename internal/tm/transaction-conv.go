package tm

import (
	"time"
	"tm/internal/orm/model"
	"tm/pkg/contracts"
)

func transactionToContract(t model.TmTransaction) contracts.Transaction {
	return contracts.Transaction{
		Id:          t.ID,
		Date:        t.Date.Format(time.RFC3339),
		Description: t.Description,
		AmountCents: int(t.AmountCents),
		Category:    t.CategoryID,
	}
}

func transactionsToContracts(ts []model.TmTransaction) []contracts.Transaction {
	contractTransactions := make([]contracts.Transaction, len(ts))
	for i, t := range ts {
		contractTransactions[i] = transactionToContract(t)
	}
	return contractTransactions
}
