package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "scloud-grpc/grpc/userinfo"
)

const (
	address           = "localhost:50051"
	defaultName int64 = 123456
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	r, err := c.GetUserInfo(context.Background(), &pb.UserInfoRequest{Uid: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
