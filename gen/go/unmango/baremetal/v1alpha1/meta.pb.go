// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: unmango/baremetal/v1alpha1/meta.proto

package baremetalv1alpha1

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
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unmango_baremetal_v1alpha1_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_unmango_baremetal_v1alpha1_meta_proto_msgTypes[0]
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
	return file_unmango_baremetal_v1alpha1_meta_proto_rawDescGZIP(), []int{0}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unmango_baremetal_v1alpha1_meta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_unmango_baremetal_v1alpha1_meta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_unmango_baremetal_v1alpha1_meta_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type VersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VersionRequest) Reset() {
	*x = VersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unmango_baremetal_v1alpha1_meta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionRequest) ProtoMessage() {}

func (x *VersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_unmango_baremetal_v1alpha1_meta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionRequest.ProtoReflect.Descriptor instead.
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return file_unmango_baremetal_v1alpha1_meta_proto_rawDescGZIP(), []int{2}
}

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unmango_baremetal_v1alpha1_meta_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_unmango_baremetal_v1alpha1_meta_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_unmango_baremetal_v1alpha1_meta_proto_rawDescGZIP(), []int{3}
}

func (x *VersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_unmango_baremetal_v1alpha1_meta_proto protoreflect.FileDescriptor

var file_unmango_baremetal_v1alpha1_meta_proto_rawDesc = []byte{
	0x0a, 0x25, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f,
	0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x10, 0x0a, 0x0e,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b,
	0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0xcc, 0x01, 0x0a, 0x0b,
	0x4d, 0x65, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x27, 0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x62, 0x61,
	0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x75,
	0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61,
	0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x8e, 0x02, 0x0a, 0x1e, 0x63,
	0x6f, 0x6d, 0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x09, 0x4d,
	0x65, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2f, 0x70,
	0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2d, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2f, 0x62,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3b, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x42, 0x58, 0xaa, 0x02, 0x1a, 0x55, 0x6e, 0x6d, 0x61,
	0x6e, 0x67, 0x6f, 0x2e, 0x42, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1a, 0x55, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f,
	0x5c, 0x42, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xe2, 0x02, 0x26, 0x55, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x5c, 0x42, 0x61,
	0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x55,
	0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x3a, 0x3a, 0x42, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61,
	0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_unmango_baremetal_v1alpha1_meta_proto_rawDescOnce sync.Once
	file_unmango_baremetal_v1alpha1_meta_proto_rawDescData = file_unmango_baremetal_v1alpha1_meta_proto_rawDesc
)

func file_unmango_baremetal_v1alpha1_meta_proto_rawDescGZIP() []byte {
	file_unmango_baremetal_v1alpha1_meta_proto_rawDescOnce.Do(func() {
		file_unmango_baremetal_v1alpha1_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_unmango_baremetal_v1alpha1_meta_proto_rawDescData)
	})
	return file_unmango_baremetal_v1alpha1_meta_proto_rawDescData
}

var file_unmango_baremetal_v1alpha1_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_unmango_baremetal_v1alpha1_meta_proto_goTypes = []any{
	(*PingRequest)(nil),     // 0: unmango.baremetal.v1alpha1.PingRequest
	(*PingResponse)(nil),    // 1: unmango.baremetal.v1alpha1.PingResponse
	(*VersionRequest)(nil),  // 2: unmango.baremetal.v1alpha1.VersionRequest
	(*VersionResponse)(nil), // 3: unmango.baremetal.v1alpha1.VersionResponse
}
var file_unmango_baremetal_v1alpha1_meta_proto_depIdxs = []int32{
	0, // 0: unmango.baremetal.v1alpha1.MetaService.Ping:input_type -> unmango.baremetal.v1alpha1.PingRequest
	2, // 1: unmango.baremetal.v1alpha1.MetaService.Version:input_type -> unmango.baremetal.v1alpha1.VersionRequest
	1, // 2: unmango.baremetal.v1alpha1.MetaService.Ping:output_type -> unmango.baremetal.v1alpha1.PingResponse
	3, // 3: unmango.baremetal.v1alpha1.MetaService.Version:output_type -> unmango.baremetal.v1alpha1.VersionResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_unmango_baremetal_v1alpha1_meta_proto_init() }
func file_unmango_baremetal_v1alpha1_meta_proto_init() {
	if File_unmango_baremetal_v1alpha1_meta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_unmango_baremetal_v1alpha1_meta_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_unmango_baremetal_v1alpha1_meta_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PingResponse); i {
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
		file_unmango_baremetal_v1alpha1_meta_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*VersionRequest); i {
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
		file_unmango_baremetal_v1alpha1_meta_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*VersionResponse); i {
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
			RawDescriptor: file_unmango_baremetal_v1alpha1_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_unmango_baremetal_v1alpha1_meta_proto_goTypes,
		DependencyIndexes: file_unmango_baremetal_v1alpha1_meta_proto_depIdxs,
		MessageInfos:      file_unmango_baremetal_v1alpha1_meta_proto_msgTypes,
	}.Build()
	File_unmango_baremetal_v1alpha1_meta_proto = out.File
	file_unmango_baremetal_v1alpha1_meta_proto_rawDesc = nil
	file_unmango_baremetal_v1alpha1_meta_proto_goTypes = nil
	file_unmango_baremetal_v1alpha1_meta_proto_depIdxs = nil
}
