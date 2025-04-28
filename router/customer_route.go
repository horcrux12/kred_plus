package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func CustomerRoute(privateHandler, publicHandler *mux.Router) {
	private := privateHandler.PathPrefix("/customers").Subrouter()
	//public := publicHandler.PathPrefix("/user").Subrouter()

	private.HandleFunc("", AppHandler.Customer.GetListCustomer).Methods(http.MethodGet)
	private.HandleFunc("", AppHandler.Customer.CreateCustomer).Methods(http.MethodPost)
	private.HandleFunc("/{Id:[0-9]+}", AppHandler.Customer.GetCustomerById).Methods(http.MethodGet)
	private.HandleFunc("/{Id:[0-9]+}", AppHandler.Customer.UpdateCustomer).Methods(http.MethodPut)
	private.HandleFunc("/{Id:[0-9]+}", AppHandler.Customer.DeleteCustomerById).Methods(http.MethodDelete)

}
