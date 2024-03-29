// Code generated by protoc-gen-goten-resource
// Resource: Plan
// DO NOT EDIT!!!

package plan

import (
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/runtime/api/view"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	iam_iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/common"
	common "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/common"
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
	_ = &iam_iam_common.PCR{}
	_ = &common.Allowance{}
	_ = &meta_service.Service{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *Plan_FieldMask) *Plan_FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_BASIC:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name", "service")
		break
	case view.View_DETAIL:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name", "service", "resource_limits")
		break
	case view.View_NAME:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name")
		break
	default:
		return extraMask
	}
	if extraMask != nil {
		protoFieldMask.Paths = append(protoFieldMask.Paths, extraMask.ToProtoFieldMask().Paths...)
	}
	res := &Plan_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
