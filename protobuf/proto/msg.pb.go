// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/msg.proto

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

type Role int32

const (
	Role_Server Role = 0
	Role_Client Role = 1
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "Server",
		1: "Client",
	}
	Role_value = map[string]int32{
		"Server": 0,
		"Client": 1,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_msg_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_proto_msg_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_proto_msg_proto_rawDescGZIP(), []int{0}
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`            // msg context
	Length  int32  `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`             // msg size
	Cnt     int32  `protobuf:"varint,3,opt,name=cnt,proto3" json:"cnt,omitempty"`                   // counter for msg
	Role    Role   `protobuf:"varint,4,opt,name=role,proto3,enum=proto.Role" json:"role,omitempty"` //
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_proto_msg_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserInfo) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *UserInfo) GetCnt() int32 {
	if x != nil {
		return x.Cnt
	}
	return 0
}

func (x *UserInfo) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_Server
}

var File_proto_msg_proto protoreflect.FileDescriptor

var file_proto_msg_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x2a, 0x1e, 0x0a, 0x04, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_msg_proto_rawDescOnce sync.Once
	file_proto_msg_proto_rawDescData = file_proto_msg_proto_rawDesc
)

func file_proto_msg_proto_rawDescGZIP() []byte {
	file_proto_msg_proto_rawDescOnce.Do(func() {
		file_proto_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_msg_proto_rawDescData)
	})
	return file_proto_msg_proto_rawDescData
}

var file_proto_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_msg_proto_goTypes = []interface{}{
	(Role)(0),        // 0: proto.Role
	(*UserInfo)(nil), // 1: proto.UserInfo
}
var file_proto_msg_proto_depIdxs = []int32{
	0, // 0: proto.UserInfo.role:type_name -> proto.Role
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_msg_proto_init() }
func file_proto_msg_proto_init() {
	if File_proto_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
			RawDescriptor: file_proto_msg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_msg_proto_goTypes,
		DependencyIndexes: file_proto_msg_proto_depIdxs,
		EnumInfos:         file_proto_msg_proto_enumTypes,
		MessageInfos:      file_proto_msg_proto_msgTypes,
	}.Build()
	File_proto_msg_proto = out.File
	file_proto_msg_proto_rawDesc = nil
	file_proto_msg_proto_goTypes = nil
	file_proto_msg_proto_depIdxs = nil
}
