package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"kredi-plus.com/be/dto/app_model"
	"kredi-plus.com/be/dto/in"
	"kredi-plus.com/be/lib/exception"
	"kredi-plus.com/be/lib/helper"
	"kredi-plus.com/be/model"
	"kredi-plus.com/be/repository"
)

type Auth interface {
	Login(ctx context.Context, req in.LoginRequest) (token string, err error)
	Logout(ctx context.Context) (err error)
	GetUserById(ctx context.Context, userId int64) (res app_model.UserSession, err error)
}

type auth struct {
	userRepo repository.User
}

func NewAuth() Auth {
	return &auth{
		userRepo: repository.NewUser(),
	}
}

func (service auth) Login(ctx context.Context, req in.LoginRequest) (token string, err error) {
	userModel, err := service.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = exception.NotFoundError("username")
		}
		return
	}

	err = helper.CompareHashAndPassword(userModel.Password, req.Password)
	if err != nil {
		err = exception.ForbiddenAccess
		return
	}

	token, err = helper.GenerateJWT(userModel.Id, userModel.Username)
	return
}

func (service auth) Logout(ctx context.Context) (err error) {
	//TODO implement me
	panic("implement me")
}

func (service auth) GetUserById(ctx context.Context, userId int64) (res app_model.UserSession, err error) {
	result, err := service.userRepo.GetDetailById(ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = exception.NotFoundError("User Id")
		}
		return
	}

	res = toUserSession(result)
	return
}

func toUserSession(userModel model.User) app_model.UserSession {
	return app_model.UserSession{
		UserId:     userModel.Id,
		Username:   userModel.Username,
		IsAdmin:    userModel.IsAdmin,
		CustomerId: userModel.CustomerId,
	}
}
