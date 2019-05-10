// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usrmgn.proto

/*
Package usrmgn is a generated protocol buffer package.

It is generated from these files:
	usrmgn.proto

It has these top-level messages:
	LoginRequest
	LoginReply
	ProfileRequest
	NicknameRequest
	CommReply
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
	Username   string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password   string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Authorized bool   `protobuf:"varint,3,opt,name=authorized" json:"authorized,omitempty"`
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

func (m *LoginRequest) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
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

type ProfileRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Profile  string `protobuf:"bytes,2,opt,name=profile" json:"profile,omitempty"`
}

func (m *ProfileRequest) Reset()                    { *m = ProfileRequest{} }
func (m *ProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*ProfileRequest) ProtoMessage()               {}
func (*ProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProfileRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ProfileRequest) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

type NicknameRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
}

func (m *NicknameRequest) Reset()                    { *m = NicknameRequest{} }
func (m *NicknameRequest) String() string            { return proto.CompactTextString(m) }
func (*NicknameRequest) ProtoMessage()               {}
func (*NicknameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *NicknameRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *NicknameRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

type CommReply struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *CommReply) Reset()                    { *m = CommReply{} }
func (m *CommReply) String() string            { return proto.CompactTextString(m) }
func (*CommReply) ProtoMessage()               {}
func (*CommReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CommReply) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "usrmgn.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "usrmgn.LoginReply")
	proto.RegisterType((*ProfileRequest)(nil), "usrmgn.ProfileRequest")
	proto.RegisterType((*NicknameRequest)(nil), "usrmgn.NicknameRequest")
	proto.RegisterType((*CommReply)(nil), "usrmgn.CommReply")
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
	UploadProfile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*CommReply, error)
	ChangeNickname(ctx context.Context, in *NicknameRequest, opts ...grpc.CallOption) (*CommReply, error)
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

func (c *usrmgnClient) UploadProfile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*CommReply, error) {
	out := new(CommReply)
	err := grpc.Invoke(ctx, "/usrmgn.Usrmgn/UploadProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usrmgnClient) ChangeNickname(ctx context.Context, in *NicknameRequest, opts ...grpc.CallOption) (*CommReply, error) {
	out := new(CommReply)
	err := grpc.Invoke(ctx, "/usrmgn.Usrmgn/ChangeNickname", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Usrmgn service

type UsrmgnServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	UploadProfile(context.Context, *ProfileRequest) (*CommReply, error)
	ChangeNickname(context.Context, *NicknameRequest) (*CommReply, error)
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

func _Usrmgn_UploadProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsrmgnServer).UploadProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usrmgn.Usrmgn/UploadProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsrmgnServer).UploadProfile(ctx, req.(*ProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usrmgn_ChangeNickname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NicknameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsrmgnServer).ChangeNickname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usrmgn.Usrmgn/ChangeNickname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsrmgnServer).ChangeNickname(ctx, req.(*NicknameRequest))
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
		{
			MethodName: "UploadProfile",
			Handler:    _Usrmgn_UploadProfile_Handler,
		},
		{
			MethodName: "ChangeNickname",
			Handler:    _Usrmgn_ChangeNickname_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usrmgn.proto",
}

func init() { proto.RegisterFile("usrmgn.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x6d, 0xfa, 0xf1, 0xc5, 0x66, 0xa8, 0x15, 0x07, 0xa9, 0x21, 0x07, 0x29, 0xeb, 0xa5, 0xa7,
	0x1e, 0xec, 0x55, 0xbc, 0x14, 0x04, 0x41, 0x44, 0x02, 0xbd, 0x0a, 0xab, 0xd9, 0xa6, 0xc1, 0xcd,
	0xee, 0xba, 0x9b, 0x45, 0xea, 0x5f, 0xf3, 0xcf, 0x89, 0x9b, 0xdd, 0x92, 0x16, 0xc5, 0xde, 0xf2,
	0x66, 0x26, 0xf3, 0xe6, 0xbd, 0xb7, 0x30, 0xb4, 0x46, 0xd7, 0xa5, 0x98, 0x29, 0x2d, 0x1b, 0x89,
	0x71, 0x8b, 0xc8, 0x0a, 0x86, 0xf7, 0xb2, 0xac, 0x44, 0xce, 0xde, 0x2c, 0x33, 0x0d, 0x66, 0x30,
	0xb0, 0x86, 0x69, 0x41, 0x6b, 0x96, 0x46, 0x93, 0x68, 0x9a, 0xe4, 0x5b, 0xfc, 0xdd, 0x53, 0xd4,
	0x98, 0x77, 0xa9, 0x8b, 0xb4, 0xdf, 0xf6, 0x02, 0xc6, 0x0b, 0x00, 0x6a, 0x9b, 0xb5, 0xd4, 0xd5,
	0x07, 0x2b, 0xd2, 0x7f, 0x93, 0x68, 0x3a, 0xc8, 0x3b, 0x15, 0xf2, 0x04, 0xe0, 0x79, 0x14, 0xdf,
	0xfc, 0xc5, 0x22, 0xaa, 0x97, 0x57, 0xd7, 0xf3, 0x2c, 0x01, 0x63, 0x0a, 0x47, 0x4a, 0xcb, 0x55,
	0xc5, 0x99, 0xa3, 0x48, 0xf2, 0x00, 0xc9, 0x2d, 0x8c, 0x1e, 0xdb, 0xcf, 0x43, 0x94, 0x74, 0xf6,
	0xf4, 0x77, 0xf7, 0xdc, 0xc1, 0xc9, 0x83, 0x67, 0x3b, 0xd0, 0x92, 0xdf, 0x8e, 0x25, 0x97, 0x90,
	0x2c, 0x64, 0x5d, 0xb7, 0x8a, 0xc7, 0x10, 0x6b, 0x66, 0x2c, 0x6f, 0xfc, 0x0a, 0x8f, 0xae, 0x3e,
	0x23, 0x88, 0x97, 0x2e, 0x0a, 0x9c, 0xc3, 0x7f, 0x67, 0x11, 0x9e, 0xcd, 0x7c, 0x54, 0xdd, 0x64,
	0x32, 0xdc, 0xab, 0x2a, 0xbe, 0x21, 0x3d, 0xbc, 0x86, 0xe3, 0xa5, 0xe2, 0x92, 0x16, 0x5e, 0x3d,
	0x8e, 0xc3, 0xd8, 0xae, 0x1d, 0xd9, 0x69, 0xa8, 0x6f, 0x6f, 0x22, 0x3d, 0xbc, 0x81, 0xd1, 0x62,
	0x4d, 0x45, 0xc9, 0x82, 0x66, 0x3c, 0x0f, 0x63, 0x7b, 0x2e, 0xfc, 0xf8, 0xff, 0x73, 0xec, 0x1e,
	0xd3, 0xfc, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xd8, 0x26, 0xf1, 0x5c, 0x02, 0x00, 0x00,
}
