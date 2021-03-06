// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/razeencheng/demo-go/grpc/demo3/helloworld/hello_world.proto

package helloworld // import "github.com/razeencheng/demo-go/grpc/demo3/helloworld"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

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

type HelloWorldRequest struct {
	Greeting             string            `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	Infos                map[string]string `protobuf:"bytes,2,rep,name=infos,proto3" json:"infos,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HelloWorldRequest) Reset()         { *m = HelloWorldRequest{} }
func (m *HelloWorldRequest) String() string { return proto.CompactTextString(m) }
func (*HelloWorldRequest) ProtoMessage()    {}
func (*HelloWorldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_world_718a08d278cf1204, []int{0}
}
func (m *HelloWorldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorldRequest.Unmarshal(m, b)
}
func (m *HelloWorldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorldRequest.Marshal(b, m, deterministic)
}
func (dst *HelloWorldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldRequest.Merge(dst, src)
}
func (m *HelloWorldRequest) XXX_Size() int {
	return xxx_messageInfo_HelloWorldRequest.Size(m)
}
func (m *HelloWorldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldRequest proto.InternalMessageInfo

func (m *HelloWorldRequest) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func (m *HelloWorldRequest) GetInfos() map[string]string {
	if m != nil {
		return m.Infos
	}
	return nil
}

type HelloWorldResponse struct {
	Reply                string     `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	Details              []*any.Any `protobuf:"bytes,2,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HelloWorldResponse) Reset()         { *m = HelloWorldResponse{} }
func (m *HelloWorldResponse) String() string { return proto.CompactTextString(m) }
func (*HelloWorldResponse) ProtoMessage()    {}
func (*HelloWorldResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_world_718a08d278cf1204, []int{1}
}
func (m *HelloWorldResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorldResponse.Unmarshal(m, b)
}
func (m *HelloWorldResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorldResponse.Marshal(b, m, deterministic)
}
func (dst *HelloWorldResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldResponse.Merge(dst, src)
}
func (m *HelloWorldResponse) XXX_Size() int {
	return xxx_messageInfo_HelloWorldResponse.Size(m)
}
func (m *HelloWorldResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldResponse proto.InternalMessageInfo

func (m *HelloWorldResponse) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func (m *HelloWorldResponse) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

type HelloWorld struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloWorld) Reset()         { *m = HelloWorld{} }
func (m *HelloWorld) String() string { return proto.CompactTextString(m) }
func (*HelloWorld) ProtoMessage()    {}
func (*HelloWorld) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_world_718a08d278cf1204, []int{2}
}
func (m *HelloWorld) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorld.Unmarshal(m, b)
}
func (m *HelloWorld) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorld.Marshal(b, m, deterministic)
}
func (dst *HelloWorld) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorld.Merge(dst, src)
}
func (m *HelloWorld) XXX_Size() int {
	return xxx_messageInfo_HelloWorld.Size(m)
}
func (m *HelloWorld) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorld.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorld proto.InternalMessageInfo

func (m *HelloWorld) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Error struct {
	Msg                  []string `protobuf:"bytes,1,rep,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_world_718a08d278cf1204, []int{3}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (dst *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(dst, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetMsg() []string {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloWorldRequest)(nil), "helloworld.HelloWorldRequest")
	proto.RegisterMapType((map[string]string)(nil), "helloworld.HelloWorldRequest.InfosEntry")
	proto.RegisterType((*HelloWorldResponse)(nil), "helloworld.HelloWorldResponse")
	proto.RegisterType((*HelloWorld)(nil), "helloworld.HelloWorld")
	proto.RegisterType((*Error)(nil), "helloworld.Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloWorldServiceClient is the client API for HelloWorldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloWorldServiceClient interface {
	SayHelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
	ListHello(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (HelloWorldService_ListHelloClient, error)
	SayMoreHello(ctx context.Context, opts ...grpc.CallOption) (HelloWorldService_SayMoreHelloClient, error)
	SayHelloChat(ctx context.Context, opts ...grpc.CallOption) (HelloWorldService_SayHelloChatClient, error)
}

type helloWorldServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloWorldServiceClient(cc *grpc.ClientConn) HelloWorldServiceClient {
	return &helloWorldServiceClient{cc}
}

func (c *helloWorldServiceClient) SayHelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, "/helloworld.HelloWorldService/SayHelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloWorldServiceClient) ListHello(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (HelloWorldService_ListHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloWorldService_serviceDesc.Streams[0], "/helloworld.HelloWorldService/ListHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloWorldServiceListHelloClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloWorldService_ListHelloClient interface {
	Recv() (*HelloWorldResponse, error)
	grpc.ClientStream
}

type helloWorldServiceListHelloClient struct {
	grpc.ClientStream
}

func (x *helloWorldServiceListHelloClient) Recv() (*HelloWorldResponse, error) {
	m := new(HelloWorldResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloWorldServiceClient) SayMoreHello(ctx context.Context, opts ...grpc.CallOption) (HelloWorldService_SayMoreHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloWorldService_serviceDesc.Streams[1], "/helloworld.HelloWorldService/SayMoreHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloWorldServiceSayMoreHelloClient{stream}
	return x, nil
}

type HelloWorldService_SayMoreHelloClient interface {
	Send(*HelloWorldRequest) error
	CloseAndRecv() (*HelloWorldResponse, error)
	grpc.ClientStream
}

type helloWorldServiceSayMoreHelloClient struct {
	grpc.ClientStream
}

func (x *helloWorldServiceSayMoreHelloClient) Send(m *HelloWorldRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloWorldServiceSayMoreHelloClient) CloseAndRecv() (*HelloWorldResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloWorldResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloWorldServiceClient) SayHelloChat(ctx context.Context, opts ...grpc.CallOption) (HelloWorldService_SayHelloChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloWorldService_serviceDesc.Streams[2], "/helloworld.HelloWorldService/SayHelloChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloWorldServiceSayHelloChatClient{stream}
	return x, nil
}

type HelloWorldService_SayHelloChatClient interface {
	Send(*HelloWorldRequest) error
	Recv() (*HelloWorldRequest, error)
	grpc.ClientStream
}

type helloWorldServiceSayHelloChatClient struct {
	grpc.ClientStream
}

func (x *helloWorldServiceSayHelloChatClient) Send(m *HelloWorldRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloWorldServiceSayHelloChatClient) Recv() (*HelloWorldRequest, error) {
	m := new(HelloWorldRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloWorldServiceServer is the server API for HelloWorldService service.
type HelloWorldServiceServer interface {
	SayHelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
	ListHello(*HelloWorldRequest, HelloWorldService_ListHelloServer) error
	SayMoreHello(HelloWorldService_SayMoreHelloServer) error
	SayHelloChat(HelloWorldService_SayHelloChatServer) error
}

func RegisterHelloWorldServiceServer(s *grpc.Server, srv HelloWorldServiceServer) {
	s.RegisterService(&_HelloWorldService_serviceDesc, srv)
}

func _HelloWorldService_SayHelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServiceServer).SayHelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.HelloWorldService/SayHelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServiceServer).SayHelloWorld(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloWorldService_ListHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloWorldRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloWorldServiceServer).ListHello(m, &helloWorldServiceListHelloServer{stream})
}

type HelloWorldService_ListHelloServer interface {
	Send(*HelloWorldResponse) error
	grpc.ServerStream
}

type helloWorldServiceListHelloServer struct {
	grpc.ServerStream
}

func (x *helloWorldServiceListHelloServer) Send(m *HelloWorldResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HelloWorldService_SayMoreHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloWorldServiceServer).SayMoreHello(&helloWorldServiceSayMoreHelloServer{stream})
}

type HelloWorldService_SayMoreHelloServer interface {
	SendAndClose(*HelloWorldResponse) error
	Recv() (*HelloWorldRequest, error)
	grpc.ServerStream
}

type helloWorldServiceSayMoreHelloServer struct {
	grpc.ServerStream
}

func (x *helloWorldServiceSayMoreHelloServer) SendAndClose(m *HelloWorldResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloWorldServiceSayMoreHelloServer) Recv() (*HelloWorldRequest, error) {
	m := new(HelloWorldRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloWorldService_SayHelloChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloWorldServiceServer).SayHelloChat(&helloWorldServiceSayHelloChatServer{stream})
}

type HelloWorldService_SayHelloChatServer interface {
	Send(*HelloWorldRequest) error
	Recv() (*HelloWorldRequest, error)
	grpc.ServerStream
}

type helloWorldServiceSayHelloChatServer struct {
	grpc.ServerStream
}

func (x *helloWorldServiceSayHelloChatServer) Send(m *HelloWorldRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloWorldServiceSayHelloChatServer) Recv() (*HelloWorldRequest, error) {
	m := new(HelloWorldRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _HelloWorldService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.HelloWorldService",
	HandlerType: (*HelloWorldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHelloWorld",
			Handler:    _HelloWorldService_SayHelloWorld_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListHello",
			Handler:       _HelloWorldService_ListHello_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayMoreHello",
			Handler:       _HelloWorldService_SayMoreHello_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloChat",
			Handler:       _HelloWorldService_SayHelloChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/razeencheng/demo-go/grpc/demo3/helloworld/hello_world.proto",
}

func init() {
	proto.RegisterFile("github.com/razeencheng/demo-go/grpc/demo3/helloworld/hello_world.proto", fileDescriptor_hello_world_718a08d278cf1204)
}

var fileDescriptor_hello_world_718a08d278cf1204 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x51, 0xab, 0xd3, 0x30,
	0x14, 0xc7, 0x6d, 0xc7, 0xd4, 0x7b, 0x54, 0xd0, 0x70, 0x1f, 0x66, 0xc1, 0xcb, 0xa5, 0x4f, 0x7d,
	0x31, 0x19, 0x9b, 0xc8, 0xf0, 0x41, 0x50, 0x99, 0x28, 0xa8, 0x68, 0xf7, 0x20, 0xec, 0x45, 0xb2,
	0xee, 0x2c, 0x2d, 0x66, 0x49, 0x4d, 0xd3, 0x49, 0xfd, 0x46, 0x7e, 0x27, 0x3f, 0x8c, 0x24, 0x5d,
	0x6d, 0x41, 0x9c, 0xa0, 0xf7, 0xa1, 0x70, 0x4e, 0xfe, 0xe7, 0xff, 0x3b, 0x39, 0xa7, 0x04, 0x5e,
	0x8a, 0xc2, 0xe6, 0xf5, 0x86, 0x66, 0x7a, 0xcf, 0x0c, 0xff, 0x86, 0xa8, 0xb2, 0x1c, 0x95, 0x60,
	0x5b, 0xdc, 0xeb, 0x87, 0x42, 0x33, 0x61, 0xca, 0xcc, 0x27, 0x73, 0x96, 0xa3, 0x94, 0xfa, 0xab,
	0x36, 0x72, 0xdb, 0x86, 0x9f, 0x7c, 0x4c, 0x4b, 0xa3, 0xad, 0x26, 0xd0, 0xab, 0x11, 0x1b, 0x30,
	0x85, 0x96, 0x5c, 0x09, 0xe6, 0x8b, 0x36, 0xf5, 0x8e, 0x95, 0xb6, 0x29, 0xb1, 0x62, 0x5c, 0x35,
	0xee, 0x6b, 0xcd, 0xf1, 0xf7, 0x00, 0xee, 0xbd, 0x72, 0xfe, 0x8f, 0xce, 0x9f, 0xe2, 0x97, 0x1a,
	0x2b, 0x4b, 0x22, 0xb8, 0x29, 0x0c, 0xa2, 0x2d, 0x94, 0x98, 0x04, 0x97, 0x41, 0x72, 0x96, 0xfe,
	0xca, 0xc9, 0x53, 0x18, 0x17, 0x6a, 0xa7, 0xab, 0x49, 0x78, 0x39, 0x4a, 0x6e, 0xcd, 0x12, 0xda,
	0xb7, 0xa7, 0xbf, 0x91, 0xe8, 0x6b, 0x57, 0xba, 0x54, 0xd6, 0x34, 0x69, 0x6b, 0x8b, 0x16, 0x00,
	0xfd, 0x21, 0xb9, 0x0b, 0xa3, 0xcf, 0xd8, 0x1c, 0x9b, 0xb8, 0x90, 0x9c, 0xc3, 0xf8, 0xc0, 0x65,
	0x8d, 0x93, 0xd0, 0x9f, 0xb5, 0xc9, 0x93, 0x70, 0x11, 0xc4, 0x6b, 0x20, 0xc3, 0x06, 0x55, 0xa9,
	0x55, 0x85, 0xae, 0xde, 0x60, 0x29, 0x3b, 0x46, 0x9b, 0x10, 0x0a, 0x37, 0xb6, 0x68, 0x79, 0x21,
	0xbb, 0x7b, 0x9e, 0x53, 0xa1, 0xb5, 0x90, 0x48, 0xbb, 0x7d, 0xd0, 0x67, 0xaa, 0x49, 0xbb, 0xa2,
	0xf8, 0x02, 0xa0, 0x67, 0xbb, 0x5b, 0xed, 0xab, 0x6e, 0x74, 0x17, 0xc6, 0xf7, 0x61, 0xbc, 0x34,
	0x46, 0x9b, 0x5e, 0x1a, 0x1d, 0xa5, 0xd9, 0x8f, 0x70, 0xb8, 0xc2, 0x15, 0x9a, 0x43, 0x91, 0x21,
	0x79, 0x0f, 0x77, 0x56, 0xbc, 0x19, 0x30, 0x1f, 0x9c, 0x5c, 0x54, 0x74, 0xf1, 0x27, 0xb9, 0x1d,
	0x33, 0xbe, 0x46, 0xde, 0xc1, 0xd9, 0x9b, 0xa2, 0xb2, 0x5e, 0xfb, 0x6f, 0xda, 0x34, 0x20, 0x1f,
	0xe0, 0xf6, 0x8a, 0x37, 0x6f, 0xb5, 0xc1, 0xab, 0x41, 0x26, 0x01, 0x49, 0x3d, 0xd2, 0x8b, 0x2f,
	0x72, 0x6e, 0xff, 0x86, 0x3c, 0x2d, 0x3b, 0xe2, 0x34, 0x78, 0xfe, 0x78, 0xfd, 0xe8, 0x5f, 0x1e,
	0xca, 0xe6, 0xba, 0xff, 0xd1, 0xf3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x17, 0x20, 0x36,
	0x67, 0x03, 0x00, 0x00,
}
