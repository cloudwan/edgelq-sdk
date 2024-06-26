// Code generated by protoc-gen-goten-resource
// Resource: Role
// DO NOT EDIT!!!

package role

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/condition"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/permission"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// make sure we're using proto imports
var (
	_ = &condition.Condition{}
	_ = &permission.Permission{}
	_ = &meta.Meta{}
)

const (
	NamePattern_Nil = "roles/{role}"
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
			RoleId: gotenresource.WildcardId,
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
	b.nameObj.RoleId = id
	return b
}
