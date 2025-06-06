// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0
// source: videoGoodsClassifiy.proto

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

type GoodsClassReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pid           int32                  `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GoodsClassReq) Reset() {
	*x = GoodsClassReq{}
	mi := &file_videoGoodsClassifiy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GoodsClassReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsClassReq) ProtoMessage() {}

func (x *GoodsClassReq) ProtoReflect() protoreflect.Message {
	mi := &file_videoGoodsClassifiy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsClassReq.ProtoReflect.Descriptor instead.
func (*GoodsClassReq) Descriptor() ([]byte, []int) {
	return file_videoGoodsClassifiy_proto_rawDescGZIP(), []int{0}
}

func (x *GoodsClassReq) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type GoodsClassRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Goodsclass    []*GoodsClass          `protobuf:"bytes,1,rep,name=goodsclass,proto3" json:"goodsclass,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GoodsClassRes) Reset() {
	*x = GoodsClassRes{}
	mi := &file_videoGoodsClassifiy_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GoodsClassRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsClassRes) ProtoMessage() {}

func (x *GoodsClassRes) ProtoReflect() protoreflect.Message {
	mi := &file_videoGoodsClassifiy_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsClassRes.ProtoReflect.Descriptor instead.
func (*GoodsClassRes) Descriptor() ([]byte, []int) {
	return file_videoGoodsClassifiy_proto_rawDescGZIP(), []int{1}
}

func (x *GoodsClassRes) GetGoodsclass() []*GoodsClass {
	if x != nil {
		return x.Goodsclass
	}
	return nil
}

type GoodsClass struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ClassifiyName string                 `protobuf:"bytes,2,opt,name=ClassifiyName,proto3" json:"ClassifiyName,omitempty"`
	Pid           int32                  `protobuf:"varint,3,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Soft          int32                  `protobuf:"varint,4,opt,name=Soft,proto3" json:"Soft,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GoodsClass) Reset() {
	*x = GoodsClass{}
	mi := &file_videoGoodsClassifiy_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GoodsClass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsClass) ProtoMessage() {}

func (x *GoodsClass) ProtoReflect() protoreflect.Message {
	mi := &file_videoGoodsClassifiy_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsClass.ProtoReflect.Descriptor instead.
func (*GoodsClass) Descriptor() ([]byte, []int) {
	return file_videoGoodsClassifiy_proto_rawDescGZIP(), []int{2}
}

func (x *GoodsClass) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GoodsClass) GetClassifiyName() string {
	if x != nil {
		return x.ClassifiyName
	}
	return ""
}

func (x *GoodsClass) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *GoodsClass) GetSoft() int32 {
	if x != nil {
		return x.Soft
	}
	return 0
}

var File_videoGoodsClassifiy_proto protoreflect.FileDescriptor

const file_videoGoodsClassifiy_proto_rawDesc = "" +
	"\n" +
	"\x19videoGoodsClassifiy.proto\"!\n" +
	"\rGoodsClassReq\x12\x10\n" +
	"\x03Pid\x18\x01 \x01(\x05R\x03Pid\"<\n" +
	"\rGoodsClassRes\x12+\n" +
	"\n" +
	"goodsclass\x18\x01 \x03(\v2\v.GoodsClassR\n" +
	"goodsclass\"h\n" +
	"\n" +
	"GoodsClass\x12\x0e\n" +
	"\x02Id\x18\x01 \x01(\x05R\x02Id\x12$\n" +
	"\rClassifiyName\x18\x02 \x01(\tR\rClassifiyName\x12\x10\n" +
	"\x03Pid\x18\x03 \x01(\x05R\x03Pid\x12\x12\n" +
	"\x04Soft\x18\x04 \x01(\x05R\x04Soft2A\n" +
	"\x0egoodsClassifiy\x12/\n" +
	"\rGetGoodsClass\x12\x0e.GoodsClassReq\x1a\x0e.GoodsClassResB\x03Z\x01.b\x06proto3"

var (
	file_videoGoodsClassifiy_proto_rawDescOnce sync.Once
	file_videoGoodsClassifiy_proto_rawDescData []byte
)

func file_videoGoodsClassifiy_proto_rawDescGZIP() []byte {
	file_videoGoodsClassifiy_proto_rawDescOnce.Do(func() {
		file_videoGoodsClassifiy_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_videoGoodsClassifiy_proto_rawDesc), len(file_videoGoodsClassifiy_proto_rawDesc)))
	})
	return file_videoGoodsClassifiy_proto_rawDescData
}

var file_videoGoodsClassifiy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_videoGoodsClassifiy_proto_goTypes = []any{
	(*GoodsClassReq)(nil), // 0: GoodsClassReq
	(*GoodsClassRes)(nil), // 1: GoodsClassRes
	(*GoodsClass)(nil),    // 2: GoodsClass
}
var file_videoGoodsClassifiy_proto_depIdxs = []int32{
	2, // 0: GoodsClassRes.goodsclass:type_name -> GoodsClass
	0, // 1: goodsClassifiy.GetGoodsClass:input_type -> GoodsClassReq
	1, // 2: goodsClassifiy.GetGoodsClass:output_type -> GoodsClassRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_videoGoodsClassifiy_proto_init() }
func file_videoGoodsClassifiy_proto_init() {
	if File_videoGoodsClassifiy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_videoGoodsClassifiy_proto_rawDesc), len(file_videoGoodsClassifiy_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_videoGoodsClassifiy_proto_goTypes,
		DependencyIndexes: file_videoGoodsClassifiy_proto_depIdxs,
		MessageInfos:      file_videoGoodsClassifiy_proto_msgTypes,
	}.Build()
	File_videoGoodsClassifiy_proto = out.File
	file_videoGoodsClassifiy_proto_goTypes = nil
	file_videoGoodsClassifiy_proto_depIdxs = nil
}
