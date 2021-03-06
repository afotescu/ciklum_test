// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ports.proto

package ports

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

type Port struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	City                 string    `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Country              string    `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Alias                []string  `protobuf:"bytes,4,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions              []string  `protobuf:"bytes,5,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates          []float32 `protobuf:"fixed32,6,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province             string    `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty"`
	Timezone             string    `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs               []string  `protobuf:"bytes,9,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code                 string    `protobuf:"bytes,10,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_ports_032566dcea0c276d, []int{0}
}
func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (dst *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(dst, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Port) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Port) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Port) GetAlias() []string {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *Port) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *Port) GetCoordinates() []float32 {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func (m *Port) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *Port) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *Port) GetUnlocs() []string {
	if m != nil {
		return m.Unlocs
	}
	return nil
}

func (m *Port) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type Ports struct {
	Data                 []*Port  `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ports) Reset()         { *m = Ports{} }
func (m *Ports) String() string { return proto.CompactTextString(m) }
func (*Ports) ProtoMessage()    {}
func (*Ports) Descriptor() ([]byte, []int) {
	return fileDescriptor_ports_032566dcea0c276d, []int{1}
}
func (m *Ports) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ports.Unmarshal(m, b)
}
func (m *Ports) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ports.Marshal(b, m, deterministic)
}
func (dst *Ports) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ports.Merge(dst, src)
}
func (m *Ports) XXX_Size() int {
	return xxx_messageInfo_Ports.Size(m)
}
func (m *Ports) XXX_DiscardUnknown() {
	xxx_messageInfo_Ports.DiscardUnknown(m)
}

var xxx_messageInfo_Ports proto.InternalMessageInfo

func (m *Ports) GetData() []*Port {
	if m != nil {
		return m.Data
	}
	return nil
}

type PortsPage struct {
	Limit                uint64   `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               uint64   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortsPage) Reset()         { *m = PortsPage{} }
func (m *PortsPage) String() string { return proto.CompactTextString(m) }
func (*PortsPage) ProtoMessage()    {}
func (*PortsPage) Descriptor() ([]byte, []int) {
	return fileDescriptor_ports_032566dcea0c276d, []int{2}
}
func (m *PortsPage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortsPage.Unmarshal(m, b)
}
func (m *PortsPage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortsPage.Marshal(b, m, deterministic)
}
func (dst *PortsPage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortsPage.Merge(dst, src)
}
func (m *PortsPage) XXX_Size() int {
	return xxx_messageInfo_PortsPage.Size(m)
}
func (m *PortsPage) XXX_DiscardUnknown() {
	xxx_messageInfo_PortsPage.DiscardUnknown(m)
}

var xxx_messageInfo_PortsPage proto.InternalMessageInfo

func (m *PortsPage) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *PortsPage) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type Response struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ports_032566dcea0c276d, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Port)(nil), "ports.Port")
	proto.RegisterType((*Ports)(nil), "ports.Ports")
	proto.RegisterType((*PortsPage)(nil), "ports.PortsPage")
	proto.RegisterType((*Response)(nil), "ports.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransporterClient is the client API for Transporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransporterClient interface {
	CreateOrUpdatePorts(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Response, error)
	GetPorts(ctx context.Context, in *PortsPage, opts ...grpc.CallOption) (*Ports, error)
}

type transporterClient struct {
	cc *grpc.ClientConn
}

func NewTransporterClient(cc *grpc.ClientConn) TransporterClient {
	return &transporterClient{cc}
}

func (c *transporterClient) CreateOrUpdatePorts(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ports.Transporter/CreateOrUpdatePorts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transporterClient) GetPorts(ctx context.Context, in *PortsPage, opts ...grpc.CallOption) (*Ports, error) {
	out := new(Ports)
	err := c.cc.Invoke(ctx, "/ports.Transporter/GetPorts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransporterServer is the server API for Transporter service.
type TransporterServer interface {
	CreateOrUpdatePorts(context.Context, *Port) (*Response, error)
	GetPorts(context.Context, *PortsPage) (*Ports, error)
}

func RegisterTransporterServer(s *grpc.Server, srv TransporterServer) {
	s.RegisterService(&_Transporter_serviceDesc, srv)
}

func _Transporter_CreateOrUpdatePorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Port)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransporterServer).CreateOrUpdatePorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ports.Transporter/CreateOrUpdatePorts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransporterServer).CreateOrUpdatePorts(ctx, req.(*Port))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transporter_GetPorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortsPage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransporterServer).GetPorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ports.Transporter/GetPorts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransporterServer).GetPorts(ctx, req.(*PortsPage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ports.Transporter",
	HandlerType: (*TransporterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdatePorts",
			Handler:    _Transporter_CreateOrUpdatePorts_Handler,
		},
		{
			MethodName: "GetPorts",
			Handler:    _Transporter_GetPorts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ports.proto",
}

func init() { proto.RegisterFile("ports.proto", fileDescriptor_ports_032566dcea0c276d) }

var fileDescriptor_ports_032566dcea0c276d = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0x41, 0x6b, 0xe3, 0x30,
	0x10, 0x85, 0xe3, 0xd8, 0x4e, 0xec, 0xf1, 0xc2, 0x2e, 0xda, 0x52, 0x44, 0x2e, 0x35, 0x3e, 0xf9,
	0x50, 0x72, 0x48, 0xe9, 0xa1, 0xe7, 0x1e, 0x7a, 0x6c, 0x30, 0xed, 0x0f, 0x50, 0xed, 0x49, 0x10,
	0x24, 0x1a, 0x23, 0x4d, 0x0a, 0xe9, 0x8f, 0x2f, 0x45, 0xb2, 0x5d, 0x7c, 0xd3, 0xf7, 0x66, 0xf0,
	0xbc, 0xf7, 0x30, 0x14, 0x3d, 0x59, 0x76, 0xdb, 0xde, 0x12, 0x93, 0x48, 0x03, 0x54, 0xdf, 0x11,
	0x24, 0x7b, 0xb2, 0x2c, 0x04, 0x24, 0x46, 0x9d, 0x51, 0x46, 0x65, 0x54, 0xe7, 0x4d, 0x78, 0x7b,
	0xad, 0xd5, 0x7c, 0x95, 0xcb, 0x41, 0xf3, 0x6f, 0x21, 0x61, 0xdd, 0xd2, 0xc5, 0xb0, 0xbd, 0xca,
	0x38, 0xc8, 0x13, 0x8a, 0x1b, 0x48, 0xd5, 0x49, 0x2b, 0x27, 0x93, 0x32, 0xae, 0xf3, 0x66, 0x00,
	0xbf, 0x6f, 0xf1, 0xa8, 0xc9, 0x38, 0x99, 0x06, 0x7d, 0x42, 0x51, 0x42, 0xd1, 0x12, 0xd9, 0x4e,
	0x1b, 0xc5, 0xe8, 0xe4, 0xaa, 0x8c, 0xeb, 0x65, 0x33, 0x97, 0xc4, 0x06, 0xb2, 0xde, 0xd2, 0xa7,
	0x36, 0x2d, 0xca, 0x75, 0x38, 0xf6, 0xcb, 0x7e, 0xc6, 0xfa, 0x8c, 0x5f, 0x64, 0x50, 0x66, 0xc3,
	0x6c, 0x62, 0x71, 0x0b, 0xab, 0x8b, 0x39, 0x51, 0xeb, 0x64, 0x1e, 0x4e, 0x8e, 0x14, 0xf2, 0x50,
	0x87, 0x12, 0xc6, 0x3c, 0xd4, 0x61, 0x55, 0x43, 0xea, 0xf3, 0x3b, 0x71, 0x07, 0x49, 0xa7, 0x58,
	0xc9, 0xa8, 0x8c, 0xeb, 0x62, 0x57, 0x6c, 0x87, 0xb2, 0xfc, 0xac, 0x09, 0x83, 0xea, 0x09, 0xf2,
	0xb0, 0xb9, 0x57, 0x47, 0xf4, 0x61, 0x4f, 0xfa, 0xac, 0x39, 0xf4, 0x95, 0x34, 0x03, 0xf8, 0xc3,
	0x74, 0x38, 0x38, 0xe4, 0x50, 0x59, 0xd2, 0x8c, 0x54, 0x55, 0x90, 0x35, 0xe8, 0x7a, 0x32, 0x2e,
	0x98, 0x73, 0xac, 0xf8, 0xe2, 0xc6, 0xaa, 0x47, 0xda, 0x59, 0x28, 0xde, 0xac, 0x32, 0xce, 0xdf,
	0x45, 0x2b, 0x1e, 0xe1, 0xff, 0xb3, 0x45, 0xc5, 0xf8, 0x6a, 0xdf, 0xfb, 0x4e, 0x31, 0x0e, 0x2e,
	0xe7, 0xbe, 0x36, 0x7f, 0x47, 0x98, 0xbe, 0x5d, 0x2d, 0xc4, 0x3d, 0x64, 0x2f, 0xc8, 0xc3, 0xee,
	0xbf, 0xd9, 0x6e, 0x70, 0xbd, 0xf9, 0x33, 0x57, 0xaa, 0xc5, 0xc7, 0x2a, 0xfc, 0x0b, 0x0f, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x2e, 0x8c, 0x7a, 0x1a, 0x02, 0x00, 0x00,
}
