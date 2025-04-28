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

type User interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetailById(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	DeleteById(w http.ResponseWriter, r *http.Request)
}

type user struct {
	userService service.User
}

func NewUser() User {
	return &user{
		userService: service.NewUser(),
	}
}

func (handler user) Create(w http.ResponseWriter, r *http.Request) {
	var (
		payload     in.UserRequest
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

	res, err := handler.userService.Create(r.Context(), payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, res, "Success", http.StatusCreated, nil)
}

func (handler user) GetList(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	payload := in.UserRequest{
		AbstractRequest: in.AbstractRequest{
			Page:    page,
			Limit:   limit,
			Search:  r.URL.Query().Get("search"),
			SortStr: r.URL.Query().Get("sort"),
		},
	}

	userSession := helper.GetAuthSessionModel(r.Context())
	if !userSession.IsAdmin {
		helper.WriteErrorResponse(w, exception.ForbiddenAccess, nil)
		return
	}

	res, totalData, err := handler.userService.GetList(r.Context(), payload)
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

func (handler user) GetDetailById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseInt(params["Id"], 10, 64)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("user Id", ""), nil)
		return
	}

	res, err := handler.userService.GetDetailById(r.Context(), userId)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, res, "Success", http.StatusOK, nil)
}

func (handler user) Update(w http.ResponseWriter, r *http.Request) {
	var (
		payload in.UserRequest
		params  = mux.Vars(r)
	)
	userId, err := strconv.ParseInt(params["Id"], 10, 64)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("user Id", ""), nil)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequest, nil)
		return
	}

	payload.ID = userId

	err = helper.Validate(payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	res, err := handler.userService.Update(r.Context(), payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, res, "Success", http.StatusOK, nil)
}

func (handler user) DeleteById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseInt(params["Id"], 10, 64)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequestWithMessage("user Id", ""), nil)
		return
	}

	err = handler.userService.DeleteById(r.Context(), userId)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, nil, "Success", http.StatusOK, nil)
}
