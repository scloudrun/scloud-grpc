package controllers

import (
	"golang.org/x/net/context"
	pb "scloud-grpc/grpc/userinfo"
)

func (s *Server) GetUserInfo(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoReply, error) {
	userInfo := GetUserById(in.Uid)
	return &pb.UserInfoReply{Message: "UserInfo " + userInfo}, nil
}

func (s *Server) GetUserDetail(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoReply, error) {
	userInfo := GetUserById(in.Uid)
	return &pb.UserInfoReply{Message: "UserInfo " + userInfo}, nil
}

func GetUserById(uid int64) string {
	return "get UserInfo!"
}
