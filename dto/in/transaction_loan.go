package in

type Platforms int

const (
	Ecommerce Platforms = iota + 1
	Dealer
	Market
)

type TransactionLoanRequest struct {
	AbstractRequest
	Id             int64     `json:"id"`
	CustomerId     int64     `json:"customer_id" validate:"required"`
	ContractNumber string    `json:"contract_number" validate:"required"`
	OtrPrice       float64   `json:"otr_price" validate:"required"`
	AdminFee       float64   `json:"admin_fee" validate:"required"`
	TenorMonths    int64     `json:"tenor_months"`
	InterestId     int64     `json:"interest_id" validate:"required"`
	AssetName      string    `json:"asset_name" validate:"required"`
	Platform       Platforms `json:"platform" validate:"required"`
}

func (in Platforms) String() string {
	mappingPlatforms := make(map[Platforms]string)
	mappingPlatforms[Ecommerce] = "Ecommerce"
	mappingPlatforms[Dealer] = "Dealer"
	mappingPlatforms[Market] = "Market"
	return mappingPlatforms[in]
}
