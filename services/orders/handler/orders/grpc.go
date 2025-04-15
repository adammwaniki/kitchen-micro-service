package handler

import (
	"context"

	"github.com/adammwaniki/kitchen-micro-service/services/common/genproto/orders"
	"github.com/adammwaniki/kitchen-micro-service/services/orders/types"
)

type OrdersGrpcHandler struct {
	// Service injection -- will require an interface
	ordersService 		types.OrderService
	// Unimplemented UnimplementedOrderServiceServer  --there is a corresponding OrderServiceServer interface generated in the grpc code
	// Composing the interface as follows so that it is obtained from the generated code
	orders.UnimplementedOrderServiceServer // Read on struct and interface composability --different from inheritance
}

func NewGrpcOrdersService() {
	gRPCHandler := &OrdersGrpcHandler{}

	// Register the OrderServiceServer --the UnimplementedOrderService will allow us to register our order service in this function
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	
}