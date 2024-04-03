// Code generated by protoc-gen-goten-object
// File: edgelq/limits/proto/v1/plan_assignment.proto
// DO NOT EDIT!!!

package plan_assignment

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
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	accepted_plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1/accepted_plan"
	common "github.com/cloudwan/edgelq-sdk/limits/resources/v1/common"
	plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1/plan"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
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
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &accepted_plan.AcceptedPlan{}
	_ = &common.RegionalPlanAssignment{}
	_ = &plan.Plan{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

type PlanAssignment_FieldMask struct {
	Paths []PlanAssignment_FieldPath
}

func FullPlanAssignment_FieldMask() *PlanAssignment_FieldMask {
	res := &PlanAssignment_FieldMask{}
	res.Paths = append(res.Paths, &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorName})
	res.Paths = append(res.Paths, &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorDefaultRegionalPlan})
	res.Paths = append(res.Paths, &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorService})
	res.Paths = append(res.Paths, &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorRegionalPlanOverrides})
	res.Paths = append(res.Paths, &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorExtensions})
	res.Paths = append(res.Paths, &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorAllowances})
	res.Paths = append(res.Paths, &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorAppliedRegions})
	res.Paths = append(res.Paths, &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorAppliedPlanSpecGeneration})
	res.Paths = append(res.Paths, &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorSource})
	return res
}

func (fieldMask *PlanAssignment_FieldMask) String() string {
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
func (fieldMask *PlanAssignment_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *PlanAssignment_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParsePlanAssignment_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *PlanAssignment_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 10)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*PlanAssignment_FieldTerminalPath); ok {
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

func (fieldMask *PlanAssignment_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParsePlanAssignment_FieldPath(raw)
	})
}

func (fieldMask *PlanAssignment_FieldMask) ProtoMessage() {}

func (fieldMask *PlanAssignment_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *PlanAssignment_FieldMask) Subtract(other *PlanAssignment_FieldMask) *PlanAssignment_FieldMask {
	result := &PlanAssignment_FieldMask{}
	removedSelectors := make([]bool, 10)
	otherSubMasks := map[PlanAssignment_FieldPathSelector]gotenobject.FieldMask{
		PlanAssignment_FieldPathSelectorMetadata:              &meta.Meta_FieldMask{},
		PlanAssignment_FieldPathSelectorRegionalPlanOverrides: &common.RegionalPlanAssignment_FieldMask{},
		PlanAssignment_FieldPathSelectorExtensions:            &common.Allowance_FieldMask{},
		PlanAssignment_FieldPathSelectorAllowances:            &common.Allowance_FieldMask{},
	}
	mySubMasks := map[PlanAssignment_FieldPathSelector]gotenobject.FieldMask{
		PlanAssignment_FieldPathSelectorMetadata:              &meta.Meta_FieldMask{},
		PlanAssignment_FieldPathSelectorRegionalPlanOverrides: &common.RegionalPlanAssignment_FieldMask{},
		PlanAssignment_FieldPathSelectorExtensions:            &common.Allowance_FieldMask{},
		PlanAssignment_FieldPathSelectorAllowances:            &common.Allowance_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *PlanAssignment_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *PlanAssignment_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*PlanAssignment_FieldTerminalPath); ok {
					switch tp.selector {
					case PlanAssignment_FieldPathSelectorMetadata:
						mySubMasks[PlanAssignment_FieldPathSelectorMetadata] = meta.FullMeta_FieldMask()
					case PlanAssignment_FieldPathSelectorRegionalPlanOverrides:
						mySubMasks[PlanAssignment_FieldPathSelectorRegionalPlanOverrides] = common.FullRegionalPlanAssignment_FieldMask()
					case PlanAssignment_FieldPathSelectorExtensions:
						mySubMasks[PlanAssignment_FieldPathSelectorExtensions] = common.FullAllowance_FieldMask()
					case PlanAssignment_FieldPathSelectorAllowances:
						mySubMasks[PlanAssignment_FieldPathSelectorAllowances] = common.FullAllowance_FieldMask()
					}
				} else if tp, ok := path.(*PlanAssignment_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &PlanAssignment_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *PlanAssignment_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*PlanAssignment_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *PlanAssignment_FieldMask) FilterInputFields() *PlanAssignment_FieldMask {
	result := &PlanAssignment_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case PlanAssignment_FieldPathSelectorDefaultRegionalPlan:
		case PlanAssignment_FieldPathSelectorService:
		case PlanAssignment_FieldPathSelectorRegionalPlanOverrides:
		case PlanAssignment_FieldPathSelectorExtensions:
		case PlanAssignment_FieldPathSelectorAllowances:
		case PlanAssignment_FieldPathSelectorAppliedRegions:
		case PlanAssignment_FieldPathSelectorAppliedPlanSpecGeneration:
		case PlanAssignment_FieldPathSelectorMetadata:
			if _, ok := path.(*PlanAssignment_FieldTerminalPath); ok {
				for _, subpath := range meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &PlanAssignment_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*PlanAssignment_FieldSubPath); ok {
				selectedMask := &meta.Meta_FieldMask{
					Paths: []meta.Meta_FieldPath{sub.subPath.(meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &PlanAssignment_FieldSubPath{selector: PlanAssignment_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *PlanAssignment_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *PlanAssignment_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]PlanAssignment_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParsePlanAssignment_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask PlanAssignment_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *PlanAssignment_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *PlanAssignment_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask PlanAssignment_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *PlanAssignment_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *PlanAssignment_FieldMask) AppendPath(path PlanAssignment_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *PlanAssignment_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(PlanAssignment_FieldPath))
}

func (fieldMask *PlanAssignment_FieldMask) GetPaths() []PlanAssignment_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *PlanAssignment_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *PlanAssignment_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParsePlanAssignment_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *PlanAssignment_FieldMask) Set(target, source *PlanAssignment) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *PlanAssignment_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*PlanAssignment), source.(*PlanAssignment))
}

func (fieldMask *PlanAssignment_FieldMask) Project(source *PlanAssignment) *PlanAssignment {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &PlanAssignment{}
	metadataMask := &meta.Meta_FieldMask{}
	wholeMetadataAccepted := false
	regionalPlanOverridesMask := &common.RegionalPlanAssignment_FieldMask{}
	wholeRegionalPlanOverridesAccepted := false
	extensionsMask := &common.Allowance_FieldMask{}
	wholeExtensionsAccepted := false
	allowancesMask := &common.Allowance_FieldMask{}
	wholeAllowancesAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *PlanAssignment_FieldTerminalPath:
			switch tp.selector {
			case PlanAssignment_FieldPathSelectorName:
				result.Name = source.Name
			case PlanAssignment_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case PlanAssignment_FieldPathSelectorDefaultRegionalPlan:
				result.DefaultRegionalPlan = source.DefaultRegionalPlan
			case PlanAssignment_FieldPathSelectorService:
				result.Service = source.Service
			case PlanAssignment_FieldPathSelectorRegionalPlanOverrides:
				result.RegionalPlanOverrides = source.RegionalPlanOverrides
				wholeRegionalPlanOverridesAccepted = true
			case PlanAssignment_FieldPathSelectorExtensions:
				result.Extensions = source.Extensions
				wholeExtensionsAccepted = true
			case PlanAssignment_FieldPathSelectorAllowances:
				result.Allowances = source.Allowances
				wholeAllowancesAccepted = true
			case PlanAssignment_FieldPathSelectorAppliedRegions:
				result.AppliedRegions = source.AppliedRegions
			case PlanAssignment_FieldPathSelectorAppliedPlanSpecGeneration:
				result.AppliedPlanSpecGeneration = source.AppliedPlanSpecGeneration
			case PlanAssignment_FieldPathSelectorSource:
				result.Source = source.Source
			}
		case *PlanAssignment_FieldSubPath:
			switch tp.selector {
			case PlanAssignment_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(meta.Meta_FieldPath))
			case PlanAssignment_FieldPathSelectorRegionalPlanOverrides:
				regionalPlanOverridesMask.AppendPath(tp.subPath.(common.RegionalPlanAssignment_FieldPath))
			case PlanAssignment_FieldPathSelectorExtensions:
				extensionsMask.AppendPath(tp.subPath.(common.Allowance_FieldPath))
			case PlanAssignment_FieldPathSelectorAllowances:
				allowancesMask.AppendPath(tp.subPath.(common.Allowance_FieldPath))
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	if wholeRegionalPlanOverridesAccepted == false && len(regionalPlanOverridesMask.Paths) > 0 {
		for _, sourceItem := range source.GetRegionalPlanOverrides() {
			result.RegionalPlanOverrides = append(result.RegionalPlanOverrides, regionalPlanOverridesMask.Project(sourceItem))
		}
	}
	if wholeExtensionsAccepted == false && len(extensionsMask.Paths) > 0 {
		for _, sourceItem := range source.GetExtensions() {
			result.Extensions = append(result.Extensions, extensionsMask.Project(sourceItem))
		}
	}
	if wholeAllowancesAccepted == false && len(allowancesMask.Paths) > 0 {
		for _, sourceItem := range source.GetAllowances() {
			result.Allowances = append(result.Allowances, allowancesMask.Project(sourceItem))
		}
	}
	return result
}

func (fieldMask *PlanAssignment_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*PlanAssignment))
}

func (fieldMask *PlanAssignment_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}