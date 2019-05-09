// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usrmgn.proto

/*
Package usrmgn is a generated protocol buffer package.

It is generated from these files:
	usrmgn.proto

It has these top-level messages:
	LoginRequest
	LoginReply
*/
package usrmgn

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

type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginReply struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	Profile  string `protobuf:"bytes,3,opt,name=profile" json:"profile,omitempty"`
}

func (m *LoginReply) Reset()                    { *m = LoginReply{} }
func (m *LoginReply) String() string            { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()               {}
func (*LoginReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginReply) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginReply) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *LoginReply) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "usrmgn.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "usrmgn.LoginReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Usrmgn service

type UsrmgnClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type usrmgnClient struct {
	cc *grpc.ClientConn
}

func NewUsrmgnClient(cc *grpc.ClientConn) UsrmgnClient {
	return &usrmgnClient{cc}
}

func (c *usrmgnClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/usrmgn.Usrmgn/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Usrmgn service

type UsrmgnServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
}

func RegisterUsrmgnServer(s *grpc.Server, srv UsrmgnServer) {
	s.RegisterService(&_Usrmgn_serviceDesc, srv)
}

func _Usrmgn_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsrmgnServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usrmgn.Usrmgn/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsrmgnServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Usrmgn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "usrmgn.Usrmgn",
	HandlerType: (*UsrmgnServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Usrmgn_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usrmgn.proto",
}

func init() { proto.RegisterFile("usrmgn.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2d, 0x2e, 0xca,
	0x4d, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xdc, 0xb8, 0x78,
	0x7c, 0xf2, 0xd3, 0x33, 0xf3, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xa4, 0xb8, 0x38,
	0x4a, 0x8b, 0x53, 0x8b, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0,
	0x7c, 0x90, 0x5c, 0x41, 0x62, 0x71, 0x71, 0x79, 0x7e, 0x51, 0x8a, 0x04, 0x13, 0x44, 0x0e, 0xc6,
	0x57, 0x8a, 0xe3, 0xe2, 0x82, 0x9a, 0x53, 0x90, 0x53, 0x49, 0xc8, 0x94, 0xbc, 0xcc, 0xe4, 0x6c,
	0xb0, 0x1c, 0xd4, 0x14, 0x18, 0x5f, 0x48, 0x82, 0x8b, 0xbd, 0xa0, 0x28, 0x3f, 0x2d, 0x33, 0x27,
	0x55, 0x82, 0x19, 0x2c, 0x05, 0xe3, 0x1a, 0xd9, 0x72, 0xb1, 0x85, 0x82, 0x5d, 0x2c, 0x64, 0xcc,
	0xc5, 0x0a, 0xb6, 0x49, 0x48, 0x44, 0x0f, 0xea, 0x23, 0x64, 0x0f, 0x48, 0x09, 0xa1, 0x89, 0x16,
	0xe4, 0x54, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0x7d, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x37,
	0xaa, 0xe5, 0xe0, 0x05, 0x01, 0x00, 0x00,
}
