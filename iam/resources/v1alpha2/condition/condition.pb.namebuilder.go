// Code generated by protoc-gen-goten-resource
// Resource: Condition
// DO NOT EDIT!!!

package condition

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &organization.Organization{}
	_ = &project.Project{}
)

const (
	NamePattern_Nil          = "conditions/{condition}"
	NamePattern_Project      = "projects/{project}/conditions/{condition}"
	NamePattern_Organization = "organizations/{organization}/conditions/{condition}"
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
			ConditionId: gotenresource.WildcardId,
			ParentName: ParentName{
				NamePattern: NamePattern{
					// Set default pattern - just first.
					Pattern: NamePattern_Nil,
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
	b.nameObj.ConditionId = id
	return b
}

func (b *NameBuilder) SetProject(parent *project.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case project.NamePattern_Nil:
		parentName.Pattern = NamePattern_Project
	}
	parentName.ProjectId = parent.ProjectId
	return b
}

func (b *NameBuilder) SetOrganization(parent *organization.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case organization.NamePattern_Nil:
		parentName.Pattern = NamePattern_Organization
	}
	parentName.OrganizationId = parent.OrganizationId
	return b
}

func (b *NameBuilder) SetProjectId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.ProjectId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" {
		parentName.Pattern = NamePattern_Project
	}
	return b
}

func (b *NameBuilder) SetOrganizationId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.OrganizationId = id

	// Set pattern if something matches for this set of IDs
	if parentName.OrganizationId != "" {
		parentName.Pattern = NamePattern_Organization
	}
	return b
}
