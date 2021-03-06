// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha2/project_custom.proto
// DO NOT EDIT!!!

package project_client

import (
	"encoding/json"
	"strings"

	firestorepb "google.golang.org/genproto/googleapis/firestore/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// ensure the imports are used
var (
	_ = json.Marshaler(nil)
	_ = strings.Builder{}

	_ = firestorepb.Value{}
	_ = codes.NotFound
	_ = status.Status{}
	_ = proto.Message(nil)
	_ = preflect.Message(nil)
	_ = fieldmaskpb.FieldMask{}

	_ = gotenobject.FieldMask(nil)
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &field_mask.FieldMask{}
)

type ListMyProjectsRequest_FieldMask struct {
	Paths []ListMyProjectsRequest_FieldPath
}

func FullListMyProjectsRequest_FieldMask() *ListMyProjectsRequest_FieldMask {
	res := &ListMyProjectsRequest_FieldMask{}
	res.Paths = append(res.Paths, &ListMyProjectsRequest_FieldTerminalPath{selector: ListMyProjectsRequest_FieldPathSelectorFilter})
	res.Paths = append(res.Paths, &ListMyProjectsRequest_FieldTerminalPath{selector: ListMyProjectsRequest_FieldPathSelectorFieldMask})
	return res
}

func (fieldMask *ListMyProjectsRequest_FieldMask) String() string {
	if fieldMask == nil {
		return "<nil>"
	}
	pathsStr := make([]string, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		pathsStr = append(pathsStr, path.String())
	}
	return strings.Join(pathsStr, ", ")
}

// firestore encoding/decoding integration
func (fieldMask *ListMyProjectsRequest_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
	if fieldMask == nil {
		return &firestorepb.Value{ValueType: &firestorepb.Value_NullValue{}}, nil
	}
	arrayValues := make([]*firestorepb.Value, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.GetPaths() {
		arrayValues = append(arrayValues, &firestorepb.Value{ValueType: &firestorepb.Value_StringValue{StringValue: path.String()}})
	}
	return &firestorepb.Value{
		ValueType: &firestorepb.Value_ArrayValue{ArrayValue: &firestorepb.ArrayValue{Values: arrayValues}},
	}, nil
}

func (fieldMask *ListMyProjectsRequest_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseListMyProjectsRequest_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ListMyProjectsRequest_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ListMyProjectsRequest_FieldTerminalPath); ok {
			presentSelectors[int(asFinal.selector)] = true
		}
	}
	for _, flag := range presentSelectors {
		if !flag {
			return false
		}
	}
	return true
}

func (fieldMask *ListMyProjectsRequest_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseListMyProjectsRequest_FieldPath(raw)
	})
}

func (fieldMask *ListMyProjectsRequest_FieldMask) ProtoMessage() {}

func (fieldMask *ListMyProjectsRequest_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ListMyProjectsRequest_FieldMask) Subtract(other *ListMyProjectsRequest_FieldMask) *ListMyProjectsRequest_FieldMask {
	result := &ListMyProjectsRequest_FieldMask{}
	removedSelectors := make([]bool, 2)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ListMyProjectsRequest_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			result.Paths = append(result.Paths, path)
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *ListMyProjectsRequest_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ListMyProjectsRequest_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ListMyProjectsRequest_FieldMask) FilterInputFields() *ListMyProjectsRequest_FieldMask {
	result := &ListMyProjectsRequest_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ListMyProjectsRequest_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ListMyProjectsRequest_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ListMyProjectsRequest_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseListMyProjectsRequest_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ListMyProjectsRequest_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ListMyProjectsRequest_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ListMyProjectsRequest_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ListMyProjectsRequest_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ListMyProjectsRequest_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ListMyProjectsRequest_FieldMask) AppendPath(path ListMyProjectsRequest_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ListMyProjectsRequest_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ListMyProjectsRequest_FieldPath))
}

func (fieldMask *ListMyProjectsRequest_FieldMask) GetPaths() []ListMyProjectsRequest_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ListMyProjectsRequest_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ListMyProjectsRequest_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseListMyProjectsRequest_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ListMyProjectsRequest_FieldMask) Set(target, source *ListMyProjectsRequest) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ListMyProjectsRequest_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ListMyProjectsRequest), source.(*ListMyProjectsRequest))
}

func (fieldMask *ListMyProjectsRequest_FieldMask) Project(source *ListMyProjectsRequest) *ListMyProjectsRequest {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ListMyProjectsRequest{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ListMyProjectsRequest_FieldTerminalPath:
			switch tp.selector {
			case ListMyProjectsRequest_FieldPathSelectorFilter:
				result.Filter = source.Filter
			case ListMyProjectsRequest_FieldPathSelectorFieldMask:
				result.FieldMask = source.FieldMask
			}
		}
	}
	return result
}

func (fieldMask *ListMyProjectsRequest_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ListMyProjectsRequest))
}

func (fieldMask *ListMyProjectsRequest_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type ListMyProjectsResponse_FieldMask struct {
	Paths []ListMyProjectsResponse_FieldPath
}

func FullListMyProjectsResponse_FieldMask() *ListMyProjectsResponse_FieldMask {
	res := &ListMyProjectsResponse_FieldMask{}
	res.Paths = append(res.Paths, &ListMyProjectsResponse_FieldTerminalPath{selector: ListMyProjectsResponse_FieldPathSelectorProjects})
	return res
}

func (fieldMask *ListMyProjectsResponse_FieldMask) String() string {
	if fieldMask == nil {
		return "<nil>"
	}
	pathsStr := make([]string, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		pathsStr = append(pathsStr, path.String())
	}
	return strings.Join(pathsStr, ", ")
}

// firestore encoding/decoding integration
func (fieldMask *ListMyProjectsResponse_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
	if fieldMask == nil {
		return &firestorepb.Value{ValueType: &firestorepb.Value_NullValue{}}, nil
	}
	arrayValues := make([]*firestorepb.Value, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.GetPaths() {
		arrayValues = append(arrayValues, &firestorepb.Value{ValueType: &firestorepb.Value_StringValue{StringValue: path.String()}})
	}
	return &firestorepb.Value{
		ValueType: &firestorepb.Value_ArrayValue{ArrayValue: &firestorepb.ArrayValue{Values: arrayValues}},
	}, nil
}

func (fieldMask *ListMyProjectsResponse_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseListMyProjectsResponse_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ListMyProjectsResponse_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 1)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ListMyProjectsResponse_FieldTerminalPath); ok {
			presentSelectors[int(asFinal.selector)] = true
		}
	}
	for _, flag := range presentSelectors {
		if !flag {
			return false
		}
	}
	return true
}

func (fieldMask *ListMyProjectsResponse_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseListMyProjectsResponse_FieldPath(raw)
	})
}

func (fieldMask *ListMyProjectsResponse_FieldMask) ProtoMessage() {}

func (fieldMask *ListMyProjectsResponse_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ListMyProjectsResponse_FieldMask) Subtract(other *ListMyProjectsResponse_FieldMask) *ListMyProjectsResponse_FieldMask {
	result := &ListMyProjectsResponse_FieldMask{}
	removedSelectors := make([]bool, 1)
	otherSubMasks := map[ListMyProjectsResponse_FieldPathSelector]gotenobject.FieldMask{
		ListMyProjectsResponse_FieldPathSelectorProjects: &project.Project_FieldMask{},
	}
	mySubMasks := map[ListMyProjectsResponse_FieldPathSelector]gotenobject.FieldMask{
		ListMyProjectsResponse_FieldPathSelectorProjects: &project.Project_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ListMyProjectsResponse_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *ListMyProjectsResponse_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*ListMyProjectsResponse_FieldTerminalPath); ok {
					switch tp.selector {
					case ListMyProjectsResponse_FieldPathSelectorProjects:
						mySubMasks[ListMyProjectsResponse_FieldPathSelectorProjects] = project.FullProject_FieldMask()
					}
				} else if tp, ok := path.(*ListMyProjectsResponse_FieldSubPath); ok {
					mySubMasks[tp.selector].AppendRawPath(tp.subPath)
				}
			} else {
				result.Paths = append(result.Paths, path)
			}
		}
	}
	for selector, mySubMask := range mySubMasks {
		if mySubMask.PathsCount() > 0 {
			for _, allowedPath := range mySubMask.SubtractRaw(otherSubMasks[selector]).GetRawPaths() {
				result.Paths = append(result.Paths, &ListMyProjectsResponse_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *ListMyProjectsResponse_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ListMyProjectsResponse_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ListMyProjectsResponse_FieldMask) FilterInputFields() *ListMyProjectsResponse_FieldMask {
	result := &ListMyProjectsResponse_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case ListMyProjectsResponse_FieldPathSelectorProjects:
			if _, ok := path.(*ListMyProjectsResponse_FieldTerminalPath); ok {
				for _, subpath := range project.FullProject_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ListMyProjectsResponse_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*ListMyProjectsResponse_FieldSubPath); ok {
				selectedMask := &project.Project_FieldMask{
					Paths: []project.Project_FieldPath{sub.subPath.(project.Project_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ListMyProjectsResponse_FieldSubPath{selector: ListMyProjectsResponse_FieldPathSelectorProjects, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ListMyProjectsResponse_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ListMyProjectsResponse_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ListMyProjectsResponse_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseListMyProjectsResponse_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ListMyProjectsResponse_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ListMyProjectsResponse_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ListMyProjectsResponse_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ListMyProjectsResponse_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ListMyProjectsResponse_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ListMyProjectsResponse_FieldMask) AppendPath(path ListMyProjectsResponse_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ListMyProjectsResponse_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ListMyProjectsResponse_FieldPath))
}

func (fieldMask *ListMyProjectsResponse_FieldMask) GetPaths() []ListMyProjectsResponse_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ListMyProjectsResponse_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ListMyProjectsResponse_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseListMyProjectsResponse_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ListMyProjectsResponse_FieldMask) Set(target, source *ListMyProjectsResponse) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ListMyProjectsResponse_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ListMyProjectsResponse), source.(*ListMyProjectsResponse))
}

func (fieldMask *ListMyProjectsResponse_FieldMask) Project(source *ListMyProjectsResponse) *ListMyProjectsResponse {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ListMyProjectsResponse{}
	projectsMask := &project.Project_FieldMask{}
	wholeProjectsAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ListMyProjectsResponse_FieldTerminalPath:
			switch tp.selector {
			case ListMyProjectsResponse_FieldPathSelectorProjects:
				result.Projects = source.Projects
				wholeProjectsAccepted = true
			}
		case *ListMyProjectsResponse_FieldSubPath:
			switch tp.selector {
			case ListMyProjectsResponse_FieldPathSelectorProjects:
				projectsMask.AppendPath(tp.subPath.(project.Project_FieldPath))
			}
		}
	}
	if wholeProjectsAccepted == false && len(projectsMask.Paths) > 0 {
		for _, sourceItem := range source.GetProjects() {
			result.Projects = append(result.Projects, projectsMask.Project(sourceItem))
		}
	}
	return result
}

func (fieldMask *ListMyProjectsResponse_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ListMyProjectsResponse))
}

func (fieldMask *ListMyProjectsResponse_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
