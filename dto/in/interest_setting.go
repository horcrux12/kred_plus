package in

type InterestSettingRequest struct {
	Id           int64   `json:"id"`
	TenorMonths  int64   `json:"tenor_months" validate:"required"`
	InterestRate float64 `json:"interest_rate" validate:"required"`
	Description  string  `json:"description"`
	IsActive     bool    `json:"is_active"`
}
