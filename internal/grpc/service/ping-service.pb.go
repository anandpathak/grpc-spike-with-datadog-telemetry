// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demogrpc/internal/proto-files/services/ping-service.proto

package service

import (
	context "context"
	domain "demogrpc/internal/grpc/domain"
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

type PingResponse struct {
	Response             *domain.PingMessage `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_437d6b83a729624d, []int{0}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetResponse() *domain.PingMessage {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*PingResponse)(nil), "service.PingResponse")
}

func init() {
	proto.RegisterFile("demogrpc/internal/proto-files/services/ping-service.proto", fileDescriptor_437d6b83a729624d)
}

var fileDescriptor_437d6b83a729624d = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4c, 0x49, 0xcd, 0xcd,
	0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x28,
	0xca, 0x2f, 0xc9, 0xd7, 0x4d, 0xcb, 0xcc, 0x49, 0x2d, 0xd6, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0x2d, 0xd6, 0x2f, 0xc8, 0xcc, 0x4b, 0xd7, 0x85, 0xf2, 0xf4, 0xc0, 0x2a, 0x84, 0xd8, 0xa1,
	0x5c, 0x29, 0x7d, 0xfc, 0x66, 0xa4, 0xe4, 0xe7, 0x26, 0x66, 0xe6, 0x81, 0x4d, 0x80, 0xe8, 0x54,
	0xb2, 0xe7, 0xe2, 0x09, 0xc8, 0xcc, 0x4b, 0x0f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15,
	0xd2, 0xe7, 0xe2, 0x80, 0xb1, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x84, 0xf5, 0x20, 0xba,
	0xf4, 0x40, 0xea, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x83, 0xe0, 0x8a, 0x8c, 0x1c, 0xb9,
	0xb8, 0x41, 0x12, 0xc1, 0x10, 0x07, 0x08, 0x19, 0x71, 0xb1, 0x80, 0xb8, 0x42, 0xd8, 0x74, 0x49,
	0x89, 0xea, 0xc1, 0x9c, 0x8d, 0x6c, 0xa7, 0x93, 0x42, 0x94, 0x1c, 0xa6, 0xb3, 0xc1, 0x3c, 0xa8,
	0xf2, 0x24, 0x36, 0xb0, 0x63, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xff, 0xe6, 0x5e, 0xdb,
	0x23, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PingServiceClient is the client API for PingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingServiceClient interface {
	Ping(ctx context.Context, in *domain.PingMessage, opts ...grpc.CallOption) (*PingResponse, error)
}

type pingServiceClient struct {
	cc *grpc.ClientConn
}

func NewPingServiceClient(cc *grpc.ClientConn) PingServiceClient {
	return &pingServiceClient{cc}
}

func (c *pingServiceClient) Ping(ctx context.Context, in *domain.PingMessage, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/service.PingService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingServiceServer is the server API for PingService service.
type PingServiceServer interface {
	Ping(context.Context, *domain.PingMessage) (*PingResponse, error)
}

// UnimplementedPingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPingServiceServer struct {
}

func (*UnimplementedPingServiceServer) Ping(ctx context.Context, req *domain.PingMessage) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterPingServiceServer(s *grpc.Server, srv PingServiceServer) {
	s.RegisterService(&_PingService_serviceDesc, srv)
}

func _PingService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.PingMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.PingService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServiceServer).Ping(ctx, req.(*domain.PingMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.PingService",
	HandlerType: (*PingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _PingService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demogrpc/internal/proto-files/services/ping-service.proto",
}
