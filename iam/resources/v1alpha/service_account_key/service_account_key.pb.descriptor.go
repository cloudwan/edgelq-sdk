// Code generated by protoc-gen-goten-resource
// Resource: ServiceAccountKey
// DO NOT EDIT!!!

package service_account_key

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/service_account"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// ensure the imports are used
var (
	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &service_account.ServiceAccount{}
	_ = &timestamp.Timestamp{}
)

var (
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"ServiceAccountKey", "ServiceAccountKeys", "iam.edgelq.com"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&ServiceAccountKey_FieldTerminalPath{selector: ServiceAccountKey_FieldPathSelectorName},
			"pattern", "serviceAccountKeyId",
			[]string{"projectId", "serviceAccountId"},
			[]gotenresource.NamePattern{NamePattern_Project_ServiceAccount}),
	}
)

type Descriptor struct {
	nameDescriptor *gotenresource.NameDescriptor
	typeName       *gotenresource.TypeName
}

func GetDescriptor() *Descriptor {
	return descriptor
}

func (d *Descriptor) NewServiceAccountKey() *ServiceAccountKey {
	return &ServiceAccountKey{}
}

func (d *Descriptor) NewResource() gotenresource.Resource {
	return d.NewServiceAccountKey()
}

func (d *Descriptor) NewResourceName() gotenresource.Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewServiceAccountKeyName() *Name {
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
func (d *Descriptor) NewServiceAccountKeyCursor() *PagerCursor {
	return &PagerCursor{}
}

func (d *Descriptor) NewResourceCursor() gotenresource.Cursor {
	return d.NewServiceAccountKeyCursor()
}
func (d *Descriptor) NewServiceAccountKeyChange() *ServiceAccountKeyChange {
	return &ServiceAccountKeyChange{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return d.NewServiceAccountKeyChange()
}

func (d *Descriptor) NewServiceAccountKeyQueryResultSnapshot() *QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return d.NewServiceAccountKeyQueryResultSnapshot()
}
func (d *Descriptor) NewServiceAccountKeyQueryResultChange() *QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewSearchQueryResultSnapshot() gotenresource.SearchQueryResultSnapshot {
	return nil
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return d.NewServiceAccountKeyQueryResultChange()
}

func (d *Descriptor) NewServiceAccountKeyList(size, reserved int) ServiceAccountKeyList {
	return make(ServiceAccountKeyList, size, reserved)
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(ServiceAccountKeyList, size, reserved)
}
func (d *Descriptor) NewServiceAccountKeyChangeList(size, reserved int) ServiceAccountKeyChangeList {
	return make(ServiceAccountKeyChangeList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(ServiceAccountKeyChangeList, size, reserved)
}

func (d *Descriptor) NewServiceAccountKeyNameList(size, reserved int) ServiceAccountKeyNameList {
	return make(ServiceAccountKeyNameList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(ServiceAccountKeyNameList, size, reserved)
}

func (d *Descriptor) NewServiceAccountKeyReferenceList(size, reserved int) ServiceAccountKeyReferenceList {
	return make(ServiceAccountKeyReferenceList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(ServiceAccountKeyReferenceList, size, reserved)
}
func (d *Descriptor) NewServiceAccountKeyParentNameList(size, reserved int) ServiceAccountKeyParentNameList {
	return make(ServiceAccountKeyParentNameList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return make(ServiceAccountKeyParentNameList, size, reserved)
}
func (d *Descriptor) NewServiceAccountKeyParentReferenceList(size, reserved int) ServiceAccountKeyParentReferenceList {
	return make(ServiceAccountKeyParentReferenceList, size, reserved)
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return make(ServiceAccountKeyParentReferenceList, size, reserved)
}

func (d *Descriptor) NewServiceAccountKeyMap(reserved int) ServiceAccountKeyMap {
	return make(ServiceAccountKeyMap, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(ServiceAccountKeyMap, reserved)
}
func (d *Descriptor) NewServiceAccountKeyChangeMap(reserved int) ServiceAccountKeyChangeMap {
	return make(ServiceAccountKeyChangeMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(ServiceAccountKeyChangeMap, reserved)
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseServiceAccountKey_FieldPath(raw)
}

func (d *Descriptor) ParseServiceAccountKeyName(nameStr string) (*Name, error) {
	return ParseName(nameStr)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}