// Code generated by protoc-gen-goten-resource
// Resource: ResourceChangeLog
// DO NOT EDIT!!!

package resource_change_log

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	audit_common "github.com/cloudwan/edgelq-sdk/audit/common/v1alpha"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// make sure we're using proto imports
var (
	_ = &audit_common.Authentication{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &timestamp.Timestamp{}
)

const (
	NamePattern_Nil          = "resourceChangeLogs/{resource_change_log}"
	NamePattern_Project      = "projects/{project}/resourceChangeLogs/{resource_change_log}"
	NamePattern_Organization = "organizations/{organization}/resourceChangeLogs/{resource_change_log}"
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
			ResourceChangeLogId: gotenresource.WildcardId,
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
	b.nameObj.ResourceChangeLogId = id
	return b
}

func (b *NameBuilder) SetProject(parent *iam_project.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case iam_project.NamePattern_Nil:
		parentName.Pattern = NamePattern_Project
	}
	parentName.ProjectId = parent.ProjectId
	return b
}

func (b *NameBuilder) SetOrganization(parent *iam_organization.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case iam_organization.NamePattern_Nil:
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
