package grpcserver

import (
	"carsvc/internal/grpcapi"

	"google.golang.org/grpc"
)

type grpcserver struct {
	grpcapi.UnimplementedCarDataServiceServer
}

func NewGRPCServer() *grpc.Server {
	srv := grpc.NewServer()
	svc := grpcserver{}

	grpcapi.RegisterCarDataServiceServer(srv, &svc)

	return srv
}
