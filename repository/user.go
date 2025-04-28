package repository

import (
	"context"
	"kredi-plus.com/be/dto/in"
	"kredi-plus.com/be/model"
)

type User interface {
	Create(ctx context.Context, inputModel *model.User) (err error)
	GetList(ctx context.Context, query in.UserRequest) (res []model.User, totalData int64, err error)
	GetDetailById(ctx context.Context, userId int64) (res model.User, err error)
	Update(ctx context.Context, inputModel *model.User) (err error)
	DeleteById(ctx context.Context, userId int64) (err error)
	FindByUsername(ctx context.Context, username string) (res model.User, err error)
}

type user struct {
}

func NewUser() User {
	return &user{}
}

func (repo user) Create(ctx context.Context, inputModel *model.User) (err error) {
	tx := GetDB(ctx)

	err = tx.Create(inputModel).Error
	if err != nil {
		return
	}

	return
}

func (repo user) GetList(ctx context.Context, query in.UserRequest) (res []model.User, totalData int64, err error) {
	tx := GetDB(ctx)
	stmt := tx.Model(&model.User{})

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

func (repo user) GetDetailById(ctx context.Context, userId int64) (res model.User, err error) {
	tx := GetDB(ctx)
	err = tx.Model(&model.User{}).Where(&model.User{Id: userId}).Take(&res).Error
	if err != nil {
		return
	}
	return
}

func (repo user) Update(ctx context.Context, inputModel *model.User) (err error) {
	tx := GetDB(ctx)
	return tx.Save(&inputModel).Error
}

func (repo user) DeleteById(ctx context.Context, userId int64) (err error) {
	tx := GetDB(ctx)
	return tx.Delete(&model.User{Id: userId}).Error
}

func (repo user) FindByUsername(ctx context.Context, username string) (res model.User, err error) {
	tx := GetDB(ctx)
	err = tx.Model(&model.User{}).Where(&model.User{Username: username}).Take(&res).Error
	if err != nil {
		return
	}
	return
}
