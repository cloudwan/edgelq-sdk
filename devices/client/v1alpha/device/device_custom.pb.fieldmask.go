// Code generated by protoc-gen-goten-object
// File: edgelq/devices/proto/v1alpha/device_custom.proto
// DO NOT EDIT!!!

package device_client

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
	api "github.com/cloudwan/edgelq-sdk/common/api"
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha/device"
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
	_ = &api.Account{}
	_ = &device.Device{}
)

type ProvisionServiceAccountToDeviceRequest_FieldMask struct {
	Paths []ProvisionServiceAccountToDeviceRequest_FieldPath
}

func FullProvisionServiceAccountToDeviceRequest_FieldMask() *ProvisionServiceAccountToDeviceRequest_FieldMask {
	res := &ProvisionServiceAccountToDeviceRequest_FieldMask{}
	res.Paths = append(res.Paths, &ProvisionServiceAccountToDeviceRequest_FieldTerminalPath{selector: ProvisionServiceAccountToDeviceRequest_FieldPathSelectorName})
	return res
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) String() string {
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
func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseProvisionServiceAccountToDeviceRequest_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 1)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ProvisionServiceAccountToDeviceRequest_FieldTerminalPath); ok {
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

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseProvisionServiceAccountToDeviceRequest_FieldPath(raw)
	})
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) ProtoMessage() {}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) Subtract(other *ProvisionServiceAccountToDeviceRequest_FieldMask) *ProvisionServiceAccountToDeviceRequest_FieldMask {
	result := &ProvisionServiceAccountToDeviceRequest_FieldMask{}
	removedSelectors := make([]bool, 1)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ProvisionServiceAccountToDeviceRequest_FieldTerminalPath:
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

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ProvisionServiceAccountToDeviceRequest_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) FilterInputFields() *ProvisionServiceAccountToDeviceRequest_FieldMask {
	result := &ProvisionServiceAccountToDeviceRequest_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ProvisionServiceAccountToDeviceRequest_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseProvisionServiceAccountToDeviceRequest_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ProvisionServiceAccountToDeviceRequest_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ProvisionServiceAccountToDeviceRequest_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) AppendPath(path ProvisionServiceAccountToDeviceRequest_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ProvisionServiceAccountToDeviceRequest_FieldPath))
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) GetPaths() []ProvisionServiceAccountToDeviceRequest_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseProvisionServiceAccountToDeviceRequest_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) Set(target, source *ProvisionServiceAccountToDeviceRequest) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ProvisionServiceAccountToDeviceRequest), source.(*ProvisionServiceAccountToDeviceRequest))
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) Project(source *ProvisionServiceAccountToDeviceRequest) *ProvisionServiceAccountToDeviceRequest {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ProvisionServiceAccountToDeviceRequest{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ProvisionServiceAccountToDeviceRequest_FieldTerminalPath:
			switch tp.selector {
			case ProvisionServiceAccountToDeviceRequest_FieldPathSelectorName:
				result.Name = source.Name
			}
		}
	}
	return result
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ProvisionServiceAccountToDeviceRequest))
}

func (fieldMask *ProvisionServiceAccountToDeviceRequest_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type ProvisionServiceAccountToDeviceResponse_FieldMask struct {
	Paths []ProvisionServiceAccountToDeviceResponse_FieldPath
}

func FullProvisionServiceAccountToDeviceResponse_FieldMask() *ProvisionServiceAccountToDeviceResponse_FieldMask {
	res := &ProvisionServiceAccountToDeviceResponse_FieldMask{}
	res.Paths = append(res.Paths, &ProvisionServiceAccountToDeviceResponse_FieldTerminalPath{selector: ProvisionServiceAccountToDeviceResponse_FieldPathSelectorServiceAccount})
	return res
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) String() string {
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
func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseProvisionServiceAccountToDeviceResponse_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 1)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ProvisionServiceAccountToDeviceResponse_FieldTerminalPath); ok {
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

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseProvisionServiceAccountToDeviceResponse_FieldPath(raw)
	})
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) ProtoMessage() {}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) Subtract(other *ProvisionServiceAccountToDeviceResponse_FieldMask) *ProvisionServiceAccountToDeviceResponse_FieldMask {
	result := &ProvisionServiceAccountToDeviceResponse_FieldMask{}
	removedSelectors := make([]bool, 1)
	otherSubMasks := map[ProvisionServiceAccountToDeviceResponse_FieldPathSelector]gotenobject.FieldMask{
		ProvisionServiceAccountToDeviceResponse_FieldPathSelectorServiceAccount: &api.ServiceAccount_FieldMask{},
	}
	mySubMasks := map[ProvisionServiceAccountToDeviceResponse_FieldPathSelector]gotenobject.FieldMask{
		ProvisionServiceAccountToDeviceResponse_FieldPathSelectorServiceAccount: &api.ServiceAccount_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ProvisionServiceAccountToDeviceResponse_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *ProvisionServiceAccountToDeviceResponse_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*ProvisionServiceAccountToDeviceResponse_FieldTerminalPath); ok {
					switch tp.selector {
					case ProvisionServiceAccountToDeviceResponse_FieldPathSelectorServiceAccount:
						mySubMasks[ProvisionServiceAccountToDeviceResponse_FieldPathSelectorServiceAccount] = api.FullServiceAccount_FieldMask()
					}
				} else if tp, ok := path.(*ProvisionServiceAccountToDeviceResponse_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &ProvisionServiceAccountToDeviceResponse_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ProvisionServiceAccountToDeviceResponse_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) FilterInputFields() *ProvisionServiceAccountToDeviceResponse_FieldMask {
	result := &ProvisionServiceAccountToDeviceResponse_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ProvisionServiceAccountToDeviceResponse_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseProvisionServiceAccountToDeviceResponse_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ProvisionServiceAccountToDeviceResponse_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ProvisionServiceAccountToDeviceResponse_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) AppendPath(path ProvisionServiceAccountToDeviceResponse_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ProvisionServiceAccountToDeviceResponse_FieldPath))
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) GetPaths() []ProvisionServiceAccountToDeviceResponse_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseProvisionServiceAccountToDeviceResponse_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) Set(target, source *ProvisionServiceAccountToDeviceResponse) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ProvisionServiceAccountToDeviceResponse), source.(*ProvisionServiceAccountToDeviceResponse))
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) Project(source *ProvisionServiceAccountToDeviceResponse) *ProvisionServiceAccountToDeviceResponse {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ProvisionServiceAccountToDeviceResponse{}
	serviceAccountMask := &api.ServiceAccount_FieldMask{}
	wholeServiceAccountAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ProvisionServiceAccountToDeviceResponse_FieldTerminalPath:
			switch tp.selector {
			case ProvisionServiceAccountToDeviceResponse_FieldPathSelectorServiceAccount:
				result.ServiceAccount = source.ServiceAccount
				wholeServiceAccountAccepted = true
			}
		case *ProvisionServiceAccountToDeviceResponse_FieldSubPath:
			switch tp.selector {
			case ProvisionServiceAccountToDeviceResponse_FieldPathSelectorServiceAccount:
				serviceAccountMask.AppendPath(tp.subPath.(api.ServiceAccount_FieldPath))
			}
		}
	}
	if wholeServiceAccountAccepted == false && len(serviceAccountMask.Paths) > 0 {
		result.ServiceAccount = serviceAccountMask.Project(source.GetServiceAccount())
	}
	return result
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ProvisionServiceAccountToDeviceResponse))
}

func (fieldMask *ProvisionServiceAccountToDeviceResponse_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type RemoveServiceAccountFromDeviceRequest_FieldMask struct {
	Paths []RemoveServiceAccountFromDeviceRequest_FieldPath
}

func FullRemoveServiceAccountFromDeviceRequest_FieldMask() *RemoveServiceAccountFromDeviceRequest_FieldMask {
	res := &RemoveServiceAccountFromDeviceRequest_FieldMask{}
	res.Paths = append(res.Paths, &RemoveServiceAccountFromDeviceRequest_FieldTerminalPath{selector: RemoveServiceAccountFromDeviceRequest_FieldPathSelectorName})
	return res
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) String() string {
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
func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseRemoveServiceAccountFromDeviceRequest_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 1)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*RemoveServiceAccountFromDeviceRequest_FieldTerminalPath); ok {
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

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseRemoveServiceAccountFromDeviceRequest_FieldPath(raw)
	})
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) ProtoMessage() {}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) Subtract(other *RemoveServiceAccountFromDeviceRequest_FieldMask) *RemoveServiceAccountFromDeviceRequest_FieldMask {
	result := &RemoveServiceAccountFromDeviceRequest_FieldMask{}
	removedSelectors := make([]bool, 1)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *RemoveServiceAccountFromDeviceRequest_FieldTerminalPath:
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

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*RemoveServiceAccountFromDeviceRequest_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) FilterInputFields() *RemoveServiceAccountFromDeviceRequest_FieldMask {
	result := &RemoveServiceAccountFromDeviceRequest_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]RemoveServiceAccountFromDeviceRequest_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseRemoveServiceAccountFromDeviceRequest_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask RemoveServiceAccountFromDeviceRequest_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask RemoveServiceAccountFromDeviceRequest_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) AppendPath(path RemoveServiceAccountFromDeviceRequest_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(RemoveServiceAccountFromDeviceRequest_FieldPath))
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) GetPaths() []RemoveServiceAccountFromDeviceRequest_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseRemoveServiceAccountFromDeviceRequest_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) Set(target, source *RemoveServiceAccountFromDeviceRequest) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*RemoveServiceAccountFromDeviceRequest), source.(*RemoveServiceAccountFromDeviceRequest))
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) Project(source *RemoveServiceAccountFromDeviceRequest) *RemoveServiceAccountFromDeviceRequest {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &RemoveServiceAccountFromDeviceRequest{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *RemoveServiceAccountFromDeviceRequest_FieldTerminalPath:
			switch tp.selector {
			case RemoveServiceAccountFromDeviceRequest_FieldPathSelectorName:
				result.Name = source.Name
			}
		}
	}
	return result
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*RemoveServiceAccountFromDeviceRequest))
}

func (fieldMask *RemoveServiceAccountFromDeviceRequest_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type RemoveServiceAccountFromDeviceResponse_FieldMask struct {
	Paths []RemoveServiceAccountFromDeviceResponse_FieldPath
}

func FullRemoveServiceAccountFromDeviceResponse_FieldMask() *RemoveServiceAccountFromDeviceResponse_FieldMask {
	res := &RemoveServiceAccountFromDeviceResponse_FieldMask{}
	return res
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) String() string {
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
func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseRemoveServiceAccountFromDeviceResponse_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 0)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*RemoveServiceAccountFromDeviceResponse_FieldTerminalPath); ok {
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

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseRemoveServiceAccountFromDeviceResponse_FieldPath(raw)
	})
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) ProtoMessage() {}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) Subtract(other *RemoveServiceAccountFromDeviceResponse_FieldMask) *RemoveServiceAccountFromDeviceResponse_FieldMask {
	result := &RemoveServiceAccountFromDeviceResponse_FieldMask{}
	removedSelectors := make([]bool, 0)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *RemoveServiceAccountFromDeviceResponse_FieldTerminalPath:
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

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*RemoveServiceAccountFromDeviceResponse_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) FilterInputFields() *RemoveServiceAccountFromDeviceResponse_FieldMask {
	result := &RemoveServiceAccountFromDeviceResponse_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]RemoveServiceAccountFromDeviceResponse_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseRemoveServiceAccountFromDeviceResponse_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask RemoveServiceAccountFromDeviceResponse_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask RemoveServiceAccountFromDeviceResponse_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) AppendPath(path RemoveServiceAccountFromDeviceResponse_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(RemoveServiceAccountFromDeviceResponse_FieldPath))
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) GetPaths() []RemoveServiceAccountFromDeviceResponse_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseRemoveServiceAccountFromDeviceResponse_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) Set(target, source *RemoveServiceAccountFromDeviceResponse) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*RemoveServiceAccountFromDeviceResponse), source.(*RemoveServiceAccountFromDeviceResponse))
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) Project(source *RemoveServiceAccountFromDeviceResponse) *RemoveServiceAccountFromDeviceResponse {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &RemoveServiceAccountFromDeviceResponse{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *RemoveServiceAccountFromDeviceResponse_FieldTerminalPath:
			switch tp.selector {
			}
		}
	}
	return result
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*RemoveServiceAccountFromDeviceResponse))
}

func (fieldMask *RemoveServiceAccountFromDeviceResponse_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
