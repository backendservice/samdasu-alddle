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

var registeredList []*pb.RegisterRequest

type server struct{}

func (s *server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {

	if find(req.Id) != -1 {
		return &pb.RegisterReply{}, nil
	}

	registeredList = append(registeredList, req)

	log.Println("Length", len(registeredList))
	log.Println("Registered", "data=", req)
	return &pb.RegisterReply{}, nil
}

func find(id string) int {
	for i, v := range registeredList {
		if (*v).Id == id {
			return i
		}
	}
	return -1
}

func (s *server) Unregister(ctx context.Context, req *pb.UnregisterRequest) (*pb.UnregisterReply, error) {
	i := find(req.RegisterId)
	if i != -1 {
		registeredList = append(registeredList[:i], registeredList[i+1:]...)
	}
	return &pb.UnregisterReply{}, nil
}

func (s *server) MatchAndNotify(context.Context, *pb.MatchAndNotifyRequest) (*pb.MatchAndNotifyReply, error) {
	log.Println("Matched ~~, Notified to ~~")
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
