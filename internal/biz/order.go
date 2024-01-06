package biz

import "time"

type Order struct {
	// 订单ID
	Id uint
	// 订单名称
	OrderName string
	// 订单时间
	OrderCratedTime time.Time
}
