// Code generated by protoc-gen-goten-resource
// Resource: Alert
// DO NOT EDIT!!!

package alert

import (
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/runtime/api/view"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	monitoring_common "github.com/cloudwan/edgelq-sdk/monitoring/common/v3"
	alerting_condition "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_condition"
)

// ensure the imports are used
var (
	_ = fieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &alerting_condition.AlertingCondition{}
	_ = &monitoring_common.LabelDescriptor{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *Alert_FieldMask) *Alert_FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_BASIC:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name", "info.time_serie.metric.type", "state.is_firing", "state.is_acknowledged", "state.is_silenced", "state.lifetime.start_time", "state.lifetime.end_time")
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
	res := &Alert_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
