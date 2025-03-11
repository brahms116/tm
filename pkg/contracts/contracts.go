package contracts

type Transaction struct {
	Id          string  `json:"id"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	AmountCents int     `json:"amountCents"`
	Category    *string `json:"category,omitEmpty"`
}
type TimelineResponseItem struct {
	Month   string  `json:"month"`
	Summary Summary `json:"summary"`
}
type TimelineResponse struct {
	Items []TimelineResponseItem `json:"items"`
}
type ReportResponse struct {
	Summary      Summary       `json:"summary"`
	TopSpendings []Transaction `json:"topSpendings"`
	TopEarnings  []Transaction `json:"topEarnings"`
}
type Summary struct {
	SpendingCents int `json:"spendingCents"`
	EarningCents  int `json:"earningCents"`
	NetCents      int `json:"netCents"`
}
