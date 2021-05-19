//+build wireinject

package main

import (
	"net/http"

	"github.com/google/wire"
	"github.com/tabakazu/webapi-app/adapter/rest"
)

func InitializeWebApp() *http.Server {
	wire.Build(
		rest.NewRouter,
		rest.NewServer,
	)
	return &http.Server{}
}
