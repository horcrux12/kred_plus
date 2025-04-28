package service

import (
	"context"
	"kredi-plus.com/be/dto/in"
	"kredi-plus.com/be/dto/out"
	"kredi-plus.com/be/lib/exception"
	"kredi-plus.com/be/lib/helper"
	"kredi-plus.com/be/model"
	"kredi-plus.com/be/repository"
)

type TransactionLoan interface {
	Create(ctx context.Context, req in.TransactionLoanRequest) (res out.TransactionLoanResponse, err error)
	Find(ctx context.Context, req in.TransactionLoanRequest) (res []out.TransactionLoanResponse, totalData int64, err error)
	FindById(ctx context.Context, id int64) (res out.TransactionLoanResponse, err error)
	Update(ctx context.Context, req in.TransactionLoanRequest) (res out.TransactionLoanResponse, err error)
	DeleteById(ctx context.Context, id int64) (err error)
}

type transactionLoan struct {
	transactionRepo repository.TransactionLoan
	creditLimitRepo repository.CreditLimit
	customerRepo    repository.Customer
	interestSetRepo repository.InterestSetting
}

func NewTransactionLoan() TransactionLoan {
	return &transactionLoan{
		transactionRepo: repository.NewTransactionLoan(),
		creditLimitRepo: repository.NewCreditLimit(),
		customerRepo:    repository.NewCustomer(),
		interestSetRepo: repository.NewInterestSetting(),
	}
}

func (service transactionLoan) Create(ctx context.Context, req in.TransactionLoanRequest) (res out.TransactionLoanResponse, err error) {
	userSession := helper.GetAuthSessionModel(ctx)
	//check customer is exist
	_, err = service.customerRepo.GetCustomerById(ctx, req.CustomerId)
	if err != nil {
		err = checkErrorGormNotFound(err, "Customer")
		return
	}

	interestSet, err := service.interestSetRepo.FindById(ctx, req.InterestId)
	if err != nil {
		err = checkErrorGormNotFound(err, "Interest Setting")
		return
	}

	if req.TenorMonths == 0 {
		req.TenorMonths = interestSet.TenorMonths
	}

	if interestSet.TenorMonths != req.TenorMonths {
		err = exception.InvalidRequestWithMessage("Tenor months", "There is not match tenor months and interest setting")
		return
	}

	transactionModel := toTransactionLoanModel(req)
	err = Transaction(ctx, func(ctx context.Context) (trxErr error) {
		creditLimit := model.CreditLimit{
			CustomerId:  req.CustomerId,
			TenorMonths: req.TenorMonths,
		}
		// get credit limits
		creditLimit, trxErr = service.creditLimitRepo.FindByCondition(ctx, creditLimit)
		if trxErr != nil {
			trxErr = checkErrorGormNotFound(trxErr, "Credit Limit")
			return
		}

		// validate whether the credit limit is available
		if creditLimit.AvailableLimit < (transactionModel.OtrPrice + transactionModel.AdminFee) {
			trxErr = exception.InvalidRequestWithMessage("Transaction", "your credit limit is low")
			return
		}

		creditLimit.AvailableLimit -= (transactionModel.OtrPrice + transactionModel.AdminFee)

		// update credit limit
		trxErr = service.creditLimitRepo.Update(ctx, &creditLimit)
		if trxErr != nil {
			return
		}

		// insert transaction
		transactionModel.CreatedBy = userSession.UserId
		return service.transactionRepo.Create(ctx, &transactionModel)
	})
	transactionModel.InterestSetting = &interestSet

	res = toTransactionResponse(transactionModel)
	if err != nil {
		return
	}
	return
}

func (service transactionLoan) Find(ctx context.Context, req in.TransactionLoanRequest) (res []out.TransactionLoanResponse, totalData int64, err error) {
	userSession := helper.GetAuthSessionModel(ctx)
	if userSession.CustomerId != nil {
		req.CustomerId = *userSession.CustomerId
	}

	result, totalData, err := service.transactionRepo.Find(ctx, req)
	for _, loan := range result {
		res = append(res, toTransactionResponse(loan))
	}
	return
}

func (service transactionLoan) FindById(ctx context.Context, id int64) (res out.TransactionLoanResponse, err error) {
	result, err := service.transactionRepo.FindById(ctx, id)
	if err != nil {
		err = checkErrorGormNotFound(err, "Transaction")
		return
	}

	res = toTransactionResponse(result)
	return
}

func (service transactionLoan) Update(ctx context.Context, req in.TransactionLoanRequest) (res out.TransactionLoanResponse, err error) {
	userSession := helper.GetAuthSessionModel(ctx)
	//check customer is exist
	_, err = service.customerRepo.GetCustomerById(ctx, req.CustomerId)
	if err != nil {
		err = checkErrorGormNotFound(err, "Customer")
		return
	}

	interestSet, err := service.interestSetRepo.FindById(ctx, req.InterestId)
	if err != nil {
		err = checkErrorGormNotFound(err, "Interest Setting")
		return
	}

	if req.TenorMonths == 0 {
		req.TenorMonths = interestSet.TenorMonths
	}

	if interestSet.TenorMonths != req.TenorMonths {
		err = exception.InvalidRequestWithMessage("Tenor months", "There is not match tenor months and interest setting")
		return
	}

	err = Transaction(ctx, func(ctx context.Context) (trxErr error) {
		// get transaction By ID
		transactionModel := toTransactionLoanModel(req)
		transactionOnDB, trxErr := service.transactionRepo.FindById(ctx, req.Id)
		if trxErr != nil {
			trxErr = checkErrorGormNotFound(trxErr, "Transaction")
			return
		}

		if transactionOnDB.IsStarted {
			trxErr = exception.InvalidRequestWithMessage("Transaction", "Cannot update this transaction, it has started")
			return
		}

		creditLimit := model.CreditLimit{
			CustomerId:  req.CustomerId,
			TenorMonths: req.TenorMonths,
		}
		// get credit limits
		creditLimit, trxErr = service.creditLimitRepo.FindByCondition(ctx, creditLimit)
		if trxErr != nil {
			trxErr = checkErrorGormNotFound(trxErr, "Credit Limit")
			return
		}

		if transactionOnDB.TenorMonths != req.TenorMonths || transactionOnDB.CustomerId != req.CustomerId {
			creditLimitOnDB := model.CreditLimit{
				CustomerId:  transactionOnDB.CustomerId,
				TenorMonths: transactionOnDB.TenorMonths,
			}

			creditLimitOnDB, trxErr = service.creditLimitRepo.FindByCondition(ctx, creditLimitOnDB)
			if trxErr != nil {
				trxErr = checkErrorGormNotFound(trxErr, "Credit Limit")
				return
			}

			creditLimitOnDB.AvailableLimit += (transactionOnDB.OtrPrice + transactionOnDB.AdminFee)
			// update existing limit
			trxErr = service.creditLimitRepo.Update(ctx, &creditLimitOnDB)
			if trxErr != nil {
				return
			}
		} else {
			creditLimit.AvailableLimit += (transactionOnDB.OtrPrice + transactionOnDB.AdminFee)
		}

		// validate whether the credit limit is available
		if creditLimit.AvailableLimit < (transactionModel.OtrPrice + transactionModel.AdminFee) {
			trxErr = exception.InvalidRequestWithMessage("Transaction", "your credit limit is low")
			return
		}

		creditLimit.AvailableLimit -= (transactionModel.OtrPrice + transactionModel.AdminFee)

		// update credit limit
		trxErr = service.creditLimitRepo.Update(ctx, &creditLimit)
		if trxErr != nil {
			return
		}

		// update transaction
		mergedTransaction(&transactionOnDB, transactionModel)
		transactionOnDB.UpdatedBy = userSession.UserId
		trxErr = service.transactionRepo.Update(ctx, &transactionOnDB)
		if trxErr != nil {
			return
		}

		transactionOnDB.InterestSetting = &interestSet

		res = toTransactionResponse(transactionOnDB)
		return
	})

	if err != nil {
		return
	}
	return
}

func (service transactionLoan) DeleteById(ctx context.Context, id int64) (err error) {
	transactionOnDB, trxErr := service.transactionRepo.FindById(ctx, id)
	if trxErr != nil {
		trxErr = checkErrorGormNotFound(trxErr, "Transaction")
		return
	}

	if transactionOnDB.IsStarted {
		trxErr = exception.InvalidRequestWithMessage("Transaction", "Cannot delete this transaction, it has started")
		return
	}

	return service.DeleteById(ctx, id)
}

func toTransactionLoanModel(req in.TransactionLoanRequest) model.TransactionLoan {
	output := model.TransactionLoan{
		Id:             req.Id,
		CustomerId:     req.CustomerId,
		ContractNumber: req.ContractNumber,
		OtrPrice:       req.OtrPrice,
		AdminFee:       req.AdminFee,
		TenorMonths:    req.TenorMonths,
		InterestId:     req.InterestId,
		AssetName:      req.AssetName,
		Platform:       int64(req.Platform),
	}

	calculateInstallments(&output)
	return output
}

func toTransactionResponse(inputModel model.TransactionLoan) out.TransactionLoanResponse {
	output := out.TransactionLoanResponse{
		Id:                inputModel.Id,
		CustomerId:        inputModel.CustomerId,
		ContractNumber:    inputModel.ContractNumber,
		OtrPrice:          inputModel.OtrPrice,
		AdminFee:          inputModel.AdminFee,
		InterestAmount:    inputModel.InterestAmount,
		InstallmentAmount: inputModel.InstallmentAmount,
		TotalLoan:         inputModel.TotalLoan,
		TenorMonths:       inputModel.TenorMonths,
		AssetName:         inputModel.AssetName,
		Platform:          in.Platforms(inputModel.Platform).String(),
		IsStarted:         inputModel.IsStarted,
		IsAlreadyPaid:     inputModel.IsAlreadyPaid,
	}

	if inputModel.Customer != nil {
		output.CustomerName = inputModel.Customer.FullName
	}

	if inputModel.InterestSetting != nil {
		output.InterestRate = inputModel.InterestSetting.InterestRate
	}

	return output
}

func calculateInstallments(inputModel *model.TransactionLoan) {
	if inputModel.InterestSetting == nil {
		return
	}

	rate := inputModel.InterestSetting.InterestRate
	inputModel.InterestAmount = (inputModel.OtrPrice + inputModel.AdminFee) * (rate / 100.0)
	inputModel.TotalLoan = inputModel.OtrPrice + inputModel.AdminFee + inputModel.InterestAmount
	inputModel.InstallmentAmount = inputModel.TotalLoan / float64(inputModel.TenorMonths)
}

func mergedTransaction(existing *model.TransactionLoan, req model.TransactionLoan) {
	existing.CustomerId = mergedOrNot(existing.CustomerId, req.CustomerId)
	existing.ContractNumber = mergedOrNot(existing.ContractNumber, req.ContractNumber)
	existing.OtrPrice = mergedOrNot(existing.OtrPrice, req.OtrPrice)
	existing.AdminFee = mergedOrNot(existing.AdminFee, req.AdminFee)
	existing.InterestAmount = mergedOrNot(existing.InterestAmount, req.InterestAmount)
	existing.InstallmentAmount = mergedOrNot(existing.InstallmentAmount, req.InstallmentAmount)
	existing.TotalLoan = mergedOrNot(existing.TotalLoan, req.TotalLoan)
	existing.TenorMonths = mergedOrNot(existing.TenorMonths, req.TenorMonths)
	existing.InterestId = mergedOrNot(existing.InterestId, req.InterestId)
	existing.AssetName = mergedOrNot(existing.AssetName, req.AssetName)
	existing.Platform = mergedOrNot(existing.Platform, req.Platform)
	existing.IsStarted = mergedOrNot(existing.IsStarted, req.IsStarted)
	existing.IsAlreadyPaid = mergedOrNot(existing.IsAlreadyPaid, req.IsAlreadyPaid)
}
