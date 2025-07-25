// Code generated by protoc-gen-goten-resource
// Resource: LogCondition
// DO NOT EDIT!!!

package log_condition

import (
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/types/view"
)

// proto imports
import (
	rcommon "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/common"
	document "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/document"
	log_condition_template "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/log_condition_template"
	policy "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/policy"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ensure the imports are used
var (
	_ = googlefieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &document.Document{}
	_ = &log_condition_template.LogConditionTemplate{}
	_ = &policy.Policy{}
	_ = &rcommon.LogCndSpec{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &meta.Meta{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *LogCondition_FieldMask) *LogCondition_FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_BASIC:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name", "spec.query.filter", "spec.query.trigger")
		break
	case view.View_DETAIL:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name", "description", "supporting_docs", "spec")
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
	res := &LogCondition_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
