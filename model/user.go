package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id         int64           `json:"id" gorm:"primaryKey"`
	Username   string          `json:"username"`
	Password   string          `json:"password"`
	IsAdmin    bool            `json:"is_admin"`
	CreatedBy  int64           `json:"created_by"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedBy  int64           `json:"updated_by"`
	UpdatedAt  time.Time       `json:"updated_at"`
	DeletedBy  int64           `json:"deleted_by"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CustomerId *int64          `json:"customer_id"`

	Customer *Customer `json:"customer,omitempty" gorm:"foreignKey:CustomerId;references:Id"`
}

func (u User) TableName() string {
	return "user"
}
