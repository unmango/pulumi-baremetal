// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: unmango/baremetal/v1alpha1/command.proto

package baremetalv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *TeeRequest) Reset() {
	*x = TeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeeRequest) ProtoMessage() {}

func (x *TeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeeRequest.ProtoReflect.Descriptor instead.
func (*TeeRequest) Descriptor() ([]byte, []int) {
	return file_unmango_baremetal_v1alpha1_command_proto_rawDescGZIP(), []int{0}
}

func (x *TeeRequest) GetState() *State {
	if x != nil {
		return x.State
	}
	return nil
}

type TeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *TeeResponse) Reset() {
	*x = TeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeeResponse) ProtoMessage() {}

func (x *TeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeeResponse.ProtoReflect.Descriptor instead.
func (*TeeResponse) Descriptor() ([]byte, []int) {
	return file_unmango_baremetal_v1alpha1_command_proto_rawDescGZIP(), []int{1}
}

func (x *TeeResponse) GetState() *State {
	if x != nil {
		return x.State
	}
	return nil
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pulumi *anypb.Any `protobuf:"bytes,1,opt,name=pulumi,proto3" json:"pulumi,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_unmango_baremetal_v1alpha1_command_proto_rawDescGZIP(), []int{2}
}

func (x *State) GetPulumi() *anypb.Any {
	if x != nil {
		return x.Pulumi
	}
	return nil
}

var File_unmango_baremetal_v1alpha1_command_proto protoreflect.FileDescriptor

var file_unmango_baremetal_v1alpha1_command_proto_rawDesc = []byte{
	0x0a, 0x28, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x75, 0x6e, 0x6d, 0x61,
	0x6e, 0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x45, 0x0a, 0x0a, 0x54, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x37, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x46, 0x0a, 0x0b, 0x54, 0x65, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f,
	0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x22, 0x35, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x70, 0x75, 0x6c,
	0x75, 0x6d, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x06, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x32, 0x68, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x03, 0x54, 0x65, 0x65,
	0x12, 0x26, 0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x65,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e,
	0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x91, 0x02, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67,
	0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2d,
	0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x62, 0x61, 0x72, 0x65,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03,
	0x55, 0x42, 0x58, 0xaa, 0x02, 0x1a, 0x55, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x42, 0x61,
	0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xca, 0x02, 0x1a, 0x55, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x5c, 0x42, 0x61, 0x72, 0x65, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x26,
	0x55, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x5c, 0x42, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61,
	0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x55, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f,
	0x3a, 0x3a, 0x42, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_unmango_baremetal_v1alpha1_command_proto_rawDescOnce sync.Once
	file_unmango_baremetal_v1alpha1_command_proto_rawDescData = file_unmango_baremetal_v1alpha1_command_proto_rawDesc
)

func file_unmango_baremetal_v1alpha1_command_proto_rawDescGZIP() []byte {
	file_unmango_baremetal_v1alpha1_command_proto_rawDescOnce.Do(func() {
		file_unmango_baremetal_v1alpha1_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_unmango_baremetal_v1alpha1_command_proto_rawDescData)
	})
	return file_unmango_baremetal_v1alpha1_command_proto_rawDescData
}

var file_unmango_baremetal_v1alpha1_command_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_unmango_baremetal_v1alpha1_command_proto_goTypes = []any{
	(*TeeRequest)(nil),  // 0: unmango.baremetal.v1alpha1.TeeRequest
	(*TeeResponse)(nil), // 1: unmango.baremetal.v1alpha1.TeeResponse
	(*State)(nil),       // 2: unmango.baremetal.v1alpha1.State
	(*anypb.Any)(nil),   // 3: google.protobuf.Any
}
var file_unmango_baremetal_v1alpha1_command_proto_depIdxs = []int32{
	2, // 0: unmango.baremetal.v1alpha1.TeeRequest.state:type_name -> unmango.baremetal.v1alpha1.State
	2, // 1: unmango.baremetal.v1alpha1.TeeResponse.state:type_name -> unmango.baremetal.v1alpha1.State
	3, // 2: unmango.baremetal.v1alpha1.State.pulumi:type_name -> google.protobuf.Any
	0, // 3: unmango.baremetal.v1alpha1.CommandService.Tee:input_type -> unmango.baremetal.v1alpha1.TeeRequest
	1, // 4: unmango.baremetal.v1alpha1.CommandService.Tee:output_type -> unmango.baremetal.v1alpha1.TeeResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_unmango_baremetal_v1alpha1_command_proto_init() }
func file_unmango_baremetal_v1alpha1_command_proto_init() {
	if File_unmango_baremetal_v1alpha1_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_unmango_baremetal_v1alpha1_command_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TeeRequest); i {
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
		file_unmango_baremetal_v1alpha1_command_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TeeResponse); i {
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
		file_unmango_baremetal_v1alpha1_command_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*State); i {
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
			RawDescriptor: file_unmango_baremetal_v1alpha1_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_unmango_baremetal_v1alpha1_command_proto_goTypes,
		DependencyIndexes: file_unmango_baremetal_v1alpha1_command_proto_depIdxs,
		MessageInfos:      file_unmango_baremetal_v1alpha1_command_proto_msgTypes,
	}.Build()
	File_unmango_baremetal_v1alpha1_command_proto = out.File
	file_unmango_baremetal_v1alpha1_command_proto_rawDesc = nil
	file_unmango_baremetal_v1alpha1_command_proto_goTypes = nil
	file_unmango_baremetal_v1alpha1_command_proto_depIdxs = nil
}
