package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"kredi-plus.com/be/dto/in"
	"kredi-plus.com/be/dto/out"
	"kredi-plus.com/be/lib/exception"
	"kredi-plus.com/be/lib/helper"
	"kredi-plus.com/be/service"
	"net/http"
	"strconv"
)

type TransactionLoan interface {
	Create(w http.ResponseWriter, r *http.Request)
	Find(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	DeleteById(w http.ResponseWriter, r *http.Request)
}

type transactionLoan struct {
	transactionService service.TransactionLoan
}

func NewTransactionLoan() TransactionLoan {
	return &transactionLoan{
		transactionService: service.NewTransactionLoan(),
	}
}

func (handler transactionLoan) Create(w http.ResponseWriter, r *http.Request) {
	var (
		payload     in.TransactionLoanRequest
		userSession = helper.GetAuthSessionModel(r.Context())
	)

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequest, nil)
		return
	}

	if userSession.CustomerId != nil {
		payload.CustomerId = *userSession.CustomerId
	}

	err = helper.Validate(payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	res, err := handler.transactionService.Create(r.Context(), payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, res, "Success", http.StatusCreated, nil)
}

func (handler transactionLoan) Find(w http.ResponseWriter, r *http.Request) {
	var (
		userSession = helper.GetAuthSessionModel(r.Context())
	)

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	payload := in.TransactionLoanRequest{
		AbstractRequest: in.AbstractRequest{
			Page:    page,
			Limit:   limit,
			Search:  r.URL.Query().Get("search"),
			SortStr: r.URL.Query().Get("sort"),
		},
	}

	if userSession.CustomerId != nil {
		payload.CustomerId = *userSession.CustomerId
	}

	res, totalData, err := handler.transactionService.Find(r.Context(), payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	metaData := out.WebMetaData{
		TotalData: int(totalData),
		Page:      payload.Page,
		Limit:     payload.Limit,
	}

	helper.WriteSuccessResponse(w, res, "Success", http.StatusOK, &metaData)
}

func (handler transactionLoan) FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["Id"], 10, 64)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("transaction id", ""), nil)
		return
	}

	res, err := handler.transactionService.FindById(r.Context(), id)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, res, "Success", http.StatusOK, nil)
}

func (handler transactionLoan) Update(w http.ResponseWriter, r *http.Request) {
	var (
		payload     in.TransactionLoanRequest
		userSession = helper.GetAuthSessionModel(r.Context())
	)

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequest, nil)
		return
	}

	if userSession.CustomerId != nil {
		payload.CustomerId = *userSession.CustomerId
	}

	err = helper.Validate(payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["Id"], 10, 64)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("transaction id", ""), nil)
		return
	}
	payload.Id = id

	res, err := handler.transactionService.Update(r.Context(), payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, res, "Success", http.StatusCreated, nil)
}

func (handler transactionLoan) DeleteById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["Id"], 10, 64)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("transaction id", ""), nil)
		return
	}

	err = handler.transactionService.DeleteById(r.Context(), id)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, nil, "Success", http.StatusOK, nil)
}
