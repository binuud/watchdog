/// Please use the following editor setup for this file:

// Tab size=2; Tabs as spaces; Clean up trailing whitepsace
// 'make proto' will run clang-format to fix formatiing

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: v1/watchdog/watchdogService.proto

package watchdog

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type HealthResponse_HealthStatus int32

const (
	HealthResponse_HealthStatusIgnore HealthResponse_HealthStatus = 0
	HealthResponse_Active             HealthResponse_HealthStatus = 1
	HealthResponse_Error              HealthResponse_HealthStatus = 2
)

// Enum value maps for HealthResponse_HealthStatus.
var (
	HealthResponse_HealthStatus_name = map[int32]string{
		0: "HealthStatusIgnore",
		1: "Active",
		2: "Error",
	}
	HealthResponse_HealthStatus_value = map[string]int32{
		"HealthStatusIgnore": 0,
		"Active":             1,
		"Error":              2,
	}
)

func (x HealthResponse_HealthStatus) Enum() *HealthResponse_HealthStatus {
	p := new(HealthResponse_HealthStatus)
	*p = x
	return p
}

func (x HealthResponse_HealthStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthResponse_HealthStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_watchdog_watchdogService_proto_enumTypes[0].Descriptor()
}

func (HealthResponse_HealthStatus) Type() protoreflect.EnumType {
	return &file_v1_watchdog_watchdogService_proto_enumTypes[0]
}

func (x HealthResponse_HealthStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthResponse_HealthStatus.Descriptor instead.
func (HealthResponse_HealthStatus) EnumDescriptor() ([]byte, []int) {
	return file_v1_watchdog_watchdogService_proto_rawDescGZIP(), []int{9, 0}
}

type ReloadRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReloadRequest) Reset() {
	*x = ReloadRequest{}
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReloadRequest) ProtoMessage() {}

func (x *ReloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReloadRequest.ProtoReflect.Descriptor instead.
func (*ReloadRequest) Descriptor() ([]byte, []int) {
	return file_v1_watchdog_watchdogService_proto_rawDescGZIP(), []int{0}
}

type ReloadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReloadResponse) Reset() {
	*x = ReloadResponse{}
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReloadResponse) ProtoMessage() {}

func (x *ReloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReloadResponse.ProtoReflect.Descriptor instead.
func (*ReloadResponse) Descriptor() ([]byte, []int) {
	return file_v1_watchdog_watchdogService_proto_rawDescGZIP(), []int{1}
}

type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_v1_watchdog_watchdogService_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Summary       *DomainSummary         `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_v1_watchdog_watchdogService_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetSummary() *DomainSummary {
	if x != nil {
		return x.Summary
	}
	return nil
}

type GetDetailsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDetailsRequest) Reset() {
	*x = GetDetailsRequest{}
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDetailsRequest) ProtoMessage() {}

func (x *GetDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetDetailsRequest) Descriptor() ([]byte, []int) {
	return file_v1_watchdog_watchdogService_proto_rawDescGZIP(), []int{4}
}

func (x *GetDetailsRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetDetailsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetDetailsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Domain        *DomainRow             `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDetailsResponse) Reset() {
	*x = GetDetailsResponse{}
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDetailsResponse) ProtoMessage() {}

func (x *GetDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetDetailsResponse) Descriptor() ([]byte, []int) {
	return file_v1_watchdog_watchdogService_proto_rawDescGZIP(), []int{5}
}

func (x *GetDetailsResponse) GetDomain() *DomainRow {
	if x != nil {
		return x.Domain
	}
	return nil
}

type ListSummariesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// current page, page starts from 1
	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	// number of items per page
	PerPage       int64 `protobuf:"varint,2,opt,name=perPage,proto3" json:"perPage,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSummariesRequest) Reset() {
	*x = ListSummariesRequest{}
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSummariesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSummariesRequest) ProtoMessage() {}

func (x *ListSummariesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSummariesRequest.ProtoReflect.Descriptor instead.
func (*ListSummariesRequest) Descriptor() ([]byte, []int) {
	return file_v1_watchdog_watchdogService_proto_rawDescGZIP(), []int{6}
}

func (x *ListSummariesRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListSummariesRequest) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type ListSummariesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int64                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage       int64                  `protobuf:"varint,2,opt,name=perPage,proto3" json:"perPage,omitempty"`
	Summaries     []*DomainSummary       `protobuf:"bytes,3,rep,name=summaries,proto3" json:"summaries,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSummariesResponse) Reset() {
	*x = ListSummariesResponse{}
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSummariesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSummariesResponse) ProtoMessage() {}

func (x *ListSummariesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSummariesResponse.ProtoReflect.Descriptor instead.
func (*ListSummariesResponse) Descriptor() ([]byte, []int) {
	return file_v1_watchdog_watchdogService_proto_rawDescGZIP(), []int{7}
}

func (x *ListSummariesResponse) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListSummariesResponse) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *ListSummariesResponse) GetSummaries() []*DomainSummary {
	if x != nil {
		return x.Summaries
	}
	return nil
}

type HealthRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_v1_watchdog_watchdogService_proto_rawDescGZIP(), []int{8}
}

type HealthResponse struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	CreatedAt int64                  `protobuf:"varint,1,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// number of domains being watched
	NumDomains    int64                       `protobuf:"varint,2,opt,name=numDomains,proto3" json:"numDomains,omitempty"`
	Status        HealthResponse_HealthStatus `protobuf:"varint,3,opt,name=status,proto3,enum=watchdog.HealthResponse_HealthStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_watchdog_watchdogService_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_v1_watchdog_watchdogService_proto_rawDescGZIP(), []int{9}
}

func (x *HealthResponse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *HealthResponse) GetNumDomains() int64 {
	if x != nil {
		return x.NumDomains
	}
	return 0
}

func (x *HealthResponse) GetStatus() HealthResponse_HealthStatus {
	if x != nil {
		return x.Status
	}
	return HealthResponse_HealthStatusIgnore
}

var File_v1_watchdog_watchdogService_proto protoreflect.FileDescriptor

var file_v1_watchdog_watchdogService_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x76, 0x31, 0x2f,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x40, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x22, 0x3b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x41, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x6f, 0x77, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x22, 0x44, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x7c, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x35, 0x0a, 0x09, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x09, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xcc, 0x01, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x75,
	0x6d, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x64, 0x6f, 0x67, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3d, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0x02, 0x32, 0xe8, 0x03, 0x0a, 0x08, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x44, 0x6f, 0x67, 0x12, 0x5b, 0x0a, 0x06, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x17, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f,
	0x67, 0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31,
	0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x4c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64,
	0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76,
	0x31, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x68,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1b, 0x2e, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x64, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12,
	0x17, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x67, 0x65,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x6d, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x64, 0x6f, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x64, 0x6f, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2f, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x58, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x12, 0x17, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76,
	0x31, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x69, 0x6e, 0x75, 0x75, 0x64, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64,
	0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_v1_watchdog_watchdogService_proto_rawDescOnce sync.Once
	file_v1_watchdog_watchdogService_proto_rawDescData []byte
)

func file_v1_watchdog_watchdogService_proto_rawDescGZIP() []byte {
	file_v1_watchdog_watchdogService_proto_rawDescOnce.Do(func() {
		file_v1_watchdog_watchdogService_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_watchdog_watchdogService_proto_rawDesc), len(file_v1_watchdog_watchdogService_proto_rawDesc)))
	})
	return file_v1_watchdog_watchdogService_proto_rawDescData
}

var file_v1_watchdog_watchdogService_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_watchdog_watchdogService_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_v1_watchdog_watchdogService_proto_goTypes = []any{
	(HealthResponse_HealthStatus)(0), // 0: watchdog.HealthResponse.HealthStatus
	(*ReloadRequest)(nil),            // 1: watchdog.ReloadRequest
	(*ReloadResponse)(nil),           // 2: watchdog.ReloadResponse
	(*GetRequest)(nil),               // 3: watchdog.GetRequest
	(*GetResponse)(nil),              // 4: watchdog.GetResponse
	(*GetDetailsRequest)(nil),        // 5: watchdog.GetDetailsRequest
	(*GetDetailsResponse)(nil),       // 6: watchdog.GetDetailsResponse
	(*ListSummariesRequest)(nil),     // 7: watchdog.ListSummariesRequest
	(*ListSummariesResponse)(nil),    // 8: watchdog.ListSummariesResponse
	(*HealthRequest)(nil),            // 9: watchdog.HealthRequest
	(*HealthResponse)(nil),           // 10: watchdog.HealthResponse
	(*DomainSummary)(nil),            // 11: watchdog.DomainSummary
	(*DomainRow)(nil),                // 12: watchdog.DomainRow
}
var file_v1_watchdog_watchdogService_proto_depIdxs = []int32{
	11, // 0: watchdog.GetResponse.summary:type_name -> watchdog.DomainSummary
	12, // 1: watchdog.GetDetailsResponse.domain:type_name -> watchdog.DomainRow
	11, // 2: watchdog.ListSummariesResponse.summaries:type_name -> watchdog.DomainSummary
	0,  // 3: watchdog.HealthResponse.status:type_name -> watchdog.HealthResponse.HealthStatus
	1,  // 4: watchdog.WatchDog.Reload:input_type -> watchdog.ReloadRequest
	3,  // 5: watchdog.WatchDog.Get:input_type -> watchdog.GetRequest
	5,  // 6: watchdog.WatchDog.GetDetails:input_type -> watchdog.GetDetailsRequest
	7,  // 7: watchdog.WatchDog.ListSummaries:input_type -> watchdog.ListSummariesRequest
	9,  // 8: watchdog.WatchDog.Health:input_type -> watchdog.HealthRequest
	2,  // 9: watchdog.WatchDog.Reload:output_type -> watchdog.ReloadResponse
	4,  // 10: watchdog.WatchDog.Get:output_type -> watchdog.GetResponse
	6,  // 11: watchdog.WatchDog.GetDetails:output_type -> watchdog.GetDetailsResponse
	8,  // 12: watchdog.WatchDog.ListSummaries:output_type -> watchdog.ListSummariesResponse
	10, // 13: watchdog.WatchDog.Health:output_type -> watchdog.HealthResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_v1_watchdog_watchdogService_proto_init() }
func file_v1_watchdog_watchdogService_proto_init() {
	if File_v1_watchdog_watchdogService_proto != nil {
		return
	}
	file_v1_watchdog_watchdog_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_watchdog_watchdogService_proto_rawDesc), len(file_v1_watchdog_watchdogService_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_watchdog_watchdogService_proto_goTypes,
		DependencyIndexes: file_v1_watchdog_watchdogService_proto_depIdxs,
		EnumInfos:         file_v1_watchdog_watchdogService_proto_enumTypes,
		MessageInfos:      file_v1_watchdog_watchdogService_proto_msgTypes,
	}.Build()
	File_v1_watchdog_watchdogService_proto = out.File
	file_v1_watchdog_watchdogService_proto_goTypes = nil
	file_v1_watchdog_watchdogService_proto_depIdxs = nil
}
