package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"persons-service/internal/conf"
)

// ProviderSetPerson is data providers.
var ProviderSetPerson = wire.NewSet(NewData, NewPersonRepo)

var ProviderSetAccount = wire.NewSet(NewData, newAccountRepo)

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
