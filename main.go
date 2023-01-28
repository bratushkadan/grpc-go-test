package main

import (
	"log"
	"net"
	pb "productinfo/service/ecommerce"

	"google.golang.org/grpc"
)

const port string = ":50001"

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, MakeServer())

	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
