// Code generated by protoc-gen-goten-object
// File: edgelq/logging/proto/v1alpha2/log.proto
// DO NOT EDIT!!!

package log

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	log_descriptor "github.com/cloudwan/edgelq-sdk/logging/resources/v1alpha2/log_descriptor"
	structpb "google.golang.org/protobuf/types/known/structpb"
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
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &log_descriptor.LogDescriptor{}
	_ = &structpb.Struct{}
	_ = &timestamppb.Timestamp{}
)

func (o *Log) GotenObjectExt() {}

func (o *Log) MakeFullFieldMask() *Log_FieldMask {
	return FullLog_FieldMask()
}

func (o *Log) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLog_FieldMask()
}

func (o *Log) MakeDiffFieldMask(other *Log) *Log_FieldMask {
	if o == nil && other == nil {
		return &Log_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLog_FieldMask()
	}

	res := &Log_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &Log_FieldTerminalPath{selector: Log_FieldPathSelectorName})
	}
	if o.GetScope() != other.GetScope() {
		res.Paths = append(res.Paths, &Log_FieldTerminalPath{selector: Log_FieldPathSelectorScope})
	}
	if o.GetService() != other.GetService() {
		res.Paths = append(res.Paths, &Log_FieldTerminalPath{selector: Log_FieldPathSelectorService})
	}
	if o.GetRegion() != other.GetRegion() {
		res.Paths = append(res.Paths, &Log_FieldTerminalPath{selector: Log_FieldPathSelectorRegion})
	}
	if o.GetVersion() != other.GetVersion() {
		res.Paths = append(res.Paths, &Log_FieldTerminalPath{selector: Log_FieldPathSelectorVersion})
	}
	if o.GetLogDescriptor().String() != other.GetLogDescriptor().String() {
		res.Paths = append(res.Paths, &Log_FieldTerminalPath{selector: Log_FieldPathSelectorLogDescriptor})
	}

	if len(o.GetLabels()) == len(other.GetLabels()) {
		for i, lValue := range o.GetLabels() {
			rValue := other.GetLabels()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &Log_FieldTerminalPath{selector: Log_FieldPathSelectorLabels})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Log_FieldTerminalPath{selector: Log_FieldPathSelectorLabels})
	}
	if !proto.Equal(o.GetTime(), other.GetTime()) {
		res.Paths = append(res.Paths, &Log_FieldTerminalPath{selector: Log_FieldPathSelectorTime})
	}
	if !proto.Equal(o.GetPayload(), other.GetPayload()) {
		res.Paths = append(res.Paths, &Log_FieldTerminalPath{selector: Log_FieldPathSelectorPayload})
	}
	return res
}

func (o *Log) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Log))
}

func (o *Log) Clone() *Log {
	if o == nil {
		return nil
	}
	result := &Log{}
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
	result.Scope = o.Scope
	result.Service = o.Service
	result.Region = o.Region
	result.Version = o.Version
	if o.LogDescriptor == nil {
		result.LogDescriptor = nil
	} else if data, err := o.LogDescriptor.ProtoString(); err != nil {
		panic(err)
	} else {
		result.LogDescriptor = &log_descriptor.Reference{}
		if err := result.LogDescriptor.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Labels = map[string]string{}
	for key, sourceValue := range o.Labels {
		result.Labels[key] = sourceValue
	}
	result.Time = proto.Clone(o.Time).(*timestamppb.Timestamp)
	result.Payload = proto.Clone(o.Payload).(*structpb.Struct)
	return result
}

func (o *Log) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Log) Merge(source *Log) {
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
	o.Scope = source.GetScope()
	o.Service = source.GetService()
	o.Region = source.GetRegion()
	o.Version = source.GetVersion()
	if source.GetLogDescriptor() != nil {
		if data, err := source.GetLogDescriptor().ProtoString(); err != nil {
			panic(err)
		} else {
			o.LogDescriptor = &log_descriptor.Reference{}
			if err := o.LogDescriptor.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.LogDescriptor = nil
	}
	if source.GetLabels() != nil {
		if o.Labels == nil {
			o.Labels = make(map[string]string, len(source.GetLabels()))
		}
		for key, sourceValue := range source.GetLabels() {
			o.Labels[key] = sourceValue
		}
	}
	if source.GetTime() != nil {
		if o.Time == nil {
			o.Time = new(timestamppb.Timestamp)
		}
		proto.Merge(o.Time, source.GetTime())
	}
	if source.GetPayload() != nil {
		if o.Payload == nil {
			o.Payload = new(structpb.Struct)
		}
		proto.Merge(o.Payload, source.GetPayload())
	}
}

func (o *Log) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Log))
}
