// Code generated by protoc-gen-goten-resource
// Resource: MemberAssignment
// DO NOT EDIT!!!

package member_assignment

import (
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/types/view"
)

// proto imports
import (
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1/common"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	role_binding "github.com/cloudwan/edgelq-sdk/iam/resources/v1/role_binding"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = googlefieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &iam_common.PCR{}
	_ = &organization.Organization{}
	_ = &role_binding.RoleBinding{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *MemberAssignment_FieldMask) *MemberAssignment_FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_NAME:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name")
		break
	default:
		return extraMask
	}
	if extraMask != nil {
		protoFieldMask.Paths = append(protoFieldMask.Paths, extraMask.ToProtoFieldMask().Paths...)
	}
	res := &MemberAssignment_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
