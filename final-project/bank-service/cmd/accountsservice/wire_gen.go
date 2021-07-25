// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"bankservice/internal/biz"
	"bankservice/internal/conf"
	"bankservice/internal/data"
	"bankservice/internal/server"
	"bankservice/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	accountRepo := data.NewAccountRepo(dataData, logger)
	accountUsercase := biz.NewAccountUsercase(accountRepo, logger)
	accountsService := service.NewAccountsService(accountUsercase, logger)
	grpcServer := server.NewAccountsGRPCServer(confServer, accountsService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}