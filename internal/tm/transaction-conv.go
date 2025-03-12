package tm

import (
	"time"
	"tm/internal/data"
	"tm/pkg/contracts"
)

func transactionToContract(t data.TmTransaction) contracts.Transaction {
	return contracts.Transaction{
		Id:          t.ID,
		Date:        t.Date.Format(time.RFC3339),
		Description: t.Description,
		AmountCents: int(t.AmountCents),
		Category:    t.CategoryID,
	}
}

func transactionsToContracts(ts []data.TmTransaction) []contracts.Transaction {
	contractTransactions := make([]contracts.Transaction, len(ts))
	for i, t := range ts {
		contractTransactions[i] = transactionToContract(t)
	}
	return contractTransactions
}
