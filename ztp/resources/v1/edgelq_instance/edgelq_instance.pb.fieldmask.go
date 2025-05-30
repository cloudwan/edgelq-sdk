// Code generated by protoc-gen-goten-object
// File: edgelq/ztp/proto/v1/edgelq_instance.proto
// DO NOT EDIT!!!

package edgelq_instance

import (
	"encoding/json"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(json.Marshaler)
	_ = strings.Builder{}

	_ = codes.NotFound
	_ = status.Status{}
	_ = new(proto.Message)
	_ = new(preflect.Message)
	_ = googlefieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldMask)
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &meta.Meta{}
)

type EdgelqInstance_FieldMask struct {
	Paths []EdgelqInstance_FieldPath
}

func FullEdgelqInstance_FieldMask() *EdgelqInstance_FieldMask {
	res := &EdgelqInstance_FieldMask{}
	res.Paths = append(res.Paths, &EdgelqInstance_FieldTerminalPath{selector: EdgelqInstance_FieldPathSelectorName})
	res.Paths = append(res.Paths, &EdgelqInstance_FieldTerminalPath{selector: EdgelqInstance_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &EdgelqInstance_FieldTerminalPath{selector: EdgelqInstance_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &EdgelqInstance_FieldTerminalPath{selector: EdgelqInstance_FieldPathSelectorControllerDomain})
	res.Paths = append(res.Paths, &EdgelqInstance_FieldTerminalPath{selector: EdgelqInstance_FieldPathSelectorEndpoints})
	return res
}

func (fieldMask *EdgelqInstance_FieldMask) String() string {
	if fieldMask == nil {
		return "<nil>"
	}
	pathsStr := make([]string, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		pathsStr = append(pathsStr, path.String())
	}
	return strings.Join(pathsStr, ", ")
}

func (fieldMask *EdgelqInstance_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 5)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*EdgelqInstance_FieldTerminalPath); ok {
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

func (fieldMask *EdgelqInstance_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseEdgelqInstance_FieldPath(raw)
	})
}

func (fieldMask *EdgelqInstance_FieldMask) ProtoMessage() {}

func (fieldMask *EdgelqInstance_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *EdgelqInstance_FieldMask) Subtract(other *EdgelqInstance_FieldMask) *EdgelqInstance_FieldMask {
	result := &EdgelqInstance_FieldMask{}
	removedSelectors := make([]bool, 5)
	otherSubMasks := map[EdgelqInstance_FieldPathSelector]gotenobject.FieldMask{
		EdgelqInstance_FieldPathSelectorMetadata: &meta.Meta_FieldMask{},
	}
	mySubMasks := map[EdgelqInstance_FieldPathSelector]gotenobject.FieldMask{
		EdgelqInstance_FieldPathSelectorMetadata: &meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *EdgelqInstance_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *EdgelqInstance_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*EdgelqInstance_FieldTerminalPath); ok {
					switch tp.selector {
					case EdgelqInstance_FieldPathSelectorMetadata:
						mySubMasks[EdgelqInstance_FieldPathSelectorMetadata] = meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*EdgelqInstance_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &EdgelqInstance_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *EdgelqInstance_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*EdgelqInstance_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *EdgelqInstance_FieldMask) FilterInputFields() *EdgelqInstance_FieldMask {
	result := &EdgelqInstance_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case EdgelqInstance_FieldPathSelectorMetadata:
			if _, ok := path.(*EdgelqInstance_FieldTerminalPath); ok {
				for _, subpath := range meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &EdgelqInstance_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*EdgelqInstance_FieldSubPath); ok {
				selectedMask := &meta.Meta_FieldMask{
					Paths: []meta.Meta_FieldPath{sub.subPath.(meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &EdgelqInstance_FieldSubPath{selector: EdgelqInstance_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *EdgelqInstance_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *EdgelqInstance_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]EdgelqInstance_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseEdgelqInstance_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask EdgelqInstance_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *EdgelqInstance_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *EdgelqInstance_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask EdgelqInstance_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *EdgelqInstance_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *EdgelqInstance_FieldMask) AppendPath(path EdgelqInstance_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *EdgelqInstance_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(EdgelqInstance_FieldPath))
}

func (fieldMask *EdgelqInstance_FieldMask) GetPaths() []EdgelqInstance_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *EdgelqInstance_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *EdgelqInstance_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseEdgelqInstance_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *EdgelqInstance_FieldMask) Set(target, source *EdgelqInstance) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *EdgelqInstance_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*EdgelqInstance), source.(*EdgelqInstance))
}

func (fieldMask *EdgelqInstance_FieldMask) Project(source *EdgelqInstance) *EdgelqInstance {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &EdgelqInstance{}
	metadataMask := &meta.Meta_FieldMask{}
	wholeMetadataAccepted := false
	var endpointsMapKeys []string
	wholeEndpointsAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *EdgelqInstance_FieldTerminalPath:
			switch tp.selector {
			case EdgelqInstance_FieldPathSelectorName:
				result.Name = source.Name
			case EdgelqInstance_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case EdgelqInstance_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case EdgelqInstance_FieldPathSelectorControllerDomain:
				result.ControllerDomain = source.ControllerDomain
			case EdgelqInstance_FieldPathSelectorEndpoints:
				result.Endpoints = source.Endpoints
				wholeEndpointsAccepted = true
			}
		case *EdgelqInstance_FieldSubPath:
			switch tp.selector {
			case EdgelqInstance_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(meta.Meta_FieldPath))
			}
		case *EdgelqInstance_FieldPathMap:
			switch tp.selector {
			case EdgelqInstance_FieldPathSelectorEndpoints:
				endpointsMapKeys = append(endpointsMapKeys, tp.key)
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	if wholeEndpointsAccepted == false && len(endpointsMapKeys) > 0 && source.GetEndpoints() != nil {
		copiedMap := map[string]string{}
		sourceMap := source.GetEndpoints()
		for _, key := range endpointsMapKeys {
			copiedMap[key] = sourceMap[key]
		}
		result.Endpoints = copiedMap
	}
	return result
}

func (fieldMask *EdgelqInstance_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*EdgelqInstance))
}

func (fieldMask *EdgelqInstance_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
