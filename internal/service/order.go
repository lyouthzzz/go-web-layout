package service

import (
	"context"
	orderV1 "github.com/lyouthzzz/go-web-layout/api/order/v1"
	orderbiz "github.com/lyouthzzz/go-web-layout/internal/biz/order"
)

type OrderService struct {
	orderV1.UnimplementedOrderServiceServer
	order *orderbiz.Usecase
}

func NewOrderService(order *orderbiz.Usecase) *OrderService {
	return &OrderService{order: order}
}

func (order *OrderService) GetUser(ctx context.Context, req *orderV1.GetOrderRequest) (*orderV1.GetOrderReply, error) {
	return &orderV1.GetOrderReply{}, nil
}
