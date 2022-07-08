package order

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/http"
	orderdomain "github.com/lyouthzzz/go-web-layout/internal/domain/order"
	"time"
)

var _ orderdomain.IOrderClient = (*Client)(nil)

func NewClient() orderdomain.IOrderClient {
	return &Client{c: nil}
}

type Client struct {
	c *http.Client
}

func (c *Client) GetOrder(ctx context.Context, oid uint) (*orderdomain.Order, error) {
	return &orderdomain.Order{Id: 1, OrderName: "order_name", OrderCratedTime: time.Now()}, nil
}
