package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func TransactionLoanRoute(privateHandler, publicHandler *mux.Router) {
	private := privateHandler.PathPrefix("/transactions").Subrouter()
	//public := publicHandler.PathPrefix("/user").Subrouter()

	private.HandleFunc("", AppHandler.TransactionLoan.Find).Methods(http.MethodGet)
	private.HandleFunc("", AppHandler.TransactionLoan.Create).Methods(http.MethodPost)
	private.HandleFunc("/{Id:[0-9]+}", AppHandler.TransactionLoan.FindById).Methods(http.MethodGet)
	private.HandleFunc("/{Id:[0-9]+}", AppHandler.TransactionLoan.Update).Methods(http.MethodPut)
	private.HandleFunc("/{Id:[0-9]+}", AppHandler.TransactionLoan.DeleteById).Methods(http.MethodDelete)

}
