// Code generated by protoc-gen-goten-object
// File: edgelq/devices/proto/v1alpha/device_custom.proto
// DO NOT EDIT!!!

package device_client

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	api "github.com/cloudwan/edgelq-sdk/common/api"
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha/device"
)

// ensure the imports are used
var (
	_ = new(fmt.Stringer)
	_ = new(sort.Interface)

	_ = new(proto.Message)
	_ = fieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &api.Account{}
	_ = &device.Device{}
)

func (o *ProvisionServiceAccountToDeviceRequest) GotenObjectExt() {}

func (o *ProvisionServiceAccountToDeviceRequest) MakeFullFieldMask() *ProvisionServiceAccountToDeviceRequest_FieldMask {
	return FullProvisionServiceAccountToDeviceRequest_FieldMask()
}

func (o *ProvisionServiceAccountToDeviceRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProvisionServiceAccountToDeviceRequest_FieldMask()
}

func (o *ProvisionServiceAccountToDeviceRequest) MakeDiffFieldMask(other *ProvisionServiceAccountToDeviceRequest) *ProvisionServiceAccountToDeviceRequest_FieldMask {
	if o == nil && other == nil {
		return &ProvisionServiceAccountToDeviceRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProvisionServiceAccountToDeviceRequest_FieldMask()
	}

	res := &ProvisionServiceAccountToDeviceRequest_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &ProvisionServiceAccountToDeviceRequest_FieldTerminalPath{selector: ProvisionServiceAccountToDeviceRequest_FieldPathSelectorName})
	}
	return res
}

func (o *ProvisionServiceAccountToDeviceRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ProvisionServiceAccountToDeviceRequest))
}

func (o *ProvisionServiceAccountToDeviceRequest) Clone() *ProvisionServiceAccountToDeviceRequest {
	if o == nil {
		return nil
	}
	result := &ProvisionServiceAccountToDeviceRequest{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &device.Reference{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	return result
}

func (o *ProvisionServiceAccountToDeviceRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ProvisionServiceAccountToDeviceRequest) Merge(source *ProvisionServiceAccountToDeviceRequest) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &device.Reference{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
}

func (o *ProvisionServiceAccountToDeviceRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProvisionServiceAccountToDeviceRequest))
}

func (o *ProvisionServiceAccountToDeviceResponse) GotenObjectExt() {}

func (o *ProvisionServiceAccountToDeviceResponse) MakeFullFieldMask() *ProvisionServiceAccountToDeviceResponse_FieldMask {
	return FullProvisionServiceAccountToDeviceResponse_FieldMask()
}

func (o *ProvisionServiceAccountToDeviceResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProvisionServiceAccountToDeviceResponse_FieldMask()
}

func (o *ProvisionServiceAccountToDeviceResponse) MakeDiffFieldMask(other *ProvisionServiceAccountToDeviceResponse) *ProvisionServiceAccountToDeviceResponse_FieldMask {
	if o == nil && other == nil {
		return &ProvisionServiceAccountToDeviceResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProvisionServiceAccountToDeviceResponse_FieldMask()
	}

	res := &ProvisionServiceAccountToDeviceResponse_FieldMask{}
	{
		subMask := o.GetServiceAccount().MakeDiffFieldMask(other.GetServiceAccount())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProvisionServiceAccountToDeviceResponse_FieldTerminalPath{selector: ProvisionServiceAccountToDeviceResponse_FieldPathSelectorServiceAccount})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProvisionServiceAccountToDeviceResponse_FieldSubPath{selector: ProvisionServiceAccountToDeviceResponse_FieldPathSelectorServiceAccount, subPath: subpath})
			}
		}
	}
	return res
}

func (o *ProvisionServiceAccountToDeviceResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ProvisionServiceAccountToDeviceResponse))
}

func (o *ProvisionServiceAccountToDeviceResponse) Clone() *ProvisionServiceAccountToDeviceResponse {
	if o == nil {
		return nil
	}
	result := &ProvisionServiceAccountToDeviceResponse{}
	result.ServiceAccount = o.ServiceAccount.Clone()
	return result
}

func (o *ProvisionServiceAccountToDeviceResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ProvisionServiceAccountToDeviceResponse) Merge(source *ProvisionServiceAccountToDeviceResponse) {
	if source.GetServiceAccount() != nil {
		if o.ServiceAccount == nil {
			o.ServiceAccount = new(api.ServiceAccount)
		}
		o.ServiceAccount.Merge(source.GetServiceAccount())
	}
}

func (o *ProvisionServiceAccountToDeviceResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProvisionServiceAccountToDeviceResponse))
}

func (o *RemoveServiceAccountFromDeviceRequest) GotenObjectExt() {}

func (o *RemoveServiceAccountFromDeviceRequest) MakeFullFieldMask() *RemoveServiceAccountFromDeviceRequest_FieldMask {
	return FullRemoveServiceAccountFromDeviceRequest_FieldMask()
}

func (o *RemoveServiceAccountFromDeviceRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRemoveServiceAccountFromDeviceRequest_FieldMask()
}

func (o *RemoveServiceAccountFromDeviceRequest) MakeDiffFieldMask(other *RemoveServiceAccountFromDeviceRequest) *RemoveServiceAccountFromDeviceRequest_FieldMask {
	if o == nil && other == nil {
		return &RemoveServiceAccountFromDeviceRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRemoveServiceAccountFromDeviceRequest_FieldMask()
	}

	res := &RemoveServiceAccountFromDeviceRequest_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &RemoveServiceAccountFromDeviceRequest_FieldTerminalPath{selector: RemoveServiceAccountFromDeviceRequest_FieldPathSelectorName})
	}
	return res
}

func (o *RemoveServiceAccountFromDeviceRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RemoveServiceAccountFromDeviceRequest))
}

func (o *RemoveServiceAccountFromDeviceRequest) Clone() *RemoveServiceAccountFromDeviceRequest {
	if o == nil {
		return nil
	}
	result := &RemoveServiceAccountFromDeviceRequest{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &device.Reference{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	return result
}

func (o *RemoveServiceAccountFromDeviceRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RemoveServiceAccountFromDeviceRequest) Merge(source *RemoveServiceAccountFromDeviceRequest) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &device.Reference{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
}

func (o *RemoveServiceAccountFromDeviceRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RemoveServiceAccountFromDeviceRequest))
}

func (o *RemoveServiceAccountFromDeviceResponse) GotenObjectExt() {}

func (o *RemoveServiceAccountFromDeviceResponse) MakeFullFieldMask() *RemoveServiceAccountFromDeviceResponse_FieldMask {
	return FullRemoveServiceAccountFromDeviceResponse_FieldMask()
}

func (o *RemoveServiceAccountFromDeviceResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRemoveServiceAccountFromDeviceResponse_FieldMask()
}

func (o *RemoveServiceAccountFromDeviceResponse) MakeDiffFieldMask(other *RemoveServiceAccountFromDeviceResponse) *RemoveServiceAccountFromDeviceResponse_FieldMask {
	if o == nil && other == nil {
		return &RemoveServiceAccountFromDeviceResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRemoveServiceAccountFromDeviceResponse_FieldMask()
	}

	res := &RemoveServiceAccountFromDeviceResponse_FieldMask{}
	return res
}

func (o *RemoveServiceAccountFromDeviceResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RemoveServiceAccountFromDeviceResponse))
}

func (o *RemoveServiceAccountFromDeviceResponse) Clone() *RemoveServiceAccountFromDeviceResponse {
	if o == nil {
		return nil
	}
	result := &RemoveServiceAccountFromDeviceResponse{}
	return result
}

func (o *RemoveServiceAccountFromDeviceResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RemoveServiceAccountFromDeviceResponse) Merge(source *RemoveServiceAccountFromDeviceResponse) {
}

func (o *RemoveServiceAccountFromDeviceResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RemoveServiceAccountFromDeviceResponse))
}
