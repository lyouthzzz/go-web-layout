//+build wireinject
//
//The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lyouthzzz/go-web-layout/internal/biz"
	"github.com/lyouthzzz/go-web-layout/internal/conf"
	"github.com/lyouthzzz/go-web-layout/internal/data"
	"github.com/lyouthzzz/go-web-layout/internal/facade"
	"github.com/lyouthzzz/go-web-layout/internal/server"
	"github.com/lyouthzzz/go-web-layout/internal/service"
)

func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(data.ProviderSet, facade.ProviderSet, biz.ProviderSet, service.ProviderSet, server.ProviderSet, newApp))
}
