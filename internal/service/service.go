package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUserService, wire.Struct(new(Set), "*"))

type Set struct {
	User *UserService
}
