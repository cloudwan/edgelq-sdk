// Code generated by protoc-gen-goten-resource
// Resource: MetricDescriptor
// DO NOT EDIT!!!

package metric_descriptor

import (
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/types/view"
)

// proto imports
import (
	api "github.com/cloudwan/edgelq-sdk/common/api"
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/common"
	monitored_resource_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/monitored_resource_descriptor"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = googlefieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = api.LaunchStage(0)
	_ = &common.LabelDescriptor{}
	_ = &monitored_resource_descriptor.MonitoredResourceDescriptor{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *MetricDescriptor_FieldMask) *MetricDescriptor_FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_BASIC:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name", "type", "metric_kind", "value_type", "unit", "metric_descriptor_metadata.launch_stage")
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
	res := &MetricDescriptor_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}