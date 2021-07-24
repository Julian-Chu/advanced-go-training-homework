package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	accounts "persons-service/api/accountsservice"
	"persons-service/internal/conf"
)

// ProviderSetPerson is data providers.
var ProviderSetPerson = wire.NewSet(NewData, NewAccountsServiceClient, NewPersonRepo)

var ProviderSetAccount = wire.NewSet(NewData, NewAccountRepo)

// Data .
type Data struct {
	// TODO warpped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func NewAccountsServiceClient() accounts.AccountClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9001"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return accounts.NewAccountClient(conn)
}
