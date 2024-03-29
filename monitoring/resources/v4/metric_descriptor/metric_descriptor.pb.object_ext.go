// Code generated by protoc-gen-goten-object
// File: edgelq/monitoring/proto/v4/metric_descriptor.proto
// DO NOT EDIT!!!

package metric_descriptor

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	api "github.com/cloudwan/edgelq-sdk/common/api"
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/common"
	monitored_resource_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/monitored_resource_descriptor"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
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
	_ = api.LaunchStage(0)
	_ = &common.LabelDescriptor{}
	_ = &monitored_resource_descriptor.MonitoredResourceDescriptor{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

func (o *MetricDescriptor) GotenObjectExt() {}

func (o *MetricDescriptor) MakeFullFieldMask() *MetricDescriptor_FieldMask {
	return FullMetricDescriptor_FieldMask()
}

func (o *MetricDescriptor) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullMetricDescriptor_FieldMask()
}

func (o *MetricDescriptor) MakeDiffFieldMask(other *MetricDescriptor) *MetricDescriptor_FieldMask {
	if o == nil && other == nil {
		return &MetricDescriptor_FieldMask{}
	}
	if o == nil || other == nil {
		return FullMetricDescriptor_FieldMask()
	}

	res := &MetricDescriptor_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &MetricDescriptor_FieldSubPath{selector: MetricDescriptor_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetType() != other.GetType() {
		res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorType})
	}

	if len(o.GetResourceTypes()) == len(other.GetResourceTypes()) {
		for i, lValue := range o.GetResourceTypes() {
			rValue := other.GetResourceTypes()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorResourceTypes})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorResourceTypes})
	}

	if len(o.GetLabels()) == len(other.GetLabels()) {
		for i, lValue := range o.GetLabels() {
			rValue := other.GetLabels()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorLabels})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorLabels})
	}
	if o.GetMetricKind() != other.GetMetricKind() {
		res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorMetricKind})
	}
	if o.GetValueType() != other.GetValueType() {
		res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorValueType})
	}
	if o.GetUnit() != other.GetUnit() {
		res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorUnit})
	}
	if o.GetDescription() != other.GetDescription() {
		res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorDescription})
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorDisplayName})
	}
	{
		subMask := o.GetMetricDescriptorMetadata().MakeDiffFieldMask(other.GetMetricDescriptorMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorMetricDescriptorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &MetricDescriptor_FieldSubPath{selector: MetricDescriptor_FieldPathSelectorMetricDescriptorMetadata, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetDistributionBucketOptions().MakeDiffFieldMask(other.GetDistributionBucketOptions())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorDistributionBucketOptions})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &MetricDescriptor_FieldSubPath{selector: MetricDescriptor_FieldPathSelectorDistributionBucketOptions, subPath: subpath})
			}
		}
	}

	if len(o.GetPromotedLabelKeySets()) == len(other.GetPromotedLabelKeySets()) {
		for i, lValue := range o.GetPromotedLabelKeySets() {
			rValue := other.GetPromotedLabelKeySets()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorPromotedLabelKeySets})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorPromotedLabelKeySets})
	}
	{
		subMask := o.GetIndexSpec().MakeDiffFieldMask(other.GetIndexSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorIndexSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &MetricDescriptor_FieldSubPath{selector: MetricDescriptor_FieldPathSelectorIndexSpec, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetStorageConfig().MakeDiffFieldMask(other.GetStorageConfig())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &MetricDescriptor_FieldTerminalPath{selector: MetricDescriptor_FieldPathSelectorStorageConfig})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &MetricDescriptor_FieldSubPath{selector: MetricDescriptor_FieldPathSelectorStorageConfig, subPath: subpath})
			}
		}
	}
	return res
}

func (o *MetricDescriptor) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*MetricDescriptor))
}

func (o *MetricDescriptor) Clone() *MetricDescriptor {
	if o == nil {
		return nil
	}
	result := &MetricDescriptor{}
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
	result.Type = o.Type
	result.ResourceTypes = make([]string, len(o.ResourceTypes))
	for i, sourceValue := range o.ResourceTypes {
		result.ResourceTypes[i] = sourceValue
	}
	result.Labels = make([]*common.LabelDescriptor, len(o.Labels))
	for i, sourceValue := range o.Labels {
		result.Labels[i] = sourceValue.Clone()
	}
	result.MetricKind = o.MetricKind
	result.ValueType = o.ValueType
	result.Unit = o.Unit
	result.Description = o.Description
	result.DisplayName = o.DisplayName
	result.MetricDescriptorMetadata = o.MetricDescriptorMetadata.Clone()
	result.DistributionBucketOptions = o.DistributionBucketOptions.Clone()
	result.PromotedLabelKeySets = make([]*common.LabelKeySet, len(o.PromotedLabelKeySets))
	for i, sourceValue := range o.PromotedLabelKeySets {
		result.PromotedLabelKeySets[i] = sourceValue.Clone()
	}
	result.IndexSpec = o.IndexSpec.Clone()
	result.StorageConfig = o.StorageConfig.Clone()
	return result
}

func (o *MetricDescriptor) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *MetricDescriptor) Merge(source *MetricDescriptor) {
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
	o.Type = source.GetType()
	for _, sourceValue := range source.GetResourceTypes() {
		exists := false
		for _, currentValue := range o.ResourceTypes {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.ResourceTypes = append(o.ResourceTypes, newDstElement)
		}
	}

	for _, sourceValue := range source.GetLabels() {
		exists := false
		for _, currentValue := range o.Labels {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *common.LabelDescriptor
			if sourceValue != nil {
				newDstElement = new(common.LabelDescriptor)
				newDstElement.Merge(sourceValue)
			}
			o.Labels = append(o.Labels, newDstElement)
		}
	}

	o.MetricKind = source.GetMetricKind()
	o.ValueType = source.GetValueType()
	o.Unit = source.GetUnit()
	o.Description = source.GetDescription()
	o.DisplayName = source.GetDisplayName()
	if source.GetMetricDescriptorMetadata() != nil {
		if o.MetricDescriptorMetadata == nil {
			o.MetricDescriptorMetadata = new(MetricDescriptor_MetricDescriptorMetadata)
		}
		o.MetricDescriptorMetadata.Merge(source.GetMetricDescriptorMetadata())
	}
	if source.GetDistributionBucketOptions() != nil {
		if o.DistributionBucketOptions == nil {
			o.DistributionBucketOptions = new(common.Distribution_BucketOptions)
		}
		o.DistributionBucketOptions.Merge(source.GetDistributionBucketOptions())
	}
	for _, sourceValue := range source.GetPromotedLabelKeySets() {
		exists := false
		for _, currentValue := range o.PromotedLabelKeySets {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *common.LabelKeySet
			if sourceValue != nil {
				newDstElement = new(common.LabelKeySet)
				newDstElement.Merge(sourceValue)
			}
			o.PromotedLabelKeySets = append(o.PromotedLabelKeySets, newDstElement)
		}
	}

	if source.GetIndexSpec() != nil {
		if o.IndexSpec == nil {
			o.IndexSpec = new(MetricDescriptor_IndexSpec)
		}
		o.IndexSpec.Merge(source.GetIndexSpec())
	}
	if source.GetStorageConfig() != nil {
		if o.StorageConfig == nil {
			o.StorageConfig = new(MetricDescriptor_StorageConfig)
		}
		o.StorageConfig.Merge(source.GetStorageConfig())
	}
}

func (o *MetricDescriptor) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*MetricDescriptor))
}

func (o *MetricDescriptor_MetricDescriptorMetadata) GotenObjectExt() {}

func (o *MetricDescriptor_MetricDescriptorMetadata) MakeFullFieldMask() *MetricDescriptor_MetricDescriptorMetadata_FieldMask {
	return FullMetricDescriptor_MetricDescriptorMetadata_FieldMask()
}

func (o *MetricDescriptor_MetricDescriptorMetadata) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullMetricDescriptor_MetricDescriptorMetadata_FieldMask()
}

func (o *MetricDescriptor_MetricDescriptorMetadata) MakeDiffFieldMask(other *MetricDescriptor_MetricDescriptorMetadata) *MetricDescriptor_MetricDescriptorMetadata_FieldMask {
	if o == nil && other == nil {
		return &MetricDescriptor_MetricDescriptorMetadata_FieldMask{}
	}
	if o == nil || other == nil {
		return FullMetricDescriptor_MetricDescriptorMetadata_FieldMask()
	}

	res := &MetricDescriptor_MetricDescriptorMetadata_FieldMask{}
	if o.GetLaunchStage() != other.GetLaunchStage() {
		res.Paths = append(res.Paths, &MetricDescriptorMetricDescriptorMetadata_FieldTerminalPath{selector: MetricDescriptorMetricDescriptorMetadata_FieldPathSelectorLaunchStage})
	}
	return res
}

func (o *MetricDescriptor_MetricDescriptorMetadata) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*MetricDescriptor_MetricDescriptorMetadata))
}

func (o *MetricDescriptor_MetricDescriptorMetadata) Clone() *MetricDescriptor_MetricDescriptorMetadata {
	if o == nil {
		return nil
	}
	result := &MetricDescriptor_MetricDescriptorMetadata{}
	result.LaunchStage = o.LaunchStage
	return result
}

func (o *MetricDescriptor_MetricDescriptorMetadata) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *MetricDescriptor_MetricDescriptorMetadata) Merge(source *MetricDescriptor_MetricDescriptorMetadata) {
	o.LaunchStage = source.GetLaunchStage()
}

func (o *MetricDescriptor_MetricDescriptorMetadata) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*MetricDescriptor_MetricDescriptorMetadata))
}

func (o *MetricDescriptor_IndexSpec) GotenObjectExt() {}

func (o *MetricDescriptor_IndexSpec) MakeFullFieldMask() *MetricDescriptor_IndexSpec_FieldMask {
	return FullMetricDescriptor_IndexSpec_FieldMask()
}

func (o *MetricDescriptor_IndexSpec) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullMetricDescriptor_IndexSpec_FieldMask()
}

func (o *MetricDescriptor_IndexSpec) MakeDiffFieldMask(other *MetricDescriptor_IndexSpec) *MetricDescriptor_IndexSpec_FieldMask {
	if o == nil && other == nil {
		return &MetricDescriptor_IndexSpec_FieldMask{}
	}
	if o == nil || other == nil {
		return FullMetricDescriptor_IndexSpec_FieldMask()
	}

	res := &MetricDescriptor_IndexSpec_FieldMask{}

	if len(o.GetPerResource()) == len(other.GetPerResource()) {
		for i, lValue := range o.GetPerResource() {
			rValue := other.GetPerResource()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &MetricDescriptorIndexSpec_FieldTerminalPath{selector: MetricDescriptorIndexSpec_FieldPathSelectorPerResource})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &MetricDescriptorIndexSpec_FieldTerminalPath{selector: MetricDescriptorIndexSpec_FieldPathSelectorPerResource})
	}
	return res
}

func (o *MetricDescriptor_IndexSpec) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*MetricDescriptor_IndexSpec))
}

func (o *MetricDescriptor_IndexSpec) Clone() *MetricDescriptor_IndexSpec {
	if o == nil {
		return nil
	}
	result := &MetricDescriptor_IndexSpec{}
	result.PerResource = make([]*MetricDescriptor_IndexSpec_PerMonitoredResource, len(o.PerResource))
	for i, sourceValue := range o.PerResource {
		result.PerResource[i] = sourceValue.Clone()
	}
	return result
}

func (o *MetricDescriptor_IndexSpec) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *MetricDescriptor_IndexSpec) Merge(source *MetricDescriptor_IndexSpec) {
	for _, sourceValue := range source.GetPerResource() {
		exists := false
		for _, currentValue := range o.PerResource {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *MetricDescriptor_IndexSpec_PerMonitoredResource
			if sourceValue != nil {
				newDstElement = new(MetricDescriptor_IndexSpec_PerMonitoredResource)
				newDstElement.Merge(sourceValue)
			}
			o.PerResource = append(o.PerResource, newDstElement)
		}
	}

}

func (o *MetricDescriptor_IndexSpec) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*MetricDescriptor_IndexSpec))
}

func (o *MetricDescriptor_StorageConfig) GotenObjectExt() {}

func (o *MetricDescriptor_StorageConfig) MakeFullFieldMask() *MetricDescriptor_StorageConfig_FieldMask {
	return FullMetricDescriptor_StorageConfig_FieldMask()
}

func (o *MetricDescriptor_StorageConfig) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullMetricDescriptor_StorageConfig_FieldMask()
}

func (o *MetricDescriptor_StorageConfig) MakeDiffFieldMask(other *MetricDescriptor_StorageConfig) *MetricDescriptor_StorageConfig_FieldMask {
	if o == nil && other == nil {
		return &MetricDescriptor_StorageConfig_FieldMask{}
	}
	if o == nil || other == nil {
		return FullMetricDescriptor_StorageConfig_FieldMask()
	}

	res := &MetricDescriptor_StorageConfig_FieldMask{}
	if o.GetStoreRawPoints() != other.GetStoreRawPoints() {
		res.Paths = append(res.Paths, &MetricDescriptorStorageConfig_FieldTerminalPath{selector: MetricDescriptorStorageConfig_FieldPathSelectorStoreRawPoints})
	}
	return res
}

func (o *MetricDescriptor_StorageConfig) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*MetricDescriptor_StorageConfig))
}

func (o *MetricDescriptor_StorageConfig) Clone() *MetricDescriptor_StorageConfig {
	if o == nil {
		return nil
	}
	result := &MetricDescriptor_StorageConfig{}
	result.StoreRawPoints = o.StoreRawPoints
	return result
}

func (o *MetricDescriptor_StorageConfig) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *MetricDescriptor_StorageConfig) Merge(source *MetricDescriptor_StorageConfig) {
	o.StoreRawPoints = source.GetStoreRawPoints()
}

func (o *MetricDescriptor_StorageConfig) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*MetricDescriptor_StorageConfig))
}

func (o *MetricDescriptor_IndexSpec_Index) GotenObjectExt() {}

func (o *MetricDescriptor_IndexSpec_Index) MakeFullFieldMask() *MetricDescriptor_IndexSpec_Index_FieldMask {
	return FullMetricDescriptor_IndexSpec_Index_FieldMask()
}

func (o *MetricDescriptor_IndexSpec_Index) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullMetricDescriptor_IndexSpec_Index_FieldMask()
}

func (o *MetricDescriptor_IndexSpec_Index) MakeDiffFieldMask(other *MetricDescriptor_IndexSpec_Index) *MetricDescriptor_IndexSpec_Index_FieldMask {
	if o == nil && other == nil {
		return &MetricDescriptor_IndexSpec_Index_FieldMask{}
	}
	if o == nil || other == nil {
		return FullMetricDescriptor_IndexSpec_Index_FieldMask()
	}

	res := &MetricDescriptor_IndexSpec_Index_FieldMask{}

	if len(o.GetPromotedLabels()) == len(other.GetPromotedLabels()) {
		for i, lValue := range o.GetPromotedLabels() {
			rValue := other.GetPromotedLabels()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &MetricDescriptorIndexSpecIndex_FieldTerminalPath{selector: MetricDescriptorIndexSpecIndex_FieldPathSelectorPromotedLabels})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &MetricDescriptorIndexSpecIndex_FieldTerminalPath{selector: MetricDescriptorIndexSpecIndex_FieldPathSelectorPromotedLabels})
	}
	if o.GetWriteOnly() != other.GetWriteOnly() {
		res.Paths = append(res.Paths, &MetricDescriptorIndexSpecIndex_FieldTerminalPath{selector: MetricDescriptorIndexSpecIndex_FieldPathSelectorWriteOnly})
	}
	return res
}

func (o *MetricDescriptor_IndexSpec_Index) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*MetricDescriptor_IndexSpec_Index))
}

func (o *MetricDescriptor_IndexSpec_Index) Clone() *MetricDescriptor_IndexSpec_Index {
	if o == nil {
		return nil
	}
	result := &MetricDescriptor_IndexSpec_Index{}
	result.PromotedLabels = make([]string, len(o.PromotedLabels))
	for i, sourceValue := range o.PromotedLabels {
		result.PromotedLabels[i] = sourceValue
	}
	result.WriteOnly = o.WriteOnly
	return result
}

func (o *MetricDescriptor_IndexSpec_Index) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *MetricDescriptor_IndexSpec_Index) Merge(source *MetricDescriptor_IndexSpec_Index) {
	for _, sourceValue := range source.GetPromotedLabels() {
		exists := false
		for _, currentValue := range o.PromotedLabels {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.PromotedLabels = append(o.PromotedLabels, newDstElement)
		}
	}

	o.WriteOnly = source.GetWriteOnly()
}

func (o *MetricDescriptor_IndexSpec_Index) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*MetricDescriptor_IndexSpec_Index))
}

func (o *MetricDescriptor_IndexSpec_PerMonitoredResource) GotenObjectExt() {}

func (o *MetricDescriptor_IndexSpec_PerMonitoredResource) MakeFullFieldMask() *MetricDescriptor_IndexSpec_PerMonitoredResource_FieldMask {
	return FullMetricDescriptor_IndexSpec_PerMonitoredResource_FieldMask()
}

func (o *MetricDescriptor_IndexSpec_PerMonitoredResource) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullMetricDescriptor_IndexSpec_PerMonitoredResource_FieldMask()
}

func (o *MetricDescriptor_IndexSpec_PerMonitoredResource) MakeDiffFieldMask(other *MetricDescriptor_IndexSpec_PerMonitoredResource) *MetricDescriptor_IndexSpec_PerMonitoredResource_FieldMask {
	if o == nil && other == nil {
		return &MetricDescriptor_IndexSpec_PerMonitoredResource_FieldMask{}
	}
	if o == nil || other == nil {
		return FullMetricDescriptor_IndexSpec_PerMonitoredResource_FieldMask()
	}

	res := &MetricDescriptor_IndexSpec_PerMonitoredResource_FieldMask{}
	if o.GetResource().String() != other.GetResource().String() {
		res.Paths = append(res.Paths, &MetricDescriptorIndexSpecPerMonitoredResource_FieldTerminalPath{selector: MetricDescriptorIndexSpecPerMonitoredResource_FieldPathSelectorResource})
	}

	if len(o.GetIndices()) == len(other.GetIndices()) {
		for i, lValue := range o.GetIndices() {
			rValue := other.GetIndices()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &MetricDescriptorIndexSpecPerMonitoredResource_FieldTerminalPath{selector: MetricDescriptorIndexSpecPerMonitoredResource_FieldPathSelectorIndices})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &MetricDescriptorIndexSpecPerMonitoredResource_FieldTerminalPath{selector: MetricDescriptorIndexSpecPerMonitoredResource_FieldPathSelectorIndices})
	}
	return res
}

func (o *MetricDescriptor_IndexSpec_PerMonitoredResource) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*MetricDescriptor_IndexSpec_PerMonitoredResource))
}

func (o *MetricDescriptor_IndexSpec_PerMonitoredResource) Clone() *MetricDescriptor_IndexSpec_PerMonitoredResource {
	if o == nil {
		return nil
	}
	result := &MetricDescriptor_IndexSpec_PerMonitoredResource{}
	if o.Resource == nil {
		result.Resource = nil
	} else if data, err := o.Resource.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Resource = &monitored_resource_descriptor.Reference{}
		if err := result.Resource.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Indices = make([]*MetricDescriptor_IndexSpec_Index, len(o.Indices))
	for i, sourceValue := range o.Indices {
		result.Indices[i] = sourceValue.Clone()
	}
	return result
}

func (o *MetricDescriptor_IndexSpec_PerMonitoredResource) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *MetricDescriptor_IndexSpec_PerMonitoredResource) Merge(source *MetricDescriptor_IndexSpec_PerMonitoredResource) {
	if source.GetResource() != nil {
		if data, err := source.GetResource().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Resource = &monitored_resource_descriptor.Reference{}
			if err := o.Resource.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Resource = nil
	}
	for _, sourceValue := range source.GetIndices() {
		exists := false
		for _, currentValue := range o.Indices {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *MetricDescriptor_IndexSpec_Index
			if sourceValue != nil {
				newDstElement = new(MetricDescriptor_IndexSpec_Index)
				newDstElement.Merge(sourceValue)
			}
			o.Indices = append(o.Indices, newDstElement)
		}
	}

}

func (o *MetricDescriptor_IndexSpec_PerMonitoredResource) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*MetricDescriptor_IndexSpec_PerMonitoredResource))
}
