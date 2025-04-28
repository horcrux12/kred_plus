package repository

import (
	"context"
	"gorm.io/gorm/clause"
	"kredi-plus.com/be/dto/in"
	"kredi-plus.com/be/model"
)

type Customer interface {
	Create(ctx context.Context, inputModel *model.Customer) (err error)
	GetList(ctx context.Context, query in.CustomerRequest) (res []model.Customer, totalData int64, err error)
	GetCustomerById(ctx context.Context, customerId int64) (res model.Customer, err error)
	Update(ctx context.Context, inputModel *model.Customer) (err error)
	DeleteById(ctx context.Context, customerId int64) (err error)
}

type customer struct {
}

func NewCustomer() Customer {
	return &customer{}
}

func (repo customer) Create(ctx context.Context, inputModel *model.Customer) (err error) {
	tx := GetDB(ctx)
	err = tx.Create(inputModel).Error
	if err != nil {
		return
	}
	return
}

func (repo customer) GetList(ctx context.Context, query in.CustomerRequest) (res []model.Customer, totalData int64, err error) {
	tx := GetDB(ctx)
	stmt := tx.Model(&model.Customer{})

	err = stmt.Count(&totalData).Error
	if err != nil {
		return
	}

	err = stmt.Limit(query.GetLimit()).Offset(query.GetOffset()).Find(&res).Error
	if err != nil {
		return
	}
	return
}

func (repo customer) GetCustomerById(ctx context.Context, customerId int64) (res model.Customer, err error) {
	tx, isTransaction := GetDBWithStatus(ctx)
	stmt := tx.Model(&model.Customer{}).Preload("CreditLimits")

	if isTransaction {
		stmt = stmt.Clauses(clause.Locking{
			Strength: "UPDATE",
			Options:  "NOWAIT",
		})
	}
	err = stmt.Where(&model.Customer{Id: customerId}).Take(&res).Error
	if err != nil {
		return
	}
	return
}

func (repo customer) Update(ctx context.Context, inputModel *model.Customer) (err error) {
	tx := GetDB(ctx)
	return tx.Save(&inputModel).Error
}

func (repo customer) DeleteById(ctx context.Context, customerId int64) (err error) {
	tx := GetDB(ctx)
	return tx.Delete(&model.Customer{Id: customerId}).Error
}
