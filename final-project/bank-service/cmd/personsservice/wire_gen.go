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
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger, traceTracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	accountClient := data.NewAccountsServiceClient()
	asyncProducer := data.NewKafkaProducer()
	personRepo := data.NewPersonRepo(dataData, logger, accountClient, asyncProducer)
	personUsercase := biz.NewPersonUsercase(personRepo, logger)
	personsService := service.NewPersonsService(personUsercase, logger)
	httpServer := server.NewPersonHTTPServer(confServer, personsService, logger, traceTracerProvider)
	grpcServer := server.NewPersonsGRPCServer(confServer, personsService, logger, traceTracerProvider)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
