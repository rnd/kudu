// Code generated by protoc-gen-go.
// source: user/user.proto
// DO NOT EDIT!

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user/user.proto

It has these top-level messages:
	User
	RegisterRequest
	RegisterResponse
	LoginRequest
	LoginResponse
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import kudu_type "github.com/rnd/kudu/golang/protogen/type/creds"

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

// User hold user information.
type User struct {
	// id is user unique identifier.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// email is user email.
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	// username is user unique username.
	Username string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
	// first_name is user first name.
	FirstName string `protobuf:"bytes,4,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	// last_name is user last name.
	LastName string `protobuf:"bytes,5,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	// created is user registration timestamp
	Created *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=created" json:"created,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

// RegisterRequest holds user registration information.
type RegisterRequest struct {
	// credential is user credential.
	Credential *kudu_type.Credential `protobuf:"bytes,1,opt,name=credential" json:"credential,omitempty"`
	// user is user profile information.
	User *User `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterRequest) GetCredential() *kudu_type.Credential {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *RegisterRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// RegisterResponse holds the response of user registration.
type RegisterResponse struct {
	// status is the registration response status.
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	// token is valid jwt token.
	Token string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegisterResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *RegisterResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// LoginRequest holds login request information.
type LoginRequest struct {
	// credential is user claimed credential.
	Credential *kudu_type.Credential `protobuf:"bytes,1,opt,name=credential" json:"credential,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LoginRequest) GetCredential() *kudu_type.Credential {
	if m != nil {
		return m.Credential
	}
	return nil
}

// LoginResponse holds login request information.
type LoginResponse struct {
	// status is the registration response status.
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LoginResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "kudu.item.User")
	proto.RegisterType((*RegisterRequest)(nil), "kudu.item.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "kudu.item.RegisterResponse")
	proto.RegisterType((*LoginRequest)(nil), "kudu.item.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "kudu.item.LoginResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := grpc.Invoke(ctx, "/kudu.item.UserService/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/kudu.item.UserService/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kudu.item.UserService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kudu.item.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kudu.item.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0x87, 0x49, 0x69, 0x4b, 0x33, 0x01, 0x5a, 0x59, 0xfc, 0x89, 0x52, 0x21, 0xaa, 0x70, 0xa0,
	0xe2, 0xe0, 0x48, 0x01, 0x6e, 0x1c, 0x10, 0xa8, 0x37, 0xc4, 0x21, 0xc0, 0x85, 0x0b, 0x72, 0x93,
	0xa9, 0xb1, 0x9a, 0xd8, 0xc1, 0x76, 0x90, 0x78, 0x0d, 0x1e, 0x87, 0xa7, 0x43, 0x76, 0x9a, 0x6e,
	0xb4, 0xbb, 0xda, 0xcb, 0x5e, 0x22, 0xcd, 0x7c, 0xe3, 0x8c, 0xbf, 0x9f, 0x0c, 0xcb, 0xce, 0xa0,
	0xce, 0xdc, 0x87, 0xb6, 0x5a, 0x59, 0x45, 0xc2, 0x63, 0x57, 0x75, 0x54, 0x58, 0x6c, 0x92, 0xe7,
	0x5c, 0x29, 0x5e, 0x63, 0xe6, 0xc1, 0xbe, 0x3b, 0x64, 0x56, 0x34, 0x68, 0x2c, 0x6b, 0xda, 0x7e,
	0x36, 0x59, 0xd9, 0x3f, 0x2d, 0x66, 0xa5, 0xc6, 0xca, 0xf4, 0x9d, 0xf4, 0x5f, 0x00, 0xd3, 0x6f,
	0x06, 0x35, 0x79, 0x08, 0x13, 0x51, 0xc5, 0xc1, 0x26, 0xd8, 0x86, 0xc5, 0x44, 0x54, 0xe4, 0x11,
	0xcc, 0xb0, 0x61, 0xa2, 0x8e, 0x27, 0xbe, 0xd5, 0x17, 0x24, 0x81, 0x85, 0x5b, 0x2d, 0x59, 0x83,
	0xf1, 0x5d, 0x0f, 0xce, 0x35, 0x79, 0x06, 0x70, 0x10, 0xda, 0xd8, 0x1f, 0x9e, 0x4e, 0x3d, 0x0d,
	0x7d, 0xe7, 0xb3, 0xc3, 0x6b, 0x08, 0x6b, 0x36, 0xd0, 0x59, 0x7f, 0xd6, 0x35, 0x3c, 0x7c, 0x03,
	0xf7, 0x4a, 0x8d, 0xcc, 0x62, 0x15, 0xcf, 0x37, 0xc1, 0x36, 0xca, 0x13, 0xda, 0xbb, 0xd0, 0xc1,
	0x85, 0x7e, 0x1d, 0x5c, 0x8a, 0x61, 0x34, 0x6d, 0x60, 0x59, 0x20, 0x17, 0xc6, 0xa2, 0x2e, 0xf0,
	0x57, 0x87, 0xc6, 0x92, 0xb7, 0x00, 0x4e, 0x0f, 0xa5, 0x15, 0xac, 0xf6, 0x3a, 0x51, 0xfe, 0x98,
	0xfa, 0x88, 0x9c, 0x3b, 0xfd, 0x78, 0x86, 0xc5, 0x68, 0x90, 0xbc, 0x80, 0xa9, 0xf3, 0xf0, 0xb2,
	0x51, 0xbe, 0xa4, 0xe7, 0x4c, 0xa9, 0x0b, 0xa7, 0xf0, 0x30, 0x7d, 0x0f, 0xab, 0x8b, 0x75, 0xa6,
	0x55, 0xd2, 0x20, 0x79, 0x02, 0x73, 0x63, 0x99, 0xed, 0xcc, 0x29, 0xba, 0x53, 0xe5, 0xe2, 0xb3,
	0xea, 0x88, 0x72, 0x88, 0xcf, 0x17, 0xe9, 0x0e, 0xee, 0x7f, 0x52, 0x5c, 0xc8, 0xdb, 0xdd, 0x36,
	0x7d, 0x09, 0x0f, 0x4e, 0xbf, 0xb9, 0xf9, 0x16, 0xf9, 0xdf, 0x00, 0x22, 0x27, 0xf0, 0x05, 0xf5,
	0x6f, 0x51, 0x22, 0xd9, 0xc1, 0x62, 0x30, 0x20, 0xc9, 0x48, 0xf2, 0x52, 0x8a, 0xc9, 0xfa, 0x5a,
	0xd6, 0x2f, 0x4b, 0xef, 0x90, 0x77, 0x30, 0xf3, 0xfb, 0xc9, 0xd3, 0xd1, 0xdc, 0x58, 0x2c, 0x89,
	0xaf, 0x82, 0xe1, 0xf4, 0x87, 0x57, 0xdf, 0xb7, 0x5c, 0xd8, 0x9f, 0xdd, 0x9e, 0x96, 0xaa, 0xc9,
	0xb4, 0xac, 0x32, 0x37, 0x9b, 0x71, 0x55, 0x33, 0xc9, 0xfb, 0xb7, 0xcb, 0x51, 0xfa, 0x27, 0xbe,
	0x9f, 0xfb, 0xf2, 0xf5, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xe1, 0x58, 0x0e, 0xf6, 0x02,
	0x00, 0x00,
}
