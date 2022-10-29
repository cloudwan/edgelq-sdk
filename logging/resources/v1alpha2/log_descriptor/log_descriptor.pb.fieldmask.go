// Code generated by protoc-gen-goten-object
// File: edgelq/logging/proto/v1alpha2/log_descriptor.proto
// DO NOT EDIT!!!

package log_descriptor

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
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	logging_common "github.com/cloudwan/edgelq-sdk/logging/common/v1alpha2"
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
	_ = &ntt_meta.Meta{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &logging_common.LabelDescriptor{}
)

type LogDescriptor_FieldMask struct {
	Paths []LogDescriptor_FieldPath
}

func FullLogDescriptor_FieldMask() *LogDescriptor_FieldMask {
	res := &LogDescriptor_FieldMask{}
	res.Paths = append(res.Paths, &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorName})
	res.Paths = append(res.Paths, &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorDescription})
	res.Paths = append(res.Paths, &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorLabels})
	res.Paths = append(res.Paths, &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorPromotedLabelKeySets})
	res.Paths = append(res.Paths, &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorMetadata})
	return res
}

func (fieldMask *LogDescriptor_FieldMask) String() string {
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
func (fieldMask *LogDescriptor_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *LogDescriptor_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseLogDescriptor_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *LogDescriptor_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 6)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*LogDescriptor_FieldTerminalPath); ok {
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

func (fieldMask *LogDescriptor_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseLogDescriptor_FieldPath(raw)
	})
}

func (fieldMask *LogDescriptor_FieldMask) ProtoMessage() {}

func (fieldMask *LogDescriptor_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *LogDescriptor_FieldMask) Subtract(other *LogDescriptor_FieldMask) *LogDescriptor_FieldMask {
	result := &LogDescriptor_FieldMask{}
	removedSelectors := make([]bool, 6)
	otherSubMasks := map[LogDescriptor_FieldPathSelector]gotenobject.FieldMask{
		LogDescriptor_FieldPathSelectorLabels:               &logging_common.LabelDescriptor_FieldMask{},
		LogDescriptor_FieldPathSelectorPromotedLabelKeySets: &logging_common.LabelKeySet_FieldMask{},
		LogDescriptor_FieldPathSelectorMetadata:             &ntt_meta.Meta_FieldMask{},
	}
	mySubMasks := map[LogDescriptor_FieldPathSelector]gotenobject.FieldMask{
		LogDescriptor_FieldPathSelectorLabels:               &logging_common.LabelDescriptor_FieldMask{},
		LogDescriptor_FieldPathSelectorPromotedLabelKeySets: &logging_common.LabelKeySet_FieldMask{},
		LogDescriptor_FieldPathSelectorMetadata:             &ntt_meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *LogDescriptor_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *LogDescriptor_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*LogDescriptor_FieldTerminalPath); ok {
					switch tp.selector {
					case LogDescriptor_FieldPathSelectorLabels:
						mySubMasks[LogDescriptor_FieldPathSelectorLabels] = logging_common.FullLabelDescriptor_FieldMask()
					case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
						mySubMasks[LogDescriptor_FieldPathSelectorPromotedLabelKeySets] = logging_common.FullLabelKeySet_FieldMask()
					case LogDescriptor_FieldPathSelectorMetadata:
						mySubMasks[LogDescriptor_FieldPathSelectorMetadata] = ntt_meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*LogDescriptor_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &LogDescriptor_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *LogDescriptor_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*LogDescriptor_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *LogDescriptor_FieldMask) FilterInputFields() *LogDescriptor_FieldMask {
	result := &LogDescriptor_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case LogDescriptor_FieldPathSelectorMetadata:
			if _, ok := path.(*LogDescriptor_FieldTerminalPath); ok {
				for _, subpath := range ntt_meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &LogDescriptor_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*LogDescriptor_FieldSubPath); ok {
				selectedMask := &ntt_meta.Meta_FieldMask{
					Paths: []ntt_meta.Meta_FieldPath{sub.subPath.(ntt_meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &LogDescriptor_FieldSubPath{selector: LogDescriptor_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *LogDescriptor_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *LogDescriptor_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]LogDescriptor_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseLogDescriptor_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask LogDescriptor_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *LogDescriptor_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *LogDescriptor_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask LogDescriptor_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *LogDescriptor_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *LogDescriptor_FieldMask) AppendPath(path LogDescriptor_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *LogDescriptor_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(LogDescriptor_FieldPath))
}

func (fieldMask *LogDescriptor_FieldMask) GetPaths() []LogDescriptor_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *LogDescriptor_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *LogDescriptor_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseLogDescriptor_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *LogDescriptor_FieldMask) Set(target, source *LogDescriptor) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *LogDescriptor_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*LogDescriptor), source.(*LogDescriptor))
}

func (fieldMask *LogDescriptor_FieldMask) Project(source *LogDescriptor) *LogDescriptor {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &LogDescriptor{}
	labelsMask := &logging_common.LabelDescriptor_FieldMask{}
	wholeLabelsAccepted := false
	promotedLabelKeySetsMask := &logging_common.LabelKeySet_FieldMask{}
	wholePromotedLabelKeySetsAccepted := false
	metadataMask := &ntt_meta.Meta_FieldMask{}
	wholeMetadataAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *LogDescriptor_FieldTerminalPath:
			switch tp.selector {
			case LogDescriptor_FieldPathSelectorName:
				result.Name = source.Name
			case LogDescriptor_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case LogDescriptor_FieldPathSelectorDescription:
				result.Description = source.Description
			case LogDescriptor_FieldPathSelectorLabels:
				result.Labels = source.Labels
				wholeLabelsAccepted = true
			case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
				result.PromotedLabelKeySets = source.PromotedLabelKeySets
				wholePromotedLabelKeySetsAccepted = true
			case LogDescriptor_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			}
		case *LogDescriptor_FieldSubPath:
			switch tp.selector {
			case LogDescriptor_FieldPathSelectorLabels:
				labelsMask.AppendPath(tp.subPath.(logging_common.LabelDescriptor_FieldPath))
			case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
				promotedLabelKeySetsMask.AppendPath(tp.subPath.(logging_common.LabelKeySet_FieldPath))
			case LogDescriptor_FieldPathSelectorMetadata:
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

func (fieldMask *LogDescriptor_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*LogDescriptor))
}

func (fieldMask *LogDescriptor_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
