// Code generated by protoc-gen-goten-object
// File: edgelq/monitoring/proto/v3/project.proto
// DO NOT EDIT!!!

package project

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
	multi_region_policy "github.com/cloudwan/goten-sdk/types/multi_region_policy"
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
	_ = &multi_region_policy.MultiRegionPolicy{}
)

type Project_FieldMask struct {
	Paths []Project_FieldPath
}

func FullProject_FieldMask() *Project_FieldMask {
	res := &Project_FieldMask{}
	res.Paths = append(res.Paths, &Project_FieldTerminalPath{selector: Project_FieldPathSelectorName})
	res.Paths = append(res.Paths, &Project_FieldTerminalPath{selector: Project_FieldPathSelectorTitle})
	res.Paths = append(res.Paths, &Project_FieldTerminalPath{selector: Project_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &Project_FieldTerminalPath{selector: Project_FieldPathSelectorMultiRegionPolicy})
	return res
}

func (fieldMask *Project_FieldMask) String() string {
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
func (fieldMask *Project_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *Project_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseProject_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Project_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 4)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*Project_FieldTerminalPath); ok {
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

func (fieldMask *Project_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseProject_FieldPath(raw)
	})
}

func (fieldMask *Project_FieldMask) ProtoMessage() {}

func (fieldMask *Project_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Project_FieldMask) Subtract(other *Project_FieldMask) *Project_FieldMask {
	result := &Project_FieldMask{}
	removedSelectors := make([]bool, 4)
	otherSubMasks := map[Project_FieldPathSelector]gotenobject.FieldMask{
		Project_FieldPathSelectorMetadata:          &meta.Meta_FieldMask{},
		Project_FieldPathSelectorMultiRegionPolicy: &multi_region_policy.MultiRegionPolicy_FieldMask{},
	}
	mySubMasks := map[Project_FieldPathSelector]gotenobject.FieldMask{
		Project_FieldPathSelectorMetadata:          &meta.Meta_FieldMask{},
		Project_FieldPathSelectorMultiRegionPolicy: &multi_region_policy.MultiRegionPolicy_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *Project_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *Project_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*Project_FieldTerminalPath); ok {
					switch tp.selector {
					case Project_FieldPathSelectorMetadata:
						mySubMasks[Project_FieldPathSelectorMetadata] = meta.FullMeta_FieldMask()
					case Project_FieldPathSelectorMultiRegionPolicy:
						mySubMasks[Project_FieldPathSelectorMultiRegionPolicy] = multi_region_policy.FullMultiRegionPolicy_FieldMask()
					}
				} else if tp, ok := path.(*Project_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &Project_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *Project_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Project_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Project_FieldMask) FilterInputFields() *Project_FieldMask {
	result := &Project_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case Project_FieldPathSelectorMetadata:
			if _, ok := path.(*Project_FieldTerminalPath); ok {
				for _, subpath := range meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &Project_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*Project_FieldSubPath); ok {
				selectedMask := &meta.Meta_FieldMask{
					Paths: []meta.Meta_FieldPath{sub.subPath.(meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &Project_FieldSubPath{selector: Project_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Project_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Project_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]Project_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseProject_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Project_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Project_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Project_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Project_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Project_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Project_FieldMask) AppendPath(path Project_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Project_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(Project_FieldPath))
}

func (fieldMask *Project_FieldMask) GetPaths() []Project_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Project_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Project_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseProject_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Project_FieldMask) Set(target, source *Project) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Project_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Project), source.(*Project))
}

func (fieldMask *Project_FieldMask) Project(source *Project) *Project {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Project{}
	metadataMask := &meta.Meta_FieldMask{}
	wholeMetadataAccepted := false
	multiRegionPolicyMask := &multi_region_policy.MultiRegionPolicy_FieldMask{}
	wholeMultiRegionPolicyAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *Project_FieldTerminalPath:
			switch tp.selector {
			case Project_FieldPathSelectorName:
				result.Name = source.Name
			case Project_FieldPathSelectorTitle:
				result.Title = source.Title
			case Project_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case Project_FieldPathSelectorMultiRegionPolicy:
				result.MultiRegionPolicy = source.MultiRegionPolicy
				wholeMultiRegionPolicyAccepted = true
			}
		case *Project_FieldSubPath:
			switch tp.selector {
			case Project_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(meta.Meta_FieldPath))
			case Project_FieldPathSelectorMultiRegionPolicy:
				multiRegionPolicyMask.AppendPath(tp.subPath.(multi_region_policy.MultiRegionPolicy_FieldPath))
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	if wholeMultiRegionPolicyAccepted == false && len(multiRegionPolicyMask.Paths) > 0 {
		result.MultiRegionPolicy = multiRegionPolicyMask.Project(source.GetMultiRegionPolicy())
	}
	return result
}

func (fieldMask *Project_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Project))
}

func (fieldMask *Project_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
