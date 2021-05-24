package server

import (
	"github.com/douyu/jupiter/pkg/server"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
	v1 "github.com/lyouthzzz/go-web-layout/api/user/v1"
	"github.com/lyouthzzz/go-web-layout/internal/service"
)

func NewGRPCServer(service *service.UserService) (server.Server, error) {
	server := xgrpc.StdConfig("grpcserver").Build()
	v1.RegisterUserServer(server, service)
	return server, nil
}
