package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	orderV1 "github.com/lyouthzzz/go-web-layout/api/order/v1"
	userV1 "github.com/lyouthzzz/go-web-layout/api/user/v1"
	"github.com/lyouthzzz/go-web-layout/internal/conf"
	"github.com/lyouthzzz/go-web-layout/internal/service"
)

func NewGRPCServer(config *conf.Server, user *service.UserService, order *service.OrderService) (*grpc.Server, error) {
	grpcServer := grpc.NewServer(grpc.Network(config.Grpc.Network), grpc.Address(config.Grpc.Addr), grpc.Timeout(config.Grpc.Timeout.AsDuration()))
	userV1.RegisterUserServiceServer(grpcServer, user)
	orderV1.RegisterOrderServiceServer(grpcServer, order)
	return grpcServer, nil
}
