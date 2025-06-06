// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0
// source: videoUser.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginRep struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	UserCode      string                 `protobuf:"bytes,2,opt,name=UserCode,proto3" json:"UserCode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRep) Reset() {
	*x = LoginRep{}
	mi := &file_videoUser_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRep) ProtoMessage() {}

func (x *LoginRep) ProtoReflect() protoreflect.Message {
	mi := &file_videoUser_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRep.ProtoReflect.Descriptor instead.
func (*LoginRep) Descriptor() ([]byte, []int) {
	return file_videoUser_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRep) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginRep) GetUserCode() string {
	if x != nil {
		return x.UserCode
	}
	return ""
}

type LoginRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRes) Reset() {
	*x = LoginRes{}
	mi := &file_videoUser_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRes) ProtoMessage() {}

func (x *LoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_videoUser_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRes.ProtoReflect.Descriptor instead.
func (*LoginRes) Descriptor() ([]byte, []int) {
	return file_videoUser_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRes) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_videoUser_proto protoreflect.FileDescriptor

const file_videoUser_proto_rawDesc = "" +
	"\n" +
	"\x0fvideoUser.proto\":\n" +
	"\bLoginRep\x12\x12\n" +
	"\x04Name\x18\x01 \x01(\tR\x04Name\x12\x1a\n" +
	"\bUserCode\x18\x02 \x01(\tR\bUserCode\"\x1a\n" +
	"\bLoginRes\x12\x0e\n" +
	"\x02Id\x18\x01 \x01(\x05R\x02Id2*\n" +
	"\tVideoUser\x12\x1d\n" +
	"\x05Login\x12\t.LoginRep\x1a\t.LoginResB\x03Z\x01.b\x06proto3"

var (
	file_videoUser_proto_rawDescOnce sync.Once
	file_videoUser_proto_rawDescData []byte
)

func file_videoUser_proto_rawDescGZIP() []byte {
	file_videoUser_proto_rawDescOnce.Do(func() {
		file_videoUser_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_videoUser_proto_rawDesc), len(file_videoUser_proto_rawDesc)))
	})
	return file_videoUser_proto_rawDescData
}

var file_videoUser_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_videoUser_proto_goTypes = []any{
	(*LoginRep)(nil), // 0: LoginRep
	(*LoginRes)(nil), // 1: LoginRes
}
var file_videoUser_proto_depIdxs = []int32{
	0, // 0: VideoUser.Login:input_type -> LoginRep
	1, // 1: VideoUser.Login:output_type -> LoginRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_videoUser_proto_init() }
func file_videoUser_proto_init() {
	if File_videoUser_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_videoUser_proto_rawDesc), len(file_videoUser_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_videoUser_proto_goTypes,
		DependencyIndexes: file_videoUser_proto_depIdxs,
		MessageInfos:      file_videoUser_proto_msgTypes,
	}.Build()
	File_videoUser_proto = out.File
	file_videoUser_proto_goTypes = nil
	file_videoUser_proto_depIdxs = nil
}
