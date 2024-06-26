// Code generated by protoc-gen-goten-object
// File: edgelq/devices/proto/v1/os_version.proto
// DO NOT EDIT!!!

package os_version

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	device_type "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device_type"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(fmt.Stringer)
	_ = new(sort.Interface)

	_ = new(proto.Message)
	_ = googlefieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &device_type.DeviceType{}
	_ = &meta.Meta{}
)

func (o *OsVersion) GotenObjectExt() {}

func (o *OsVersion) MakeFullFieldMask() *OsVersion_FieldMask {
	return FullOsVersion_FieldMask()
}

func (o *OsVersion) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullOsVersion_FieldMask()
}

func (o *OsVersion) MakeDiffFieldMask(other *OsVersion) *OsVersion_FieldMask {
	if o == nil && other == nil {
		return &OsVersion_FieldMask{}
	}
	if o == nil || other == nil {
		return FullOsVersion_FieldMask()
	}

	res := &OsVersion_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &OsVersion_FieldTerminalPath{selector: OsVersion_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &OsVersion_FieldTerminalPath{selector: OsVersion_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &OsVersion_FieldSubPath{selector: OsVersion_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetVersion() != other.GetVersion() {
		res.Paths = append(res.Paths, &OsVersion_FieldTerminalPath{selector: OsVersion_FieldPathSelectorVersion})
	}
	if o.GetDeviceType().String() != other.GetDeviceType().String() {
		res.Paths = append(res.Paths, &OsVersion_FieldTerminalPath{selector: OsVersion_FieldPathSelectorDeviceType})
	}
	if o.GetMinimumPreviousVersion() != other.GetMinimumPreviousVersion() {
		res.Paths = append(res.Paths, &OsVersion_FieldTerminalPath{selector: OsVersion_FieldPathSelectorMinimumPreviousVersion})
	}
	if o.GetChannel() != other.GetChannel() {
		res.Paths = append(res.Paths, &OsVersion_FieldTerminalPath{selector: OsVersion_FieldPathSelectorChannel})
	}
	return res
}

func (o *OsVersion) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*OsVersion))
}

func (o *OsVersion) Clone() *OsVersion {
	if o == nil {
		return nil
	}
	result := &OsVersion{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &Name{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Metadata = o.Metadata.Clone()
	result.Version = o.Version
	if o.DeviceType == nil {
		result.DeviceType = nil
	} else if data, err := o.DeviceType.ProtoString(); err != nil {
		panic(err)
	} else {
		result.DeviceType = &device_type.Reference{}
		if err := result.DeviceType.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.MinimumPreviousVersion = o.MinimumPreviousVersion
	result.Channel = o.Channel
	return result
}

func (o *OsVersion) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *OsVersion) Merge(source *OsVersion) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &Name{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
	o.Version = source.GetVersion()
	if source.GetDeviceType() != nil {
		if data, err := source.GetDeviceType().ProtoString(); err != nil {
			panic(err)
		} else {
			o.DeviceType = &device_type.Reference{}
			if err := o.DeviceType.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.DeviceType = nil
	}
	o.MinimumPreviousVersion = source.GetMinimumPreviousVersion()
	o.Channel = source.GetChannel()
}

func (o *OsVersion) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*OsVersion))
}
