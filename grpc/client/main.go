package main

import (
	"context"
	"log"
	pb "productinfo/client/ecommerce"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to client %v", err)
	}

	defer conn.Close()
	c := pb.NewProductInfoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.AddProduct(ctx, &pb.Product{Name: "Test", Price: 10.50, Description: "This is a test", Id: "1234"})
	if err != nil {
		log.Fatalf("Failed to save product %v", err)
	}

	log.Printf("Product successfully saved")

	p, err := c.GetProduct(ctx, &pb.ProductID{Value: "1234"})

	log.Print(p.String())
}
