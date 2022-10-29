// Code generated by protoc-gen-goten-object
// File: edgelq/audit/proto/v1alpha/method_descriptor.proto
// DO NOT EDIT!!!

package method_descriptor

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
	audit_common "github.com/cloudwan/edgelq-sdk/audit/common/v1alpha"
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
	_ = &audit_common.Authentication{}
	_ = &ntt_meta.Meta{}
)

type MethodDescriptor_FieldMask struct {
	Paths []MethodDescriptor_FieldPath
}

func FullMethodDescriptor_FieldMask() *MethodDescriptor_FieldMask {
	res := &MethodDescriptor_FieldMask{}
	res.Paths = append(res.Paths, &MethodDescriptor_FieldTerminalPath{selector: MethodDescriptor_FieldPathSelectorName})
	res.Paths = append(res.Paths, &MethodDescriptor_FieldTerminalPath{selector: MethodDescriptor_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &MethodDescriptor_FieldTerminalPath{selector: MethodDescriptor_FieldPathSelectorDescription})
	res.Paths = append(res.Paths, &MethodDescriptor_FieldTerminalPath{selector: MethodDescriptor_FieldPathSelectorLabels})
	res.Paths = append(res.Paths, &MethodDescriptor_FieldTerminalPath{selector: MethodDescriptor_FieldPathSelectorPromotedLabelKeySets})
	res.Paths = append(res.Paths, &MethodDescriptor_FieldTerminalPath{selector: MethodDescriptor_FieldPathSelectorVersions})
	res.Paths = append(res.Paths, &MethodDescriptor_FieldTerminalPath{selector: MethodDescriptor_FieldPathSelectorMetadata})
	return res
}

func (fieldMask *MethodDescriptor_FieldMask) String() string {
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
func (fieldMask *MethodDescriptor_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *MethodDescriptor_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseMethodDescriptor_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *MethodDescriptor_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 7)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*MethodDescriptor_FieldTerminalPath); ok {
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

func (fieldMask *MethodDescriptor_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseMethodDescriptor_FieldPath(raw)
	})
}

func (fieldMask *MethodDescriptor_FieldMask) ProtoMessage() {}

func (fieldMask *MethodDescriptor_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *MethodDescriptor_FieldMask) Subtract(other *MethodDescriptor_FieldMask) *MethodDescriptor_FieldMask {
	result := &MethodDescriptor_FieldMask{}
	removedSelectors := make([]bool, 7)
	otherSubMasks := map[MethodDescriptor_FieldPathSelector]gotenobject.FieldMask{
		MethodDescriptor_FieldPathSelectorLabels:               &audit_common.LabelDescriptor_FieldMask{},
		MethodDescriptor_FieldPathSelectorPromotedLabelKeySets: &audit_common.LabelKeySet_FieldMask{},
		MethodDescriptor_FieldPathSelectorMetadata:             &ntt_meta.Meta_FieldMask{},
	}
	mySubMasks := map[MethodDescriptor_FieldPathSelector]gotenobject.FieldMask{
		MethodDescriptor_FieldPathSelectorLabels:               &audit_common.LabelDescriptor_FieldMask{},
		MethodDescriptor_FieldPathSelectorPromotedLabelKeySets: &audit_common.LabelKeySet_FieldMask{},
		MethodDescriptor_FieldPathSelectorMetadata:             &ntt_meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *MethodDescriptor_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *MethodDescriptor_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*MethodDescriptor_FieldTerminalPath); ok {
					switch tp.selector {
					case MethodDescriptor_FieldPathSelectorLabels:
						mySubMasks[MethodDescriptor_FieldPathSelectorLabels] = audit_common.FullLabelDescriptor_FieldMask()
					case MethodDescriptor_FieldPathSelectorPromotedLabelKeySets:
						mySubMasks[MethodDescriptor_FieldPathSelectorPromotedLabelKeySets] = audit_common.FullLabelKeySet_FieldMask()
					case MethodDescriptor_FieldPathSelectorMetadata:
						mySubMasks[MethodDescriptor_FieldPathSelectorMetadata] = ntt_meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*MethodDescriptor_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &MethodDescriptor_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *MethodDescriptor_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*MethodDescriptor_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *MethodDescriptor_FieldMask) FilterInputFields() *MethodDescriptor_FieldMask {
	result := &MethodDescriptor_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case MethodDescriptor_FieldPathSelectorMetadata:
			if _, ok := path.(*MethodDescriptor_FieldTerminalPath); ok {
				for _, subpath := range ntt_meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &MethodDescriptor_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*MethodDescriptor_FieldSubPath); ok {
				selectedMask := &ntt_meta.Meta_FieldMask{
					Paths: []ntt_meta.Meta_FieldPath{sub.subPath.(ntt_meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &MethodDescriptor_FieldSubPath{selector: MethodDescriptor_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *MethodDescriptor_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *MethodDescriptor_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]MethodDescriptor_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseMethodDescriptor_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask MethodDescriptor_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *MethodDescriptor_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *MethodDescriptor_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask MethodDescriptor_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *MethodDescriptor_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *MethodDescriptor_FieldMask) AppendPath(path MethodDescriptor_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *MethodDescriptor_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(MethodDescriptor_FieldPath))
}

func (fieldMask *MethodDescriptor_FieldMask) GetPaths() []MethodDescriptor_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *MethodDescriptor_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *MethodDescriptor_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseMethodDescriptor_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *MethodDescriptor_FieldMask) Set(target, source *MethodDescriptor) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *MethodDescriptor_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*MethodDescriptor), source.(*MethodDescriptor))
}

func (fieldMask *MethodDescriptor_FieldMask) Project(source *MethodDescriptor) *MethodDescriptor {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &MethodDescriptor{}
	labelsMask := &audit_common.LabelDescriptor_FieldMask{}
	wholeLabelsAccepted := false
	promotedLabelKeySetsMask := &audit_common.LabelKeySet_FieldMask{}
	wholePromotedLabelKeySetsAccepted := false
	metadataMask := &ntt_meta.Meta_FieldMask{}
	wholeMetadataAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *MethodDescriptor_FieldTerminalPath:
			switch tp.selector {
			case MethodDescriptor_FieldPathSelectorName:
				result.Name = source.Name
			case MethodDescriptor_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case MethodDescriptor_FieldPathSelectorDescription:
				result.Description = source.Description
			case MethodDescriptor_FieldPathSelectorLabels:
				result.Labels = source.Labels
				wholeLabelsAccepted = true
			case MethodDescriptor_FieldPathSelectorPromotedLabelKeySets:
				result.PromotedLabelKeySets = source.PromotedLabelKeySets
				wholePromotedLabelKeySetsAccepted = true
			case MethodDescriptor_FieldPathSelectorVersions:
				result.Versions = source.Versions
			case MethodDescriptor_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			}
		case *MethodDescriptor_FieldSubPath:
			switch tp.selector {
			case MethodDescriptor_FieldPathSelectorLabels:
				labelsMask.AppendPath(tp.subPath.(audit_common.LabelDescriptor_FieldPath))
			case MethodDescriptor_FieldPathSelectorPromotedLabelKeySets:
				promotedLabelKeySetsMask.AppendPath(tp.subPath.(audit_common.LabelKeySet_FieldPath))
			case MethodDescriptor_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(ntt_meta.Meta_FieldPath))
			}
		}
	}
	if wholeLabelsAccepted == false && len(labelsMask.Paths) > 0 {
		for _, sourceItem := range source.GetLabels() {
			result.Labels = append(result.Labels, labelsMask.Project(sourceItem))
		}
	}
	if wholePromotedLabelKeySetsAccepted == false && len(promotedLabelKeySetsMask.Paths) > 0 {
		for _, sourceItem := range source.GetPromotedLabelKeySets() {
			result.PromotedLabelKeySets = append(result.PromotedLabelKeySets, promotedLabelKeySetsMask.Project(sourceItem))
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	return result
}

func (fieldMask *MethodDescriptor_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*MethodDescriptor))
}

func (fieldMask *MethodDescriptor_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
