package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"kredi-plus.com/be/dto/in"
	"kredi-plus.com/be/lib/exception"
	"kredi-plus.com/be/lib/helper"
	"kredi-plus.com/be/service"
	"net/http"
	"strconv"
)

type InterestSetting interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	DeleteById(w http.ResponseWriter, r *http.Request)
	ChangeStatusById(w http.ResponseWriter, r *http.Request)
}

type interestSetting struct {
	interestSetService service.InterestSetting
}

func NewInterestSetting() InterestSetting {
	return &interestSetting{
		interestSetService: service.NewInterestSetting(),
	}
}

func (handler interestSetting) Create(w http.ResponseWriter, r *http.Request) {
	var (
		payload     in.InterestSettingRequest
		userSession = helper.GetAuthSessionModel(r.Context())
	)

	if !userSession.IsAdmin {
		helper.WriteErrorResponse(w, exception.ForbiddenAccess, nil)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequest, nil)
		return
	}

	err = helper.Validate(payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	res, err := handler.interestSetService.Create(r.Context(), payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, res, "Success", http.StatusCreated, nil)
}

func (handler interestSetting) FindAll(w http.ResponseWriter, r *http.Request) {
	userSession := helper.GetAuthSessionModel(r.Context())

	if !userSession.IsAdmin {
		helper.WriteErrorResponse(w, exception.ForbiddenAccess, nil)
		return
	}

	res, err := handler.interestSetService.FindAll(r.Context())
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, res, "Success", http.StatusOK, nil)
}

func (handler interestSetting) FindById(w http.ResponseWriter, r *http.Request) {
	userSession := helper.GetAuthSessionModel(r.Context())

	if !userSession.IsAdmin {
		helper.WriteErrorResponse(w, exception.ForbiddenAccess, nil)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["Id"], 10, 64)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("Interest Settings Id", ""), nil)
		return
	}

	res, err := handler.interestSetService.FindById(r.Context(), id)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, res, "Success", http.StatusOK, nil)
}

func (handler interestSetting) Update(w http.ResponseWriter, r *http.Request) {
	var (
		payload     in.InterestSettingRequest
		userSession = helper.GetAuthSessionModel(r.Context())
	)

	if !userSession.IsAdmin {
		helper.WriteErrorResponse(w, exception.ForbiddenAccess, nil)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequest, nil)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["Id"], 10, 64)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("Interest Settings Id", ""), nil)
		return
	}

	payload.Id = id

	err = helper.Validate(payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	res, err := handler.interestSetService.Update(r.Context(), payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, res, "Success", http.StatusOK, nil)
}

func (handler interestSetting) DeleteById(w http.ResponseWriter, r *http.Request) {
	userSession := helper.GetAuthSessionModel(r.Context())

	if !userSession.IsAdmin {
		helper.WriteErrorResponse(w, exception.ForbiddenAccess, nil)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["Id"], 10, 64)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("Interest Settings Id", ""), nil)
		return
	}

	err = handler.interestSetService.DeleteById(r.Context(), id)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, nil, "Success", http.StatusOK, nil)
}

func (handler interestSetting) ChangeStatusById(w http.ResponseWriter, r *http.Request) {
	var (
		payload     in.InterestSettingRequest
		userSession = helper.GetAuthSessionModel(r.Context())
	)

	if !userSession.IsAdmin {
		helper.WriteErrorResponse(w, exception.ForbiddenAccess, nil)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequest, nil)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["Id"], 10, 64)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("Interest Settings Id", ""), nil)
		return
	}

	payload.Id = id

	res, err := handler.interestSetService.ChangeStatusById(r.Context(), payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, res, "Success", http.StatusOK, nil)
}
