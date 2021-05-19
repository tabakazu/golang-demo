package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer(r *mux.Router) *http.Server {
	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}
	return srv
}
