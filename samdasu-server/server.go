package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net"

	"github.com/SherClockHolmes/webpush-go"
	pb "github.com/backendservice/samdasu-alddle"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
	// vapidPrivateKey = "<YOUR VAPID PRIVATE KEY>"
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

	subJSON := `{
		"endpoint": "https://fcm.googleapis.com/fcm/send/fazRZn_0WNg:APA91bFNIkM1CpfnXFgRjeGRlxjNHoIrYUZZrDUK2_wMNDzglHGxCzKpCe_163qU1FVusiiHcrxm2h5ar1fko83L7m-XtJoulvEmCIBc-lQ-ka9nQAB1OL2I7DW4bTFo_5-JeFKt4DUbAxNjRRp22y3rvOp0K4zIJQ",
		"expirationTime": null,
		"keys": {
		  "p256dh": "BM0wnPcg7JKFFXaCM9NYy9ePuV-HEgnfQR0BbwVUy7k4tu_ZF-x-oklQ1lk326JnpB5p10QbISLrijoDiTb4SFo",
		  "auth": "Gtkcbdw7JiOHQDZCbVxjhA"
		}
	  }`

	privateKey, _, vapidErr := webpush.GenerateVAPIDKeys()
	if vapidErr != nil {
		// TODO: Handle failure!
	}

	vapidPrivateKey := privateKey

	// Decode subscription
	subscription := webpush.Subscription{}
	if err := json.NewDecoder(bytes.NewBufferString(subJSON)).Decode(&s); err != nil {
		log.Fatal(err)
	}

	// Send Notification
	_, subErr := webpush.SendNotification([]byte("Test"), &subscription, &webpush.Options{
		Subscriber:      "<EMAIL@EXAMPLE.COM>",
		VAPIDPrivateKey: vapidPrivateKey,
	})
	if subErr != nil {
		log.Fatal(subErr)
	}
}
