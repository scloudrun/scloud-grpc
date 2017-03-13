// Code generated by protoc-gen-go.
// source: userinfo.proto
// DO NOT EDIT!

/*
Package userinfo is a generated protocol buffer package.

It is generated from these files:
	userinfo.proto

It has these top-level messages:
	UserInfoRequest
	UserInfoReply
*/
package userinfo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type UserInfoRequest struct {
	Uid int64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *UserInfoRequest) Reset()                    { *m = UserInfoRequest{} }
func (m *UserInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*UserInfoRequest) ProtoMessage()               {}
func (*UserInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserInfoRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

// The response message containing the greetings
type UserInfoReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *UserInfoReply) Reset()                    { *m = UserInfoReply{} }
func (m *UserInfoReply) String() string            { return proto.CompactTextString(m) }
func (*UserInfoReply) ProtoMessage()               {}
func (*UserInfoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserInfoReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*UserInfoRequest)(nil), "userinfo.UserInfoRequest")
	proto.RegisterType((*UserInfoReply)(nil), "userinfo.UserInfoReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error)
	GetUserDetail(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := grpc.Invoke(ctx, "/userinfo.Greeter/GetUserInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetUserDetail(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := grpc.Invoke(ctx, "/userinfo.Greeter/GetUserDetail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	GetUserInfo(context.Context, *UserInfoRequest) (*UserInfoReply, error)
	GetUserDetail(context.Context, *UserInfoRequest) (*UserInfoReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userinfo.Greeter/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetUserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetUserDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetUserDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userinfo.Greeter/GetUserDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetUserDetail(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userinfo.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _Greeter_GetUserInfo_Handler,
		},
		{
			MethodName: "GetUserDetail",
			Handler:    _Greeter_GetUserDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userinfo.proto",
}

func init() { proto.RegisterFile("userinfo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2d, 0x4e, 0x2d,
	0xca, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x94,
	0xb9, 0xf8, 0x43, 0x8b, 0x53, 0x8b, 0x3c, 0xf3, 0xd2, 0xf2, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x84, 0x04, 0xb8, 0x98, 0x4b, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x40,
	0x4c, 0x25, 0x4d, 0x2e, 0x5e, 0x84, 0xa2, 0x82, 0x9c, 0x4a, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4,
	0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0xb0, 0x32, 0xce, 0x20, 0x18, 0xd7, 0x68, 0x2a, 0x23, 0x17, 0xbb,
	0x7b, 0x51, 0x6a, 0x6a, 0x49, 0x6a, 0x91, 0x90, 0x33, 0x17, 0xb7, 0x7b, 0x6a, 0x09, 0x4c, 0xa7,
	0x90, 0xa4, 0x1e, 0xdc, 0x15, 0x68, 0x56, 0x4a, 0x89, 0x63, 0x93, 0x2a, 0xc8, 0xa9, 0x54, 0x62,
	0x10, 0x72, 0xe5, 0xe2, 0x85, 0x1a, 0xe2, 0x92, 0x5a, 0x92, 0x98, 0x99, 0x43, 0x9e, 0x31, 0x4e,
	0x82, 0x4e, 0x70, 0x2f, 0x04, 0x80, 0x82, 0x20, 0x80, 0x31, 0x89, 0x0d, 0x1c, 0x16, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x3d, 0x6d, 0x14, 0x1d, 0x01, 0x00, 0x00,
}