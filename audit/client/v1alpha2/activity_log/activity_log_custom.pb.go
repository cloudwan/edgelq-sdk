// Code generated by protoc-gen-goten-go
// File: edgelq/audit/proto/v1alpha2/activity_log_custom.proto
// DO NOT EDIT!!!

package activity_log_client

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
	activity_log "github.com/cloudwan/edgelq-sdk/audit/resources/v1alpha2/activity_log"
	common "github.com/cloudwan/edgelq-sdk/audit/resources/v1alpha2/common"
	rpc "github.com/cloudwan/edgelq-sdk/common/rpc"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
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
	_ = &activity_log.ActivityLog{}
	_ = &common.Authentication{}
	_ = &rpc.Status{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method
// [ListActivityLogs][ntt.audit.v1alpha2.ListActivityLogs]
//
// Returns activities for specified time range and within specified filter.
// Activity logs are stored only in region which executed them and never
// duplicated. Because of that its important to pay attention to region IDs in a
// request object.
//
// Basic supported filters are:
// * --filter 'service.name=[SERVICE_NAME]' (what is happening in this service)
// * --filter 'service.name=[SERVICE_NAME] and method.type=[METHOD_NAME]' (what
// is happening for this API call)
// * --filter 'authentication.principal=[PRINCIPAL_NAME]' (what that person is
// doing)
// * --filter 'request_id=[REQUEST_ID]' (I have request ID, what is actually
// this?)
// * --filter 'service.name=[SERVICE_NAME] and
// resource.name=[FULL_RESOURCE_NAME]' (can I see activities on this resource?)
//
// Its also possible to filter logs by their region of activity - by using field
// service.region_id in a filter field. Its important to note that logs may be
// present in multiple locations, if request was routed somewhere else or split
// & merged across many regions. Those activity logs may have different activity
// log names, but they will share same values in fields request_id and
// request_routing.
//
// Be aware, that server will append scope filter condition (and scope=...) to
// the filter. Scope(s) will be extracted from fields parents in
// ListActivityLogsRequest object. Note you can query for multiple at once, both
// projects and organizations.
//
// For all of the above filters you can replace filter condition compare (=)
// with IN operator. You can therefore query for multiple services, methods or
// users at once. Above filters are also preferred as we have optimization for
// them.
//
// Activity logs can be filtered by custom labels (field labels in ActivityLog).
// Labels are defined per each API method - so you must specify service.name and
// method.type conditions to be able to query by labels.
//
// For example, suppose you have a CreateVM method, which creates resource "VM".
// Suppose there is a field "group" within resource body, which is reference to
// other resource. If you want to make a query like "who was creating VMs for
// that group", then you need to create label "group" inside resource body. Then
// you will be able to make a query with following filter condition:
// *--filter 'service.name=vms.domain.com and method.type=CreateVM and
// labels.group=mySpecialVMGroup'*.
//
// Be aware, that Create/Update requests, which have resource object in their
// own bodies, will automatically inherit resource labels. So, basically you
// need to define "group" label in resource spec, not inside request. This is
// useful, as both Create/Update methods will have this label. This also allows
// code-gen to continue maintaining *derived.proto files.
//
// Examples of usage (with cuttle - we are interested only in one region and
// scope):
//
// * Checks connections to all devices within ssh-demo project starting from 8th
// of September 12 UTC time
//
// $ cuttle audit query activity-log --parents 'project/ssh-demo'
//   --filter 'service.name="devices.edgelq.com" and
//   method.type="ConnectToDevice" and service.regionId="us-west"'
//   --interval '{"startTime":"2020-09-08T12:00:00Z"}' -o json
//
// *Checks connections to device demo-device within ssh-demo project starting
// from 8th of September 12 UTC time
//
// $ cuttle audit query activity-log --parents 'project/ssh-demo'
//   --filter 'service.name="devices.edgelq.com" and
//   method.type="ConnectToDevice" and service.regionId="us-west" and
//     resource.name="projects/ssh-demo/devices/demo-device"'
//   --interval '{"startTime":"2020-09-08T12:00:00Z"}' -o json
//
// * Checks what is happening within whole iam service for project demo starting
// from 8th of September 12 UTC time
//
// $ cuttle audit query activity-log --parents 'project/demo'
//   --filter 'service.name="iam.edgelq.com" and service.regionId="us-west"'
//   --filter 'service.name="iam.edgelq.com"'
//   --interval '{"startTime":"2020-09-08T12:00:00Z"}' -o json
//
// * Checks activities within one hour for whole iam service for selected
// methods
//
// $ cuttle audit query activity-log --parents 'project/demo'
//   --filter 'service.name="iam.edgelq.com" and method.type IN
//   ["CreateRoleBinding", "UpdateRoleBinding", "DeleteRoleBinding"]
//   and service.regionId="us-west"'
//   --interval '{"startTime":"2020-09-08T12:00:00Z",
//   "endTime":"2020-09-08T13:00:00Z"}' -o json
//
// * Checks modification of RoleBinding
//
// $ cuttle audit query activity-log --parents 'project/demo'
//   --filter 'service.name="iam.edgelq.com" and method.type="UpdateRoleBinding"
//   and labels.resource_name="projects/x/roleBindings/myRB"'
//   --interval '{"startTime":"2020-09-08T12:00:00Z"}' -o json
//
// * Checks what was happening with some device
//
// $ cuttle audit query activity-log --parents 'project/demo'
//   --filter 'service.name="devices.edgelq.com" and
//   resource.name="projects/x/devices/myDevice" and service.regionId="us-west"'
//   --interval '{"startTime":"2020-09-08T12:00:00Z"}' -o json
//
// * Checks activities made by specific user (we need their email)
//
// $ cuttle audit query activity-log --parents 'project/demo'
//   --filter 'authentication.principal="user:we.know.who@domain.com" and
//   service.regionId="us-west"'
//   --interval '{"startTime":"2020-09-08T12:00:00Z"}' -o json
//
// * Checks activities made by specific service account (we need it's email)
//
// $ cuttle audit query activity-log --parents 'project/demo'
//   --filter
//   'authentication.principal="serviceAccount:myServiceAccount@domain.com" and
//   service.regionId="us-west"'
type ListActivityLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Parent references of ntt.audit.v1alpha2.ActivityLog - provides list of all
	// scopes we want to query about
	Parents []*activity_log.ParentName `protobuf:"bytes,1,rep,customtype=ParentName,name=parents,proto3" json:"parents,omitempty" firestore:"parents"`
	// A audit filter that specifies which activity logs should be returned
	Filter *activity_log.Filter `protobuf:"bytes,2,opt,customtype=Filter,name=filter,proto3" json:"filter,omitempty" firestore:"filter"`
	// The time interval for which results should be returned. Only logs
	// that contain data points in the specified interval are included
	// in the response.
	Interval *common.TimeInterval `protobuf:"bytes,4,opt,name=interval,proto3" json:"interval,omitempty" firestore:"interval"`
	// Cap on a number of activity logs to be included in a response.
	// Number of logs in an actual response can be higher, since logs are
	// read in bulk with second precision - exceed logs above the limit will share
	// same timestamp as the logs below the limit.
	//
	// Results will be adjusted to the "end time" taken from interval field
	// (adjusted also by page_token if provided).
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" firestore:"pageSize"`
	// Token which identifies next page with further results. Token should be
	// taken from
	// [ListActivityLogsResponse.next_page_token][ntt.audit.v1alpha2.ListActivityLogsResponse.next_page_token].
	PageToken string `protobuf:"bytes,6,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty" firestore:"pageToken"`
}

func (m *ListActivityLogsRequest) Reset() {
	*m = ListActivityLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListActivityLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListActivityLogsRequest) ProtoMessage() {}

func (m *ListActivityLogsRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListActivityLogsRequest) GotenMessage() {}

// Deprecated, Use ListActivityLogsRequest.ProtoReflect.Descriptor instead.
func (*ListActivityLogsRequest) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha2_activity_log_custom_proto_rawDescGZIP(), []int{0}
}

func (m *ListActivityLogsRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListActivityLogsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListActivityLogsRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListActivityLogsRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListActivityLogsRequest) GetParents() []*activity_log.ParentName {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *ListActivityLogsRequest) GetFilter() *activity_log.Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListActivityLogsRequest) GetInterval() *common.TimeInterval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *ListActivityLogsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return int32(0)
}

func (m *ListActivityLogsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListActivityLogsRequest) SetParents(fv []*activity_log.ParentName) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Parents", "ListActivityLogsRequest"))
	}
	m.Parents = fv
}

func (m *ListActivityLogsRequest) SetFilter(fv *activity_log.Filter) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Filter", "ListActivityLogsRequest"))
	}
	m.Filter = fv
}

func (m *ListActivityLogsRequest) SetInterval(fv *common.TimeInterval) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Interval", "ListActivityLogsRequest"))
	}
	m.Interval = fv
}

func (m *ListActivityLogsRequest) SetPageSize(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PageSize", "ListActivityLogsRequest"))
	}
	m.PageSize = fv
}

func (m *ListActivityLogsRequest) SetPageToken(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PageToken", "ListActivityLogsRequest"))
	}
	m.PageToken = fv
}

// Response message for method
// [ListActivityLogs][ntt.audit.v1alpha2.ListActivityLogs]
type ListActivityLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// One or more activity method logs that match the filter included in the
	// request. Contains results from all queried regions. Its possible however
	// that some logs may be missing, for this see execution_errors.
	ActivityLogs []*activity_log.ActivityLog `protobuf:"bytes,1,rep,name=activity_logs,json=activityLogs,proto3" json:"activity_logs,omitempty" firestore:"activityLogs"`
	// If there are more results than have been returned, then this field is set
	// to a non-empty value. To see the additional results,
	// use that value as `pageToken` in the next call to this method.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty" firestore:"nextPageToken"`
	// Query execution errors that may have caused the response data returned to
	// be incomplete. Because logs are stored only locally (for each region), all
	// activity log queries are split and merged by a receiving request server
	// according to the queried regions. Its possible that some regions will fail
	// when request is redirected to them, but others not. For each failed region,
	// one execution error will be appended. In each ntt.rpc.Status message,
	// fields code and message will contain error obtained from failed regional
	// server, while field details will contain always one item and this item will
	// be of type ErrorDetails.
	ExecutionErrors []*rpc.Status `protobuf:"bytes,3,rep,name=execution_errors,json=executionErrors,proto3" json:"execution_errors,omitempty" firestore:"executionErrors"`
}

func (m *ListActivityLogsResponse) Reset() {
	*m = ListActivityLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListActivityLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListActivityLogsResponse) ProtoMessage() {}

func (m *ListActivityLogsResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListActivityLogsResponse) GotenMessage() {}

// Deprecated, Use ListActivityLogsResponse.ProtoReflect.Descriptor instead.
func (*ListActivityLogsResponse) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha2_activity_log_custom_proto_rawDescGZIP(), []int{1}
}

func (m *ListActivityLogsResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListActivityLogsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListActivityLogsResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListActivityLogsResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListActivityLogsResponse) GetActivityLogs() []*activity_log.ActivityLog {
	if m != nil {
		return m.ActivityLogs
	}
	return nil
}

func (m *ListActivityLogsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *ListActivityLogsResponse) GetExecutionErrors() []*rpc.Status {
	if m != nil {
		return m.ExecutionErrors
	}
	return nil
}

func (m *ListActivityLogsResponse) SetActivityLogs(fv []*activity_log.ActivityLog) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ActivityLogs", "ListActivityLogsResponse"))
	}
	m.ActivityLogs = fv
}

func (m *ListActivityLogsResponse) SetNextPageToken(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "NextPageToken", "ListActivityLogsResponse"))
	}
	m.NextPageToken = fv
}

func (m *ListActivityLogsResponse) SetExecutionErrors(fv []*rpc.Status) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ExecutionErrors", "ListActivityLogsResponse"))
	}
	m.ExecutionErrors = fv
}

// Request message for method
// [CreateActivityLogs][ntt.audit.v1alpha2.CreateActivityLogs]
//
// Creates many activity logs at once - or appends existing, if some of the
// activity logs already exist (their name is already known).
//
// This request should not be used by regular users - only API services should
// be able to submit activity logs. Developers of services should use logs
// exporter package offered along other Audit service packages instead of
// developing own components.
type CreateActivityLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// List of activity logs to be added to service. Can be coming from different
	// scopes but must be submitted to the same region.
	ActivityLogs []*activity_log.ActivityLog `protobuf:"bytes,1,rep,name=activity_logs,json=activityLogs,proto3" json:"activity_logs,omitempty" firestore:"activityLogs"`
}

func (m *CreateActivityLogsRequest) Reset() {
	*m = CreateActivityLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *CreateActivityLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*CreateActivityLogsRequest) ProtoMessage() {}

func (m *CreateActivityLogsRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*CreateActivityLogsRequest) GotenMessage() {}

// Deprecated, Use CreateActivityLogsRequest.ProtoReflect.Descriptor instead.
func (*CreateActivityLogsRequest) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha2_activity_log_custom_proto_rawDescGZIP(), []int{2}
}

func (m *CreateActivityLogsRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *CreateActivityLogsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *CreateActivityLogsRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *CreateActivityLogsRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *CreateActivityLogsRequest) GetActivityLogs() []*activity_log.ActivityLog {
	if m != nil {
		return m.ActivityLogs
	}
	return nil
}

func (m *CreateActivityLogsRequest) SetActivityLogs(fv []*activity_log.ActivityLog) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ActivityLogs", "CreateActivityLogsRequest"))
	}
	m.ActivityLogs = fv
}

// Response message for method
// [CreateActivityLogs][ntt.audit.v1alpha2.CreateActivityLogs]
type CreateActivityLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Activity log names - one name per each activity log, in same order
	// as in the request
	LogNames []*activity_log.Name `protobuf:"bytes,1,rep,customtype=Name,name=log_names,json=logNames,proto3" json:"log_names,omitempty" firestore:"logNames"`
}

func (m *CreateActivityLogsResponse) Reset() {
	*m = CreateActivityLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *CreateActivityLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*CreateActivityLogsResponse) ProtoMessage() {}

func (m *CreateActivityLogsResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*CreateActivityLogsResponse) GotenMessage() {}

// Deprecated, Use CreateActivityLogsResponse.ProtoReflect.Descriptor instead.
func (*CreateActivityLogsResponse) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha2_activity_log_custom_proto_rawDescGZIP(), []int{3}
}

func (m *CreateActivityLogsResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *CreateActivityLogsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *CreateActivityLogsResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *CreateActivityLogsResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *CreateActivityLogsResponse) GetLogNames() []*activity_log.Name {
	if m != nil {
		return m.LogNames
	}
	return nil
}

func (m *CreateActivityLogsResponse) SetLogNames(fv []*activity_log.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "LogNames", "CreateActivityLogsResponse"))
	}
	m.LogNames = fv
}

// ErrorDetails is used when one of the queried regions fails to produce
// results. It is used in execution_errors field (see subfield
// ntt.rpc.Status.details).
type ListActivityLogsResponse_ErrorDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// region id which failed to give results.
	RegionId string `protobuf:"bytes,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty" firestore:"regionId"`
}

func (m *ListActivityLogsResponse_ErrorDetails) Reset() {
	*m = ListActivityLogsResponse_ErrorDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ListActivityLogsResponse_ErrorDetails) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ListActivityLogsResponse_ErrorDetails) ProtoMessage() {}

func (m *ListActivityLogsResponse_ErrorDetails) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ListActivityLogsResponse_ErrorDetails) GotenMessage() {}

// Deprecated, Use ListActivityLogsResponse_ErrorDetails.ProtoReflect.Descriptor instead.
func (*ListActivityLogsResponse_ErrorDetails) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha2_activity_log_custom_proto_rawDescGZIP(), []int{1, 0}
}

func (m *ListActivityLogsResponse_ErrorDetails) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ListActivityLogsResponse_ErrorDetails) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ListActivityLogsResponse_ErrorDetails) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ListActivityLogsResponse_ErrorDetails) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ListActivityLogsResponse_ErrorDetails) GetRegionId() string {
	if m != nil {
		return m.RegionId
	}
	return ""
}

func (m *ListActivityLogsResponse_ErrorDetails) SetRegionId(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RegionId", "ListActivityLogsResponse_ErrorDetails"))
	}
	m.RegionId = fv
}

var edgelq_audit_proto_v1alpha2_activity_log_custom_proto preflect.FileDescriptor

var edgelq_audit_proto_v1alpha2_activity_log_custom_proto_rawDesc = []byte{
	0x0a, 0x35, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x20, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff,
	0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x13, 0xb2, 0xda, 0x21,
	0x0f, 0x3a, 0x0d, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67,
	0x52, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xb2, 0xda, 0x21, 0x0f, 0x1a,
	0x0d, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x08, 0xca, 0xc6, 0x27, 0x04,
	0x62, 0x02, 0x08, 0x01, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04,
	0x22, 0xf1, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a,
	0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c,
	0x6f, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3a, 0x0a, 0x10, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x2b, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x22, 0x61, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x44, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x6f,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0x4e, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x13, 0xb2, 0xda, 0x21, 0x0f, 0x0a, 0x0d,
	0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x08, 0x6c,
	0x6f, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x42, 0x88, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x16, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c,
	0x6f, 0x67, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a,
	0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x3b, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_audit_proto_v1alpha2_activity_log_custom_proto_rawDescOnce sync.Once
	edgelq_audit_proto_v1alpha2_activity_log_custom_proto_rawDescData = edgelq_audit_proto_v1alpha2_activity_log_custom_proto_rawDesc
)

func edgelq_audit_proto_v1alpha2_activity_log_custom_proto_rawDescGZIP() []byte {
	edgelq_audit_proto_v1alpha2_activity_log_custom_proto_rawDescOnce.Do(func() {
		edgelq_audit_proto_v1alpha2_activity_log_custom_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_audit_proto_v1alpha2_activity_log_custom_proto_rawDescData)
	})
	return edgelq_audit_proto_v1alpha2_activity_log_custom_proto_rawDescData
}

var edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var edgelq_audit_proto_v1alpha2_activity_log_custom_proto_goTypes = []interface{}{
	(*ListActivityLogsRequest)(nil),               // 0: ntt.audit.v1alpha2.ListActivityLogsRequest
	(*ListActivityLogsResponse)(nil),              // 1: ntt.audit.v1alpha2.ListActivityLogsResponse
	(*CreateActivityLogsRequest)(nil),             // 2: ntt.audit.v1alpha2.CreateActivityLogsRequest
	(*CreateActivityLogsResponse)(nil),            // 3: ntt.audit.v1alpha2.CreateActivityLogsResponse
	(*ListActivityLogsResponse_ErrorDetails)(nil), // 4: ntt.audit.v1alpha2.ListActivityLogsResponse.ErrorDetails
	(*common.TimeInterval)(nil),                   // 5: ntt.audit.v1alpha2.TimeInterval
	(*activity_log.ActivityLog)(nil),              // 6: ntt.audit.v1alpha2.ActivityLog
	(*rpc.Status)(nil),                            // 7: ntt.rpc.Status
}
var edgelq_audit_proto_v1alpha2_activity_log_custom_proto_depIdxs = []int32{
	5, // 0: ntt.audit.v1alpha2.ListActivityLogsRequest.interval:type_name -> ntt.audit.v1alpha2.TimeInterval
	6, // 1: ntt.audit.v1alpha2.ListActivityLogsResponse.activity_logs:type_name -> ntt.audit.v1alpha2.ActivityLog
	7, // 2: ntt.audit.v1alpha2.ListActivityLogsResponse.execution_errors:type_name -> ntt.rpc.Status
	6, // 3: ntt.audit.v1alpha2.CreateActivityLogsRequest.activity_logs:type_name -> ntt.audit.v1alpha2.ActivityLog
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { edgelq_audit_proto_v1alpha2_activity_log_custom_proto_init() }
func edgelq_audit_proto_v1alpha2_activity_log_custom_proto_init() {
	if edgelq_audit_proto_v1alpha2_activity_log_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActivityLogsRequest); i {
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
		edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActivityLogsResponse); i {
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
		edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateActivityLogsRequest); i {
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
		edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateActivityLogsResponse); i {
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
		edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActivityLogsResponse_ErrorDetails); i {
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
			RawDescriptor: edgelq_audit_proto_v1alpha2_activity_log_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_audit_proto_v1alpha2_activity_log_custom_proto_goTypes,
		DependencyIndexes: edgelq_audit_proto_v1alpha2_activity_log_custom_proto_depIdxs,
		MessageInfos:      edgelq_audit_proto_v1alpha2_activity_log_custom_proto_msgTypes,
	}.Build()
	edgelq_audit_proto_v1alpha2_activity_log_custom_proto = out.File
	edgelq_audit_proto_v1alpha2_activity_log_custom_proto_rawDesc = nil
	edgelq_audit_proto_v1alpha2_activity_log_custom_proto_goTypes = nil
	edgelq_audit_proto_v1alpha2_activity_log_custom_proto_depIdxs = nil
}
