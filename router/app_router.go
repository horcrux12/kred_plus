package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"kredi-plus.com/be/config"
	"kredi-plus.com/be/lib/helper"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func AppRouter() {
	api := mux.NewRouter()

	restPrefix := fmt.Sprintf("/%s/%s",
		config.Attr.App.Prefix,
		config.Attr.App.Version)

	publicHandler := api.PathPrefix(restPrefix).Subrouter()
	privateHandler := api.PathPrefix(restPrefix).Subrouter()

	// sub router
	UserRoute(privateHandler, publicHandler)
	CustomerRoute(privateHandler, publicHandler)
	InterestSetting(privateHandler, publicHandler)
	TransactionLoanRoute(privateHandler, publicHandler)

	// health
	publicHandler.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		helper.WriteSuccessResponse(writer, nil, "Everything is good", 200, nil)
	}).Methods(http.MethodGet)

	// Auth
	publicHandler.HandleFunc("/auth/login", AppHandler.Auth.Login).Methods(http.MethodPost)

	// Private Image link
	imgLink := api.PathPrefix("").Subrouter()
	imgLink.HandleFunc("/img/{funcName}/{filename}", ProtectedFileHandler).Methods(http.MethodGet)

	publicHandler.Use(PublicMiddleware)
	privateHandler.Use(PrivateMiddleware)
	imgLink.Use(PrivateMiddleware)

	// Serve router
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", config.Attr.App.Port),
		Handler:      api,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	fmt.Println("Application has started in port : ", config.Attr.App.Port)
	fmt.Println(fmt.Sprintf("Main route : 0.0.0.0:%s/%s/%s",
		config.Attr.App.Port,
		config.Attr.App.Prefix,
		config.Attr.App.Version))
	srv.ListenAndServe()
}

// --- Protected File Serve ---
func ProtectedFileHandler(w http.ResponseWriter, r *http.Request) {
	const uploadPath = "./uploads/"
	vars := mux.Vars(r)
	filename := vars["filename"]
	funcName := vars["funcName"]

	filePath := filepath.Join(uploadPath+"/"+funcName+"/", filename)

	// cek file ada atau tidak
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, filePath)
}
