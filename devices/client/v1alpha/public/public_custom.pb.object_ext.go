// Code generated by protoc-gen-goten-object
// File: edgelq/devices/proto/v1alpha/public_custom.proto
// DO NOT EDIT!!!

package public_client

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
	view "github.com/cloudwan/goten-sdk/runtime/api/view"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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
	_ = &field_mask.FieldMask{}
	_ = view.View(0)
)

func (o *ListPublicDevicesRequest) GotenObjectExt() {}

func (o *ListPublicDevicesRequest) MakeFullFieldMask() *ListPublicDevicesRequest_FieldMask {
	return FullListPublicDevicesRequest_FieldMask()
}

func (o *ListPublicDevicesRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullListPublicDevicesRequest_FieldMask()
}

func (o *ListPublicDevicesRequest) MakeDiffFieldMask(other *ListPublicDevicesRequest) *ListPublicDevicesRequest_FieldMask {
	if o == nil && other == nil {
		return &ListPublicDevicesRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullListPublicDevicesRequest_FieldMask()
	}

	res := &ListPublicDevicesRequest_FieldMask{}
	if o.GetParent().String() != other.GetParent().String() {
		res.Paths = append(res.Paths, &ListPublicDevicesRequest_FieldTerminalPath{selector: ListPublicDevicesRequest_FieldPathSelectorParent})
	}
	if o.GetPageSize() != other.GetPageSize() {
		res.Paths = append(res.Paths, &ListPublicDevicesRequest_FieldTerminalPath{selector: ListPublicDevicesRequest_FieldPathSelectorPageSize})
	}
	if o.GetPageToken().String() != other.GetPageToken().String() {
		res.Paths = append(res.Paths, &ListPublicDevicesRequest_FieldTerminalPath{selector: ListPublicDevicesRequest_FieldPathSelectorPageToken})
	}
	if o.GetOrderBy().String() != other.GetOrderBy().String() {
		res.Paths = append(res.Paths, &ListPublicDevicesRequest_FieldTerminalPath{selector: ListPublicDevicesRequest_FieldPathSelectorOrderBy})
	}
	if o.GetFilter().String() != other.GetFilter().String() {
		res.Paths = append(res.Paths, &ListPublicDevicesRequest_FieldTerminalPath{selector: ListPublicDevicesRequest_FieldPathSelectorFilter})
	}
	if !proto.Equal(o.GetFieldMask(), other.GetFieldMask()) {
		res.Paths = append(res.Paths, &ListPublicDevicesRequest_FieldTerminalPath{selector: ListPublicDevicesRequest_FieldPathSelectorFieldMask})
	}
	if o.GetView() != other.GetView() {
		res.Paths = append(res.Paths, &ListPublicDevicesRequest_FieldTerminalPath{selector: ListPublicDevicesRequest_FieldPathSelectorView})
	}
	return res
}

func (o *ListPublicDevicesRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ListPublicDevicesRequest))
}

func (o *ListPublicDevicesRequest) Clone() *ListPublicDevicesRequest {
	if o == nil {
		return nil
	}
	result := &ListPublicDevicesRequest{}
	if o.Parent == nil {
		result.Parent = nil
	} else if data, err := o.Parent.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Parent = &device.ParentName{}
		if err := result.Parent.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.PageSize = o.PageSize
	if o.PageToken == nil {
		result.PageToken = nil
	} else if data, err := o.PageToken.ProtoString(); err != nil {
		panic(err)
	} else {
		result.PageToken = &device.PagerCursor{}
		if err := result.PageToken.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	if o.OrderBy == nil {
		result.OrderBy = nil
	} else if data, err := o.OrderBy.ProtoString(); err != nil {
		panic(err)
	} else {
		result.OrderBy = &device.OrderBy{}
		if err := result.OrderBy.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	if o.Filter == nil {
		result.Filter = nil
	} else if data, err := o.Filter.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Filter = &device.Filter{}
		if err := result.Filter.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.FieldMask = proto.Clone(o.FieldMask).(*device.Device_FieldMask)
	result.View = o.View
	return result
}

func (o *ListPublicDevicesRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ListPublicDevicesRequest) Merge(source *ListPublicDevicesRequest) {
	if source.GetParent() != nil {
		if data, err := source.GetParent().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Parent = &device.ParentName{}
			if err := o.Parent.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Parent = nil
	}
	o.PageSize = source.GetPageSize()
	if source.GetPageToken() != nil {
		if data, err := source.GetPageToken().ProtoString(); err != nil {
			panic(err)
		} else {
			o.PageToken = &device.PagerCursor{}
			if err := o.PageToken.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.PageToken = nil
	}
	if source.GetOrderBy() != nil {
		if data, err := source.GetOrderBy().ProtoString(); err != nil {
			panic(err)
		} else {
			o.OrderBy = &device.OrderBy{}
			if err := o.OrderBy.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.OrderBy = nil
	}
	if source.GetFilter() != nil {
		if data, err := source.GetFilter().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Filter = &device.Filter{}
			if err := o.Filter.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Filter = nil
	}
	if source.GetFieldMask() != nil {
		if o.FieldMask == nil {
			o.FieldMask = new(device.Device_FieldMask)
		}
		mergedMask := fieldmaskpb.Union(source.GetFieldMask().ToProtoFieldMask(), o.FieldMask.ToProtoFieldMask())
		if err := o.FieldMask.FromProtoFieldMask(mergedMask); err != nil {
			panic(err)
		}
	}
	o.View = source.GetView()
}

func (o *ListPublicDevicesRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ListPublicDevicesRequest))
}

func (o *ListPublicDevicesResponse) GotenObjectExt() {}

func (o *ListPublicDevicesResponse) MakeFullFieldMask() *ListPublicDevicesResponse_FieldMask {
	return FullListPublicDevicesResponse_FieldMask()
}

func (o *ListPublicDevicesResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullListPublicDevicesResponse_FieldMask()
}

func (o *ListPublicDevicesResponse) MakeDiffFieldMask(other *ListPublicDevicesResponse) *ListPublicDevicesResponse_FieldMask {
	if o == nil && other == nil {
		return &ListPublicDevicesResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullListPublicDevicesResponse_FieldMask()
	}

	res := &ListPublicDevicesResponse_FieldMask{}

	if len(o.GetDevices()) == len(other.GetDevices()) {
		for i, lValue := range o.GetDevices() {
			rValue := other.GetDevices()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &ListPublicDevicesResponse_FieldTerminalPath{selector: ListPublicDevicesResponse_FieldPathSelectorDevices})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ListPublicDevicesResponse_FieldTerminalPath{selector: ListPublicDevicesResponse_FieldPathSelectorDevices})
	}
	if o.GetPrevPageToken().String() != other.GetPrevPageToken().String() {
		res.Paths = append(res.Paths, &ListPublicDevicesResponse_FieldTerminalPath{selector: ListPublicDevicesResponse_FieldPathSelectorPrevPageToken})
	}
	if o.GetNextPageToken().String() != other.GetNextPageToken().String() {
		res.Paths = append(res.Paths, &ListPublicDevicesResponse_FieldTerminalPath{selector: ListPublicDevicesResponse_FieldPathSelectorNextPageToken})
	}
	return res
}

func (o *ListPublicDevicesResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ListPublicDevicesResponse))
}

func (o *ListPublicDevicesResponse) Clone() *ListPublicDevicesResponse {
	if o == nil {
		return nil
	}
	result := &ListPublicDevicesResponse{}
	result.Devices = make([]*device.Device, len(o.Devices))
	for i, sourceValue := range o.Devices {
		result.Devices[i] = sourceValue.Clone()
	}
	if o.PrevPageToken == nil {
		result.PrevPageToken = nil
	} else if data, err := o.PrevPageToken.ProtoString(); err != nil {
		panic(err)
	} else {
		result.PrevPageToken = &device.PagerCursor{}
		if err := result.PrevPageToken.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	if o.NextPageToken == nil {
		result.NextPageToken = nil
	} else if data, err := o.NextPageToken.ProtoString(); err != nil {
		panic(err)
	} else {
		result.NextPageToken = &device.PagerCursor{}
		if err := result.NextPageToken.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	return result
}

func (o *ListPublicDevicesResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ListPublicDevicesResponse) Merge(source *ListPublicDevicesResponse) {
	for _, sourceValue := range source.GetDevices() {
		exists := false
		for _, currentValue := range o.Devices {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *device.Device
			if sourceValue != nil {
				newDstElement = new(device.Device)
				newDstElement.Merge(sourceValue)
			}
			o.Devices = append(o.Devices, newDstElement)
		}
	}

	if source.GetPrevPageToken() != nil {
		if data, err := source.GetPrevPageToken().ProtoString(); err != nil {
			panic(err)
		} else {
			o.PrevPageToken = &device.PagerCursor{}
			if err := o.PrevPageToken.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.PrevPageToken = nil
	}
	if source.GetNextPageToken() != nil {
		if data, err := source.GetNextPageToken().ProtoString(); err != nil {
			panic(err)
		} else {
			o.NextPageToken = &device.PagerCursor{}
			if err := o.NextPageToken.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.NextPageToken = nil
	}
}

func (o *ListPublicDevicesResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ListPublicDevicesResponse))
}
