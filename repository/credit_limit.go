package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm/clause"
	"kredi-plus.com/be/model"
)

type CreditLimit interface {
	CreateBulk(ctx context.Context, inputModels []model.CreditLimit) (err error)
	FindByCustomerId(ctx context.Context, customerId int64) (res []model.CreditLimit, err error)
	FindById(ctx context.Context, id int64) (res model.CreditLimit, err error)
	FindByCondition(ctx context.Context, query model.CreditLimit) (res model.CreditLimit, err error)
	SaveAll(ctx context.Context, inputModels []model.CreditLimit) (err error)
	DeleteByCustomerId(ctx context.Context, customerId int64) (err error)
}

type creditLimit struct {
}

func NewCreditLimit() CreditLimit {
	return &creditLimit{}
}

func (repo creditLimit) CreateBulk(ctx context.Context, inputModels []model.CreditLimit) (err error) {
	tx := GetDB(ctx)
	err = tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "customer_id"}, {Name: "tenor_months"}},
		UpdateAll: true,
	}).Create(&inputModels).Error

	if err != nil {
		fmt.Print("Err : ", err)
		return
	}
	return
}

func (repo creditLimit) FindByCustomerId(ctx context.Context, customerId int64) (res []model.CreditLimit, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo creditLimit) FindById(ctx context.Context, id int64) (res model.CreditLimit, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo creditLimit) FindByCondition(ctx context.Context, query model.CreditLimit) (res model.CreditLimit, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo creditLimit) SaveAll(ctx context.Context, inputModels []model.CreditLimit) (err error) {
	//TODO implement me
	panic("implement me")
}

func (repo creditLimit) DeleteByCustomerId(ctx context.Context, customerId int64) (err error) {
	tx := GetDB(ctx)

	return tx.Model(&model.CreditLimit{}).
		Where(&model.CreditLimit{CustomerId: customerId}).
		Delete(&model.CreditLimit{}).Error
}
