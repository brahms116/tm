package contracts

import (
	"time"
	"tm/internal/data"
)

type MonthReport struct {
	Month             time.Time `json:"month"`
	Summary           SummaryOld   `json:"summary"`
	SummaryComparison SummaryOld   `json:"summaryComparison"`

	Periods []MonthPeriodReport `json:"periods"`
}

type SummaryOld struct {
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
