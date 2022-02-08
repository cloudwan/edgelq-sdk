// Code generated by protoc-gen-goten-resource
// Resource: Permission
// DO NOT EDIT!!!

package permission

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import ()

// make sure we're using proto imports
var ()

const (
	NamePattern_Nil = "permissions/{permission}"
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
			PermissionId: gotenresource.WildcardId,
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
	b.nameObj.PermissionId = id
	return b
}