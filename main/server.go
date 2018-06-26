package main

import (
	"context"
	"log"
	"net"

	pb "github.com/backendservice/samdasu-alddle"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Register(context.Context, *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}

func (s *server) Unregister(context.Context, *pb.UnregisterRequest) (*pb.UnregisterReply, error) {
	return &pb.UnregisterReply{}, nil
}

func (s *server) MatchAndNotify(context.Context, *pb.MatchAndNotifyRequest) (*pb.MatchAndNotifyReply, error) {
	return &pb.MatchAndNotifyReply{}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAlddleServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
