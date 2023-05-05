// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tbd/service/v1/project.proto

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

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config       *Configuration          `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	CustomTests  map[string]*CustomTest  `protobuf:"bytes,2,rep,name=custom_tests,json=customTests,proto3" json:"custom_tests,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Seeds        map[string]*Seed        `protobuf:"bytes,3,rep,name=seeds,proto3" json:"seeds,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Models       map[string]*Model       `protobuf:"bytes,4,rep,name=models,proto3" json:"models,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tests        map[string]*Test        `protobuf:"bytes,5,rep,name=tests,proto3" json:"tests,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Sources      map[string]*Source      `protobuf:"bytes,6,rep,name=sources,proto3" json:"sources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ProjectFiles map[string]*ProjectFile `protobuf:"bytes,7,rep,name=project_files,json=projectFiles,proto3" json:"project_files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tbd_service_v1_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_tbd_service_v1_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_tbd_service_v1_project_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetConfig() *Configuration {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Project) GetCustomTests() map[string]*CustomTest {
	if x != nil {
		return x.CustomTests
	}
	return nil
}

func (x *Project) GetSeeds() map[string]*Seed {
	if x != nil {
		return x.Seeds
	}
	return nil
}

func (x *Project) GetModels() map[string]*Model {
	if x != nil {
		return x.Models
	}
	return nil
}

func (x *Project) GetTests() map[string]*Test {
	if x != nil {
		return x.Tests
	}
	return nil
}

func (x *Project) GetSources() map[string]*Source {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *Project) GetProjectFiles() map[string]*ProjectFile {
	if x != nil {
		return x.ProjectFiles
	}
	return nil
}

type CustomTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Sql      string `protobuf:"bytes,3,opt,name=sql,proto3" json:"sql,omitempty"`
}

func (x *CustomTest) Reset() {
	*x = CustomTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tbd_service_v1_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomTest) ProtoMessage() {}

func (x *CustomTest) ProtoReflect() protoreflect.Message {
	mi := &file_tbd_service_v1_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomTest.ProtoReflect.Descriptor instead.
func (*CustomTest) Descriptor() ([]byte, []int) {
	return file_tbd_service_v1_project_proto_rawDescGZIP(), []int{1}
}

func (x *CustomTest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *CustomTest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomTest) GetSql() string {
	if x != nil {
		return x.Sql
	}
	return ""
}

var File_tbd_service_v1_project_proto protoreflect.FileDescriptor

var file_tbd_service_v1_project_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x62, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x74, 0x62, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x22,
	0x74, 0x62, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x74, 0x62, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x74, 0x62, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xce, 0x07, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x35, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x74, 0x62, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x4b, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x62, 0x64,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x65, 0x73, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x65, 0x73, 0x74,
	0x73, 0x12, 0x38, 0x0a, 0x05, 0x73, 0x65, 0x65, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x74, 0x62, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x65, 0x64, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x73, 0x65, 0x65, 0x64, 0x73, 0x12, 0x3b, 0x0a, 0x06, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x62,
	0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x38, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x62, 0x64, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x12, 0x3e, 0x0a, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x62, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x12, 0x4e, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x62, 0x64, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x1a, 0x5a, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x65, 0x73, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x62, 0x64, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4e,
	0x0a, 0x0a, 0x53, 0x65, 0x65, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x74, 0x62, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x65, 0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x50,
	0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x74, 0x62, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x4e, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x74, 0x62, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x52, 0x0a, 0x0c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x74, 0x62, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5c, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x62, 0x64,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x4f, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x71, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x71, 0x6c, 0x42, 0x42, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x6e, 0x66, 0x64, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x62,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74,
	0x62, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tbd_service_v1_project_proto_rawDescOnce sync.Once
	file_tbd_service_v1_project_proto_rawDescData = file_tbd_service_v1_project_proto_rawDesc
)

func file_tbd_service_v1_project_proto_rawDescGZIP() []byte {
	file_tbd_service_v1_project_proto_rawDescOnce.Do(func() {
		file_tbd_service_v1_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_tbd_service_v1_project_proto_rawDescData)
	})
	return file_tbd_service_v1_project_proto_rawDescData
}

var file_tbd_service_v1_project_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tbd_service_v1_project_proto_goTypes = []interface{}{
	(*Project)(nil),       // 0: tbd.service.v1.Project
	(*CustomTest)(nil),    // 1: tbd.service.v1.CustomTest
	nil,                   // 2: tbd.service.v1.Project.CustomTestsEntry
	nil,                   // 3: tbd.service.v1.Project.SeedsEntry
	nil,                   // 4: tbd.service.v1.Project.ModelsEntry
	nil,                   // 5: tbd.service.v1.Project.TestsEntry
	nil,                   // 6: tbd.service.v1.Project.SourcesEntry
	nil,                   // 7: tbd.service.v1.Project.ProjectFilesEntry
	(*Configuration)(nil), // 8: tbd.service.v1.Configuration
	(*Seed)(nil),          // 9: tbd.service.v1.Seed
	(*Model)(nil),         // 10: tbd.service.v1.Model
	(*Test)(nil),          // 11: tbd.service.v1.Test
	(*Source)(nil),        // 12: tbd.service.v1.Source
	(*ProjectFile)(nil),   // 13: tbd.service.v1.ProjectFile
}
var file_tbd_service_v1_project_proto_depIdxs = []int32{
	8,  // 0: tbd.service.v1.Project.config:type_name -> tbd.service.v1.Configuration
	2,  // 1: tbd.service.v1.Project.custom_tests:type_name -> tbd.service.v1.Project.CustomTestsEntry
	3,  // 2: tbd.service.v1.Project.seeds:type_name -> tbd.service.v1.Project.SeedsEntry
	4,  // 3: tbd.service.v1.Project.models:type_name -> tbd.service.v1.Project.ModelsEntry
	5,  // 4: tbd.service.v1.Project.tests:type_name -> tbd.service.v1.Project.TestsEntry
	6,  // 5: tbd.service.v1.Project.sources:type_name -> tbd.service.v1.Project.SourcesEntry
	7,  // 6: tbd.service.v1.Project.project_files:type_name -> tbd.service.v1.Project.ProjectFilesEntry
	1,  // 7: tbd.service.v1.Project.CustomTestsEntry.value:type_name -> tbd.service.v1.CustomTest
	9,  // 8: tbd.service.v1.Project.SeedsEntry.value:type_name -> tbd.service.v1.Seed
	10, // 9: tbd.service.v1.Project.ModelsEntry.value:type_name -> tbd.service.v1.Model
	11, // 10: tbd.service.v1.Project.TestsEntry.value:type_name -> tbd.service.v1.Test
	12, // 11: tbd.service.v1.Project.SourcesEntry.value:type_name -> tbd.service.v1.Source
	13, // 12: tbd.service.v1.Project.ProjectFilesEntry.value:type_name -> tbd.service.v1.ProjectFile
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_tbd_service_v1_project_proto_init() }
func file_tbd_service_v1_project_proto_init() {
	if File_tbd_service_v1_project_proto != nil {
		return
	}
	file_tbd_service_v1_configuration_proto_init()
	file_tbd_service_v1_project_file_proto_init()
	file_tbd_service_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tbd_service_v1_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_tbd_service_v1_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomTest); i {
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
			RawDescriptor: file_tbd_service_v1_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tbd_service_v1_project_proto_goTypes,
		DependencyIndexes: file_tbd_service_v1_project_proto_depIdxs,
		MessageInfos:      file_tbd_service_v1_project_proto_msgTypes,
	}.Build()
	File_tbd_service_v1_project_proto = out.File
	file_tbd_service_v1_project_proto_rawDesc = nil
	file_tbd_service_v1_project_proto_goTypes = nil
	file_tbd_service_v1_project_proto_depIdxs = nil
}
