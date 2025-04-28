package model

import (
	"gorm.io/gorm"
	"time"
)

type InterestSetting struct {
	Id           int64           `json:"id" gorm:"primaryKey"`
	TenorMonths  int64           `json:"tenor_months"`
	InterestRate float64         `json:"interest_rate"`
	Description  string          `json:"description"`
	IsActive     bool            `json:"is_active"`
	CreatedBy    int64           `json:"created_by"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedBy    int64           `json:"updated_by"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedBy    int64           `json:"deleted_by"`
	DeletedAt    *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (i InterestSetting) TableName() string {
	return "interest_setting"
}
