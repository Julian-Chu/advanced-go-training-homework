package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"todo/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewTodoRepo)

// Data .
type Data struct {
	todoStore *InMemoryTodoStore
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		logger.Log("msg", "closing the data resources")
	}
	store := NewInMemoryTodoStore()
	return &Data{todoStore: store}, cleanup, nil
}
