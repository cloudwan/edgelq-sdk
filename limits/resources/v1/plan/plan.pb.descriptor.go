// Code generated by protoc-gen-goten-resource
// Resource: Plan
// DO NOT EDIT!!!

package plan

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1/common"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	meta_resource "github.com/cloudwan/goten-sdk/meta-service/resources/v1/resource"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(gotenobject.FieldPath)
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &iam_iam_common.PCR{}
	_ = &iam_organization.Organization{}
	_ = &meta_resource.Resource{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

var (
	descriptor *Descriptor
)

func (r *Plan) GetRawName() gotenresource.Name {
	return r.GetName()
}

func (r *Plan) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (r *Plan) EnsureMetadata() *meta.Meta {
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
	return &Plan{}
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
	return nil
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
	return &Plan_FieldMask{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return &PlanChange{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(PlanList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(PlanChangeList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(PlanNameList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(PlanReferenceList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return make(PlanParentNameList, size, reserved)
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return make(PlanParentReferenceList, size, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(PlanMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(PlanChangeMap, reserved)
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
		meta_service.GetDescriptor(),
		iam_organization.GetDescriptor(),
	}
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParsePlan_FieldPath(raw)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}

func (d *Descriptor) SupportsMetadata() bool {
	return true
}

func (d *Descriptor) SupportsDbConstraints() bool {
	return true
}

func initPlanDescriptor() {
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"Plan", "Plans", "limits.edgelq.com", "v1"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&Plan_FieldTerminalPath{selector: Plan_FieldPathSelectorName},
			"pattern", "planId",
			[]string{"serviceId", "organizationId"},
			[]gotenresource.NamePattern{NamePattern_Service, NamePattern_Organization}),
	}
	gotenresource.GetRegistry().RegisterDescriptor(descriptor)
}

func init() {
	initPlanDescriptor()
}
