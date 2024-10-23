// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: chittychat.proto

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

type Students struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Students []string `protobuf:"bytes,1,rep,name=students,proto3" json:"students,omitempty"`
}

func (x *Students) Reset() {
	*x = Students{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chittychat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Students) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Students) ProtoMessage() {}

func (x *Students) ProtoReflect() protoreflect.Message {
	mi := &file_chittychat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Students.ProtoReflect.Descriptor instead.
func (*Students) Descriptor() ([]byte, []int) {
	return file_chittychat_proto_rawDescGZIP(), []int{0}
}

func (x *Students) GetStudents() []string {
	if x != nil {
		return x.Students
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
	if protoimpl.UnsafeEnabled {
		mi := &file_chittychat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_chittychat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_chittychat_proto_rawDescGZIP(), []int{1}
}

var File_chittychat_proto protoreflect.FileDescriptor

var file_chittychat_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x26, 0x0a, 0x08, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x32, 0x30, 0x0a, 0x0a, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61,
	0x74, 0x12, 0x22, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x09, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43,
	0x68, 0x61, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_chittychat_proto_rawDescOnce sync.Once
	file_chittychat_proto_rawDescData = file_chittychat_proto_rawDesc
)

func file_chittychat_proto_rawDescGZIP() []byte {
	file_chittychat_proto_rawDescOnce.Do(func() {
		file_chittychat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chittychat_proto_rawDescData)
	})
	return file_chittychat_proto_rawDescData
}

var file_chittychat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chittychat_proto_goTypes = []interface{}{
	(*Students)(nil), // 0: Students
	(*Empty)(nil),    // 1: Empty
}
var file_chittychat_proto_depIdxs = []int32{
	1, // 0: ChittyChat.GetStudents:input_type -> Empty
	0, // 1: ChittyChat.GetStudents:output_type -> Students
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chittychat_proto_init() }
func file_chittychat_proto_init() {
	if File_chittychat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chittychat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Students); i {
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
		file_chittychat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_chittychat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chittychat_proto_goTypes,
		DependencyIndexes: file_chittychat_proto_depIdxs,
		MessageInfos:      file_chittychat_proto_msgTypes,
	}.Build()
	File_chittychat_proto = out.File
	file_chittychat_proto_rawDesc = nil
	file_chittychat_proto_goTypes = nil
	file_chittychat_proto_depIdxs = nil
}
