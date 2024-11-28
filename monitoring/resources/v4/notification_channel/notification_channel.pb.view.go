// Code generated by protoc-gen-goten-resource
// Resource: NotificationChannel
// DO NOT EDIT!!!

package notification_channel

import (
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/types/view"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
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
	_ = &project.Project{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &timestamppb.Timestamp{}
	_ = &meta.Meta{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *NotificationChannel_FieldMask) *NotificationChannel_FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_BASIC:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name", "spec.enabled", "spec.type")
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
	res := &NotificationChannel_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
