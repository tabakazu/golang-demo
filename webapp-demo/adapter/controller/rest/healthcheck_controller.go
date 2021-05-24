package rest

import (
	"encoding/json"
	"net/http"
)

type HealthCheckController interface {
	CheckHandler(http.ResponseWriter, *http.Request)
}

type healthCheckController struct{}

func NewHealthCheckController() HealthCheckController {
	return &healthCheckController{}
}

func (ctrl healthCheckController) CheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"msg": "ok!"})
}
