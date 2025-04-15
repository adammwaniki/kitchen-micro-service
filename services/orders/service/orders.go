/*
- Business logic will be held here
- It shall be separate from the transport layer, database layer, gRPC layer etc.
*/

package service

type OrderService struct {
	//store dependency injection
}

func NewOrderService() *OrderService {
	return &OrderService{}
}