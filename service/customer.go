package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/lithammer/shortuuid/v4"
	"gorm.io/gorm"
	"kredi-plus.com/be/dto/in"
	"kredi-plus.com/be/dto/out"
	"kredi-plus.com/be/lib/exception"
	"kredi-plus.com/be/lib/helper"
	"kredi-plus.com/be/model"
	"kredi-plus.com/be/repository"
	"path/filepath"
	"time"
)

type Customer interface {
	Create(ctx context.Context, req in.CustomerRequest) (res out.CustomerDetailResponse, err error)
	GetList(ctx context.Context, req in.CustomerRequest) (res []out.CustomerResponse, totalData int64, err error)
	GetDetailById(ctx context.Context, customerId int64) (res out.CustomerDetailResponse, err error)
	Update(ctx context.Context, payload in.CustomerRequest) (res out.CustomerDetailResponse, err error)
	DeleteById(ctx context.Context, customerId int64) (err error)
}

type customer struct {
	customerRepo    repository.Customer
	userRepo        repository.User
	creditLimitRepo repository.CreditLimit
}

func NewCustomer() Customer {
	return &customer{
		customerRepo:    repository.NewCustomer(),
		userRepo:        repository.NewUser(),
		creditLimitRepo: repository.NewCreditLimit(),
	}
}

var customerFuncName = "customer"

func (service customer) Create(ctx context.Context, req in.CustomerRequest) (res out.CustomerDetailResponse, err error) {
	userSession := helper.GetAuthSessionModel(ctx)
	entity := toCustomerModel(req)
	entity.CreatedBy = userSession.UserId
	for i := range entity.CreditLimits {
		entity.CreditLimits[i].CreatedBy = userSession.UserId
	}

	var (
		identityCardFilename, selfiePhotoFilename string
		timeNow                                   = time.Now()
		uuidStr                                   = shortuuid.New()
	)

	identityCardFilename = fmt.Sprintf("%s_indentity_card_%d%s", uuidStr, timeNow.UnixNano(), filepath.Ext(req.IdentityCardFile.Filename))
	selfiePhotoFilename = fmt.Sprintf("%s_selfie_photo_%d%s", uuidStr, timeNow.UnixNano(), filepath.Ext(req.SelfiePhotoFile.Filename))

	entity.IdentityCardUrl, err = helper.ProcessFileUpload(req.IdentityCardFile, customerFuncName, identityCardFilename)
	if err != nil {
		return
	}

	entity.SelfiePhotoUrl, err = helper.ProcessFileUpload(req.SelfiePhotoFile, customerFuncName, selfiePhotoFilename)
	if err != nil {
		return
	}

	err = service.customerRepo.Create(ctx, &entity)
	if err != nil {
		err = errorHandlerForConstraint(err)
		return
	}

	if req.Username != "" && req.Password != "" {
		var pass string
		pass, err = helper.GenerateHashFromString(req.Password)
		if err != nil {
			return
		}
		err = service.userRepo.Create(ctx, &model.User{
			Username:   req.Username,
			Password:   pass,
			CustomerId: &entity.Id,
		})
		if err != nil {
			return
		}
	}

	res = toCustomerDetailResponse(entity)
	return
}

func (service customer) GetList(ctx context.Context, req in.CustomerRequest) (res []out.CustomerResponse, totalData int64, err error) {
	result, totalData, err := service.customerRepo.GetList(ctx, req)
	if err != nil {
		return
	}

	for _, v := range result {
		res = append(res, toCustomerResponse(v))
	}
	return
}

func (service customer) GetDetailById(ctx context.Context, customerId int64) (res out.CustomerDetailResponse, err error) {
	result, err := service.customerRepo.GetCustomerById(ctx, customerId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = exception.NotFoundError("Customer Id")
		}
		return
	}

	res = toCustomerDetailResponse(result)
	return
}

func (service customer) Update(ctx context.Context, req in.CustomerRequest) (res out.CustomerDetailResponse, err error) {
	// handle upload file
	var (
		identityCardFilename, selfiePhotoFilename string
		timeNow                                   = time.Now()
		uuidStr                                   = shortuuid.New()
		creditLimits                              []model.CreditLimit
		userSession                               = helper.GetAuthSessionModel(ctx)
		customerOnDB                              model.Customer
	)

	err = Transaction(ctx, func(ctx context.Context) (trxErr error) {
		customerOnDB, trxErr = service.customerRepo.GetCustomerById(ctx, req.ID)
		if trxErr != nil {
			if errors.Is(trxErr, gorm.ErrRecordNotFound) {
				trxErr = exception.NotFoundError("Customer Id")
			}
			return trxErr
		}

		mergedCustomerModel(&customerOnDB, req)
		customerOnDB.UpdatedBy = userSession.UserId

		if req.IdentityCardFile != nil {
			identityCardFilename = fmt.Sprintf("%s_indentity_card_%d%s", uuidStr, timeNow.UnixNano(), filepath.Ext(req.IdentityCardFile.Filename))
			customerOnDB.IdentityCardUrl, trxErr = helper.ProcessFileUpload(req.IdentityCardFile, customerFuncName, identityCardFilename)
			if trxErr != nil {
				return trxErr
			}
		}

		if req.SelfiePhotoFile != nil {
			selfiePhotoFilename = fmt.Sprintf("%s_selfie_photo_%d%s", uuidStr, timeNow.UnixNano(), filepath.Ext(req.SelfiePhotoFile.Filename))
			customerOnDB.SelfiePhotoUrl, trxErr = helper.ProcessFileUpload(req.SelfiePhotoFile, customerFuncName, selfiePhotoFilename)
			if trxErr != nil {
				return trxErr
			}
		}

		// validate credit limit
		mappingCreditLimits := make(map[int64]model.CreditLimit)
		for _, limit := range customerOnDB.CreditLimits {
			mappingCreditLimits[limit.TenorMonths] = limit
		}

		for _, limit := range req.CreditLimits {
			creditLimitOnDB := mappingCreditLimits[limit.TenorMonths]
			usedLimit := creditLimitOnDB.LimitAmount - creditLimitOnDB.AvailableLimit
			if creditLimitOnDB.Id > 0 {
				if creditLimitOnDB.LimitAmount > limit.LimitAmount && usedLimit > limit.LimitAmount {
					return exception.InvalidRequestWithMessage("credit limit", fmt.Sprintf("limit has been used by '%.2f', limit amount must be greater than used limit", usedLimit))
				} else {
					creditLimitOnDB.UpdatedBy = userSession.UserId
					creditLimitOnDB.LimitAmount = limit.LimitAmount
					creditLimitOnDB.AvailableLimit = limit.LimitAmount - usedLimit
					creditLimits = append(creditLimits, creditLimitOnDB)
				}
			} else {
				creditLimits = append(creditLimits, model.CreditLimit{
					TenorMonths:    limit.TenorMonths,
					LimitAmount:    limit.LimitAmount,
					AvailableLimit: limit.LimitAmount,
					CreatedBy:      userSession.UserId,
					CustomerId:     customerOnDB.Id,
				})
			}
		}

		// delete all credit limit by customer id
		trxErr = service.creditLimitRepo.DeleteByCustomerId(ctx, customerOnDB.Id)
		if trxErr != nil {
			return trxErr
		}

		// upsert credit limit
		trxErr = service.creditLimitRepo.CreateBulk(ctx, creditLimits)
		if trxErr != nil {
			return trxErr
		}

		customerOnDB.CreditLimits = creditLimits

		// do update
		trxErr = service.customerRepo.Update(ctx, &customerOnDB)
		if trxErr != nil {
			trxErr = errorHandlerForConstraint(trxErr)
			return trxErr
		}

		return
	})
	if err != nil {
		return
	}

	res = toCustomerDetailResponse(customerOnDB)
	return
}

func (service customer) DeleteById(ctx context.Context, customerId int64) (err error) {
	_, err = service.customerRepo.GetCustomerById(ctx, customerId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = exception.NotFoundError("Customer Id")
		}
		return
	}

	return service.customerRepo.DeleteById(ctx, customerId)
}

func toCustomerModel(req in.CustomerRequest) model.Customer {
	output := model.Customer{
		Id:             req.ID,
		NIK:            req.NIK,
		FullName:       req.FullName,
		LegalName:      req.LegalName,
		PlaceOfBirth:   req.PlaceOfBirth,
		DateOfBirth:    req.DateOfBirth,
		CustomerSalary: req.CustomerSalary,
	}

	for _, limit := range req.CreditLimits {
		output.CreditLimits = append(output.CreditLimits, model.CreditLimit{
			TenorMonths:    limit.TenorMonths,
			LimitAmount:    limit.LimitAmount,
			AvailableLimit: limit.LimitAmount,
		})
	}
	return output
}

func toCustomerResponse(inputModel model.Customer) out.CustomerResponse {
	output := out.CustomerResponse{
		ID:              inputModel.Id,
		NIK:             inputModel.NIK,
		FullName:        inputModel.FullName,
		LegalName:       inputModel.LegalName,
		PlaceOfBirth:    inputModel.PlaceOfBirth,
		DateOfBirth:     inputModel.DateOfBirth,
		CustomerSalary:  inputModel.CustomerSalary,
		IdentityCardUrl: inputModel.IdentityCardUrl,
		SelfiePhotoUrl:  inputModel.SelfiePhotoUrl,
	}

	return output
}

func mergedCustomerModel(dataOnDB *model.Customer, req in.CustomerRequest) {
	dataOnDB.NIK = mergedOrNot(dataOnDB.NIK, req.NIK)
	dataOnDB.FullName = mergedOrNot(dataOnDB.FullName, req.FullName)
	dataOnDB.LegalName = mergedOrNot(dataOnDB.LegalName, req.LegalName)
	dataOnDB.PlaceOfBirth = mergedOrNot(dataOnDB.PlaceOfBirth, req.PlaceOfBirth)
	if !req.DateOfBirth.IsZero() {
		dataOnDB.DateOfBirth = mergedOrNot(dataOnDB.DateOfBirth, req.DateOfBirth)
	}
	dataOnDB.CustomerSalary = mergedOrNot(dataOnDB.CustomerSalary, req.CustomerSalary)
}

func toCustomerDetailResponse(inputModel model.Customer) out.CustomerDetailResponse {
	output := out.CustomerDetailResponse{
		ID:              inputModel.Id,
		NIK:             inputModel.NIK,
		FullName:        inputModel.FullName,
		LegalName:       inputModel.LegalName,
		PlaceOfBirth:    inputModel.PlaceOfBirth,
		DateOfBirth:     inputModel.DateOfBirth,
		CustomerSalary:  inputModel.CustomerSalary,
		IdentityCardUrl: inputModel.IdentityCardUrl,
		SelfiePhotoUrl:  inputModel.SelfiePhotoUrl,
	}

	for _, limit := range inputModel.CreditLimits {
		output.CreditLimits = append(output.CreditLimits, out.CreditLimitResponse{
			TenorMonths:    limit.TenorMonths,
			LimitAmount:    limit.LimitAmount,
			AvailableLimit: limit.AvailableLimit,
		})
	}

	return output
}
