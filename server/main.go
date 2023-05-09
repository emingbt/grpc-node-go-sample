package main

import (
	"context"
	"net"
	"fmt"
	"bending/api"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	api.RegisterBendingServiceServer(grpcServer, &BendingService{})
	fmt.Println("Starting server...")
	grpcServer.Serve(lis)
}

type BendingService struct {
	api.UnimplementedBendingServiceServer
}

func (s *BendingService) CalculateBendingType(ctx context.Context,
				req *api.BendingRequest) (*api.BendingResponse, error) {
	BendingTypes := []string{"Air", "Water", "Earth", "Fire"}
	return &api.BendingResponse{
		BendingType: BendingTypes[req.BirthYear % 4],
	}, nil
}
