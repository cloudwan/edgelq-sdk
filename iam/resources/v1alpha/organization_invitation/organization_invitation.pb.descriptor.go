// Code generated by protoc-gen-goten-resource
// Resource: OrganizationInvitation
// DO NOT EDIT!!!

package organization_invitation

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/common"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
)

// ensure the imports are used
var (
	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &iam_common.Actor{}
	_ = &organization.Organization{}
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
	return &OrganizationInvitation{}
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

func (d *Descriptor) NewResourceFieldMask() gotenobject.FieldMask {
	return &OrganizationInvitation_FieldMask{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return &OrganizationInvitationChange{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewSearchQueryResultSnapshot() gotenresource.SearchQueryResultSnapshot {
	return nil
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(OrganizationInvitationList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(OrganizationInvitationChangeList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(OrganizationInvitationNameList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(OrganizationInvitationReferenceList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return make(OrganizationInvitationParentNameList, size, reserved)
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return make(OrganizationInvitationParentReferenceList, size, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(OrganizationInvitationMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(OrganizationInvitationChangeMap, reserved)
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseOrganizationInvitation_FieldPath(raw)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}

func initOrganizationInvitationDescriptor() {
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"OrganizationInvitation", "OrganizationInvitations", "iam.edgelq.com", "v1alpha"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&OrganizationInvitation_FieldTerminalPath{selector: OrganizationInvitation_FieldPathSelectorName},
			"pattern", "organizationInvitationId",
			[]string{"organizationId"},
			[]gotenresource.NamePattern{NamePattern_Organization}),
	}
	gotenresource.GetRegistry().RegisterDescriptor(descriptor)
}

func init() {
	initOrganizationInvitationDescriptor()
}
