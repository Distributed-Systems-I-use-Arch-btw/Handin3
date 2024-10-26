// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: gRPC/chittychat.proto

package proto

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

type Messages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []string `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *Messages) Reset() {
	*x = Messages{}
	mi := &file_gRPC_chittychat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Messages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Messages) ProtoMessage() {}

func (x *Messages) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_chittychat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Messages.ProtoReflect.Descriptor instead.
func (*Messages) Descriptor() ([]byte, []int) {
	return file_gRPC_chittychat_proto_rawDescGZIP(), []int{0}
}

func (x *Messages) GetMessages() []string {
	if x != nil {
		return x.Messages
	}
	return nil
}

type VectorClock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vectorclock []int32 `protobuf:"varint,1,rep,packed,name=vectorclock,proto3" json:"vectorclock,omitempty"`
}

func (x *VectorClock) Reset() {
	*x = VectorClock{}
	mi := &file_gRPC_chittychat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VectorClock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VectorClock) ProtoMessage() {}

func (x *VectorClock) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_chittychat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VectorClock.ProtoReflect.Descriptor instead.
func (*VectorClock) Descriptor() ([]byte, []int) {
	return file_gRPC_chittychat_proto_rawDescGZIP(), []int{1}
}

func (x *VectorClock) GetVectorclock() []int32 {
	if x != nil {
		return x.Vectorclock
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_gRPC_chittychat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_chittychat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_gRPC_chittychat_proto_rawDescGZIP(), []int{2}
}

type MessagePackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message     *Messages    `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Vectorclock *VectorClock `protobuf:"bytes,2,opt,name=vectorclock,proto3" json:"vectorclock,omitempty"`
}

func (x *MessagePackage) Reset() {
	*x = MessagePackage{}
	mi := &file_gRPC_chittychat_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessagePackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagePackage) ProtoMessage() {}

func (x *MessagePackage) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_chittychat_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagePackage.ProtoReflect.Descriptor instead.
func (*MessagePackage) Descriptor() ([]byte, []int) {
	return file_gRPC_chittychat_proto_rawDescGZIP(), []int{3}
}

func (x *MessagePackage) GetMessage() *Messages {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *MessagePackage) GetVectorclock() *VectorClock {
	if x != nil {
		return x.Vectorclock
	}
	return nil
}

type ClientId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clientid int32 `protobuf:"varint,1,opt,name=clientid,proto3" json:"clientid,omitempty"`
}

func (x *ClientId) Reset() {
	*x = ClientId{}
	mi := &file_gRPC_chittychat_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientId) ProtoMessage() {}

func (x *ClientId) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_chittychat_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientId.ProtoReflect.Descriptor instead.
func (*ClientId) Descriptor() ([]byte, []int) {
	return file_gRPC_chittychat_proto_rawDescGZIP(), []int{4}
}

func (x *ClientId) GetClientid() int32 {
	if x != nil {
		return x.Clientid
	}
	return 0
}

var File_gRPC_chittychat_proto protoreflect.FileDescriptor

var file_gRPC_chittychat_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22,
	0x2f, 0x0a, 0x0b, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x20,
	0x0a, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x65, 0x0a, 0x0e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2e, 0x0a, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x63, 0x6c, 0x6f, 0x63, 0x6b,
	0x22, 0x26, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x32, 0x89, 0x01, 0x0a, 0x0a, 0x43, 0x68, 0x69,
	0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x12, 0x28, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x12, 0x22, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x09, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x09, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68,
	0x61, 0x74, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gRPC_chittychat_proto_rawDescOnce sync.Once
	file_gRPC_chittychat_proto_rawDescData = file_gRPC_chittychat_proto_rawDesc
)

func file_gRPC_chittychat_proto_rawDescGZIP() []byte {
	file_gRPC_chittychat_proto_rawDescOnce.Do(func() {
		file_gRPC_chittychat_proto_rawDescData = protoimpl.X.CompressGZIP(file_gRPC_chittychat_proto_rawDescData)
	})
	return file_gRPC_chittychat_proto_rawDescData
}

var file_gRPC_chittychat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gRPC_chittychat_proto_goTypes = []any{
	(*Messages)(nil),       // 0: Messages
	(*VectorClock)(nil),    // 1: VectorClock
	(*Empty)(nil),          // 2: Empty
	(*MessagePackage)(nil), // 3: MessagePackage
	(*ClientId)(nil),       // 4: ClientId
}
var file_gRPC_chittychat_proto_depIdxs = []int32{
	0, // 0: MessagePackage.message:type_name -> Messages
	1, // 1: MessagePackage.vectorclock:type_name -> VectorClock
	2, // 2: ChittyChat.GetMessages:input_type -> Empty
	0, // 3: ChittyChat.PostMessage:input_type -> Messages
	2, // 4: ChittyChat.CreateClientIdentifier:input_type -> Empty
	3, // 5: ChittyChat.GetMessages:output_type -> MessagePackage
	2, // 6: ChittyChat.PostMessage:output_type -> Empty
	4, // 7: ChittyChat.CreateClientIdentifier:output_type -> ClientId
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gRPC_chittychat_proto_init() }
func file_gRPC_chittychat_proto_init() {
	if File_gRPC_chittychat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gRPC_chittychat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gRPC_chittychat_proto_goTypes,
		DependencyIndexes: file_gRPC_chittychat_proto_depIdxs,
		MessageInfos:      file_gRPC_chittychat_proto_msgTypes,
	}.Build()
	File_gRPC_chittychat_proto = out.File
	file_gRPC_chittychat_proto_rawDesc = nil
	file_gRPC_chittychat_proto_goTypes = nil
	file_gRPC_chittychat_proto_depIdxs = nil
}
