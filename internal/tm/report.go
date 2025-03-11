package tm

import (
	"context"
	"fmt"
	"time"
	"tm/internal/data"
	"tm/pkg/contracts"
)

const SMALL_SPEND_THRESHOLD = -15000

func (t *tm) ReportText(ctx context.Context, dateMonth time.Time) (string, error) {
	report, err := t.Report(ctx, dateMonth)
	if err != nil {
		return "", fmt.Errorf("error generating report: %w", err)
	}
	return textReport(report), nil
}

func (t *tm) Report(ctx context.Context, dateMonth time.Time) (contracts.MonthReport, error) {
	currStart, currEnd := monthSpan(dateMonth)
	params := data.ListTransactionsParams{
		Date:   currStart,
		Date_2: currEnd,
	}

	curr, err := data.New().ListTransactions(ctx, t.conn, params)
	if err != nil {
		return contracts.MonthReport{}, fmt.Errorf("error listing transactions: %w", err)
	}

	prevStart, prevEnd := monthSpan(dateMonth.AddDate(0, -1, 0))
	params = data.ListTransactionsParams{
		Date:   prevStart,
		Date_2: prevEnd,
	}
	prev, err := data.New().ListTransactions(ctx, t.conn, params)
	if err != nil {
		return contracts.MonthReport{}, fmt.Errorf("error listing transactions: %w", err)
	}

	report := generateMonthReport(curr, prev, currStart, currEnd)
	return report, nil
}

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

func summary(transactions []data.TmTransaction) contracts.SummaryOld {
	spends := spends(transactions)
	smallSpends := smallSpends(transactions)
	earns := earns(transactions)
	return contracts.SummaryOld{
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
func generateMonthReport(curr, prev []data.TmTransaction, monthStart, monthEnd time.Time) contracts.MonthReport {
	currSummary := summary(curr)
	prevSummary := summary(prev)
	diffSummary := contracts.SummaryOld{
		Spending:      currSummary.Spending - prevSummary.Spending,
		SmallSpending: currSummary.SmallSpending - prevSummary.SmallSpending,
		Earning:       currSummary.Earning - prevSummary.Earning,
		Net:           currSummary.Net - prevSummary.Net,
	}

	periods := monthlyPeriodReports(curr, monthStart, monthEnd)
	return contracts.MonthReport{
		Month:             monthStart,
		Summary:           currSummary,
		SummaryComparison: diffSummary,
		Periods:           periods,
	}
}

func nextMonthPeriodEnd(monthStart, currPeriodEnd time.Time) time.Time {
	var prevMonday time.Time
	if currPeriodEnd.Weekday() == 0 {
		prevMonday = currPeriodEnd.AddDate(0, 0, -7)
	} else {
		prevMonday = currPeriodEnd.AddDate(0, 0, -int(currPeriodEnd.Weekday()))
	}

	if prevMonday.After(monthStart) {
		return prevMonday
	}
	return monthStart
}

func monthlyPeriodReports(trans []data.TmTransaction, monthStart, periodEnd time.Time) []contracts.MonthPeriodReport {
	nextPeriodEnd := nextMonthPeriodEnd(monthStart, periodEnd)

	currTs, nextTs := splitTransactionsByTime(trans, nextPeriodEnd)
	currSummary := monthPeriodSummary(currTs, periodEnd.Sub(nextPeriodEnd))
	currSmallSpends := smallSpends(currTs)
	currReport := contracts.MonthPeriodReport{
		StartDate:   nextPeriodEnd,
		EndDate:     periodEnd,
		Summary:     currSummary,
		SmallSpends: currSmallSpends,
	}

	if nextPeriodEnd.After(monthStart) {
		return append([]contracts.MonthPeriodReport{currReport}, monthlyPeriodReports(nextTs, monthStart, nextPeriodEnd)...)
	}
	return []contracts.MonthPeriodReport{currReport}
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
