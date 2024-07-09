// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1/role_binding.proto
// DO NOT EDIT!!!

package role_binding

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
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1/condition"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1/role"
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
	_ = &condition.Condition{}
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &role.Role{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

type RoleBinding_FieldMask struct {
	Paths []RoleBinding_FieldPath
}

func FullRoleBinding_FieldMask() *RoleBinding_FieldMask {
	res := &RoleBinding_FieldMask{}
	res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorName})
	res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorRole})
	res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorOwnedObjects})
	res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorMember})
	res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorScopeParams})
	res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorExecutableConditions})
	res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorMemberType})
	res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorAncestryPath})
	res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorSpecGeneration})
	res.Paths = append(res.Paths, &RoleBinding_FieldTerminalPath{selector: RoleBinding_FieldPathSelectorHasOwnedObjects})
	return res
}

func (fieldMask *RoleBinding_FieldMask) String() string {
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
func (fieldMask *RoleBinding_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *RoleBinding_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseRoleBinding_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *RoleBinding_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 11)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*RoleBinding_FieldTerminalPath); ok {
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

func (fieldMask *RoleBinding_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseRoleBinding_FieldPath(raw)
	})
}

func (fieldMask *RoleBinding_FieldMask) ProtoMessage() {}

func (fieldMask *RoleBinding_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *RoleBinding_FieldMask) Subtract(other *RoleBinding_FieldMask) *RoleBinding_FieldMask {
	result := &RoleBinding_FieldMask{}
	removedSelectors := make([]bool, 11)
	otherSubMasks := map[RoleBinding_FieldPathSelector]gotenobject.FieldMask{
		RoleBinding_FieldPathSelectorMetadata:             &meta.Meta_FieldMask{},
		RoleBinding_FieldPathSelectorScopeParams:          &role.ScopeParam_FieldMask{},
		RoleBinding_FieldPathSelectorExecutableConditions: &condition.ExecutableCondition_FieldMask{},
		RoleBinding_FieldPathSelectorAncestryPath:         &RoleBinding_Parent_FieldMask{},
	}
	mySubMasks := map[RoleBinding_FieldPathSelector]gotenobject.FieldMask{
		RoleBinding_FieldPathSelectorMetadata:             &meta.Meta_FieldMask{},
		RoleBinding_FieldPathSelectorScopeParams:          &role.ScopeParam_FieldMask{},
		RoleBinding_FieldPathSelectorExecutableConditions: &condition.ExecutableCondition_FieldMask{},
		RoleBinding_FieldPathSelectorAncestryPath:         &RoleBinding_Parent_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *RoleBinding_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *RoleBinding_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*RoleBinding_FieldTerminalPath); ok {
					switch tp.selector {
					case RoleBinding_FieldPathSelectorMetadata:
						mySubMasks[RoleBinding_FieldPathSelectorMetadata] = meta.FullMeta_FieldMask()
					case RoleBinding_FieldPathSelectorScopeParams:
						mySubMasks[RoleBinding_FieldPathSelectorScopeParams] = role.FullScopeParam_FieldMask()
					case RoleBinding_FieldPathSelectorExecutableConditions:
						mySubMasks[RoleBinding_FieldPathSelectorExecutableConditions] = condition.FullExecutableCondition_FieldMask()
					case RoleBinding_FieldPathSelectorAncestryPath:
						mySubMasks[RoleBinding_FieldPathSelectorAncestryPath] = FullRoleBinding_Parent_FieldMask()
					}
				} else if tp, ok := path.(*RoleBinding_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &RoleBinding_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *RoleBinding_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*RoleBinding_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *RoleBinding_FieldMask) FilterInputFields() *RoleBinding_FieldMask {
	result := &RoleBinding_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case RoleBinding_FieldPathSelectorOwnedObjects:
		case RoleBinding_FieldPathSelectorMemberType:
		case RoleBinding_FieldPathSelectorSpecGeneration:
		case RoleBinding_FieldPathSelectorHasOwnedObjects:
		case RoleBinding_FieldPathSelectorMetadata:
			if _, ok := path.(*RoleBinding_FieldTerminalPath); ok {
				for _, subpath := range meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &RoleBinding_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*RoleBinding_FieldSubPath); ok {
				selectedMask := &meta.Meta_FieldMask{
					Paths: []meta.Meta_FieldPath{sub.subPath.(meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &RoleBinding_FieldSubPath{selector: RoleBinding_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *RoleBinding_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *RoleBinding_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]RoleBinding_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseRoleBinding_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask RoleBinding_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *RoleBinding_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RoleBinding_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask RoleBinding_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *RoleBinding_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RoleBinding_FieldMask) AppendPath(path RoleBinding_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *RoleBinding_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(RoleBinding_FieldPath))
}

func (fieldMask *RoleBinding_FieldMask) GetPaths() []RoleBinding_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *RoleBinding_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *RoleBinding_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseRoleBinding_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *RoleBinding_FieldMask) Set(target, source *RoleBinding) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *RoleBinding_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*RoleBinding), source.(*RoleBinding))
}

func (fieldMask *RoleBinding_FieldMask) Project(source *RoleBinding) *RoleBinding {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &RoleBinding{}
	metadataMask := &meta.Meta_FieldMask{}
	wholeMetadataAccepted := false
	scopeParamsMask := &role.ScopeParam_FieldMask{}
	wholeScopeParamsAccepted := false
	executableConditionsMask := &condition.ExecutableCondition_FieldMask{}
	wholeExecutableConditionsAccepted := false
	ancestryPathMask := &RoleBinding_Parent_FieldMask{}
	wholeAncestryPathAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *RoleBinding_FieldTerminalPath:
			switch tp.selector {
			case RoleBinding_FieldPathSelectorName:
				result.Name = source.Name
			case RoleBinding_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case RoleBinding_FieldPathSelectorRole:
				result.Role = source.Role
			case RoleBinding_FieldPathSelectorOwnedObjects:
				result.OwnedObjects = source.OwnedObjects
			case RoleBinding_FieldPathSelectorMember:
				result.Member = source.Member
			case RoleBinding_FieldPathSelectorScopeParams:
				result.ScopeParams = source.ScopeParams
				wholeScopeParamsAccepted = true
			case RoleBinding_FieldPathSelectorExecutableConditions:
				result.ExecutableConditions = source.ExecutableConditions
				wholeExecutableConditionsAccepted = true
			case RoleBinding_FieldPathSelectorMemberType:
				result.MemberType = source.MemberType
			case RoleBinding_FieldPathSelectorAncestryPath:
				result.AncestryPath = source.AncestryPath
				wholeAncestryPathAccepted = true
			case RoleBinding_FieldPathSelectorSpecGeneration:
				result.SpecGeneration = source.SpecGeneration
			case RoleBinding_FieldPathSelectorHasOwnedObjects:
				result.HasOwnedObjects = source.HasOwnedObjects
			}
		case *RoleBinding_FieldSubPath:
			switch tp.selector {
			case RoleBinding_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(meta.Meta_FieldPath))
			case RoleBinding_FieldPathSelectorScopeParams:
				scopeParamsMask.AppendPath(tp.subPath.(role.ScopeParam_FieldPath))
			case RoleBinding_FieldPathSelectorExecutableConditions:
				executableConditionsMask.AppendPath(tp.subPath.(condition.ExecutableCondition_FieldPath))
			case RoleBinding_FieldPathSelectorAncestryPath:
				ancestryPathMask.AppendPath(tp.subPath.(RoleBindingParent_FieldPath))
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	if wholeScopeParamsAccepted == false && len(scopeParamsMask.Paths) > 0 {
		for _, sourceItem := range source.GetScopeParams() {
			result.ScopeParams = append(result.ScopeParams, scopeParamsMask.Project(sourceItem))
		}
	}
	if wholeExecutableConditionsAccepted == false && len(executableConditionsMask.Paths) > 0 {
		for _, sourceItem := range source.GetExecutableConditions() {
			result.ExecutableConditions = append(result.ExecutableConditions, executableConditionsMask.Project(sourceItem))
		}
	}
	if wholeAncestryPathAccepted == false && len(ancestryPathMask.Paths) > 0 {
		for _, sourceItem := range source.GetAncestryPath() {
			result.AncestryPath = append(result.AncestryPath, ancestryPathMask.Project(sourceItem))
		}
	}
	return result
}

func (fieldMask *RoleBinding_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*RoleBinding))
}

func (fieldMask *RoleBinding_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type RoleBinding_Parent_FieldMask struct {
	Paths []RoleBindingParent_FieldPath
}

func FullRoleBinding_Parent_FieldMask() *RoleBinding_Parent_FieldMask {
	res := &RoleBinding_Parent_FieldMask{}
	res.Paths = append(res.Paths, &RoleBindingParent_FieldTerminalPath{selector: RoleBindingParent_FieldPathSelectorParent})
	res.Paths = append(res.Paths, &RoleBindingParent_FieldTerminalPath{selector: RoleBindingParent_FieldPathSelectorMember})
	return res
}

func (fieldMask *RoleBinding_Parent_FieldMask) String() string {
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
func (fieldMask *RoleBinding_Parent_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *RoleBinding_Parent_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseRoleBindingParent_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *RoleBinding_Parent_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*RoleBindingParent_FieldTerminalPath); ok {
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

func (fieldMask *RoleBinding_Parent_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseRoleBindingParent_FieldPath(raw)
	})
}

func (fieldMask *RoleBinding_Parent_FieldMask) ProtoMessage() {}

func (fieldMask *RoleBinding_Parent_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *RoleBinding_Parent_FieldMask) Subtract(other *RoleBinding_Parent_FieldMask) *RoleBinding_Parent_FieldMask {
	result := &RoleBinding_Parent_FieldMask{}
	removedSelectors := make([]bool, 2)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *RoleBindingParent_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			result.Paths = append(result.Paths, path)
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *RoleBinding_Parent_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*RoleBinding_Parent_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *RoleBinding_Parent_FieldMask) FilterInputFields() *RoleBinding_Parent_FieldMask {
	result := &RoleBinding_Parent_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *RoleBinding_Parent_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *RoleBinding_Parent_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]RoleBindingParent_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseRoleBindingParent_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask RoleBinding_Parent_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *RoleBinding_Parent_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RoleBinding_Parent_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask RoleBinding_Parent_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *RoleBinding_Parent_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RoleBinding_Parent_FieldMask) AppendPath(path RoleBindingParent_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *RoleBinding_Parent_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(RoleBindingParent_FieldPath))
}

func (fieldMask *RoleBinding_Parent_FieldMask) GetPaths() []RoleBindingParent_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *RoleBinding_Parent_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *RoleBinding_Parent_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseRoleBindingParent_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *RoleBinding_Parent_FieldMask) Set(target, source *RoleBinding_Parent) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *RoleBinding_Parent_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*RoleBinding_Parent), source.(*RoleBinding_Parent))
}

func (fieldMask *RoleBinding_Parent_FieldMask) Project(source *RoleBinding_Parent) *RoleBinding_Parent {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &RoleBinding_Parent{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *RoleBindingParent_FieldTerminalPath:
			switch tp.selector {
			case RoleBindingParent_FieldPathSelectorParent:
				result.Parent = source.Parent
			case RoleBindingParent_FieldPathSelectorMember:
				result.Member = source.Member
			}
		}
	}
	return result
}

func (fieldMask *RoleBinding_Parent_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*RoleBinding_Parent))
}

func (fieldMask *RoleBinding_Parent_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
