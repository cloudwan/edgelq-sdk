// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha/project_invitation.proto
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
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/common"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
)

// ensure the imports are used
var (
	_ = fmt.Stringer(nil)
	_ = sort.Interface(nil)

	_ = proto.Message(nil)
	_ = fieldmaskpb.FieldMask{}

	_ = gotenobject.FieldPath(nil)
)

// make sure we're using proto imports
var (
	_ = &iam_common.Actor{}
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
	result.Invitation = o.Invitation.Clone()
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
	if source.GetInvitation() != nil {
		if o.Invitation == nil {
			o.Invitation = new(iam_common.Invitation)
		}
		o.Invitation.Merge(source.GetInvitation())
	}
}

func (o *ProjectInvitation) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProjectInvitation))
}
