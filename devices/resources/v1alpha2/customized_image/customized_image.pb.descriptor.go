// Code generated by protoc-gen-goten-resource
// Resource: CustomizedImage
// DO NOT EDIT!!!

package customized_image

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(gotenobject.FieldPath)
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &meta.Meta{}
)

var (
	descriptor *Descriptor
)

func (r *CustomizedImage) GetRawName() gotenresource.Name {
	return r.GetName()
}

func (r *CustomizedImage) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (r *CustomizedImage) EnsureMetadata() *meta.Meta {
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
	return &CustomizedImage{}
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
	return &CustomizedImage_FieldMask{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return &CustomizedImageChange{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(CustomizedImageList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(CustomizedImageChangeList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(CustomizedImageNameList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(CustomizedImageReferenceList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return make(CustomizedImageParentNameList, size, reserved)
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return make(CustomizedImageParentReferenceList, size, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(CustomizedImageMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(CustomizedImageChangeMap, reserved)
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
		project.GetDescriptor(),
	}
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseCustomizedImage_FieldPath(raw)
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

func initCustomizedImageDescriptor() {
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"CustomizedImage", "CustomizedImages", "devices.edgelq.com", "v1alpha2"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&CustomizedImage_FieldTerminalPath{selector: CustomizedImage_FieldPathSelectorName},
			"pattern", "customizedImageId",
			[]string{"projectId", "regionId"},
			[]gotenresource.NamePattern{NamePattern_Project_Region}),
	}
	gotenresource.GetRegistry().RegisterDescriptor(descriptor)
}

func init() {
	initCustomizedImageDescriptor()
}
