// Code generated by protoc-gen-goten-object
// File: edgelq/common/rpc/status.proto
// DO NOT EDIT!!!

package rpc

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	anypb "google.golang.org/protobuf/types/known/anypb"
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
	_ = &anypb.Any{}
)

func (o *Status) GotenObjectExt() {}

func (o *Status) MakeFullFieldMask() *Status_FieldMask {
	return FullStatus_FieldMask()
}

func (o *Status) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullStatus_FieldMask()
}

func (o *Status) MakeDiffFieldMask(other *Status) *Status_FieldMask {
	if o == nil && other == nil {
		return &Status_FieldMask{}
	}
	if o == nil || other == nil {
		return FullStatus_FieldMask()
	}

	res := &Status_FieldMask{}
	if o.GetCode() != other.GetCode() {
		res.Paths = append(res.Paths, &Status_FieldTerminalPath{selector: Status_FieldPathSelectorCode})
	}
	if o.GetMessage() != other.GetMessage() {
		res.Paths = append(res.Paths, &Status_FieldTerminalPath{selector: Status_FieldPathSelectorMessage})
	}

	if len(o.GetDetails()) == len(other.GetDetails()) {
		for i, lValue := range o.GetDetails() {
			rValue := other.GetDetails()[i]
			if !proto.Equal(lValue, rValue) {
				res.Paths = append(res.Paths, &Status_FieldTerminalPath{selector: Status_FieldPathSelectorDetails})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Status_FieldTerminalPath{selector: Status_FieldPathSelectorDetails})
	}
	return res
}

func (o *Status) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Status))
}

func (o *Status) Clone() *Status {
	if o == nil {
		return nil
	}
	result := &Status{}
	result.Code = o.Code
	result.Message = o.Message
	result.Details = make([]*anypb.Any, len(o.Details))
	for i, sourceValue := range o.Details {
		result.Details[i] = proto.Clone(sourceValue).(*anypb.Any)
	}
	return result
}

func (o *Status) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Status) Merge(source *Status) {
	o.Code = source.GetCode()
	o.Message = source.GetMessage()
	for _, sourceValue := range source.GetDetails() {
		exists := false
		for _, currentValue := range o.Details {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *anypb.Any
			if sourceValue != nil {
				newDstElement = new(anypb.Any)
				proto.Merge(newDstElement, sourceValue)
			}
			o.Details = append(o.Details, newDstElement)
		}
	}

}

func (o *Status) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Status))
}
