package tm

import (
	"context"
	"fmt"
	"time"
	"tm/internal/data"
	"tm/pkg/contracts"
)

func (t *tm) ReportPeriod(
	ctx context.Context,
	start, end time.Time, u100 bool,
) (contracts.ReportResponse, error) {

	var summary contracts.Summary
	if u100 {
		data, err := data.New().SummariseTransactionsU100(ctx, t.conn, data.SummariseTransactionsU100Params{
			Date:   start,
			Date_2: end,
		})
		if err != nil {
			return contracts.ReportResponse{}, fmt.Errorf("error retrieving summary: %w", err)
		}
		summary.SpendingCents = int(data.Spendings)
		summary.EarningCents = int(data.Earnings)
	} else {
		data, err := data.New().SummariseTransactions(ctx, t.conn, data.SummariseTransactionsParams{
			Date:   start,
			Date_2: end,
		})
		if err != nil {
			return contracts.ReportResponse{}, fmt.Errorf("error retrieving summary: %w", err)
		}
		summary.SpendingCents = int(data.Spendings)
		summary.EarningCents = int(data.Earnings)
	}
	summary.NetCents = summary.EarningCents - summary.SpendingCents

	topSpendings := []data.TmTransaction{}
	topEarnings := []data.TmTransaction{}
	if u100 {
		spendings, err := data.New().TopSpendingsU100(ctx, t.conn, data.TopSpendingsU100Params{
			Date:   start,
			Date_2: end,
		})
		if err != nil {
			return contracts.ReportResponse{}, fmt.Errorf("error retrieving top spendings %w", err)
		}
		topSpendings = spendings
	} else {
		spendings, err := data.New().TopSpendings(ctx, t.conn, data.TopSpendingsParams{
			Date:   start,
			Date_2: end,
		})
		if err != nil {
			return contracts.ReportResponse{}, fmt.Errorf("error retrieving top spendings %w", err)
		}
		topSpendings = spendings

		earnings, err := data.New().TopEarnings(ctx, t.conn, data.TopEarningsParams{
			Date:   start,
			Date_2: end,
		})
		if err != nil {
			return contracts.ReportResponse{}, fmt.Errorf("error retrieving top earnings %w", err)
		}
		topEarnings = earnings
	}

	return contracts.ReportResponse{
		Summary:      summary,
		TopSpendings: transactionsToContracts(topSpendings),
		TopEarnings:  transactionsToContracts(topEarnings),
	}, nil
}
