// Code generated by protoc-gen-goten-resource
// Resource: Organization
// DO NOT EDIT!!!

package organization

import (
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/runtime/api/view"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	policy "github.com/cloudwan/edgelq-sdk/meta/multi_region/proto/policy"
)

// ensure the imports are used
var (
	_ = fieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &policy.Policy{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *Organization_FieldMask) *Organization_FieldMask {
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
	res := &Organization_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
