package handler

import (
	"encoding/json"
	"github.com/go-playground/form"
	"github.com/gorilla/mux"
	"kredi-plus.com/be/dto/in"
	"kredi-plus.com/be/dto/out"
	"kredi-plus.com/be/lib/exception"
	"kredi-plus.com/be/lib/helper"
	"kredi-plus.com/be/service"
	"net/http"
	"strconv"
	"time"
)

type Customer interface {
	CreateCustomer(w http.ResponseWriter, r *http.Request)
	GetListCustomer(w http.ResponseWriter, r *http.Request)
	GetCustomerById(w http.ResponseWriter, r *http.Request)
	UpdateCustomer(w http.ResponseWriter, r *http.Request)
	DeleteCustomerById(w http.ResponseWriter, r *http.Request)
}

type customer struct {
	customerService service.Customer
}

func NewCustomer() Customer {
	return &customer{
		customerService: service.NewCustomer(),
	}
}

func (handler customer) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		payload in.CustomerRequest
	)

	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	if err = form.NewDecoder().Decode(&payload, r.Form); err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	if _, payload.IdentityCardFile, err = r.FormFile("identity_card_file"); err != nil {
		helper.WriteErrorResponse(w, exception.MandatoryError("Identity Card"), nil)
		return
	}

	if _, payload.SelfiePhotoFile, err = r.FormFile("selfie_photo_file"); err != nil {
		helper.WriteErrorResponse(w, exception.MandatoryError("Selfie Photo"), nil)
		return
	}

	credits := r.Form["credit_limits"]
	for _, v := range credits {
		var tmpData in.CreditLimitDetail
		err = json.Unmarshal([]byte(v), &tmpData)
		if err != nil {
			helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("Credit limits", ""), nil)
			return
		}
		payload.CreditLimits = append(payload.CreditLimits, tmpData)
	}

	payload.DateOfBirth, err = time.Parse(helper.ConvertDateFormat("DD-MM-YYYY"), payload.DateOfBirthStr)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("date of birth", ""), nil)
		return
	}

	err = helper.Validate(payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	res, err := handler.customerService.Create(r.Context(), payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, res, "Success", http.StatusCreated, nil)
}

func (handler customer) GetListCustomer(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	payload := in.CustomerRequest{
		AbstractRequest: in.AbstractRequest{
			Page:    page,
			Limit:   limit,
			Search:  r.URL.Query().Get("search"),
			SortStr: r.URL.Query().Get("sort"),
		},
	}

	res, totalData, err := handler.customerService.GetList(r.Context(), payload)
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

func (handler customer) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerId, err := strconv.ParseInt(params["Id"], 10, 64)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("customer Id", ""), nil)
		return
	}

	res, err := handler.customerService.GetDetailById(r.Context(), customerId)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, res, "Success", http.StatusOK, nil)
}

func (handler customer) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		payload in.CustomerRequest
	)

	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	if err = form.NewDecoder().Decode(&payload, r.Form); err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	_, payload.IdentityCardFile, _ = r.FormFile("identity_card_file")

	_, payload.SelfiePhotoFile, _ = r.FormFile("selfie_photo_file")

	payload.DateOfBirth, err = time.Parse(helper.ConvertDateFormat("DD-MM-YYYY"), payload.DateOfBirthStr)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("date of birth", ""), nil)
		return
	}

	credits := r.Form["credit_limits"]
	for _, v := range credits {
		var tmpData in.CreditLimitDetail
		err = json.Unmarshal([]byte(v), &tmpData)
		if err != nil {
			helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("Credit limits", ""), nil)
			return
		}
		payload.CreditLimits = append(payload.CreditLimits, tmpData)
	}

	params := mux.Vars(r)
	customerId, err := strconv.ParseInt(params["Id"], 10, 64)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("customer Id", ""), nil)
		return
	}
	payload.ID = customerId

	err = helper.Validate(payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	res, err := handler.customerService.Update(r.Context(), payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, res, "Success", http.StatusCreated, nil)
}

func (handler customer) DeleteCustomerById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerId, err := strconv.ParseInt(params["Id"], 10, 64)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("customer Id", ""), nil)
		return
	}

	err = handler.customerService.DeleteById(r.Context(), customerId)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, nil, "Success", http.StatusOK, nil)
}
