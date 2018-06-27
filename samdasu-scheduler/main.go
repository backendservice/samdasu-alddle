package main

import (
	"log"
	"time"

	pb "github.com/backendservice/samdasu-alddle"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	for {
		// Set up a connection to the server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Println("did not connect: ", err)
		}
		defer conn.Close()
		c := pb.NewAlddleClient(conn)
		_, err = call(context.Background(), c, time.Minute)
		if err != nil {
			log.Println("could not call: ", err)
		} else {
			log.Println("called")
		}
		time.Sleep(1 * time.Second)
	}
}

func call(ctx context.Context, c pb.AlddleClient, timeout time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	_, err := c.MatchAndNotify(ctx, &pb.MatchAndNotifyRequest{})
	if err != nil {
		return "", err
	}
	return "", nil
}
