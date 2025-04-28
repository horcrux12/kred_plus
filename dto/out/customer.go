package out

import (
	"time"
)

type CustomerResponse struct {
	ID              int64     `json:"id"`
	NIK             string    `json:"nik"`
	FullName        string    `json:"full_name"`
	LegalName       string    `json:"legal_name"`
	PlaceOfBirth    string    `json:"place_of_birth"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	CustomerSalary  float64   `json:"customer_salary"`
	IdentityCardUrl string    `json:"identity_card_url"`
	SelfiePhotoUrl  string    `json:"selfie_photo_url"`
}

type CustomerDetailResponse struct {
	ID              int64                 `json:"id"`
	NIK             string                `json:"nik"`
	FullName        string                `json:"full_name"`
	LegalName       string                `json:"legal_name"`
	PlaceOfBirth    string                `json:"place_of_birth"`
	DateOfBirth     time.Time             `json:"date_of_birth"`
	CustomerSalary  float64               `json:"customer_salary"`
	IdentityCardUrl string                `json:"identity_card_url"`
	SelfiePhotoUrl  string                `json:"selfie_photo_url"`
	CreditLimits    []CreditLimitResponse `json:"credit_limits"`
}

type CreditLimitResponse struct {
	TenorMonths    int64   `json:"tenor_months"`
	LimitAmount    float64 `json:"limit_amount"`
	AvailableLimit float64 `json:"available_limit"`
}
