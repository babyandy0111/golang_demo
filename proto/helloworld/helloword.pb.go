// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloword.proto

package helloworld

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

type AndyTestRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AndyTestRequest) Reset()         { *m = AndyTestRequest{} }
func (m *AndyTestRequest) String() string { return proto.CompactTextString(m) }
func (*AndyTestRequest) ProtoMessage()    {}
func (*AndyTestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_408dc8e96988f433, []int{0}
}

func (m *AndyTestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AndyTestRequest.Unmarshal(m, b)
}
func (m *AndyTestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AndyTestRequest.Marshal(b, m, deterministic)
}
func (m *AndyTestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AndyTestRequest.Merge(m, src)
}
func (m *AndyTestRequest) XXX_Size() int {
	return xxx_messageInfo_AndyTestRequest.Size(m)
}
func (m *AndyTestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AndyTestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AndyTestRequest proto.InternalMessageInfo

func (m *AndyTestRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AndyTestReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AndyTestReply) Reset()         { *m = AndyTestReply{} }
func (m *AndyTestReply) String() string { return proto.CompactTextString(m) }
func (*AndyTestReply) ProtoMessage()    {}
func (*AndyTestReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_408dc8e96988f433, []int{1}
}

func (m *AndyTestReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AndyTestReply.Unmarshal(m, b)
}
func (m *AndyTestReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AndyTestReply.Marshal(b, m, deterministic)
}
func (m *AndyTestReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AndyTestReply.Merge(m, src)
}
func (m *AndyTestReply) XXX_Size() int {
	return xxx_messageInfo_AndyTestReply.Size(m)
}
func (m *AndyTestReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AndyTestReply.DiscardUnknown(m)
}

var xxx_messageInfo_AndyTestReply proto.InternalMessageInfo

func (m *AndyTestReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*AndyTestRequest)(nil), "helloworld.AndyTestRequest")
	proto.RegisterType((*AndyTestReply)(nil), "helloworld.AndyTestReply")
}

func init() { proto.RegisterFile("helloword.proto", fileDescriptor_408dc8e96988f433) }

var fileDescriptor_408dc8e96988f433 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0x09, 0xe4,
	0xa4, 0x28, 0xa9, 0x72, 0xf1, 0x3b, 0xe6, 0xa5, 0x54, 0x86, 0xa4, 0x16, 0x97, 0x04, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x9a, 0x5c, 0xbc, 0x08, 0x65, 0x05, 0x39, 0x95, 0x42, 0x12,
	0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0x30, 0x75, 0x30, 0xae, 0x51, 0x18, 0x17, 0x17,
	0x48, 0x69, 0x70, 0x6a, 0x51, 0x59, 0x6a, 0x91, 0x90, 0x07, 0x17, 0x4f, 0x70, 0x62, 0x65, 0x70,
	0x7e, 0x6e, 0x6a, 0x48, 0x46, 0x66, 0x5e, 0xba, 0x90, 0xb4, 0x1e, 0xc2, 0x72, 0x3d, 0x34, 0x9b,
	0xa5, 0x24, 0xb1, 0x4b, 0x16, 0xe4, 0x54, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0x1d, 0x6f, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x62, 0xe9, 0x0b, 0x23, 0xcf, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AndyServerClient is the client API for AndyServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AndyServerClient interface {
	SaySomeThing(ctx context.Context, in *AndyTestRequest, opts ...grpc.CallOption) (*AndyTestReply, error)
}

type andyServerClient struct {
	cc *grpc.ClientConn
}

func NewAndyServerClient(cc *grpc.ClientConn) AndyServerClient {
	return &andyServerClient{cc}
}

func (c *andyServerClient) SaySomeThing(ctx context.Context, in *AndyTestRequest, opts ...grpc.CallOption) (*AndyTestReply, error) {
	out := new(AndyTestReply)
	err := c.cc.Invoke(ctx, "/helloworld.AndyServer/SaySomeThing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AndyServerServer is the server API for AndyServer service.
type AndyServerServer interface {
	SaySomeThing(context.Context, *AndyTestRequest) (*AndyTestReply, error)
}

// UnimplementedAndyServerServer can be embedded to have forward compatible implementations.
type UnimplementedAndyServerServer struct {
}

func (*UnimplementedAndyServerServer) SaySomeThing(ctx context.Context, req *AndyTestRequest) (*AndyTestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaySomeThing not implemented")
}

func RegisterAndyServerServer(s *grpc.Server, srv AndyServerServer) {
	s.RegisterService(&_AndyServer_serviceDesc, srv)
}

func _AndyServer_SaySomeThing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AndyTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AndyServerServer).SaySomeThing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.AndyServer/SaySomeThing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AndyServerServer).SaySomeThing(ctx, req.(*AndyTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AndyServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.AndyServer",
	HandlerType: (*AndyServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaySomeThing",
			Handler:    _AndyServer_SaySomeThing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloword.proto",
}
