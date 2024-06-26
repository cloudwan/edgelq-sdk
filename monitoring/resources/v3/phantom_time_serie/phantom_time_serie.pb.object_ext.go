// Code generated by protoc-gen-goten-object
// File: edgelq/monitoring/proto/v3/phantom_time_serie.proto
// DO NOT EDIT!!!

package phantom_time_serie

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/common"
	metric_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/metric_descriptor"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/project"
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
	_ = &common.LabelDescriptor{}
	_ = &metric_descriptor.MetricDescriptor{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

func (o *PhantomTimeSerie) GotenObjectExt() {}

func (o *PhantomTimeSerie) MakeFullFieldMask() *PhantomTimeSerie_FieldMask {
	return FullPhantomTimeSerie_FieldMask()
}

func (o *PhantomTimeSerie) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPhantomTimeSerie_FieldMask()
}

func (o *PhantomTimeSerie) MakeDiffFieldMask(other *PhantomTimeSerie) *PhantomTimeSerie_FieldMask {
	if o == nil && other == nil {
		return &PhantomTimeSerie_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPhantomTimeSerie_FieldMask()
	}

	res := &PhantomTimeSerie_FieldMask{}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &PhantomTimeSerie_FieldSubPath{selector: PhantomTimeSerie_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorName})
	}
	if string(o.GetKey()) != string(other.GetKey()) {
		res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorKey})
	}
	if o.GetProject() != other.GetProject() {
		res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorProject})
	}
	{
		subMask := o.GetMetric().MakeDiffFieldMask(other.GetMetric())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorMetric})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &PhantomTimeSerie_FieldSubPath{selector: PhantomTimeSerie_FieldPathSelectorMetric, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetResource().MakeDiffFieldMask(other.GetResource())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorResource})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &PhantomTimeSerie_FieldSubPath{selector: PhantomTimeSerie_FieldPathSelectorResource, subPath: subpath})
			}
		}
	}
	if o.GetMetricKind() != other.GetMetricKind() {
		res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorMetricKind})
	}
	if o.GetValueType() != other.GetValueType() {
		res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorValueType})
	}
	{
		subMask := o.GetValue().MakeDiffFieldMask(other.GetValue())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorValue})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &PhantomTimeSerie_FieldSubPath{selector: PhantomTimeSerie_FieldPathSelectorValue, subPath: subpath})
			}
		}
	}
	return res
}

func (o *PhantomTimeSerie) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*PhantomTimeSerie))
}

func (o *PhantomTimeSerie) Clone() *PhantomTimeSerie {
	if o == nil {
		return nil
	}
	result := &PhantomTimeSerie{}
	result.Metadata = o.Metadata.Clone()
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
	result.Key = make([]byte, len(o.Key))
	for i, bt := range o.Key {
		result.Key[i] = bt
	}
	result.Project = o.Project
	result.Metric = o.Metric.Clone()
	result.Resource = o.Resource.Clone()
	result.MetricKind = o.MetricKind
	result.ValueType = o.ValueType
	result.Value = o.Value.Clone()
	return result
}

func (o *PhantomTimeSerie) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *PhantomTimeSerie) Merge(source *PhantomTimeSerie) {
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
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
	o.Key = make([]byte, len(source.GetKey()))
	for i, bt := range source.GetKey() {
		o.Key[i] = bt
	}
	o.Project = source.GetProject()
	if source.GetMetric() != nil {
		if o.Metric == nil {
			o.Metric = new(common.Metric)
		}
		o.Metric.Merge(source.GetMetric())
	}
	if source.GetResource() != nil {
		if o.Resource == nil {
			o.Resource = new(common.MonitoredResource)
		}
		o.Resource.Merge(source.GetResource())
	}
	o.MetricKind = source.GetMetricKind()
	o.ValueType = source.GetValueType()
	if source.GetValue() != nil {
		if o.Value == nil {
			o.Value = new(common.TypedValue)
		}
		o.Value.Merge(source.GetValue())
	}
}

func (o *PhantomTimeSerie) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*PhantomTimeSerie))
}
