package in

import (
	"mime/multipart"
	"time"
)

type CustomerRequest struct {
	AbstractRequest
	ID               int64                 `json:"id" form:"id"`
	NIK              string                `json:"nik" form:"nik" validate:"required,nik"`
	FullName         string                `json:"full_name" form:"full_name" validate:"required"`
	LegalName        string                `json:"legal_name" form:"legal_name"`
	PlaceOfBirth     string                `json:"place_of_birth" form:"place_of_birth"`
	DateOfBirthStr   string                `json:"date_of_birth" form:"date_of_birth"`
	DateOfBirth      time.Time             `json:"-"`
	CustomerSalary   float64               `json:"customer_salary" form:"customer_salary"`
	IdentityCardUrl  string                `json:"identity_card_url"`
	SelfiePhotoUrl   string                `json:"selfie_photo_url"`
	IdentityCardFile *multipart.FileHeader `json:"identity_card_file"`
	SelfiePhotoFile  *multipart.FileHeader `json:"selfie_photo_file"`
	Username         string                `json:"username" form:"username"`
	Password         string                `json:"password" form:"password"`
	CreditLimits     []CreditLimitDetail   `json:"credit_limits"`
}
