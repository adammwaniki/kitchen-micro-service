package handler

import (
	"context"

	"github.com/adammwaniki/kitchen-micro-service/services/common/genproto/orders"
	"github.com/adammwaniki/kitchen-micro-service/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	// Service injection -- will require an interface
	ordersService 	types.OrderService
	// Unimplemented UnimplementedOrderServiceServer  --there is a corresponding OrderServiceServer interface generated in the grpc code
	// Composing the interface as follows so that it is obtained from the generated code
	orders.UnimplementedOrderServiceServer // Read on struct and interface composability --different from inheritance
}

func NewGrpcOrdersService(grpc *grpc.Server, ordersService types.OrderService, ) { // needs to receive the service from OrdersGrpcHandler above and the grpc server from services/orders/grpc.go
	gRPCHandler := &OrdersGrpcHandler{
		ordersService: ordersService,
	}

	// Register the OrderServiceServer --the UnimplementedOrderService will allow us to register our order service in this function
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

// The new GetOrders grpc method implementation
func (h *OrdersGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrderResponse, error) {
	o := h.ordersService.GetOrders(ctx) // not calling it orders to prevent naming conflict with the orders package
	res := &orders.GetOrderResponse{
		Orders: o,
	}

	return res, nil
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	// We can start implementing and consuming the order service
	// Dummy data to experiment with
	order := &orders.Order{
		OrderID: 42,
		CustomerID: 2,
		ProductID: 1,
		Quantity: 10,
	}

	err := h.ordersService.CreateOrder(ctx, order) // CreateOrder will need to be implemented on the service
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}