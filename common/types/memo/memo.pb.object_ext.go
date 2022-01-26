// Code generated by protoc-gen-goten-object
// File: edgelq/common/types/memo.proto
// DO NOT EDIT!!!

package ntt_memo

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

func (o *Memo) GotenObjectExt() {}

func (o *Memo) MakeFullFieldMask() *Memo_FieldMask {
	return FullMemo_FieldMask()
}

func (o *Memo) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullMemo_FieldMask()
}

func (o *Memo) MakeDiffFieldMask(other *Memo) *Memo_FieldMask {
	if o == nil && other == nil {
		return &Memo_FieldMask{}
	}
	if o == nil || other == nil {
		return FullMemo_FieldMask()
	}

	res := &Memo_FieldMask{}
	if !proto.Equal(o.GetCreateTime(), other.GetCreateTime()) {
		res.Paths = append(res.Paths, &Memo_FieldTerminalPath{selector: Memo_FieldPathSelectorCreateTime})
	}
	if !proto.Equal(o.GetUpdateTime(), other.GetUpdateTime()) {
		res.Paths = append(res.Paths, &Memo_FieldTerminalPath{selector: Memo_FieldPathSelectorUpdateTime})
	}
	if o.GetMessage() != other.GetMessage() {
		res.Paths = append(res.Paths, &Memo_FieldTerminalPath{selector: Memo_FieldPathSelectorMessage})
	}
	if o.GetCreatedBy() != other.GetCreatedBy() {
		res.Paths = append(res.Paths, &Memo_FieldTerminalPath{selector: Memo_FieldPathSelectorCreatedBy})
	}
	return res
}

func (o *Memo) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Memo))
}

func (o *Memo) Clone() *Memo {
	if o == nil {
		return nil
	}
	result := &Memo{}
	result.CreateTime = proto.Clone(o.CreateTime).(*timestamp.Timestamp)
	result.UpdateTime = proto.Clone(o.UpdateTime).(*timestamp.Timestamp)
	result.Message = o.Message
	result.CreatedBy = o.CreatedBy
	return result
}

func (o *Memo) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Memo) Merge(source *Memo) {
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
	o.Message = source.GetMessage()
	o.CreatedBy = source.GetCreatedBy()
}

func (o *Memo) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Memo))
}
