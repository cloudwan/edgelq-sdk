// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha2/project_invitation.proto
// DO NOT EDIT!!!

package project_invitation

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
	iam_invitation "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/invitation"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
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
	_ = &iam_invitation.Actor{}
	_ = &project.Project{}
)

func (o *ProjectInvitation) GotenObjectExt() {}

func (o *ProjectInvitation) MakeFullFieldMask() *ProjectInvitation_FieldMask {
	return FullProjectInvitation_FieldMask()
}

func (o *ProjectInvitation) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProjectInvitation_FieldMask()
}

func (o *ProjectInvitation) MakeDiffFieldMask(other *ProjectInvitation) *ProjectInvitation_FieldMask {
	if o == nil && other == nil {
		return &ProjectInvitation_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProjectInvitation_FieldMask()
	}

	res := &ProjectInvitation_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &ProjectInvitation_FieldTerminalPath{selector: ProjectInvitation_FieldPathSelectorName})
	}
	if o.GetProjectDisplayName() != other.GetProjectDisplayName() {
		res.Paths = append(res.Paths, &ProjectInvitation_FieldTerminalPath{selector: ProjectInvitation_FieldPathSelectorProjectDisplayName})
	}
	{
		subMask := o.GetInvitation().MakeDiffFieldMask(other.GetInvitation())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProjectInvitation_FieldTerminalPath{selector: ProjectInvitation_FieldPathSelectorInvitation})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProjectInvitation_FieldSubPath{selector: ProjectInvitation_FieldPathSelectorInvitation, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProjectInvitation_FieldTerminalPath{selector: ProjectInvitation_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProjectInvitation_FieldSubPath{selector: ProjectInvitation_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	return res
}

func (o *ProjectInvitation) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ProjectInvitation))
}

func (o *ProjectInvitation) Clone() *ProjectInvitation {
	if o == nil {
		return nil
	}
	result := &ProjectInvitation{}
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
	result.ProjectDisplayName = o.ProjectDisplayName
	result.Invitation = o.Invitation.Clone()
	result.Metadata = o.Metadata.Clone()
	return result
}

func (o *ProjectInvitation) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ProjectInvitation) Merge(source *ProjectInvitation) {
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
	o.ProjectDisplayName = source.GetProjectDisplayName()
	if source.GetInvitation() != nil {
		if o.Invitation == nil {
			o.Invitation = new(iam_invitation.Invitation)
		}
		o.Invitation.Merge(source.GetInvitation())
	}
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(ntt_meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
}

func (o *ProjectInvitation) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProjectInvitation))
}
