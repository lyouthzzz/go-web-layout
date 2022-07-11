package order

import (
	"context"
	"time"
)

//go:generate mockgen -source order.go -destination ../../mock/domain/order/order_mock.go -package=order

type IOrderClient interface {
	GetOrder(context.Context, uint) (*Order, error)
}

type Order struct {
	// 订单ID
	Id uint
	// 订单名称
	OrderName string
	// 订单时间
	OrderCratedTime time.Time
}
