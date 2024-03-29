// Code generated by protoc-gen-goten-resource
// Resource: GroupMember
// DO NOT EDIT!!!

package group_member

import (
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/types/view"
)

// proto imports
import (
	group "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/group"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = googlefieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &group.Group{}
	_ = &meta.Meta{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *GroupMember_FieldMask) *GroupMember_FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_BASIC:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "member")
		break
	case view.View_DETAIL:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "member", "metadata")
		break
	case view.View_NAME:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name")
		break
	default:
		return extraMask
	}
	if extraMask != nil {
		protoFieldMask.Paths = append(protoFieldMask.Paths, extraMask.ToProtoFieldMask().Paths...)
	}
	res := &GroupMember_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
