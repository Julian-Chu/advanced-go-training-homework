// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"persons-service/internal/biz"
	"persons-service/internal/conf"
	"persons-service/internal/data"
	"persons-service/internal/server"
	"persons-service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	personRepo := data.NewPersonRepo(dataData, logger)
	personUsercase := biz.NewPersonUsercase(personRepo, logger)
	personsService := service.NewPersonsService(personUsercase, logger)
	httpServer := server.NewHTTPServer(confServer, personsService, logger)
	grpcServer := server.NewGRPCServer(confServer, personsService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
