// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha2/condition.proto
// DO NOT EDIT!!!

package condition

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
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	structpb "google.golang.org/protobuf/types/known/structpb"
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
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &structpb.Struct{}
	_ = &meta.Meta{}
)

type Condition_FieldMask struct {
	Paths []Condition_FieldPath
}

func FullCondition_FieldMask() *Condition_FieldMask {
	res := &Condition_FieldMask{}
	res.Paths = append(res.Paths, &Condition_FieldTerminalPath{selector: Condition_FieldPathSelectorName})
	res.Paths = append(res.Paths, &Condition_FieldTerminalPath{selector: Condition_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &Condition_FieldTerminalPath{selector: Condition_FieldPathSelectorDescription})
	res.Paths = append(res.Paths, &Condition_FieldTerminalPath{selector: Condition_FieldPathSelectorExpression})
	res.Paths = append(res.Paths, &Condition_FieldTerminalPath{selector: Condition_FieldPathSelectorParameterDeclarations})
	res.Paths = append(res.Paths, &Condition_FieldTerminalPath{selector: Condition_FieldPathSelectorMetadata})
	return res
}

func (fieldMask *Condition_FieldMask) String() string {
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
func (fieldMask *Condition_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *Condition_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseCondition_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Condition_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 6)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*Condition_FieldTerminalPath); ok {
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

func (fieldMask *Condition_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseCondition_FieldPath(raw)
	})
}

func (fieldMask *Condition_FieldMask) ProtoMessage() {}

func (fieldMask *Condition_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Condition_FieldMask) Subtract(other *Condition_FieldMask) *Condition_FieldMask {
	result := &Condition_FieldMask{}
	removedSelectors := make([]bool, 6)
	otherSubMasks := map[Condition_FieldPathSelector]gotenobject.FieldMask{
		Condition_FieldPathSelectorParameterDeclarations: &Condition_ParameterDeclaration_FieldMask{},
		Condition_FieldPathSelectorMetadata:              &meta.Meta_FieldMask{},
	}
	mySubMasks := map[Condition_FieldPathSelector]gotenobject.FieldMask{
		Condition_FieldPathSelectorParameterDeclarations: &Condition_ParameterDeclaration_FieldMask{},
		Condition_FieldPathSelectorMetadata:              &meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *Condition_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *Condition_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*Condition_FieldTerminalPath); ok {
					switch tp.selector {
					case Condition_FieldPathSelectorParameterDeclarations:
						mySubMasks[Condition_FieldPathSelectorParameterDeclarations] = FullCondition_ParameterDeclaration_FieldMask()
					case Condition_FieldPathSelectorMetadata:
						mySubMasks[Condition_FieldPathSelectorMetadata] = meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*Condition_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &Condition_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *Condition_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Condition_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Condition_FieldMask) FilterInputFields() *Condition_FieldMask {
	result := &Condition_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case Condition_FieldPathSelectorMetadata:
			if _, ok := path.(*Condition_FieldTerminalPath); ok {
				for _, subpath := range meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &Condition_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*Condition_FieldSubPath); ok {
				selectedMask := &meta.Meta_FieldMask{
					Paths: []meta.Meta_FieldPath{sub.subPath.(meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &Condition_FieldSubPath{selector: Condition_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Condition_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Condition_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]Condition_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseCondition_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Condition_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Condition_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Condition_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Condition_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Condition_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Condition_FieldMask) AppendPath(path Condition_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Condition_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(Condition_FieldPath))
}

func (fieldMask *Condition_FieldMask) GetPaths() []Condition_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Condition_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Condition_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseCondition_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Condition_FieldMask) Set(target, source *Condition) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Condition_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Condition), source.(*Condition))
}

func (fieldMask *Condition_FieldMask) Project(source *Condition) *Condition {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Condition{}
	parameterDeclarationsMask := &Condition_ParameterDeclaration_FieldMask{}
	wholeParameterDeclarationsAccepted := false
	metadataMask := &meta.Meta_FieldMask{}
	wholeMetadataAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *Condition_FieldTerminalPath:
			switch tp.selector {
			case Condition_FieldPathSelectorName:
				result.Name = source.Name
			case Condition_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case Condition_FieldPathSelectorDescription:
				result.Description = source.Description
			case Condition_FieldPathSelectorExpression:
				result.Expression = source.Expression
			case Condition_FieldPathSelectorParameterDeclarations:
				result.ParameterDeclarations = source.ParameterDeclarations
				wholeParameterDeclarationsAccepted = true
			case Condition_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			}
		case *Condition_FieldSubPath:
			switch tp.selector {
			case Condition_FieldPathSelectorParameterDeclarations:
				parameterDeclarationsMask.AppendPath(tp.subPath.(ConditionParameterDeclaration_FieldPath))
			case Condition_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(meta.Meta_FieldPath))
			}
		}
	}
	if wholeParameterDeclarationsAccepted == false && len(parameterDeclarationsMask.Paths) > 0 {
		for _, sourceItem := range source.GetParameterDeclarations() {
			result.ParameterDeclarations = append(result.ParameterDeclarations, parameterDeclarationsMask.Project(sourceItem))
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	return result
}

func (fieldMask *Condition_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Condition))
}

func (fieldMask *Condition_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type Condition_ParameterDeclaration_FieldMask struct {
	Paths []ConditionParameterDeclaration_FieldPath
}

func FullCondition_ParameterDeclaration_FieldMask() *Condition_ParameterDeclaration_FieldMask {
	res := &Condition_ParameterDeclaration_FieldMask{}
	res.Paths = append(res.Paths, &ConditionParameterDeclaration_FieldTerminalPath{selector: ConditionParameterDeclaration_FieldPathSelectorKey})
	res.Paths = append(res.Paths, &ConditionParameterDeclaration_FieldTerminalPath{selector: ConditionParameterDeclaration_FieldPathSelectorType})
	return res
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) String() string {
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
func (fieldMask *Condition_ParameterDeclaration_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *Condition_ParameterDeclaration_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseConditionParameterDeclaration_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ConditionParameterDeclaration_FieldTerminalPath); ok {
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

func (fieldMask *Condition_ParameterDeclaration_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseConditionParameterDeclaration_FieldPath(raw)
	})
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) ProtoMessage() {}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) Subtract(other *Condition_ParameterDeclaration_FieldMask) *Condition_ParameterDeclaration_FieldMask {
	result := &Condition_ParameterDeclaration_FieldMask{}
	removedSelectors := make([]bool, 2)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ConditionParameterDeclaration_FieldTerminalPath:
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

func (fieldMask *Condition_ParameterDeclaration_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Condition_ParameterDeclaration_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Condition_ParameterDeclaration_FieldMask) FilterInputFields() *Condition_ParameterDeclaration_FieldMask {
	result := &Condition_ParameterDeclaration_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Condition_ParameterDeclaration_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ConditionParameterDeclaration_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseConditionParameterDeclaration_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Condition_ParameterDeclaration_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Condition_ParameterDeclaration_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) AppendPath(path ConditionParameterDeclaration_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ConditionParameterDeclaration_FieldPath))
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) GetPaths() []ConditionParameterDeclaration_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseConditionParameterDeclaration_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) Set(target, source *Condition_ParameterDeclaration) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Condition_ParameterDeclaration), source.(*Condition_ParameterDeclaration))
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) Project(source *Condition_ParameterDeclaration) *Condition_ParameterDeclaration {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Condition_ParameterDeclaration{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ConditionParameterDeclaration_FieldTerminalPath:
			switch tp.selector {
			case ConditionParameterDeclaration_FieldPathSelectorKey:
				result.Key = source.Key
			case ConditionParameterDeclaration_FieldPathSelectorType:
				result.Type = source.Type
			}
		}
	}
	return result
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Condition_ParameterDeclaration))
}

func (fieldMask *Condition_ParameterDeclaration_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type ConditionBinding_FieldMask struct {
	Paths []ConditionBinding_FieldPath
}

func FullConditionBinding_FieldMask() *ConditionBinding_FieldMask {
	res := &ConditionBinding_FieldMask{}
	res.Paths = append(res.Paths, &ConditionBinding_FieldTerminalPath{selector: ConditionBinding_FieldPathSelectorCondition})
	res.Paths = append(res.Paths, &ConditionBinding_FieldTerminalPath{selector: ConditionBinding_FieldPathSelectorParameters})
	res.Paths = append(res.Paths, &ConditionBinding_FieldTerminalPath{selector: ConditionBinding_FieldPathSelectorParams})
	return res
}

func (fieldMask *ConditionBinding_FieldMask) String() string {
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
func (fieldMask *ConditionBinding_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ConditionBinding_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseConditionBinding_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ConditionBinding_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 3)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ConditionBinding_FieldTerminalPath); ok {
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

func (fieldMask *ConditionBinding_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseConditionBinding_FieldPath(raw)
	})
}

func (fieldMask *ConditionBinding_FieldMask) ProtoMessage() {}

func (fieldMask *ConditionBinding_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ConditionBinding_FieldMask) Subtract(other *ConditionBinding_FieldMask) *ConditionBinding_FieldMask {
	result := &ConditionBinding_FieldMask{}
	removedSelectors := make([]bool, 3)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ConditionBinding_FieldTerminalPath:
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

func (fieldMask *ConditionBinding_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ConditionBinding_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ConditionBinding_FieldMask) FilterInputFields() *ConditionBinding_FieldMask {
	result := &ConditionBinding_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ConditionBinding_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ConditionBinding_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ConditionBinding_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseConditionBinding_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ConditionBinding_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ConditionBinding_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ConditionBinding_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ConditionBinding_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ConditionBinding_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ConditionBinding_FieldMask) AppendPath(path ConditionBinding_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ConditionBinding_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ConditionBinding_FieldPath))
}

func (fieldMask *ConditionBinding_FieldMask) GetPaths() []ConditionBinding_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ConditionBinding_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ConditionBinding_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseConditionBinding_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ConditionBinding_FieldMask) Set(target, source *ConditionBinding) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ConditionBinding_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ConditionBinding), source.(*ConditionBinding))
}

func (fieldMask *ConditionBinding_FieldMask) Project(source *ConditionBinding) *ConditionBinding {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ConditionBinding{}
	var parametersMapKeys []string
	wholeParametersAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ConditionBinding_FieldTerminalPath:
			switch tp.selector {
			case ConditionBinding_FieldPathSelectorCondition:
				result.Condition = source.Condition
			case ConditionBinding_FieldPathSelectorParameters:
				result.Parameters = source.Parameters
				wholeParametersAccepted = true
			case ConditionBinding_FieldPathSelectorParams:
				result.Params = source.Params
			}
		case *ConditionBinding_FieldPathMap:
			switch tp.selector {
			case ConditionBinding_FieldPathSelectorParameters:
				parametersMapKeys = append(parametersMapKeys, tp.key)
			}
		}
	}
	if wholeParametersAccepted == false && len(parametersMapKeys) > 0 && source.GetParameters() != nil {
		copiedMap := map[string]string{}
		sourceMap := source.GetParameters()
		for _, key := range parametersMapKeys {
			copiedMap[key] = sourceMap[key]
		}
		result.Parameters = copiedMap
	}
	return result
}

func (fieldMask *ConditionBinding_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ConditionBinding))
}

func (fieldMask *ConditionBinding_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
