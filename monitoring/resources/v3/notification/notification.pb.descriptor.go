// Code generated by protoc-gen-goten-resource
// Resource: Notification
// DO NOT EDIT!!!

package notification

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	alert "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alert"
	alerting_condition "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_condition"
	alerting_policy "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_policy"
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/common"
	notification_channel "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/notification_channel"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(gotenobject.FieldPath)
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &alert.Alert{}
	_ = &alerting_condition.AlertingCondition{}
	_ = &alerting_policy.AlertingPolicy{}
	_ = &common.LabelDescriptor{}
	_ = &notification_channel.NotificationChannel{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

var (
	descriptor *Descriptor
)

func (r *Notification) GetRawName() gotenresource.Name {
	return r.GetName()
}

func (r *Notification) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (r *Notification) EnsureMetadata() *meta.Meta {
	if r.Metadata == nil {
		r.Metadata = &meta.Meta{}
	}
	if r.Metadata.Lifecycle == nil {
		r.Metadata.Lifecycle = &meta.Lifecycle{}
	}
	return r.Metadata
}

type Descriptor struct {
	nameDescriptor *gotenresource.NameDescriptor
	typeName       *gotenresource.TypeName
}

func GetDescriptor() *Descriptor {
	return descriptor
}

func (d *Descriptor) NewResource() gotenresource.Resource {
	return &Notification{}
}

func (d *Descriptor) NewResourceName() gotenresource.Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewGetQuery() gotenresource.GetQuery {
	return &GetQuery{}
}

func (d *Descriptor) NewListQuery() gotenresource.ListQuery {
	return &ListQuery{}
}

func (d *Descriptor) NewSearchQuery() gotenresource.SearchQuery {
	return &SearchQuery{}
}

func (d *Descriptor) NewWatchQuery() gotenresource.WatchQuery {
	return &WatchQuery{}
}

func (d *Descriptor) NewResourceCursor() gotenresource.Cursor {
	return &PagerCursor{}
}

func (d *Descriptor) NewResourceFilter() gotenresource.Filter {
	return &Filter{}
}

func (d *Descriptor) NewResourceOrderBy() gotenresource.OrderBy {
	return &OrderBy{}
}

func (d *Descriptor) NewResourcePager() gotenresource.PagerQuery {
	return MakePagerQuery(nil, nil, 100, true)
}

func (d *Descriptor) NewResourceFieldMask() gotenobject.FieldMask {
	return &Notification_FieldMask{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return &NotificationChange{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(NotificationList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(NotificationChangeList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(NotificationNameList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(NotificationReferenceList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return make(NotificationParentNameList, size, reserved)
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return make(NotificationParentReferenceList, size, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(NotificationMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(NotificationChangeMap, reserved)
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) CanBeParentless() bool {
	return false
}

func (d *Descriptor) GetParentResDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		alerting_policy.GetDescriptor(),
	}
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseNotification_FieldPath(raw)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}

func (d *Descriptor) SupportsMetadata() bool {
	return true
}

func initNotificationDescriptor() {
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"Notification", "Notifications", "monitoring.edgelq.com", "v3"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&Notification_FieldTerminalPath{selector: Notification_FieldPathSelectorName},
			"pattern", "notificationId",
			[]string{"projectId", "regionId", "alertingPolicyId"},
			[]gotenresource.NamePattern{NamePattern_Project_Region_AlertingPolicy}),
	}
	gotenresource.GetRegistry().RegisterDescriptor(descriptor)
}

func init() {
	initNotificationDescriptor()
}
