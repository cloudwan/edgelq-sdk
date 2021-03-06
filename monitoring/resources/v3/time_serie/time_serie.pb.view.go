// Code generated by protoc-gen-goten-resource
// Resource: TimeSerie
// DO NOT EDIT!!!

package time_serie

import (
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/runtime/api/view"
)

// proto imports
import (
	monitoring_common "github.com/cloudwan/edgelq-sdk/monitoring/common/v3"
	metric_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/metric_descriptor"
)

// ensure the imports are used
var (
	_ = fieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &monitoring_common.LabelDescriptor{}
	_ = &metric_descriptor.MetricDescriptor{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *TimeSerie_FieldMask) *TimeSerie_FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}

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
