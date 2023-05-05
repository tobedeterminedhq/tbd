// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tbd/service/v1/connection_config.proto

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

// Configuration describes the configuration of the project.
type ConnectionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to Config:
	//	*ConnectionConfig_Sqlite
	//	*ConnectionConfig_SqliteInMemory
	//	*ConnectionConfig_Duckdb
	//	*ConnectionConfig_Postgres
	//	*ConnectionConfig_Mysql
	//	*ConnectionConfig_BigQuery
	Config isConnectionConfig_Config `protobuf_oneof:"config"`
}

func (x *ConnectionConfig) Reset() {
	*x = ConnectionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tbd_service_v1_connection_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionConfig) ProtoMessage() {}

func (x *ConnectionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tbd_service_v1_connection_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionConfig.ProtoReflect.Descriptor instead.
func (*ConnectionConfig) Descriptor() ([]byte, []int) {
	return file_tbd_service_v1_connection_config_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *ConnectionConfig) GetConfig() isConnectionConfig_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *ConnectionConfig) GetSqlite() *ConnectionConfig_ConnectionConfigSqLite {
	if x, ok := x.GetConfig().(*ConnectionConfig_Sqlite); ok {
		return x.Sqlite
	}
	return nil
}

func (x *ConnectionConfig) GetSqliteInMemory() *ConnectionConfig_ConnectionConfigSqLiteInMemory {
	if x, ok := x.GetConfig().(*ConnectionConfig_SqliteInMemory); ok {
		return x.SqliteInMemory
	}
	return nil
}

func (x *ConnectionConfig) GetDuckdb() *ConnectionConfig_ConnectionConfigDuckDB {
	if x, ok := x.GetConfig().(*ConnectionConfig_Duckdb); ok {
		return x.Duckdb
	}
	return nil
}

func (x *ConnectionConfig) GetPostgres() *ConnectionConfig_ConnectionConfigPostgres {
	if x, ok := x.GetConfig().(*ConnectionConfig_Postgres); ok {
		return x.Postgres
	}
	return nil
}

func (x *ConnectionConfig) GetMysql() *ConnectionConfig_ConnectionConfigMySql {
	if x, ok := x.GetConfig().(*ConnectionConfig_Mysql); ok {
		return x.Mysql
	}
	return nil
}

func (x *ConnectionConfig) GetBigQuery() *ConnectionConfig_ConnectionConfigBigQuery {
	if x, ok := x.GetConfig().(*ConnectionConfig_BigQuery); ok {
		return x.BigQuery
	}
	return nil
}

type isConnectionConfig_Config interface {
	isConnectionConfig_Config()
}

type ConnectionConfig_Sqlite struct {
	Sqlite *ConnectionConfig_ConnectionConfigSqLite `protobuf:"bytes,2,opt,name=sqlite,proto3,oneof"`
}

type ConnectionConfig_SqliteInMemory struct {
	SqliteInMemory *ConnectionConfig_ConnectionConfigSqLiteInMemory `protobuf:"bytes,3,opt,name=sqlite_in_memory,json=sqliteInMemory,proto3,oneof"`
}

type ConnectionConfig_Duckdb struct {
	Duckdb *ConnectionConfig_ConnectionConfigDuckDB `protobuf:"bytes,4,opt,name=duckdb,proto3,oneof"`
}

type ConnectionConfig_Postgres struct {
	Postgres *ConnectionConfig_ConnectionConfigPostgres `protobuf:"bytes,5,opt,name=postgres,proto3,oneof"`
}

type ConnectionConfig_Mysql struct {
	Mysql *ConnectionConfig_ConnectionConfigMySql `protobuf:"bytes,6,opt,name=mysql,proto3,oneof"`
}

type ConnectionConfig_BigQuery struct {
	BigQuery *ConnectionConfig_ConnectionConfigBigQuery `protobuf:"bytes,7,opt,name=big_query,json=bigQuery,proto3,oneof"`
}

func (*ConnectionConfig_Sqlite) isConnectionConfig_Config() {}

func (*ConnectionConfig_SqliteInMemory) isConnectionConfig_Config() {}

func (*ConnectionConfig_Duckdb) isConnectionConfig_Config() {}

func (*ConnectionConfig_Postgres) isConnectionConfig_Config() {}

func (*ConnectionConfig_Mysql) isConnectionConfig_Config() {}

func (*ConnectionConfig_BigQuery) isConnectionConfig_Config() {}

type ConnectionConfig_ConnectionConfigSqLite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ConnectionConfig_ConnectionConfigSqLite) Reset() {
	*x = ConnectionConfig_ConnectionConfigSqLite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tbd_service_v1_connection_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionConfig_ConnectionConfigSqLite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionConfig_ConnectionConfigSqLite) ProtoMessage() {}

func (x *ConnectionConfig_ConnectionConfigSqLite) ProtoReflect() protoreflect.Message {
	mi := &file_tbd_service_v1_connection_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionConfig_ConnectionConfigSqLite.ProtoReflect.Descriptor instead.
func (*ConnectionConfig_ConnectionConfigSqLite) Descriptor() ([]byte, []int) {
	return file_tbd_service_v1_connection_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ConnectionConfig_ConnectionConfigSqLite) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ConnectionConfig_ConnectionConfigSqLiteInMemory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectionConfig_ConnectionConfigSqLiteInMemory) Reset() {
	*x = ConnectionConfig_ConnectionConfigSqLiteInMemory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tbd_service_v1_connection_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionConfig_ConnectionConfigSqLiteInMemory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionConfig_ConnectionConfigSqLiteInMemory) ProtoMessage() {}

func (x *ConnectionConfig_ConnectionConfigSqLiteInMemory) ProtoReflect() protoreflect.Message {
	mi := &file_tbd_service_v1_connection_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionConfig_ConnectionConfigSqLiteInMemory.ProtoReflect.Descriptor instead.
func (*ConnectionConfig_ConnectionConfigSqLiteInMemory) Descriptor() ([]byte, []int) {
	return file_tbd_service_v1_connection_config_proto_rawDescGZIP(), []int{0, 1}
}

type ConnectionConfig_ConnectionConfigDuckDB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// params are configuration options for the database. The options are available to see here: https://duckdb.org/docs/sql/configuration.
	Params map[string]string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConnectionConfig_ConnectionConfigDuckDB) Reset() {
	*x = ConnectionConfig_ConnectionConfigDuckDB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tbd_service_v1_connection_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionConfig_ConnectionConfigDuckDB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionConfig_ConnectionConfigDuckDB) ProtoMessage() {}

func (x *ConnectionConfig_ConnectionConfigDuckDB) ProtoReflect() protoreflect.Message {
	mi := &file_tbd_service_v1_connection_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionConfig_ConnectionConfigDuckDB.ProtoReflect.Descriptor instead.
func (*ConnectionConfig_ConnectionConfigDuckDB) Descriptor() ([]byte, []int) {
	return file_tbd_service_v1_connection_config_proto_rawDescGZIP(), []int{0, 2}
}

func (x *ConnectionConfig_ConnectionConfigDuckDB) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ConnectionConfig_ConnectionConfigDuckDB) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type ConnectionConfig_ConnectionConfigPostgres struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host     string            `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port     string            `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	User     string            `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Password string            `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Database string            `protobuf:"bytes,5,opt,name=database,proto3" json:"database,omitempty"`
	Params   map[string]string `protobuf:"bytes,6,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConnectionConfig_ConnectionConfigPostgres) Reset() {
	*x = ConnectionConfig_ConnectionConfigPostgres{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tbd_service_v1_connection_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionConfig_ConnectionConfigPostgres) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionConfig_ConnectionConfigPostgres) ProtoMessage() {}

func (x *ConnectionConfig_ConnectionConfigPostgres) ProtoReflect() protoreflect.Message {
	mi := &file_tbd_service_v1_connection_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionConfig_ConnectionConfigPostgres.ProtoReflect.Descriptor instead.
func (*ConnectionConfig_ConnectionConfigPostgres) Descriptor() ([]byte, []int) {
	return file_tbd_service_v1_connection_config_proto_rawDescGZIP(), []int{0, 3}
}

func (x *ConnectionConfig_ConnectionConfigPostgres) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ConnectionConfig_ConnectionConfigPostgres) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *ConnectionConfig_ConnectionConfigPostgres) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ConnectionConfig_ConnectionConfigPostgres) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ConnectionConfig_ConnectionConfigPostgres) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *ConnectionConfig_ConnectionConfigPostgres) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type ConnectionConfig_ConnectionConfigMySql struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string            `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string            `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Protocol string            `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Host     string            `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	Port     string            `protobuf:"bytes,5,opt,name=port,proto3" json:"port,omitempty"`
	Database string            `protobuf:"bytes,6,opt,name=database,proto3" json:"database,omitempty"`
	Params   map[string]string `protobuf:"bytes,7,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConnectionConfig_ConnectionConfigMySql) Reset() {
	*x = ConnectionConfig_ConnectionConfigMySql{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tbd_service_v1_connection_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionConfig_ConnectionConfigMySql) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionConfig_ConnectionConfigMySql) ProtoMessage() {}

func (x *ConnectionConfig_ConnectionConfigMySql) ProtoReflect() protoreflect.Message {
	mi := &file_tbd_service_v1_connection_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionConfig_ConnectionConfigMySql.ProtoReflect.Descriptor instead.
func (*ConnectionConfig_ConnectionConfigMySql) Descriptor() ([]byte, []int) {
	return file_tbd_service_v1_connection_config_proto_rawDescGZIP(), []int{0, 4}
}

func (x *ConnectionConfig_ConnectionConfigMySql) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ConnectionConfig_ConnectionConfigMySql) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ConnectionConfig_ConnectionConfigMySql) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ConnectionConfig_ConnectionConfigMySql) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ConnectionConfig_ConnectionConfigMySql) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *ConnectionConfig_ConnectionConfigMySql) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *ConnectionConfig_ConnectionConfigMySql) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type ConnectionConfig_ConnectionConfigBigQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	DatasetId string `protobuf:"bytes,2,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
}

func (x *ConnectionConfig_ConnectionConfigBigQuery) Reset() {
	*x = ConnectionConfig_ConnectionConfigBigQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tbd_service_v1_connection_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionConfig_ConnectionConfigBigQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionConfig_ConnectionConfigBigQuery) ProtoMessage() {}

func (x *ConnectionConfig_ConnectionConfigBigQuery) ProtoReflect() protoreflect.Message {
	mi := &file_tbd_service_v1_connection_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionConfig_ConnectionConfigBigQuery.ProtoReflect.Descriptor instead.
func (*ConnectionConfig_ConnectionConfigBigQuery) Descriptor() ([]byte, []int) {
	return file_tbd_service_v1_connection_config_proto_rawDescGZIP(), []int{0, 5}
}

func (x *ConnectionConfig_ConnectionConfigBigQuery) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ConnectionConfig_ConnectionConfigBigQuery) GetDatasetId() string {
	if x != nil {
		return x.DatasetId
	}
	return ""
}

var File_tbd_service_v1_connection_config_proto protoreflect.FileDescriptor

var file_tbd_service_v1_connection_config_proto_rawDesc = []byte{
	0x0a, 0x26, 0x74, 0x62, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x62, 0x64, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xab, 0x0c, 0x0a, 0x10, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x51, 0x0a, 0x06, 0x73, 0x71, 0x6c, 0x69, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x74, 0x62, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x53, 0x71, 0x4c, 0x69, 0x74, 0x65, 0x48, 0x00, 0x52, 0x06, 0x73, 0x71,
	0x6c, 0x69, 0x74, 0x65, 0x12, 0x6b, 0x0a, 0x10, 0x73, 0x71, 0x6c, 0x69, 0x74, 0x65, 0x5f, 0x69,
	0x6e, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f,
	0x2e, 0x74, 0x62, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x71, 0x4c, 0x69, 0x74, 0x65, 0x49, 0x6e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x48,
	0x00, 0x52, 0x0e, 0x73, 0x71, 0x6c, 0x69, 0x74, 0x65, 0x49, 0x6e, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x12, 0x51, 0x0a, 0x06, 0x64, 0x75, 0x63, 0x6b, 0x64, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x74, 0x62, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x44, 0x75, 0x63, 0x6b, 0x44, 0x42, 0x48, 0x00, 0x52, 0x06, 0x64, 0x75,
	0x63, 0x6b, 0x64, 0x62, 0x12, 0x57, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x74, 0x62, 0x64, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65,
	0x73, 0x48, 0x00, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x12, 0x4e, 0x0a,
	0x05, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x74,
	0x62, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d,
	0x79, 0x53, 0x71, 0x6c, 0x48, 0x00, 0x52, 0x05, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x12, 0x58, 0x0a,
	0x09, 0x62, 0x69, 0x67, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x39, 0x2e, 0x74, 0x62, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x08, 0x62,
	0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x2c, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x71, 0x4c, 0x69, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x1a, 0x20, 0x0a, 0x1e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x71, 0x4c, 0x69, 0x74, 0x65, 0x49,
	0x6e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x1a, 0xc4, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x75, 0x63, 0x6b,
	0x44, 0x42, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x5b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x74, 0x62, 0x64, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x75, 0x63, 0x6b, 0x44, 0x42, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xa8,
	0x02, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x5d, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x45, 0x2e, 0x74, 0x62, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39,
	0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xc6, 0x02, 0x0a, 0x15, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x79,
	0x53, 0x71, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x74, 0x62,
	0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x79,
	0x53, 0x71, 0x6c, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x58, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x64, 0x42, 0x08, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x42, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x6e, 0x66, 0x64, 0x6b, 0x69, 0x6e, 0x67,
	0x2f, 0x74, 0x62, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x74, 0x62, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tbd_service_v1_connection_config_proto_rawDescOnce sync.Once
	file_tbd_service_v1_connection_config_proto_rawDescData = file_tbd_service_v1_connection_config_proto_rawDesc
)

func file_tbd_service_v1_connection_config_proto_rawDescGZIP() []byte {
	file_tbd_service_v1_connection_config_proto_rawDescOnce.Do(func() {
		file_tbd_service_v1_connection_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_tbd_service_v1_connection_config_proto_rawDescData)
	})
	return file_tbd_service_v1_connection_config_proto_rawDescData
}

var file_tbd_service_v1_connection_config_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_tbd_service_v1_connection_config_proto_goTypes = []interface{}{
	(*ConnectionConfig)(nil),                                // 0: tbd.service.v1.ConnectionConfig
	(*ConnectionConfig_ConnectionConfigSqLite)(nil),         // 1: tbd.service.v1.ConnectionConfig.ConnectionConfigSqLite
	(*ConnectionConfig_ConnectionConfigSqLiteInMemory)(nil), // 2: tbd.service.v1.ConnectionConfig.ConnectionConfigSqLiteInMemory
	(*ConnectionConfig_ConnectionConfigDuckDB)(nil),         // 3: tbd.service.v1.ConnectionConfig.ConnectionConfigDuckDB
	(*ConnectionConfig_ConnectionConfigPostgres)(nil),       // 4: tbd.service.v1.ConnectionConfig.ConnectionConfigPostgres
	(*ConnectionConfig_ConnectionConfigMySql)(nil),          // 5: tbd.service.v1.ConnectionConfig.ConnectionConfigMySql
	(*ConnectionConfig_ConnectionConfigBigQuery)(nil),       // 6: tbd.service.v1.ConnectionConfig.ConnectionConfigBigQuery
	nil, // 7: tbd.service.v1.ConnectionConfig.ConnectionConfigDuckDB.ParamsEntry
	nil, // 8: tbd.service.v1.ConnectionConfig.ConnectionConfigPostgres.ParamsEntry
	nil, // 9: tbd.service.v1.ConnectionConfig.ConnectionConfigMySql.ParamsEntry
}
var file_tbd_service_v1_connection_config_proto_depIdxs = []int32{
	1, // 0: tbd.service.v1.ConnectionConfig.sqlite:type_name -> tbd.service.v1.ConnectionConfig.ConnectionConfigSqLite
	2, // 1: tbd.service.v1.ConnectionConfig.sqlite_in_memory:type_name -> tbd.service.v1.ConnectionConfig.ConnectionConfigSqLiteInMemory
	3, // 2: tbd.service.v1.ConnectionConfig.duckdb:type_name -> tbd.service.v1.ConnectionConfig.ConnectionConfigDuckDB
	4, // 3: tbd.service.v1.ConnectionConfig.postgres:type_name -> tbd.service.v1.ConnectionConfig.ConnectionConfigPostgres
	5, // 4: tbd.service.v1.ConnectionConfig.mysql:type_name -> tbd.service.v1.ConnectionConfig.ConnectionConfigMySql
	6, // 5: tbd.service.v1.ConnectionConfig.big_query:type_name -> tbd.service.v1.ConnectionConfig.ConnectionConfigBigQuery
	7, // 6: tbd.service.v1.ConnectionConfig.ConnectionConfigDuckDB.params:type_name -> tbd.service.v1.ConnectionConfig.ConnectionConfigDuckDB.ParamsEntry
	8, // 7: tbd.service.v1.ConnectionConfig.ConnectionConfigPostgres.params:type_name -> tbd.service.v1.ConnectionConfig.ConnectionConfigPostgres.ParamsEntry
	9, // 8: tbd.service.v1.ConnectionConfig.ConnectionConfigMySql.params:type_name -> tbd.service.v1.ConnectionConfig.ConnectionConfigMySql.ParamsEntry
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_tbd_service_v1_connection_config_proto_init() }
func file_tbd_service_v1_connection_config_proto_init() {
	if File_tbd_service_v1_connection_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tbd_service_v1_connection_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionConfig); i {
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
		file_tbd_service_v1_connection_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionConfig_ConnectionConfigSqLite); i {
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
		file_tbd_service_v1_connection_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionConfig_ConnectionConfigSqLiteInMemory); i {
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
		file_tbd_service_v1_connection_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionConfig_ConnectionConfigDuckDB); i {
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
		file_tbd_service_v1_connection_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionConfig_ConnectionConfigPostgres); i {
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
		file_tbd_service_v1_connection_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionConfig_ConnectionConfigMySql); i {
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
		file_tbd_service_v1_connection_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionConfig_ConnectionConfigBigQuery); i {
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
	file_tbd_service_v1_connection_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ConnectionConfig_Sqlite)(nil),
		(*ConnectionConfig_SqliteInMemory)(nil),
		(*ConnectionConfig_Duckdb)(nil),
		(*ConnectionConfig_Postgres)(nil),
		(*ConnectionConfig_Mysql)(nil),
		(*ConnectionConfig_BigQuery)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tbd_service_v1_connection_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tbd_service_v1_connection_config_proto_goTypes,
		DependencyIndexes: file_tbd_service_v1_connection_config_proto_depIdxs,
		MessageInfos:      file_tbd_service_v1_connection_config_proto_msgTypes,
	}.Build()
	File_tbd_service_v1_connection_config_proto = out.File
	file_tbd_service_v1_connection_config_proto_rawDesc = nil
	file_tbd_service_v1_connection_config_proto_goTypes = nil
	file_tbd_service_v1_connection_config_proto_depIdxs = nil
}
