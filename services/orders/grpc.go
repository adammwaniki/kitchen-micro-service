package main

import "google.golang.org/grpc"

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
	grpcServer := grpc.NewServer()

	return grpcServer.Serve() // Serve needs a listener
}