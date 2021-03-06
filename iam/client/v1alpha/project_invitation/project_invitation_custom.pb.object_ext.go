// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha/project_invitation_custom.proto
// DO NOT EDIT!!!

package project_invitation_client

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	project_invitation "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project_invitation"
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
	_ = &project.Project{}
	_ = &project_invitation.ProjectInvitation{}
)

func (o *AcceptProjectInvitationRequest) GotenObjectExt() {}

func (o *AcceptProjectInvitationRequest) MakeFullFieldMask() *AcceptProjectInvitationRequest_FieldMask {
	return FullAcceptProjectInvitationRequest_FieldMask()
}

func (o *AcceptProjectInvitationRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAcceptProjectInvitationRequest_FieldMask()
}

func (o *AcceptProjectInvitationRequest) MakeDiffFieldMask(other *AcceptProjectInvitationRequest) *AcceptProjectInvitationRequest_FieldMask {
	if o == nil && other == nil {
		return &AcceptProjectInvitationRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAcceptProjectInvitationRequest_FieldMask()
	}

	res := &AcceptProjectInvitationRequest_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &AcceptProjectInvitationRequest_FieldTerminalPath{selector: AcceptProjectInvitationRequest_FieldPathSelectorName})
	}
	return res
}

func (o *AcceptProjectInvitationRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AcceptProjectInvitationRequest))
}

func (o *AcceptProjectInvitationRequest) Clone() *AcceptProjectInvitationRequest {
	if o == nil {
		return nil
	}
	result := &AcceptProjectInvitationRequest{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &project_invitation.Reference{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	return result
}

func (o *AcceptProjectInvitationRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AcceptProjectInvitationRequest) Merge(source *AcceptProjectInvitationRequest) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &project_invitation.Reference{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
}

func (o *AcceptProjectInvitationRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AcceptProjectInvitationRequest))
}

func (o *AcceptProjectInvitationResponse) GotenObjectExt() {}

func (o *AcceptProjectInvitationResponse) MakeFullFieldMask() *AcceptProjectInvitationResponse_FieldMask {
	return FullAcceptProjectInvitationResponse_FieldMask()
}

func (o *AcceptProjectInvitationResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullAcceptProjectInvitationResponse_FieldMask()
}

func (o *AcceptProjectInvitationResponse) MakeDiffFieldMask(other *AcceptProjectInvitationResponse) *AcceptProjectInvitationResponse_FieldMask {
	if o == nil && other == nil {
		return &AcceptProjectInvitationResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullAcceptProjectInvitationResponse_FieldMask()
	}

	res := &AcceptProjectInvitationResponse_FieldMask{}
	return res
}

func (o *AcceptProjectInvitationResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*AcceptProjectInvitationResponse))
}

func (o *AcceptProjectInvitationResponse) Clone() *AcceptProjectInvitationResponse {
	if o == nil {
		return nil
	}
	result := &AcceptProjectInvitationResponse{}
	return result
}

func (o *AcceptProjectInvitationResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *AcceptProjectInvitationResponse) Merge(source *AcceptProjectInvitationResponse) {
}

func (o *AcceptProjectInvitationResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*AcceptProjectInvitationResponse))
}

func (o *DeclineProjectInvitationRequest) GotenObjectExt() {}

func (o *DeclineProjectInvitationRequest) MakeFullFieldMask() *DeclineProjectInvitationRequest_FieldMask {
	return FullDeclineProjectInvitationRequest_FieldMask()
}

func (o *DeclineProjectInvitationRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullDeclineProjectInvitationRequest_FieldMask()
}

func (o *DeclineProjectInvitationRequest) MakeDiffFieldMask(other *DeclineProjectInvitationRequest) *DeclineProjectInvitationRequest_FieldMask {
	if o == nil && other == nil {
		return &DeclineProjectInvitationRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullDeclineProjectInvitationRequest_FieldMask()
	}

	res := &DeclineProjectInvitationRequest_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &DeclineProjectInvitationRequest_FieldTerminalPath{selector: DeclineProjectInvitationRequest_FieldPathSelectorName})
	}
	return res
}

func (o *DeclineProjectInvitationRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*DeclineProjectInvitationRequest))
}

func (o *DeclineProjectInvitationRequest) Clone() *DeclineProjectInvitationRequest {
	if o == nil {
		return nil
	}
	result := &DeclineProjectInvitationRequest{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &project_invitation.Reference{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	return result
}

func (o *DeclineProjectInvitationRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *DeclineProjectInvitationRequest) Merge(source *DeclineProjectInvitationRequest) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &project_invitation.Reference{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
}

func (o *DeclineProjectInvitationRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*DeclineProjectInvitationRequest))
}

func (o *DeclineProjectInvitationResponse) GotenObjectExt() {}

func (o *DeclineProjectInvitationResponse) MakeFullFieldMask() *DeclineProjectInvitationResponse_FieldMask {
	return FullDeclineProjectInvitationResponse_FieldMask()
}

func (o *DeclineProjectInvitationResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullDeclineProjectInvitationResponse_FieldMask()
}

func (o *DeclineProjectInvitationResponse) MakeDiffFieldMask(other *DeclineProjectInvitationResponse) *DeclineProjectInvitationResponse_FieldMask {
	if o == nil && other == nil {
		return &DeclineProjectInvitationResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullDeclineProjectInvitationResponse_FieldMask()
	}

	res := &DeclineProjectInvitationResponse_FieldMask{}
	return res
}

func (o *DeclineProjectInvitationResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*DeclineProjectInvitationResponse))
}

func (o *DeclineProjectInvitationResponse) Clone() *DeclineProjectInvitationResponse {
	if o == nil {
		return nil
	}
	result := &DeclineProjectInvitationResponse{}
	return result
}

func (o *DeclineProjectInvitationResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *DeclineProjectInvitationResponse) Merge(source *DeclineProjectInvitationResponse) {
}

func (o *DeclineProjectInvitationResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*DeclineProjectInvitationResponse))
}

func (o *ListMyProjectInvitationsRequest) GotenObjectExt() {}

func (o *ListMyProjectInvitationsRequest) MakeFullFieldMask() *ListMyProjectInvitationsRequest_FieldMask {
	return FullListMyProjectInvitationsRequest_FieldMask()
}

func (o *ListMyProjectInvitationsRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullListMyProjectInvitationsRequest_FieldMask()
}

func (o *ListMyProjectInvitationsRequest) MakeDiffFieldMask(other *ListMyProjectInvitationsRequest) *ListMyProjectInvitationsRequest_FieldMask {
	if o == nil && other == nil {
		return &ListMyProjectInvitationsRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullListMyProjectInvitationsRequest_FieldMask()
	}

	res := &ListMyProjectInvitationsRequest_FieldMask{}
	if o.GetParent().String() != other.GetParent().String() {
		res.Paths = append(res.Paths, &ListMyProjectInvitationsRequest_FieldTerminalPath{selector: ListMyProjectInvitationsRequest_FieldPathSelectorParent})
	}
	if o.GetFilter().String() != other.GetFilter().String() {
		res.Paths = append(res.Paths, &ListMyProjectInvitationsRequest_FieldTerminalPath{selector: ListMyProjectInvitationsRequest_FieldPathSelectorFilter})
	}
	return res
}

func (o *ListMyProjectInvitationsRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ListMyProjectInvitationsRequest))
}

func (o *ListMyProjectInvitationsRequest) Clone() *ListMyProjectInvitationsRequest {
	if o == nil {
		return nil
	}
	result := &ListMyProjectInvitationsRequest{}
	if o.Parent == nil {
		result.Parent = nil
	} else if data, err := o.Parent.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Parent = &project_invitation.ParentReference{}
		if err := result.Parent.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	if o.Filter == nil {
		result.Filter = nil
	} else if data, err := o.Filter.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Filter = &project_invitation.Filter{}
		if err := result.Filter.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	return result
}

func (o *ListMyProjectInvitationsRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ListMyProjectInvitationsRequest) Merge(source *ListMyProjectInvitationsRequest) {
	if source.GetParent() != nil {
		if data, err := source.GetParent().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Parent = &project_invitation.ParentReference{}
			if err := o.Parent.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Parent = nil
	}
	if source.GetFilter() != nil {
		if data, err := source.GetFilter().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Filter = &project_invitation.Filter{}
			if err := o.Filter.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Filter = nil
	}
}

func (o *ListMyProjectInvitationsRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ListMyProjectInvitationsRequest))
}

func (o *ListMyProjectInvitationsResponse) GotenObjectExt() {}

func (o *ListMyProjectInvitationsResponse) MakeFullFieldMask() *ListMyProjectInvitationsResponse_FieldMask {
	return FullListMyProjectInvitationsResponse_FieldMask()
}

func (o *ListMyProjectInvitationsResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullListMyProjectInvitationsResponse_FieldMask()
}

func (o *ListMyProjectInvitationsResponse) MakeDiffFieldMask(other *ListMyProjectInvitationsResponse) *ListMyProjectInvitationsResponse_FieldMask {
	if o == nil && other == nil {
		return &ListMyProjectInvitationsResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullListMyProjectInvitationsResponse_FieldMask()
	}

	res := &ListMyProjectInvitationsResponse_FieldMask{}

	if len(o.GetProjectInvitations()) == len(other.GetProjectInvitations()) {
		for i, lValue := range o.GetProjectInvitations() {
			rValue := other.GetProjectInvitations()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &ListMyProjectInvitationsResponse_FieldTerminalPath{selector: ListMyProjectInvitationsResponse_FieldPathSelectorProjectInvitations})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ListMyProjectInvitationsResponse_FieldTerminalPath{selector: ListMyProjectInvitationsResponse_FieldPathSelectorProjectInvitations})
	}
	return res
}

func (o *ListMyProjectInvitationsResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ListMyProjectInvitationsResponse))
}

func (o *ListMyProjectInvitationsResponse) Clone() *ListMyProjectInvitationsResponse {
	if o == nil {
		return nil
	}
	result := &ListMyProjectInvitationsResponse{}
	result.ProjectInvitations = make([]*project_invitation.ProjectInvitation, len(o.ProjectInvitations))
	for i, sourceValue := range o.ProjectInvitations {
		result.ProjectInvitations[i] = sourceValue.Clone()
	}
	return result
}

func (o *ListMyProjectInvitationsResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ListMyProjectInvitationsResponse) Merge(source *ListMyProjectInvitationsResponse) {
	for _, sourceValue := range source.GetProjectInvitations() {
		exists := false
		for _, currentValue := range o.ProjectInvitations {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *project_invitation.ProjectInvitation
			if sourceValue != nil {
				newDstElement = new(project_invitation.ProjectInvitation)
				newDstElement.Merge(sourceValue)
			}
			o.ProjectInvitations = append(o.ProjectInvitations, newDstElement)
		}
	}

}

func (o *ListMyProjectInvitationsResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ListMyProjectInvitationsResponse))
}
