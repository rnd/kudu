// Code generated by protoc-gen-go.
// source: item.proto
// DO NOT EDIT!

/*
Package item is a generated protocol buffer package.

It is generated from these files:
	item.proto

It has these top-level messages:
	Item
	ListRequest
	ListResponse
	AddRequest
	AddResponse
	GetRequest
	GetResponse
*/
package item

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import kudu_type "github.com/rnd/kudu/golang/protogen/type/date"

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

// Item holds item information.
type Item struct {
	Id      string          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Goal    string          `protobuf:"bytes,2,opt,name=goal" json:"goal,omitempty"`
	Url     string          `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Tag     string          `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
	Notes   string          `protobuf:"bytes,5,opt,name=notes" json:"notes,omitempty"`
	NotesMd string          `protobuf:"bytes,6,opt,name=notes_md,json=notesMd" json:"notes_md,omitempty"`
	Date    *kudu_type.Date `protobuf:"bytes,7,opt,name=date" json:"date,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetGoal() string {
	if m != nil {
		return m.Goal
	}
	return ""
}

func (m *Item) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Item) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Item) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

func (m *Item) GetNotesMd() string {
	if m != nil {
		return m.NotesMd
	}
	return ""
}

func (m *Item) GetDate() *kudu_type.Date {
	if m != nil {
		return m.Date
	}
	return nil
}

// ListRequest holds item list filter value.
type ListRequest struct {
	Goal string `protobuf:"bytes,1,opt,name=goal" json:"goal,omitempty"`
	Tag  string `protobuf:"bytes,2,opt,name=tag" json:"tag,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListRequest) GetGoal() string {
	if m != nil {
		return m.Goal
	}
	return ""
}

func (m *ListRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

// ListResponse holds item list information.
type ListResponse struct {
	Items []*Item `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListResponse) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

// AddRequest holds information of item that is going to be added.
type AddRequest struct {
	Item *Item `protobuf:"bytes,1,opt,name=item" json:"item,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AddRequest) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

// AddResponse holds newly added item id.
type AddResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *AddResponse) Reset()                    { *m = AddResponse{} }
func (m *AddResponse) String() string            { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()               {}
func (*AddResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AddResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// GetRequest holds item filter value.
type GetRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// GetResponse holds item information.
type GetResponse struct {
	Item *Item `protobuf:"bytes,1,opt,name=item" json:"item,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetResponse) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func init() {
	proto.RegisterType((*Item)(nil), "kudu.item.Item")
	proto.RegisterType((*ListRequest)(nil), "kudu.item.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "kudu.item.ListResponse")
	proto.RegisterType((*AddRequest)(nil), "kudu.item.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "kudu.item.AddResponse")
	proto.RegisterType((*GetRequest)(nil), "kudu.item.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "kudu.item.GetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ItemService service

type ItemServiceClient interface {
	ListItem(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	AddItem(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	GetItem(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type itemServiceClient struct {
	cc *grpc.ClientConn
}

func NewItemServiceClient(cc *grpc.ClientConn) ItemServiceClient {
	return &itemServiceClient{cc}
}

func (c *itemServiceClient) ListItem(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/kudu.item.ItemService/ListItem", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) AddItem(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := grpc.Invoke(ctx, "/kudu.item.ItemService/AddItem", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetItem(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/kudu.item.ItemService/GetItem", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ItemService service

type ItemServiceServer interface {
	ListItem(context.Context, *ListRequest) (*ListResponse, error)
	AddItem(context.Context, *AddRequest) (*AddResponse, error)
	GetItem(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterItemServiceServer(s *grpc.Server, srv ItemServiceServer) {
	s.RegisterService(&_ItemService_serviceDesc, srv)
}

func _ItemService_ListItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).ListItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kudu.item.ItemService/ListItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).ListItem(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kudu.item.ItemService/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).AddItem(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kudu.item.ItemService/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetItem(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ItemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kudu.item.ItemService",
	HandlerType: (*ItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListItem",
			Handler:    _ItemService_ListItem_Handler,
		},
		{
			MethodName: "AddItem",
			Handler:    _ItemService_AddItem_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _ItemService_GetItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "item.proto",
}

func init() { proto.RegisterFile("item.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x6d, 0xd2, 0xf4, 0x6b, 0x22, 0x56, 0x16, 0xad, 0x6b, 0x51, 0x28, 0x5b, 0x84, 0xe2, 0x21,
	0xc1, 0x16, 0x2f, 0x82, 0x87, 0x8a, 0x50, 0x04, 0xbd, 0xc4, 0x9b, 0x17, 0x49, 0xbb, 0x43, 0x0c,
	0x36, 0x49, 0x4d, 0x36, 0x82, 0xff, 0xc7, 0x9f, 0xe2, 0x0f, 0x93, 0x9d, 0xd4, 0x36, 0x6d, 0x3d,
	0x78, 0x9b, 0x79, 0x6f, 0x1f, 0xef, 0xcd, 0x4b, 0x00, 0x42, 0x85, 0x91, 0xb3, 0x48, 0x13, 0x95,
	0xb0, 0xd6, 0x5b, 0x2e, 0x73, 0x47, 0x03, 0xdd, 0xb6, 0xfa, 0x5c, 0xa0, 0x2b, 0x7d, 0x85, 0x05,
	0x27, 0xbe, 0x0c, 0xb0, 0xee, 0x15, 0x46, 0x6c, 0x1f, 0xcc, 0x50, 0x72, 0xa3, 0x67, 0x0c, 0x5a,
	0x9e, 0x19, 0x4a, 0xc6, 0xc0, 0x0a, 0x12, 0x7f, 0xce, 0x4d, 0x42, 0x68, 0x66, 0x07, 0x50, 0xcd,
	0xd3, 0x39, 0xaf, 0x12, 0xa4, 0x47, 0x8d, 0x28, 0x3f, 0xe0, 0x56, 0x81, 0x28, 0x3f, 0x60, 0x87,
	0x50, 0x8b, 0x13, 0x85, 0x19, 0xaf, 0x11, 0x56, 0x2c, 0xec, 0x04, 0x9a, 0x34, 0xbc, 0x44, 0x92,
	0xd7, 0x89, 0x68, 0xd0, 0xfe, 0x28, 0x59, 0x1f, 0x2c, 0x9d, 0x87, 0x37, 0x7a, 0xc6, 0xc0, 0x1e,
	0xb6, 0x1d, 0x0a, 0xab, 0x63, 0x3a, 0x77, 0xbe, 0x42, 0x8f, 0x48, 0x31, 0x02, 0xfb, 0x21, 0xcc,
	0x94, 0x87, 0xef, 0x39, 0x66, 0x6a, 0x15, 0xce, 0xd8, 0x0c, 0xa7, 0xa3, 0x98, 0xab, 0x28, 0xe2,
	0x0a, 0xf6, 0x0a, 0x51, 0xb6, 0x48, 0xe2, 0x0c, 0xd9, 0x39, 0xd4, 0x74, 0x09, 0x19, 0x37, 0x7a,
	0xd5, 0xb5, 0x15, 0x15, 0xa5, 0x2b, 0xf0, 0x0a, 0x56, 0x5c, 0x02, 0x8c, 0xa5, 0xfc, 0xb5, 0xea,
	0x83, 0xa5, 0x61, 0xb2, 0xfa, 0x43, 0x43, 0xa4, 0x38, 0x03, 0x9b, 0x24, 0x4b, 0xa3, 0xad, 0x2e,
	0xc5, 0x29, 0xc0, 0x04, 0x57, 0xe1, 0xb7, 0xd9, 0x21, 0xd8, 0xc4, 0x2e, 0xc5, 0xff, 0x31, 0x1c,
	0x7e, 0x1b, 0x60, 0xeb, 0xf5, 0x09, 0xd3, 0x8f, 0x70, 0x86, 0xec, 0x06, 0x9a, 0xfa, 0x54, 0xfa,
	0x92, 0x9d, 0x92, 0xa4, 0x54, 0x5a, 0xf7, 0x78, 0x07, 0x2f, 0x1c, 0x45, 0x85, 0x5d, 0x43, 0x63,
	0x2c, 0x25, 0xa9, 0x8f, 0x4a, 0xaf, 0xd6, 0x35, 0x74, 0x3b, 0xdb, 0x70, 0x59, 0x3b, 0x41, 0xb5,
	0xa3, 0x5d, 0x1f, 0xbc, 0xa1, 0x2d, 0x5d, 0x2a, 0x2a, 0xb7, 0x17, 0xcf, 0x83, 0x20, 0x54, 0xaf,
	0xf9, 0xd4, 0x99, 0x25, 0x91, 0x9b, 0xc6, 0xd2, 0xd5, 0x2f, 0xdd, 0x20, 0x99, 0xfb, 0x71, 0xe0,
	0xd2, 0x1f, 0x1a, 0x60, 0xec, 0x6a, 0xe5, 0xb4, 0x4e, 0xeb, 0xe8, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x7e, 0x1e, 0x8d, 0xe3, 0xda, 0x02, 0x00, 0x00,
}
