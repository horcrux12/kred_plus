package out

type TransactionLoanResponse struct {
	Id                int64   `json:"id"`
	CustomerId        int64   `json:"customer_id"`
	CustomerName      string  `json:"customer_name"`
	ContractNumber    string  `json:"contract_number"`
	OtrPrice          float64 `json:"otr_price"`
	AdminFee          float64 `json:"admin_fee"`
	InterestAmount    float64 `json:"interest_amount"`
	InstallmentAmount float64 `json:"installment_amount"`
	TotalLoan         float64 `json:"total_loan"`
	TenorMonths       int64   `json:"tenor_months"`
	InterestRate      float64 `json:"interest_rate"`
	AssetName         string  `json:"asset_name"`
	Platform          string  `json:"platform"`
	IsStarted         bool    `json:"is_started"`
	IsAlreadyPaid     bool    `json:"is_already_paid"`
}
