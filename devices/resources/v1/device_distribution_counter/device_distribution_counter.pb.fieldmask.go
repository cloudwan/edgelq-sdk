// Code generated by protoc-gen-goten-object
// File: edgelq/devices/proto/v1/device_distribution_counter.proto
// DO NOT EDIT!!!

package device_distribution_counter

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
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1/project"
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

type DeviceDistributionCounter_FieldMask struct {
	Paths []DeviceDistributionCounter_FieldPath
}

func FullDeviceDistributionCounter_FieldMask() *DeviceDistributionCounter_FieldMask {
	res := &DeviceDistributionCounter_FieldMask{}
	res.Paths = append(res.Paths, &DeviceDistributionCounter_FieldTerminalPath{selector: DeviceDistributionCounter_FieldPathSelectorName})
	res.Paths = append(res.Paths, &DeviceDistributionCounter_FieldTerminalPath{selector: DeviceDistributionCounter_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &DeviceDistributionCounter_FieldTerminalPath{selector: DeviceDistributionCounter_FieldPathSelectorTotalCount})
	res.Paths = append(res.Paths, &DeviceDistributionCounter_FieldTerminalPath{selector: DeviceDistributionCounter_FieldPathSelectorOnlineCount})
	return res
}

func (fieldMask *DeviceDistributionCounter_FieldMask) String() string {
	if fieldMask == nil {
		return "<nil>"
	}
	pathsStr := make([]string, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		pathsStr = append(pathsStr, path.String())
	}
	return strings.Join(pathsStr, ", ")
}

func (fieldMask *DeviceDistributionCounter_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 4)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*DeviceDistributionCounter_FieldTerminalPath); ok {
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

func (fieldMask *DeviceDistributionCounter_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseDeviceDistributionCounter_FieldPath(raw)
	})
}

func (fieldMask *DeviceDistributionCounter_FieldMask) ProtoMessage() {}

func (fieldMask *DeviceDistributionCounter_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *DeviceDistributionCounter_FieldMask) Subtract(other *DeviceDistributionCounter_FieldMask) *DeviceDistributionCounter_FieldMask {
	result := &DeviceDistributionCounter_FieldMask{}
	removedSelectors := make([]bool, 4)
	otherSubMasks := map[DeviceDistributionCounter_FieldPathSelector]gotenobject.FieldMask{
		DeviceDistributionCounter_FieldPathSelectorMetadata: &meta.Meta_FieldMask{},
	}
	mySubMasks := map[DeviceDistributionCounter_FieldPathSelector]gotenobject.FieldMask{
		DeviceDistributionCounter_FieldPathSelectorMetadata: &meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *DeviceDistributionCounter_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *DeviceDistributionCounter_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*DeviceDistributionCounter_FieldTerminalPath); ok {
					switch tp.selector {
					case DeviceDistributionCounter_FieldPathSelectorMetadata:
						mySubMasks[DeviceDistributionCounter_FieldPathSelectorMetadata] = meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*DeviceDistributionCounter_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &DeviceDistributionCounter_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *DeviceDistributionCounter_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*DeviceDistributionCounter_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *DeviceDistributionCounter_FieldMask) FilterInputFields() *DeviceDistributionCounter_FieldMask {
	result := &DeviceDistributionCounter_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case DeviceDistributionCounter_FieldPathSelectorMetadata:
			if _, ok := path.(*DeviceDistributionCounter_FieldTerminalPath); ok {
				for _, subpath := range meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &DeviceDistributionCounter_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*DeviceDistributionCounter_FieldSubPath); ok {
				selectedMask := &meta.Meta_FieldMask{
					Paths: []meta.Meta_FieldPath{sub.subPath.(meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &DeviceDistributionCounter_FieldSubPath{selector: DeviceDistributionCounter_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *DeviceDistributionCounter_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *DeviceDistributionCounter_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]DeviceDistributionCounter_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseDeviceDistributionCounter_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask DeviceDistributionCounter_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *DeviceDistributionCounter_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *DeviceDistributionCounter_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask DeviceDistributionCounter_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *DeviceDistributionCounter_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *DeviceDistributionCounter_FieldMask) AppendPath(path DeviceDistributionCounter_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *DeviceDistributionCounter_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(DeviceDistributionCounter_FieldPath))
}

func (fieldMask *DeviceDistributionCounter_FieldMask) GetPaths() []DeviceDistributionCounter_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *DeviceDistributionCounter_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *DeviceDistributionCounter_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseDeviceDistributionCounter_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *DeviceDistributionCounter_FieldMask) Set(target, source *DeviceDistributionCounter) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *DeviceDistributionCounter_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*DeviceDistributionCounter), source.(*DeviceDistributionCounter))
}

func (fieldMask *DeviceDistributionCounter_FieldMask) Project(source *DeviceDistributionCounter) *DeviceDistributionCounter {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &DeviceDistributionCounter{}
	metadataMask := &meta.Meta_FieldMask{}
	wholeMetadataAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *DeviceDistributionCounter_FieldTerminalPath:
			switch tp.selector {
			case DeviceDistributionCounter_FieldPathSelectorName:
				result.Name = source.Name
			case DeviceDistributionCounter_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case DeviceDistributionCounter_FieldPathSelectorTotalCount:
				result.TotalCount = source.TotalCount
			case DeviceDistributionCounter_FieldPathSelectorOnlineCount:
				result.OnlineCount = source.OnlineCount
			}
		case *DeviceDistributionCounter_FieldSubPath:
			switch tp.selector {
			case DeviceDistributionCounter_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(meta.Meta_FieldPath))
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	return result
}

func (fieldMask *DeviceDistributionCounter_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*DeviceDistributionCounter))
}

func (fieldMask *DeviceDistributionCounter_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
