package main

import (
	"log"
	"os"
	"time"

	"github.com/satori/go.uuid"

	pb "github.com/backendservice/samdasu-alddle"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	//	address     = "ec2-18-191-204-27.us-east-2.compute.amazonaws.com:50051"
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAlddleClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	msg, err := register(context.Background(), c, time.Minute, name)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("Started : {}", msg)
}

func register(ctx context.Context, c pb.AlddleClient, timeout time.Duration, name string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	id, _ := uuid.NewV4()
	_, err := c.Register(ctx, &pb.RegisterRequest{
		Id:          id.String(),
		Email:       "hodol.kang@gmail.com",
		Departure:   "Seoul",
		Destination: "Osaka",
		Expense:     1000000,
		Duration:    4,
		FromDate:    "20180701",
		ToDate:      "20180831",
	})
	if err != nil {
		return "aa", err
	}
	return "registered", nil
}
