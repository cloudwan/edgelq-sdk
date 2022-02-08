// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha/project_custom.proto
// DO NOT EDIT!!!

package project_client

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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
	_ = &project.Project{}
	_ = &field_mask.FieldMask{}
)

func (o *ListMyProjectsRequest) GotenObjectExt() {}

func (o *ListMyProjectsRequest) MakeFullFieldMask() *ListMyProjectsRequest_FieldMask {
	return FullListMyProjectsRequest_FieldMask()
}

func (o *ListMyProjectsRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullListMyProjectsRequest_FieldMask()
}

func (o *ListMyProjectsRequest) MakeDiffFieldMask(other *ListMyProjectsRequest) *ListMyProjectsRequest_FieldMask {
	if o == nil && other == nil {
		return &ListMyProjectsRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullListMyProjectsRequest_FieldMask()
	}

	res := &ListMyProjectsRequest_FieldMask{}
	if o.GetFilter().String() != other.GetFilter().String() {
		res.Paths = append(res.Paths, &ListMyProjectsRequest_FieldTerminalPath{selector: ListMyProjectsRequest_FieldPathSelectorFilter})
	}
	if !proto.Equal(o.GetFieldMask(), other.GetFieldMask()) {
		res.Paths = append(res.Paths, &ListMyProjectsRequest_FieldTerminalPath{selector: ListMyProjectsRequest_FieldPathSelectorFieldMask})
	}
	return res
}

func (o *ListMyProjectsRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ListMyProjectsRequest))
}

func (o *ListMyProjectsRequest) Clone() *ListMyProjectsRequest {
	if o == nil {
		return nil
	}
	result := &ListMyProjectsRequest{}
	if o.Filter == nil {
		result.Filter = nil
	} else if data, err := o.Filter.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Filter = &project.Filter{}
		if err := result.Filter.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.FieldMask = proto.Clone(o.FieldMask).(*project.Project_FieldMask)
	return result
}

func (o *ListMyProjectsRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ListMyProjectsRequest) Merge(source *ListMyProjectsRequest) {
	if source.GetFilter() != nil {
		if data, err := source.GetFilter().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Filter = &project.Filter{}
			if err := o.Filter.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Filter = nil
	}
	if source.GetFieldMask() != nil {
		if o.FieldMask == nil {
			o.FieldMask = new(project.Project_FieldMask)
		}
		mergedMask := fieldmaskpb.Union(source.GetFieldMask().ToProtoFieldMask(), o.FieldMask.ToProtoFieldMask())
		if err := o.FieldMask.FromProtoFieldMask(mergedMask); err != nil {
			panic(err)
		}
	}
}

func (o *ListMyProjectsRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ListMyProjectsRequest))
}

func (o *ListMyProjectsResponse) GotenObjectExt() {}

func (o *ListMyProjectsResponse) MakeFullFieldMask() *ListMyProjectsResponse_FieldMask {
	return FullListMyProjectsResponse_FieldMask()
}

func (o *ListMyProjectsResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullListMyProjectsResponse_FieldMask()
}

func (o *ListMyProjectsResponse) MakeDiffFieldMask(other *ListMyProjectsResponse) *ListMyProjectsResponse_FieldMask {
	if o == nil && other == nil {
		return &ListMyProjectsResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullListMyProjectsResponse_FieldMask()
	}

	res := &ListMyProjectsResponse_FieldMask{}

	if len(o.GetProjects()) == len(other.GetProjects()) {
		for i, lValue := range o.GetProjects() {
			rValue := other.GetProjects()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &ListMyProjectsResponse_FieldTerminalPath{selector: ListMyProjectsResponse_FieldPathSelectorProjects})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ListMyProjectsResponse_FieldTerminalPath{selector: ListMyProjectsResponse_FieldPathSelectorProjects})
	}
	return res
}

func (o *ListMyProjectsResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ListMyProjectsResponse))
}

func (o *ListMyProjectsResponse) Clone() *ListMyProjectsResponse {
	if o == nil {
		return nil
	}
	result := &ListMyProjectsResponse{}
	result.Projects = make([]*project.Project, len(o.Projects))
	for i, sourceValue := range o.Projects {
		result.Projects[i] = sourceValue.Clone()
	}
	return result
}

func (o *ListMyProjectsResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ListMyProjectsResponse) Merge(source *ListMyProjectsResponse) {
	for _, sourceValue := range source.GetProjects() {
		exists := false
		for _, currentValue := range o.Projects {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *project.Project
			if sourceValue != nil {
				newDstElement = new(project.Project)
				newDstElement.Merge(sourceValue)
			}
			o.Projects = append(o.Projects, newDstElement)
		}
	}

}

func (o *ListMyProjectsResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ListMyProjectsResponse))
}