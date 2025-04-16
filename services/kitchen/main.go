package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials())) // This particular implementation allows us to not need credentials since it's educational
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	return conn
}

func main () {
	httpServer := NewHttpServer(":8060") // Needs to be different from the other servers because it is an entirely different service. NB ports below 1024 need root access on unix systems for security reasons
	if err := httpServer.Run(); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}