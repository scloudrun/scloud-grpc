package main

import (
	config "github.com/astaxie/beego/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"scloud-grpc/controllers"
	pb_hw "scloud-grpc/grpc/helloworld"
	pb_ui "scloud-grpc/grpc/userinfo"
)

func main() {
	iniconf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		log.Fatalf("failed to iniconf: %v", err)
	}
	lis, err := net.Listen("tcp", iniconf.String("httpport"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb_hw.RegisterGreeterServer(s, &controllers.Server{})
	pb_ui.RegisterGreeterServer(s, &controllers.Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
