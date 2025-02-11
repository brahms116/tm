package contracts

import "time"

type ReportResponse struct {
	Total TotalSpendings    `json:"total"`
	Net   int               `json:"net"`
	Weeks []WeeklySpendings `json:"weeks"`
}

type Spending struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	AmountCents int       `json:"amountCents"`
}

type TotalSpendings struct {
	Total          int `json:"total"`
	SmallSpendings int `json:"smallSpendings"`
}

type WeeklySpendings struct {
	WeekStart time.Time      `json:"weekStart"`
	Total     TotalSpendings `json:"total"`

	Spendings []Spending `json:"spendings"`
}
