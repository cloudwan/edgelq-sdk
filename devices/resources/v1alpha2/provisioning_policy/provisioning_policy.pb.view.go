// Code generated by protoc-gen-goten-resource
// Resource: ProvisioningPolicy
// DO NOT EDIT!!!

package provisioning_policy

import (
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/types/view"
)

// proto imports
import (
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/device"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/project"
	iam_condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/condition"
	iam_role "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/role"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/service_account"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = googlefieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &device.Device{}
	_ = &project.Project{}
	_ = &iam_condition.Condition{}
	_ = &iam_role.Role{}
	_ = &iam_service_account.ServiceAccount{}
	_ = &meta.Meta{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *ProvisioningPolicy_FieldMask) *ProvisioningPolicy_FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_NAME:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name")
		break
	default:
		return extraMask
	}
	if extraMask != nil {
		protoFieldMask.Paths = append(protoFieldMask.Paths, extraMask.ToProtoFieldMask().Paths...)
	}
	res := &ProvisioningPolicy_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
