// Code generated by protoc-gen-goten-go
// File: edgelq/alerting/proto/v1/edge_watch_service_custom.proto
// DO NOT EDIT!!!

package edge_watch_service_client

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
	alert "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/alert"
	log_condition "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/log_condition"
	policy "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/policy"
	ts_condition "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/ts_condition"
	ts_entry "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/ts_entry"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
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
	_ = &alert.Alert{}
	_ = &log_condition.LogCondition{}
	_ = &policy.Policy{}
	_ = &ts_condition.TsCondition{}
	_ = &ts_entry.TsEntry{}
	_ = &iam_project.Project{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A request message of the WatchAlertData method.
// This is special combined watch of 5 streams dedicated specifically
// for edge alerter component. It reduces number of watch streams and
// therefore connections maintained by server (and slightly agent).
type WatchAlertDataRequest struct {
	state                protoimpl.MessageState
	sizeCache            protoimpl.SizeCache
	unknownFields        protoimpl.UnknownFields
	Project              *iam_project.Name `protobuf:"bytes,1,opt,customtype=Name,name=project,proto3" json:"project,omitempty"`
	AlertingResourceName string            `protobuf:"bytes,2,opt,name=alerting_resource_name,json=alertingResourceName,proto3" json:"alerting_resource_name,omitempty"`
	AlertingResourceKind string            `protobuf:"bytes,3,opt,name=alerting_resource_kind,json=alertingResourceKind,proto3" json:"alerting_resource_kind,omitempty"`
}

func (m *WatchAlertDataRequest) Reset() {
	*m = WatchAlertDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_alerting_proto_v1_edge_watch_service_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *WatchAlertDataRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*WatchAlertDataRequest) ProtoMessage() {}

func (m *WatchAlertDataRequest) ProtoReflect() preflect.Message {
	mi := &edgelq_alerting_proto_v1_edge_watch_service_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*WatchAlertDataRequest) GotenMessage() {}

// Deprecated, Use WatchAlertDataRequest.ProtoReflect.Descriptor instead.
func (*WatchAlertDataRequest) Descriptor() ([]byte, []int) {
	return edgelq_alerting_proto_v1_edge_watch_service_custom_proto_rawDescGZIP(), []int{0}
}

func (m *WatchAlertDataRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *WatchAlertDataRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *WatchAlertDataRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *WatchAlertDataRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *WatchAlertDataRequest) GetProject() *iam_project.Name {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *WatchAlertDataRequest) GetAlertingResourceName() string {
	if m != nil {
		return m.AlertingResourceName
	}
	return ""
}

func (m *WatchAlertDataRequest) GetAlertingResourceKind() string {
	if m != nil {
		return m.AlertingResourceKind
	}
	return ""
}

func (m *WatchAlertDataRequest) SetProject(fv *iam_project.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Project", "WatchAlertDataRequest"))
	}
	m.Project = fv
}

func (m *WatchAlertDataRequest) SetAlertingResourceName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AlertingResourceName", "WatchAlertDataRequest"))
	}
	m.AlertingResourceName = fv
}

func (m *WatchAlertDataRequest) SetAlertingResourceKind(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AlertingResourceKind", "WatchAlertDataRequest"))
	}
	m.AlertingResourceKind = fv
}

// A response message of the WatchAlertData method.
type WatchAlertDataResponse struct {
	state                protoimpl.MessageState
	sizeCache            protoimpl.SizeCache
	unknownFields        protoimpl.UnknownFields
	AlertsToAdd          []*alert.Alert                `protobuf:"bytes,1,rep,name=alerts_to_add,json=alertsToAdd,proto3" json:"alerts_to_add,omitempty"`
	TsEntriesToAdd       []*ts_entry.TsEntry           `protobuf:"bytes,2,rep,name=ts_entries_to_add,json=tsEntriesToAdd,proto3" json:"ts_entries_to_add,omitempty"`
	PoliciesToAdd        []*policy.Policy              `protobuf:"bytes,3,rep,name=policies_to_add,json=policiesToAdd,proto3" json:"policies_to_add,omitempty"`
	TsCndsToAdd          []*ts_condition.TsCondition   `protobuf:"bytes,4,rep,name=ts_cnds_to_add,json=tsCndsToAdd,proto3" json:"ts_cnds_to_add,omitempty"`
	LogCndsToAdd         []*log_condition.LogCondition `protobuf:"bytes,5,rep,name=log_cnds_to_add,json=logCndsToAdd,proto3" json:"log_cnds_to_add,omitempty"`
	AlertNamesToRemove   []*alert.Name                 `protobuf:"bytes,6,rep,customtype=Name,name=alert_names_to_remove,json=alertNamesToRemove,proto3" json:"alert_names_to_remove,omitempty"`
	TsEntryNamesToRemove []*ts_entry.Name              `protobuf:"bytes,7,rep,customtype=Name,name=ts_entry_names_to_remove,json=tsEntryNamesToRemove,proto3" json:"ts_entry_names_to_remove,omitempty"`
	PolicyNamesToRemove  []*policy.Name                `protobuf:"bytes,8,rep,customtype=Name,name=policy_names_to_remove,json=policyNamesToRemove,proto3" json:"policy_names_to_remove,omitempty"`
	TsCndNamesToRemove   []*ts_condition.Name          `protobuf:"bytes,9,rep,customtype=Name,name=ts_cnd_names_to_remove,json=tsCndNamesToRemove,proto3" json:"ts_cnd_names_to_remove,omitempty"`
	LogCndNamesToRemove  []*log_condition.Name         `protobuf:"bytes,10,rep,customtype=Name,name=log_cnd_names_to_remove,json=logCndNamesToRemove,proto3" json:"log_cnd_names_to_remove,omitempty"`
	ResetAlerts          bool                          `protobuf:"varint,11,opt,name=reset_alerts,json=resetAlerts,proto3" json:"reset_alerts,omitempty"`
	ResetTsEntries       bool                          `protobuf:"varint,12,opt,name=reset_ts_entries,json=resetTsEntries,proto3" json:"reset_ts_entries,omitempty"`
	ResetPolicies        bool                          `protobuf:"varint,13,opt,name=reset_policies,json=resetPolicies,proto3" json:"reset_policies,omitempty"`
	ResetTsCnds          bool                          `protobuf:"varint,14,opt,name=reset_ts_cnds,json=resetTsCnds,proto3" json:"reset_ts_cnds,omitempty"`
	ResetLogCnds         bool                          `protobuf:"varint,15,opt,name=reset_log_cnds,json=resetLogCnds,proto3" json:"reset_log_cnds,omitempty"`
}

func (m *WatchAlertDataResponse) Reset() {
	*m = WatchAlertDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &edgelq_alerting_proto_v1_edge_watch_service_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *WatchAlertDataResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*WatchAlertDataResponse) ProtoMessage() {}

func (m *WatchAlertDataResponse) ProtoReflect() preflect.Message {
	mi := &edgelq_alerting_proto_v1_edge_watch_service_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*WatchAlertDataResponse) GotenMessage() {}

// Deprecated, Use WatchAlertDataResponse.ProtoReflect.Descriptor instead.
func (*WatchAlertDataResponse) Descriptor() ([]byte, []int) {
	return edgelq_alerting_proto_v1_edge_watch_service_custom_proto_rawDescGZIP(), []int{1}
}

func (m *WatchAlertDataResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *WatchAlertDataResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *WatchAlertDataResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *WatchAlertDataResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *WatchAlertDataResponse) GetAlertsToAdd() []*alert.Alert {
	if m != nil {
		return m.AlertsToAdd
	}
	return nil
}

func (m *WatchAlertDataResponse) GetTsEntriesToAdd() []*ts_entry.TsEntry {
	if m != nil {
		return m.TsEntriesToAdd
	}
	return nil
}

func (m *WatchAlertDataResponse) GetPoliciesToAdd() []*policy.Policy {
	if m != nil {
		return m.PoliciesToAdd
	}
	return nil
}

func (m *WatchAlertDataResponse) GetTsCndsToAdd() []*ts_condition.TsCondition {
	if m != nil {
		return m.TsCndsToAdd
	}
	return nil
}

func (m *WatchAlertDataResponse) GetLogCndsToAdd() []*log_condition.LogCondition {
	if m != nil {
		return m.LogCndsToAdd
	}
	return nil
}

func (m *WatchAlertDataResponse) GetAlertNamesToRemove() []*alert.Name {
	if m != nil {
		return m.AlertNamesToRemove
	}
	return nil
}

func (m *WatchAlertDataResponse) GetTsEntryNamesToRemove() []*ts_entry.Name {
	if m != nil {
		return m.TsEntryNamesToRemove
	}
	return nil
}

func (m *WatchAlertDataResponse) GetPolicyNamesToRemove() []*policy.Name {
	if m != nil {
		return m.PolicyNamesToRemove
	}
	return nil
}

func (m *WatchAlertDataResponse) GetTsCndNamesToRemove() []*ts_condition.Name {
	if m != nil {
		return m.TsCndNamesToRemove
	}
	return nil
}

func (m *WatchAlertDataResponse) GetLogCndNamesToRemove() []*log_condition.Name {
	if m != nil {
		return m.LogCndNamesToRemove
	}
	return nil
}

func (m *WatchAlertDataResponse) GetResetAlerts() bool {
	if m != nil {
		return m.ResetAlerts
	}
	return false
}

func (m *WatchAlertDataResponse) GetResetTsEntries() bool {
	if m != nil {
		return m.ResetTsEntries
	}
	return false
}

func (m *WatchAlertDataResponse) GetResetPolicies() bool {
	if m != nil {
		return m.ResetPolicies
	}
	return false
}

func (m *WatchAlertDataResponse) GetResetTsCnds() bool {
	if m != nil {
		return m.ResetTsCnds
	}
	return false
}

func (m *WatchAlertDataResponse) GetResetLogCnds() bool {
	if m != nil {
		return m.ResetLogCnds
	}
	return false
}

func (m *WatchAlertDataResponse) SetAlertsToAdd(fv []*alert.Alert) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AlertsToAdd", "WatchAlertDataResponse"))
	}
	m.AlertsToAdd = fv
}

func (m *WatchAlertDataResponse) SetTsEntriesToAdd(fv []*ts_entry.TsEntry) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TsEntriesToAdd", "WatchAlertDataResponse"))
	}
	m.TsEntriesToAdd = fv
}

func (m *WatchAlertDataResponse) SetPoliciesToAdd(fv []*policy.Policy) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PoliciesToAdd", "WatchAlertDataResponse"))
	}
	m.PoliciesToAdd = fv
}

func (m *WatchAlertDataResponse) SetTsCndsToAdd(fv []*ts_condition.TsCondition) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TsCndsToAdd", "WatchAlertDataResponse"))
	}
	m.TsCndsToAdd = fv
}

func (m *WatchAlertDataResponse) SetLogCndsToAdd(fv []*log_condition.LogCondition) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "LogCndsToAdd", "WatchAlertDataResponse"))
	}
	m.LogCndsToAdd = fv
}

func (m *WatchAlertDataResponse) SetAlertNamesToRemove(fv []*alert.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AlertNamesToRemove", "WatchAlertDataResponse"))
	}
	m.AlertNamesToRemove = fv
}

func (m *WatchAlertDataResponse) SetTsEntryNamesToRemove(fv []*ts_entry.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TsEntryNamesToRemove", "WatchAlertDataResponse"))
	}
	m.TsEntryNamesToRemove = fv
}

func (m *WatchAlertDataResponse) SetPolicyNamesToRemove(fv []*policy.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PolicyNamesToRemove", "WatchAlertDataResponse"))
	}
	m.PolicyNamesToRemove = fv
}

func (m *WatchAlertDataResponse) SetTsCndNamesToRemove(fv []*ts_condition.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TsCndNamesToRemove", "WatchAlertDataResponse"))
	}
	m.TsCndNamesToRemove = fv
}

func (m *WatchAlertDataResponse) SetLogCndNamesToRemove(fv []*log_condition.Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "LogCndNamesToRemove", "WatchAlertDataResponse"))
	}
	m.LogCndNamesToRemove = fv
}

func (m *WatchAlertDataResponse) SetResetAlerts(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ResetAlerts", "WatchAlertDataResponse"))
	}
	m.ResetAlerts = fv
}

func (m *WatchAlertDataResponse) SetResetTsEntries(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ResetTsEntries", "WatchAlertDataResponse"))
	}
	m.ResetTsEntries = fv
}

func (m *WatchAlertDataResponse) SetResetPolicies(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ResetPolicies", "WatchAlertDataResponse"))
	}
	m.ResetPolicies = fv
}

func (m *WatchAlertDataResponse) SetResetTsCnds(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ResetTsCnds", "WatchAlertDataResponse"))
	}
	m.ResetTsCnds = fv
}

func (m *WatchAlertDataResponse) SetResetLogCnds(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ResetLogCnds", "WatchAlertDataResponse"))
	}
	m.ResetLogCnds = fv
}

var edgelq_alerting_proto_v1_edge_watch_service_custom_proto preflect.FileDescriptor

var edgelq_alerting_proto_v1_edge_watch_service_custom_proto_rawDesc = []byte{
	0x0a, 0x38, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x5f,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6e, 0x74, 0x74, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x25, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x15, 0x57, 0x61, 0x74, 0x63, 0x68, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1e, 0xb2, 0xda, 0x21, 0x1a, 0x0a, 0x18, 0x0a, 0x16, 0x69, 0x61, 0x6d, 0x2e, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34,
	0x0a, 0x16, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4b, 0x69, 0x6e, 0x64, 0x22, 0x86, 0x07, 0x0a, 0x16, 0x57, 0x61, 0x74, 0x63, 0x68, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3a, 0x0a, 0x0d, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x0b,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x12, 0x43, 0x0a, 0x11, 0x74,
	0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64,
	0x12, 0x3f, 0x0a, 0x0f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x5f,
	0x61, 0x64, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x0d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x54, 0x6f, 0x41, 0x64,
	0x64, 0x12, 0x41, 0x0a, 0x0e, 0x74, 0x73, 0x5f, 0x63, 0x6e, 0x64, 0x73, 0x5f, 0x74, 0x6f, 0x5f,
	0x61, 0x64, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x73, 0x43, 0x6e, 0x64, 0x73, 0x54,
	0x6f, 0x41, 0x64, 0x64, 0x12, 0x44, 0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6e, 0x64, 0x73,
	0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6c, 0x6f,
	0x67, 0x43, 0x6e, 0x64, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x12, 0x40, 0x0a, 0x15, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0d, 0xb2, 0xda, 0x21, 0x09, 0x0a,
	0x07, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x12, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x47, 0x0a, 0x18,
	0x74, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x5f, 0x74,
	0x6f, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0f,
	0xb2, 0xda, 0x21, 0x0b, 0x0a, 0x09, 0x0a, 0x07, 0x54, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x14, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x54, 0x6f, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x43, 0x0a, 0x16, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0e, 0xb2, 0xda, 0x21, 0x0a, 0x0a, 0x08, 0x0a, 0x06, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x13, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x47, 0x0a, 0x16, 0x74, 0x73,
	0x5f, 0x63, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x42, 0x13, 0xb2, 0xda, 0x21, 0x0f,
	0x0a, 0x0d, 0x0a, 0x0b, 0x54, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x12, 0x74, 0x73, 0x43, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x12, 0x4a, 0x0a, 0x17, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6e, 0x64, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x09, 0x42, 0x14, 0xb2, 0xda, 0x21, 0x10, 0x0a, 0x0e, 0x0a, 0x0c, 0x4c, 0x6f,
	0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x6c, 0x6f, 0x67, 0x43,
	0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x74, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x73, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x54, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x73, 0x5f,
	0x63, 0x6e, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x54, 0x73, 0x43, 0x6e, 0x64, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6e, 0x64, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x43, 0x6e, 0x64, 0x73, 0x42, 0x97, 0x01,
	0xe8, 0xde, 0x21, 0x00, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x1b, 0x45, 0x64,
	0x67, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x5a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x5f,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x65, 0x64,
	0x67, 0x65, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	edgelq_alerting_proto_v1_edge_watch_service_custom_proto_rawDescOnce sync.Once
	edgelq_alerting_proto_v1_edge_watch_service_custom_proto_rawDescData = edgelq_alerting_proto_v1_edge_watch_service_custom_proto_rawDesc
)

func edgelq_alerting_proto_v1_edge_watch_service_custom_proto_rawDescGZIP() []byte {
	edgelq_alerting_proto_v1_edge_watch_service_custom_proto_rawDescOnce.Do(func() {
		edgelq_alerting_proto_v1_edge_watch_service_custom_proto_rawDescData = protoimpl.X.CompressGZIP(edgelq_alerting_proto_v1_edge_watch_service_custom_proto_rawDescData)
	})
	return edgelq_alerting_proto_v1_edge_watch_service_custom_proto_rawDescData
}

var edgelq_alerting_proto_v1_edge_watch_service_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var edgelq_alerting_proto_v1_edge_watch_service_custom_proto_goTypes = []interface{}{
	(*WatchAlertDataRequest)(nil),      // 0: ntt.alerting.v1.WatchAlertDataRequest
	(*WatchAlertDataResponse)(nil),     // 1: ntt.alerting.v1.WatchAlertDataResponse
	(*alert.Alert)(nil),                // 2: ntt.alerting.v1.Alert
	(*ts_entry.TsEntry)(nil),           // 3: ntt.alerting.v1.TsEntry
	(*policy.Policy)(nil),              // 4: ntt.alerting.v1.Policy
	(*ts_condition.TsCondition)(nil),   // 5: ntt.alerting.v1.TsCondition
	(*log_condition.LogCondition)(nil), // 6: ntt.alerting.v1.LogCondition
}
var edgelq_alerting_proto_v1_edge_watch_service_custom_proto_depIdxs = []int32{
	2, // 0: ntt.alerting.v1.WatchAlertDataResponse.alerts_to_add:type_name -> ntt.alerting.v1.Alert
	3, // 1: ntt.alerting.v1.WatchAlertDataResponse.ts_entries_to_add:type_name -> ntt.alerting.v1.TsEntry
	4, // 2: ntt.alerting.v1.WatchAlertDataResponse.policies_to_add:type_name -> ntt.alerting.v1.Policy
	5, // 3: ntt.alerting.v1.WatchAlertDataResponse.ts_cnds_to_add:type_name -> ntt.alerting.v1.TsCondition
	6, // 4: ntt.alerting.v1.WatchAlertDataResponse.log_cnds_to_add:type_name -> ntt.alerting.v1.LogCondition
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { edgelq_alerting_proto_v1_edge_watch_service_custom_proto_init() }
func edgelq_alerting_proto_v1_edge_watch_service_custom_proto_init() {
	if edgelq_alerting_proto_v1_edge_watch_service_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		edgelq_alerting_proto_v1_edge_watch_service_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchAlertDataRequest); i {
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
		edgelq_alerting_proto_v1_edge_watch_service_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchAlertDataResponse); i {
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
			RawDescriptor: edgelq_alerting_proto_v1_edge_watch_service_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           edgelq_alerting_proto_v1_edge_watch_service_custom_proto_goTypes,
		DependencyIndexes: edgelq_alerting_proto_v1_edge_watch_service_custom_proto_depIdxs,
		MessageInfos:      edgelq_alerting_proto_v1_edge_watch_service_custom_proto_msgTypes,
	}.Build()
	edgelq_alerting_proto_v1_edge_watch_service_custom_proto = out.File
	edgelq_alerting_proto_v1_edge_watch_service_custom_proto_rawDesc = nil
	edgelq_alerting_proto_v1_edge_watch_service_custom_proto_goTypes = nil
	edgelq_alerting_proto_v1_edge_watch_service_custom_proto_depIdxs = nil
}
