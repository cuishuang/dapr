// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dapr/proto/daprinternal/v1/daprinternal.proto

package v1

import (
	context "context"
	fmt "fmt"
	v1 "github.com/dapr/dapr/pkg/proto/common/v1"
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

// Actor represents actor using actor_type and actor_id
type Actor struct {
	// actor_type is the type of actor.
	//
	// This field is required.
	ActorType string `protobuf:"bytes,1,opt,name=actor_type,json=actorType,proto3" json:"actor_type,omitempty"`
	// actor_id is the id of actor type (actor_type)
	//
	// This field is required.
	ActorId              string   `protobuf:"bytes,2,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Actor) Reset()         { *m = Actor{} }
func (m *Actor) String() string { return proto.CompactTextString(m) }
func (*Actor) ProtoMessage()    {}
func (*Actor) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c6da3b6bd4beea4, []int{0}
}

func (m *Actor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Actor.Unmarshal(m, b)
}
func (m *Actor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Actor.Marshal(b, m, deterministic)
}
func (m *Actor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Actor.Merge(m, src)
}
func (m *Actor) XXX_Size() int {
	return xxx_messageInfo_Actor.Size(m)
}
func (m *Actor) XXX_DiscardUnknown() {
	xxx_messageInfo_Actor.DiscardUnknown(m)
}

var xxx_messageInfo_Actor proto.InternalMessageInfo

func (m *Actor) GetActorType() string {
	if m != nil {
		return m.ActorType
	}
	return ""
}

func (m *Actor) GetActorId() string {
	if m != nil {
		return m.ActorId
	}
	return ""
}

// InternalInvokeRequest is the message to transfer caller's data to callee
// for service invocaton. This includes callee's app id and caller's request data.
type InternalInvokeRequest struct {
	// ver is the version of Dapr API.
	//
	// This field is required.
	Ver APIVersion `protobuf:"varint,1,opt,name=ver,proto3,enum=dapr.proto.daprinternal.v1.APIVersion" json:"ver,omitempty"`
	// metadata holds caller's HTTP headers or gRPC metadata.
	//
	// This field is required.
	Metadata map[string]*ListStringValue `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// message includes caller's invocation request message.
	//
	// This field is required.
	Message *v1.InvokeRequest `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// actor includes actor type and id. This field is used only for
	// actor service invocation.
	//
	// This field is optional.
	Actor                *Actor   `protobuf:"bytes,4,opt,name=actor,proto3" json:"actor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalInvokeRequest) Reset()         { *m = InternalInvokeRequest{} }
func (m *InternalInvokeRequest) String() string { return proto.CompactTextString(m) }
func (*InternalInvokeRequest) ProtoMessage()    {}
func (*InternalInvokeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c6da3b6bd4beea4, []int{1}
}

func (m *InternalInvokeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalInvokeRequest.Unmarshal(m, b)
}
func (m *InternalInvokeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalInvokeRequest.Marshal(b, m, deterministic)
}
func (m *InternalInvokeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalInvokeRequest.Merge(m, src)
}
func (m *InternalInvokeRequest) XXX_Size() int {
	return xxx_messageInfo_InternalInvokeRequest.Size(m)
}
func (m *InternalInvokeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalInvokeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InternalInvokeRequest proto.InternalMessageInfo

func (m *InternalInvokeRequest) GetVer() APIVersion {
	if m != nil {
		return m.Ver
	}
	return APIVersion_UNKNOWN
}

func (m *InternalInvokeRequest) GetMetadata() map[string]*ListStringValue {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *InternalInvokeRequest) GetMessage() *v1.InvokeRequest {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *InternalInvokeRequest) GetActor() *Actor {
	if m != nil {
		return m.Actor
	}
	return nil
}

// InternalInvokeResponse is the message to transfer callee's response to caller
// for service invocaton.
type InternalInvokeResponse struct {
	// status is http/grpc status.
	//
	// This field is required.
	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// headers holds app channel response headers.
	//
	// This field is required.
	Headers map[string]*ListStringValue `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// trailer holds app channel response trailers.
	// this will be used only for gRPC app channel
	//
	// This field is optional.
	Trailers map[string]*ListStringValue `protobuf:"bytes,3,rep,name=trailers,proto3" json:"trailers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// message includes callee's invocation response message.
	//
	// This field is required.
	Message              *v1.InvokeResponse `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *InternalInvokeResponse) Reset()         { *m = InternalInvokeResponse{} }
func (m *InternalInvokeResponse) String() string { return proto.CompactTextString(m) }
func (*InternalInvokeResponse) ProtoMessage()    {}
func (*InternalInvokeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c6da3b6bd4beea4, []int{2}
}

func (m *InternalInvokeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalInvokeResponse.Unmarshal(m, b)
}
func (m *InternalInvokeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalInvokeResponse.Marshal(b, m, deterministic)
}
func (m *InternalInvokeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalInvokeResponse.Merge(m, src)
}
func (m *InternalInvokeResponse) XXX_Size() int {
	return xxx_messageInfo_InternalInvokeResponse.Size(m)
}
func (m *InternalInvokeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalInvokeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InternalInvokeResponse proto.InternalMessageInfo

func (m *InternalInvokeResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *InternalInvokeResponse) GetHeaders() map[string]*ListStringValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *InternalInvokeResponse) GetTrailers() map[string]*ListStringValue {
	if m != nil {
		return m.Trailers
	}
	return nil
}

func (m *InternalInvokeResponse) GetMessage() *v1.InvokeResponse {
	if m != nil {
		return m.Message
	}
	return nil
}

// ListStringValue represents string value array
type ListStringValue struct {
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListStringValue) Reset()         { *m = ListStringValue{} }
func (m *ListStringValue) String() string { return proto.CompactTextString(m) }
func (*ListStringValue) ProtoMessage()    {}
func (*ListStringValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c6da3b6bd4beea4, []int{3}
}

func (m *ListStringValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStringValue.Unmarshal(m, b)
}
func (m *ListStringValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStringValue.Marshal(b, m, deterministic)
}
func (m *ListStringValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStringValue.Merge(m, src)
}
func (m *ListStringValue) XXX_Size() int {
	return xxx_messageInfo_ListStringValue.Size(m)
}
func (m *ListStringValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStringValue.DiscardUnknown(m)
}

var xxx_messageInfo_ListStringValue proto.InternalMessageInfo

func (m *ListStringValue) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Actor)(nil), "dapr.proto.daprinternal.v1.Actor")
	proto.RegisterType((*InternalInvokeRequest)(nil), "dapr.proto.daprinternal.v1.InternalInvokeRequest")
	proto.RegisterMapType((map[string]*ListStringValue)(nil), "dapr.proto.daprinternal.v1.InternalInvokeRequest.MetadataEntry")
	proto.RegisterType((*InternalInvokeResponse)(nil), "dapr.proto.daprinternal.v1.InternalInvokeResponse")
	proto.RegisterMapType((map[string]*ListStringValue)(nil), "dapr.proto.daprinternal.v1.InternalInvokeResponse.HeadersEntry")
	proto.RegisterMapType((map[string]*ListStringValue)(nil), "dapr.proto.daprinternal.v1.InternalInvokeResponse.TrailersEntry")
	proto.RegisterType((*ListStringValue)(nil), "dapr.proto.daprinternal.v1.ListStringValue")
}

func init() {
	proto.RegisterFile("dapr/proto/daprinternal/v1/daprinternal.proto", fileDescriptor_3c6da3b6bd4beea4)
}

var fileDescriptor_3c6da3b6bd4beea4 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xe9, 0x6a, 0xdb, 0x40,
	0x10, 0xc7, 0x2b, 0x2b, 0xbe, 0xc6, 0xe9, 0xc1, 0x42, 0x83, 0x2a, 0x28, 0x38, 0x6a, 0x69, 0x5d,
	0x42, 0xe5, 0x5a, 0xfd, 0xd0, 0x10, 0xe8, 0xe1, 0x1e, 0x50, 0x43, 0x0a, 0x65, 0x13, 0x02, 0x3d,
	0xa0, 0x6c, 0xac, 0x45, 0x16, 0x96, 0x25, 0x75, 0x77, 0x2d, 0xf0, 0x5b, 0xf4, 0x11, 0xfa, 0x9c,
	0xfd, 0x54, 0xf6, 0x90, 0x6b, 0x85, 0x54, 0x90, 0x40, 0xf2, 0xc5, 0x8c, 0x76, 0x66, 0x7e, 0xb3,
	0x33, 0xff, 0xf1, 0xc2, 0xd3, 0x90, 0xe4, 0x6c, 0x98, 0xb3, 0x4c, 0x64, 0x43, 0x69, 0xc6, 0xa9,
	0xa0, 0x2c, 0x25, 0xc9, 0xb0, 0x18, 0x55, 0xbe, 0x7d, 0x15, 0x82, 0x5c, 0x79, 0xa6, 0x6d, 0xbf,
	0xe2, 0x2e, 0x46, 0xee, 0xee, 0x06, 0x6a, 0x9a, 0x2d, 0x16, 0x59, 0x2a, 0x21, 0xda, 0xd2, 0x29,
	0xee, 0x5e, 0x4d, 0x35, 0x92, 0xc7, 0x05, 0x65, 0x3c, 0x5e, 0x07, 0x3f, 0xae, 0x09, 0xe6, 0x82,
	0x88, 0x25, 0xd7, 0x81, 0xde, 0x18, 0x9a, 0xe3, 0xa9, 0xc8, 0x18, 0xba, 0x0f, 0x40, 0xa4, 0xf1,
	0x43, 0xac, 0x72, 0xea, 0x58, 0x7d, 0x6b, 0xd0, 0xc5, 0x5d, 0x75, 0x72, 0xbc, 0xca, 0x29, 0xba,
	0x07, 0x1d, 0xed, 0x8e, 0x43, 0xa7, 0xa1, 0x9c, 0x6d, 0xf5, 0x3d, 0x09, 0xbd, 0x5f, 0x36, 0xdc,
	0x9d, 0x18, 0xfe, 0x24, 0x2d, 0xb2, 0x39, 0xc5, 0xf4, 0xe7, 0x92, 0x72, 0x81, 0xf6, 0xc1, 0x2e,
	0x28, 0x53, 0xb0, 0x5b, 0xc1, 0x23, 0xff, 0xff, 0xfd, 0xfb, 0xe3, 0xcf, 0x93, 0x13, 0xdd, 0x00,
	0x96, 0x29, 0xe8, 0x1b, 0x74, 0x16, 0x54, 0x90, 0x90, 0x08, 0xe2, 0x34, 0xfa, 0xf6, 0xa0, 0x17,
	0xbc, 0xae, 0x4b, 0x3f, 0xb7, 0xbc, 0xff, 0xc9, 0x10, 0x3e, 0xa4, 0x82, 0xad, 0xf0, 0x1a, 0x88,
	0x5e, 0x42, 0x7b, 0x41, 0x39, 0x27, 0x11, 0x75, 0xec, 0xbe, 0x35, 0xe8, 0x05, 0x0f, 0x36, 0xd9,
	0x66, 0xe8, 0x8a, 0xba, 0x41, 0xc3, 0x65, 0x0e, 0x7a, 0x01, 0x4d, 0xd5, 0xba, 0xb3, 0xa5, 0x92,
	0x77, 0x6b, 0xfb, 0x92, 0x81, 0x58, 0xc7, 0xbb, 0x33, 0xb8, 0x59, 0xb9, 0x12, 0xba, 0x03, 0xf6,
	0x9c, 0xae, 0xcc, 0xb0, 0xa5, 0x89, 0xc6, 0xd0, 0x2c, 0x48, 0xb2, 0xa4, 0x6a, 0xc6, 0xbd, 0x60,
	0xaf, 0x8e, 0x7d, 0x18, 0x73, 0x71, 0x24, 0x58, 0x9c, 0x46, 0x27, 0x32, 0x05, 0xeb, 0xcc, 0x83,
	0xc6, 0xbe, 0xe5, 0xfd, 0xde, 0x82, 0x9d, 0xb3, 0x33, 0xe1, 0x79, 0x96, 0x72, 0x8a, 0x0e, 0xa0,
	0xa5, 0x17, 0x40, 0x95, 0xed, 0x05, 0x5e, 0x5d, 0x89, 0x23, 0x15, 0x89, 0x4d, 0x06, 0xfa, 0x02,
	0xed, 0x19, 0x25, 0x21, 0x65, 0xfc, 0x32, 0xa2, 0xe8, 0x0b, 0xf8, 0x1f, 0x35, 0x41, 0x8b, 0x52,
	0xf2, 0xd0, 0x77, 0xe8, 0x08, 0x46, 0xe2, 0x44, 0xb2, 0x6d, 0xc5, 0x7e, 0x73, 0x09, 0xf6, 0xb1,
	0x41, 0x18, 0xc5, 0x4b, 0x22, 0x7a, 0xf5, 0x4f, 0x71, 0x2d, 0xda, 0xc3, 0x7a, 0xc5, 0x35, 0x6e,
	0x2d, 0xb9, 0x1b, 0xc1, 0xf6, 0xe6, 0xb5, 0xaf, 0x4c, 0x38, 0xb9, 0x22, 0x95, 0x1e, 0xae, 0x6e,
	0x45, 0x9e, 0xc0, 0xed, 0x33, 0x5e, 0xb4, 0x03, 0x2d, 0xe5, 0x97, 0xab, 0x61, 0x0f, 0xba, 0xd8,
	0x7c, 0x05, 0x7f, 0x2c, 0xd8, 0x7e, 0x4f, 0x72, 0x56, 0x0e, 0x1d, 0x09, 0xe8, 0xbe, 0x23, 0x49,
	0xa2, 0x1f, 0x8e, 0xd1, 0x85, 0xff, 0x98, 0x6e, 0x70, 0x71, 0x69, 0xbd, 0x1b, 0x65, 0xd5, 0xc3,
	0x6c, 0x4a, 0x92, 0x6b, 0xab, 0xfa, 0xf6, 0xd9, 0x57, 0x3f, 0x8a, 0xc5, 0x6c, 0x79, 0x2a, 0xd7,
	0x44, 0xbd, 0xa5, 0xfa, 0x27, 0x9f, 0x47, 0xe7, 0xbf, 0xaf, 0xa7, 0x2d, 0x75, 0xfc, 0xfc, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xc0, 0x72, 0x2e, 0x1f, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DaprInternalClient is the client API for DaprInternal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DaprInternalClient interface {
	CallActor(ctx context.Context, in *InternalInvokeRequest, opts ...grpc.CallOption) (*InternalInvokeResponse, error)
	CallLocal(ctx context.Context, in *InternalInvokeRequest, opts ...grpc.CallOption) (*InternalInvokeResponse, error)
}

type daprInternalClient struct {
	cc *grpc.ClientConn
}

func NewDaprInternalClient(cc *grpc.ClientConn) DaprInternalClient {
	return &daprInternalClient{cc}
}

func (c *daprInternalClient) CallActor(ctx context.Context, in *InternalInvokeRequest, opts ...grpc.CallOption) (*InternalInvokeResponse, error) {
	out := new(InternalInvokeResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.daprinternal.v1.DaprInternal/CallActor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprInternalClient) CallLocal(ctx context.Context, in *InternalInvokeRequest, opts ...grpc.CallOption) (*InternalInvokeResponse, error) {
	out := new(InternalInvokeResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.daprinternal.v1.DaprInternal/CallLocal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaprInternalServer is the server API for DaprInternal service.
type DaprInternalServer interface {
	CallActor(context.Context, *InternalInvokeRequest) (*InternalInvokeResponse, error)
	CallLocal(context.Context, *InternalInvokeRequest) (*InternalInvokeResponse, error)
}

// UnimplementedDaprInternalServer can be embedded to have forward compatible implementations.
type UnimplementedDaprInternalServer struct {
}

func (*UnimplementedDaprInternalServer) CallActor(ctx context.Context, req *InternalInvokeRequest) (*InternalInvokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallActor not implemented")
}
func (*UnimplementedDaprInternalServer) CallLocal(ctx context.Context, req *InternalInvokeRequest) (*InternalInvokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallLocal not implemented")
}

func RegisterDaprInternalServer(s *grpc.Server, srv DaprInternalServer) {
	s.RegisterService(&_DaprInternal_serviceDesc, srv)
}

func _DaprInternal_CallActor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalInvokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprInternalServer).CallActor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.daprinternal.v1.DaprInternal/CallActor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprInternalServer).CallActor(ctx, req.(*InternalInvokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaprInternal_CallLocal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalInvokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprInternalServer).CallLocal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.daprinternal.v1.DaprInternal/CallLocal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprInternalServer).CallLocal(ctx, req.(*InternalInvokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DaprInternal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dapr.proto.daprinternal.v1.DaprInternal",
	HandlerType: (*DaprInternalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallActor",
			Handler:    _DaprInternal_CallActor_Handler,
		},
		{
			MethodName: "CallLocal",
			Handler:    _DaprInternal_CallLocal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dapr/proto/daprinternal/v1/daprinternal.proto",
}