// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/grpc-go-examples.proto

package grpc_go_examples

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Input struct {
	MyMessage            string   `protobuf:"bytes,1,opt,name=myMessage,proto3" json:"myMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_279ce708195876ca, []int{0}
}

func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetMyMessage() string {
	if m != nil {
		return m.MyMessage
	}
	return ""
}

type Output struct {
	MyResponse           string   `protobuf:"bytes,1,opt,name=myResponse,proto3" json:"myResponse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_279ce708195876ca, []int{1}
}

func (m *Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Output.Unmarshal(m, b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Output.Marshal(b, m, deterministic)
}
func (m *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(m, src)
}
func (m *Output) XXX_Size() int {
	return xxx_messageInfo_Output.Size(m)
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetMyResponse() string {
	if m != nil {
		return m.MyResponse
	}
	return ""
}

type MessageList struct {
	Output               []string `protobuf:"bytes,1,rep,name=Output,proto3" json:"Output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageList) Reset()         { *m = MessageList{} }
func (m *MessageList) String() string { return proto.CompactTextString(m) }
func (*MessageList) ProtoMessage()    {}
func (*MessageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_279ce708195876ca, []int{2}
}

func (m *MessageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageList.Unmarshal(m, b)
}
func (m *MessageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageList.Marshal(b, m, deterministic)
}
func (m *MessageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageList.Merge(m, src)
}
func (m *MessageList) XXX_Size() int {
	return xxx_messageInfo_MessageList.Size(m)
}
func (m *MessageList) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageList.DiscardUnknown(m)
}

var xxx_messageInfo_MessageList proto.InternalMessageInfo

func (m *MessageList) GetOutput() []string {
	if m != nil {
		return m.Output
	}
	return nil
}

func init() {
	proto.RegisterType((*Input)(nil), "grpcgoexamples.Input")
	proto.RegisterType((*Output)(nil), "grpcgoexamples.Output")
	proto.RegisterType((*MessageList)(nil), "grpcgoexamples.MessageList")
}

func init() {
	proto.RegisterFile("proto/grpc-go-examples.proto", fileDescriptor_279ce708195876ca)
}

var fileDescriptor_279ce708195876ca = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x4d, 0xcf, 0xd7, 0x4d, 0xad, 0x48, 0xcc, 0x2d, 0xc8,
	0x49, 0x2d, 0xd6, 0x03, 0x0b, 0x0b, 0xf1, 0x81, 0xc4, 0xd3, 0xf3, 0x61, 0xa2, 0x4a, 0xaa, 0x5c,
	0xac, 0x9e, 0x79, 0x05, 0xa5, 0x25, 0x42, 0x32, 0x5c, 0x9c, 0xb9, 0x95, 0xbe, 0xa9, 0xc5, 0xc5,
	0x89, 0xe9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x08, 0x01, 0x25, 0x0d, 0x2e, 0x36,
	0xff, 0xd2, 0x12, 0x90, 0x3a, 0x39, 0x2e, 0xae, 0xdc, 0xca, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc,
	0x62, 0x98, 0x42, 0x24, 0x11, 0x25, 0x55, 0x2e, 0x6e, 0xa8, 0x26, 0x9f, 0xcc, 0xe2, 0x12, 0x21,
	0x31, 0x98, 0x46, 0x09, 0x46, 0x05, 0x66, 0x0d, 0xce, 0x20, 0x28, 0xcf, 0x68, 0x16, 0x13, 0x17,
	0x67, 0x70, 0x26, 0xc8, 0x0d, 0x41, 0x01, 0xce, 0x42, 0x66, 0x5c, 0xcc, 0xee, 0xa9, 0x25, 0x42,
	0xa2, 0x7a, 0xa8, 0xae, 0xd3, 0x03, 0x3b, 0x4d, 0x4a, 0x0c, 0x5d, 0x18, 0x62, 0x86, 0x12, 0x83,
	0x90, 0x23, 0x17, 0x4f, 0x70, 0x6a, 0x51, 0x59, 0x6a, 0x51, 0x70, 0x49, 0x51, 0x6a, 0x62, 0x2e,
	0xc9, 0x06, 0x18, 0x30, 0x0a, 0xb9, 0x71, 0xf1, 0x38, 0xe7, 0x64, 0xa6, 0xe6, 0x95, 0xe0, 0x37,
	0x42, 0x1a, 0x5d, 0x18, 0xc9, 0x93, 0x4a, 0x0c, 0x1a, 0x8c, 0x42, 0x5e, 0x5c, 0xc2, 0x4e, 0x99,
	0x29, 0x99, 0x45, 0xa9, 0xc9, 0x25, 0x99, 0xf9, 0x79, 0x89, 0x39, 0x64, 0xba, 0x48, 0x83, 0xd1,
	0x80, 0xd1, 0x49, 0x28, 0x4a, 0x00, 0x3d, 0xfa, 0x92, 0xd8, 0xc0, 0xf1, 0x67, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xf0, 0xa8, 0x6d, 0x58, 0xdf, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SimpleRPCClient is the client API for SimpleRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SimpleRPCClient interface {
	Get(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error)
	ServerStream(ctx context.Context, in *Input, opts ...grpc.CallOption) (SimpleRPC_ServerStreamClient, error)
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (SimpleRPC_ClientStreamClient, error)
	BidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (SimpleRPC_BidirectionalStreamClient, error)
}

type simpleRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleRPCClient(cc grpc.ClientConnInterface) SimpleRPCClient {
	return &simpleRPCClient{cc}
}

func (c *simpleRPCClient) Get(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error) {
	out := new(Output)
	err := c.cc.Invoke(ctx, "/grpcgoexamples.SimpleRPC/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleRPCClient) ServerStream(ctx context.Context, in *Input, opts ...grpc.CallOption) (SimpleRPC_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SimpleRPC_serviceDesc.Streams[0], "/grpcgoexamples.SimpleRPC/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &simpleRPCServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SimpleRPC_ServerStreamClient interface {
	Recv() (*Output, error)
	grpc.ClientStream
}

type simpleRPCServerStreamClient struct {
	grpc.ClientStream
}

func (x *simpleRPCServerStreamClient) Recv() (*Output, error) {
	m := new(Output)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *simpleRPCClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (SimpleRPC_ClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SimpleRPC_serviceDesc.Streams[1], "/grpcgoexamples.SimpleRPC/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &simpleRPCClientStreamClient{stream}
	return x, nil
}

type SimpleRPC_ClientStreamClient interface {
	Send(*Input) error
	CloseAndRecv() (*MessageList, error)
	grpc.ClientStream
}

type simpleRPCClientStreamClient struct {
	grpc.ClientStream
}

func (x *simpleRPCClientStreamClient) Send(m *Input) error {
	return x.ClientStream.SendMsg(m)
}

func (x *simpleRPCClientStreamClient) CloseAndRecv() (*MessageList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessageList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *simpleRPCClient) BidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (SimpleRPC_BidirectionalStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SimpleRPC_serviceDesc.Streams[2], "/grpcgoexamples.SimpleRPC/BidirectionalStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &simpleRPCBidirectionalStreamClient{stream}
	return x, nil
}

type SimpleRPC_BidirectionalStreamClient interface {
	Send(*Input) error
	Recv() (*Output, error)
	grpc.ClientStream
}

type simpleRPCBidirectionalStreamClient struct {
	grpc.ClientStream
}

func (x *simpleRPCBidirectionalStreamClient) Send(m *Input) error {
	return x.ClientStream.SendMsg(m)
}

func (x *simpleRPCBidirectionalStreamClient) Recv() (*Output, error) {
	m := new(Output)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SimpleRPCServer is the server API for SimpleRPC service.
type SimpleRPCServer interface {
	Get(context.Context, *Input) (*Output, error)
	ServerStream(*Input, SimpleRPC_ServerStreamServer) error
	ClientStream(SimpleRPC_ClientStreamServer) error
	BidirectionalStream(SimpleRPC_BidirectionalStreamServer) error
}

// UnimplementedSimpleRPCServer can be embedded to have forward compatible implementations.
type UnimplementedSimpleRPCServer struct {
}

func (*UnimplementedSimpleRPCServer) Get(ctx context.Context, req *Input) (*Output, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedSimpleRPCServer) ServerStream(req *Input, srv SimpleRPC_ServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}
func (*UnimplementedSimpleRPCServer) ClientStream(srv SimpleRPC_ClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStream not implemented")
}
func (*UnimplementedSimpleRPCServer) BidirectionalStream(srv SimpleRPC_BidirectionalStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStream not implemented")
}

func RegisterSimpleRPCServer(s *grpc.Server, srv SimpleRPCServer) {
	s.RegisterService(&_SimpleRPC_serviceDesc, srv)
}

func _SimpleRPC_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleRPCServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcgoexamples.SimpleRPC/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleRPCServer).Get(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleRPC_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Input)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SimpleRPCServer).ServerStream(m, &simpleRPCServerStreamServer{stream})
}

type SimpleRPC_ServerStreamServer interface {
	Send(*Output) error
	grpc.ServerStream
}

type simpleRPCServerStreamServer struct {
	grpc.ServerStream
}

func (x *simpleRPCServerStreamServer) Send(m *Output) error {
	return x.ServerStream.SendMsg(m)
}

func _SimpleRPC_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SimpleRPCServer).ClientStream(&simpleRPCClientStreamServer{stream})
}

type SimpleRPC_ClientStreamServer interface {
	SendAndClose(*MessageList) error
	Recv() (*Input, error)
	grpc.ServerStream
}

type simpleRPCClientStreamServer struct {
	grpc.ServerStream
}

func (x *simpleRPCClientStreamServer) SendAndClose(m *MessageList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *simpleRPCClientStreamServer) Recv() (*Input, error) {
	m := new(Input)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SimpleRPC_BidirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SimpleRPCServer).BidirectionalStream(&simpleRPCBidirectionalStreamServer{stream})
}

type SimpleRPC_BidirectionalStreamServer interface {
	Send(*Output) error
	Recv() (*Input, error)
	grpc.ServerStream
}

type simpleRPCBidirectionalStreamServer struct {
	grpc.ServerStream
}

func (x *simpleRPCBidirectionalStreamServer) Send(m *Output) error {
	return x.ServerStream.SendMsg(m)
}

func (x *simpleRPCBidirectionalStreamServer) Recv() (*Input, error) {
	m := new(Input)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SimpleRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcgoexamples.SimpleRPC",
	HandlerType: (*SimpleRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SimpleRPC_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStream",
			Handler:       _SimpleRPC_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStream",
			Handler:       _SimpleRPC_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStream",
			Handler:       _SimpleRPC_BidirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/grpc-go-examples.proto",
}
