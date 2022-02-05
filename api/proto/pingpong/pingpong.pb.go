// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: api/proto/pingpong/pingpong.proto

package pingpong

import (
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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_pingpong_pingpong_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_pingpong_pingpong_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_pingpong_pingpong_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PongResponse) Reset() {
	*x = PongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_pingpong_pingpong_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongResponse) ProtoMessage() {}

func (x *PongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_pingpong_pingpong_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongResponse.ProtoReflect.Descriptor instead.
func (*PongResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_pingpong_pingpong_proto_rawDescGZIP(), []int{1}
}

func (x *PongResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_api_proto_pingpong_pingpong_proto protoreflect.FileDescriptor

var file_api_proto_pingpong_pingpong_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69, 0x6e, 0x67,
	0x70, 0x6f, 0x6e, 0x67, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x1f, 0x0a,
	0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x20,
	0x0a, 0x0c, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x32, 0x86, 0x01, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x37, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70,
	0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x50, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x69,
	0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x61, 0x72, 0x64, 0x6f, 0x7a, 0x6f,
	0x73, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x3b, 0x70,
	0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x3b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_pingpong_pingpong_proto_rawDescOnce sync.Once
	file_api_proto_pingpong_pingpong_proto_rawDescData = file_api_proto_pingpong_pingpong_proto_rawDesc
)

func file_api_proto_pingpong_pingpong_proto_rawDescGZIP() []byte {
	file_api_proto_pingpong_pingpong_proto_rawDescOnce.Do(func() {
		file_api_proto_pingpong_pingpong_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_pingpong_pingpong_proto_rawDescData)
	})
	return file_api_proto_pingpong_pingpong_proto_rawDescData
}

var file_api_proto_pingpong_pingpong_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_pingpong_pingpong_proto_goTypes = []interface{}{
	(*PingRequest)(nil),  // 0: pingpong.PingRequest
	(*PongResponse)(nil), // 1: pingpong.PongResponse
}
var file_api_proto_pingpong_pingpong_proto_depIdxs = []int32{
	0, // 0: pingpong.PingPong.Ping:input_type -> pingpong.PingRequest
	0, // 1: pingpong.PingPong.PingStream:input_type -> pingpong.PingRequest
	1, // 2: pingpong.PingPong.Ping:output_type -> pingpong.PongResponse
	1, // 3: pingpong.PingPong.PingStream:output_type -> pingpong.PongResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_pingpong_pingpong_proto_init() }
func file_api_proto_pingpong_pingpong_proto_init() {
	if File_api_proto_pingpong_pingpong_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_pingpong_pingpong_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_api_proto_pingpong_pingpong_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongResponse); i {
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
			RawDescriptor: file_api_proto_pingpong_pingpong_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_pingpong_pingpong_proto_goTypes,
		DependencyIndexes: file_api_proto_pingpong_pingpong_proto_depIdxs,
		MessageInfos:      file_api_proto_pingpong_pingpong_proto_msgTypes,
	}.Build()
	File_api_proto_pingpong_pingpong_proto = out.File
	file_api_proto_pingpong_pingpong_proto_rawDesc = nil
	file_api_proto_pingpong_pingpong_proto_goTypes = nil
	file_api_proto_pingpong_pingpong_proto_depIdxs = nil
}
