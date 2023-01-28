package main

import (
	"context"

	pb "productinfo/service/ecommerce"

	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Server struct {
	productMap map[string]*pb.Product

	pb.UnimplementedProductInfoServer
}

func MakeServer() *Server {
	return &Server{
		productMap: make(map[string]*pb.Product),
	}
}

func (s *Server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}

	in.Id = out.String()
	s.productMap[in.Id] = in

	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *Server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}

func (s *Server) GreetUser(ctx context.Context, _ *empty.Empty) (*wrapperspb.StringValue, error) {
	return &wrapperspb.StringValue{Value: "Hello, world!"}, status.New(codes.OK, "").Err()
}
