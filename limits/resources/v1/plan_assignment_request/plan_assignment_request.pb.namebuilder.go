// Code generated by protoc-gen-goten-resource
// Resource: PlanAssignmentRequest
// DO NOT EDIT!!!

package plan_assignment_request

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	accepted_plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1/accepted_plan"
	common "github.com/cloudwan/edgelq-sdk/limits/resources/v1/common"
	plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1/plan"
	plan_assignment "github.com/cloudwan/edgelq-sdk/limits/resources/v1/plan_assignment"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// make sure we're using proto imports
var (
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &accepted_plan.AcceptedPlan{}
	_ = &common.RegionalPlanAssignment{}
	_ = &plan.Plan{}
	_ = &plan_assignment.PlanAssignment{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

const (
	NamePattern_Project      = "projects/{project}/planAssignmentRequests/{plan_assignment_request}"
	NamePattern_Organization = "organizations/{organization}/planAssignmentRequests/{plan_assignment_request}"
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
			PlanAssignmentRequestId: gotenresource.WildcardId,
			ParentName: ParentName{
				NamePattern: NamePattern{
					// Set default pattern - just first.
					Pattern: NamePattern_Project,
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
	b.nameObj.PlanAssignmentRequestId = id
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
