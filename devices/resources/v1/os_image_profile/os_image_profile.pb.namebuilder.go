// Code generated by protoc-gen-goten-resource
// Resource: OsImageProfile
// DO NOT EDIT!!!

package os_image_profile

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	device_type "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device_type"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// make sure we're using proto imports
var (
	_ = &device_type.DeviceType{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

const (
	NamePattern_Project_Region = "projects/{project}/regions/{region}/osImageProfiles/{os_image_profile}"
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
			OsImageProfileId: gotenresource.WildcardId,
			ParentName: ParentName{
				NamePattern: NamePattern{
					// Set default pattern - just first.
					Pattern: NamePattern_Project_Region,
				},
				RegionId: gotenresource.WildcardId,
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
	b.nameObj.OsImageProfileId = id
	return b
}

func (b *NameBuilder) SetProject(parent *project.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case project.NamePattern_Nil:
		parentName.Pattern = NamePattern_Project_Region
	}
	parentName.ProjectId = parent.ProjectId
	return b
}

func (b *NameBuilder) SetProjectId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.ProjectId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.RegionId != "" {
		parentName.Pattern = NamePattern_Project_Region
	}
	return b
}

func (b *NameBuilder) SetRegionId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.RegionId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.RegionId != "" {
		parentName.Pattern = NamePattern_Project_Region
	}
	return b
}
