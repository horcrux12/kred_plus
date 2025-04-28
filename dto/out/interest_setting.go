package out

type InterestSettingResponse struct {
	Id           int64   `json:"id"`
	TenorMonths  int64   `json:"tenor_months"`
	InterestRate float64 `json:"interest_rate"`
	Description  string  `json:"description"`
	IsActive     bool    `json:"is_active"`
}
