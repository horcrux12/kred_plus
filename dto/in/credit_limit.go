package in

type CreditLimitDetail struct {
	TenorMonths int64   `json:"tenor_months" validate:"required"`
	LimitAmount float64 `json:"limit_amount" validate:"required"`
}
