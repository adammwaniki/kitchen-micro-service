package types

import (
	"context"

	"github.com/adammwaniki/kitchen-micro-service/services/common/genproto/orders"
)

type OrderService interface {
	// CreateOrder service will be injected into the handler
	// The method has been generated in gRPC and so we need to basically replicate it in golang so that the input and output matches whatever we want from the business logic
	// The second parameter will be the order payload to create the order. This comes from the public struct methods in the generated code
	CreateOrder(context.Context, *orders.Order) error // passing in context to improve compatibility with programs around the world. If any method implements or receives context we can pass it around
	GetOrders(context.Context) []*orders.Order // We could add the customerID as a parameter but for simplicity of the tutorial we shan't
}