package main

import (
	"log"
	"net/http"

	handler "github.com/adammwaniki/kitchen-micro-service/services/orders/handler/orders"
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
	orderHandler := handler.NewHttpOrdersHandler(orderService) // Register the http handler
	orderHandler.RegisterRouter(router) // Register the router by using the handler

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}