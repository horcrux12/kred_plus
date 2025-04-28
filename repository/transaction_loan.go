package repository

import (
	"context"
	"gorm.io/gorm/clause"
	"kredi-plus.com/be/dto/in"
	"kredi-plus.com/be/model"
)

type TransactionLoan interface {
	Create(ctx context.Context, inputModel *model.TransactionLoan) (err error)
	Find(ctx context.Context, query in.TransactionLoanRequest) (res []model.TransactionLoan, totalData int64, err error)
	FindById(ctx context.Context, id int64) (res model.TransactionLoan, err error)
	Update(ctx context.Context, inputModel *model.TransactionLoan) (err error)
	DeleteById(ctx context.Context, id int64) (err error)
}

type transactionLoan struct {
}

func NewTransactionLoan() TransactionLoan {
	return &transactionLoan{}
}

func (repo transactionLoan) Create(ctx context.Context, inputModel *model.TransactionLoan) (err error) {
	tx := GetDB(ctx)
	return tx.Create(inputModel).Error
}

func (repo transactionLoan) Find(ctx context.Context, query in.TransactionLoanRequest) (res []model.TransactionLoan, totalData int64, err error) {
	tx := GetDB(ctx)

	stmt := tx.Model(&model.TransactionLoan{}).Preload("InterestSetting").Preload("Customer")

	if query.CustomerId > 0 {
		stmt = stmt.Where("customer_id = ?", query.CustomerId)
	}

	err = stmt.Count(&totalData).Error
	if err != nil {
		return
	}

	err = stmt.Limit(query.GetLimit()).Offset(query.GetOffset()).
		Find(&res).Error
	if err != nil {
		return
	}

	return
}

func (repo transactionLoan) FindById(ctx context.Context, id int64) (res model.TransactionLoan, err error) {
	tx, isTransaction := GetDBWithStatus(ctx)
	stmt := tx.Model(&model.TransactionLoan{}).Preload("InterestSetting").Preload("Customer")

	if isTransaction {
		stmt = stmt.Clauses(clause.Locking{
			Strength: "UPDATE",
			Options:  "NOWAIT",
		})
	}

	err = stmt.Where(&model.TransactionLoan{Id: id}).Take(&res).Error
	if err != nil {
		return
	}
	return
}

func (repo transactionLoan) Update(ctx context.Context, inputModel *model.TransactionLoan) (err error) {
	tx := GetDB(ctx)
	return tx.Save(inputModel).Error
}

func (repo transactionLoan) DeleteById(ctx context.Context, id int64) (err error) {
	tx := GetDB(ctx)
	return tx.Delete(&model.InterestSetting{Id: id}).Error
}
