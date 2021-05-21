package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *router
}

type router struct {
	*mux.Router
}

func (r *router) GET(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.HandleFunc(path, handler).Methods("GET")
}

func (r *router) POST(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.HandleFunc(path, handler).Methods("POST")
}

func NewServer() *server {
	r := mux.NewRouter()
	router := &router{r}
	return &server{
		router: router,
	}
}

func (s *server) Router() *router {
	return s.router
}

func (s *server) ListenAndServe() error {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: s.router,
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
