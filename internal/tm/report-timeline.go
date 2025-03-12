package tm

import (
	"context"
	"fmt"
	"time"
	"tm/internal/data"
	"tm/pkg/contracts"
)

func (t *tm) ReportTimeline(
	ctx context.Context,
	start, end time.Time,
) (contracts.TimelineResponse, error) {
	timelineRows, err := data.New().YearlyTimeline(
		ctx,
		t.conn,
		data.YearlyTimelineParams{
			Date:   start,
			Date_2: end,
		},
	)
	if err != nil {
		return contracts.TimelineResponse{}, fmt.Errorf("error retrieving timeline data %w", err)
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
