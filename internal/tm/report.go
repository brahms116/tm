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

func earns(transactions []data.TmTransaction) []data.TmTransaction {
	earns := []data.TmTransaction{}
	for _, transaction := range transactions {
		if transaction.AmountCents > 0 {
			earns = append(earns, transaction)
		}
	}
	return earns
}

func sumAmounts(transactions []data.TmTransaction) int {
	sum := 0
	for _, transaction := range transactions {
		sum += int(transaction.AmountCents)
	}
	return sum
}

func summary(transactions []data.TmTransaction) contracts.Summary {
	spends := spends(transactions)
	smallSpends := smallSpends(transactions)
	earns := earns(transactions)
	return contracts.Summary{
		Spending:      -1 * sumAmounts(spends),
		SmallSpending: -1 * sumAmounts(smallSpends),
		Earning:       sumAmounts(earns),
		Net:           sumAmounts(transactions),
	}
}

func monthSpan(dateMonth time.Time) (time.Time, time.Time) {
	start := time.Date(dateMonth.Year(), dateMonth.Month(), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)
	return start, end
}

func (t *tm) Report(ctx context.Context, dateMonth time.Time) (contracts.MonthReport, error) {
	start, end := monthSpan(dateMonth)
	params := data.ListTransactionsParams{
		Date:   start,
		Date_2: end,
	}

	curr, err := data.New().ListTransactions(ctx, t.conn, params)
	if err != nil {
		return contracts.MonthReport{}, fmt.Errorf("error listing transactions: %w", err)
	}

	start, end = monthSpan(dateMonth.AddDate(0, -1, 0))
	params = data.ListTransactionsParams{
		Date:   start,
		Date_2: end,
	}
	prev, err := data.New().ListTransactions(ctx, t.conn, params)
	if err != nil {
		return contracts.MonthReport{}, fmt.Errorf("error listing transactions: %w", err)
	}

	report := generateMonthReport(curr, prev, end)
  return report, nil
}

func generateMonthReport(curr, prev []data.TmTransaction, end time.Time) contracts.MonthReport {
	currSummary := summary(curr)
	prevSummary := summary(prev)
	diffSummary := contracts.Summary{
		Spending:      currSummary.Spending - prevSummary.Spending,
		SmallSpending: currSummary.SmallSpending - prevSummary.SmallSpending,
		Earning:       currSummary.Earning - prevSummary.Earning,
		Net:           currSummary.Net - prevSummary.Net,
	}

	periods := monthlyPeriodReports(curr, end)
	return contracts.MonthReport{
		Summary:           currSummary,
		SummaryComparison: diffSummary,
		Periods:           periods,
	}
}

func monthPeriodStart(end time.Time) time.Time {
	var prevMonday time.Time
	if end.Weekday() == 0 {
		prevMonday = end.AddDate(0, 0, -7)
	} else {
		prevMonday = end.AddDate(0, 0, -int(end.Weekday()))
	}

	monthStart := time.Date(end.Year(), end.Month(), 1, 0, 0, 0, 0, time.UTC)
	if prevMonday.After(monthStart) {
		return prevMonday
	}
	return monthStart
}

func monthlyPeriodReports(trans []data.TmTransaction, end time.Time) []contracts.MonthPeriodReport {
	periodStart := monthPeriodStart(end)
	curr, next := splitTransactionsByTime(trans, periodStart)
	currSummary := monthPeriodSummary(curr, end.Sub(periodStart))
	currSmallSpends := smallSpends(curr)
	currReport := contracts.MonthPeriodReport{
		StartDate:   periodStart,
		EndDate:     end,
		Summary:     currSummary,
		SmallSpends: currSmallSpends,
	}

	return append([]contracts.MonthPeriodReport{currReport}, monthlyPeriodReports(next, periodStart)...)
}

func splitTransactionsByTime(transactions []data.TmTransaction, at time.Time) ([]data.TmTransaction, []data.TmTransaction) {
	currTrans := []data.TmTransaction{}
	prevTrans := []data.TmTransaction{}
	for _, transaction := range transactions {
		if transaction.Date.After(at) {
			currTrans = append(currTrans, transaction)
		} else {
			prevTrans = append(prevTrans, transaction)
		}
	}
	return currTrans, prevTrans
}

func monthPeriodSummary(transactions []data.TmTransaction, duration time.Duration) contracts.MonthPeriodSummary {
	spends := spends(transactions)
	smallSpends := smallSpends(transactions)
	return contracts.MonthPeriodSummary{
		SpendingPerDay:      -1 * sumAmounts(spends) / int(duration.Hours()/24),
		SmallSpendingPerDay: -1 * sumAmounts(smallSpends) / int(duration.Hours()/24),
	}
}
