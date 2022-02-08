// Code generated by protoc-gen-goten-object
// File: edgelq/devices/proto/v1alpha/provisioning_approval_request_custom.proto
// DO NOT EDIT!!!

package provisioning_approval_request_client

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha/device"
	provisioning_approval_request "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha/provisioning_approval_request"
)

// ensure the imports are used
var (
	_ = fmt.Stringer(nil)
	_ = sort.Interface(nil)

	_ = proto.Message(nil)
	_ = fieldmaskpb.FieldMask{}

	_ = gotenobject.FieldPath(nil)
)

// make sure we're using proto imports
var (
	_ = &device.Device{}
	_ = &provisioning_approval_request.ProvisioningApprovalRequest{}
)

func (o *ProvisionDeviceForApprovedRequestRequest) GotenObjectExt() {}

func (o *ProvisionDeviceForApprovedRequestRequest) MakeFullFieldMask() *ProvisionDeviceForApprovedRequestRequest_FieldMask {
	return FullProvisionDeviceForApprovedRequestRequest_FieldMask()
}

func (o *ProvisionDeviceForApprovedRequestRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProvisionDeviceForApprovedRequestRequest_FieldMask()
}

func (o *ProvisionDeviceForApprovedRequestRequest) MakeDiffFieldMask(other *ProvisionDeviceForApprovedRequestRequest) *ProvisionDeviceForApprovedRequestRequest_FieldMask {
	if o == nil && other == nil {
		return &ProvisionDeviceForApprovedRequestRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProvisionDeviceForApprovedRequestRequest_FieldMask()
	}

	res := &ProvisionDeviceForApprovedRequestRequest_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &ProvisionDeviceForApprovedRequestRequest_FieldTerminalPath{selector: ProvisionDeviceForApprovedRequestRequest_FieldPathSelectorName})
	}
	{
		subMask := o.GetDeviceStatus().MakeDiffFieldMask(other.GetDeviceStatus())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProvisionDeviceForApprovedRequestRequest_FieldTerminalPath{selector: ProvisionDeviceForApprovedRequestRequest_FieldPathSelectorDeviceStatus})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProvisionDeviceForApprovedRequestRequest_FieldSubPath{selector: ProvisionDeviceForApprovedRequestRequest_FieldPathSelectorDeviceStatus, subPath: subpath})
			}
		}
	}
	return res
}

func (o *ProvisionDeviceForApprovedRequestRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ProvisionDeviceForApprovedRequestRequest))
}

func (o *ProvisionDeviceForApprovedRequestRequest) Clone() *ProvisionDeviceForApprovedRequestRequest {
	if o == nil {
		return nil
	}
	result := &ProvisionDeviceForApprovedRequestRequest{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &provisioning_approval_request.Reference{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.DeviceStatus = o.DeviceStatus.Clone()
	return result
}

func (o *ProvisionDeviceForApprovedRequestRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ProvisionDeviceForApprovedRequestRequest) Merge(source *ProvisionDeviceForApprovedRequestRequest) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &provisioning_approval_request.Reference{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
	if source.GetDeviceStatus() != nil {
		if o.DeviceStatus == nil {
			o.DeviceStatus = new(device.Device_Status)
		}
		o.DeviceStatus.Merge(source.GetDeviceStatus())
	}
}

func (o *ProvisionDeviceForApprovedRequestRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProvisionDeviceForApprovedRequestRequest))
}

func (o *ProvisionDeviceForApprovedRequestResponse) GotenObjectExt() {}

func (o *ProvisionDeviceForApprovedRequestResponse) MakeFullFieldMask() *ProvisionDeviceForApprovedRequestResponse_FieldMask {
	return FullProvisionDeviceForApprovedRequestResponse_FieldMask()
}

func (o *ProvisionDeviceForApprovedRequestResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProvisionDeviceForApprovedRequestResponse_FieldMask()
}

func (o *ProvisionDeviceForApprovedRequestResponse) MakeDiffFieldMask(other *ProvisionDeviceForApprovedRequestResponse) *ProvisionDeviceForApprovedRequestResponse_FieldMask {
	if o == nil && other == nil {
		return &ProvisionDeviceForApprovedRequestResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProvisionDeviceForApprovedRequestResponse_FieldMask()
	}

	res := &ProvisionDeviceForApprovedRequestResponse_FieldMask{}
	{
		subMask := o.GetDevice().MakeDiffFieldMask(other.GetDevice())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProvisionDeviceForApprovedRequestResponse_FieldTerminalPath{selector: ProvisionDeviceForApprovedRequestResponse_FieldPathSelectorDevice})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProvisionDeviceForApprovedRequestResponse_FieldSubPath{selector: ProvisionDeviceForApprovedRequestResponse_FieldPathSelectorDevice, subPath: subpath})
			}
		}
	}
	return res
}

func (o *ProvisionDeviceForApprovedRequestResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ProvisionDeviceForApprovedRequestResponse))
}

func (o *ProvisionDeviceForApprovedRequestResponse) Clone() *ProvisionDeviceForApprovedRequestResponse {
	if o == nil {
		return nil
	}
	result := &ProvisionDeviceForApprovedRequestResponse{}
	result.Device = o.Device.Clone()
	return result
}

func (o *ProvisionDeviceForApprovedRequestResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ProvisionDeviceForApprovedRequestResponse) Merge(source *ProvisionDeviceForApprovedRequestResponse) {
	if source.GetDevice() != nil {
		if o.Device == nil {
			o.Device = new(device.Device)
		}
		o.Device.Merge(source.GetDevice())
	}
}

func (o *ProvisionDeviceForApprovedRequestResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProvisionDeviceForApprovedRequestResponse))
}