package main

func main() {
	httpServer := NewHttpServer(":8000") // Needs to be different from the grpc server else conflict
	go httpServer.Run() // the http and grpc servers need to run concurrently. Without concurrency only this one will run

	grpcServer := NewGRPCServer(":9000")
	grpcServer.Run()
}