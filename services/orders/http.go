package main

import (
	"log"
	"net/http"

	"github.com/adammwaniki/kitchen-micro-service/services/orders/service"
)

// Implementations will be similar to grpc
type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService() // Implementing the orderService this way allows separation of concerns from business logic in the services

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}