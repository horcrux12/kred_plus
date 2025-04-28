package handler

import (
	"encoding/json"
	"kredi-plus.com/be/dto/in"
	"kredi-plus.com/be/lib/exception"
	"kredi-plus.com/be/lib/helper"
	"kredi-plus.com/be/service"
	"net/http"
)

type Auth interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type auth struct {
	service service.Auth
}

func NewAuth() Auth {
	return &auth{
		service: service.NewAuth(),
	}
}

func (handler auth) Login(w http.ResponseWriter, r *http.Request) {
	var payload in.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		helper.WriteErrorResponse(w, exception.InvalidRequest, nil)
		return
	}

	token, err := handler.service.Login(r.Context(), payload)
	if err != nil {
		helper.WriteErrorResponse(w, err, nil)
		return
	}

	helper.WriteSuccessResponse(w, token, "Success", http.StatusOK, nil)
}

func (handler auth) Logout(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
