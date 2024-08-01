// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: unmango/baremetal/v1alpha1/command.proto

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

type Bin int32

const (
	Bin_BIN_UNSPECIFIED Bin = 0
	Bin_BIN_TEE         Bin = 1
	Bin_BIN_RM          Bin = 2
)

// Enum value maps for Bin.
var (
	Bin_name = map[int32]string{
		0: "BIN_UNSPECIFIED",
		1: "BIN_TEE",
		2: "BIN_RM",
	}
	Bin_value = map[string]int32{
		"BIN_UNSPECIFIED": 0,
		"BIN_TEE":         1,
		"BIN_RM":          2,
	}
)

func (x Bin) Enum() *Bin {
	p := new(Bin)
	*p = x
	return p
}

func (x Bin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Bin) Descriptor() protoreflect.EnumDescriptor {
	return file_unmango_baremetal_v1alpha1_command_proto_enumTypes[0].Descriptor()
}

func (Bin) Type() protoreflect.EnumType {
	return &file_unmango_baremetal_v1alpha1_command_proto_enumTypes[0]
}

func (x Bin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Bin.Descriptor instead.
func (Bin) EnumDescriptor() ([]byte, []int) {
	return file_unmango_baremetal_v1alpha1_command_proto_rawDescGZIP(), []int{0}
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command *Command `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_unmango_baremetal_v1alpha1_command_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetCommand() *Command {
	if x != nil {
		return x.Command
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_unmango_baremetal_v1alpha1_command_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Create *Command `protobuf:"bytes,1,opt,name=create,proto3" json:"create,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_unmango_baremetal_v1alpha1_command_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteRequest) GetCreate() *Command {
	if x != nil {
		return x.Create
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op *Operation `protobuf:"bytes,1,opt,name=op,proto3,oneof" json:"op,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_unmango_baremetal_v1alpha1_command_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteResponse) GetOp() *Operation {
	if x != nil {
		return x.Op
	}
	return nil
}

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  *Result  `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Command *Command `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_unmango_baremetal_v1alpha1_command_proto_rawDescGZIP(), []int{4}
}

func (x *Operation) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *Operation) GetCommand() *Command {
	if x != nil {
		return x.Command
	}
	return nil
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bin   Bin      `protobuf:"varint,1,opt,name=bin,proto3,enum=unmango.baremetal.v1alpha1.Bin" json:"bin,omitempty"`
	Args  []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	Stdin *string  `protobuf:"bytes,3,opt,name=stdin,proto3,oneof" json:"stdin,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_unmango_baremetal_v1alpha1_command_proto_rawDescGZIP(), []int{5}
}

func (x *Command) GetBin() Bin {
	if x != nil {
		return x.Bin
	}
	return Bin_BIN_UNSPECIFIED
}

func (x *Command) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *Command) GetStdin() string {
	if x != nil && x.Stdin != nil {
		return *x.Stdin
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExitCode int32  `protobuf:"varint,1,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	Stdout   string `protobuf:"bytes,2,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr   string `protobuf:"bytes,3,opt,name=stderr,proto3" json:"stderr,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_unmango_baremetal_v1alpha1_command_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_unmango_baremetal_v1alpha1_command_proto_rawDescGZIP(), []int{6}
}

func (x *Result) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *Result) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *Result) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

var File_unmango_baremetal_v1alpha1_command_proto protoreflect.FileDescriptor

var file_unmango_baremetal_v1alpha1_command_proto_rawDesc = []byte{
	0x0a, 0x28, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x75, 0x6e, 0x6d, 0x61,
	0x6e, 0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x4e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e,
	0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x4c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e,
	0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x4c, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e,
	0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x22, 0x53, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x02, 0x6f, 0x70, 0x88, 0x01, 0x01,
	0x42, 0x05, 0x0a, 0x03, 0x5f, 0x6f, 0x70, 0x22, 0x86, 0x01, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e,
	0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72,
	0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x22, 0x75, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x31, 0x0a, 0x03, 0x62,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e,
	0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x52, 0x03, 0x62, 0x69, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72,
	0x67, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x22, 0x55, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x2a, 0x33,
	0x0a, 0x03, 0x42, 0x69, 0x6e, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x49, 0x4e, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x49,
	0x4e, 0x5f, 0x54, 0x45, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x49, 0x4e, 0x5f, 0x52,
	0x4d, 0x10, 0x02, 0x32, 0xd2, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x29, 0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x75, 0x6e,
	0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x29, 0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x75,
	0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x91, 0x02, 0x0a, 0x1e, 0x63, 0x6f, 0x6d,
	0x2e, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x57, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2f,
	0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2d, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2f,
	0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x3b, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x42, 0x58, 0xaa, 0x02, 0x1a, 0x55, 0x6e, 0x6d,
	0x61, 0x6e, 0x67, 0x6f, 0x2e, 0x42, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1a, 0x55, 0x6e, 0x6d, 0x61, 0x6e, 0x67,
	0x6f, 0x5c, 0x42, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x26, 0x55, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x5c, 0x42,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c,
	0x55, 0x6e, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x3a, 0x3a, 0x42, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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

var file_unmango_baremetal_v1alpha1_command_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_unmango_baremetal_v1alpha1_command_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_unmango_baremetal_v1alpha1_command_proto_goTypes = []any{
	(Bin)(0),               // 0: unmango.baremetal.v1alpha1.Bin
	(*CreateRequest)(nil),  // 1: unmango.baremetal.v1alpha1.CreateRequest
	(*CreateResponse)(nil), // 2: unmango.baremetal.v1alpha1.CreateResponse
	(*DeleteRequest)(nil),  // 3: unmango.baremetal.v1alpha1.DeleteRequest
	(*DeleteResponse)(nil), // 4: unmango.baremetal.v1alpha1.DeleteResponse
	(*Operation)(nil),      // 5: unmango.baremetal.v1alpha1.Operation
	(*Command)(nil),        // 6: unmango.baremetal.v1alpha1.Command
	(*Result)(nil),         // 7: unmango.baremetal.v1alpha1.Result
}
var file_unmango_baremetal_v1alpha1_command_proto_depIdxs = []int32{
	6, // 0: unmango.baremetal.v1alpha1.CreateRequest.command:type_name -> unmango.baremetal.v1alpha1.Command
	7, // 1: unmango.baremetal.v1alpha1.CreateResponse.result:type_name -> unmango.baremetal.v1alpha1.Result
	6, // 2: unmango.baremetal.v1alpha1.DeleteRequest.create:type_name -> unmango.baremetal.v1alpha1.Command
	5, // 3: unmango.baremetal.v1alpha1.DeleteResponse.op:type_name -> unmango.baremetal.v1alpha1.Operation
	7, // 4: unmango.baremetal.v1alpha1.Operation.result:type_name -> unmango.baremetal.v1alpha1.Result
	6, // 5: unmango.baremetal.v1alpha1.Operation.command:type_name -> unmango.baremetal.v1alpha1.Command
	0, // 6: unmango.baremetal.v1alpha1.Command.bin:type_name -> unmango.baremetal.v1alpha1.Bin
	1, // 7: unmango.baremetal.v1alpha1.CommandService.Create:input_type -> unmango.baremetal.v1alpha1.CreateRequest
	3, // 8: unmango.baremetal.v1alpha1.CommandService.Delete:input_type -> unmango.baremetal.v1alpha1.DeleteRequest
	2, // 9: unmango.baremetal.v1alpha1.CommandService.Create:output_type -> unmango.baremetal.v1alpha1.CreateResponse
	4, // 10: unmango.baremetal.v1alpha1.CommandService.Delete:output_type -> unmango.baremetal.v1alpha1.DeleteResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_unmango_baremetal_v1alpha1_command_proto_init() }
func file_unmango_baremetal_v1alpha1_command_proto_init() {
	if File_unmango_baremetal_v1alpha1_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_unmango_baremetal_v1alpha1_command_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRequest); i {
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
			switch v := v.(*CreateResponse); i {
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
			switch v := v.(*DeleteRequest); i {
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
		file_unmango_baremetal_v1alpha1_command_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteResponse); i {
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
		file_unmango_baremetal_v1alpha1_command_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Operation); i {
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
		file_unmango_baremetal_v1alpha1_command_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Command); i {
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
		file_unmango_baremetal_v1alpha1_command_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Result); i {
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
	file_unmango_baremetal_v1alpha1_command_proto_msgTypes[3].OneofWrappers = []any{}
	file_unmango_baremetal_v1alpha1_command_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_unmango_baremetal_v1alpha1_command_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_unmango_baremetal_v1alpha1_command_proto_goTypes,
		DependencyIndexes: file_unmango_baremetal_v1alpha1_command_proto_depIdxs,
		EnumInfos:         file_unmango_baremetal_v1alpha1_command_proto_enumTypes,
		MessageInfos:      file_unmango_baremetal_v1alpha1_command_proto_msgTypes,
	}.Build()
	File_unmango_baremetal_v1alpha1_command_proto = out.File
	file_unmango_baremetal_v1alpha1_command_proto_rawDesc = nil
	file_unmango_baremetal_v1alpha1_command_proto_goTypes = nil
	file_unmango_baremetal_v1alpha1_command_proto_depIdxs = nil
}
