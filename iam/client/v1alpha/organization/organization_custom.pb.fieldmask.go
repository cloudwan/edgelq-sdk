// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha/organization_custom.proto
// DO NOT EDIT!!!

package organization_client

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
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
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
	_ = &organization.Organization{}
	_ = &field_mask.FieldMask{}
)

type ListMyOrganizationsRequest_FieldMask struct {
	Paths []ListMyOrganizationsRequest_FieldPath
}

func FullListMyOrganizationsRequest_FieldMask() *ListMyOrganizationsRequest_FieldMask {
	res := &ListMyOrganizationsRequest_FieldMask{}
	res.Paths = append(res.Paths, &ListMyOrganizationsRequest_FieldTerminalPath{selector: ListMyOrganizationsRequest_FieldPathSelectorFilter})
	res.Paths = append(res.Paths, &ListMyOrganizationsRequest_FieldTerminalPath{selector: ListMyOrganizationsRequest_FieldPathSelectorFieldMask})
	return res
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) String() string {
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
func (fieldMask *ListMyOrganizationsRequest_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ListMyOrganizationsRequest_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseListMyOrganizationsRequest_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ListMyOrganizationsRequest_FieldTerminalPath); ok {
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

func (fieldMask *ListMyOrganizationsRequest_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseListMyOrganizationsRequest_FieldPath(raw)
	})
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) ProtoMessage() {}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) Subtract(other *ListMyOrganizationsRequest_FieldMask) *ListMyOrganizationsRequest_FieldMask {
	result := &ListMyOrganizationsRequest_FieldMask{}
	removedSelectors := make([]bool, 2)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ListMyOrganizationsRequest_FieldTerminalPath:
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

func (fieldMask *ListMyOrganizationsRequest_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ListMyOrganizationsRequest_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ListMyOrganizationsRequest_FieldMask) FilterInputFields() *ListMyOrganizationsRequest_FieldMask {
	result := &ListMyOrganizationsRequest_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ListMyOrganizationsRequest_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ListMyOrganizationsRequest_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseListMyOrganizationsRequest_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ListMyOrganizationsRequest_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ListMyOrganizationsRequest_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) AppendPath(path ListMyOrganizationsRequest_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ListMyOrganizationsRequest_FieldPath))
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) GetPaths() []ListMyOrganizationsRequest_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseListMyOrganizationsRequest_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) Set(target, source *ListMyOrganizationsRequest) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ListMyOrganizationsRequest), source.(*ListMyOrganizationsRequest))
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) Project(source *ListMyOrganizationsRequest) *ListMyOrganizationsRequest {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ListMyOrganizationsRequest{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ListMyOrganizationsRequest_FieldTerminalPath:
			switch tp.selector {
			case ListMyOrganizationsRequest_FieldPathSelectorFilter:
				result.Filter = source.Filter
			case ListMyOrganizationsRequest_FieldPathSelectorFieldMask:
				result.FieldMask = source.FieldMask
			}
		}
	}
	return result
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ListMyOrganizationsRequest))
}

func (fieldMask *ListMyOrganizationsRequest_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type ListMyOrganizationsResponse_FieldMask struct {
	Paths []ListMyOrganizationsResponse_FieldPath
}

func FullListMyOrganizationsResponse_FieldMask() *ListMyOrganizationsResponse_FieldMask {
	res := &ListMyOrganizationsResponse_FieldMask{}
	res.Paths = append(res.Paths, &ListMyOrganizationsResponse_FieldTerminalPath{selector: ListMyOrganizationsResponse_FieldPathSelectorOrganizations})
	return res
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) String() string {
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
func (fieldMask *ListMyOrganizationsResponse_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ListMyOrganizationsResponse_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseListMyOrganizationsResponse_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 1)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ListMyOrganizationsResponse_FieldTerminalPath); ok {
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

func (fieldMask *ListMyOrganizationsResponse_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseListMyOrganizationsResponse_FieldPath(raw)
	})
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) ProtoMessage() {}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) Subtract(other *ListMyOrganizationsResponse_FieldMask) *ListMyOrganizationsResponse_FieldMask {
	result := &ListMyOrganizationsResponse_FieldMask{}
	removedSelectors := make([]bool, 1)
	otherSubMasks := map[ListMyOrganizationsResponse_FieldPathSelector]gotenobject.FieldMask{
		ListMyOrganizationsResponse_FieldPathSelectorOrganizations: &organization.Organization_FieldMask{},
	}
	mySubMasks := map[ListMyOrganizationsResponse_FieldPathSelector]gotenobject.FieldMask{
		ListMyOrganizationsResponse_FieldPathSelectorOrganizations: &organization.Organization_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ListMyOrganizationsResponse_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *ListMyOrganizationsResponse_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*ListMyOrganizationsResponse_FieldTerminalPath); ok {
					switch tp.selector {
					case ListMyOrganizationsResponse_FieldPathSelectorOrganizations:
						mySubMasks[ListMyOrganizationsResponse_FieldPathSelectorOrganizations] = organization.FullOrganization_FieldMask()
					}
				} else if tp, ok := path.(*ListMyOrganizationsResponse_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &ListMyOrganizationsResponse_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ListMyOrganizationsResponse_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ListMyOrganizationsResponse_FieldMask) FilterInputFields() *ListMyOrganizationsResponse_FieldMask {
	result := &ListMyOrganizationsResponse_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case ListMyOrganizationsResponse_FieldPathSelectorOrganizations:
			if _, ok := path.(*ListMyOrganizationsResponse_FieldTerminalPath); ok {
				for _, subpath := range organization.FullOrganization_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ListMyOrganizationsResponse_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*ListMyOrganizationsResponse_FieldSubPath); ok {
				selectedMask := &organization.Organization_FieldMask{
					Paths: []organization.Organization_FieldPath{sub.subPath.(organization.Organization_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ListMyOrganizationsResponse_FieldSubPath{selector: ListMyOrganizationsResponse_FieldPathSelectorOrganizations, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ListMyOrganizationsResponse_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ListMyOrganizationsResponse_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseListMyOrganizationsResponse_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ListMyOrganizationsResponse_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ListMyOrganizationsResponse_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) AppendPath(path ListMyOrganizationsResponse_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ListMyOrganizationsResponse_FieldPath))
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) GetPaths() []ListMyOrganizationsResponse_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseListMyOrganizationsResponse_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) Set(target, source *ListMyOrganizationsResponse) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ListMyOrganizationsResponse), source.(*ListMyOrganizationsResponse))
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) Project(source *ListMyOrganizationsResponse) *ListMyOrganizationsResponse {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ListMyOrganizationsResponse{}
	organizationsMask := &organization.Organization_FieldMask{}
	wholeOrganizationsAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ListMyOrganizationsResponse_FieldTerminalPath:
			switch tp.selector {
			case ListMyOrganizationsResponse_FieldPathSelectorOrganizations:
				result.Organizations = source.Organizations
				wholeOrganizationsAccepted = true
			}
		case *ListMyOrganizationsResponse_FieldSubPath:
			switch tp.selector {
			case ListMyOrganizationsResponse_FieldPathSelectorOrganizations:
				organizationsMask.AppendPath(tp.subPath.(organization.Organization_FieldPath))
			}
		}
	}
	if wholeOrganizationsAccepted == false && len(organizationsMask.Paths) > 0 {
		for _, sourceItem := range source.GetOrganizations() {
			result.Organizations = append(result.Organizations, organizationsMask.Project(sourceItem))
		}
	}
	return result
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ListMyOrganizationsResponse))
}

func (fieldMask *ListMyOrganizationsResponse_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
