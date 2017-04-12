// Code generated by protoc-gen-go.
// source: page.proto
// DO NOT EDIT!

/*
Package page is a generated protocol buffer package.

It is generated from these files:
	page.proto

It has these top-level messages:
	PageResponse
	ItemListPageRequest
*/
package page

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import kudu_item "github.com/rnd/kudu/golang/protogen/item"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

// PageResponse holds html content and response meta data.
type PageResponse struct {
	Html      string                     `protobuf:"bytes,1,opt,name=html" json:"html,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *PageResponse) Reset()                    { *m = PageResponse{} }
func (m *PageResponse) String() string            { return proto.CompactTextString(m) }
func (*PageResponse) ProtoMessage()               {}
func (*PageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PageResponse) GetHtml() string {
	if m != nil {
		return m.Html
	}
	return ""
}

func (m *PageResponse) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

// ItemListPageRequest holds list of item.
type ItemListPageRequest struct {
	Item []*kudu_item.Item `protobuf:"bytes,1,rep,name=item" json:"item,omitempty"`
}

func (m *ItemListPageRequest) Reset()                    { *m = ItemListPageRequest{} }
func (m *ItemListPageRequest) String() string            { return proto.CompactTextString(m) }
func (*ItemListPageRequest) ProtoMessage()               {}
func (*ItemListPageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ItemListPageRequest) GetItem() []*kudu_item.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func init() {
	proto.RegisterType((*PageResponse)(nil), "kudu.page.PageResponse")
	proto.RegisterType((*ItemListPageRequest)(nil), "kudu.page.ItemListPageRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PageService service

type PageServiceClient interface {
	ItemPageList(ctx context.Context, in *ItemListPageRequest, opts ...grpc.CallOption) (PageService_ItemPageListClient, error)
}

type pageServiceClient struct {
	cc *grpc.ClientConn
}

func NewPageServiceClient(cc *grpc.ClientConn) PageServiceClient {
	return &pageServiceClient{cc}
}

func (c *pageServiceClient) ItemPageList(ctx context.Context, in *ItemListPageRequest, opts ...grpc.CallOption) (PageService_ItemPageListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PageService_serviceDesc.Streams[0], c.cc, "/kudu.page.PageService/ItemPageList", opts...)
	if err != nil {
		return nil, err
	}
	x := &pageServiceItemPageListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PageService_ItemPageListClient interface {
	Recv() (*PageResponse, error)
	grpc.ClientStream
}

type pageServiceItemPageListClient struct {
	grpc.ClientStream
}

func (x *pageServiceItemPageListClient) Recv() (*PageResponse, error) {
	m := new(PageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PageService service

type PageServiceServer interface {
	ItemPageList(*ItemListPageRequest, PageService_ItemPageListServer) error
}

func RegisterPageServiceServer(s *grpc.Server, srv PageServiceServer) {
	s.RegisterService(&_PageService_serviceDesc, srv)
}

func _PageService_ItemPageList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ItemListPageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PageServiceServer).ItemPageList(m, &pageServiceItemPageListServer{stream})
}

type PageService_ItemPageListServer interface {
	Send(*PageResponse) error
	grpc.ServerStream
}

type pageServiceItemPageListServer struct {
	grpc.ServerStream
}

func (x *pageServiceItemPageListServer) Send(m *PageResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _PageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kudu.page.PageService",
	HandlerType: (*PageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ItemPageList",
			Handler:       _PageService_ItemPageList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "page.proto",
}

func init() { proto.RegisterFile("page.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x50, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x50, 0x21, 0xc5, 0xa9, 0x54, 0xc9, 0x0c, 0x44, 0x19, 0x20, 0x0a, 0x4b, 0xc4, 0x60,
	0xa3, 0xb0, 0x20, 0x46, 0x36, 0x04, 0x03, 0x0a, 0x4c, 0x15, 0x4b, 0xd2, 0x1e, 0xae, 0x45, 0x1c,
	0x87, 0xd8, 0xe6, 0xf7, 0xa3, 0xb3, 0xd5, 0x92, 0xa1, 0x8b, 0xe5, 0xbb, 0xf7, 0x71, 0xf7, 0x8e,
	0xd2, 0xb1, 0x95, 0xc0, 0xc7, 0xc9, 0x38, 0xc3, 0x92, 0x6f, 0xbf, 0xf5, 0x1c, 0x1b, 0xf9, 0x4a,
	0x39, 0xd0, 0x02, 0x9f, 0x88, 0xe5, 0xd7, 0xd2, 0x18, 0xd9, 0x83, 0x08, 0x55, 0xe7, 0xbf, 0x84,
	0x53, 0x1a, 0xac, 0x6b, 0xf5, 0x18, 0x09, 0xe5, 0x27, 0x5d, 0xbe, 0xb5, 0x12, 0x1a, 0xb0, 0xa3,
	0x19, 0x2c, 0x30, 0x46, 0x17, 0x3b, 0xa7, 0xfb, 0x8c, 0x14, 0xa4, 0x4a, 0x9a, 0xf0, 0x67, 0x0f,
	0x34, 0x39, 0xc8, 0xb2, 0xd3, 0x82, 0x54, 0x69, 0x9d, 0xf3, 0x68, 0xcc, 0xf7, 0xc6, 0xfc, 0x63,
	0xcf, 0x68, 0xfe, 0xc9, 0xe5, 0x23, 0xbd, 0x78, 0x76, 0xa0, 0x5f, 0x95, 0x75, 0x71, 0xca, 0x8f,
	0x07, 0xeb, 0xd8, 0x0d, 0x5d, 0xe0, 0x8e, 0x19, 0x29, 0xce, 0xaa, 0xb4, 0x5e, 0xf1, 0x10, 0x20,
	0x6c, 0x8d, 0xec, 0x26, 0x80, 0xf5, 0x9a, 0xa6, 0xa8, 0x79, 0x87, 0xe9, 0x57, 0x6d, 0x80, 0xbd,
	0xd0, 0x25, 0x82, 0xd8, 0x42, 0x3b, 0x76, 0xc5, 0x0f, 0xb1, 0xf9, 0x91, 0x19, 0xf9, 0xe5, 0x0c,
	0x9f, 0x27, 0x2c, 0x4f, 0xee, 0xc8, 0xd3, 0xed, 0xba, 0x92, 0xca, 0xed, 0x7c, 0xc7, 0x37, 0x46,
	0x8b, 0x69, 0xd8, 0x0a, 0x24, 0x0b, 0x69, 0xfa, 0x76, 0x90, 0xf1, 0x58, 0x12, 0x06, 0x81, 0xe2,
	0xee, 0x3c, 0x94, 0xf7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x78, 0x3d, 0x9d, 0xc5, 0x73, 0x01,
	0x00, 0x00,
}
