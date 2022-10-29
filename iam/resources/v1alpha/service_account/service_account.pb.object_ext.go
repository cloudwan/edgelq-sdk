// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha/service_account.proto
// DO NOT EDIT!!!

package service_account

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
)

// ensure the imports are used
var (
	_ = new(fmt.Stringer)
	_ = new(sort.Interface)

	_ = new(proto.Message)
	_ = fieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &project.Project{}
)

func (o *ServiceAccount) GotenObjectExt() {}

func (o *ServiceAccount) MakeFullFieldMask() *ServiceAccount_FieldMask {
	return FullServiceAccount_FieldMask()
}

func (o *ServiceAccount) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullServiceAccount_FieldMask()
}

func (o *ServiceAccount) MakeDiffFieldMask(other *ServiceAccount) *ServiceAccount_FieldMask {
	if o == nil && other == nil {
		return &ServiceAccount_FieldMask{}
	}
	if o == nil || other == nil {
		return FullServiceAccount_FieldMask()
	}

	res := &ServiceAccount_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorName})
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorDisplayName})
	}
	if o.GetEmail() != other.GetEmail() {
		res.Paths = append(res.Paths, &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorEmail})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ServiceAccount_FieldSubPath{selector: ServiceAccount_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	return res
}

func (o *ServiceAccount) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ServiceAccount))
}

func (o *ServiceAccount) Clone() *ServiceAccount {
	if o == nil {
		return nil
	}
	result := &ServiceAccount{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &Name{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.DisplayName = o.DisplayName
	result.Email = o.Email
	result.Metadata = o.Metadata.Clone()
	return result
}

func (o *ServiceAccount) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ServiceAccount) Merge(source *ServiceAccount) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &Name{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
	o.DisplayName = source.GetDisplayName()
	o.Email = source.GetEmail()
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(ntt_meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
}

func (o *ServiceAccount) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ServiceAccount))
}
