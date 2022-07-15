package order

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	orderdomain "github.com/lyouthzzz/go-web-layout/internal/domain/order"
	"github.com/lyouthzzz/go-web-layout/internal/mock/domain/order"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUsecase_GetOrder(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	orderClient := order.NewMockIOrderClient(ctl)
	orderUc := NewOrderUsecase(orderClient)

	o := &orderdomain.Order{Id: 1, OrderName: "name"}
	retErr := errors.New("order: not found")
	gomock.InOrder(
		orderClient.EXPECT().GetOrder(gomock.Any(), gomock.Any()).Return(o, nil),
		orderClient.EXPECT().GetOrder(gomock.Any(), gomock.Eq(uint(1))).Return(o, retErr).AnyTimes(),
	)

	_order, err := orderUc.GetOrder(context.Background(), uint(1))
	require.NoError(t, err)
	require.Equal(t, o, _order)

	_order, err = orderUc.GetOrder(context.Background(), uint(1))
	require.Equal(t, retErr, err)
	require.Equal(t, o, _order)

	_order, err = orderUc.GetOrder(context.Background(), uint(1))
	require.Equal(t, retErr, err)
	require.Equal(t, o, _order)
}

func TestOrderMock(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	orderClient := order.NewMockIOrderClient(ctl)

	orderUc := NewOrderUsecase(orderClient)

	orderClient.EXPECT().GetOrder(gomock.Any(), gomock.Eq(uint(111))).Return(nil, nil).AnyTimes()
	o, err := orderUc.GetOrder(context.Background(), 111)
	require.Nil(t, o)
	require.Nil(t, err)

	orderClient.EXPECT().GetOrder(gomock.Eq(context.Background()), gomock.Not(uint(111))).DoAndReturn(func(ctx context.Context, id uint) (*orderdomain.Order, error) {
		if id == 1111 {
			fmt.Println("1111")
		} else {
			fmt.Println(id)
		}
		return nil, nil
	})
	o, err = orderUc.GetOrder(context.Background(), 1111)
	require.Nil(t, o)
	require.Nil(t, err)
}
