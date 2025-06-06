// Code generated by protoc-gen-goten-resource
// Resource: User
// DO NOT EDIT!!!

package user

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	meta "github.com/cloudwan/goten-sdk/types/meta"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// ensure the imports are used
var (
	_ = new(gotenobject.FieldPath)
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &timestamppb.Timestamp{}
	_ = &meta.Meta{}
)

var (
	descriptor *Descriptor
)

func (r *User) GetRawName() gotenresource.Name {
	return r.GetName()
}

func (r *User) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (r *User) EnsureMetadata() *meta.Meta {
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
	return &User{}
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
	return &User_FieldMask{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return &UserChange{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(UserList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(UserChangeList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(UserNameList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(UserReferenceList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return nil
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return nil
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(UserMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(UserChangeMap, reserved)
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) CanBeParentless() bool {
	return true
}

func (d *Descriptor) GetParentResDescriptors() []gotenresource.Descriptor {
	return nil
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseUser_FieldPath(raw)
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

func initUserDescriptor() {
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"User", "Users", "iam.edgelq.com", "v1"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&User_FieldTerminalPath{selector: User_FieldPathSelectorName},
			"pattern", "userId",
			[]string{},
			[]gotenresource.NamePattern{NamePattern_Nil}),
	}
	gotenresource.GetRegistry().RegisterDescriptor(descriptor)
}

func init() {
	initUserDescriptor()
}
