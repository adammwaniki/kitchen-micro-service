/*
- Business logic will be held here
- It shall be separate from the transport layer, database layer, gRPC layer etc.
*/

package service

import (
	"context"

	"github.com/adammwaniki/kitchen-micro-service/services/common/genproto/orders"
)

var ordersDb = make([]*orders.Order, 0) // For learning purposes. Creating a simple slice of orders in memory so that even if the server dies we won't necessarily lose the orders

type OrderService struct {
	//store dependency injection
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDb = append(ordersDb, order)

	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*orders.Order {
	return ordersDb
}