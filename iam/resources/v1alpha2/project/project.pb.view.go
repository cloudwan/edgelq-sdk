// Code generated by protoc-gen-goten-resource
// Resource: Project
// DO NOT EDIT!!!

package project

import (
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/runtime/api/view"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	multi_region_policy "github.com/cloudwan/edgelq-sdk/common/types/multi_region_policy"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
)

// ensure the imports are used
var (
	_ = fieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
	_ = &organization.Organization{}
	_ = &meta_service.Service{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *Project_FieldMask) *Project_FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_NAME:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "title")
		break
	case view.View_BASIC:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "title", "parent_organization")
		break
	default:
		return extraMask
	}
	if extraMask != nil {
		protoFieldMask.Paths = append(protoFieldMask.Paths, extraMask.ToProtoFieldMask().Paths...)
	}
	res := &Project_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
