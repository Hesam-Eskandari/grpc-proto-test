// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: src/mahsa/mahsa.proto

package mahsa

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloMahsaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greet string `protobuf:"bytes,1,opt,name=greet,proto3" json:"greet,omitempty"`
}

func (x *HelloMahsaRequest) Reset() {
	*x = HelloMahsaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_mahsa_mahsa_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloMahsaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloMahsaRequest) ProtoMessage() {}

func (x *HelloMahsaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_mahsa_mahsa_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloMahsaRequest.ProtoReflect.Descriptor instead.
func (*HelloMahsaRequest) Descriptor() ([]byte, []int) {
	return file_src_mahsa_mahsa_proto_rawDescGZIP(), []int{0}
}

func (x *HelloMahsaRequest) GetGreet() string {
	if x != nil {
		return x.Greet
	}
	return ""
}

type HelloMahsaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nemikham string `protobuf:"bytes,1,opt,name=nemikham,proto3" json:"nemikham,omitempty"`
}

func (x *HelloMahsaResponse) Reset() {
	*x = HelloMahsaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_mahsa_mahsa_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloMahsaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloMahsaResponse) ProtoMessage() {}

func (x *HelloMahsaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_mahsa_mahsa_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloMahsaResponse.ProtoReflect.Descriptor instead.
func (*HelloMahsaResponse) Descriptor() ([]byte, []int) {
	return file_src_mahsa_mahsa_proto_rawDescGZIP(), []int{1}
}

func (x *HelloMahsaResponse) GetNemikham() string {
	if x != nil {
		return x.Nemikham
	}
	return ""
}

var File_src_mahsa_mahsa_proto protoreflect.FileDescriptor

var file_src_mahsa_mahsa_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x61, 0x68, 0x73, 0x61, 0x2f, 0x6d, 0x61, 0x68, 0x73,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x4d, 0x61, 0x68, 0x73, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x22, 0x30, 0x0a, 0x12, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x61, 0x68, 0x73, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x6d, 0x69,
	0x6b, 0x68, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x6d, 0x69,
	0x6b, 0x68, 0x61, 0x6d, 0x32, 0x50, 0x0a, 0x10, 0x47, 0x52, 0x50, 0x43, 0x4d, 0x61, 0x68, 0x73,
	0x61, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0f, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x54, 0x6f, 0x4d, 0x61, 0x68, 0x73, 0x61, 0x12, 0x12, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x4d, 0x61, 0x68, 0x73, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x61, 0x68, 0x73, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x61,
	0x68, 0x73, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_mahsa_mahsa_proto_rawDescOnce sync.Once
	file_src_mahsa_mahsa_proto_rawDescData = file_src_mahsa_mahsa_proto_rawDesc
)

func file_src_mahsa_mahsa_proto_rawDescGZIP() []byte {
	file_src_mahsa_mahsa_proto_rawDescOnce.Do(func() {
		file_src_mahsa_mahsa_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_mahsa_mahsa_proto_rawDescData)
	})
	return file_src_mahsa_mahsa_proto_rawDescData
}

var file_src_mahsa_mahsa_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_mahsa_mahsa_proto_goTypes = []interface{}{
	(*HelloMahsaRequest)(nil),  // 0: HelloMahsaRequest
	(*HelloMahsaResponse)(nil), // 1: HelloMahsaResponse
}
var file_src_mahsa_mahsa_proto_depIdxs = []int32{
	0, // 0: GRPCMahsaGreeter.SayHelloToMahsa:input_type -> HelloMahsaRequest
	1, // 1: GRPCMahsaGreeter.SayHelloToMahsa:output_type -> HelloMahsaResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_mahsa_mahsa_proto_init() }
func file_src_mahsa_mahsa_proto_init() {
	if File_src_mahsa_mahsa_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_mahsa_mahsa_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloMahsaRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_mahsa_mahsa_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloMahsaResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_mahsa_mahsa_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_mahsa_mahsa_proto_goTypes,
		DependencyIndexes: file_src_mahsa_mahsa_proto_depIdxs,
		MessageInfos:      file_src_mahsa_mahsa_proto_msgTypes,
	}.Build()
	File_src_mahsa_mahsa_proto = out.File
	file_src_mahsa_mahsa_proto_rawDesc = nil
	file_src_mahsa_mahsa_proto_goTypes = nil
	file_src_mahsa_mahsa_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GRPCMahsaGreeterClient is the client API for GRPCMahsaGreeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GRPCMahsaGreeterClient interface {
	SayHelloToMahsa(ctx context.Context, in *HelloMahsaRequest, opts ...grpc.CallOption) (*HelloMahsaResponse, error)
}

type gRPCMahsaGreeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCMahsaGreeterClient(cc grpc.ClientConnInterface) GRPCMahsaGreeterClient {
	return &gRPCMahsaGreeterClient{cc}
}

func (c *gRPCMahsaGreeterClient) SayHelloToMahsa(ctx context.Context, in *HelloMahsaRequest, opts ...grpc.CallOption) (*HelloMahsaResponse, error) {
	out := new(HelloMahsaResponse)
	err := c.cc.Invoke(ctx, "/GRPCMahsaGreeter/SayHelloToMahsa", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCMahsaGreeterServer is the server API for GRPCMahsaGreeter service.
type GRPCMahsaGreeterServer interface {
	SayHelloToMahsa(context.Context, *HelloMahsaRequest) (*HelloMahsaResponse, error)
}

// UnimplementedGRPCMahsaGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGRPCMahsaGreeterServer struct {
}

func (*UnimplementedGRPCMahsaGreeterServer) SayHelloToMahsa(context.Context, *HelloMahsaRequest) (*HelloMahsaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloToMahsa not implemented")
}

func RegisterGRPCMahsaGreeterServer(s *grpc.Server, srv GRPCMahsaGreeterServer) {
	s.RegisterService(&_GRPCMahsaGreeter_serviceDesc, srv)
}

func _GRPCMahsaGreeter_SayHelloToMahsa_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloMahsaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCMahsaGreeterServer).SayHelloToMahsa(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GRPCMahsaGreeter/SayHelloToMahsa",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCMahsaGreeterServer).SayHelloToMahsa(ctx, req.(*HelloMahsaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GRPCMahsaGreeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GRPCMahsaGreeter",
	HandlerType: (*GRPCMahsaGreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHelloToMahsa",
			Handler:    _GRPCMahsaGreeter_SayHelloToMahsa_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/mahsa/mahsa.proto",
}
