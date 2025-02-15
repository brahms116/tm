package contracts

import (
	"time"
	"tm/internal/data"
)

type MonthReport struct {
	Month             time.Time `json:"month"`
	Summary           Summary   `json:"summary"`
	SummaryComparison Summary   `json:"summaryComparison"`

	Periods []MonthPeriodReport `json:"periods"`
}

type Summary struct {
	Spending      int `json:"spending"`
	SmallSpending int `json:"smallSpending"`
	Earning       int `json:"earning"`
	Net           int `json:"net"`
}

type MonthPeriodSummary struct {
	SpendingPerDay      int `json:"spendingPerDay"`
	SmallSpendingPerDay int `json:"smallSpendingPerDay"`
}

type MonthPeriodReport struct {
	StartDate   time.Time            `json:"startDate"`
	EndDate     time.Time            `json:"endDate"`
	Summary     MonthPeriodSummary   `json:"summary"`
	SmallSpends []data.TmTransaction `json:"smallSpends"`
}
