package tm

import (
	"context"
	"fmt"
	"time"
	"tm/internal/orm/model"
	"tm/pkg/contracts"
)

type monthlyTimelineRow struct {
	Month     time.Time
	Earnings  int32
	Spendings int32
}

func (t *tm) ReportTimeline(
	ctx context.Context,
	start, end time.Time,
) (contracts.TimelineResponse, error) {
	var timelineRows []monthlyTimelineRow

	result := t.db.Model(&model.TmTransaction{}).
		Select(`
    date_trunc('month', date)::date as month,
    sum(case when amount_cents > 0 then amount_cents else 0 end)::int as earnings,
    (-1 * sum(case when amount_cents < 0 then amount_cents else 0 end))::int as spendings
  `).
		Where("date >= ? and date < ?", start, end).
		Group("month").
		Order("month asc").
		Scan(&timelineRows)

	if result.Error != nil {
		return contracts.TimelineResponse{}, fmt.Errorf("error retrieving timeline data %w", result.Error)
	}

	timelineItems := make([]contracts.TimelineResponseItem, len(timelineRows))
	for i := range timelineRows {
		earnings := timelineRows[i].Earnings
		spendings := timelineRows[i].Spendings
		net := earnings - spendings
		timelineItems[i] = contracts.TimelineResponseItem{
			Month: timelineRows[i].Month.Format(time.RFC3339),
			Summary: contracts.Summary{
				SpendingCents: int(spendings),
				EarningCents:  int(earnings),
				NetCents:      int(net),
			},
		}
	}

	return contracts.TimelineResponse{
		Items: timelineItems,
	}, nil
}
