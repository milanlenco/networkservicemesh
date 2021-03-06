// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networkservice.proto

package networkservice

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	connection "github.com/networkservicemesh/networkservicemesh/controlplane/api/local/connection"
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

type NetworkServiceRequest struct {
	Connection           *connection.Connection  `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	MechanismPreferences []*connection.Mechanism `protobuf:"bytes,2,rep,name=mechanism_preferences,json=mechanismPreferences,proto3" json:"mechanism_preferences,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *NetworkServiceRequest) Reset()         { *m = NetworkServiceRequest{} }
func (m *NetworkServiceRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceRequest) ProtoMessage()    {}
func (*NetworkServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_361e8247d5a9945c, []int{0}
}

func (m *NetworkServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServiceRequest.Unmarshal(m, b)
}
func (m *NetworkServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServiceRequest.Marshal(b, m, deterministic)
}
func (m *NetworkServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServiceRequest.Merge(m, src)
}
func (m *NetworkServiceRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkServiceRequest.Size(m)
}
func (m *NetworkServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServiceRequest proto.InternalMessageInfo

func (m *NetworkServiceRequest) GetConnection() *connection.Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (m *NetworkServiceRequest) GetMechanismPreferences() []*connection.Mechanism {
	if m != nil {
		return m.MechanismPreferences
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkServiceRequest)(nil), "local.networkservice.NetworkServiceRequest")
}

func init() { proto.RegisterFile("networkservice.proto", fileDescriptor_361e8247d5a9945c) }

var fileDescriptor_361e8247d5a9945c = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcf, 0x4b, 0xfb, 0x30,
	0x14, 0xa7, 0xdf, 0x2f, 0x2a, 0x64, 0x30, 0x24, 0x74, 0x52, 0xaa, 0x87, 0xe1, 0x69, 0x20, 0xa4,
	0x50, 0xaf, 0x7a, 0x71, 0x78, 0x54, 0x46, 0xf5, 0xe4, 0x45, 0xda, 0xf0, 0xd6, 0x06, 0x93, 0xbc,
	0x98, 0xa4, 0xca, 0xfe, 0x20, 0xf1, 0xdf, 0x94, 0x2d, 0x6e, 0x6b, 0xb1, 0xec, 0x12, 0x1e, 0xef,
	0x7d, 0xde, 0xe7, 0xc7, 0x0b, 0x89, 0x35, 0xf8, 0x4f, 0xb4, 0x6f, 0x0e, 0xec, 0x87, 0xe0, 0xc0,
	0x8c, 0x45, 0x8f, 0x34, 0x96, 0xc8, 0x4b, 0xc9, 0xfa, 0xb3, 0x94, 0xd7, 0xc2, 0x37, 0x6d, 0xc5,
	0x38, 0xaa, 0xac, 0x3f, 0x52, 0xe0, 0x9a, 0xa1, 0x16, 0x47, 0xed, 0x2d, 0x4a, 0x23, 0x4b, 0x0d,
	0x59, 0x69, 0x44, 0xb6, 0xe1, 0x5d, 0xb7, 0x35, 0x70, 0x2f, 0x50, 0x77, 0xca, 0x20, 0x9d, 0x26,
	0xc6, 0xaf, 0x0c, 0xb8, 0x0c, 0x94, 0xf1, 0xab, 0xf0, 0x86, 0xc9, 0xe5, 0x77, 0x44, 0x26, 0x8f,
	0x41, 0xe3, 0x29, 0x68, 0x14, 0xf0, 0xde, 0x82, 0xf3, 0xf4, 0x86, 0x90, 0x3d, 0x4f, 0x12, 0x4d,
	0xa3, 0xd9, 0x28, 0xbf, 0x60, 0x21, 0x43, 0x47, 0x60, 0xbe, 0x2b, 0x8b, 0x0e, 0x9e, 0x2e, 0xc8,
	0x44, 0x01, 0x6f, 0x4a, 0x2d, 0x9c, 0x7a, 0x35, 0x16, 0x96, 0x60, 0x41, 0x73, 0x70, 0xc9, 0xbf,
	0xe9, 0xff, 0xd9, 0x28, 0x3f, 0xff, 0x4b, 0xf4, 0xb0, 0x85, 0x17, 0xf1, 0x6e, 0x73, 0xb1, 0x5f,
	0xcc, 0xbf, 0x22, 0x32, 0xee, 0x3b, 0xa5, 0xcf, 0xe4, 0x64, 0xeb, 0xf6, 0x8a, 0x0d, 0x5d, 0x97,
	0x0d, 0x46, 0x4b, 0x0f, 0xc6, 0xa0, 0xb7, 0xe4, 0x68, 0x2e, 0xd1, 0x01, 0x3d, 0x08, 0x4b, 0xcf,
	0x58, 0x8d, 0x58, 0xcb, 0xdf, 0xdf, 0xad, 0xda, 0x25, 0xbb, 0x5f, 0xdf, 0xf5, 0xee, 0xf4, 0x65,
	0xdc, 0x37, 0x51, 0x1d, 0x6f, 0x10, 0xd7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xce, 0xc9, 0x0d,
	0x33, 0x17, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkServiceClient is the client API for NetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServiceClient interface {
	Request(ctx context.Context, in *NetworkServiceRequest, opts ...grpc.CallOption) (*connection.Connection, error)
	Close(ctx context.Context, in *connection.Connection, opts ...grpc.CallOption) (*empty.Empty, error)
}

type networkServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceClient(cc *grpc.ClientConn) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) Request(ctx context.Context, in *NetworkServiceRequest, opts ...grpc.CallOption) (*connection.Connection, error) {
	out := new(connection.Connection)
	err := c.cc.Invoke(ctx, "/local.networkservice.NetworkService/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Close(ctx context.Context, in *connection.Connection, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/local.networkservice.NetworkService/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceServer is the server API for NetworkService service.
type NetworkServiceServer interface {
	Request(context.Context, *NetworkServiceRequest) (*connection.Connection, error)
	Close(context.Context, *connection.Connection) (*empty.Empty, error)
}

// UnimplementedNetworkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkServiceServer struct {
}

func (*UnimplementedNetworkServiceServer) Request(ctx context.Context, req *NetworkServiceRequest) (*connection.Connection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (*UnimplementedNetworkServiceServer) Close(ctx context.Context, req *connection.Connection) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}

func RegisterNetworkServiceServer(s *grpc.Server, srv NetworkServiceServer) {
	s.RegisterService(&_NetworkService_serviceDesc, srv)
}

func _NetworkService_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/local.networkservice.NetworkService/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Request(ctx, req.(*NetworkServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(connection.Connection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/local.networkservice.NetworkService/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Close(ctx, req.(*connection.Connection))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "local.networkservice.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _NetworkService_Request_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _NetworkService_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "networkservice.proto",
}
