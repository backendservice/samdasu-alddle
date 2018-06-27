package main

import (
	"log"
	"os"
	"strconv"
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

	if len(os.Args) < 3 {
		log.Println("usage: go run main.go [EMAIL] [EXPENSE]")
		return
	}

	email := os.Args[1]
	expense, _ := strconv.Atoi(os.Args[2])

	_, err = register(context.Background(), c, time.Minute, email, int32(expense))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	} else {
		log.Println("Enjoy holiday with samdasu")
	}
}

func register(ctx context.Context, c pb.AlddleClient, timeout time.Duration, email string, expense int32) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	id := uuid.NewV4()
	_, err := c.Register(ctx, &pb.RegisterRequest{
		Id:          id.String(),
		Email:       email,
		Departure:   "Seoul",
		Destination: "Osaka",
		Expense:     expense,
		Duration:    4,
		FromDate:    "20180701",
		ToDate:      "20180831",
	})
	if err != nil {
		return "aa", err
	}
	return "registered", nil
}
