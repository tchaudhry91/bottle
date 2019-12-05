// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crate.proto

package pb

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

type Bottle struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              []byte   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bottle) Reset()         { *m = Bottle{} }
func (m *Bottle) String() string { return proto.CompactTextString(m) }
func (*Bottle) ProtoMessage()    {}
func (*Bottle) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0391c64e2e3c0e7, []int{0}
}

func (m *Bottle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bottle.Unmarshal(m, b)
}
func (m *Bottle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bottle.Marshal(b, m, deterministic)
}
func (m *Bottle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bottle.Merge(m, src)
}
func (m *Bottle) XXX_Size() int {
	return xxx_messageInfo_Bottle.Size(m)
}
func (m *Bottle) XXX_DiscardUnknown() {
	xxx_messageInfo_Bottle.DiscardUnknown(m)
}

var xxx_messageInfo_Bottle proto.InternalMessageInfo

func (m *Bottle) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Bottle) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type StoreResponse struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreResponse) Reset()         { *m = StoreResponse{} }
func (m *StoreResponse) String() string { return proto.CompactTextString(m) }
func (*StoreResponse) ProtoMessage()    {}
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0391c64e2e3c0e7, []int{1}
}

func (m *StoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreResponse.Unmarshal(m, b)
}
func (m *StoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreResponse.Marshal(b, m, deterministic)
}
func (m *StoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreResponse.Merge(m, src)
}
func (m *StoreResponse) XXX_Size() int {
	return xxx_messageInfo_StoreResponse.Size(m)
}
func (m *StoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoreResponse proto.InternalMessageInfo

func (m *StoreResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0391c64e2e3c0e7, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetResponse struct {
	Bottle               *Bottle  `protobuf:"bytes,1,opt,name=bottle,proto3" json:"bottle,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0391c64e2e3c0e7, []int{3}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetBottle() *Bottle {
	if m != nil {
		return m.Bottle
	}
	return nil
}

func (m *GetResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*Bottle)(nil), "pb.Bottle")
	proto.RegisterType((*StoreResponse)(nil), "pb.StoreResponse")
	proto.RegisterType((*GetRequest)(nil), "pb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "pb.GetResponse")
}

func init() { proto.RegisterFile("crate.proto", fileDescriptor_f0391c64e2e3c0e7) }

var fileDescriptor_f0391c64e2e3c0e7 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xbd, 0x4e, 0x86, 0x30,
	0x14, 0xfd, 0xa8, 0x82, 0xf1, 0xa2, 0x5f, 0xf4, 0x4e, 0x84, 0x38, 0x60, 0xa3, 0x09, 0x13, 0x03,
	0xbe, 0x81, 0x98, 0xb8, 0x9a, 0xfa, 0x04, 0x14, 0xee, 0x40, 0x62, 0xda, 0x5a, 0x2e, 0xab, 0xcf,
	0x6e, 0x5a, 0x41, 0xd1, 0xc1, 0xed, 0xb4, 0xa7, 0xe7, 0xaf, 0x90, 0x0f, 0xbe, 0x67, 0x6a, 0x9c,
	0xb7, 0x6c, 0x51, 0x38, 0x2d, 0x5b, 0xc8, 0x1e, 0x2d, 0xf3, 0x1b, 0xe1, 0x11, 0xc4, 0x34, 0x16,
	0x49, 0x95, 0xd4, 0xe7, 0x4a, 0x4c, 0x23, 0x16, 0x70, 0x36, 0x58, 0xc3, 0x64, 0xb8, 0x10, 0x55,
	0x52, 0x5f, 0xa8, 0xed, 0x28, 0x6f, 0xe1, 0xf2, 0x95, 0xad, 0x27, 0x45, 0xb3, 0xb3, 0x66, 0x26,
	0xbc, 0x82, 0x13, 0xf2, 0x7e, 0xd5, 0x06, 0x28, 0x6f, 0x00, 0x9e, 0x89, 0x15, 0xbd, 0x2f, 0x34,
	0xf3, 0x5f, 0x6b, 0xd9, 0x41, 0x1e, 0xd9, 0x55, 0x2e, 0x21, 0xd3, 0xb1, 0x43, 0x7c, 0x92, 0xb7,
	0xd0, 0x38, 0xdd, 0x7c, 0xb5, 0x52, 0x2b, 0xb3, 0x45, 0x88, 0xef, 0x88, 0xf6, 0x03, 0xd2, 0x2e,
	0x8c, 0xc1, 0x3b, 0x38, 0x7d, 0xb1, 0x8b, 0xc7, 0x63, 0x90, 0xfd, 0xa4, 0x96, 0x3b, 0x1b, 0x79,
	0xc0, 0x1a, 0xd2, 0x58, 0x1a, 0x77, 0xd7, 0xe5, 0x75, 0xc0, 0xbf, 0xb6, 0xc8, 0x03, 0xde, 0x43,
	0xfa, 0xe4, 0xfb, 0xc9, 0xfc, 0x6f, 0xa8, 0xb3, 0xf8, 0x89, 0x0f, 0x9f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x88, 0xc7, 0x9c, 0x8a, 0x53, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CrateClient is the client API for Crate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CrateClient interface {
	Pour(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Bottle, error)
	Store(ctx context.Context, in *Bottle, opts ...grpc.CallOption) (*StoreResponse, error)
	Drain(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Bottle, error)
}

type crateClient struct {
	cc *grpc.ClientConn
}

func NewCrateClient(cc *grpc.ClientConn) CrateClient {
	return &crateClient{cc}
}

func (c *crateClient) Pour(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Bottle, error) {
	out := new(Bottle)
	err := c.cc.Invoke(ctx, "/pb.Crate/Pour", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crateClient) Store(ctx context.Context, in *Bottle, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/pb.Crate/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crateClient) Drain(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Bottle, error) {
	out := new(Bottle)
	err := c.cc.Invoke(ctx, "/pb.Crate/Drain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrateServer is the server API for Crate service.
type CrateServer interface {
	Pour(context.Context, *GetRequest) (*Bottle, error)
	Store(context.Context, *Bottle) (*StoreResponse, error)
	Drain(context.Context, *GetRequest) (*Bottle, error)
}

// UnimplementedCrateServer can be embedded to have forward compatible implementations.
type UnimplementedCrateServer struct {
}

func (*UnimplementedCrateServer) Pour(ctx context.Context, req *GetRequest) (*Bottle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pour not implemented")
}
func (*UnimplementedCrateServer) Store(ctx context.Context, req *Bottle) (*StoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (*UnimplementedCrateServer) Drain(ctx context.Context, req *GetRequest) (*Bottle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Drain not implemented")
}

func RegisterCrateServer(s *grpc.Server, srv CrateServer) {
	s.RegisterService(&_Crate_serviceDesc, srv)
}

func _Crate_Pour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrateServer).Pour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Crate/Pour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrateServer).Pour(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crate_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bottle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrateServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Crate/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrateServer).Store(ctx, req.(*Bottle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crate_Drain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrateServer).Drain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Crate/Drain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrateServer).Drain(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Crate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Crate",
	HandlerType: (*CrateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pour",
			Handler:    _Crate_Pour_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _Crate_Store_Handler,
		},
		{
			MethodName: "Drain",
			Handler:    _Crate_Drain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crate.proto",
}