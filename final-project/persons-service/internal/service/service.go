package service

import "github.com/google/wire"

// ProviderSetPerson is service providers.
var ProviderSetPerson = wire.NewSet(NewPersonsService)

// ProviderSetAccount is service providers.
var ProviderSetAccount = wire.NewSet(NewAccountsService)
