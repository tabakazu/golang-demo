//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tabakazu/webapi-app/infra/web"
)

func InitializeWebApp() *web.Server {
	wire.Build(
		web.NewServer,
	)
	return &web.Server{}
}
