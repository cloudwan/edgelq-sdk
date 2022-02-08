// Code generated by protoc-gen-goten-resource
// Resource: AuditedResourceDescriptor
// DO NOT EDIT!!!

package audited_resource_descriptor

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	audit_common "github.com/cloudwan/edgelq-sdk/audit/common/v1alpha"
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
)

// make sure we're using proto imports
var (
	_ = &audit_common.Authentication{}
	_ = &ntt_meta.Meta{}
)

const (
	NamePattern_Nil = "auditedResourceDescriptors/{audited_resource_descriptor}"
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
			AuditedResourceDescriptorId: gotenresource.WildcardId,
			NamePattern: NamePattern{
				// Set default pattern - just first.
				Pattern: NamePattern_Nil,
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

func (b *NameBuilder) SetId(id string) *NameBuilder {
	b.nameObj.AuditedResourceDescriptorId = id
	return b
}