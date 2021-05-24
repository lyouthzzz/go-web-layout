package server

import (
	"github.com/douyu/jupiter/pkg/server"
	"github.com/douyu/jupiter/pkg/server/xgin"
	v1 "github.com/lyouthzzz/go-web-layout/api/user/v1"
	"github.com/lyouthzzz/go-web-layout/internal/service"
)

func NewHTTPServer(service *service.UserService) (server.Server, error) {
	server := xgin.StdConfig("httpserver").Build()
	userServer := v1.NewUserHTTPServer(server, service)
	userServer.BuildRoute()
	return userServer, nil
}
