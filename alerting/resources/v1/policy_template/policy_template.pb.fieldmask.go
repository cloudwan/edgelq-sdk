// Code generated by protoc-gen-goten-object
// File: edgelq/alerting/proto/v1/policy_template.proto
// DO NOT EDIT!!!

package policy_template

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
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
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
	_ = &rcommon.LogCndSpec{}
	_ = &iam_project.Project{}
	_ = &meta.Meta{}
)

type PolicyTemplate_FieldMask struct {
	Paths []PolicyTemplate_FieldPath
}

func FullPolicyTemplate_FieldMask() *PolicyTemplate_FieldMask {
	res := &PolicyTemplate_FieldMask{}
	res.Paths = append(res.Paths, &PolicyTemplate_FieldTerminalPath{selector: PolicyTemplate_FieldPathSelectorName})
	res.Paths = append(res.Paths, &PolicyTemplate_FieldTerminalPath{selector: PolicyTemplate_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &PolicyTemplate_FieldTerminalPath{selector: PolicyTemplate_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &PolicyTemplate_FieldTerminalPath{selector: PolicyTemplate_FieldPathSelectorDescription})
	res.Paths = append(res.Paths, &PolicyTemplate_FieldTerminalPath{selector: PolicyTemplate_FieldPathSelectorSupportingDocs})
	res.Paths = append(res.Paths, &PolicyTemplate_FieldTerminalPath{selector: PolicyTemplate_FieldPathSelectorSpecTemplate})
	return res
}

func (fieldMask *PolicyTemplate_FieldMask) String() string {
	if fieldMask == nil {
		return "<nil>"
	}
	pathsStr := make([]string, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		pathsStr = append(pathsStr, path.String())
	}
	return strings.Join(pathsStr, ", ")
}

func (fieldMask *PolicyTemplate_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 6)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*PolicyTemplate_FieldTerminalPath); ok {
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

func (fieldMask *PolicyTemplate_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParsePolicyTemplate_FieldPath(raw)
	})
}

func (fieldMask *PolicyTemplate_FieldMask) ProtoMessage() {}

func (fieldMask *PolicyTemplate_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *PolicyTemplate_FieldMask) Subtract(other *PolicyTemplate_FieldMask) *PolicyTemplate_FieldMask {
	result := &PolicyTemplate_FieldMask{}
	removedSelectors := make([]bool, 6)
	otherSubMasks := map[PolicyTemplate_FieldPathSelector]gotenobject.FieldMask{
		PolicyTemplate_FieldPathSelectorMetadata:     &meta.Meta_FieldMask{},
		PolicyTemplate_FieldPathSelectorSpecTemplate: &rcommon.PolicySpec_FieldMask{},
	}
	mySubMasks := map[PolicyTemplate_FieldPathSelector]gotenobject.FieldMask{
		PolicyTemplate_FieldPathSelectorMetadata:     &meta.Meta_FieldMask{},
		PolicyTemplate_FieldPathSelectorSpecTemplate: &rcommon.PolicySpec_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *PolicyTemplate_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *PolicyTemplate_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*PolicyTemplate_FieldTerminalPath); ok {
					switch tp.selector {
					case PolicyTemplate_FieldPathSelectorMetadata:
						mySubMasks[PolicyTemplate_FieldPathSelectorMetadata] = meta.FullMeta_FieldMask()
					case PolicyTemplate_FieldPathSelectorSpecTemplate:
						mySubMasks[PolicyTemplate_FieldPathSelectorSpecTemplate] = rcommon.FullPolicySpec_FieldMask()
					}
				} else if tp, ok := path.(*PolicyTemplate_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &PolicyTemplate_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *PolicyTemplate_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*PolicyTemplate_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *PolicyTemplate_FieldMask) FilterInputFields() *PolicyTemplate_FieldMask {
	result := &PolicyTemplate_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case PolicyTemplate_FieldPathSelectorMetadata:
			if _, ok := path.(*PolicyTemplate_FieldTerminalPath); ok {
				for _, subpath := range meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &PolicyTemplate_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*PolicyTemplate_FieldSubPath); ok {
				selectedMask := &meta.Meta_FieldMask{
					Paths: []meta.Meta_FieldPath{sub.subPath.(meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &PolicyTemplate_FieldSubPath{selector: PolicyTemplate_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *PolicyTemplate_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *PolicyTemplate_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]PolicyTemplate_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParsePolicyTemplate_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask PolicyTemplate_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *PolicyTemplate_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *PolicyTemplate_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask PolicyTemplate_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *PolicyTemplate_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *PolicyTemplate_FieldMask) AppendPath(path PolicyTemplate_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *PolicyTemplate_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(PolicyTemplate_FieldPath))
}

func (fieldMask *PolicyTemplate_FieldMask) GetPaths() []PolicyTemplate_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *PolicyTemplate_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *PolicyTemplate_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParsePolicyTemplate_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *PolicyTemplate_FieldMask) Set(target, source *PolicyTemplate) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *PolicyTemplate_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*PolicyTemplate), source.(*PolicyTemplate))
}

func (fieldMask *PolicyTemplate_FieldMask) Project(source *PolicyTemplate) *PolicyTemplate {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &PolicyTemplate{}
	metadataMask := &meta.Meta_FieldMask{}
	wholeMetadataAccepted := false
	specTemplateMask := &rcommon.PolicySpec_FieldMask{}
	wholeSpecTemplateAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *PolicyTemplate_FieldTerminalPath:
			switch tp.selector {
			case PolicyTemplate_FieldPathSelectorName:
				result.Name = source.Name
			case PolicyTemplate_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case PolicyTemplate_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case PolicyTemplate_FieldPathSelectorDescription:
				result.Description = source.Description
			case PolicyTemplate_FieldPathSelectorSupportingDocs:
				result.SupportingDocs = source.SupportingDocs
			case PolicyTemplate_FieldPathSelectorSpecTemplate:
				result.SpecTemplate = source.SpecTemplate
				wholeSpecTemplateAccepted = true
			}
		case *PolicyTemplate_FieldSubPath:
			switch tp.selector {
			case PolicyTemplate_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(meta.Meta_FieldPath))
			case PolicyTemplate_FieldPathSelectorSpecTemplate:
				specTemplateMask.AppendPath(tp.subPath.(rcommon.PolicySpec_FieldPath))
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

func (fieldMask *PolicyTemplate_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*PolicyTemplate))
}

func (fieldMask *PolicyTemplate_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
