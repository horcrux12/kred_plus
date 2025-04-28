package model

import (
	"gorm.io/gorm"
	"time"
)

type CreditLimit struct {
	Id             int64           `json:"id" gorm:"primaryKey"`
	TenorMonths    int64           `json:"tenor_months"`
	LimitAmount    float64         `json:"limit_amount"`
	AvailableLimit float64         `json:"available_limit"`
	CreatedBy      int64           `json:"created_by"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedBy      int64           `json:"updated_by"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedBy      int64           `json:"deleted_by"`
	DeletedAt      *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CustomerId     int64           `json:"customer_id"`

	Customer *Customer `json:"customer,omitempty" gorm:"foreignKey:CustomerId;references:Id"`
}

func (c CreditLimit) TableName() string {
	return "credit_limit"
}
