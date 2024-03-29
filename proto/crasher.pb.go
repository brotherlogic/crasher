// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        (unknown)
// source: crasher.proto

package crasher

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CrashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CrashRequest) Reset() {
	*x = CrashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crasher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrashRequest) ProtoMessage() {}

func (x *CrashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crasher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrashRequest.ProtoReflect.Descriptor instead.
func (*CrashRequest) Descriptor() ([]byte, []int) {
	return file_crasher_proto_rawDescGZIP(), []int{0}
}

type CrashResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CrashResponse) Reset() {
	*x = CrashResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crasher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrashResponse) ProtoMessage() {}

func (x *CrashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crasher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrashResponse.ProtoReflect.Descriptor instead.
func (*CrashResponse) Descriptor() ([]byte, []int) {
	return file_crasher_proto_rawDescGZIP(), []int{1}
}

var File_crasher_proto protoreflect.FileDescriptor

var file_crasher_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x72, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x72, 0x61, 0x73, 0x68, 0x65, 0x72, 0x22, 0x0e, 0x0a, 0x0c, 0x43, 0x72, 0x61, 0x73,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x43, 0x72, 0x61, 0x73,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x48, 0x0a, 0x0c, 0x43, 0x72, 0x61,
	0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x43, 0x72, 0x61,
	0x73, 0x68, 0x12, 0x15, 0x2e, 0x63, 0x72, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x61,
	0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x72, 0x61, 0x73,
	0x68, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crasher_proto_rawDescOnce sync.Once
	file_crasher_proto_rawDescData = file_crasher_proto_rawDesc
)

func file_crasher_proto_rawDescGZIP() []byte {
	file_crasher_proto_rawDescOnce.Do(func() {
		file_crasher_proto_rawDescData = protoimpl.X.CompressGZIP(file_crasher_proto_rawDescData)
	})
	return file_crasher_proto_rawDescData
}

var file_crasher_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_crasher_proto_goTypes = []interface{}{
	(*CrashRequest)(nil),  // 0: crasher.CrashRequest
	(*CrashResponse)(nil), // 1: crasher.CrashResponse
}
var file_crasher_proto_depIdxs = []int32{
	0, // 0: crasher.CrashService.Crash:input_type -> crasher.CrashRequest
	1, // 1: crasher.CrashService.Crash:output_type -> crasher.CrashResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crasher_proto_init() }
func file_crasher_proto_init() {
	if File_crasher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crasher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrashRequest); i {
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
		file_crasher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrashResponse); i {
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
			RawDescriptor: file_crasher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crasher_proto_goTypes,
		DependencyIndexes: file_crasher_proto_depIdxs,
		MessageInfos:      file_crasher_proto_msgTypes,
	}.Build()
	File_crasher_proto = out.File
	file_crasher_proto_rawDesc = nil
	file_crasher_proto_goTypes = nil
	file_crasher_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CrashServiceClient is the client API for CrashService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CrashServiceClient interface {
	Crash(ctx context.Context, in *CrashRequest, opts ...grpc.CallOption) (*CrashResponse, error)
}

type crashServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrashServiceClient(cc grpc.ClientConnInterface) CrashServiceClient {
	return &crashServiceClient{cc}
}

func (c *crashServiceClient) Crash(ctx context.Context, in *CrashRequest, opts ...grpc.CallOption) (*CrashResponse, error) {
	out := new(CrashResponse)
	err := c.cc.Invoke(ctx, "/crasher.CrashService/Crash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrashServiceServer is the server API for CrashService service.
type CrashServiceServer interface {
	Crash(context.Context, *CrashRequest) (*CrashResponse, error)
}

// UnimplementedCrashServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCrashServiceServer struct {
}

func (*UnimplementedCrashServiceServer) Crash(context.Context, *CrashRequest) (*CrashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Crash not implemented")
}

func RegisterCrashServiceServer(s *grpc.Server, srv CrashServiceServer) {
	s.RegisterService(&_CrashService_serviceDesc, srv)
}

func _CrashService_Crash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CrashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrashServiceServer).Crash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crasher.CrashService/Crash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrashServiceServer).Crash(ctx, req.(*CrashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CrashService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crasher.CrashService",
	HandlerType: (*CrashServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Crash",
			Handler:    _CrashService_Crash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crasher.proto",
}
