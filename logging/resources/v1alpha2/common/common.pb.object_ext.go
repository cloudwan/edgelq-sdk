// Code generated by protoc-gen-goten-object
// File: edgelq/logging/proto/v1alpha2/common.proto
// DO NOT EDIT!!!

package common

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	_ = &timestamppb.Timestamp{}
)

func (o *LabelDescriptor) GotenObjectExt() {}

func (o *LabelDescriptor) MakeFullFieldMask() *LabelDescriptor_FieldMask {
	return FullLabelDescriptor_FieldMask()
}

func (o *LabelDescriptor) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLabelDescriptor_FieldMask()
}

func (o *LabelDescriptor) MakeDiffFieldMask(other *LabelDescriptor) *LabelDescriptor_FieldMask {
	if o == nil && other == nil {
		return &LabelDescriptor_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLabelDescriptor_FieldMask()
	}

	res := &LabelDescriptor_FieldMask{}
	if o.GetKey() != other.GetKey() {
		res.Paths = append(res.Paths, &LabelDescriptor_FieldTerminalPath{selector: LabelDescriptor_FieldPathSelectorKey})
	}
	if o.GetDescription() != other.GetDescription() {
		res.Paths = append(res.Paths, &LabelDescriptor_FieldTerminalPath{selector: LabelDescriptor_FieldPathSelectorDescription})
	}
	return res
}

func (o *LabelDescriptor) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*LabelDescriptor))
}

func (o *LabelDescriptor) Clone() *LabelDescriptor {
	if o == nil {
		return nil
	}
	result := &LabelDescriptor{}
	result.Key = o.Key
	result.Description = o.Description
	return result
}

func (o *LabelDescriptor) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *LabelDescriptor) Merge(source *LabelDescriptor) {
	o.Key = source.GetKey()
	o.Description = source.GetDescription()
}

func (o *LabelDescriptor) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*LabelDescriptor))
}

func (o *LabelKeySet) GotenObjectExt() {}

func (o *LabelKeySet) MakeFullFieldMask() *LabelKeySet_FieldMask {
	return FullLabelKeySet_FieldMask()
}

func (o *LabelKeySet) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLabelKeySet_FieldMask()
}

func (o *LabelKeySet) MakeDiffFieldMask(other *LabelKeySet) *LabelKeySet_FieldMask {
	if o == nil && other == nil {
		return &LabelKeySet_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLabelKeySet_FieldMask()
	}

	res := &LabelKeySet_FieldMask{}

	if len(o.GetLabelKeys()) == len(other.GetLabelKeys()) {
		for i, lValue := range o.GetLabelKeys() {
			rValue := other.GetLabelKeys()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &LabelKeySet_FieldTerminalPath{selector: LabelKeySet_FieldPathSelectorLabelKeys})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &LabelKeySet_FieldTerminalPath{selector: LabelKeySet_FieldPathSelectorLabelKeys})
	}
	return res
}

func (o *LabelKeySet) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*LabelKeySet))
}

func (o *LabelKeySet) Clone() *LabelKeySet {
	if o == nil {
		return nil
	}
	result := &LabelKeySet{}
	result.LabelKeys = make([]string, len(o.LabelKeys))
	for i, sourceValue := range o.LabelKeys {
		result.LabelKeys[i] = sourceValue
	}
	return result
}

func (o *LabelKeySet) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *LabelKeySet) Merge(source *LabelKeySet) {
	for _, sourceValue := range source.GetLabelKeys() {
		exists := false
		for _, currentValue := range o.LabelKeys {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.LabelKeys = append(o.LabelKeys, newDstElement)
		}
	}

}

func (o *LabelKeySet) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*LabelKeySet))
}

func (o *TimeInterval) GotenObjectExt() {}

func (o *TimeInterval) MakeFullFieldMask() *TimeInterval_FieldMask {
	return FullTimeInterval_FieldMask()
}

func (o *TimeInterval) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullTimeInterval_FieldMask()
}

func (o *TimeInterval) MakeDiffFieldMask(other *TimeInterval) *TimeInterval_FieldMask {
	if o == nil && other == nil {
		return &TimeInterval_FieldMask{}
	}
	if o == nil || other == nil {
		return FullTimeInterval_FieldMask()
	}

	res := &TimeInterval_FieldMask{}
	if !proto.Equal(o.GetEndTime(), other.GetEndTime()) {
		res.Paths = append(res.Paths, &TimeInterval_FieldTerminalPath{selector: TimeInterval_FieldPathSelectorEndTime})
	}
	if !proto.Equal(o.GetStartTime(), other.GetStartTime()) {
		res.Paths = append(res.Paths, &TimeInterval_FieldTerminalPath{selector: TimeInterval_FieldPathSelectorStartTime})
	}
	return res
}

func (o *TimeInterval) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*TimeInterval))
}

func (o *TimeInterval) Clone() *TimeInterval {
	if o == nil {
		return nil
	}
	result := &TimeInterval{}
	result.EndTime = proto.Clone(o.EndTime).(*timestamppb.Timestamp)
	result.StartTime = proto.Clone(o.StartTime).(*timestamppb.Timestamp)
	return result
}

func (o *TimeInterval) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *TimeInterval) Merge(source *TimeInterval) {
	if source.GetEndTime() != nil {
		if o.EndTime == nil {
			o.EndTime = new(timestamppb.Timestamp)
		}
		proto.Merge(o.EndTime, source.GetEndTime())
	}
	if source.GetStartTime() != nil {
		if o.StartTime == nil {
			o.StartTime = new(timestamppb.Timestamp)
		}
		proto.Merge(o.StartTime, source.GetStartTime())
	}
}

func (o *TimeInterval) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*TimeInterval))
}