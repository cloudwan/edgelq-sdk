// Code generated by protoc-gen-goten-object
// File: edgelq/common/types/meta.proto
// DO NOT EDIT!!!

package ntt_meta

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	_ = &timestamp.Timestamp{}
)

func (o *Meta) GotenObjectExt() {}

func (o *Meta) MakeFullFieldMask() *Meta_FieldMask {
	return FullMeta_FieldMask()
}

func (o *Meta) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullMeta_FieldMask()
}

func (o *Meta) MakeDiffFieldMask(other *Meta) *Meta_FieldMask {
	if o == nil && other == nil {
		return &Meta_FieldMask{}
	}
	if o == nil || other == nil {
		return FullMeta_FieldMask()
	}

	res := &Meta_FieldMask{}
	if !proto.Equal(o.GetCreateTime(), other.GetCreateTime()) {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorCreateTime})
	}
	if !proto.Equal(o.GetUpdateTime(), other.GetUpdateTime()) {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorUpdateTime})
	}
	if o.GetUuid() != other.GetUuid() {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorUuid})
	}

	if len(o.GetTags()) == len(other.GetTags()) {
		for i, lValue := range o.GetTags() {
			rValue := other.GetTags()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorTags})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorTags})
	}

	if len(o.GetLabels()) == len(other.GetLabels()) {
		for i, lValue := range o.GetLabels() {
			rValue := other.GetLabels()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorLabels})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorLabels})
	}

	if len(o.GetAnnotations()) == len(other.GetAnnotations()) {
		for i, lValue := range o.GetAnnotations() {
			rValue := other.GetAnnotations()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorAnnotations})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorAnnotations})
	}
	if o.GetGeneration() != other.GetGeneration() {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorGeneration})
	}
	if o.GetResourceVersion() != other.GetResourceVersion() {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorResourceVersion})
	}

	if len(o.GetOwnerReferences()) == len(other.GetOwnerReferences()) {
		for i, lValue := range o.GetOwnerReferences() {
			rValue := other.GetOwnerReferences()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorOwnerReferences})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorOwnerReferences})
	}

	if len(o.GetShards()) == len(other.GetShards()) {
		for i, lValue := range o.GetShards() {
			rValue := other.GetShards()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorShards})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorShards})
	}
	{
		subMask := o.GetSyncing().MakeDiffFieldMask(other.GetSyncing())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorSyncing})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Meta_FieldSubPath{selector: Meta_FieldPathSelectorSyncing, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Meta) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Meta))
}

func (o *Meta) Clone() *Meta {
	if o == nil {
		return nil
	}
	result := &Meta{}
	result.CreateTime = proto.Clone(o.CreateTime).(*timestamp.Timestamp)
	result.UpdateTime = proto.Clone(o.UpdateTime).(*timestamp.Timestamp)
	result.Uuid = o.Uuid
	result.Tags = make([]string, len(o.Tags))
	for i, sourceValue := range o.Tags {
		result.Tags[i] = sourceValue
	}
	result.Labels = map[string]string{}
	for key, sourceValue := range o.Labels {
		result.Labels[key] = sourceValue
	}
	result.Annotations = map[string]string{}
	for key, sourceValue := range o.Annotations {
		result.Annotations[key] = sourceValue
	}
	result.Generation = o.Generation
	result.ResourceVersion = o.ResourceVersion
	result.OwnerReferences = make([]*OwnerReference, len(o.OwnerReferences))
	for i, sourceValue := range o.OwnerReferences {
		result.OwnerReferences[i] = sourceValue.Clone()
	}
	result.Shards = map[string]int64{}
	for key, sourceValue := range o.Shards {
		result.Shards[key] = sourceValue
	}
	result.Syncing = o.Syncing.Clone()
	return result
}

func (o *Meta) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Meta) Merge(source *Meta) {
	if source.GetCreateTime() != nil {
		if o.CreateTime == nil {
			o.CreateTime = new(timestamp.Timestamp)
		}
		proto.Merge(o.CreateTime, source.GetCreateTime())
	}
	if source.GetUpdateTime() != nil {
		if o.UpdateTime == nil {
			o.UpdateTime = new(timestamp.Timestamp)
		}
		proto.Merge(o.UpdateTime, source.GetUpdateTime())
	}
	o.Uuid = source.GetUuid()
	for _, sourceValue := range source.GetTags() {
		exists := false
		for _, currentValue := range o.Tags {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.Tags = append(o.Tags, newDstElement)
		}
	}

	if source.GetLabels() != nil {
		if o.Labels == nil {
			o.Labels = make(map[string]string, len(source.GetLabels()))
		}
		for key, sourceValue := range source.GetLabels() {
			o.Labels[key] = sourceValue
		}
	}
	if source.GetAnnotations() != nil {
		if o.Annotations == nil {
			o.Annotations = make(map[string]string, len(source.GetAnnotations()))
		}
		for key, sourceValue := range source.GetAnnotations() {
			o.Annotations[key] = sourceValue
		}
	}
	o.Generation = source.GetGeneration()
	o.ResourceVersion = source.GetResourceVersion()
	sourceOwnerReferences, origOwnerReferencesKeys := map[string]*OwnerReference{}, map[string]bool{}
	newOwnerReferences := make([]*OwnerReference, 0, len(o.OwnerReferences))
	for _, sourceValue := range source.GetOwnerReferences() {
		key := fmt.Sprintf("%s", sourceValue.GetName())
		sourceOwnerReferences[key] = sourceValue
	}
	for _, origValue := range o.OwnerReferences {
		key := fmt.Sprintf("%s", origValue.GetName())
		origOwnerReferencesKeys[key] = true
		sourceValue := sourceOwnerReferences[key]
		if sourceValue != nil {
			if origValue == nil {
				origValue = new(OwnerReference)
			}
			origValue.Merge(sourceValue)
		}
		newOwnerReferences = append(newOwnerReferences, origValue)
	}
	for key, sourceValue := range sourceOwnerReferences {
		if origOwnerReferencesKeys[key] == false {
			newOwnerReferences = append(newOwnerReferences, sourceValue.Clone())
		}
	}
	o.OwnerReferences = newOwnerReferences

	if source.GetShards() != nil {
		if o.Shards == nil {
			o.Shards = make(map[string]int64, len(source.GetShards()))
		}
		for key, sourceValue := range source.GetShards() {
			o.Shards[key] = sourceValue
		}
	}
	if source.GetSyncing() != nil {
		if o.Syncing == nil {
			o.Syncing = new(SyncingMeta)
		}
		o.Syncing.Merge(source.GetSyncing())
	}
}

func (o *Meta) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Meta))
}

func (o *LabelSelector) GotenObjectExt() {}

func (o *LabelSelector) MakeFullFieldMask() *LabelSelector_FieldMask {
	return FullLabelSelector_FieldMask()
}

func (o *LabelSelector) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLabelSelector_FieldMask()
}

func (o *LabelSelector) MakeDiffFieldMask(other *LabelSelector) *LabelSelector_FieldMask {
	if o == nil && other == nil {
		return &LabelSelector_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLabelSelector_FieldMask()
	}

	res := &LabelSelector_FieldMask{}

	if len(o.GetMatchLabels()) == len(other.GetMatchLabels()) {
		for i, lValue := range o.GetMatchLabels() {
			rValue := other.GetMatchLabels()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &LabelSelector_FieldTerminalPath{selector: LabelSelector_FieldPathSelectorMatchLabels})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &LabelSelector_FieldTerminalPath{selector: LabelSelector_FieldPathSelectorMatchLabels})
	}

	if len(o.GetMatchExpressions()) == len(other.GetMatchExpressions()) {
		for i, lValue := range o.GetMatchExpressions() {
			rValue := other.GetMatchExpressions()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &LabelSelector_FieldTerminalPath{selector: LabelSelector_FieldPathSelectorMatchExpressions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &LabelSelector_FieldTerminalPath{selector: LabelSelector_FieldPathSelectorMatchExpressions})
	}
	return res
}

func (o *LabelSelector) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*LabelSelector))
}

func (o *LabelSelector) Clone() *LabelSelector {
	if o == nil {
		return nil
	}
	result := &LabelSelector{}
	result.MatchLabels = map[string]string{}
	for key, sourceValue := range o.MatchLabels {
		result.MatchLabels[key] = sourceValue
	}
	result.MatchExpressions = make([]*LabelSelectorRequirement, len(o.MatchExpressions))
	for i, sourceValue := range o.MatchExpressions {
		result.MatchExpressions[i] = sourceValue.Clone()
	}
	return result
}

func (o *LabelSelector) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *LabelSelector) Merge(source *LabelSelector) {
	if source.GetMatchLabels() != nil {
		if o.MatchLabels == nil {
			o.MatchLabels = make(map[string]string, len(source.GetMatchLabels()))
		}
		for key, sourceValue := range source.GetMatchLabels() {
			o.MatchLabels[key] = sourceValue
		}
	}
	sourceMatchExpressions, origMatchExpressionsKeys := map[string]*LabelSelectorRequirement{}, map[string]bool{}
	newMatchExpressions := make([]*LabelSelectorRequirement, 0, len(o.MatchExpressions))
	for _, sourceValue := range source.GetMatchExpressions() {
		key := fmt.Sprintf("%s", sourceValue.GetKey())
		sourceMatchExpressions[key] = sourceValue
	}
	for _, origValue := range o.MatchExpressions {
		key := fmt.Sprintf("%s", origValue.GetKey())
		origMatchExpressionsKeys[key] = true
		sourceValue := sourceMatchExpressions[key]
		if sourceValue != nil {
			if origValue == nil {
				origValue = new(LabelSelectorRequirement)
			}
			origValue.Merge(sourceValue)
		}
		newMatchExpressions = append(newMatchExpressions, origValue)
	}
	for key, sourceValue := range sourceMatchExpressions {
		if origMatchExpressionsKeys[key] == false {
			newMatchExpressions = append(newMatchExpressions, sourceValue.Clone())
		}
	}
	o.MatchExpressions = newMatchExpressions

}

func (o *LabelSelector) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*LabelSelector))
}

func (o *LabelSelectorRequirement) GotenObjectExt() {}

func (o *LabelSelectorRequirement) MakeFullFieldMask() *LabelSelectorRequirement_FieldMask {
	return FullLabelSelectorRequirement_FieldMask()
}

func (o *LabelSelectorRequirement) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLabelSelectorRequirement_FieldMask()
}

func (o *LabelSelectorRequirement) MakeDiffFieldMask(other *LabelSelectorRequirement) *LabelSelectorRequirement_FieldMask {
	if o == nil && other == nil {
		return &LabelSelectorRequirement_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLabelSelectorRequirement_FieldMask()
	}

	res := &LabelSelectorRequirement_FieldMask{}
	if o.GetKey() != other.GetKey() {
		res.Paths = append(res.Paths, &LabelSelectorRequirement_FieldTerminalPath{selector: LabelSelectorRequirement_FieldPathSelectorKey})
	}
	if o.GetOperator() != other.GetOperator() {
		res.Paths = append(res.Paths, &LabelSelectorRequirement_FieldTerminalPath{selector: LabelSelectorRequirement_FieldPathSelectorOperator})
	}

	if len(o.GetValues()) == len(other.GetValues()) {
		for i, lValue := range o.GetValues() {
			rValue := other.GetValues()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &LabelSelectorRequirement_FieldTerminalPath{selector: LabelSelectorRequirement_FieldPathSelectorValues})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &LabelSelectorRequirement_FieldTerminalPath{selector: LabelSelectorRequirement_FieldPathSelectorValues})
	}
	return res
}

func (o *LabelSelectorRequirement) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*LabelSelectorRequirement))
}

func (o *LabelSelectorRequirement) Clone() *LabelSelectorRequirement {
	if o == nil {
		return nil
	}
	result := &LabelSelectorRequirement{}
	result.Key = o.Key
	result.Operator = o.Operator
	result.Values = make([]string, len(o.Values))
	for i, sourceValue := range o.Values {
		result.Values[i] = sourceValue
	}
	return result
}

func (o *LabelSelectorRequirement) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *LabelSelectorRequirement) Merge(source *LabelSelectorRequirement) {
	o.Key = source.GetKey()
	o.Operator = source.GetOperator()
	o.Values = make([]string, 0, len(source.GetValues()))
	for _, sourceValue := range source.GetValues() {
		var newDstElement string
		newDstElement = sourceValue
		o.Values = append(o.Values, newDstElement)
	}

}

func (o *LabelSelectorRequirement) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*LabelSelectorRequirement))
}

func (o *OwnerReference) GotenObjectExt() {}

func (o *OwnerReference) MakeFullFieldMask() *OwnerReference_FieldMask {
	return FullOwnerReference_FieldMask()
}

func (o *OwnerReference) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullOwnerReference_FieldMask()
}

func (o *OwnerReference) MakeDiffFieldMask(other *OwnerReference) *OwnerReference_FieldMask {
	if o == nil && other == nil {
		return &OwnerReference_FieldMask{}
	}
	if o == nil || other == nil {
		return FullOwnerReference_FieldMask()
	}

	res := &OwnerReference_FieldMask{}
	if o.GetApiVersion() != other.GetApiVersion() {
		res.Paths = append(res.Paths, &OwnerReference_FieldTerminalPath{selector: OwnerReference_FieldPathSelectorApiVersion})
	}
	if o.GetKind() != other.GetKind() {
		res.Paths = append(res.Paths, &OwnerReference_FieldTerminalPath{selector: OwnerReference_FieldPathSelectorKind})
	}
	if o.GetName() != other.GetName() {
		res.Paths = append(res.Paths, &OwnerReference_FieldTerminalPath{selector: OwnerReference_FieldPathSelectorName})
	}
	if o.GetUid() != other.GetUid() {
		res.Paths = append(res.Paths, &OwnerReference_FieldTerminalPath{selector: OwnerReference_FieldPathSelectorUid})
	}
	if o.GetController() != other.GetController() {
		res.Paths = append(res.Paths, &OwnerReference_FieldTerminalPath{selector: OwnerReference_FieldPathSelectorController})
	}
	if o.GetBlockOwnerDeletion() != other.GetBlockOwnerDeletion() {
		res.Paths = append(res.Paths, &OwnerReference_FieldTerminalPath{selector: OwnerReference_FieldPathSelectorBlockOwnerDeletion})
	}
	return res
}

func (o *OwnerReference) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*OwnerReference))
}

func (o *OwnerReference) Clone() *OwnerReference {
	if o == nil {
		return nil
	}
	result := &OwnerReference{}
	result.ApiVersion = o.ApiVersion
	result.Kind = o.Kind
	result.Name = o.Name
	result.Uid = o.Uid
	result.Controller = o.Controller
	result.BlockOwnerDeletion = o.BlockOwnerDeletion
	return result
}

func (o *OwnerReference) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *OwnerReference) Merge(source *OwnerReference) {
	o.ApiVersion = source.GetApiVersion()
	o.Kind = source.GetKind()
	o.Name = source.GetName()
	o.Uid = source.GetUid()
	o.Controller = source.GetController()
	o.BlockOwnerDeletion = source.GetBlockOwnerDeletion()
}

func (o *OwnerReference) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*OwnerReference))
}

func (o *SyncingMeta) GotenObjectExt() {}

func (o *SyncingMeta) MakeFullFieldMask() *SyncingMeta_FieldMask {
	return FullSyncingMeta_FieldMask()
}

func (o *SyncingMeta) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullSyncingMeta_FieldMask()
}

func (o *SyncingMeta) MakeDiffFieldMask(other *SyncingMeta) *SyncingMeta_FieldMask {
	if o == nil && other == nil {
		return &SyncingMeta_FieldMask{}
	}
	if o == nil || other == nil {
		return FullSyncingMeta_FieldMask()
	}

	res := &SyncingMeta_FieldMask{}
	if o.GetOwningRegion() != other.GetOwningRegion() {
		res.Paths = append(res.Paths, &SyncingMeta_FieldTerminalPath{selector: SyncingMeta_FieldPathSelectorOwningRegion})
	}

	if len(o.GetRegions()) == len(other.GetRegions()) {
		for i, lValue := range o.GetRegions() {
			rValue := other.GetRegions()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &SyncingMeta_FieldTerminalPath{selector: SyncingMeta_FieldPathSelectorRegions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &SyncingMeta_FieldTerminalPath{selector: SyncingMeta_FieldPathSelectorRegions})
	}
	return res
}

func (o *SyncingMeta) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*SyncingMeta))
}

func (o *SyncingMeta) Clone() *SyncingMeta {
	if o == nil {
		return nil
	}
	result := &SyncingMeta{}
	result.OwningRegion = o.OwningRegion
	result.Regions = make([]string, len(o.Regions))
	for i, sourceValue := range o.Regions {
		result.Regions[i] = sourceValue
	}
	return result
}

func (o *SyncingMeta) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *SyncingMeta) Merge(source *SyncingMeta) {
	o.OwningRegion = source.GetOwningRegion()
	for _, sourceValue := range source.GetRegions() {
		exists := false
		for _, currentValue := range o.Regions {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.Regions = append(o.Regions, newDstElement)
		}
	}

}

func (o *SyncingMeta) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*SyncingMeta))
}
