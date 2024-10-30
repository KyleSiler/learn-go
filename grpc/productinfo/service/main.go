package main

import (
	"context"
	"log"
	"net"
	pb "productinfo/service/ecommerce"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	productMap map[string]*pb.Product
	pb.UnimplementedProductInfoServer
}

func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	log.Printf("Inserted product %s", in.Name)

	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	product, exists := s.productMap[in.Value]
	if exists && product != nil {
		log.Printf("Product %v : %v - Retrieved.", product.Id, product.Name)
		return product, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}

func main() {
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen to %v", err)
	}
	service := grpc.NewServer()
	pb.RegisterProductInfoServer(service, &server{})
	log.Printf("Listening at %v", list.Addr())
	if err = service.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
