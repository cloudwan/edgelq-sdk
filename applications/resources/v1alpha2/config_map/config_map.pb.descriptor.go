// Code generated by protoc-gen-goten-resource
// Resource: ConfigMap
// DO NOT EDIT!!!

package config_map

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha2/project"
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
)

// ensure the imports are used
var (
	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &ntt_meta.Meta{}
)

var (
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"ConfigMap", "ConfigMaps", "applications.edgelq.com"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorName},
			"pattern", "configMapId",
			[]string{"projectId", "regionId"},
			[]gotenresource.NamePattern{NamePattern_Project_Region}),
	}
)

type Descriptor struct {
	nameDescriptor *gotenresource.NameDescriptor
	typeName       *gotenresource.TypeName
}

func GetDescriptor() *Descriptor {
	return descriptor
}

func (d *Descriptor) NewConfigMap() *ConfigMap {
	return &ConfigMap{}
}

func (d *Descriptor) NewResource() gotenresource.Resource {
	return d.NewConfigMap()
}

func (d *Descriptor) NewResourceName() gotenresource.Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewConfigMapName() *Name {
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
func (d *Descriptor) NewConfigMapCursor() *PagerCursor {
	return &PagerCursor{}
}

func (d *Descriptor) NewResourceCursor() gotenresource.Cursor {
	return d.NewConfigMapCursor()
}
func (d *Descriptor) NewConfigMapChange() *ConfigMapChange {
	return &ConfigMapChange{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return d.NewConfigMapChange()
}

func (d *Descriptor) NewConfigMapQueryResultSnapshot() *QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return d.NewConfigMapQueryResultSnapshot()
}
func (d *Descriptor) NewConfigMapQueryResultChange() *QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewSearchQueryResultSnapshot() gotenresource.SearchQueryResultSnapshot {
	return nil
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return d.NewConfigMapQueryResultChange()
}

func (d *Descriptor) NewConfigMapList(size, reserved int) ConfigMapList {
	return make(ConfigMapList, size, reserved)
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(ConfigMapList, size, reserved)
}
func (d *Descriptor) NewConfigMapChangeList(size, reserved int) ConfigMapChangeList {
	return make(ConfigMapChangeList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(ConfigMapChangeList, size, reserved)
}

func (d *Descriptor) NewConfigMapNameList(size, reserved int) ConfigMapNameList {
	return make(ConfigMapNameList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(ConfigMapNameList, size, reserved)
}

func (d *Descriptor) NewConfigMapReferenceList(size, reserved int) ConfigMapReferenceList {
	return make(ConfigMapReferenceList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(ConfigMapReferenceList, size, reserved)
}
func (d *Descriptor) NewConfigMapParentNameList(size, reserved int) ConfigMapParentNameList {
	return make(ConfigMapParentNameList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return make(ConfigMapParentNameList, size, reserved)
}
func (d *Descriptor) NewConfigMapParentReferenceList(size, reserved int) ConfigMapParentReferenceList {
	return make(ConfigMapParentReferenceList, size, reserved)
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return make(ConfigMapParentReferenceList, size, reserved)
}

func (d *Descriptor) NewConfigMapMap(reserved int) ConfigMapMap {
	return make(ConfigMapMap, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(ConfigMapMap, reserved)
}
func (d *Descriptor) NewConfigMapChangeMap(reserved int) ConfigMapChangeMap {
	return make(ConfigMapChangeMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(ConfigMapChangeMap, reserved)
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseConfigMap_FieldPath(raw)
}

func (d *Descriptor) ParseConfigMapName(nameStr string) (*Name, error) {
	return ParseName(nameStr)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}
