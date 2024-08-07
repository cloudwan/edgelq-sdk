// Code generated by protoc-gen-goten-resource
// Resource: Role
// DO NOT EDIT!!!

package role

import (
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/types/view"
)

// proto imports
import (
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1/condition"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1/permission"
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
	_ = &condition.Condition{}
	_ = &permission.Permission{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *Role_FieldMask) *Role_FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_BASIC:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name", "services")
		break
	case view.View_DETAIL:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name", "services", "scope_params", "grants")
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
	res := &Role_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
