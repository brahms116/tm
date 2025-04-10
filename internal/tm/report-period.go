package tm

import (
	"context"
	"fmt"
	"time"
	"tm/internal/orm/model"
	"tm/pkg/contracts"
)

type summariseTransactionsRow struct {
	Earnings  int64
	Spendings int32
}

func (t *tm) ReportPeriod(
	ctx context.Context,
	start, end time.Time, u100 bool,
) (contracts.ReportResponse, error) {

	query := t.db.Model(&model.TmTransaction{}).Select(`
    sum(case when amount_cents > 0 then amount_cents else 0 end) as earnings,
    -1 * sum(case when amount_cents < 0 then amount_cents else 0 end) as spendings
  `).
		Where("date >= ? and date < ?", start, end)

	if u100 {
		query = query.Where("amount_cents > -10000 and amount_cents < 0")
	}
	var resultRow summariseTransactionsRow
	result := query.Scan(&resultRow)

	if result.Error != nil {
		return contracts.ReportResponse{}, fmt.Errorf("error retrieving summary: %w", result.Error)
	}

	summary := contracts.Summary{
		SpendingCents: int(resultRow.Spendings),
		EarningCents:  int(resultRow.Earnings),
	}
	summary.NetCents = summary.EarningCents - summary.SpendingCents

	topSpendings := []model.TmTransaction{}
	topSpendingsQuery := t.db.Where("date >= ? and date < ?", start, end).Order("amount_cents desc").Where("amount_cents < 0")
	if u100 {
		topSpendingsQuery = topSpendingsQuery.Where("amount_cents > -10000")
	}
	result = topSpendingsQuery.Find(&topSpendings)
	if result.Error != nil {
		return contracts.ReportResponse{}, fmt.Errorf("error retrieving top spendings: %w", result.Error)
	}

	topEarnings := []model.TmTransaction{}
	if !u100 {
		topEarningsQuery := t.db.Where("date >= ? and date < ?", start, end).Order("amount_cents desc").Where("amount_cents > 0")
		result = topEarningsQuery.Find(&topEarnings)
		if result.Error != nil {
			return contracts.ReportResponse{}, fmt.Errorf("error retrieving top earnings: %w", result.Error)
		}
	}

	return contracts.ReportResponse{
		Summary:      summary,
		TopSpendings: transactionsToContracts(topSpendings),
		TopEarnings:  transactionsToContracts(topEarnings),
	}, nil
}
