package repository

import (
	"context"
	"gorm.io/gorm/clause"
	"kredi-plus.com/be/model"
)

type InterestSetting interface {
	Create(ctx context.Context, inputModel *model.InterestSetting) (err error)
	FindAll(ctx context.Context) (res []model.InterestSetting, err error)
	FindActiveDataByTenorMonths(ctx context.Context, tenorMonths int64) (res model.InterestSetting, err error)
	FindById(ctx context.Context, id int64) (res model.InterestSetting, err error)
	Update(ctx context.Context, inputModel *model.InterestSetting) (err error)
	DeleteById(ctx context.Context, id int64) (err error)
	FindActiveOtherDataByTenorMonths(ctx context.Context, tenorMonths, existingId int64) (res model.InterestSetting, err error)
}

type interestSetting struct {
}

func NewInterestSetting() InterestSetting {
	return &interestSetting{}
}

func (repo interestSetting) Create(ctx context.Context, inputModel *model.InterestSetting) (err error) {
	tx := GetDB(ctx)
	return tx.Create(inputModel).Error
}

func (repo interestSetting) FindAll(ctx context.Context) (res []model.InterestSetting, err error) {
	tx := GetDB(ctx)
	err = tx.Model(&model.InterestSetting{}).Find(&res).Error
	if err != nil {
		return
	}
	return
}

func (repo interestSetting) FindActiveDataByTenorMonths(ctx context.Context, tenorMonths int64) (res model.InterestSetting, err error) {
	tx, isTransaction := GetDBWithStatus(ctx)
	stmt := tx.Model(&model.InterestSetting{})

	if isTransaction {
		stmt = stmt.Clauses(clause.Locking{
			Strength: "UPDATE",
			Options:  "NOWAIT",
		})
	}

	err = stmt.Where(&model.InterestSetting{TenorMonths: tenorMonths, IsActive: true}).Take(&res).Error
	if err != nil {
		return
	}
	return
}

func (repo interestSetting) FindById(ctx context.Context, id int64) (res model.InterestSetting, err error) {
	tx, isTransaction := GetDBWithStatus(ctx)
	stmt := tx.Model(&model.InterestSetting{})

	if isTransaction {
		stmt = stmt.Clauses(clause.Locking{
			Strength: "UPDATE",
			Options:  "NOWAIT",
		})
	}

	err = stmt.Where(&model.InterestSetting{Id: id}).Take(&res).Error
	if err != nil {
		return
	}
	return
}

func (repo interestSetting) Update(ctx context.Context, inputModel *model.InterestSetting) (err error) {
	tx := GetDB(ctx)
	return tx.Save(inputModel).Error
}

func (repo interestSetting) DeleteById(ctx context.Context, id int64) (err error) {
	tx := GetDB(ctx)
	return tx.Delete(&model.InterestSetting{Id: id}).Error
}

func (repo interestSetting) FindActiveOtherDataByTenorMonths(ctx context.Context, tenorMonths, existingId int64) (res model.InterestSetting, err error) {
	tx, isTransaction := GetDBWithStatus(ctx)
	stmt := tx.Model(&model.InterestSetting{})

	if isTransaction {
		stmt = stmt.Clauses(clause.Locking{
			Strength: "UPDATE",
			Options:  "NOWAIT",
		})
	}

	err = stmt.Where(&model.InterestSetting{TenorMonths: tenorMonths, IsActive: true}).
		Where("id != ?", existingId).Take(&res).Error
	if err != nil {
		return
	}
	return
}
