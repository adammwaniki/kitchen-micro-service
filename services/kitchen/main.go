package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials())) // This particular implementation allows us to not need credentials since it's educational
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}

func main () {

}