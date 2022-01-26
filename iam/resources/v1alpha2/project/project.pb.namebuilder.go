// Code generated by protoc-gen-goten-resource
// Resource: Project
// DO NOT EDIT!!!

package project

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	policy "github.com/cloudwan/edgelq-sdk/meta/multi_region/proto/policy"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &organization.Organization{}
	_ = &policy.Policy{}
)

const (
	NamePattern_Nil = "projects/{project}"
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
			ProjectId: gotenresource.WildcardId,
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
	b.nameObj.ProjectId = id
	return b
}
