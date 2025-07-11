// Code generated by protoc-gen-goten-object
// File: edgelq/alerting/proto/v1/log_condition_template.proto
// DO NOT EDIT!!!

package log_condition_template

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
	rcommon "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/common"
	document "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/document"
	policy_template "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/policy_template"
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
	_ = &document.Document{}
	_ = &policy_template.PolicyTemplate{}
	_ = &rcommon.LogCndSpec{}
	_ = &meta.Meta{}
)

type LogConditionTemplate_FieldMask struct {
	Paths []LogConditionTemplate_FieldPath
}

func FullLogConditionTemplate_FieldMask() *LogConditionTemplate_FieldMask {
	res := &LogConditionTemplate_FieldMask{}
	res.Paths = append(res.Paths, &LogConditionTemplate_FieldTerminalPath{selector: LogConditionTemplate_FieldPathSelectorName})
	res.Paths = append(res.Paths, &LogConditionTemplate_FieldTerminalPath{selector: LogConditionTemplate_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &LogConditionTemplate_FieldTerminalPath{selector: LogConditionTemplate_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &LogConditionTemplate_FieldTerminalPath{selector: LogConditionTemplate_FieldPathSelectorDescription})
	res.Paths = append(res.Paths, &LogConditionTemplate_FieldTerminalPath{selector: LogConditionTemplate_FieldPathSelectorSupportingDocs})
	res.Paths = append(res.Paths, &LogConditionTemplate_FieldTerminalPath{selector: LogConditionTemplate_FieldPathSelectorSpecTemplate})
	return res
}

func (fieldMask *LogConditionTemplate_FieldMask) String() string {
	if fieldMask == nil {
		return "<nil>"
	}
	pathsStr := make([]string, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		pathsStr = append(pathsStr, path.String())
	}
	return strings.Join(pathsStr, ", ")
}

func (fieldMask *LogConditionTemplate_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 6)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*LogConditionTemplate_FieldTerminalPath); ok {
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

func (fieldMask *LogConditionTemplate_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseLogConditionTemplate_FieldPath(raw)
	})
}

func (fieldMask *LogConditionTemplate_FieldMask) ProtoMessage() {}

func (fieldMask *LogConditionTemplate_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *LogConditionTemplate_FieldMask) Subtract(other *LogConditionTemplate_FieldMask) *LogConditionTemplate_FieldMask {
	result := &LogConditionTemplate_FieldMask{}
	removedSelectors := make([]bool, 6)
	otherSubMasks := map[LogConditionTemplate_FieldPathSelector]gotenobject.FieldMask{
		LogConditionTemplate_FieldPathSelectorMetadata:     &meta.Meta_FieldMask{},
		LogConditionTemplate_FieldPathSelectorSpecTemplate: &rcommon.LogCndSpec_FieldMask{},
	}
	mySubMasks := map[LogConditionTemplate_FieldPathSelector]gotenobject.FieldMask{
		LogConditionTemplate_FieldPathSelectorMetadata:     &meta.Meta_FieldMask{},
		LogConditionTemplate_FieldPathSelectorSpecTemplate: &rcommon.LogCndSpec_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *LogConditionTemplate_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *LogConditionTemplate_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*LogConditionTemplate_FieldTerminalPath); ok {
					switch tp.selector {
					case LogConditionTemplate_FieldPathSelectorMetadata:
						mySubMasks[LogConditionTemplate_FieldPathSelectorMetadata] = meta.FullMeta_FieldMask()
					case LogConditionTemplate_FieldPathSelectorSpecTemplate:
						mySubMasks[LogConditionTemplate_FieldPathSelectorSpecTemplate] = rcommon.FullLogCndSpec_FieldMask()
					}
				} else if tp, ok := path.(*LogConditionTemplate_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &LogConditionTemplate_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *LogConditionTemplate_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*LogConditionTemplate_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *LogConditionTemplate_FieldMask) FilterInputFields() *LogConditionTemplate_FieldMask {
	result := &LogConditionTemplate_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case LogConditionTemplate_FieldPathSelectorMetadata:
			if _, ok := path.(*LogConditionTemplate_FieldTerminalPath); ok {
				for _, subpath := range meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &LogConditionTemplate_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*LogConditionTemplate_FieldSubPath); ok {
				selectedMask := &meta.Meta_FieldMask{
					Paths: []meta.Meta_FieldPath{sub.subPath.(meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &LogConditionTemplate_FieldSubPath{selector: LogConditionTemplate_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *LogConditionTemplate_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *LogConditionTemplate_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]LogConditionTemplate_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseLogConditionTemplate_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask LogConditionTemplate_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *LogConditionTemplate_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *LogConditionTemplate_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask LogConditionTemplate_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *LogConditionTemplate_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *LogConditionTemplate_FieldMask) AppendPath(path LogConditionTemplate_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *LogConditionTemplate_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(LogConditionTemplate_FieldPath))
}

func (fieldMask *LogConditionTemplate_FieldMask) GetPaths() []LogConditionTemplate_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *LogConditionTemplate_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *LogConditionTemplate_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseLogConditionTemplate_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *LogConditionTemplate_FieldMask) Set(target, source *LogConditionTemplate) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *LogConditionTemplate_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*LogConditionTemplate), source.(*LogConditionTemplate))
}

func (fieldMask *LogConditionTemplate_FieldMask) Project(source *LogConditionTemplate) *LogConditionTemplate {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &LogConditionTemplate{}
	metadataMask := &meta.Meta_FieldMask{}
	wholeMetadataAccepted := false
	specTemplateMask := &rcommon.LogCndSpec_FieldMask{}
	wholeSpecTemplateAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *LogConditionTemplate_FieldTerminalPath:
			switch tp.selector {
			case LogConditionTemplate_FieldPathSelectorName:
				result.Name = source.Name
			case LogConditionTemplate_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case LogConditionTemplate_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case LogConditionTemplate_FieldPathSelectorDescription:
				result.Description = source.Description
			case LogConditionTemplate_FieldPathSelectorSupportingDocs:
				result.SupportingDocs = source.SupportingDocs
			case LogConditionTemplate_FieldPathSelectorSpecTemplate:
				result.SpecTemplate = source.SpecTemplate
				wholeSpecTemplateAccepted = true
			}
		case *LogConditionTemplate_FieldSubPath:
			switch tp.selector {
			case LogConditionTemplate_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(meta.Meta_FieldPath))
			case LogConditionTemplate_FieldPathSelectorSpecTemplate:
				specTemplateMask.AppendPath(tp.subPath.(rcommon.LogCndSpec_FieldPath))
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	if wholeSpecTemplateAccepted == false && len(specTemplateMask.Paths) > 0 {
		result.SpecTemplate = specTemplateMask.Project(source.GetSpecTemplate())
	}
	return result
}

func (fieldMask *LogConditionTemplate_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*LogConditionTemplate))
}

func (fieldMask *LogConditionTemplate_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
