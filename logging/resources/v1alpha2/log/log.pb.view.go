// Code generated by protoc-gen-goten-resource
// Resource: Log
// DO NOT EDIT!!!

package log

import (
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/runtime/api/view"
)

// proto imports
import (
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	log_descriptor "github.com/cloudwan/edgelq-sdk/logging/resources/v1alpha2/log_descriptor"
	structpb "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// ensure the imports are used
var (
	_ = fieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &log_descriptor.LogDescriptor{}
	_ = &structpb.Struct{}
	_ = &timestamp.Timestamp{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *Log_FieldMask) *Log_FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}

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
	res := &Log_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
