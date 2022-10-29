// Code generated by protoc-gen-goten-resource
// Resource: ActivityLog
// DO NOT EDIT!!!

package activity_log

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	audit_common "github.com/cloudwan/edgelq-sdk/audit/common/v1alpha2"
	rpc "github.com/cloudwan/edgelq-sdk/common/rpc"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// ensure the imports are used
var (
	_ = new(gotenobject.FieldPath)
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &audit_common.Authentication{}
	_ = &rpc.Status{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &any.Any{}
	_ = &field_mask.FieldMask{}
	_ = &timestamp.Timestamp{}
)

var (
	descriptor *Descriptor
)

type Descriptor struct {
	nameDescriptor *gotenresource.NameDescriptor
	typeName       *gotenresource.TypeName
}

func GetDescriptor() *Descriptor {
	return descriptor
}

func (d *Descriptor) NewResource() gotenresource.Resource {
	return &ActivityLog{}
}

func (d *Descriptor) NewResourceName() gotenresource.Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewGetQuery() gotenresource.GetQuery {
	return nil
}

func (d *Descriptor) NewListQuery() gotenresource.ListQuery {
	return nil
}

func (d *Descriptor) NewSearchQuery() gotenresource.SearchQuery {
	return nil
}

func (d *Descriptor) NewWatchQuery() gotenresource.WatchQuery {
	return nil
}

func (d *Descriptor) NewResourceCursor() gotenresource.Cursor {
	return nil
}

func (d *Descriptor) NewResourceFilter() gotenresource.Filter {
	return &Filter{}
}

func (d *Descriptor) NewResourceOrderBy() gotenresource.OrderBy {
	return nil
}

func (d *Descriptor) NewResourcePager() gotenresource.PagerQuery {
	return nil
}

func (d *Descriptor) NewResourceFieldMask() gotenobject.FieldMask {
	return &ActivityLog_FieldMask{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return nil
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return nil
}

func (d *Descriptor) NewSearchQueryResultSnapshot() gotenresource.SearchQueryResultSnapshot {
	return nil
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return nil
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(ActivityLogList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return nil
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(ActivityLogNameList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(ActivityLogReferenceList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return make(ActivityLogParentNameList, size, reserved)
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return make(ActivityLogParentReferenceList, size, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(ActivityLogMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return nil
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseActivityLog_FieldPath(raw)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}

func initActivityLogDescriptor() {
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"ActivityLog", "ActivityLogs", "audit.edgelq.com", "v1alpha2"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&ActivityLog_FieldTerminalPath{selector: ActivityLog_FieldPathSelectorName},
			"pattern", "activityLogId",
			[]string{"projectId", "organizationId"},
			[]gotenresource.NamePattern{NamePattern_Nil, NamePattern_Project, NamePattern_Organization}),
	}
	gotenresource.GetRegistry().RegisterDescriptor(descriptor)
}

func init() {
	initActivityLogDescriptor()
}
