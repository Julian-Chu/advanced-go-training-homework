// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"bankservice/internal/biz"
	"bankservice/internal/conf"
	"bankservice/internal/data"
	"bankservice/internal/server"
	"bankservice/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSetPerson, data.ProviderSetPerson, biz.ProviderSetPerson, service.ProviderSetPerson, newApp))
}
