package main

import (
	"context"
	pbPing "github.com/Gruzchick/learns-grpc/ping"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type server struct {
	pbPing.UnimplementedServiceServer
}

func (s *server) Ping(ctx context.Context, in *emptypb.Empty) (*pbPing.PingResponse, error) {
	return &pbPing.PingResponse{Message: "Pinged"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pbPing.RegisterServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
