// Code generated by protoc-gen-go.
// source: add.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	add.proto

It has these top-level messages:
	SumRequest
	SumReply
*/
package pb

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

type SumRequest struct {
	A int64 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B int64 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
}

func (m *SumRequest) Reset()                    { *m = SumRequest{} }
func (m *SumRequest) String() string            { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()               {}
func (*SumRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SumReply struct {
	V   int64  `protobuf:"varint,1,opt,name=v" json:"v,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *SumReply) Reset()                    { *m = SumReply{} }
func (m *SumReply) String() string            { return proto.CompactTextString(m) }
func (*SumReply) ProtoMessage()               {}
func (*SumReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*SumRequest)(nil), "pb.SumRequest")
	proto.RegisterType((*SumReply)(nil), "pb.SumReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Add service

type AddClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error)
}

type addClient struct {
	cc *grpc.ClientConn
}

func NewAddClient(cc *grpc.ClientConn) AddClient {
	return &addClient{cc}
}

func (c *addClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error) {
	out := new(SumReply)
	err := grpc.Invoke(ctx, "/pb.Add/Sum", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Add service

type AddServer interface {
	Sum(context.Context, *SumRequest) (*SumReply, error)
}

func RegisterAddServer(s *grpc.Server, srv AddServer) {
	s.RegisterService(&_Add_serviceDesc, srv)
}

func _Add_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Add/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Add_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Add",
	HandlerType: (*AddServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Add_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("add.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x4c, 0x49, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe0, 0xe2, 0x0a, 0x2e, 0xcd,
	0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe1, 0x62, 0x4c, 0x94, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0e, 0x62, 0x4c, 0x04, 0xf1, 0x92, 0x24, 0x98, 0x20, 0xbc, 0x24, 0x25, 0x2d, 0x2e,
	0x0e, 0xb0, 0xca, 0x82, 0x9c, 0x4a, 0x90, 0x4c, 0x19, 0x4c, 0x5d, 0x99, 0x90, 0x00, 0x17, 0x73,
	0x6a, 0x51, 0x11, 0x58, 0x25, 0x67, 0x10, 0x88, 0x69, 0xa4, 0xc3, 0xc5, 0xec, 0x98, 0x92, 0x22,
	0xa4, 0xca, 0xc5, 0x0c, 0xd4, 0x22, 0xc4, 0xa7, 0x57, 0x90, 0xa4, 0x87, 0xb0, 0x45, 0x8a, 0x07,
	0xce, 0x07, 0x9a, 0xa5, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0x8e, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x49, 0xd1, 0xd9, 0x3f, 0x9b, 0x00, 0x00, 0x00,
}
