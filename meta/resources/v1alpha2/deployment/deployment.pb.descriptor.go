// Code generated by protoc-gen-goten-resource
// Resource: Deployment
// DO NOT EDIT!!!

package deployment

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	region "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/region"
	service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
)

// ensure the imports are used
var (
	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &region.Region{}
	_ = &service.Service{}
)

var (
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"Deployment", "Deployments", "meta.edgelq.com"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorName},
			"pattern", "deploymentId",
			[]string{"serviceId"},
			[]gotenresource.NamePattern{NamePattern_Service}),
	}
)

type Descriptor struct {
	nameDescriptor *gotenresource.NameDescriptor
	typeName       *gotenresource.TypeName
}

func GetDescriptor() *Descriptor {
	return descriptor
}

func (d *Descriptor) NewDeployment() *Deployment {
	return &Deployment{}
}

func (d *Descriptor) NewResource() gotenresource.Resource {
	return d.NewDeployment()
}

func (d *Descriptor) NewResourceName() gotenresource.Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewDeploymentName() *Name {
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
func (d *Descriptor) NewDeploymentCursor() *PagerCursor {
	return &PagerCursor{}
}

func (d *Descriptor) NewResourceCursor() gotenresource.Cursor {
	return d.NewDeploymentCursor()
}
func (d *Descriptor) NewDeploymentChange() *DeploymentChange {
	return &DeploymentChange{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return d.NewDeploymentChange()
}

func (d *Descriptor) NewDeploymentQueryResultSnapshot() *QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return d.NewDeploymentQueryResultSnapshot()
}
func (d *Descriptor) NewDeploymentQueryResultChange() *QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewSearchQueryResultSnapshot() gotenresource.SearchQueryResultSnapshot {
	return nil
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return d.NewDeploymentQueryResultChange()
}

func (d *Descriptor) NewDeploymentList(size, reserved int) DeploymentList {
	return make(DeploymentList, size, reserved)
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(DeploymentList, size, reserved)
}
func (d *Descriptor) NewDeploymentChangeList(size, reserved int) DeploymentChangeList {
	return make(DeploymentChangeList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(DeploymentChangeList, size, reserved)
}

func (d *Descriptor) NewDeploymentNameList(size, reserved int) DeploymentNameList {
	return make(DeploymentNameList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(DeploymentNameList, size, reserved)
}

func (d *Descriptor) NewDeploymentReferenceList(size, reserved int) DeploymentReferenceList {
	return make(DeploymentReferenceList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(DeploymentReferenceList, size, reserved)
}
func (d *Descriptor) NewDeploymentParentNameList(size, reserved int) DeploymentParentNameList {
	return make(DeploymentParentNameList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return make(DeploymentParentNameList, size, reserved)
}
func (d *Descriptor) NewDeploymentParentReferenceList(size, reserved int) DeploymentParentReferenceList {
	return make(DeploymentParentReferenceList, size, reserved)
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return make(DeploymentParentReferenceList, size, reserved)
}

func (d *Descriptor) NewDeploymentMap(reserved int) DeploymentMap {
	return make(DeploymentMap, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(DeploymentMap, reserved)
}
func (d *Descriptor) NewDeploymentChangeMap(reserved int) DeploymentChangeMap {
	return make(DeploymentChangeMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(DeploymentChangeMap, reserved)
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseDeployment_FieldPath(raw)
}

func (d *Descriptor) ParseDeploymentName(nameStr string) (*Name, error) {
	return ParseName(nameStr)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}
