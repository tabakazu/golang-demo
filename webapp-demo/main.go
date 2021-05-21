package main

import (
	"github.com/tabakazu/webapi-app/adapter/controller/rest"
)

func main() {
	hcc := rest.NewHealthCheckController()
	srv := rest.NewServer()
	rest.SetupHealthCheckRoutes(srv.Router(), hcc)
	srv.ListenAndServe()
}
