// Code generated by protoc-gen-goten-object
// File: edgelq/devices/proto/v1/device_type.proto
// DO NOT EDIT!!!

package device_type

import (
	"encoding/json"
	"strings"

	firestorepb "google.golang.org/genproto/googleapis/firestore/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	meta "github.com/cloudwan/goten-sdk/types/meta"
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
	_ = googlefieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldMask)
)

// make sure we're using proto imports
var (
	_ = &meta.Meta{}
)

type DeviceType_FieldMask struct {
	Paths []DeviceType_FieldPath
}

func FullDeviceType_FieldMask() *DeviceType_FieldMask {
	res := &DeviceType_FieldMask{}
	res.Paths = append(res.Paths, &DeviceType_FieldTerminalPath{selector: DeviceType_FieldPathSelectorName})
	res.Paths = append(res.Paths, &DeviceType_FieldTerminalPath{selector: DeviceType_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &DeviceType_FieldTerminalPath{selector: DeviceType_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &DeviceType_FieldTerminalPath{selector: DeviceType_FieldPathSelectorHardware})
	res.Paths = append(res.Paths, &DeviceType_FieldTerminalPath{selector: DeviceType_FieldPathSelectorArchitecture})
	res.Paths = append(res.Paths, &DeviceType_FieldTerminalPath{selector: DeviceType_FieldPathSelectorDescription})
	return res
}

func (fieldMask *DeviceType_FieldMask) String() string {
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
func (fieldMask *DeviceType_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *DeviceType_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseDeviceType_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *DeviceType_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 6)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*DeviceType_FieldTerminalPath); ok {
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

func (fieldMask *DeviceType_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseDeviceType_FieldPath(raw)
	})
}

func (fieldMask *DeviceType_FieldMask) ProtoMessage() {}

func (fieldMask *DeviceType_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *DeviceType_FieldMask) Subtract(other *DeviceType_FieldMask) *DeviceType_FieldMask {
	result := &DeviceType_FieldMask{}
	removedSelectors := make([]bool, 6)
	otherSubMasks := map[DeviceType_FieldPathSelector]gotenobject.FieldMask{
		DeviceType_FieldPathSelectorMetadata: &meta.Meta_FieldMask{},
	}
	mySubMasks := map[DeviceType_FieldPathSelector]gotenobject.FieldMask{
		DeviceType_FieldPathSelectorMetadata: &meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *DeviceType_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *DeviceType_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*DeviceType_FieldTerminalPath); ok {
					switch tp.selector {
					case DeviceType_FieldPathSelectorMetadata:
						mySubMasks[DeviceType_FieldPathSelectorMetadata] = meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*DeviceType_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &DeviceType_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *DeviceType_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*DeviceType_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *DeviceType_FieldMask) FilterInputFields() *DeviceType_FieldMask {
	result := &DeviceType_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case DeviceType_FieldPathSelectorMetadata:
			if _, ok := path.(*DeviceType_FieldTerminalPath); ok {
				for _, subpath := range meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &DeviceType_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*DeviceType_FieldSubPath); ok {
				selectedMask := &meta.Meta_FieldMask{
					Paths: []meta.Meta_FieldPath{sub.subPath.(meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &DeviceType_FieldSubPath{selector: DeviceType_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *DeviceType_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *DeviceType_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]DeviceType_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseDeviceType_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask DeviceType_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *DeviceType_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *DeviceType_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask DeviceType_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *DeviceType_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *DeviceType_FieldMask) AppendPath(path DeviceType_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *DeviceType_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(DeviceType_FieldPath))
}

func (fieldMask *DeviceType_FieldMask) GetPaths() []DeviceType_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *DeviceType_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *DeviceType_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseDeviceType_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *DeviceType_FieldMask) Set(target, source *DeviceType) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *DeviceType_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*DeviceType), source.(*DeviceType))
}

func (fieldMask *DeviceType_FieldMask) Project(source *DeviceType) *DeviceType {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &DeviceType{}
	metadataMask := &meta.Meta_FieldMask{}
	wholeMetadataAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *DeviceType_FieldTerminalPath:
			switch tp.selector {
			case DeviceType_FieldPathSelectorName:
				result.Name = source.Name
			case DeviceType_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case DeviceType_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case DeviceType_FieldPathSelectorHardware:
				result.Hardware = source.Hardware
			case DeviceType_FieldPathSelectorArchitecture:
				result.Architecture = source.Architecture
			case DeviceType_FieldPathSelectorDescription:
				result.Description = source.Description
			}
		case *DeviceType_FieldSubPath:
			switch tp.selector {
			case DeviceType_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(meta.Meta_FieldPath))
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	return result
}

func (fieldMask *DeviceType_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*DeviceType))
}

func (fieldMask *DeviceType_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
