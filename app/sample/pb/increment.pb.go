// Code generated by protoc-gen-go. DO NOT EDIT.
// source: increment.proto

package increment

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

type IncrementRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncrementRequest) Reset()         { *m = IncrementRequest{} }
func (m *IncrementRequest) String() string { return proto.CompactTextString(m) }
func (*IncrementRequest) ProtoMessage()    {}
func (*IncrementRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fe32297f269ab62, []int{0}
}

func (m *IncrementRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncrementRequest.Unmarshal(m, b)
}
func (m *IncrementRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncrementRequest.Marshal(b, m, deterministic)
}
func (m *IncrementRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncrementRequest.Merge(m, src)
}
func (m *IncrementRequest) XXX_Size() int {
	return xxx_messageInfo_IncrementRequest.Size(m)
}
func (m *IncrementRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IncrementRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IncrementRequest proto.InternalMessageInfo

func (m *IncrementRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type IncrementResponse struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncrementResponse) Reset()         { *m = IncrementResponse{} }
func (m *IncrementResponse) String() string { return proto.CompactTextString(m) }
func (*IncrementResponse) ProtoMessage()    {}
func (*IncrementResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fe32297f269ab62, []int{1}
}

func (m *IncrementResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncrementResponse.Unmarshal(m, b)
}
func (m *IncrementResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncrementResponse.Marshal(b, m, deterministic)
}
func (m *IncrementResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncrementResponse.Merge(m, src)
}
func (m *IncrementResponse) XXX_Size() int {
	return xxx_messageInfo_IncrementResponse.Size(m)
}
func (m *IncrementResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IncrementResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IncrementResponse proto.InternalMessageInfo

func (m *IncrementResponse) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*IncrementRequest)(nil), "increment.IncrementRequest")
	proto.RegisterType((*IncrementResponse)(nil), "increment.IncrementResponse")
}

func init() { proto.RegisterFile("increment.proto", fileDescriptor_7fe32297f269ab62) }

var fileDescriptor_7fe32297f269ab62 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0xcc, 0x4b, 0x2e,
	0x4a, 0xcd, 0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x69, 0x71, 0x09, 0x78, 0xc2, 0x38, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x62, 0x5c,
	0x6c, 0x79, 0xa5, 0xb9, 0x49, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x50, 0x9e,
	0x92, 0x36, 0x97, 0x20, 0x92, 0xda, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x5c, 0x8a, 0x8d, 0xa2,
	0x90, 0x0c, 0x0e, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x72, 0xe3, 0xe2, 0x84, 0x8b, 0x09,
	0x49, 0xeb, 0x21, 0x9c, 0x85, 0xee, 0x04, 0x29, 0x19, 0xec, 0x92, 0x10, 0x3b, 0x93, 0xd8, 0xc0,
	0xde, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x30, 0xc8, 0xec, 0xd9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IncrementServiceClient is the client API for IncrementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IncrementServiceClient interface {
	Increment(ctx context.Context, in *IncrementRequest, opts ...grpc.CallOption) (*IncrementResponse, error)
}

type incrementServiceClient struct {
	cc *grpc.ClientConn
}

func NewIncrementServiceClient(cc *grpc.ClientConn) IncrementServiceClient {
	return &incrementServiceClient{cc}
}

func (c *incrementServiceClient) Increment(ctx context.Context, in *IncrementRequest, opts ...grpc.CallOption) (*IncrementResponse, error) {
	out := new(IncrementResponse)
	err := c.cc.Invoke(ctx, "/increment.IncrementService/Increment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IncrementServiceServer is the server API for IncrementService service.
type IncrementServiceServer interface {
	Increment(context.Context, *IncrementRequest) (*IncrementResponse, error)
}

// UnimplementedIncrementServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIncrementServiceServer struct {
}

func (*UnimplementedIncrementServiceServer) Increment(ctx context.Context, req *IncrementRequest) (*IncrementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increment not implemented")
}

func RegisterIncrementServiceServer(s *grpc.Server, srv IncrementServiceServer) {
	s.RegisterService(&_IncrementService_serviceDesc, srv)
}

func _IncrementService_Increment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementServiceServer).Increment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/increment.IncrementService/Increment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementServiceServer).Increment(ctx, req.(*IncrementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IncrementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "increment.IncrementService",
	HandlerType: (*IncrementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Increment",
			Handler:    _IncrementService_Increment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "increment.proto",
}
