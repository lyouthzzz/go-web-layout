package order

import (
	"context"
	orderdomain "github.com/lyouthzzz/go-web-layout/internal/domain/order"
)

type Usecase struct {
	order orderdomain.IOrderClient
}

func NewOrderUsecase(order orderdomain.IOrderClient) *Usecase {
	return &Usecase{order: order}
}

func (uc *Usecase) GetOrder(ctx context.Context, id uint) (*orderdomain.Order, error) {
	return uc.order.GetOrder(ctx, id)
}
