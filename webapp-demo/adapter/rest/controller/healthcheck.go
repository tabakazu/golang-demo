package controller

import (
	"encoding/json"
	"net/http"
)

type HealthCheckController interface {
	Show(w http.ResponseWriter, r *http.Request)
}

type healthCheck struct{}

func NewHealthCheck() *healthCheck {
	return &healthCheck{}
}

func (ctrl healthCheck) Show(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"msg": "ok"})
}
