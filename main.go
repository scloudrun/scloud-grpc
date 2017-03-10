package main

import (
	config "github.com/astaxie/beego/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"scloud-grpc/controllers"
	pb "scloud-grpc/grpc/helloworld"
)

func main() {
	iniconf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	lis, err := net.Listen("tcp", iniconf.String("httpport"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &controllers.Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
