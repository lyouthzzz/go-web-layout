package order

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	orderdomain "github.com/lyouthzzz/go-web-layout/internal/domain/order"
	"github.com/lyouthzzz/go-web-layout/internal/mock/domain/order"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUsecase_GetOrder(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	o := &orderdomain.Order{Id: 1, OrderName: "name"}

	orderClient := order.NewMockIOrderClient(ctl)

	// input any matcher
	orderClient.EXPECT().GetOrder(gomock.Any(), gomock.Any()).Return(o, nil)
	orderUc := NewOrderUsecase(orderClient)

	_order, err := orderUc.GetOrder(context.Background(), uint(1))
	require.NoError(t, err)
	require.Equal(t, o, _order)

	// input eq matcher output error
	var retErr = errors.New("order: not found")
	orderClient.EXPECT().GetOrder(gomock.Any(), gomock.Eq(uint(1))).Return(o, retErr)
	orderUc = NewOrderUsecase(orderClient)

	_order, err = orderUc.GetOrder(context.Background(), uint(2))
	require.Equal(t, retErr, err)
	require.Equal(t, o, _order)
}
