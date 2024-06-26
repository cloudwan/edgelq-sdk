// Code generated by protoc-gen-goten-resource
// Resource: Bucket
// DO NOT EDIT!!!

package bucket

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	log_descriptor "github.com/cloudwan/edgelq-sdk/logging/resources/v1/log_descriptor"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// make sure we're using proto imports
var (
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &log_descriptor.LogDescriptor{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

const (
	NamePattern_Project_Region      = "projects/{project}/regions/{region}/buckets/{bucket}"
	NamePattern_Organization_Region = "organizations/{organization}/regions/{region}/buckets/{bucket}"
	NamePattern_Service_Region      = "services/{service}/regions/{region}/buckets/{bucket}"
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
			BucketId: gotenresource.WildcardId,
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
	b.nameObj.BucketId = id
	return b
}

func (b *NameBuilder) SetProject(parent *iam_project.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case iam_project.NamePattern_Nil:
		parentName.Pattern = NamePattern_Project_Region
	}
	parentName.ProjectId = parent.ProjectId
	return b
}

func (b *NameBuilder) SetOrganization(parent *iam_organization.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case iam_organization.NamePattern_Nil:
		parentName.Pattern = NamePattern_Organization_Region
	}
	parentName.OrganizationId = parent.OrganizationId
	return b
}

func (b *NameBuilder) SetService(parent *meta_service.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case meta_service.NamePattern_Nil:
		parentName.Pattern = NamePattern_Service_Region
	}
	parentName.ServiceId = parent.ServiceId
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
	if parentName.OrganizationId != "" && parentName.RegionId != "" {
		parentName.Pattern = NamePattern_Organization_Region
	}
	if parentName.ServiceId != "" && parentName.RegionId != "" {
		parentName.Pattern = NamePattern_Service_Region
	}
	return b
}

func (b *NameBuilder) SetOrganizationId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.OrganizationId = id

	// Set pattern if something matches for this set of IDs
	if parentName.OrganizationId != "" && parentName.RegionId != "" {
		parentName.Pattern = NamePattern_Organization_Region
	}
	return b
}

func (b *NameBuilder) SetServiceId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.ServiceId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ServiceId != "" && parentName.RegionId != "" {
		parentName.Pattern = NamePattern_Service_Region
	}
	return b
}
