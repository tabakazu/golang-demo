package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tabakazu/webapi-app/adapter/rest/controller"
)

func NewServer() *http.Server {
	r := mux.NewRouter()

	healthCheckCtrl := controller.NewHealthCheck()
	r.HandleFunc("/healthcheck", healthCheckCtrl.Show)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}
	return srv
}
