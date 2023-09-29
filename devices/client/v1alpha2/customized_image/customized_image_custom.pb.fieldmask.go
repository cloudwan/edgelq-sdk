// Code generated by protoc-gen-goten-object
// File: edgelq/devices/proto/v1alpha2/customized_image_custom.proto
// DO NOT EDIT!!!

package customized_image_client

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
	customized_image "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/customized_image"
)

// ensure the imports are used
var (
	_ = new(json.Marshaler)
	_ = strings.Builder{}

	_ = firestorepb.Value{}
	_ = codes.NotFound
	_ = status.Status{}
	_ = new(proto.Message)
	_ = new(preflect.Message)
	_ = fieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldMask)
)

// make sure we're using proto imports
var (
	_ = &customized_image.CustomizedImage{}
)

type RequestUrlRequest_FieldMask struct {
	Paths []RequestUrlRequest_FieldPath
}

func FullRequestUrlRequest_FieldMask() *RequestUrlRequest_FieldMask {
	res := &RequestUrlRequest_FieldMask{}
	res.Paths = append(res.Paths, &RequestUrlRequest_FieldTerminalPath{selector: RequestUrlRequest_FieldPathSelectorName})
	return res
}

func (fieldMask *RequestUrlRequest_FieldMask) String() string {
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
func (fieldMask *RequestUrlRequest_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *RequestUrlRequest_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseRequestUrlRequest_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *RequestUrlRequest_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 1)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*RequestUrlRequest_FieldTerminalPath); ok {
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

func (fieldMask *RequestUrlRequest_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseRequestUrlRequest_FieldPath(raw)
	})
}

func (fieldMask *RequestUrlRequest_FieldMask) ProtoMessage() {}

func (fieldMask *RequestUrlRequest_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *RequestUrlRequest_FieldMask) Subtract(other *RequestUrlRequest_FieldMask) *RequestUrlRequest_FieldMask {
	result := &RequestUrlRequest_FieldMask{}
	removedSelectors := make([]bool, 1)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *RequestUrlRequest_FieldTerminalPath:
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

func (fieldMask *RequestUrlRequest_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*RequestUrlRequest_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *RequestUrlRequest_FieldMask) FilterInputFields() *RequestUrlRequest_FieldMask {
	result := &RequestUrlRequest_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *RequestUrlRequest_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *RequestUrlRequest_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]RequestUrlRequest_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseRequestUrlRequest_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask RequestUrlRequest_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *RequestUrlRequest_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RequestUrlRequest_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask RequestUrlRequest_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *RequestUrlRequest_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RequestUrlRequest_FieldMask) AppendPath(path RequestUrlRequest_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *RequestUrlRequest_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(RequestUrlRequest_FieldPath))
}

func (fieldMask *RequestUrlRequest_FieldMask) GetPaths() []RequestUrlRequest_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *RequestUrlRequest_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *RequestUrlRequest_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseRequestUrlRequest_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *RequestUrlRequest_FieldMask) Set(target, source *RequestUrlRequest) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *RequestUrlRequest_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*RequestUrlRequest), source.(*RequestUrlRequest))
}

func (fieldMask *RequestUrlRequest_FieldMask) Project(source *RequestUrlRequest) *RequestUrlRequest {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &RequestUrlRequest{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *RequestUrlRequest_FieldTerminalPath:
			switch tp.selector {
			case RequestUrlRequest_FieldPathSelectorName:
				result.Name = source.Name
			}
		}
	}
	return result
}

func (fieldMask *RequestUrlRequest_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*RequestUrlRequest))
}

func (fieldMask *RequestUrlRequest_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type RequestUrlResponse_FieldMask struct {
	Paths []RequestUrlResponse_FieldPath
}

func FullRequestUrlResponse_FieldMask() *RequestUrlResponse_FieldMask {
	res := &RequestUrlResponse_FieldMask{}
	res.Paths = append(res.Paths, &RequestUrlResponse_FieldTerminalPath{selector: RequestUrlResponse_FieldPathSelectorUrl})
	res.Paths = append(res.Paths, &RequestUrlResponse_FieldTerminalPath{selector: RequestUrlResponse_FieldPathSelectorHeaders})
	return res
}

func (fieldMask *RequestUrlResponse_FieldMask) String() string {
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
func (fieldMask *RequestUrlResponse_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *RequestUrlResponse_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseRequestUrlResponse_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *RequestUrlResponse_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*RequestUrlResponse_FieldTerminalPath); ok {
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

func (fieldMask *RequestUrlResponse_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseRequestUrlResponse_FieldPath(raw)
	})
}

func (fieldMask *RequestUrlResponse_FieldMask) ProtoMessage() {}

func (fieldMask *RequestUrlResponse_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *RequestUrlResponse_FieldMask) Subtract(other *RequestUrlResponse_FieldMask) *RequestUrlResponse_FieldMask {
	result := &RequestUrlResponse_FieldMask{}
	removedSelectors := make([]bool, 2)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *RequestUrlResponse_FieldTerminalPath:
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

func (fieldMask *RequestUrlResponse_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*RequestUrlResponse_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *RequestUrlResponse_FieldMask) FilterInputFields() *RequestUrlResponse_FieldMask {
	result := &RequestUrlResponse_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *RequestUrlResponse_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *RequestUrlResponse_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]RequestUrlResponse_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseRequestUrlResponse_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask RequestUrlResponse_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *RequestUrlResponse_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RequestUrlResponse_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask RequestUrlResponse_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *RequestUrlResponse_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RequestUrlResponse_FieldMask) AppendPath(path RequestUrlResponse_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *RequestUrlResponse_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(RequestUrlResponse_FieldPath))
}

func (fieldMask *RequestUrlResponse_FieldMask) GetPaths() []RequestUrlResponse_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *RequestUrlResponse_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *RequestUrlResponse_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseRequestUrlResponse_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *RequestUrlResponse_FieldMask) Set(target, source *RequestUrlResponse) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *RequestUrlResponse_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*RequestUrlResponse), source.(*RequestUrlResponse))
}

func (fieldMask *RequestUrlResponse_FieldMask) Project(source *RequestUrlResponse) *RequestUrlResponse {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &RequestUrlResponse{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *RequestUrlResponse_FieldTerminalPath:
			switch tp.selector {
			case RequestUrlResponse_FieldPathSelectorUrl:
				result.Url = source.Url
			case RequestUrlResponse_FieldPathSelectorHeaders:
				result.Headers = source.Headers
			}
		}
	}
	return result
}

func (fieldMask *RequestUrlResponse_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*RequestUrlResponse))
}

func (fieldMask *RequestUrlResponse_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}