package main

import (
	"log"
	"net/http"
)

// The Kitchen will connect to the orders via grpc while the end clients will interact with the browser via http (the orders are received by the browser in http)
type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGRPCClient(":9000") // This will establish a link to the grpc server that we are running on the orders
	defer conn.Close() // We defer this connection to allow us to release the resources

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// We will create a grpc connection to our orders
	})

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

// Template for vewing the orders. The user will view them as a table. We can connect to any frontend framework e.g., ReactJS etc.
// Template is for simplicity
var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`