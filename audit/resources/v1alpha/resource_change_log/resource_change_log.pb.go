// Code generated by protoc-gen-goten-go
// File: edgelq/audit/proto/v1alpha/resource_change_log.proto
// DO NOT EDIT!!!

package resource_change_log

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
	audit_common "github.com/cloudwan/edgelq-sdk/audit/common/v1alpha"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	_ = &audit_common.Authentication{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &timestamp.Timestamp{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Type of change
type ResourceChangeLog_ResourceChange_Action int32

const (
	ResourceChangeLog_ResourceChange_UNDEFINED ResourceChangeLog_ResourceChange_Action = 0
	ResourceChangeLog_ResourceChange_CREATE    ResourceChangeLog_ResourceChange_Action = 1
	ResourceChangeLog_ResourceChange_UPDATE    ResourceChangeLog_ResourceChange_Action = 2
	ResourceChangeLog_ResourceChange_DELETE    ResourceChangeLog_ResourceChange_Action = 3
)

var (
	ResourceChangeLog_ResourceChange_Action_name = map[int32]string{
		0: "UNDEFINED",
		1: "CREATE",
		2: "UPDATE",
		3: "DELETE",
	}

	ResourceChangeLog_ResourceChange_Action_value = map[string]int32{
		"UNDEFINED": 0,
		"CREATE":    1,
		"UPDATE":    2,
		"DELETE":    3,
	}
)

func (x ResourceChangeLog_ResourceChange_Action) Enum() *ResourceChangeLog_ResourceChange_Action {
	p := new(ResourceChangeLog_ResourceChange_Action)
	*p = x
	return p
}

func (x ResourceChangeLog_ResourceChange_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), preflect.EnumNumber(x))
}

func (ResourceChangeLog_ResourceChange_Action) Descriptor() preflect.EnumDescriptor {
	return edgelq_audit_proto_v1alpha_resource_change_log_proto_enumTypes[0].Descriptor()
}

func (ResourceChangeLog_ResourceChange_Action) Type() preflect.EnumType {
	return &edgelq_audit_proto_v1alpha_resource_change_log_proto_enumTypes[0]
}

func (x ResourceChangeLog_ResourceChange_Action) Number() preflect.EnumNumber {
	return preflect.EnumNumber(x)
}

// Deprecated, Use ResourceChangeLog_ResourceChange_Action.ProtoReflect.Descriptor instead.
func (ResourceChangeLog_ResourceChange_Action) EnumDescriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha_resource_change_log_proto_rawDescGZIP(), []int{0, 0, 0}
}

// State of the transaction.
type ResourceChangeLog_TransactionInfo_State int32

const (
	ResourceChangeLog_TransactionInfo_UNDEFINED ResourceChangeLog_TransactionInfo_State = 0
	// Indicates that this change did not happen -
	// it is just proposal of the change.
	// Such a log should be followed by another
	// ResourceChangeLog with value COMMITTED
	// or ROLLED_BACK.
	// If one transaction has been retried
	// multiple times, then there may be multiple
	// records with PRE_COMMITTED, last record
	// should indicate final transaction state.
	ResourceChangeLog_TransactionInfo_PRE_COMMITTED ResourceChangeLog_TransactionInfo_State = 1
	// Indicates change has been committed
	// successfully.
	ResourceChangeLog_TransactionInfo_COMMITTED ResourceChangeLog_TransactionInfo_State = 2
	// Indicates that change did not happen.
	// Log of this type should be treated as
	// attempt of change.
	ResourceChangeLog_TransactionInfo_ROLLED_BACK ResourceChangeLog_TransactionInfo_State = 3
)

var (
	ResourceChangeLog_TransactionInfo_State_name = map[int32]string{
		0: "UNDEFINED",
		1: "PRE_COMMITTED",
		2: "COMMITTED",
		3: "ROLLED_BACK",
	}

	ResourceChangeLog_TransactionInfo_State_value = map[string]int32{
		"UNDEFINED":     0,
		"PRE_COMMITTED": 1,
		"COMMITTED":     2,
		"ROLLED_BACK":   3,
	}
)

func (x ResourceChangeLog_TransactionInfo_State) Enum() *ResourceChangeLog_TransactionInfo_State {
	p := new(ResourceChangeLog_TransactionInfo_State)
	*p = x
	return p
}

func (x ResourceChangeLog_TransactionInfo_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), preflect.EnumNumber(x))
}

func (ResourceChangeLog_TransactionInfo_State) Descriptor() preflect.EnumDescriptor {
	return edgelq_audit_proto_v1alpha_resource_change_log_proto_enumTypes[1].Descriptor()
}

func (ResourceChangeLog_TransactionInfo_State) Type() preflect.EnumType {
	return &edgelq_audit_proto_v1alpha_resource_change_log_proto_enumTypes[1]
}

func (x ResourceChangeLog_TransactionInfo_State) Number() preflect.EnumNumber {
	return preflect.EnumNumber(x)
}

// Deprecated, Use ResourceChangeLog_TransactionInfo_State.ProtoReflect.Descriptor instead.
func (ResourceChangeLog_TransactionInfo_State) EnumDescriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha_resource_change_log_proto_rawDescGZIP(), []int{0, 1, 0}
}

// ResourceChangeLog Resource
type ResourceChangeLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of ResourceChangeLog. It contains scope + ID of the log.
	// ID is a base64 encoded unique key that identifies tuple:
	//   scope
	//   request_id
	//   authentication.principal
	//   service.name
	//   resource.name
	//   resource.type
	//   resource.pre.labels
	//   resource.post.labels
	//
	// ID part should not be decoded, but treated as opaque string
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Contains scope from name field without resource ID.
	// Used for internal purpose for filtering (logs are using custom store).
	// Supported formats are:
	// - organization/umbrella
	// - projects/mars_exploration
	// - <system>
	Scope string `protobuf:"bytes,2,opt,name=scope,proto3" json:"scope,omitempty" firestore:"scope"`
	// Unique identifier of request - it must match the one
	// in the associated activity log.
	RequestId uint64 `protobuf:"varint,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty" firestore:"requestId"`
	// Time of the change - equal to request timestamp (activity log)
	Timestamp *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty" firestore:"timestamp"`
	// Authentication data - informs who made a change
	Authentication *audit_common.Authentication `protobuf:"bytes,5,opt,name=authentication,proto3" json:"authentication,omitempty" firestore:"authentication"`
	// Information about the service
	Service *audit_common.ServiceData `protobuf:"bytes,6,opt,name=service,proto3" json:"service,omitempty" firestore:"service"`
	// Describes change on the resource
	Resource *ResourceChangeLog_ResourceChange `protobuf:"bytes,7,opt,name=resource,proto3" json:"resource,omitempty" firestore:"resource"`
	// Describes state of the transaction
	Transaction *ResourceChangeLog_TransactionInfo `protobuf:"bytes,8,opt,name=transaction,proto3" json:"transaction,omitempty" firestore:"transaction"`
}

func (m *ResourceChangeLog) Reset() {
	*m = ResourceChangeLog{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha_resource_change_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ResourceChangeLog) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ResourceChangeLog) ProtoMessage() {}

func (m *ResourceChangeLog) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha_resource_change_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ResourceChangeLog) GotenMessage() {}

// Deprecated, Use ResourceChangeLog.ProtoReflect.Descriptor instead.
func (*ResourceChangeLog) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha_resource_change_log_proto_rawDescGZIP(), []int{0}
}

func (m *ResourceChangeLog) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ResourceChangeLog) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ResourceChangeLog) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ResourceChangeLog) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ResourceChangeLog) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ResourceChangeLog) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *ResourceChangeLog) GetRequestId() uint64 {
	if m != nil {
		return m.RequestId
	}
	return uint64(0)
}

func (m *ResourceChangeLog) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ResourceChangeLog) GetAuthentication() *audit_common.Authentication {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *ResourceChangeLog) GetService() *audit_common.ServiceData {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ResourceChangeLog) GetResource() *ResourceChangeLog_ResourceChange {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *ResourceChangeLog) GetTransaction() *ResourceChangeLog_TransactionInfo {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *ResourceChangeLog) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ResourceChangeLog"))
	}
	m.Name = fv
}

func (m *ResourceChangeLog) SetScope(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Scope", "ResourceChangeLog"))
	}
	m.Scope = fv
}

func (m *ResourceChangeLog) SetRequestId(fv uint64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RequestId", "ResourceChangeLog"))
	}
	m.RequestId = fv
}

func (m *ResourceChangeLog) SetTimestamp(fv *timestamp.Timestamp) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Timestamp", "ResourceChangeLog"))
	}
	m.Timestamp = fv
}

func (m *ResourceChangeLog) SetAuthentication(fv *audit_common.Authentication) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Authentication", "ResourceChangeLog"))
	}
	m.Authentication = fv
}

func (m *ResourceChangeLog) SetService(fv *audit_common.ServiceData) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Service", "ResourceChangeLog"))
	}
	m.Service = fv
}

func (m *ResourceChangeLog) SetResource(fv *ResourceChangeLog_ResourceChange) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Resource", "ResourceChangeLog"))
	}
	m.Resource = fv
}

func (m *ResourceChangeLog) SetTransaction(fv *ResourceChangeLog_TransactionInfo) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Transaction", "ResourceChangeLog"))
	}
	m.Transaction = fv
}

// Description of change on the resource
type ResourceChangeLog_ResourceChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Fully qualified name of the resource (eg. "RoleBinding/Public")
	// that has changed from this request (if successful)
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Name of the resource type for example "RoleBinding".
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty" firestore:"type"`
	// Action on the resource
	Action ResourceChangeLog_ResourceChange_Action `protobuf:"varint,3,opt,name=action,proto3,enum=ntt.audit.v1alpha.ResourceChangeLog_ResourceChange_Action" json:"action,omitempty" firestore:"action"`
	// State of the resource before change.
	// It is empty if action is CREATE
	Pre *audit_common.ObjectState `protobuf:"bytes,4,opt,name=pre,proto3" json:"pre,omitempty" firestore:"pre"`
	// State of the resource after change.
	// It is empty if action is DELETE
	Post *audit_common.ObjectState `protobuf:"bytes,5,opt,name=post,proto3" json:"post,omitempty" firestore:"post"`
}

func (m *ResourceChangeLog_ResourceChange) Reset() {
	*m = ResourceChangeLog_ResourceChange{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha_resource_change_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ResourceChangeLog_ResourceChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ResourceChangeLog_ResourceChange) ProtoMessage() {}

func (m *ResourceChangeLog_ResourceChange) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha_resource_change_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ResourceChangeLog_ResourceChange) GotenMessage() {}

// Deprecated, Use ResourceChangeLog_ResourceChange.ProtoReflect.Descriptor instead.
func (*ResourceChangeLog_ResourceChange) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha_resource_change_log_proto_rawDescGZIP(), []int{0, 0}
}

func (m *ResourceChangeLog_ResourceChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ResourceChangeLog_ResourceChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ResourceChangeLog_ResourceChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ResourceChangeLog_ResourceChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ResourceChangeLog_ResourceChange) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceChangeLog_ResourceChange) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ResourceChangeLog_ResourceChange) GetAction() ResourceChangeLog_ResourceChange_Action {
	if m != nil {
		return m.Action
	}
	return ResourceChangeLog_ResourceChange_UNDEFINED
}

func (m *ResourceChangeLog_ResourceChange) GetPre() *audit_common.ObjectState {
	if m != nil {
		return m.Pre
	}
	return nil
}

func (m *ResourceChangeLog_ResourceChange) GetPost() *audit_common.ObjectState {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *ResourceChangeLog_ResourceChange) SetName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ResourceChangeLog_ResourceChange"))
	}
	m.Name = fv
}

func (m *ResourceChangeLog_ResourceChange) SetType(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Type", "ResourceChangeLog_ResourceChange"))
	}
	m.Type = fv
}

func (m *ResourceChangeLog_ResourceChange) SetAction(fv ResourceChangeLog_ResourceChange_Action) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Action", "ResourceChangeLog_ResourceChange"))
	}
	m.Action = fv
}

func (m *ResourceChangeLog_ResourceChange) SetPre(fv *audit_common.ObjectState) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Pre", "ResourceChangeLog_ResourceChange"))
	}
	m.Pre = fv
}

func (m *ResourceChangeLog_ResourceChange) SetPost(fv *audit_common.ObjectState) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Post", "ResourceChangeLog_ResourceChange"))
	}
	m.Post = fv
}

// Information about transaction where change
// has been executed
type ResourceChangeLog_TransactionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// unique identifier of the transaction.
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty" firestore:"identifier"`
	// Indicator of try counter. If transaction has been
	// concluded at first try, try_counter will be 1. If
	// on the second try, then number will be 2 (etc).
	TryCounter int32 `protobuf:"varint,2,opt,name=try_counter,json=tryCounter,proto3" json:"try_counter,omitempty" firestore:"tryCounter"`
	// State of the transaction.
	State ResourceChangeLog_TransactionInfo_State `protobuf:"varint,3,opt,name=state,proto3,enum=ntt.audit.v1alpha.ResourceChangeLog_TransactionInfo_State" json:"state,omitempty" firestore:"state"`
}

func (m *ResourceChangeLog_TransactionInfo) Reset() {
	*m = ResourceChangeLog_TransactionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_audit_proto_v1alpha_resource_change_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ResourceChangeLog_TransactionInfo) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ResourceChangeLog_TransactionInfo) ProtoMessage() {}

func (m *ResourceChangeLog_TransactionInfo) ProtoReflect() preflect.Message {
	mi := &edgelq_audit_proto_v1alpha_resource_change_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ResourceChangeLog_TransactionInfo) GotenMessage() {}

// Deprecated, Use ResourceChangeLog_TransactionInfo.ProtoReflect.Descriptor instead.
func (*ResourceChangeLog_TransactionInfo) Descriptor() ([]byte, []int) {
	return edgelq_audit_proto_v1alpha_resource_change_log_proto_rawDescGZIP(), []int{0, 1}
}

func (m *ResourceChangeLog_TransactionInfo) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ResourceChangeLog_TransactionInfo) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ResourceChangeLog_TransactionInfo) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ResourceChangeLog_TransactionInfo) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ResourceChangeLog_TransactionInfo) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *ResourceChangeLog_TransactionInfo) GetTryCounter() int32 {
	if m != nil {
		return m.TryCounter
	}
	return int32(0)
}

func (m *ResourceChangeLog_TransactionInfo) GetState() ResourceChangeLog_TransactionInfo_State {
	if m != nil {
		return m.State
	}
	return ResourceChangeLog_TransactionInfo_UNDEFINED
}

func (m *ResourceChangeLog_TransactionInfo) SetIdentifier(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Identifier", "ResourceChangeLog_TransactionInfo"))
	}
	m.Identifier = fv
}

func (m *ResourceChangeLog_TransactionInfo) SetTryCounter(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TryCounter", "ResourceChangeLog_TransactionInfo"))
	}
	m.TryCounter = fv
}

func (m *ResourceChangeLog_TransactionInfo) SetState(fv ResourceChangeLog_TransactionInfo_State) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "State", "ResourceChangeLog_TransactionInfo"))
	}
	m.State = fv
}

var edgelq_audit_proto_v1alpha_resource_change_log_proto preflect.FileDescriptor

var edgelq_audit_proto_v1alpha_resource_change_log_proto_rawDesc = []byte{
	0x0a, 0x34, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x0a, 0x0a,
	0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c,
	0x6f, 0x67, 0x12, 0x2d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x19, 0xb2, 0xda, 0x21, 0x15, 0x0a, 0x13, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x49, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f,
	0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0xaf,
	0x02, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x52, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a,
	0x03, 0x70, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x03, 0x70, 0x72, 0x65, 0x12,
	0x32, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x04, 0x70,
	0x6f, 0x73, 0x74, 0x22, 0x3b, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x0a,
	0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03,
	0x1a, 0xef, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x72, 0x79, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x50, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x49, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x11, 0x0a, 0x0d, 0x50, 0x52, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x42, 0x41, 0x43, 0x4b,
	0x10, 0x03, 0x3a, 0xf0, 0x02, 0xea, 0x41, 0xd2, 0x01, 0x0a, 0x22, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x28, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67,
	0x73, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x7d, 0x12, 0x3b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x2f, 0x7b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x6c, 0x6f, 0x67, 0x7d, 0x12, 0x45, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x4c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x7d, 0x92, 0xd9, 0x21, 0x7f, 0x0a,
	0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c,
	0x6f, 0x67, 0x73, 0x12, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x1a, 0x12, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x6f, 0x6e, 0x65, 0x1a, 0x17, 0x6e, 0x74, 0x74,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x1a, 0x1c, 0x6e, 0x74, 0x74, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x28, 0x01, 0x5a, 0x02, 0x08, 0x01, 0xfa, 0xde,
	0x21, 0x13, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4c, 0x6f, 0x67, 0x42, 0x98, 0x04, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02,
	0x5f, 0x0a, 0x19, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x42, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61,
	0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67,
	0x92, 0x8c, 0xd1, 0x02, 0x69, 0x0a, 0x1e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x0a, 0x18,
	0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x62,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x16, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0xd2, 0x84,
	0xd1, 0x02, 0x47, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x73, 0x12, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x80, 0xd1, 0x02, 0x61, 0x0a,
	0x1a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x43, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_audit_proto_v1alpha_resource_change_log_proto_rawDescOnce sync.Once
	edgelq_audit_proto_v1alpha_resource_change_log_proto_rawDescData = edgelq_audit_proto_v1alpha_resource_change_log_proto_rawDesc
)

func edgelq_audit_proto_v1alpha_resource_change_log_proto_rawDescGZIP() []byte {
	edgelq_audit_proto_v1alpha_resource_change_log_proto_rawDescOnce.Do(func() {
		edgelq_audit_proto_v1alpha_resource_change_log_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_audit_proto_v1alpha_resource_change_log_proto_rawDescData)
	})
	return edgelq_audit_proto_v1alpha_resource_change_log_proto_rawDescData
}

var edgelq_audit_proto_v1alpha_resource_change_log_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var edgelq_audit_proto_v1alpha_resource_change_log_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var edgelq_audit_proto_v1alpha_resource_change_log_proto_goTypes = []interface{}{
	(ResourceChangeLog_ResourceChange_Action)(0), // 0: ntt.audit.v1alpha.ResourceChangeLog_ResourceChange_Action
	(ResourceChangeLog_TransactionInfo_State)(0), // 1: ntt.audit.v1alpha.ResourceChangeLog_TransactionInfo_State
	(*ResourceChangeLog)(nil),                    // 2: ntt.audit.v1alpha.ResourceChangeLog
	(*ResourceChangeLog_ResourceChange)(nil),     // 3: ntt.audit.v1alpha.ResourceChangeLog.ResourceChange
	(*ResourceChangeLog_TransactionInfo)(nil),    // 4: ntt.audit.v1alpha.ResourceChangeLog.TransactionInfo
	(*timestamp.Timestamp)(nil),                  // 5: google.protobuf.Timestamp
	(*audit_common.Authentication)(nil),          // 6: ntt.audit.v1alpha.Authentication
	(*audit_common.ServiceData)(nil),             // 7: ntt.audit.v1alpha.ServiceData
	(*audit_common.ObjectState)(nil),             // 8: ntt.audit.v1alpha.ObjectState
}
var edgelq_audit_proto_v1alpha_resource_change_log_proto_depIdxs = []int32{
	5, // 0: ntt.audit.v1alpha.ResourceChangeLog.timestamp:type_name -> google.protobuf.Timestamp
	6, // 1: ntt.audit.v1alpha.ResourceChangeLog.authentication:type_name -> ntt.audit.v1alpha.Authentication
	7, // 2: ntt.audit.v1alpha.ResourceChangeLog.service:type_name -> ntt.audit.v1alpha.ServiceData
	3, // 3: ntt.audit.v1alpha.ResourceChangeLog.resource:type_name -> ntt.audit.v1alpha.ResourceChangeLog.ResourceChange
	4, // 4: ntt.audit.v1alpha.ResourceChangeLog.transaction:type_name -> ntt.audit.v1alpha.ResourceChangeLog.TransactionInfo
	0, // 5: ntt.audit.v1alpha.ResourceChangeLog.ResourceChange.action:type_name -> ntt.audit.v1alpha.ResourceChangeLog_ResourceChange_Action
	8, // 6: ntt.audit.v1alpha.ResourceChangeLog.ResourceChange.pre:type_name -> ntt.audit.v1alpha.ObjectState
	8, // 7: ntt.audit.v1alpha.ResourceChangeLog.ResourceChange.post:type_name -> ntt.audit.v1alpha.ObjectState
	1, // 8: ntt.audit.v1alpha.ResourceChangeLog.TransactionInfo.state:type_name -> ntt.audit.v1alpha.ResourceChangeLog_TransactionInfo_State
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { edgelq_audit_proto_v1alpha_resource_change_log_proto_init() }
func edgelq_audit_proto_v1alpha_resource_change_log_proto_init() {
	if edgelq_audit_proto_v1alpha_resource_change_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_audit_proto_v1alpha_resource_change_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceChangeLog); i {
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
		edgelq_audit_proto_v1alpha_resource_change_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceChangeLog_ResourceChange); i {
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
		edgelq_audit_proto_v1alpha_resource_change_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceChangeLog_TransactionInfo); i {
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
			RawDescriptor: edgelq_audit_proto_v1alpha_resource_change_log_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_audit_proto_v1alpha_resource_change_log_proto_goTypes,
		DependencyIndexes: edgelq_audit_proto_v1alpha_resource_change_log_proto_depIdxs,
		EnumInfos:         edgelq_audit_proto_v1alpha_resource_change_log_proto_enumTypes,
		MessageInfos:      edgelq_audit_proto_v1alpha_resource_change_log_proto_msgTypes,
	}.Build()
	edgelq_audit_proto_v1alpha_resource_change_log_proto = out.File
	edgelq_audit_proto_v1alpha_resource_change_log_proto_rawDesc = nil
	edgelq_audit_proto_v1alpha_resource_change_log_proto_goTypes = nil
	edgelq_audit_proto_v1alpha_resource_change_log_proto_depIdxs = nil
}
