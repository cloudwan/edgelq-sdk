// Code generated by protoc-gen-goten-resource
// Resource: GroupMember
// DO NOT EDIT!!!

package group_member

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	group "github.com/cloudwan/edgelq-sdk/iam/resources/v1/group"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// make sure we're using proto imports
var (
	_ = &group.Group{}
	_ = &meta.Meta{}
)

const (
	NamePattern_Group              = "groups/{group}/groupMembers/{group_member}"
	NamePattern_Project_Group      = "projects/{project}/groups/{group}/groupMembers/{group_member}"
	NamePattern_Organization_Group = "organizations/{organization}/groups/{group}/groupMembers/{group_member}"
	NamePattern_Service_Group      = "services/{service}/groups/{group}/groupMembers/{group_member}"
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
			GroupMemberId: gotenresource.WildcardId,
			ParentName: ParentName{
				NamePattern: NamePattern{
					// Set default pattern - just first.
					Pattern: NamePattern_Group,
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
	b.nameObj.GroupMemberId = id
	return b
}

func (b *NameBuilder) SetGroup(parent *group.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case group.NamePattern_Nil:
		parentName.Pattern = NamePattern_Group
	case group.NamePattern_Project:
		parentName.Pattern = NamePattern_Project_Group
	case group.NamePattern_Organization:
		parentName.Pattern = NamePattern_Organization_Group
	case group.NamePattern_Service:
		parentName.Pattern = NamePattern_Service_Group
	}
	parentName.GroupId = parent.GroupId
	parentName.ProjectId = parent.ProjectId
	parentName.OrganizationId = parent.OrganizationId
	parentName.ServiceId = parent.ServiceId
	return b
}

func (b *NameBuilder) SetGroupId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.GroupId = id

	// Set pattern if something matches for this set of IDs
	if parentName.GroupId != "" {
		parentName.Pattern = NamePattern_Group
	}
	if parentName.ProjectId != "" && parentName.GroupId != "" {
		parentName.Pattern = NamePattern_Project_Group
	}
	if parentName.OrganizationId != "" && parentName.GroupId != "" {
		parentName.Pattern = NamePattern_Organization_Group
	}
	if parentName.ServiceId != "" && parentName.GroupId != "" {
		parentName.Pattern = NamePattern_Service_Group
	}
	return b
}

func (b *NameBuilder) SetProjectId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.ProjectId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.GroupId != "" {
		parentName.Pattern = NamePattern_Project_Group
	}
	return b
}

func (b *NameBuilder) SetOrganizationId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.OrganizationId = id

	// Set pattern if something matches for this set of IDs
	if parentName.OrganizationId != "" && parentName.GroupId != "" {
		parentName.Pattern = NamePattern_Organization_Group
	}
	return b
}

func (b *NameBuilder) SetServiceId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.ServiceId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ServiceId != "" && parentName.GroupId != "" {
		parentName.Pattern = NamePattern_Service_Group
	}
	return b
}
