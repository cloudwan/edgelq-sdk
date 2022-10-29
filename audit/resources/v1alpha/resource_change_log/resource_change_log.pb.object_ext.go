// Code generated by protoc-gen-goten-object
// File: edgelq/audit/proto/v1alpha/resource_change_log.proto
// DO NOT EDIT!!!

package resource_change_log

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	audit_common "github.com/cloudwan/edgelq-sdk/audit/common/v1alpha"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
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
	_ = &audit_common.Authentication{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &timestamp.Timestamp{}
)

func (o *ResourceChangeLog) GotenObjectExt() {}

func (o *ResourceChangeLog) MakeFullFieldMask() *ResourceChangeLog_FieldMask {
	return FullResourceChangeLog_FieldMask()
}

func (o *ResourceChangeLog) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullResourceChangeLog_FieldMask()
}

func (o *ResourceChangeLog) MakeDiffFieldMask(other *ResourceChangeLog) *ResourceChangeLog_FieldMask {
	if o == nil && other == nil {
		return &ResourceChangeLog_FieldMask{}
	}
	if o == nil || other == nil {
		return FullResourceChangeLog_FieldMask()
	}

	res := &ResourceChangeLog_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &ResourceChangeLog_FieldTerminalPath{selector: ResourceChangeLog_FieldPathSelectorName})
	}
	if o.GetScope() != other.GetScope() {
		res.Paths = append(res.Paths, &ResourceChangeLog_FieldTerminalPath{selector: ResourceChangeLog_FieldPathSelectorScope})
	}
	if o.GetRequestId() != other.GetRequestId() {
		res.Paths = append(res.Paths, &ResourceChangeLog_FieldTerminalPath{selector: ResourceChangeLog_FieldPathSelectorRequestId})
	}
	if !proto.Equal(o.GetTimestamp(), other.GetTimestamp()) {
		res.Paths = append(res.Paths, &ResourceChangeLog_FieldTerminalPath{selector: ResourceChangeLog_FieldPathSelectorTimestamp})
	}
	{
		subMask := o.GetAuthentication().MakeDiffFieldMask(other.GetAuthentication())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ResourceChangeLog_FieldTerminalPath{selector: ResourceChangeLog_FieldPathSelectorAuthentication})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ResourceChangeLog_FieldSubPath{selector: ResourceChangeLog_FieldPathSelectorAuthentication, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetService().MakeDiffFieldMask(other.GetService())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ResourceChangeLog_FieldTerminalPath{selector: ResourceChangeLog_FieldPathSelectorService})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ResourceChangeLog_FieldSubPath{selector: ResourceChangeLog_FieldPathSelectorService, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetResource().MakeDiffFieldMask(other.GetResource())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ResourceChangeLog_FieldTerminalPath{selector: ResourceChangeLog_FieldPathSelectorResource})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ResourceChangeLog_FieldSubPath{selector: ResourceChangeLog_FieldPathSelectorResource, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetTransaction().MakeDiffFieldMask(other.GetTransaction())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ResourceChangeLog_FieldTerminalPath{selector: ResourceChangeLog_FieldPathSelectorTransaction})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ResourceChangeLog_FieldSubPath{selector: ResourceChangeLog_FieldPathSelectorTransaction, subPath: subpath})
			}
		}
	}
	return res
}

func (o *ResourceChangeLog) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ResourceChangeLog))
}

func (o *ResourceChangeLog) Clone() *ResourceChangeLog {
	if o == nil {
		return nil
	}
	result := &ResourceChangeLog{}
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
	result.RequestId = o.RequestId
	result.Timestamp = proto.Clone(o.Timestamp).(*timestamp.Timestamp)
	result.Authentication = o.Authentication.Clone()
	result.Service = o.Service.Clone()
	result.Resource = o.Resource.Clone()
	result.Transaction = o.Transaction.Clone()
	return result
}

func (o *ResourceChangeLog) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ResourceChangeLog) Merge(source *ResourceChangeLog) {
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
	o.RequestId = source.GetRequestId()
	if source.GetTimestamp() != nil {
		if o.Timestamp == nil {
			o.Timestamp = new(timestamp.Timestamp)
		}
		proto.Merge(o.Timestamp, source.GetTimestamp())
	}
	if source.GetAuthentication() != nil {
		if o.Authentication == nil {
			o.Authentication = new(audit_common.Authentication)
		}
		o.Authentication.Merge(source.GetAuthentication())
	}
	if source.GetService() != nil {
		if o.Service == nil {
			o.Service = new(audit_common.ServiceData)
		}
		o.Service.Merge(source.GetService())
	}
	if source.GetResource() != nil {
		if o.Resource == nil {
			o.Resource = new(ResourceChangeLog_ResourceChange)
		}
		o.Resource.Merge(source.GetResource())
	}
	if source.GetTransaction() != nil {
		if o.Transaction == nil {
			o.Transaction = new(ResourceChangeLog_TransactionInfo)
		}
		o.Transaction.Merge(source.GetTransaction())
	}
}

func (o *ResourceChangeLog) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ResourceChangeLog))
}

func (o *ResourceChangeLog_ResourceChange) GotenObjectExt() {}

func (o *ResourceChangeLog_ResourceChange) MakeFullFieldMask() *ResourceChangeLog_ResourceChange_FieldMask {
	return FullResourceChangeLog_ResourceChange_FieldMask()
}

func (o *ResourceChangeLog_ResourceChange) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullResourceChangeLog_ResourceChange_FieldMask()
}

func (o *ResourceChangeLog_ResourceChange) MakeDiffFieldMask(other *ResourceChangeLog_ResourceChange) *ResourceChangeLog_ResourceChange_FieldMask {
	if o == nil && other == nil {
		return &ResourceChangeLog_ResourceChange_FieldMask{}
	}
	if o == nil || other == nil {
		return FullResourceChangeLog_ResourceChange_FieldMask()
	}

	res := &ResourceChangeLog_ResourceChange_FieldMask{}
	if o.GetName() != other.GetName() {
		res.Paths = append(res.Paths, &ResourceChangeLogResourceChange_FieldTerminalPath{selector: ResourceChangeLogResourceChange_FieldPathSelectorName})
	}
	if o.GetType() != other.GetType() {
		res.Paths = append(res.Paths, &ResourceChangeLogResourceChange_FieldTerminalPath{selector: ResourceChangeLogResourceChange_FieldPathSelectorType})
	}
	if o.GetAction() != other.GetAction() {
		res.Paths = append(res.Paths, &ResourceChangeLogResourceChange_FieldTerminalPath{selector: ResourceChangeLogResourceChange_FieldPathSelectorAction})
	}
	{
		subMask := o.GetPre().MakeDiffFieldMask(other.GetPre())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ResourceChangeLogResourceChange_FieldTerminalPath{selector: ResourceChangeLogResourceChange_FieldPathSelectorPre})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ResourceChangeLogResourceChange_FieldSubPath{selector: ResourceChangeLogResourceChange_FieldPathSelectorPre, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetPost().MakeDiffFieldMask(other.GetPost())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ResourceChangeLogResourceChange_FieldTerminalPath{selector: ResourceChangeLogResourceChange_FieldPathSelectorPost})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ResourceChangeLogResourceChange_FieldSubPath{selector: ResourceChangeLogResourceChange_FieldPathSelectorPost, subPath: subpath})
			}
		}
	}
	return res
}

func (o *ResourceChangeLog_ResourceChange) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ResourceChangeLog_ResourceChange))
}

func (o *ResourceChangeLog_ResourceChange) Clone() *ResourceChangeLog_ResourceChange {
	if o == nil {
		return nil
	}
	result := &ResourceChangeLog_ResourceChange{}
	result.Name = o.Name
	result.Type = o.Type
	result.Action = o.Action
	result.Pre = o.Pre.Clone()
	result.Post = o.Post.Clone()
	return result
}

func (o *ResourceChangeLog_ResourceChange) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ResourceChangeLog_ResourceChange) Merge(source *ResourceChangeLog_ResourceChange) {
	o.Name = source.GetName()
	o.Type = source.GetType()
	o.Action = source.GetAction()
	if source.GetPre() != nil {
		if o.Pre == nil {
			o.Pre = new(audit_common.ObjectState)
		}
		o.Pre.Merge(source.GetPre())
	}
	if source.GetPost() != nil {
		if o.Post == nil {
			o.Post = new(audit_common.ObjectState)
		}
		o.Post.Merge(source.GetPost())
	}
}

func (o *ResourceChangeLog_ResourceChange) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ResourceChangeLog_ResourceChange))
}

func (o *ResourceChangeLog_TransactionInfo) GotenObjectExt() {}

func (o *ResourceChangeLog_TransactionInfo) MakeFullFieldMask() *ResourceChangeLog_TransactionInfo_FieldMask {
	return FullResourceChangeLog_TransactionInfo_FieldMask()
}

func (o *ResourceChangeLog_TransactionInfo) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullResourceChangeLog_TransactionInfo_FieldMask()
}

func (o *ResourceChangeLog_TransactionInfo) MakeDiffFieldMask(other *ResourceChangeLog_TransactionInfo) *ResourceChangeLog_TransactionInfo_FieldMask {
	if o == nil && other == nil {
		return &ResourceChangeLog_TransactionInfo_FieldMask{}
	}
	if o == nil || other == nil {
		return FullResourceChangeLog_TransactionInfo_FieldMask()
	}

	res := &ResourceChangeLog_TransactionInfo_FieldMask{}
	if o.GetIdentifier() != other.GetIdentifier() {
		res.Paths = append(res.Paths, &ResourceChangeLogTransactionInfo_FieldTerminalPath{selector: ResourceChangeLogTransactionInfo_FieldPathSelectorIdentifier})
	}
	if o.GetTryCounter() != other.GetTryCounter() {
		res.Paths = append(res.Paths, &ResourceChangeLogTransactionInfo_FieldTerminalPath{selector: ResourceChangeLogTransactionInfo_FieldPathSelectorTryCounter})
	}
	if o.GetState() != other.GetState() {
		res.Paths = append(res.Paths, &ResourceChangeLogTransactionInfo_FieldTerminalPath{selector: ResourceChangeLogTransactionInfo_FieldPathSelectorState})
	}
	return res
}

func (o *ResourceChangeLog_TransactionInfo) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ResourceChangeLog_TransactionInfo))
}

func (o *ResourceChangeLog_TransactionInfo) Clone() *ResourceChangeLog_TransactionInfo {
	if o == nil {
		return nil
	}
	result := &ResourceChangeLog_TransactionInfo{}
	result.Identifier = o.Identifier
	result.TryCounter = o.TryCounter
	result.State = o.State
	return result
}

func (o *ResourceChangeLog_TransactionInfo) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ResourceChangeLog_TransactionInfo) Merge(source *ResourceChangeLog_TransactionInfo) {
	o.Identifier = source.GetIdentifier()
	o.TryCounter = source.GetTryCounter()
	o.State = source.GetState()
}

func (o *ResourceChangeLog_TransactionInfo) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ResourceChangeLog_TransactionInfo))
}
