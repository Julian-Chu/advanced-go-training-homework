// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"persons-service/internal/biz"
	"persons-service/internal/conf"
	"persons-service/internal/data"
	"persons-service/internal/server"
	"persons-service/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSetAccount,
		data.ProviderSetAccount,
		biz.ProviderSetAccount,
		service.ProviderSetAccount,
		newApp))
}
