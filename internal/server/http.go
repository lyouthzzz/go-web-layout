package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	orderV1 "github.com/lyouthzzz/go-web-layout/api/order/v1"
	userV1 "github.com/lyouthzzz/go-web-layout/api/user/v1"
	"github.com/lyouthzzz/go-web-layout/internal/conf"
	"github.com/lyouthzzz/go-web-layout/internal/service"
)

func NewHTTPServer(config *conf.Server, user *service.UserService, order *service.OrderService) (*http.Server, error) {
	httpServer := http.NewServer(http.Network(config.Http.Network), http.Address(config.Http.Addr), http.Timeout(config.Http.Timeout.AsDuration()))
	userV1.RegisterUserServiceHTTPServer(httpServer, user)
	orderV1.RegisterOrderServiceHTTPServer(httpServer, order)
	return httpServer, nil
}
