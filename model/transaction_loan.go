package model

import (
	"gorm.io/gorm"
	"time"
)

type TransactionLoan struct {
	Id                int64           `json:"id" gorm:"primaryKey"`
	CustomerId        int64           `json:"customer_id"`
	ContractNumber    string          `json:"contract_number"`
	OtrPrice          float64         `json:"otr_price"`
	AdminFee          float64         `json:"admin_fee"`
	InterestAmount    float64         `json:"interest_amount"`    //total bunga
	InstallmentAmount float64         `json:"installment_amount"` //cicilan per-bulan
	TotalLoan         float64         `json:"total_loan"`         //total yang dicicil
	TenorMonths       int64           `json:"tenor_months"`
	InterestId        int64           `json:"interest_id"`
	AssetName         string          `json:"asset_name"`
	Platform          int64           `json:"platform"`
	IsStarted         bool            `json:"is_started"`
	IsAlreadyPaid     bool            `json:"is_already_paid"`
	CreatedBy         int64           `json:"created_by"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedBy         int64           `json:"updated_by"`
	UpdatedAt         time.Time       `json:"updated_at"`
	DeletedBy         int64           `json:"deleted_by"`
	DeletedAt         *gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	InterestSetting *InterestSetting `json:"interest_setting" gorm:"foreignKey:InterestId;references:Id"`
	Customer        *Customer        `json:"customer" gorm:"foreignKey:CustomerId;references:Id"`
}

func (t TransactionLoan) TableName() string {
	return "transaction_loan"
}
