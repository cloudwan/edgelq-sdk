// Code generated by protoc-gen-goten-object
// File: edgelq/logging/proto/v1alpha2/log_custom.proto
// DO NOT EDIT!!!

package log_client

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	rpc "github.com/cloudwan/edgelq-sdk/common/rpc"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	logging_common "github.com/cloudwan/edgelq-sdk/logging/common/v1alpha2"
	log "github.com/cloudwan/edgelq-sdk/logging/resources/v1alpha2/log"
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
	_ = &rpc.Status{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &logging_common.LabelDescriptor{}
	_ = &log.Log{}
)

func (o *ListLogsRequest) GotenObjectExt() {}

func (o *ListLogsRequest) MakeFullFieldMask() *ListLogsRequest_FieldMask {
	return FullListLogsRequest_FieldMask()
}

func (o *ListLogsRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullListLogsRequest_FieldMask()
}

func (o *ListLogsRequest) MakeDiffFieldMask(other *ListLogsRequest) *ListLogsRequest_FieldMask {
	if o == nil && other == nil {
		return &ListLogsRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullListLogsRequest_FieldMask()
	}

	res := &ListLogsRequest_FieldMask{}

	if len(o.GetParents()) == len(other.GetParents()) {
		for i, lValue := range o.GetParents() {
			rValue := other.GetParents()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &ListLogsRequest_FieldTerminalPath{selector: ListLogsRequest_FieldPathSelectorParents})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ListLogsRequest_FieldTerminalPath{selector: ListLogsRequest_FieldPathSelectorParents})
	}
	if o.GetFilter().String() != other.GetFilter().String() {
		res.Paths = append(res.Paths, &ListLogsRequest_FieldTerminalPath{selector: ListLogsRequest_FieldPathSelectorFilter})
	}
	{
		subMask := o.GetInterval().MakeDiffFieldMask(other.GetInterval())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ListLogsRequest_FieldTerminalPath{selector: ListLogsRequest_FieldPathSelectorInterval})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ListLogsRequest_FieldSubPath{selector: ListLogsRequest_FieldPathSelectorInterval, subPath: subpath})
			}
		}
	}
	if o.GetPageSize() != other.GetPageSize() {
		res.Paths = append(res.Paths, &ListLogsRequest_FieldTerminalPath{selector: ListLogsRequest_FieldPathSelectorPageSize})
	}
	if o.GetPageToken() != other.GetPageToken() {
		res.Paths = append(res.Paths, &ListLogsRequest_FieldTerminalPath{selector: ListLogsRequest_FieldPathSelectorPageToken})
	}
	return res
}

func (o *ListLogsRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ListLogsRequest))
}

func (o *ListLogsRequest) Clone() *ListLogsRequest {
	if o == nil {
		return nil
	}
	result := &ListLogsRequest{}
	result.Parents = make([]*log.ParentReference, len(o.Parents))
	for i, sourceValue := range o.Parents {
		if sourceValue == nil {
			result.Parents[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.Parents[i] = &log.ParentReference{}
			if err := result.Parents[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	if o.Filter == nil {
		result.Filter = nil
	} else if data, err := o.Filter.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Filter = &log.Filter{}
		if err := result.Filter.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Interval = o.Interval.Clone()
	result.PageSize = o.PageSize
	result.PageToken = o.PageToken
	return result
}

func (o *ListLogsRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ListLogsRequest) Merge(source *ListLogsRequest) {
	for _, sourceValue := range source.GetParents() {
		exists := false
		for _, currentValue := range o.Parents {
			leftProtoStr, _ := currentValue.ProtoString()
			rightProtoStr, _ := sourceValue.ProtoString()
			if leftProtoStr == rightProtoStr {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *log.ParentReference
			if sourceValue != nil {
				if data, err := sourceValue.ProtoString(); err != nil {
					panic(err)
				} else {
					newDstElement = &log.ParentReference{}
					if err := newDstElement.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			}
			o.Parents = append(o.Parents, newDstElement)
		}
	}

	if source.GetFilter() != nil {
		if data, err := source.GetFilter().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Filter = &log.Filter{}
			if err := o.Filter.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Filter = nil
	}
	if source.GetInterval() != nil {
		if o.Interval == nil {
			o.Interval = new(logging_common.TimeInterval)
		}
		o.Interval.Merge(source.GetInterval())
	}
	o.PageSize = source.GetPageSize()
	o.PageToken = source.GetPageToken()
}

func (o *ListLogsRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ListLogsRequest))
}

func (o *ListLogsResponse) GotenObjectExt() {}

func (o *ListLogsResponse) MakeFullFieldMask() *ListLogsResponse_FieldMask {
	return FullListLogsResponse_FieldMask()
}

func (o *ListLogsResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullListLogsResponse_FieldMask()
}

func (o *ListLogsResponse) MakeDiffFieldMask(other *ListLogsResponse) *ListLogsResponse_FieldMask {
	if o == nil && other == nil {
		return &ListLogsResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullListLogsResponse_FieldMask()
	}

	res := &ListLogsResponse_FieldMask{}

	if len(o.GetLogs()) == len(other.GetLogs()) {
		for i, lValue := range o.GetLogs() {
			rValue := other.GetLogs()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &ListLogsResponse_FieldTerminalPath{selector: ListLogsResponse_FieldPathSelectorLogs})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ListLogsResponse_FieldTerminalPath{selector: ListLogsResponse_FieldPathSelectorLogs})
	}
	if o.GetNextPageToken() != other.GetNextPageToken() {
		res.Paths = append(res.Paths, &ListLogsResponse_FieldTerminalPath{selector: ListLogsResponse_FieldPathSelectorNextPageToken})
	}

	if len(o.GetExecutionErrors()) == len(other.GetExecutionErrors()) {
		for i, lValue := range o.GetExecutionErrors() {
			rValue := other.GetExecutionErrors()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &ListLogsResponse_FieldTerminalPath{selector: ListLogsResponse_FieldPathSelectorExecutionErrors})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ListLogsResponse_FieldTerminalPath{selector: ListLogsResponse_FieldPathSelectorExecutionErrors})
	}
	return res
}

func (o *ListLogsResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ListLogsResponse))
}

func (o *ListLogsResponse) Clone() *ListLogsResponse {
	if o == nil {
		return nil
	}
	result := &ListLogsResponse{}
	result.Logs = make([]*log.Log, len(o.Logs))
	for i, sourceValue := range o.Logs {
		result.Logs[i] = sourceValue.Clone()
	}
	result.NextPageToken = o.NextPageToken
	result.ExecutionErrors = make([]*rpc.Status, len(o.ExecutionErrors))
	for i, sourceValue := range o.ExecutionErrors {
		result.ExecutionErrors[i] = sourceValue.Clone()
	}
	return result
}

func (o *ListLogsResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ListLogsResponse) Merge(source *ListLogsResponse) {
	for _, sourceValue := range source.GetLogs() {
		exists := false
		for _, currentValue := range o.Logs {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *log.Log
			if sourceValue != nil {
				newDstElement = new(log.Log)
				newDstElement.Merge(sourceValue)
			}
			o.Logs = append(o.Logs, newDstElement)
		}
	}

	o.NextPageToken = source.GetNextPageToken()
	for _, sourceValue := range source.GetExecutionErrors() {
		exists := false
		for _, currentValue := range o.ExecutionErrors {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *rpc.Status
			if sourceValue != nil {
				newDstElement = new(rpc.Status)
				newDstElement.Merge(sourceValue)
			}
			o.ExecutionErrors = append(o.ExecutionErrors, newDstElement)
		}
	}

}

func (o *ListLogsResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ListLogsResponse))
}

func (o *ListLogsResponse_ErrorDetails) GotenObjectExt() {}

func (o *ListLogsResponse_ErrorDetails) MakeFullFieldMask() *ListLogsResponse_ErrorDetails_FieldMask {
	return FullListLogsResponse_ErrorDetails_FieldMask()
}

func (o *ListLogsResponse_ErrorDetails) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullListLogsResponse_ErrorDetails_FieldMask()
}

func (o *ListLogsResponse_ErrorDetails) MakeDiffFieldMask(other *ListLogsResponse_ErrorDetails) *ListLogsResponse_ErrorDetails_FieldMask {
	if o == nil && other == nil {
		return &ListLogsResponse_ErrorDetails_FieldMask{}
	}
	if o == nil || other == nil {
		return FullListLogsResponse_ErrorDetails_FieldMask()
	}

	res := &ListLogsResponse_ErrorDetails_FieldMask{}
	if o.GetRegionId() != other.GetRegionId() {
		res.Paths = append(res.Paths, &ListLogsResponseErrorDetails_FieldTerminalPath{selector: ListLogsResponseErrorDetails_FieldPathSelectorRegionId})
	}
	return res
}

func (o *ListLogsResponse_ErrorDetails) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ListLogsResponse_ErrorDetails))
}

func (o *ListLogsResponse_ErrorDetails) Clone() *ListLogsResponse_ErrorDetails {
	if o == nil {
		return nil
	}
	result := &ListLogsResponse_ErrorDetails{}
	result.RegionId = o.RegionId
	return result
}

func (o *ListLogsResponse_ErrorDetails) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ListLogsResponse_ErrorDetails) Merge(source *ListLogsResponse_ErrorDetails) {
	o.RegionId = source.GetRegionId()
}

func (o *ListLogsResponse_ErrorDetails) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ListLogsResponse_ErrorDetails))
}

func (o *CreateLogsRequest) GotenObjectExt() {}

func (o *CreateLogsRequest) MakeFullFieldMask() *CreateLogsRequest_FieldMask {
	return FullCreateLogsRequest_FieldMask()
}

func (o *CreateLogsRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullCreateLogsRequest_FieldMask()
}

func (o *CreateLogsRequest) MakeDiffFieldMask(other *CreateLogsRequest) *CreateLogsRequest_FieldMask {
	if o == nil && other == nil {
		return &CreateLogsRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullCreateLogsRequest_FieldMask()
	}

	res := &CreateLogsRequest_FieldMask{}
	if o.GetParent().String() != other.GetParent().String() {
		res.Paths = append(res.Paths, &CreateLogsRequest_FieldTerminalPath{selector: CreateLogsRequest_FieldPathSelectorParent})
	}

	if len(o.GetLogs()) == len(other.GetLogs()) {
		for i, lValue := range o.GetLogs() {
			rValue := other.GetLogs()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &CreateLogsRequest_FieldTerminalPath{selector: CreateLogsRequest_FieldPathSelectorLogs})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &CreateLogsRequest_FieldTerminalPath{selector: CreateLogsRequest_FieldPathSelectorLogs})
	}
	return res
}

func (o *CreateLogsRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*CreateLogsRequest))
}

func (o *CreateLogsRequest) Clone() *CreateLogsRequest {
	if o == nil {
		return nil
	}
	result := &CreateLogsRequest{}
	if o.Parent == nil {
		result.Parent = nil
	} else if data, err := o.Parent.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Parent = &log.ParentReference{}
		if err := result.Parent.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Logs = make([]*log.Log, len(o.Logs))
	for i, sourceValue := range o.Logs {
		result.Logs[i] = sourceValue.Clone()
	}
	return result
}

func (o *CreateLogsRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *CreateLogsRequest) Merge(source *CreateLogsRequest) {
	if source.GetParent() != nil {
		if data, err := source.GetParent().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Parent = &log.ParentReference{}
			if err := o.Parent.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Parent = nil
	}
	for _, sourceValue := range source.GetLogs() {
		exists := false
		for _, currentValue := range o.Logs {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *log.Log
			if sourceValue != nil {
				newDstElement = new(log.Log)
				newDstElement.Merge(sourceValue)
			}
			o.Logs = append(o.Logs, newDstElement)
		}
	}

}

func (o *CreateLogsRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*CreateLogsRequest))
}

func (o *CreateLogsResponse) GotenObjectExt() {}

func (o *CreateLogsResponse) MakeFullFieldMask() *CreateLogsResponse_FieldMask {
	return FullCreateLogsResponse_FieldMask()
}

func (o *CreateLogsResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullCreateLogsResponse_FieldMask()
}

func (o *CreateLogsResponse) MakeDiffFieldMask(other *CreateLogsResponse) *CreateLogsResponse_FieldMask {
	if o == nil && other == nil {
		return &CreateLogsResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullCreateLogsResponse_FieldMask()
	}

	res := &CreateLogsResponse_FieldMask{}

	if len(o.GetLogNames()) == len(other.GetLogNames()) {
		for i, lValue := range o.GetLogNames() {
			rValue := other.GetLogNames()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &CreateLogsResponse_FieldTerminalPath{selector: CreateLogsResponse_FieldPathSelectorLogNames})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &CreateLogsResponse_FieldTerminalPath{selector: CreateLogsResponse_FieldPathSelectorLogNames})
	}

	if len(o.GetFailedLogs()) == len(other.GetFailedLogs()) {
		for i, lValue := range o.GetFailedLogs() {
			rValue := other.GetFailedLogs()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &CreateLogsResponse_FieldTerminalPath{selector: CreateLogsResponse_FieldPathSelectorFailedLogs})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &CreateLogsResponse_FieldTerminalPath{selector: CreateLogsResponse_FieldPathSelectorFailedLogs})
	}
	return res
}

func (o *CreateLogsResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*CreateLogsResponse))
}

func (o *CreateLogsResponse) Clone() *CreateLogsResponse {
	if o == nil {
		return nil
	}
	result := &CreateLogsResponse{}
	result.LogNames = map[uint32]*log.Reference{}
	for key, sourceValue := range o.LogNames {
		if sourceValue == nil {
			result.LogNames[key] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.LogNames[key] = &log.Reference{}
			if err := result.LogNames[key].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	result.FailedLogs = make([]*CreateLogsResponse_CreateError, len(o.FailedLogs))
	for i, sourceValue := range o.FailedLogs {
		result.FailedLogs[i] = sourceValue.Clone()
	}
	return result
}

func (o *CreateLogsResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *CreateLogsResponse) Merge(source *CreateLogsResponse) {
	if source.GetLogNames() != nil {
		if o.LogNames == nil {
			o.LogNames = make(map[uint32]*log.Reference, len(source.GetLogNames()))
		}
		for key, sourceValue := range source.GetLogNames() {
			if sourceValue != nil {
				if data, err := sourceValue.ProtoString(); err != nil {
					panic(err)
				} else {
					o.LogNames[key] = &log.Reference{}
					if err := o.LogNames[key].ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			} else {
				o.LogNames[key] = nil
			}
		}
	}
	for _, sourceValue := range source.GetFailedLogs() {
		exists := false
		for _, currentValue := range o.FailedLogs {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *CreateLogsResponse_CreateError
			if sourceValue != nil {
				newDstElement = new(CreateLogsResponse_CreateError)
				newDstElement.Merge(sourceValue)
			}
			o.FailedLogs = append(o.FailedLogs, newDstElement)
		}
	}

}

func (o *CreateLogsResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*CreateLogsResponse))
}

func (o *CreateLogsResponse_CreateError) GotenObjectExt() {}

func (o *CreateLogsResponse_CreateError) MakeFullFieldMask() *CreateLogsResponse_CreateError_FieldMask {
	return FullCreateLogsResponse_CreateError_FieldMask()
}

func (o *CreateLogsResponse_CreateError) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullCreateLogsResponse_CreateError_FieldMask()
}

func (o *CreateLogsResponse_CreateError) MakeDiffFieldMask(other *CreateLogsResponse_CreateError) *CreateLogsResponse_CreateError_FieldMask {
	if o == nil && other == nil {
		return &CreateLogsResponse_CreateError_FieldMask{}
	}
	if o == nil || other == nil {
		return FullCreateLogsResponse_CreateError_FieldMask()
	}

	res := &CreateLogsResponse_CreateError_FieldMask{}

	if len(o.GetLogs()) == len(other.GetLogs()) {
		for i, lValue := range o.GetLogs() {
			rValue := other.GetLogs()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &CreateLogsResponseCreateError_FieldTerminalPath{selector: CreateLogsResponseCreateError_FieldPathSelectorLogs})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &CreateLogsResponseCreateError_FieldTerminalPath{selector: CreateLogsResponseCreateError_FieldPathSelectorLogs})
	}
	{
		subMask := o.GetStatus().MakeDiffFieldMask(other.GetStatus())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &CreateLogsResponseCreateError_FieldTerminalPath{selector: CreateLogsResponseCreateError_FieldPathSelectorStatus})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &CreateLogsResponseCreateError_FieldSubPath{selector: CreateLogsResponseCreateError_FieldPathSelectorStatus, subPath: subpath})
			}
		}
	}
	return res
}

func (o *CreateLogsResponse_CreateError) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*CreateLogsResponse_CreateError))
}

func (o *CreateLogsResponse_CreateError) Clone() *CreateLogsResponse_CreateError {
	if o == nil {
		return nil
	}
	result := &CreateLogsResponse_CreateError{}
	result.Logs = make([]*log.Log, len(o.Logs))
	for i, sourceValue := range o.Logs {
		result.Logs[i] = sourceValue.Clone()
	}
	result.Status = o.Status.Clone()
	return result
}

func (o *CreateLogsResponse_CreateError) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *CreateLogsResponse_CreateError) Merge(source *CreateLogsResponse_CreateError) {
	for _, sourceValue := range source.GetLogs() {
		exists := false
		for _, currentValue := range o.Logs {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *log.Log
			if sourceValue != nil {
				newDstElement = new(log.Log)
				newDstElement.Merge(sourceValue)
			}
			o.Logs = append(o.Logs, newDstElement)
		}
	}

	if source.GetStatus() != nil {
		if o.Status == nil {
			o.Status = new(rpc.Status)
		}
		o.Status.Merge(source.GetStatus())
	}
}

func (o *CreateLogsResponse_CreateError) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*CreateLogsResponse_CreateError))
}
