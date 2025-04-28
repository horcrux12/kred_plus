package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func UserRoute(privateHandler, publicHandler *mux.Router) {
	private := privateHandler.PathPrefix("/user").Subrouter()
	//public := publicHandler.PathPrefix("/user").Subrouter()

	private.HandleFunc("", AppHandler.User.GetList).Methods(http.MethodGet)
	private.HandleFunc("", AppHandler.User.Create).Methods(http.MethodPost)
	private.HandleFunc("/{Id:[0-9]+}", AppHandler.User.GetDetailById).Methods(http.MethodGet)
	private.HandleFunc("/{Id:[0-9]+}", AppHandler.User.Update).Methods(http.MethodPut)
	private.HandleFunc("/{Id:[0-9]+}", AppHandler.User.DeleteById).Methods(http.MethodDelete)

}
