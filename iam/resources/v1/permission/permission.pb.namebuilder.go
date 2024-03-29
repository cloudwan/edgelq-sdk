// Code generated by protoc-gen-goten-resource
// Resource: Permission
// DO NOT EDIT!!!

package permission

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// make sure we're using proto imports
var (
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

const (
	NamePattern_Service = "services/{service}/permissions/{permission}"
)

type NamePattern struct {
	Pattern gotenresource.NamePattern `firestore:"pattern"`
}

type NameBuilder struct {
	nameObj Name
}

func NewNameBuilder() *NameBuilder {
	return &NameBuilder{
		nameObj: Name{
			PermissionId: gotenresource.WildcardId,
			ParentName: ParentName{
				NamePattern: NamePattern{
					// Set default pattern - just first.
					Pattern: NamePattern_Service,
				},
			},
		},
	}
}

func (b *NameBuilder) Name() *Name {
	copied := b.nameObj
	return &copied
}

func (b *NameBuilder) Reference() *Reference {
	return b.nameObj.AsReference()
}

func (b *NameBuilder) Parent() *ParentName {
	copied := b.nameObj.ParentName
	return &copied
}

func (b *NameBuilder) ParentReference() *ParentReference {
	return b.nameObj.ParentName.AsReference()
}

func (b *NameBuilder) SetId(id string) *NameBuilder {
	b.nameObj.PermissionId = id
	return b
}

func (b *NameBuilder) SetService(parent *meta_service.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case meta_service.NamePattern_Nil:
		parentName.Pattern = NamePattern_Service
	}
	parentName.ServiceId = parent.ServiceId
	return b
}

func (b *NameBuilder) SetServiceId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.ServiceId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ServiceId != "" {
		parentName.Pattern = NamePattern_Service
	}
	return b
}
