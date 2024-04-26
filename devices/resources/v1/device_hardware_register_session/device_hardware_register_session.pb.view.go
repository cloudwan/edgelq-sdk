// Code generated by protoc-gen-goten-resource
// Resource: DeviceHardwareRegisterSession
// DO NOT EDIT!!!

package device_hardware_register_session

import (
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/types/view"
)

// proto imports
import (
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device"
	device_hardware "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device_hardware"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1/project"
	provisioning_policy "github.com/cloudwan/edgelq-sdk/devices/resources/v1/provisioning_policy"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// ensure the imports are used
var (
	_ = googlefieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &device.Device{}
	_ = &device_hardware.DeviceHardware{}
	_ = &project.Project{}
	_ = &provisioning_policy.ProvisioningPolicy{}
	_ = &timestamppb.Timestamp{}
	_ = &meta.Meta{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *DeviceHardwareRegisterSession_FieldMask) *DeviceHardwareRegisterSession_FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_BASIC:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name", "start_time", "expiration_time", "user_email", "provisioning_policy_name", "device_name", "single_use", "token", "status.device_hardwares")
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
	res := &DeviceHardwareRegisterSession_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
