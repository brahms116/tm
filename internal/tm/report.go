package tm

import (
	"context"
	"fmt"
	"time"
	"tm/internal/data"
	"tm/pkg/contracts"
)

const SMALL_SPEND_THRESHOLD = -10000

func smallSpends(transactions []data.TmTransaction) []data.TmTransaction {
	smallSpends := []data.TmTransaction{}
	for _, transaction := range transactions {
		if transaction.AmountCents > SMALL_SPEND_THRESHOLD && transaction.AmountCents < 0 {
			smallSpends = append(smallSpends, transaction)
		}
	}
	return smallSpends
}

func spends(transactions []data.TmTransaction) []data.TmTransaction {
	spends := []data.TmTransaction{}
	for _, transaction := range transactions {
		if transaction.AmountCents < 0 {
			spends = append(spends, transaction)
		}
	}
	return spends
}

func sumAmounts(transactions []data.TmTransaction) int {
	sum := int(0)
	for _, transaction := range transactions {
		sum += int(transaction.AmountCents)
	}
	return sum
}

func totalSpendings(transactions []data.TmTransaction) contracts.TotalSpendings {
	spends := spends(transactions)
	smallSpends := smallSpends(transactions)
	return contracts.TotalSpendings{
		Total:          sumAmounts(spends),
		SmallSpendings: sumAmounts(smallSpends),
	}
}

func getStartEndPeriod(dateWeek time.Time, numWeeks int) (time.Time, time.Time) {
	weekStart := dateWeek.AddDate(0, 0, -int(dateWeek.Weekday()))
	end := weekStart.AddDate(0, 0, 6)
	start := weekStart.AddDate(0, 0, -7*numWeeks)
	return start, end
}

func totalAndNet(transactions []data.TmTransaction) (contracts.TotalSpendings, int) {
	return totalSpendings(transactions), sumAmounts(transactions)
}



func (t *tm) Report(ctx context.Context, dateWeek time.Time, numWeeks int) (contracts.ReportResponse, error) {
	start, end := getStartEndPeriod(dateWeek, numWeeks)
	params := data.ListTransactionsParams{
		Date:   start,
		Date_2: end,
	}

	transactions, err := data.New().ListTransactions(ctx, t.conn, params)

	if err != nil {
		return contracts.ReportResponse{}, fmt.Errorf("error listing transactions: %w", err)
	}

	return contracts.ReportResponse{}, nil
}
