// Code generated by protoc-gen-goten-resource
// Resource: Project
// DO NOT EDIT!!!

package project

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import ()

// ensure the imports are used
var (
	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var ()

var (
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"Project", "Projects", "proxies.edgelq.com"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&Project_FieldTerminalPath{selector: Project_FieldPathSelectorName},
			"pattern", "projectId",
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

func (d *Descriptor) NewProject() *Project {
	return &Project{}
}

func (d *Descriptor) NewResource() gotenresource.Resource {
	return d.NewProject()
}

func (d *Descriptor) NewResourceName() gotenresource.Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewProjectName() *Name {
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
func (d *Descriptor) NewProjectCursor() *PagerCursor {
	return &PagerCursor{}
}

func (d *Descriptor) NewResourceCursor() gotenresource.Cursor {
	return d.NewProjectCursor()
}
func (d *Descriptor) NewProjectChange() *ProjectChange {
	return &ProjectChange{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return d.NewProjectChange()
}

func (d *Descriptor) NewProjectQueryResultSnapshot() *QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return d.NewProjectQueryResultSnapshot()
}
func (d *Descriptor) NewProjectQueryResultChange() *QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewSearchQueryResultSnapshot() gotenresource.SearchQueryResultSnapshot {
	return nil
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return d.NewProjectQueryResultChange()
}

func (d *Descriptor) NewProjectList(size, reserved int) ProjectList {
	return make(ProjectList, size, reserved)
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(ProjectList, size, reserved)
}
func (d *Descriptor) NewProjectChangeList(size, reserved int) ProjectChangeList {
	return make(ProjectChangeList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(ProjectChangeList, size, reserved)
}

func (d *Descriptor) NewProjectNameList(size, reserved int) ProjectNameList {
	return make(ProjectNameList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(ProjectNameList, size, reserved)
}

func (d *Descriptor) NewProjectReferenceList(size, reserved int) ProjectReferenceList {
	return make(ProjectReferenceList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(ProjectReferenceList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return nil
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return nil
}

func (d *Descriptor) NewProjectMap(reserved int) ProjectMap {
	return make(ProjectMap, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(ProjectMap, reserved)
}
func (d *Descriptor) NewProjectChangeMap(reserved int) ProjectChangeMap {
	return make(ProjectChangeMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(ProjectChangeMap, reserved)
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseProject_FieldPath(raw)
}

func (d *Descriptor) ParseProjectName(nameStr string) (*Name, error) {
	return ParseName(nameStr)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}