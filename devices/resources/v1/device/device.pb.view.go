// Code generated by protoc-gen-goten-resource
// Resource: Device
// DO NOT EDIT!!!

package device

import (
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/types/view"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1/project"
	iam_attestation_domain "github.com/cloudwan/edgelq-sdk/iam/resources/v1/attestation_domain"
	iam_iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1/common"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1/service_account"
	logging_bucket "github.com/cloudwan/edgelq-sdk/logging/resources/v1/bucket"
	monitoring_bucket "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/bucket"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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
	_ = &iam_attestation_domain.AttestationDomain{}
	_ = &iam_iam_common.PCR{}
	_ = &iam_service_account.ServiceAccount{}
	_ = &logging_bucket.Bucket{}
	_ = &monitoring_bucket.Bucket{}
	_ = &durationpb.Duration{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &timestamppb.Timestamp{}
	_ = &latlng.LatLng{}
	_ = &meta.Meta{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *Device_FieldMask) *Device_FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_BASIC:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name", "display_name", "status.device_info.os_version", "status.device_info.hardware_information.system.product_name", "status.device_info.hardware_information.system.serial_number", "metadata.labels", "metadata.tags")
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
	res := &Device_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
