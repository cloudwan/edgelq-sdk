// Code generated by protoc-gen-goten-resource
// Resource: Device
// DO NOT EDIT!!!

package device

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/project"
	iam_attestation_domain "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/attestation_domain"
	iam_iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/common"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/service_account"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &project.Project{}
	_ = &iam_attestation_domain.AttestationDomain{}
	_ = &iam_iam_common.PCR{}
	_ = &iam_service_account.ServiceAccount{}
	_ = &duration.Duration{}
	_ = &field_mask.FieldMask{}
	_ = &timestamp.Timestamp{}
)

const (
	NamePattern_Project_Region = "projects/{project}/regions/{region}/devices/{device}"
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
			DeviceId: gotenresource.WildcardId,
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
	b.nameObj.DeviceId = id
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
