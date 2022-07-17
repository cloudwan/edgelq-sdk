// Code generated by protoc-gen-goten-resource
// Resource: Deployment
// DO NOT EDIT!!!

package deployment

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	region "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/region"
	service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &region.Region{}
	_ = &service.Service{}
)

const (
	NamePattern_Region = "regions/{region}/deployments/{deployment}"
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
			DeploymentId: gotenresource.WildcardId,
			ParentName: ParentName{
				NamePattern: NamePattern{
					// Set default pattern - just first.
					Pattern: NamePattern_Region,
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
	b.nameObj.DeploymentId = id
	return b
}

func (b *NameBuilder) SetRegion(parent *region.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case region.NamePattern_Nil:
		parentName.Pattern = NamePattern_Region
	}
	parentName.RegionId = parent.RegionId
	return b
}

func (b *NameBuilder) SetRegionId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.RegionId = id

	// Set pattern if something matches for this set of IDs
	if parentName.RegionId != "" {
		parentName.Pattern = NamePattern_Region
	}
	return b
}
