package main

import (
	"context"
	pb "ficsithub/server/pb/ficsithub"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedFicsitHubServer
}

func (s *server) GetAllPioneers(ctx context.Context, in *pb.Empty) (*pb.PioneersList, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	return &pb.PioneersList{
		Pioneers: getSamplePioneers(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterFicsitHubServer(s, &server{})
	log.Printf("Server started running on port 50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getSamplePioneers() []*pb.Pioneer {
	samplePioneers := []*pb.Pioneer{
		{
			Username: "emingbt",
		},
		{
			Username: "hegobin",
		},
	}
	return samplePioneers
}