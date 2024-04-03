// Code generated by protoc-gen-goten-resource
// Resource: OsVersion
// DO NOT EDIT!!!

package os_version

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	device_type "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device_type"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// make sure we're using proto imports
var (
	_ = &device_type.DeviceType{}
	_ = &meta.Meta{}
)

const (
	NamePattern_Nil = "osVersions/{os_version}"
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
			OsVersionId: gotenresource.WildcardId,
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
	b.nameObj.OsVersionId = id
	return b
}