// Code generated by protoc-gen-goten-resource
// Resource: Resource
// DO NOT EDIT!!!

package resource

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &service.Service{}
)

const (
	NamePattern_Service = "services/{service}/resources/{resource}"
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
			ResourceId: gotenresource.WildcardId,
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
	b.nameObj.ResourceId = id
	return b
}

func (b *NameBuilder) SetService(parent *service.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case service.NamePattern_Nil:
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
