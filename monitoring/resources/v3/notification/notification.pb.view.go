// Code generated by protoc-gen-goten-resource
// Resource: Notification
// DO NOT EDIT!!!

package notification

import (
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/runtime/api/view"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	monitoring_common "github.com/cloudwan/edgelq-sdk/monitoring/common/v3"
	alert "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alert"
	alerting_condition "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_condition"
	alerting_policy "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_policy"
	notification_channel "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/notification_channel"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/project"
)

// ensure the imports are used
var (
	_ = fieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &alert.Alert{}
	_ = &alerting_condition.AlertingCondition{}
	_ = &alerting_policy.AlertingPolicy{}
	_ = &monitoring_common.LabelDescriptor{}
	_ = &notification_channel.NotificationChannel{}
	_ = &project.Project{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *Notification_FieldMask) *Notification_FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_BASIC:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "alerting_policy", "state.notification_attempts_completed", "state.notification_state")
		break
	case view.View_NAME:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name")
		break
	default:
		return extraMask
	}
	if extraMask != nil {
		protoFieldMask.Paths = append(protoFieldMask.Paths, extraMask.ToProtoFieldMask().Paths...)
	}
	res := &Notification_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
