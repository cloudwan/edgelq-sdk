// Code generated by protoc-gen-goten-resource
// Resource: Role
// DO NOT EDIT!!!

package role

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/condition"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/permission"
)

// ensure the imports are used
var (
	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &condition.Condition{}
	_ = &permission.Permission{}
)

var (
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"Role", "Roles", "iam.edgelq.com"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&Role_FieldTerminalPath{selector: Role_FieldPathSelectorName},
			"pattern", "roleId",
			[]string{},
			[]gotenresource.NamePattern{NamePattern_Nil}),
	}
)

type Descriptor struct {
	nameDescriptor *gotenresource.NameDescriptor
	typeName       *gotenresource.TypeName
}

func GetDescriptor() *Descriptor {
	return descriptor
}

func (d *Descriptor) NewRole() *Role {
	return &Role{}
}

func (d *Descriptor) NewResource() gotenresource.Resource {
	return d.NewRole()
}

func (d *Descriptor) NewResourceName() gotenresource.Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewRoleName() *Name {
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
func (d *Descriptor) NewRoleCursor() *PagerCursor {
	return &PagerCursor{}
}

func (d *Descriptor) NewResourceCursor() gotenresource.Cursor {
	return d.NewRoleCursor()
}
func (d *Descriptor) NewRoleChange() *RoleChange {
	return &RoleChange{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return d.NewRoleChange()
}

func (d *Descriptor) NewRoleQueryResultSnapshot() *QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return d.NewRoleQueryResultSnapshot()
}
func (d *Descriptor) NewRoleQueryResultChange() *QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewSearchQueryResultSnapshot() gotenresource.SearchQueryResultSnapshot {
	return nil
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return d.NewRoleQueryResultChange()
}

func (d *Descriptor) NewRoleList(size, reserved int) RoleList {
	return make(RoleList, size, reserved)
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(RoleList, size, reserved)
}
func (d *Descriptor) NewRoleChangeList(size, reserved int) RoleChangeList {
	return make(RoleChangeList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(RoleChangeList, size, reserved)
}

func (d *Descriptor) NewRoleNameList(size, reserved int) RoleNameList {
	return make(RoleNameList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(RoleNameList, size, reserved)
}

func (d *Descriptor) NewRoleReferenceList(size, reserved int) RoleReferenceList {
	return make(RoleReferenceList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(RoleReferenceList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return nil
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return nil
}

func (d *Descriptor) NewRoleMap(reserved int) RoleMap {
	return make(RoleMap, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(RoleMap, reserved)
}
func (d *Descriptor) NewRoleChangeMap(reserved int) RoleChangeMap {
	return make(RoleChangeMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(RoleChangeMap, reserved)
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseRole_FieldPath(raw)
}

func (d *Descriptor) ParseRoleName(nameStr string) (*Name, error) {
	return ParseName(nameStr)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}
