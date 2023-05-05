// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tbd/service/v1/seeds.proto

package servicev1

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

type SeedsSQL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sql []string `protobuf:"bytes,1,rep,name=sql,proto3" json:"sql,omitempty"`
}

func (x *SeedsSQL) Reset() {
	*x = SeedsSQL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tbd_service_v1_seeds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeedsSQL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeedsSQL) ProtoMessage() {}

func (x *SeedsSQL) ProtoReflect() protoreflect.Message {
	mi := &file_tbd_service_v1_seeds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeedsSQL.ProtoReflect.Descriptor instead.
func (*SeedsSQL) Descriptor() ([]byte, []int) {
	return file_tbd_service_v1_seeds_proto_rawDescGZIP(), []int{0}
}

func (x *SeedsSQL) GetSql() []string {
	if x != nil {
		return x.Sql
	}
	return nil
}

var File_tbd_service_v1_seeds_proto protoreflect.FileDescriptor

var file_tbd_service_v1_seeds_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x62, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x62,
	0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x1c, 0x0a, 0x08,
	0x53, 0x65, 0x65, 0x64, 0x73, 0x53, 0x51, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x71, 0x6c, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x73, 0x71, 0x6c, 0x42, 0x42, 0x50, 0x01, 0x5a, 0x3e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x6e, 0x66, 0x64,
	0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x62, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x62, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tbd_service_v1_seeds_proto_rawDescOnce sync.Once
	file_tbd_service_v1_seeds_proto_rawDescData = file_tbd_service_v1_seeds_proto_rawDesc
)

func file_tbd_service_v1_seeds_proto_rawDescGZIP() []byte {
	file_tbd_service_v1_seeds_proto_rawDescOnce.Do(func() {
		file_tbd_service_v1_seeds_proto_rawDescData = protoimpl.X.CompressGZIP(file_tbd_service_v1_seeds_proto_rawDescData)
	})
	return file_tbd_service_v1_seeds_proto_rawDescData
}

var file_tbd_service_v1_seeds_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tbd_service_v1_seeds_proto_goTypes = []interface{}{
	(*SeedsSQL)(nil), // 0: tbd.service.v1.SeedsSQL
}
var file_tbd_service_v1_seeds_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tbd_service_v1_seeds_proto_init() }
func file_tbd_service_v1_seeds_proto_init() {
	if File_tbd_service_v1_seeds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tbd_service_v1_seeds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeedsSQL); i {
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
			RawDescriptor: file_tbd_service_v1_seeds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tbd_service_v1_seeds_proto_goTypes,
		DependencyIndexes: file_tbd_service_v1_seeds_proto_depIdxs,
		MessageInfos:      file_tbd_service_v1_seeds_proto_msgTypes,
	}.Build()
	File_tbd_service_v1_seeds_proto = out.File
	file_tbd_service_v1_seeds_proto_rawDesc = nil
	file_tbd_service_v1_seeds_proto_goTypes = nil
	file_tbd_service_v1_seeds_proto_depIdxs = nil
}
