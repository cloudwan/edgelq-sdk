// Code generated by protoc-gen-goten-resource
// Resource: ProvisioningApprovalRequest
// DO NOT EDIT!!!

package provisioning_approval_request

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	provisioning_policy "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/provisioning_policy"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/service_account"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// make sure we're using proto imports
var (
	_ = &provisioning_policy.ProvisioningPolicy{}
	_ = &iam_service_account.ServiceAccount{}
	_ = &meta.Meta{}
)

const (
	NamePattern_Project_Region_ProvisioningPolicy = "projects/{project}/regions/{region}/provisioningPolicies/{provisioning_policy}/provisioningApprovalRequests/{provisioning_approval_request}"
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
			ProvisioningApprovalRequestId: gotenresource.WildcardId,
			ParentName: ParentName{
				NamePattern: NamePattern{
					// Set default pattern - just first.
					Pattern: NamePattern_Project_Region_ProvisioningPolicy,
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
	b.nameObj.ProvisioningApprovalRequestId = id
	return b
}

func (b *NameBuilder) SetProvisioningPolicy(parent *provisioning_policy.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case provisioning_policy.NamePattern_Project_Region:
		parentName.Pattern = NamePattern_Project_Region_ProvisioningPolicy
	}
	parentName.ProjectId = parent.ProjectId
	parentName.RegionId = parent.RegionId
	parentName.ProvisioningPolicyId = parent.ProvisioningPolicyId
	return b
}

func (b *NameBuilder) SetProjectId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.ProjectId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.RegionId != "" && parentName.ProvisioningPolicyId != "" {
		parentName.Pattern = NamePattern_Project_Region_ProvisioningPolicy
	}
	return b
}

func (b *NameBuilder) SetRegionId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.RegionId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.RegionId != "" && parentName.ProvisioningPolicyId != "" {
		parentName.Pattern = NamePattern_Project_Region_ProvisioningPolicy
	}
	return b
}

func (b *NameBuilder) SetProvisioningPolicyId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.ProvisioningPolicyId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.RegionId != "" && parentName.ProvisioningPolicyId != "" {
		parentName.Pattern = NamePattern_Project_Region_ProvisioningPolicy
	}
	return b
}
