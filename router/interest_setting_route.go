package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InterestSetting(privateHandler, publicHandler *mux.Router) {
	private := privateHandler.PathPrefix("/interest-settings").Subrouter()
	//public := publicHandler.PathPrefix("/user").Subrouter()

	private.HandleFunc("", AppHandler.InterestSet.FindAll).Methods(http.MethodGet)
	private.HandleFunc("", AppHandler.InterestSet.Create).Methods(http.MethodPost)
	private.HandleFunc("/{Id:[0-9]+}", AppHandler.InterestSet.FindById).Methods(http.MethodGet)
	private.HandleFunc("/{Id:[0-9]+}", AppHandler.InterestSet.Update).Methods(http.MethodPut)
	private.HandleFunc("/{Id:[0-9]+}", AppHandler.InterestSet.DeleteById).Methods(http.MethodDelete)
	private.HandleFunc("/change-status/{Id:[0-9]+}", AppHandler.InterestSet.ChangeStatusById).Methods(http.MethodPut)
}
