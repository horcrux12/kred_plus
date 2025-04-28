package model

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	Id              int64           `json:"id" gorm:"primaryKey"`
	NIK             string          `json:"nik"`
	FullName        string          `json:"full_name"`
	LegalName       string          `json:"legal_name"`
	PlaceOfBirth    string          `json:"place_of_birth"`
	DateOfBirth     time.Time       `json:"date_of_birth"`
	CustomerSalary  float64         `json:"customer_salary"`
	IdentityCardUrl string          `json:"identity_card_url"`
	SelfiePhotoUrl  string          `json:"selfie_photo_url"`
	CreatedBy       int64           `json:"created_by"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedBy       int64           `json:"updated_by"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedBy       int64           `json:"deleted_by"`
	DeletedAt       *gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	User         *User         `json:"user,omitempty" gorm:"foreignKey:CustomerId;references:Id"`
	CreditLimits []CreditLimit `json:"credit_limits,omitempty" gorm:"foreignKey:CustomerId;references:Id"`
}

func (c Customer) TableName() string {
	return "customer"
}
