package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/adammwaniki/kitchen-micro-service/services/common/genproto/orders"
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
		// Typically the functions here would be kept in a service directory under the kitchen directory similar to the orders for separation of concerns
		// However, for educational purposes we will write it all here

		// pass the connection to the client of the order service. This will be gotten from the generated code
		c := orders.NewOrderServiceClient(conn)

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{ // This is similar to a POST request to the api but the difference is that grpc allows us to just call it like a function (a remote procedure call)
			CustomerID: 24,
			ProductID: 3123,
			Quantity: 2,
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		// Getting orders from the client
		res, err := c.GetOrders(ctx, &orders.GetOrdersRequest{
			CustomerID: 42, // our sample customer
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		// Create a new template and parse it
		t := template.Must(template.New("orders").Parse(ordersTemplate))

		if err := t.Execute(w, res.GetOrders()); err != nil {
			log.Fatalf("template error: %v", err)
		}
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