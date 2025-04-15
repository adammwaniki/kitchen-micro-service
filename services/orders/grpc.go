package main

import (
	"log"
	"net"

	handler "github.com/adammwaniki/kitchen-micro-service/services/orders/handler/orders"
	"github.com/adammwaniki/kitchen-micro-service/services/orders/service"
	"google.golang.org/grpc"
)

// Server struct
type gRPCServer struct {
	addr string
}

// Generator function to create a new instance of the struct
func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

// Run method to initialise it
func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)// Create a tcp connection that we can send to the Serve method and listen on
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register our grpc services
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService) // We don't need to assign it to an orderHandler variable

	log.Printf("Starting gRPC server on %v", s.addr)

	return grpcServer.Serve(lis) // Serve needs a listener
}