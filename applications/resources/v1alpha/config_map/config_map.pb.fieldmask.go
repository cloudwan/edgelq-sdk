// Code generated by protoc-gen-goten-object
// File: edgelq/applications/proto/v1alpha/config_map.proto
// DO NOT EDIT!!!

package config_map

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
	project "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha/project"
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
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
	_ = &project.Project{}
	_ = &ntt_meta.Meta{}
)

type ConfigMap_FieldMask struct {
	Paths []ConfigMap_FieldPath
}

func FullConfigMap_FieldMask() *ConfigMap_FieldMask {
	res := &ConfigMap_FieldMask{}
	res.Paths = append(res.Paths, &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorName})
	res.Paths = append(res.Paths, &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorData})
	res.Paths = append(res.Paths, &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorBinaryData})
	return res
}

func (fieldMask *ConfigMap_FieldMask) String() string {
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
func (fieldMask *ConfigMap_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ConfigMap_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseConfigMap_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ConfigMap_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 5)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ConfigMap_FieldTerminalPath); ok {
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

func (fieldMask *ConfigMap_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseConfigMap_FieldPath(raw)
	})
}

func (fieldMask *ConfigMap_FieldMask) ProtoMessage() {}

func (fieldMask *ConfigMap_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ConfigMap_FieldMask) Subtract(other *ConfigMap_FieldMask) *ConfigMap_FieldMask {
	result := &ConfigMap_FieldMask{}
	removedSelectors := make([]bool, 5)
	otherSubMasks := map[ConfigMap_FieldPathSelector]gotenobject.FieldMask{
		ConfigMap_FieldPathSelectorMetadata: &ntt_meta.Meta_FieldMask{},
	}
	mySubMasks := map[ConfigMap_FieldPathSelector]gotenobject.FieldMask{
		ConfigMap_FieldPathSelectorMetadata: &ntt_meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ConfigMap_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *ConfigMap_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*ConfigMap_FieldTerminalPath); ok {
					switch tp.selector {
					case ConfigMap_FieldPathSelectorMetadata:
						mySubMasks[ConfigMap_FieldPathSelectorMetadata] = ntt_meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*ConfigMap_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &ConfigMap_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *ConfigMap_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ConfigMap_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ConfigMap_FieldMask) FilterInputFields() *ConfigMap_FieldMask {
	result := &ConfigMap_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case ConfigMap_FieldPathSelectorMetadata:
			if _, ok := path.(*ConfigMap_FieldTerminalPath); ok {
				for _, subpath := range ntt_meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ConfigMap_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*ConfigMap_FieldSubPath); ok {
				selectedMask := &ntt_meta.Meta_FieldMask{
					Paths: []ntt_meta.Meta_FieldPath{sub.subPath.(ntt_meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ConfigMap_FieldSubPath{selector: ConfigMap_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ConfigMap_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ConfigMap_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ConfigMap_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseConfigMap_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ConfigMap_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ConfigMap_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ConfigMap_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ConfigMap_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ConfigMap_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ConfigMap_FieldMask) AppendPath(path ConfigMap_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ConfigMap_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ConfigMap_FieldPath))
}

func (fieldMask *ConfigMap_FieldMask) GetPaths() []ConfigMap_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ConfigMap_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ConfigMap_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseConfigMap_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ConfigMap_FieldMask) Set(target, source *ConfigMap) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ConfigMap_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ConfigMap), source.(*ConfigMap))
}

func (fieldMask *ConfigMap_FieldMask) Project(source *ConfigMap) *ConfigMap {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ConfigMap{}
	metadataMask := &ntt_meta.Meta_FieldMask{}
	wholeMetadataAccepted := false
	var dataMapKeys []string
	wholeDataAccepted := false
	var binaryDataMapKeys []string
	wholeBinaryDataAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ConfigMap_FieldTerminalPath:
			switch tp.selector {
			case ConfigMap_FieldPathSelectorName:
				result.Name = source.Name
			case ConfigMap_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case ConfigMap_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case ConfigMap_FieldPathSelectorData:
				result.Data = source.Data
				wholeDataAccepted = true
			case ConfigMap_FieldPathSelectorBinaryData:
				result.BinaryData = source.BinaryData
				wholeBinaryDataAccepted = true
			}
		case *ConfigMap_FieldSubPath:
			switch tp.selector {
			case ConfigMap_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(ntt_meta.Meta_FieldPath))
			}
		case *ConfigMap_FieldPathMap:
			switch tp.selector {
			case ConfigMap_FieldPathSelectorData:
				dataMapKeys = append(dataMapKeys, tp.key)
			case ConfigMap_FieldPathSelectorBinaryData:
				binaryDataMapKeys = append(binaryDataMapKeys, tp.key)
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	if wholeDataAccepted == false && len(dataMapKeys) > 0 && source.GetData() != nil {
		copiedMap := map[string]string{}
		sourceMap := source.GetData()
		for _, key := range dataMapKeys {
			copiedMap[key] = sourceMap[key]
		}
		result.Data = copiedMap
	}
	if wholeBinaryDataAccepted == false && len(binaryDataMapKeys) > 0 && source.GetBinaryData() != nil {
		copiedMap := map[string][]byte{}
		sourceMap := source.GetBinaryData()
		for _, key := range binaryDataMapKeys {
			copiedMap[key] = sourceMap[key]
		}
		result.BinaryData = copiedMap
	}
	return result
}

func (fieldMask *ConfigMap_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ConfigMap))
}

func (fieldMask *ConfigMap_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
