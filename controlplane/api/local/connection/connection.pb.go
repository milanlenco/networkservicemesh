// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connection.proto

package connection

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	connectioncontext "github.com/networkservicemesh/networkservicemesh/controlplane/api/connectioncontext"
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

type MechanismType int32

const (
	MechanismType_DEFAULT_INTERFACE MechanismType = 0
	MechanismType_KERNEL_INTERFACE  MechanismType = 1
	MechanismType_VHOST_INTERFACE   MechanismType = 2
	MechanismType_MEM_INTERFACE     MechanismType = 3
	MechanismType_SRIOV_INTERFACE   MechanismType = 4
	MechanismType_HW_INTERFACE      MechanismType = 5
)

var MechanismType_name = map[int32]string{
	0: "DEFAULT_INTERFACE",
	1: "KERNEL_INTERFACE",
	2: "VHOST_INTERFACE",
	3: "MEM_INTERFACE",
	4: "SRIOV_INTERFACE",
	5: "HW_INTERFACE",
}

var MechanismType_value = map[string]int32{
	"DEFAULT_INTERFACE": 0,
	"KERNEL_INTERFACE":  1,
	"VHOST_INTERFACE":   2,
	"MEM_INTERFACE":     3,
	"SRIOV_INTERFACE":   4,
	"HW_INTERFACE":      5,
}

func (x MechanismType) String() string {
	return proto.EnumName(MechanismType_name, int32(x))
}

func (MechanismType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{0}
}

type State int32

const (
	State_UP   State = 0
	State_DOWN State = 1
)

var State_name = map[int32]string{
	0: "UP",
	1: "DOWN",
}

var State_value = map[string]int32{
	"UP":   0,
	"DOWN": 1,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{1}
}

type ConnectionEventType int32

const (
	ConnectionEventType_INITIAL_STATE_TRANSFER ConnectionEventType = 0
	ConnectionEventType_UPDATE                 ConnectionEventType = 1
	ConnectionEventType_DELETE                 ConnectionEventType = 2
)

var ConnectionEventType_name = map[int32]string{
	0: "INITIAL_STATE_TRANSFER",
	1: "UPDATE",
	2: "DELETE",
}

var ConnectionEventType_value = map[string]int32{
	"INITIAL_STATE_TRANSFER": 0,
	"UPDATE":                 1,
	"DELETE":                 2,
}

func (x ConnectionEventType) String() string {
	return proto.EnumName(ConnectionEventType_name, int32(x))
}

func (ConnectionEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{2}
}

type Mechanism struct {
	Type                 MechanismType     `protobuf:"varint,1,opt,name=type,proto3,enum=local.connection.MechanismType" json:"type,omitempty"`
	Parameters           map[string]string `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Mechanism) Reset()         { *m = Mechanism{} }
func (m *Mechanism) String() string { return proto.CompactTextString(m) }
func (*Mechanism) ProtoMessage()    {}
func (*Mechanism) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{0}
}

func (m *Mechanism) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mechanism.Unmarshal(m, b)
}
func (m *Mechanism) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mechanism.Marshal(b, m, deterministic)
}
func (m *Mechanism) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mechanism.Merge(m, src)
}
func (m *Mechanism) XXX_Size() int {
	return xxx_messageInfo_Mechanism.Size(m)
}
func (m *Mechanism) XXX_DiscardUnknown() {
	xxx_messageInfo_Mechanism.DiscardUnknown(m)
}

var xxx_messageInfo_Mechanism proto.InternalMessageInfo

func (m *Mechanism) GetType() MechanismType {
	if m != nil {
		return m.Type
	}
	return MechanismType_DEFAULT_INTERFACE
}

func (m *Mechanism) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type Connection struct {
	Id                   string                               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NetworkService       string                               `protobuf:"bytes,2,opt,name=network_service,json=networkService,proto3" json:"network_service,omitempty"`
	Mechanism            *Mechanism                           `protobuf:"bytes,3,opt,name=mechanism,proto3" json:"mechanism,omitempty"`
	Context              *connectioncontext.ConnectionContext `protobuf:"bytes,4,opt,name=context,proto3" json:"context,omitempty"`
	Labels               map[string]string                    `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	State                State                                `protobuf:"varint,6,opt,name=state,proto3,enum=local.connection.State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}
func (*Connection) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{1}
}

func (m *Connection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connection.Unmarshal(m, b)
}
func (m *Connection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connection.Marshal(b, m, deterministic)
}
func (m *Connection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connection.Merge(m, src)
}
func (m *Connection) XXX_Size() int {
	return xxx_messageInfo_Connection.Size(m)
}
func (m *Connection) XXX_DiscardUnknown() {
	xxx_messageInfo_Connection.DiscardUnknown(m)
}

var xxx_messageInfo_Connection proto.InternalMessageInfo

func (m *Connection) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Connection) GetNetworkService() string {
	if m != nil {
		return m.NetworkService
	}
	return ""
}

func (m *Connection) GetMechanism() *Mechanism {
	if m != nil {
		return m.Mechanism
	}
	return nil
}

func (m *Connection) GetContext() *connectioncontext.ConnectionContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Connection) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Connection) GetState() State {
	if m != nil {
		return m.State
	}
	return State_UP
}

type ConnectionEvent struct {
	Type                 ConnectionEventType    `protobuf:"varint,1,opt,name=type,proto3,enum=local.connection.ConnectionEventType" json:"type,omitempty"`
	Connections          map[string]*Connection `protobuf:"bytes,2,rep,name=connections,proto3" json:"connections,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ConnectionEvent) Reset()         { *m = ConnectionEvent{} }
func (m *ConnectionEvent) String() string { return proto.CompactTextString(m) }
func (*ConnectionEvent) ProtoMessage()    {}
func (*ConnectionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{2}
}

func (m *ConnectionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionEvent.Unmarshal(m, b)
}
func (m *ConnectionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionEvent.Marshal(b, m, deterministic)
}
func (m *ConnectionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionEvent.Merge(m, src)
}
func (m *ConnectionEvent) XXX_Size() int {
	return xxx_messageInfo_ConnectionEvent.Size(m)
}
func (m *ConnectionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionEvent proto.InternalMessageInfo

func (m *ConnectionEvent) GetType() ConnectionEventType {
	if m != nil {
		return m.Type
	}
	return ConnectionEventType_INITIAL_STATE_TRANSFER
}

func (m *ConnectionEvent) GetConnections() map[string]*Connection {
	if m != nil {
		return m.Connections
	}
	return nil
}

func init() {
	proto.RegisterEnum("local.connection.MechanismType", MechanismType_name, MechanismType_value)
	proto.RegisterEnum("local.connection.State", State_name, State_value)
	proto.RegisterEnum("local.connection.ConnectionEventType", ConnectionEventType_name, ConnectionEventType_value)
	proto.RegisterType((*Mechanism)(nil), "local.connection.Mechanism")
	proto.RegisterMapType((map[string]string)(nil), "local.connection.Mechanism.ParametersEntry")
	proto.RegisterType((*Connection)(nil), "local.connection.Connection")
	proto.RegisterMapType((map[string]string)(nil), "local.connection.Connection.LabelsEntry")
	proto.RegisterType((*ConnectionEvent)(nil), "local.connection.ConnectionEvent")
	proto.RegisterMapType((map[string]*Connection)(nil), "local.connection.ConnectionEvent.ConnectionsEntry")
}

func init() { proto.RegisterFile("connection.proto", fileDescriptor_51baa40a1cc6b48b) }

var fileDescriptor_51baa40a1cc6b48b = []byte{
	// 623 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x6d, 0x4f, 0xd3, 0x50,
	0x18, 0xa5, 0xdd, 0x56, 0xe5, 0x19, 0xb0, 0x72, 0x41, 0xac, 0xd5, 0x44, 0x24, 0x1a, 0x17, 0x8c,
	0x9d, 0x29, 0x5f, 0xc4, 0x44, 0x63, 0x65, 0x97, 0xd0, 0xb0, 0x0d, 0x6c, 0x0b, 0x24, 0xc6, 0x64,
	0xe9, 0xca, 0x05, 0x1a, 0xda, 0xde, 0xa6, 0xbd, 0xa0, 0xfb, 0xe6, 0x77, 0xff, 0x9a, 0x3f, 0xc5,
	0x1f, 0x61, 0xfa, 0x02, 0xbd, 0x30, 0x32, 0xe2, 0x97, 0x65, 0x3d, 0xcf, 0x39, 0xcf, 0x73, 0xee,
	0xb9, 0x2f, 0x20, 0x7b, 0x34, 0x8a, 0x88, 0xc7, 0x7c, 0x1a, 0x69, 0x71, 0x42, 0x19, 0x45, 0x72,
	0x40, 0x3d, 0x37, 0xd0, 0x2a, 0x5c, 0x55, 0x62, 0x36, 0x8e, 0x49, 0xda, 0x21, 0x61, 0xcc, 0xc6,
	0xc5, 0x6f, 0xc1, 0x55, 0xcf, 0x4f, 0x7d, 0x76, 0x76, 0x31, 0xd2, 0x3c, 0x1a, 0x76, 0x22, 0xc2,
	0x7e, 0xd0, 0xe4, 0x3c, 0x25, 0xc9, 0xa5, 0xef, 0x91, 0x90, 0xa4, 0x67, 0x77, 0x41, 0x1e, 0x8d,
	0x58, 0x42, 0x83, 0x38, 0x70, 0x23, 0xd2, 0x71, 0x63, 0xbf, 0x53, 0xcd, 0xc9, 0x4a, 0xe4, 0x27,
	0x9b, 0x44, 0x8a, 0x61, 0x6b, 0x7f, 0x04, 0x98, 0xed, 0x13, 0xef, 0xcc, 0x8d, 0xfc, 0x34, 0x44,
	0x1b, 0x50, 0xcf, 0x5c, 0x29, 0xc2, 0xaa, 0xd0, 0x5e, 0xd0, 0x9f, 0x6b, 0xb7, 0x5d, 0x6b, 0xd7,
	0x54, 0x67, 0x1c, 0x13, 0x2b, 0x27, 0xa3, 0x5d, 0x80, 0xd8, 0x4d, 0xdc, 0x90, 0x30, 0x92, 0xa4,
	0x8a, 0xb8, 0x5a, 0x6b, 0x37, 0xf5, 0x37, 0x53, 0xa4, 0xda, 0xfe, 0x35, 0x1b, 0x47, 0x2c, 0x19,
	0x5b, 0x9c, 0x5c, 0xfd, 0x08, 0xad, 0x5b, 0x65, 0x24, 0x43, 0xed, 0x9c, 0x8c, 0x73, 0x4f, 0xb3,
	0x56, 0xf6, 0x17, 0x2d, 0x43, 0xe3, 0xd2, 0x0d, 0x2e, 0x88, 0x22, 0xe6, 0x58, 0xf1, 0xf1, 0x41,
	0x7c, 0x2f, 0xac, 0xfd, 0x15, 0x01, 0xb6, 0xae, 0x67, 0xa2, 0x05, 0x10, 0xfd, 0xe3, 0x52, 0x29,
	0xfa, 0xc7, 0xe8, 0x35, 0xb4, 0xca, 0xf8, 0x86, 0x65, 0x7e, 0x65, 0x8b, 0x85, 0x12, 0xb6, 0x0b,
	0x14, 0x6d, 0xc2, 0x6c, 0x78, 0xe5, 0x57, 0xa9, 0xad, 0x0a, 0xed, 0xa6, 0xfe, 0x74, 0xca, 0x92,
	0xac, 0x8a, 0x8d, 0x3e, 0xc1, 0x83, 0x32, 0x62, 0xa5, 0x9e, 0x0b, 0x5f, 0x6a, 0x93, 0xe1, 0x57,
	0x1e, 0xb7, 0x0a, 0xc4, 0xba, 0x12, 0xa1, 0xcf, 0x20, 0x05, 0xee, 0x88, 0x04, 0xa9, 0xd2, 0xc8,
	0xa3, 0x6c, 0x4f, 0xce, 0xad, 0xd4, 0x5a, 0x2f, 0xa7, 0x16, 0x39, 0x96, 0x3a, 0xf4, 0x16, 0x1a,
	0x29, 0x73, 0x19, 0x51, 0xa4, 0x7c, 0x1b, 0x1f, 0x4f, 0x36, 0xb0, 0xb3, 0xb2, 0x55, 0xb0, 0xd4,
	0x4d, 0x68, 0x72, 0x5d, 0xfe, 0x2b, 0xee, 0x5f, 0x22, 0xb4, 0x2a, 0x33, 0xf8, 0x92, 0x44, 0x0c,
	0x6d, 0xde, 0x38, 0x43, 0xaf, 0xa6, 0xb9, 0xcf, 0x05, 0xdc, 0x49, 0x72, 0xa0, 0x59, 0xf1, 0xae,
	0x8e, 0x92, 0x7e, 0x6f, 0x07, 0xee, 0xbb, 0x4c, 0x82, 0x6f, 0xa3, 0x7e, 0x07, 0xf9, 0x36, 0xe1,
	0x8e, 0x45, 0xea, 0xfc, 0x22, 0x9b, 0xfa, 0xb3, 0x69, 0x53, 0xb9, 0x08, 0xd6, 0x7f, 0x0b, 0x30,
	0x7f, 0xe3, 0x56, 0xa0, 0x47, 0xb0, 0xd8, 0xc5, 0xdb, 0xc6, 0x41, 0xcf, 0x19, 0x9a, 0x03, 0x07,
	0x5b, 0xdb, 0xc6, 0x16, 0x96, 0x67, 0xd0, 0x32, 0xc8, 0xbb, 0xd8, 0x1a, 0xe0, 0x1e, 0x87, 0x0a,
	0x68, 0x09, 0x5a, 0x87, 0x3b, 0x7b, 0x36, 0x4f, 0x15, 0xd1, 0x22, 0xcc, 0xf7, 0x71, 0x9f, 0x83,
	0x6a, 0x19, 0xcf, 0xb6, 0xcc, 0xbd, 0x43, 0x0e, 0xac, 0x23, 0x19, 0xe6, 0x76, 0x8e, 0x38, 0xa4,
	0xb1, 0xfe, 0x04, 0x1a, 0xf9, 0xde, 0x22, 0x09, 0xc4, 0x83, 0x7d, 0x79, 0x06, 0x3d, 0x84, 0x7a,
	0x77, 0xef, 0x68, 0x20, 0x0b, 0xeb, 0x26, 0x2c, 0xdd, 0x91, 0x3c, 0x52, 0x61, 0xc5, 0x1c, 0x98,
	0x8e, 0x69, 0xf4, 0x86, 0xb6, 0x63, 0x38, 0x78, 0xe8, 0x58, 0xc6, 0xc0, 0xde, 0xc6, 0x96, 0x3c,
	0x83, 0x00, 0xa4, 0x83, 0xfd, 0xae, 0xe1, 0x64, 0x46, 0x01, 0xa4, 0x2e, 0xee, 0x61, 0x07, 0xcb,
	0xa2, 0x7e, 0x02, 0x8b, 0x7d, 0x1a, 0xf9, 0x8c, 0x26, 0xdc, 0x5d, 0xfb, 0x0a, 0x68, 0x02, 0x4c,
	0xd1, 0x8a, 0x76, 0x4a, 0xe9, 0x69, 0x40, 0x8a, 0xe7, 0x66, 0x74, 0x71, 0xa2, 0xe1, 0xec, 0xa9,
	0x53, 0x5f, 0xdc, 0xbb, 0xab, 0xef, 0x84, 0x2f, 0x73, 0xdf, 0xa0, 0xaa, 0x8f, 0xa4, 0xbc, 0xc5,
	0xc6, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x50, 0x8e, 0x4b, 0x5e, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MonitorConnectionClient is the client API for MonitorConnection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitorConnectionClient interface {
	MonitorConnections(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (MonitorConnection_MonitorConnectionsClient, error)
}

type monitorConnectionClient struct {
	cc *grpc.ClientConn
}

func NewMonitorConnectionClient(cc *grpc.ClientConn) MonitorConnectionClient {
	return &monitorConnectionClient{cc}
}

func (c *monitorConnectionClient) MonitorConnections(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (MonitorConnection_MonitorConnectionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MonitorConnection_serviceDesc.Streams[0], "/local.connection.MonitorConnection/MonitorConnections", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitorConnectionMonitorConnectionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MonitorConnection_MonitorConnectionsClient interface {
	Recv() (*ConnectionEvent, error)
	grpc.ClientStream
}

type monitorConnectionMonitorConnectionsClient struct {
	grpc.ClientStream
}

func (x *monitorConnectionMonitorConnectionsClient) Recv() (*ConnectionEvent, error) {
	m := new(ConnectionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MonitorConnectionServer is the server API for MonitorConnection service.
type MonitorConnectionServer interface {
	MonitorConnections(*empty.Empty, MonitorConnection_MonitorConnectionsServer) error
}

// UnimplementedMonitorConnectionServer can be embedded to have forward compatible implementations.
type UnimplementedMonitorConnectionServer struct {
}

func (*UnimplementedMonitorConnectionServer) MonitorConnections(req *empty.Empty, srv MonitorConnection_MonitorConnectionsServer) error {
	return status.Errorf(codes.Unimplemented, "method MonitorConnections not implemented")
}

func RegisterMonitorConnectionServer(s *grpc.Server, srv MonitorConnectionServer) {
	s.RegisterService(&_MonitorConnection_serviceDesc, srv)
}

func _MonitorConnection_MonitorConnections_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MonitorConnectionServer).MonitorConnections(m, &monitorConnectionMonitorConnectionsServer{stream})
}

type MonitorConnection_MonitorConnectionsServer interface {
	Send(*ConnectionEvent) error
	grpc.ServerStream
}

type monitorConnectionMonitorConnectionsServer struct {
	grpc.ServerStream
}

func (x *monitorConnectionMonitorConnectionsServer) Send(m *ConnectionEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _MonitorConnection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "local.connection.MonitorConnection",
	HandlerType: (*MonitorConnectionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorConnections",
			Handler:       _MonitorConnection_MonitorConnections_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "connection.proto",
}
