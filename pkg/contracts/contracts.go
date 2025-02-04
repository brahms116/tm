package contracts

type ImportCsvResponse struct {
	Duplicates int `json:"duplicates"`
	Total      int `json:"total"`
}
