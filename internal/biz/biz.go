package biz

import (
	"github.com/google/wire"
	"github.com/lyouthzzz/go-web-layout/internal/biz/order"
	"github.com/lyouthzzz/go-web-layout/internal/biz/user"
)

var ProviderSet = wire.NewSet(user.NewUserUsecase, order.NewOrderUsecase)
