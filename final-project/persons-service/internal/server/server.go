package server

import (
	"github.com/google/wire"
)

// ProviderSetPerson is server providers.
var ProviderSetPerson = wire.NewSet(NewPersonHTTPServer, NewPersonsGRPCServer)
var ProviderSetAccount = wire.NewSet(NewAccountsGRPCServer)
