// Code generated by protoc-gen-goten-resource
// Resource: ResourceChangeLog
// DO NOT EDIT!!!

package resource_change_log

import (
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/types/view"
)

// proto imports
import (
	common "github.com/cloudwan/edgelq-sdk/audit/resources/v1/common"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	anypb "google.golang.org/protobuf/types/known/anypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// ensure the imports are used
var (
	_ = googlefieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &common.Authentication{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &anypb.Any{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &timestamppb.Timestamp{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *ResourceChangeLog_FieldMask) *ResourceChangeLog_FieldMask {
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
	res := &ResourceChangeLog_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
