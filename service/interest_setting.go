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

type InterestSetting interface {
	Create(ctx context.Context, req in.InterestSettingRequest) (res out.InterestSettingResponse, err error)
	FindAll(ctx context.Context) (res []out.InterestSettingResponse, err error)
	FindById(ctx context.Context, id int64) (res out.InterestSettingResponse, err error)
	Update(ctx context.Context, req in.InterestSettingRequest) (res out.InterestSettingResponse, err error)
	ChangeStatusById(ctx context.Context, req in.InterestSettingRequest) (res out.InterestSettingResponse, err error)
	DeleteById(ctx context.Context, id int64) (err error)
}

type interestSetting struct {
	interestSettingRepo repository.InterestSetting
}

func NewInterestSetting() InterestSetting {
	return &interestSetting{
		interestSettingRepo: repository.NewInterestSetting(),
	}
}

func (service interestSetting) Create(ctx context.Context, req in.InterestSettingRequest) (res out.InterestSettingResponse, err error) {
	userSession := helper.GetAuthSessionModel(ctx)
	interestSet, _ := service.interestSettingRepo.FindActiveDataByTenorMonths(ctx, req.TenorMonths)
	if interestSet.Id > 0 {
		err = exception.AlreadyExistWithMessage("Interest setting", "please disable the setting for the existing 4 month tenor.")
		return
	}

	interestSet = toInterestSetModel(req)
	interestSet.CreatedBy = userSession.UserId
	err = service.interestSettingRepo.Create(ctx, &interestSet)
	if err != nil {
		return
	}

	res = toInterestSetResponse(interestSet)
	return
}

func (service interestSetting) FindAll(ctx context.Context) (res []out.InterestSettingResponse, err error) {
	interestSets, err := service.interestSettingRepo.FindAll(ctx)
	if err != nil {
		return
	}

	for _, set := range interestSets {
		res = append(res, toInterestSetResponse(set))
	}

	return
}

func (service interestSetting) FindById(ctx context.Context, id int64) (res out.InterestSettingResponse, err error) {
	interestSet, err := service.interestSettingRepo.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = exception.NotFoundError("Interest settings")
		}
		return
	}

	res = toInterestSetResponse(interestSet)
	return
}

func (service interestSetting) Update(ctx context.Context, req in.InterestSettingRequest) (res out.InterestSettingResponse, err error) {
	userSession := helper.GetAuthSessionModel(ctx)
	interestSet, err := service.interestSettingRepo.FindById(ctx, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = exception.NotFoundError("Interest settings")
		}
		return
	}

	existingInterestSet, _ := service.interestSettingRepo.FindActiveOtherDataByTenorMonths(ctx, req.TenorMonths, req.Id)
	if existingInterestSet.Id > 0 {
		err = exception.AlreadyExistWithMessage("Interest setting", "please disable the setting for the existing 4 month tenor.")
		return
	}

	mergedInterestSet(&interestSet, req)
	interestSet.UpdatedBy = userSession.UserId

	err = service.interestSettingRepo.Update(ctx, &interestSet)

	res = toInterestSetResponse(interestSet)
	return
}

func (service interestSetting) ChangeStatusById(ctx context.Context, req in.InterestSettingRequest) (res out.InterestSettingResponse, err error) {
	userSession := helper.GetAuthSessionModel(ctx)
	interestSet, err := service.interestSettingRepo.FindById(ctx, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = exception.NotFoundError("Interest settings")
		}
		return
	}

	interestSet.IsActive = req.IsActive
	interestSet.UpdatedBy = userSession.UserId
	err = service.interestSettingRepo.Update(ctx, &interestSet)

	res = toInterestSetResponse(interestSet)
	return
}

func (service interestSetting) DeleteById(ctx context.Context, id int64) (err error) {
	_, err = service.interestSettingRepo.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = exception.NotFoundError("Interest settings")
		}
		return
	}

	return service.interestSettingRepo.DeleteById(ctx, id)
}

func toInterestSetModel(req in.InterestSettingRequest) model.InterestSetting {
	return model.InterestSetting{
		Id:           req.Id,
		TenorMonths:  req.TenorMonths,
		InterestRate: req.InterestRate,
		Description:  req.Description,
		IsActive:     req.IsActive,
	}
}

func toInterestSetResponse(inputModel model.InterestSetting) out.InterestSettingResponse {
	return out.InterestSettingResponse{
		Id:           inputModel.Id,
		TenorMonths:  inputModel.TenorMonths,
		InterestRate: inputModel.InterestRate,
		Description:  inputModel.Description,
		IsActive:     inputModel.IsActive,
	}
}

func mergedInterestSet(inputModel *model.InterestSetting, req in.InterestSettingRequest) {
	inputModel.TenorMonths = mergedOrNot(inputModel.TenorMonths, req.TenorMonths)
	inputModel.InterestRate = mergedOrNot(inputModel.InterestRate, req.InterestRate)
	inputModel.Description = mergedOrNot(inputModel.Description, req.Description)
	inputModel.IsActive = mergedOrNot(inputModel.IsActive, req.IsActive)
}
