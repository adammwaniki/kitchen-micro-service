package main

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

}