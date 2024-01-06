package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	apiV1 "github.com/lyouthzzz/go-web-layout/api/v1"
	"github.com/lyouthzzz/go-web-layout/internal/conf"
	"github.com/lyouthzzz/go-web-layout/internal/service"
)

func NewHTTPServer(config *conf.Server, serviceSet *service.Set) (*http.Server, error) {
	httpServer := http.NewServer(http.Network(config.Http.Network), http.Address(config.Http.Addr), http.Timeout(config.Http.Timeout.AsDuration()))

	apiV1.RegisterUserServiceHTTPServer(httpServer, serviceSet.User)

	return httpServer, nil
}
