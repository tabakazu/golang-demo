package rest

import (
	"net/http"
)

type Routeing interface {
	GET(string, func(http.ResponseWriter, *http.Request))
	POST(string, func(http.ResponseWriter, *http.Request))
}

func SetupHealthCheckRoutes(r Routeing, ctrl HealthCheckController) {
	r.GET("/health_check", ctrl.CheckHandler)
}
