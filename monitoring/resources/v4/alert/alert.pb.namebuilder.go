// Code generated by protoc-gen-goten-resource
// Resource: Alert
// DO NOT EDIT!!!

package alert

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	alerting_condition "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/alerting_condition"
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/common"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// make sure we're using proto imports
var (
	_ = &alerting_condition.AlertingCondition{}
	_ = &common.LabelDescriptor{}
	_ = &meta.Meta{}
)

const (
	NamePattern_Project_Region_AlertingPolicy_AlertingCondition = "projects/{project}/regions/{region}/alertingPolicies/{alerting_policy}/alertingConditions/{alerting_condition}/alerts/{alert}"
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
			AlertId: gotenresource.WildcardId,
			ParentName: ParentName{
				NamePattern: NamePattern{
					// Set default pattern - just first.
					Pattern: NamePattern_Project_Region_AlertingPolicy_AlertingCondition,
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
	b.nameObj.AlertId = id
	return b
}

func (b *NameBuilder) SetAlertingCondition(parent *alerting_condition.Name) *NameBuilder {
	parentName := &b.nameObj.ParentName

	switch parent.Pattern {
	case alerting_condition.NamePattern_Project_Region_AlertingPolicy:
		parentName.Pattern = NamePattern_Project_Region_AlertingPolicy_AlertingCondition
	}
	parentName.ProjectId = parent.ProjectId
	parentName.RegionId = parent.RegionId
	parentName.AlertingPolicyId = parent.AlertingPolicyId
	parentName.AlertingConditionId = parent.AlertingConditionId
	return b
}

func (b *NameBuilder) SetProjectId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.ProjectId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.RegionId != "" && parentName.AlertingPolicyId != "" && parentName.AlertingConditionId != "" {
		parentName.Pattern = NamePattern_Project_Region_AlertingPolicy_AlertingCondition
	}
	return b
}

func (b *NameBuilder) SetRegionId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.RegionId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.RegionId != "" && parentName.AlertingPolicyId != "" && parentName.AlertingConditionId != "" {
		parentName.Pattern = NamePattern_Project_Region_AlertingPolicy_AlertingCondition
	}
	return b
}

func (b *NameBuilder) SetAlertingPolicyId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.AlertingPolicyId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.RegionId != "" && parentName.AlertingPolicyId != "" && parentName.AlertingConditionId != "" {
		parentName.Pattern = NamePattern_Project_Region_AlertingPolicy_AlertingCondition
	}
	return b
}

func (b *NameBuilder) SetAlertingConditionId(id string) *NameBuilder {
	parentName := &b.nameObj.ParentName
	parentName.AlertingConditionId = id

	// Set pattern if something matches for this set of IDs
	if parentName.ProjectId != "" && parentName.RegionId != "" && parentName.AlertingPolicyId != "" && parentName.AlertingConditionId != "" {
		parentName.Pattern = NamePattern_Project_Region_AlertingPolicy_AlertingCondition
	}
	return b
}