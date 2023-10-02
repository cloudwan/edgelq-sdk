// Code generated by protoc-gen-goten-object
// File: edgelq/applications/proto/v1alpha2/pod.proto
// DO NOT EDIT!!!

package pod

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	common "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha2/common"
	distribution "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha2/distribution"
	project "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha2/project"
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	_ = &common.PodSpec{}
	_ = &distribution.Distribution{}
	_ = &project.Project{}
	_ = &ntt_meta.Meta{}
	_ = &timestamp.Timestamp{}
)

func (o *Pod) GotenObjectExt() {}

func (o *Pod) MakeFullFieldMask() *Pod_FieldMask {
	return FullPod_FieldMask()
}

func (o *Pod) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPod_FieldMask()
}

func (o *Pod) MakeDiffFieldMask(other *Pod) *Pod_FieldMask {
	if o == nil && other == nil {
		return &Pod_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPod_FieldMask()
	}

	res := &Pod_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &Pod_FieldTerminalPath{selector: Pod_FieldPathSelectorName})
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &Pod_FieldTerminalPath{selector: Pod_FieldPathSelectorDisplayName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Pod_FieldTerminalPath{selector: Pod_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Pod_FieldSubPath{selector: Pod_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetSpec().MakeDiffFieldMask(other.GetSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Pod_FieldTerminalPath{selector: Pod_FieldPathSelectorSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Pod_FieldSubPath{selector: Pod_FieldPathSelectorSpec, subPath: subpath})
			}
		}
	}
	if o.GetDistribution().String() != other.GetDistribution().String() {
		res.Paths = append(res.Paths, &Pod_FieldTerminalPath{selector: Pod_FieldPathSelectorDistribution})
	}
	{
		subMask := o.GetStatus().MakeDiffFieldMask(other.GetStatus())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Pod_FieldTerminalPath{selector: Pod_FieldPathSelectorStatus})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Pod_FieldSubPath{selector: Pod_FieldPathSelectorStatus, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Pod) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Pod))
}

func (o *Pod) Clone() *Pod {
	if o == nil {
		return nil
	}
	result := &Pod{}
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
	result.DisplayName = o.DisplayName
	result.Metadata = o.Metadata.Clone()
	result.Spec = o.Spec.Clone()
	if o.Distribution == nil {
		result.Distribution = nil
	} else if data, err := o.Distribution.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Distribution = &distribution.Reference{}
		if err := result.Distribution.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Status = o.Status.Clone()
	return result
}

func (o *Pod) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Pod) Merge(source *Pod) {
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
	o.DisplayName = source.GetDisplayName()
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(ntt_meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
	if source.GetSpec() != nil {
		if o.Spec == nil {
			o.Spec = new(common.PodSpec)
		}
		o.Spec.Merge(source.GetSpec())
	}
	if source.GetDistribution() != nil {
		if data, err := source.GetDistribution().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Distribution = &distribution.Reference{}
			if err := o.Distribution.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Distribution = nil
	}
	if source.GetStatus() != nil {
		if o.Status == nil {
			o.Status = new(Pod_Status)
		}
		o.Status.Merge(source.GetStatus())
	}
}

func (o *Pod) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Pod))
}

func (o *Pod_Status) GotenObjectExt() {}

func (o *Pod_Status) MakeFullFieldMask() *Pod_Status_FieldMask {
	return FullPod_Status_FieldMask()
}

func (o *Pod_Status) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPod_Status_FieldMask()
}

func (o *Pod_Status) MakeDiffFieldMask(other *Pod_Status) *Pod_Status_FieldMask {
	if o == nil && other == nil {
		return &Pod_Status_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPod_Status_FieldMask()
	}

	res := &Pod_Status_FieldMask{}
	if o.GetPhase() != other.GetPhase() {
		res.Paths = append(res.Paths, &PodStatus_FieldTerminalPath{selector: PodStatus_FieldPathSelectorPhase})
	}

	if len(o.GetContainerStatuses()) == len(other.GetContainerStatuses()) {
		for i, lValue := range o.GetContainerStatuses() {
			rValue := other.GetContainerStatuses()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &PodStatus_FieldTerminalPath{selector: PodStatus_FieldPathSelectorContainerStatuses})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &PodStatus_FieldTerminalPath{selector: PodStatus_FieldPathSelectorContainerStatuses})
	}
	return res
}

func (o *Pod_Status) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Pod_Status))
}

func (o *Pod_Status) Clone() *Pod_Status {
	if o == nil {
		return nil
	}
	result := &Pod_Status{}
	result.Phase = o.Phase
	result.ContainerStatuses = make([]*Pod_Status_Container, len(o.ContainerStatuses))
	for i, sourceValue := range o.ContainerStatuses {
		result.ContainerStatuses[i] = sourceValue.Clone()
	}
	return result
}

func (o *Pod_Status) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Pod_Status) Merge(source *Pod_Status) {
	o.Phase = source.GetPhase()
	for _, sourceValue := range source.GetContainerStatuses() {
		exists := false
		for _, currentValue := range o.ContainerStatuses {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *Pod_Status_Container
			if sourceValue != nil {
				newDstElement = new(Pod_Status_Container)
				newDstElement.Merge(sourceValue)
			}
			o.ContainerStatuses = append(o.ContainerStatuses, newDstElement)
		}
	}

}

func (o *Pod_Status) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Pod_Status))
}

func (o *Pod_Status_Container) GotenObjectExt() {}

func (o *Pod_Status_Container) MakeFullFieldMask() *Pod_Status_Container_FieldMask {
	return FullPod_Status_Container_FieldMask()
}

func (o *Pod_Status_Container) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPod_Status_Container_FieldMask()
}

func (o *Pod_Status_Container) MakeDiffFieldMask(other *Pod_Status_Container) *Pod_Status_Container_FieldMask {
	if o == nil && other == nil {
		return &Pod_Status_Container_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPod_Status_Container_FieldMask()
	}

	res := &Pod_Status_Container_FieldMask{}
	if o.GetName() != other.GetName() {
		res.Paths = append(res.Paths, &PodStatusContainer_FieldTerminalPath{selector: PodStatusContainer_FieldPathSelectorName})
	}
	if o.GetState() != other.GetState() {
		res.Paths = append(res.Paths, &PodStatusContainer_FieldTerminalPath{selector: PodStatusContainer_FieldPathSelectorState})
	}
	{
		subMask := o.GetWaiting().MakeDiffFieldMask(other.GetWaiting())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &PodStatusContainer_FieldTerminalPath{selector: PodStatusContainer_FieldPathSelectorWaiting})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &PodStatusContainer_FieldSubPath{selector: PodStatusContainer_FieldPathSelectorWaiting, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetRunning().MakeDiffFieldMask(other.GetRunning())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &PodStatusContainer_FieldTerminalPath{selector: PodStatusContainer_FieldPathSelectorRunning})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &PodStatusContainer_FieldSubPath{selector: PodStatusContainer_FieldPathSelectorRunning, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetTerminated().MakeDiffFieldMask(other.GetTerminated())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &PodStatusContainer_FieldTerminalPath{selector: PodStatusContainer_FieldPathSelectorTerminated})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &PodStatusContainer_FieldSubPath{selector: PodStatusContainer_FieldPathSelectorTerminated, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Pod_Status_Container) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Pod_Status_Container))
}

func (o *Pod_Status_Container) Clone() *Pod_Status_Container {
	if o == nil {
		return nil
	}
	result := &Pod_Status_Container{}
	result.Name = o.Name
	result.State = o.State
	result.Waiting = o.Waiting.Clone()
	result.Running = o.Running.Clone()
	result.Terminated = o.Terminated.Clone()
	return result
}

func (o *Pod_Status_Container) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Pod_Status_Container) Merge(source *Pod_Status_Container) {
	o.Name = source.GetName()
	o.State = source.GetState()
	if source.GetWaiting() != nil {
		if o.Waiting == nil {
			o.Waiting = new(Pod_Status_Container_StateWaiting)
		}
		o.Waiting.Merge(source.GetWaiting())
	}
	if source.GetRunning() != nil {
		if o.Running == nil {
			o.Running = new(Pod_Status_Container_StateRunning)
		}
		o.Running.Merge(source.GetRunning())
	}
	if source.GetTerminated() != nil {
		if o.Terminated == nil {
			o.Terminated = new(Pod_Status_Container_StateTerminated)
		}
		o.Terminated.Merge(source.GetTerminated())
	}
}

func (o *Pod_Status_Container) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Pod_Status_Container))
}

func (o *Pod_Status_Container_StateWaiting) GotenObjectExt() {}

func (o *Pod_Status_Container_StateWaiting) MakeFullFieldMask() *Pod_Status_Container_StateWaiting_FieldMask {
	return FullPod_Status_Container_StateWaiting_FieldMask()
}

func (o *Pod_Status_Container_StateWaiting) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPod_Status_Container_StateWaiting_FieldMask()
}

func (o *Pod_Status_Container_StateWaiting) MakeDiffFieldMask(other *Pod_Status_Container_StateWaiting) *Pod_Status_Container_StateWaiting_FieldMask {
	if o == nil && other == nil {
		return &Pod_Status_Container_StateWaiting_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPod_Status_Container_StateWaiting_FieldMask()
	}

	res := &Pod_Status_Container_StateWaiting_FieldMask{}
	if o.GetReason() != other.GetReason() {
		res.Paths = append(res.Paths, &PodStatusContainerStateWaiting_FieldTerminalPath{selector: PodStatusContainerStateWaiting_FieldPathSelectorReason})
	}
	if o.GetMessage() != other.GetMessage() {
		res.Paths = append(res.Paths, &PodStatusContainerStateWaiting_FieldTerminalPath{selector: PodStatusContainerStateWaiting_FieldPathSelectorMessage})
	}
	return res
}

func (o *Pod_Status_Container_StateWaiting) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Pod_Status_Container_StateWaiting))
}

func (o *Pod_Status_Container_StateWaiting) Clone() *Pod_Status_Container_StateWaiting {
	if o == nil {
		return nil
	}
	result := &Pod_Status_Container_StateWaiting{}
	result.Reason = o.Reason
	result.Message = o.Message
	return result
}

func (o *Pod_Status_Container_StateWaiting) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Pod_Status_Container_StateWaiting) Merge(source *Pod_Status_Container_StateWaiting) {
	o.Reason = source.GetReason()
	o.Message = source.GetMessage()
}

func (o *Pod_Status_Container_StateWaiting) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Pod_Status_Container_StateWaiting))
}

func (o *Pod_Status_Container_StateRunning) GotenObjectExt() {}

func (o *Pod_Status_Container_StateRunning) MakeFullFieldMask() *Pod_Status_Container_StateRunning_FieldMask {
	return FullPod_Status_Container_StateRunning_FieldMask()
}

func (o *Pod_Status_Container_StateRunning) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPod_Status_Container_StateRunning_FieldMask()
}

func (o *Pod_Status_Container_StateRunning) MakeDiffFieldMask(other *Pod_Status_Container_StateRunning) *Pod_Status_Container_StateRunning_FieldMask {
	if o == nil && other == nil {
		return &Pod_Status_Container_StateRunning_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPod_Status_Container_StateRunning_FieldMask()
	}

	res := &Pod_Status_Container_StateRunning_FieldMask{}
	if !proto.Equal(o.GetStartedAt(), other.GetStartedAt()) {
		res.Paths = append(res.Paths, &PodStatusContainerStateRunning_FieldTerminalPath{selector: PodStatusContainerStateRunning_FieldPathSelectorStartedAt})
	}
	return res
}

func (o *Pod_Status_Container_StateRunning) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Pod_Status_Container_StateRunning))
}

func (o *Pod_Status_Container_StateRunning) Clone() *Pod_Status_Container_StateRunning {
	if o == nil {
		return nil
	}
	result := &Pod_Status_Container_StateRunning{}
	result.StartedAt = proto.Clone(o.StartedAt).(*timestamp.Timestamp)
	return result
}

func (o *Pod_Status_Container_StateRunning) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Pod_Status_Container_StateRunning) Merge(source *Pod_Status_Container_StateRunning) {
	if source.GetStartedAt() != nil {
		if o.StartedAt == nil {
			o.StartedAt = new(timestamp.Timestamp)
		}
		proto.Merge(o.StartedAt, source.GetStartedAt())
	}
}

func (o *Pod_Status_Container_StateRunning) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Pod_Status_Container_StateRunning))
}

func (o *Pod_Status_Container_StateTerminated) GotenObjectExt() {}

func (o *Pod_Status_Container_StateTerminated) MakeFullFieldMask() *Pod_Status_Container_StateTerminated_FieldMask {
	return FullPod_Status_Container_StateTerminated_FieldMask()
}

func (o *Pod_Status_Container_StateTerminated) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullPod_Status_Container_StateTerminated_FieldMask()
}

func (o *Pod_Status_Container_StateTerminated) MakeDiffFieldMask(other *Pod_Status_Container_StateTerminated) *Pod_Status_Container_StateTerminated_FieldMask {
	if o == nil && other == nil {
		return &Pod_Status_Container_StateTerminated_FieldMask{}
	}
	if o == nil || other == nil {
		return FullPod_Status_Container_StateTerminated_FieldMask()
	}

	res := &Pod_Status_Container_StateTerminated_FieldMask{}
	if o.GetExitCode() != other.GetExitCode() {
		res.Paths = append(res.Paths, &PodStatusContainerStateTerminated_FieldTerminalPath{selector: PodStatusContainerStateTerminated_FieldPathSelectorExitCode})
	}
	if o.GetSignal() != other.GetSignal() {
		res.Paths = append(res.Paths, &PodStatusContainerStateTerminated_FieldTerminalPath{selector: PodStatusContainerStateTerminated_FieldPathSelectorSignal})
	}
	if o.GetReason() != other.GetReason() {
		res.Paths = append(res.Paths, &PodStatusContainerStateTerminated_FieldTerminalPath{selector: PodStatusContainerStateTerminated_FieldPathSelectorReason})
	}
	if o.GetMessage() != other.GetMessage() {
		res.Paths = append(res.Paths, &PodStatusContainerStateTerminated_FieldTerminalPath{selector: PodStatusContainerStateTerminated_FieldPathSelectorMessage})
	}
	if !proto.Equal(o.GetStartedAt(), other.GetStartedAt()) {
		res.Paths = append(res.Paths, &PodStatusContainerStateTerminated_FieldTerminalPath{selector: PodStatusContainerStateTerminated_FieldPathSelectorStartedAt})
	}
	if !proto.Equal(o.GetFinishedAt(), other.GetFinishedAt()) {
		res.Paths = append(res.Paths, &PodStatusContainerStateTerminated_FieldTerminalPath{selector: PodStatusContainerStateTerminated_FieldPathSelectorFinishedAt})
	}
	if o.GetContainerId() != other.GetContainerId() {
		res.Paths = append(res.Paths, &PodStatusContainerStateTerminated_FieldTerminalPath{selector: PodStatusContainerStateTerminated_FieldPathSelectorContainerId})
	}
	return res
}

func (o *Pod_Status_Container_StateTerminated) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Pod_Status_Container_StateTerminated))
}

func (o *Pod_Status_Container_StateTerminated) Clone() *Pod_Status_Container_StateTerminated {
	if o == nil {
		return nil
	}
	result := &Pod_Status_Container_StateTerminated{}
	result.ExitCode = o.ExitCode
	result.Signal = o.Signal
	result.Reason = o.Reason
	result.Message = o.Message
	result.StartedAt = proto.Clone(o.StartedAt).(*timestamp.Timestamp)
	result.FinishedAt = proto.Clone(o.FinishedAt).(*timestamp.Timestamp)
	result.ContainerId = o.ContainerId
	return result
}

func (o *Pod_Status_Container_StateTerminated) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Pod_Status_Container_StateTerminated) Merge(source *Pod_Status_Container_StateTerminated) {
	o.ExitCode = source.GetExitCode()
	o.Signal = source.GetSignal()
	o.Reason = source.GetReason()
	o.Message = source.GetMessage()
	if source.GetStartedAt() != nil {
		if o.StartedAt == nil {
			o.StartedAt = new(timestamp.Timestamp)
		}
		proto.Merge(o.StartedAt, source.GetStartedAt())
	}
	if source.GetFinishedAt() != nil {
		if o.FinishedAt == nil {
			o.FinishedAt = new(timestamp.Timestamp)
		}
		proto.Merge(o.FinishedAt, source.GetFinishedAt())
	}
	o.ContainerId = source.GetContainerId()
}

func (o *Pod_Status_Container_StateTerminated) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Pod_Status_Container_StateTerminated))
}
