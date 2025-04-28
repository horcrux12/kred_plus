package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"kredi-plus.com/be/dto/in"
	"kredi-plus.com/be/dto/out"
	"kredi-plus.com/be/lib/exception"
	"kredi-plus.com/be/lib/helper"
	"kredi-plus.com/be/model"
	"kredi-plus.com/be/repository"
)

type User interface {
	Create(ctx context.Context, req in.UserRequest) (res out.UserResponse, err error)
	GetList(ctx context.Context, req in.UserRequest) (res []out.UserResponse, totalData int64, err error)
	GetDetailById(ctx context.Context, userId int64) (res out.UserResponse, err error)
	Update(ctx context.Context, req in.UserRequest) (res out.UserResponse, err error)
	DeleteById(ctx context.Context, userId int64) (err error)
}

type user struct {
	userRepo     repository.User
	customerRepo repository.Customer
}

func NewUser() User {
	return &user{
		userRepo:     repository.NewUser(),
		customerRepo: repository.NewCustomer(),
	}
}

func (service user) Create(ctx context.Context, req in.UserRequest) (res out.UserResponse, err error) {
	userSession := helper.GetAuthSessionModel(ctx)
	entity := toUserModel(req)
	entity.CreatedBy = userSession.UserId

	entity.Password, err = helper.GenerateHashFromString(entity.Password)
	if err != nil {
		err = errors.New("Error when generate hash password")
		return
	}

	if req.CustomerId > 0 {
		var customerOnDB model.Customer
		customerOnDB, err = service.customerRepo.GetCustomerById(ctx, req.CustomerId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = exception.NotFoundError("Customer Id")
			}
			return
		}
		entity.Customer = &customerOnDB
	}

	err = service.userRepo.Create(ctx, &entity)
	if err != nil {
		err = errorHandlerForConstraint(err)
		return
	}

	res = toUserResponse(entity)
	return
}

func (service user) GetList(ctx context.Context, req in.UserRequest) (res []out.UserResponse, totalData int64, err error) {
	result, totalData, err := service.userRepo.GetList(ctx, req)
	if err != nil {
		return
	}

	for _, v := range result {
		res = append(res, toUserResponse(v))
	}
	return
}

func (service user) GetDetailById(ctx context.Context, userId int64) (res out.UserResponse, err error) {
	result, err := service.userRepo.GetDetailById(ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = exception.NotFoundError("User Id")
		}
		return
	}

	res = toUserResponse(result)
	return
}

func (service user) Update(ctx context.Context, req in.UserRequest) (res out.UserResponse, err error) {
	userSession := helper.GetAuthSessionModel(ctx)
	userModel, err := service.userRepo.GetDetailById(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = exception.NotFoundError("User Id")
		}
		return
	}

	if req.Password != "" {
		req.Password, err = helper.GenerateHashFromString(req.Password)
		if err != nil {
			err = errors.New("Error when generate hash password")
			return
		}
	}

	mergeUserModel(&userModel, req)
	userModel.UpdatedBy = userSession.UserId

	err = service.userRepo.Update(ctx, &userModel)
	if err != nil {
		err = errorHandlerForConstraint(err)
		return
	}

	res = toUserResponse(userModel)
	return
}

func (service user) DeleteById(ctx context.Context, userId int64) (err error) {
	_, err = service.userRepo.GetDetailById(ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = exception.NotFoundError("User Id")
		}
		return
	}

	return service.userRepo.DeleteById(ctx, userId)
}

func toUserModel(payload in.UserRequest) model.User {
	output := model.User{
		Id:       payload.ID,
		Username: payload.Username,
		Password: payload.Password,
		IsAdmin:  payload.IsAdmin,
	}

	if payload.CustomerId > 0 {
		output.CustomerId = &payload.CustomerId
	}

	return output
}

func toUserResponse(userModel model.User) out.UserResponse {
	res := out.UserResponse{
		ID:       userModel.Id,
		Username: userModel.Username,
		IsAdmin:  userModel.IsAdmin,
	}

	if userModel.Customer != nil {
		res.CustomerName = userModel.Customer.FullName
	}

	return res
}

func mergeUserModel(existingData *model.User, requested in.UserRequest) {
	if existingData.Username != requested.Username {
		existingData.Username = requested.Username
	}

	if existingData.IsAdmin != requested.IsAdmin {
		existingData.IsAdmin = requested.IsAdmin
	}

	var customerID int64
	if existingData.CustomerId != nil {
		customerID = *existingData.CustomerId
	}

	if customerID != requested.CustomerId {
		existingData.CustomerId = &requested.CustomerId
	}

	if requested.Password != "" && existingData.Password != requested.Password {
		existingData.Password = requested.Password
	}
}
