// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1/service_account.proto
// DO NOT EDIT!!!

package service_account

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
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
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

type ServiceAccount_FieldMask struct {
	Paths []ServiceAccount_FieldPath
}

func FullServiceAccount_FieldMask() *ServiceAccount_FieldMask {
	res := &ServiceAccount_FieldMask{}
	res.Paths = append(res.Paths, &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorName})
	res.Paths = append(res.Paths, &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorDescription})
	res.Paths = append(res.Paths, &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorEmail})
	res.Paths = append(res.Paths, &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorKind})
	return res
}

func (fieldMask *ServiceAccount_FieldMask) String() string {
	if fieldMask == nil {
		return "<nil>"
	}
	pathsStr := make([]string, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		pathsStr = append(pathsStr, path.String())
	}
	return strings.Join(pathsStr, ", ")
}

func (fieldMask *ServiceAccount_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 6)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ServiceAccount_FieldTerminalPath); ok {
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

func (fieldMask *ServiceAccount_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseServiceAccount_FieldPath(raw)
	})
}

func (fieldMask *ServiceAccount_FieldMask) ProtoMessage() {}

func (fieldMask *ServiceAccount_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ServiceAccount_FieldMask) Subtract(other *ServiceAccount_FieldMask) *ServiceAccount_FieldMask {
	result := &ServiceAccount_FieldMask{}
	removedSelectors := make([]bool, 6)
	otherSubMasks := map[ServiceAccount_FieldPathSelector]gotenobject.FieldMask{
		ServiceAccount_FieldPathSelectorMetadata: &meta.Meta_FieldMask{},
	}
	mySubMasks := map[ServiceAccount_FieldPathSelector]gotenobject.FieldMask{
		ServiceAccount_FieldPathSelectorMetadata: &meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ServiceAccount_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *ServiceAccount_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*ServiceAccount_FieldTerminalPath); ok {
					switch tp.selector {
					case ServiceAccount_FieldPathSelectorMetadata:
						mySubMasks[ServiceAccount_FieldPathSelectorMetadata] = meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*ServiceAccount_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &ServiceAccount_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *ServiceAccount_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ServiceAccount_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ServiceAccount_FieldMask) FilterInputFields() *ServiceAccount_FieldMask {
	result := &ServiceAccount_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case ServiceAccount_FieldPathSelectorEmail:
		case ServiceAccount_FieldPathSelectorMetadata:
			if _, ok := path.(*ServiceAccount_FieldTerminalPath); ok {
				for _, subpath := range meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ServiceAccount_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*ServiceAccount_FieldSubPath); ok {
				selectedMask := &meta.Meta_FieldMask{
					Paths: []meta.Meta_FieldPath{sub.subPath.(meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ServiceAccount_FieldSubPath{selector: ServiceAccount_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ServiceAccount_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ServiceAccount_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ServiceAccount_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseServiceAccount_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ServiceAccount_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ServiceAccount_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ServiceAccount_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ServiceAccount_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ServiceAccount_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ServiceAccount_FieldMask) AppendPath(path ServiceAccount_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ServiceAccount_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ServiceAccount_FieldPath))
}

func (fieldMask *ServiceAccount_FieldMask) GetPaths() []ServiceAccount_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ServiceAccount_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ServiceAccount_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseServiceAccount_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ServiceAccount_FieldMask) Set(target, source *ServiceAccount) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ServiceAccount_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ServiceAccount), source.(*ServiceAccount))
}

func (fieldMask *ServiceAccount_FieldMask) Project(source *ServiceAccount) *ServiceAccount {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ServiceAccount{}
	metadataMask := &meta.Meta_FieldMask{}
	wholeMetadataAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ServiceAccount_FieldTerminalPath:
			switch tp.selector {
			case ServiceAccount_FieldPathSelectorName:
				result.Name = source.Name
			case ServiceAccount_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case ServiceAccount_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case ServiceAccount_FieldPathSelectorDescription:
				result.Description = source.Description
			case ServiceAccount_FieldPathSelectorEmail:
				result.Email = source.Email
			case ServiceAccount_FieldPathSelectorKind:
				result.Kind = source.Kind
			}
		case *ServiceAccount_FieldSubPath:
			switch tp.selector {
			case ServiceAccount_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(meta.Meta_FieldPath))
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	return result
}

func (fieldMask *ServiceAccount_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ServiceAccount))
}

func (fieldMask *ServiceAccount_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
