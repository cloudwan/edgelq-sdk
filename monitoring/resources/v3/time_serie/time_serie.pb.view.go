// Code generated by protoc-gen-goten-resource
// Resource: TimeSerie
// DO NOT EDIT!!!

package time_serie

import (
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/types/view"
)

// proto imports
import (
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/common"
	metric_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/metric_descriptor"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = googlefieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &common.LabelDescriptor{}
	_ = &metric_descriptor.MetricDescriptor{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *TimeSerie_FieldMask) *TimeSerie_FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	default:
		return extraMask
	}
	if extraMask != nil {
		protoFieldMask.Paths = append(protoFieldMask.Paths, extraMask.ToProtoFieldMask().Paths...)
	}
	res := &TimeSerie_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
