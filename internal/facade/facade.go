package facade

import (
	"github.com/google/wire"
	"github.com/lyouthzzz/go-web-layout/internal/facade/order"
)

var ProviderSet = wire.NewSet(order.NewClient)
