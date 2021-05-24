package main

import (
	"github.com/tabakazu/webapi-app/adapter/controller/rest"
	"github.com/tabakazu/webapi-app/adapter/gateway/datastore"
)

func main() {
	db := datastore.NewConnection()
	srv := rest.NewServer()

	// HealthCheck
	hcc := rest.NewHealthCheckController()
	rest.SetupHealthCheckRoutes(srv.Router(), hcc)

	// UserAccount
	uar := datastore.NewUserAccountRepository(db)
	uac := rest.NewUserAccountController(uar)
	rest.SetupUserAccountRoutes(srv.Router(), uac)

	srv.ListenAndServe()
}
