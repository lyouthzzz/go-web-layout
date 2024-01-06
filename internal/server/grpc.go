package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	apiV1 "github.com/lyouthzzz/go-web-layout/api/v1"
	"github.com/lyouthzzz/go-web-layout/internal/conf"
	"github.com/lyouthzzz/go-web-layout/internal/service"
)

func NewGRPCServer(config *conf.Server, serviceSet *service.Set) (*grpc.Server, error) {
	grpcServer := grpc.NewServer(grpc.Network(config.Grpc.Network), grpc.Address(config.Grpc.Addr), grpc.Timeout(config.Grpc.Timeout.AsDuration()))

	apiV1.RegisterUserServiceServer(grpcServer, serviceSet.User)

	return grpcServer, nil
}
