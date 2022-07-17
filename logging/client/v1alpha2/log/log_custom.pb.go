// Code generated by protoc-gen-goten-go
// File: edgelq/logging/proto/v1alpha2/log_custom.proto
// DO NOT EDIT!!!

package log_client

import (
	"fmt"
	"reflect"
	"sync"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	rpc "github.com/cloudwan/edgelq-sdk/common/rpc"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	logging_common "github.com/cloudwan/edgelq-sdk/logging/common/v1alpha2"
	log "github.com/cloudwan/edgelq-sdk/logging/resources/v1alpha2/log"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = reflect.Method{}
	_ = sync.Once{}

	_ = protojson.MarshalOptions{}
	_ = proto.MarshalOptions{}
	_ = preflect.Value{}
	_ = protoimpl.DescBuilder{}
)

// make sure we're using proto imports
var (
	_ = &rpc.Status{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &logging_common.LabelDescriptor{}
	_ = &log.Log{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method
// [ListLogs][ntt.logging.v1alpha2.ListLogs]
type ListLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Parent references of ntt.logging.v1alpha2.Log
	Parents []*log.ParentReference `protobuf:"bytes,1,rep,customtype=ParentReference,name=parents,proto3" json:"parents,omitempty" firestore:"parents"`
	// Filter that specifies which logs should be returned
	Filter *log.Filter `protobuf:"bytes,2,opt,customtype=Filter,name=filter,proto3" json:"filter,omitempty" firestore:"filter"`
	// The time interval for which results should be returned. Only logs
	// that contain data points in the specified interval are included
	// in the response.
	Interval *logging_common.TimeInterval `protobuf:"bytes,3,opt,name=interval,proto3" json:"interval,omitempty" firestore:"interval"`
	// Cap on a number of log entries to be included in a response.
	// Number of logs in an actual response can be higher, since logs are
	// read in bulk with second precision - exceed logs above the limit will share
	// same timestamp as the logs below the limit.
	//
	// Results will be adjusted to the "end time" taken from interval field
	// (adjusted also by page_token if provided).
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" firestore:"pageSize"`
	// Token which identifies next page with further results. Token should be
	// taken from
	// [ListLogsResponse.next_page_token][ntt.logging.v1alpha2.ListLogsResponse.next_page_token].
	PageToken string `protobuf:"bytes,6,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty" firestore:"pageToken"`
}

func (m *ListLogsRequest) Reset() {
	*m = ListLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListLogsRequest) ProtoMessage() {}

func (m *ListLogsRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListLogsRequest) GotenMessage() {}

// Deprecated, Use ListLogsRequest.ProtoReflect.Descriptor instead.
func (*ListLogsRequest) Descriptor() ([]byte, []int) {
	return edgelq_logging_proto_v1alpha2_log_custom_proto_rawDescGZIP(), []int{0}
}

func (m *ListLogsRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListLogsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListLogsRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListLogsRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListLogsRequest) GetParents() []*log.ParentReference {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *ListLogsRequest) GetFilter() *log.Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListLogsRequest) GetInterval() *logging_common.TimeInterval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *ListLogsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return int32(0)
}

func (m *ListLogsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListLogsRequest) SetParents(fv []*log.ParentReference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Parents", "ListLogsRequest"))
	}
	m.Parents = fv
}

func (m *ListLogsRequest) SetFilter(fv *log.Filter) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Filter", "ListLogsRequest"))
	}
	m.Filter = fv
}

func (m *ListLogsRequest) SetInterval(fv *logging_common.TimeInterval) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Interval", "ListLogsRequest"))
	}
	m.Interval = fv
}

func (m *ListLogsRequest) SetPageSize(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PageSize", "ListLogsRequest"))
	}
	m.PageSize = fv
}

func (m *ListLogsRequest) SetPageToken(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PageToken", "ListLogsRequest"))
	}
	m.PageToken = fv
}

// Response message for method
// [ListLogs][ntt.logging.v1alpha2.ListLogs]
type ListLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Logs that match the filter included in the request.
	Logs []*log.Log `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty" firestore:"logs"`
	// If there are more results than have been returned, then this field is set
	// to a non-empty value. To see the additional results,
	// use that value as `pageToken` in the next call to this method.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty" firestore:"nextPageToken"`
	// Query execution errors that may have caused the logs data returned
	// to be incomplete.
	ExecutionErrors []*rpc.Status `protobuf:"bytes,3,rep,name=execution_errors,json=executionErrors,proto3" json:"execution_errors,omitempty" firestore:"executionErrors"`
}

func (m *ListLogsResponse) Reset() {
	*m = ListLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListLogsResponse) ProtoMessage() {}

func (m *ListLogsResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListLogsResponse) GotenMessage() {}

// Deprecated, Use ListLogsResponse.ProtoReflect.Descriptor instead.
func (*ListLogsResponse) Descriptor() ([]byte, []int) {
	return edgelq_logging_proto_v1alpha2_log_custom_proto_rawDescGZIP(), []int{1}
}

func (m *ListLogsResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListLogsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListLogsResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListLogsResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListLogsResponse) GetLogs() []*log.Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *ListLogsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *ListLogsResponse) GetExecutionErrors() []*rpc.Status {
	if m != nil {
		return m.ExecutionErrors
	}
	return nil
}

func (m *ListLogsResponse) SetLogs(fv []*log.Log) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Logs", "ListLogsResponse"))
	}
	m.Logs = fv
}

func (m *ListLogsResponse) SetNextPageToken(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "NextPageToken", "ListLogsResponse"))
	}
	m.NextPageToken = fv
}

func (m *ListLogsResponse) SetExecutionErrors(fv []*rpc.Status) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ExecutionErrors", "ListLogsResponse"))
	}
	m.ExecutionErrors = fv
}

// Request message for method
// [CreateLogs][ntt.logging.v1alpha2.CreateLogs]
type CreateLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Parent reference of ntt.logging.v1alpha2.Log
	Parent *log.ParentReference `protobuf:"bytes,1,opt,customtype=ParentReference,name=parent,proto3" json:"parent,omitempty" firestore:"parent"`
	// List of logs to create/append. If they have specified name
	// field, it must match provided parent field.
	Logs []*log.Log `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty" firestore:"logs"`
}

func (m *CreateLogsRequest) Reset() {
	*m = CreateLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *CreateLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*CreateLogsRequest) ProtoMessage() {}

func (m *CreateLogsRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*CreateLogsRequest) GotenMessage() {}

// Deprecated, Use CreateLogsRequest.ProtoReflect.Descriptor instead.
func (*CreateLogsRequest) Descriptor() ([]byte, []int) {
	return edgelq_logging_proto_v1alpha2_log_custom_proto_rawDescGZIP(), []int{2}
}

func (m *CreateLogsRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *CreateLogsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *CreateLogsRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *CreateLogsRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *CreateLogsRequest) GetParent() *log.ParentReference {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *CreateLogsRequest) GetLogs() []*log.Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *CreateLogsRequest) SetParent(fv *log.ParentReference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Parent", "CreateLogsRequest"))
	}
	m.Parent = fv
}

func (m *CreateLogsRequest) SetLogs(fv []*log.Log) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Logs", "CreateLogsRequest"))
	}
	m.Logs = fv
}

// Response message for method
// [CreateLogs][ntt.logging.v1alpha2.CreateLogs]
type CreateLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Log names indexed by Create position. All logs, except
	// failed ones will be present. If all logs were written successfully,
	// then map will have keys all from 0 to N-1.
	LogNames map[uint32]*log.Reference `protobuf:"bytes,1,rep,customtype=Reference,name=log_names,json=logNames,proto3" json:"log_names,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3,customtype=Reference" firestore:"logNames"`
	// Logs that failed to be created
	FailedLogs []*CreateLogsResponse_CreateError `protobuf:"bytes,2,rep,name=failed_logs,json=failedLogs,proto3" json:"failed_logs,omitempty" firestore:"failedLogs"`
}

func (m *CreateLogsResponse) Reset() {
	*m = CreateLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *CreateLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*CreateLogsResponse) ProtoMessage() {}

func (m *CreateLogsResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*CreateLogsResponse) GotenMessage() {}

// Deprecated, Use CreateLogsResponse.ProtoReflect.Descriptor instead.
func (*CreateLogsResponse) Descriptor() ([]byte, []int) {
	return edgelq_logging_proto_v1alpha2_log_custom_proto_rawDescGZIP(), []int{3}
}

func (m *CreateLogsResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *CreateLogsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *CreateLogsResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *CreateLogsResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *CreateLogsResponse) GetLogNames() map[uint32]*log.Reference {
	if m != nil {
		return m.LogNames
	}
	return nil
}

func (m *CreateLogsResponse) GetFailedLogs() []*CreateLogsResponse_CreateError {
	if m != nil {
		return m.FailedLogs
	}
	return nil
}

func (m *CreateLogsResponse) SetLogNames(fv map[uint32]*log.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "LogNames", "CreateLogsResponse"))
	}
	m.LogNames = fv
}

func (m *CreateLogsResponse) SetFailedLogs(fv []*CreateLogsResponse_CreateError) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FailedLogs", "CreateLogsResponse"))
	}
	m.FailedLogs = fv
}

// ErrorDetails is used when one of the queried regions fails to produce
// results. It is used in execution_errors field (see subfield
// ntt.rpc.Status.details).
type ListLogsResponse_ErrorDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// region id which failed to give results.
	RegionId string `protobuf:"bytes,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty" firestore:"regionId"`
}

func (m *ListLogsResponse_ErrorDetails) Reset() {
	*m = ListLogsResponse_ErrorDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListLogsResponse_ErrorDetails) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListLogsResponse_ErrorDetails) ProtoMessage() {}

func (m *ListLogsResponse_ErrorDetails) ProtoReflect() preflect.Message {
	mi := &edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListLogsResponse_ErrorDetails) GotenMessage() {}

// Deprecated, Use ListLogsResponse_ErrorDetails.ProtoReflect.Descriptor instead.
func (*ListLogsResponse_ErrorDetails) Descriptor() ([]byte, []int) {
	return edgelq_logging_proto_v1alpha2_log_custom_proto_rawDescGZIP(), []int{1, 0}
}

func (m *ListLogsResponse_ErrorDetails) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListLogsResponse_ErrorDetails) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListLogsResponse_ErrorDetails) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListLogsResponse_ErrorDetails) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListLogsResponse_ErrorDetails) GetRegionId() string {
	if m != nil {
		return m.RegionId
	}
	return ""
}

func (m *ListLogsResponse_ErrorDetails) SetRegionId(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RegionId", "ListLogsResponse_ErrorDetails"))
	}
	m.RegionId = fv
}

// Describes the result of a failed request to write logs.
type CreateLogsResponse_CreateError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// All logs that failed to be written. This field provides all of
	// the context that would be needed to retry the operation.
	Logs []*log.Log `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty" firestore:"logs"`
	// The status of the requested write operation.
	Status *rpc.Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty" firestore:"status"`
}

func (m *CreateLogsResponse_CreateError) Reset() {
	*m = CreateLogsResponse_CreateError{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *CreateLogsResponse_CreateError) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*CreateLogsResponse_CreateError) ProtoMessage() {}

func (m *CreateLogsResponse_CreateError) ProtoReflect() preflect.Message {
	mi := &edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*CreateLogsResponse_CreateError) GotenMessage() {}

// Deprecated, Use CreateLogsResponse_CreateError.ProtoReflect.Descriptor instead.
func (*CreateLogsResponse_CreateError) Descriptor() ([]byte, []int) {
	return edgelq_logging_proto_v1alpha2_log_custom_proto_rawDescGZIP(), []int{3, 0}
}

func (m *CreateLogsResponse_CreateError) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *CreateLogsResponse_CreateError) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *CreateLogsResponse_CreateError) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *CreateLogsResponse_CreateError) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *CreateLogsResponse_CreateError) GetLogs() []*log.Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *CreateLogsResponse_CreateError) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CreateLogsResponse_CreateError) SetLogs(fv []*log.Log) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Logs", "CreateLogsResponse_CreateError"))
	}
	m.Logs = fv
}

func (m *CreateLogsResponse_CreateError) SetStatus(fv *rpc.Status) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Status", "CreateLogsResponse_CreateError"))
	}
	m.Status = fv
}

var edgelq_logging_proto_v1alpha2_log_custom_proto preflect.FileDescriptor

var edgelq_logging_proto_v1alpha2_log_custom_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x0f, 0xb2, 0xda, 0x21, 0x07, 0x42, 0x05, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0xba, 0x9d, 0x22, 0x00,
	0x52, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xb2, 0xda, 0x21, 0x07, 0x1a,
	0x05, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x48,
	0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x42, 0x08, 0xca, 0xc6, 0x27, 0x04, 0x62, 0x02, 0x08, 0x01, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xd2, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x6c, 0x6f, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4c,
	0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x3a, 0x0a, 0x10, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x2b, 0x0a, 0x0c,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x6b, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f,
	0xb2, 0xda, 0x21, 0x07, 0x42, 0x05, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0xba, 0x9d, 0x22, 0x00, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4c, 0x6f, 0x67,
	0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0xf1, 0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a,
	0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0b, 0xb2, 0xda, 0x21, 0x07, 0x12, 0x05,
	0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12,
	0x55, 0x0a, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x0a, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x4c, 0x6f, 0x67, 0x73, 0x1a, 0x65, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x04,
	0x6c, 0x6f, 0x67, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x3b, 0x0a,
	0x0d, 0x4c, 0x6f, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xc5, 0x01, 0xe8, 0xde, 0x21,
	0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0e,
	0x4c, 0x6f, 0x67, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00,
	0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x6c, 0x6f, 0x67, 0x3b, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0xd2, 0x84, 0xd1, 0x02, 0x4a, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_logging_proto_v1alpha2_log_custom_proto_rawDescOnce sync.Once
	edgelq_logging_proto_v1alpha2_log_custom_proto_rawDescData = edgelq_logging_proto_v1alpha2_log_custom_proto_rawDesc
)

func edgelq_logging_proto_v1alpha2_log_custom_proto_rawDescGZIP() []byte {
	edgelq_logging_proto_v1alpha2_log_custom_proto_rawDescOnce.Do(func() {
		edgelq_logging_proto_v1alpha2_log_custom_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_logging_proto_v1alpha2_log_custom_proto_rawDescData)
	})
	return edgelq_logging_proto_v1alpha2_log_custom_proto_rawDescData
}

var edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var edgelq_logging_proto_v1alpha2_log_custom_proto_goTypes = []interface{}{
	(*ListLogsRequest)(nil),                // 0: ntt.logging.v1alpha2.ListLogsRequest
	(*ListLogsResponse)(nil),               // 1: ntt.logging.v1alpha2.ListLogsResponse
	(*CreateLogsRequest)(nil),              // 2: ntt.logging.v1alpha2.CreateLogsRequest
	(*CreateLogsResponse)(nil),             // 3: ntt.logging.v1alpha2.CreateLogsResponse
	(*ListLogsResponse_ErrorDetails)(nil),  // 4: ntt.logging.v1alpha2.ListLogsResponse.ErrorDetails
	(*CreateLogsResponse_CreateError)(nil), // 5: ntt.logging.v1alpha2.CreateLogsResponse.CreateError
	nil,                                    // 6: ntt.logging.v1alpha2.CreateLogsResponse.LogNamesEntry
	(*logging_common.TimeInterval)(nil),    // 7: ntt.logging.v1alpha2.TimeInterval
	(*log.Log)(nil),                        // 8: ntt.logging.v1alpha2.Log
	(*rpc.Status)(nil),                     // 9: ntt.rpc.Status
}
var edgelq_logging_proto_v1alpha2_log_custom_proto_depIdxs = []int32{
	7, // 0: ntt.logging.v1alpha2.ListLogsRequest.interval:type_name -> ntt.logging.v1alpha2.TimeInterval
	8, // 1: ntt.logging.v1alpha2.ListLogsResponse.logs:type_name -> ntt.logging.v1alpha2.Log
	9, // 2: ntt.logging.v1alpha2.ListLogsResponse.execution_errors:type_name -> ntt.rpc.Status
	8, // 3: ntt.logging.v1alpha2.CreateLogsRequest.logs:type_name -> ntt.logging.v1alpha2.Log
	6, // 4: ntt.logging.v1alpha2.CreateLogsResponse.log_names:type_name -> ntt.logging.v1alpha2.CreateLogsResponse.LogNamesEntry
	5, // 5: ntt.logging.v1alpha2.CreateLogsResponse.failed_logs:type_name -> ntt.logging.v1alpha2.CreateLogsResponse.CreateError
	8, // 6: ntt.logging.v1alpha2.CreateLogsResponse.CreateError.logs:type_name -> ntt.logging.v1alpha2.Log
	9, // 7: ntt.logging.v1alpha2.CreateLogsResponse.CreateError.status:type_name -> ntt.rpc.Status
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { edgelq_logging_proto_v1alpha2_log_custom_proto_init() }
func edgelq_logging_proto_v1alpha2_log_custom_proto_init() {
	if edgelq_logging_proto_v1alpha2_log_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLogsRequest); i {
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
		edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLogsResponse); i {
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
		edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLogsRequest); i {
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
		edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLogsResponse); i {
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
		edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLogsResponse_ErrorDetails); i {
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
		edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLogsResponse_CreateError); i {
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
			RawDescriptor: edgelq_logging_proto_v1alpha2_log_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_logging_proto_v1alpha2_log_custom_proto_goTypes,
		DependencyIndexes: edgelq_logging_proto_v1alpha2_log_custom_proto_depIdxs,
		MessageInfos:      edgelq_logging_proto_v1alpha2_log_custom_proto_msgTypes,
	}.Build()
	edgelq_logging_proto_v1alpha2_log_custom_proto = out.File
	edgelq_logging_proto_v1alpha2_log_custom_proto_rawDesc = nil
	edgelq_logging_proto_v1alpha2_log_custom_proto_goTypes = nil
	edgelq_logging_proto_v1alpha2_log_custom_proto_depIdxs = nil
}
