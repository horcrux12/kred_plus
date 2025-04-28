package router

import "kredi-plus.com/be/handler"

type appHandler struct {
	User        handler.User
	Auth        handler.Auth
	Customer    handler.Customer
	InterestSet handler.InterestSetting
}

var AppHandler = appHandler{}.New()

func (route appHandler) New() appHandler {
	return appHandler{
		User:        handler.NewUser(),
		Auth:        handler.NewAuth(),
		Customer:    handler.NewCustomer(),
		InterestSet: handler.NewInterestSetting(),
	}
}
