package biz

import "github.com/google/wire"

// ProviderSetPerson is Person biz providers.
var ProviderSetPerson = wire.NewSet(NewPersonUsercase)

// ProviderSetAccount is Account biz providers.
var ProviderSetAccount = wire.NewSet(NewAccountUsercase)
