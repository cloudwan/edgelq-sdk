// Code generated by protoc-gen-goten-object
// File: edgelq/monitoring/proto/v3/recovery_store_sharding_info.proto
// DO NOT EDIT!!!

package recovery_store_sharding_info

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
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	_ = &durationpb.Duration{}
	_ = &timestamppb.Timestamp{}
	_ = &meta.Meta{}
)

type RecoveryStoreShardingInfo_FieldMask struct {
	Paths []RecoveryStoreShardingInfo_FieldPath
}

func FullRecoveryStoreShardingInfo_FieldMask() *RecoveryStoreShardingInfo_FieldMask {
	res := &RecoveryStoreShardingInfo_FieldMask{}
	res.Paths = append(res.Paths, &RecoveryStoreShardingInfo_FieldTerminalPath{selector: RecoveryStoreShardingInfo_FieldPathSelectorName})
	res.Paths = append(res.Paths, &RecoveryStoreShardingInfo_FieldTerminalPath{selector: RecoveryStoreShardingInfo_FieldPathSelectorValidityPeriod})
	res.Paths = append(res.Paths, &RecoveryStoreShardingInfo_FieldTerminalPath{selector: RecoveryStoreShardingInfo_FieldPathSelectorSpec})
	res.Paths = append(res.Paths, &RecoveryStoreShardingInfo_FieldTerminalPath{selector: RecoveryStoreShardingInfo_FieldPathSelectorMetadata})
	return res
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) String() string {
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
func (fieldMask *RecoveryStoreShardingInfo_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseRecoveryStoreShardingInfo_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 4)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*RecoveryStoreShardingInfo_FieldTerminalPath); ok {
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

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseRecoveryStoreShardingInfo_FieldPath(raw)
	})
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) ProtoMessage() {}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) Subtract(other *RecoveryStoreShardingInfo_FieldMask) *RecoveryStoreShardingInfo_FieldMask {
	result := &RecoveryStoreShardingInfo_FieldMask{}
	removedSelectors := make([]bool, 4)
	otherSubMasks := map[RecoveryStoreShardingInfo_FieldPathSelector]gotenobject.FieldMask{
		RecoveryStoreShardingInfo_FieldPathSelectorValidityPeriod: &RecoveryStoreShardingInfo_ValidityPeriod_FieldMask{},
		RecoveryStoreShardingInfo_FieldPathSelectorSpec:           &RecoveryStoreShardingInfo_ShardingSpec_FieldMask{},
		RecoveryStoreShardingInfo_FieldPathSelectorMetadata:       &meta.Meta_FieldMask{},
	}
	mySubMasks := map[RecoveryStoreShardingInfo_FieldPathSelector]gotenobject.FieldMask{
		RecoveryStoreShardingInfo_FieldPathSelectorValidityPeriod: &RecoveryStoreShardingInfo_ValidityPeriod_FieldMask{},
		RecoveryStoreShardingInfo_FieldPathSelectorSpec:           &RecoveryStoreShardingInfo_ShardingSpec_FieldMask{},
		RecoveryStoreShardingInfo_FieldPathSelectorMetadata:       &meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *RecoveryStoreShardingInfo_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *RecoveryStoreShardingInfo_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*RecoveryStoreShardingInfo_FieldTerminalPath); ok {
					switch tp.selector {
					case RecoveryStoreShardingInfo_FieldPathSelectorValidityPeriod:
						mySubMasks[RecoveryStoreShardingInfo_FieldPathSelectorValidityPeriod] = FullRecoveryStoreShardingInfo_ValidityPeriod_FieldMask()
					case RecoveryStoreShardingInfo_FieldPathSelectorSpec:
						mySubMasks[RecoveryStoreShardingInfo_FieldPathSelectorSpec] = FullRecoveryStoreShardingInfo_ShardingSpec_FieldMask()
					case RecoveryStoreShardingInfo_FieldPathSelectorMetadata:
						mySubMasks[RecoveryStoreShardingInfo_FieldPathSelectorMetadata] = meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*RecoveryStoreShardingInfo_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &RecoveryStoreShardingInfo_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*RecoveryStoreShardingInfo_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *RecoveryStoreShardingInfo_FieldMask) FilterInputFields() *RecoveryStoreShardingInfo_FieldMask {
	result := &RecoveryStoreShardingInfo_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case RecoveryStoreShardingInfo_FieldPathSelectorMetadata:
			if _, ok := path.(*RecoveryStoreShardingInfo_FieldTerminalPath); ok {
				for _, subpath := range meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &RecoveryStoreShardingInfo_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*RecoveryStoreShardingInfo_FieldSubPath); ok {
				selectedMask := &meta.Meta_FieldMask{
					Paths: []meta.Meta_FieldPath{sub.subPath.(meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &RecoveryStoreShardingInfo_FieldSubPath{selector: RecoveryStoreShardingInfo_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *RecoveryStoreShardingInfo_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]RecoveryStoreShardingInfo_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseRecoveryStoreShardingInfo_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask RecoveryStoreShardingInfo_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask RecoveryStoreShardingInfo_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) AppendPath(path RecoveryStoreShardingInfo_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(RecoveryStoreShardingInfo_FieldPath))
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) GetPaths() []RecoveryStoreShardingInfo_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseRecoveryStoreShardingInfo_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) Set(target, source *RecoveryStoreShardingInfo) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*RecoveryStoreShardingInfo), source.(*RecoveryStoreShardingInfo))
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) Project(source *RecoveryStoreShardingInfo) *RecoveryStoreShardingInfo {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &RecoveryStoreShardingInfo{}
	validityPeriodMask := &RecoveryStoreShardingInfo_ValidityPeriod_FieldMask{}
	wholeValidityPeriodAccepted := false
	specMask := &RecoveryStoreShardingInfo_ShardingSpec_FieldMask{}
	wholeSpecAccepted := false
	metadataMask := &meta.Meta_FieldMask{}
	wholeMetadataAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *RecoveryStoreShardingInfo_FieldTerminalPath:
			switch tp.selector {
			case RecoveryStoreShardingInfo_FieldPathSelectorName:
				result.Name = source.Name
			case RecoveryStoreShardingInfo_FieldPathSelectorValidityPeriod:
				result.ValidityPeriod = source.ValidityPeriod
				wholeValidityPeriodAccepted = true
			case RecoveryStoreShardingInfo_FieldPathSelectorSpec:
				result.Spec = source.Spec
				wholeSpecAccepted = true
			case RecoveryStoreShardingInfo_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			}
		case *RecoveryStoreShardingInfo_FieldSubPath:
			switch tp.selector {
			case RecoveryStoreShardingInfo_FieldPathSelectorValidityPeriod:
				validityPeriodMask.AppendPath(tp.subPath.(RecoveryStoreShardingInfoValidityPeriod_FieldPath))
			case RecoveryStoreShardingInfo_FieldPathSelectorSpec:
				specMask.AppendPath(tp.subPath.(RecoveryStoreShardingInfoShardingSpec_FieldPath))
			case RecoveryStoreShardingInfo_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(meta.Meta_FieldPath))
			}
		}
	}
	if wholeValidityPeriodAccepted == false && len(validityPeriodMask.Paths) > 0 {
		result.ValidityPeriod = validityPeriodMask.Project(source.GetValidityPeriod())
	}
	if wholeSpecAccepted == false && len(specMask.Paths) > 0 {
		result.Spec = specMask.Project(source.GetSpec())
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	return result
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*RecoveryStoreShardingInfo))
}

func (fieldMask *RecoveryStoreShardingInfo_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type RecoveryStoreShardingInfo_ValidityPeriod_FieldMask struct {
	Paths []RecoveryStoreShardingInfoValidityPeriod_FieldPath
}

func FullRecoveryStoreShardingInfo_ValidityPeriod_FieldMask() *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask {
	res := &RecoveryStoreShardingInfo_ValidityPeriod_FieldMask{}
	res.Paths = append(res.Paths, &RecoveryStoreShardingInfoValidityPeriod_FieldTerminalPath{selector: RecoveryStoreShardingInfoValidityPeriod_FieldPathSelectorStartTime})
	res.Paths = append(res.Paths, &RecoveryStoreShardingInfoValidityPeriod_FieldTerminalPath{selector: RecoveryStoreShardingInfoValidityPeriod_FieldPathSelectorEndTime})
	return res
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) String() string {
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
func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseRecoveryStoreShardingInfoValidityPeriod_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*RecoveryStoreShardingInfoValidityPeriod_FieldTerminalPath); ok {
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

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseRecoveryStoreShardingInfoValidityPeriod_FieldPath(raw)
	})
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) ProtoMessage() {}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) Subtract(other *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask {
	result := &RecoveryStoreShardingInfo_ValidityPeriod_FieldMask{}
	removedSelectors := make([]bool, 2)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *RecoveryStoreShardingInfoValidityPeriod_FieldTerminalPath:
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

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*RecoveryStoreShardingInfo_ValidityPeriod_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) FilterInputFields() *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask {
	result := &RecoveryStoreShardingInfo_ValidityPeriod_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]RecoveryStoreShardingInfoValidityPeriod_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseRecoveryStoreShardingInfoValidityPeriod_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) AppendPath(path RecoveryStoreShardingInfoValidityPeriod_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(RecoveryStoreShardingInfoValidityPeriod_FieldPath))
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) GetPaths() []RecoveryStoreShardingInfoValidityPeriod_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseRecoveryStoreShardingInfoValidityPeriod_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) Set(target, source *RecoveryStoreShardingInfo_ValidityPeriod) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*RecoveryStoreShardingInfo_ValidityPeriod), source.(*RecoveryStoreShardingInfo_ValidityPeriod))
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) Project(source *RecoveryStoreShardingInfo_ValidityPeriod) *RecoveryStoreShardingInfo_ValidityPeriod {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &RecoveryStoreShardingInfo_ValidityPeriod{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *RecoveryStoreShardingInfoValidityPeriod_FieldTerminalPath:
			switch tp.selector {
			case RecoveryStoreShardingInfoValidityPeriod_FieldPathSelectorStartTime:
				result.StartTime = source.StartTime
			case RecoveryStoreShardingInfoValidityPeriod_FieldPathSelectorEndTime:
				result.EndTime = source.EndTime
			}
		}
	}
	return result
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*RecoveryStoreShardingInfo_ValidityPeriod))
}

func (fieldMask *RecoveryStoreShardingInfo_ValidityPeriod_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type RecoveryStoreShardingInfo_ShardingSpec_FieldMask struct {
	Paths []RecoveryStoreShardingInfoShardingSpec_FieldPath
}

func FullRecoveryStoreShardingInfo_ShardingSpec_FieldMask() *RecoveryStoreShardingInfo_ShardingSpec_FieldMask {
	res := &RecoveryStoreShardingInfo_ShardingSpec_FieldMask{}
	res.Paths = append(res.Paths, &RecoveryStoreShardingInfoShardingSpec_FieldTerminalPath{selector: RecoveryStoreShardingInfoShardingSpec_FieldPathSelectorTsBlobPeriod})
	res.Paths = append(res.Paths, &RecoveryStoreShardingInfoShardingSpec_FieldTerminalPath{selector: RecoveryStoreShardingInfoShardingSpec_FieldPathSelectorShardsCount})
	return res
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) String() string {
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
func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseRecoveryStoreShardingInfoShardingSpec_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*RecoveryStoreShardingInfoShardingSpec_FieldTerminalPath); ok {
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

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseRecoveryStoreShardingInfoShardingSpec_FieldPath(raw)
	})
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) ProtoMessage() {}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) Subtract(other *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) *RecoveryStoreShardingInfo_ShardingSpec_FieldMask {
	result := &RecoveryStoreShardingInfo_ShardingSpec_FieldMask{}
	removedSelectors := make([]bool, 2)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *RecoveryStoreShardingInfoShardingSpec_FieldTerminalPath:
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

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*RecoveryStoreShardingInfo_ShardingSpec_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) FilterInputFields() *RecoveryStoreShardingInfo_ShardingSpec_FieldMask {
	result := &RecoveryStoreShardingInfo_ShardingSpec_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]RecoveryStoreShardingInfoShardingSpec_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseRecoveryStoreShardingInfoShardingSpec_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask RecoveryStoreShardingInfo_ShardingSpec_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask RecoveryStoreShardingInfo_ShardingSpec_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) AppendPath(path RecoveryStoreShardingInfoShardingSpec_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(RecoveryStoreShardingInfoShardingSpec_FieldPath))
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) GetPaths() []RecoveryStoreShardingInfoShardingSpec_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseRecoveryStoreShardingInfoShardingSpec_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) Set(target, source *RecoveryStoreShardingInfo_ShardingSpec) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*RecoveryStoreShardingInfo_ShardingSpec), source.(*RecoveryStoreShardingInfo_ShardingSpec))
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) Project(source *RecoveryStoreShardingInfo_ShardingSpec) *RecoveryStoreShardingInfo_ShardingSpec {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &RecoveryStoreShardingInfo_ShardingSpec{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *RecoveryStoreShardingInfoShardingSpec_FieldTerminalPath:
			switch tp.selector {
			case RecoveryStoreShardingInfoShardingSpec_FieldPathSelectorTsBlobPeriod:
				result.TsBlobPeriod = source.TsBlobPeriod
			case RecoveryStoreShardingInfoShardingSpec_FieldPathSelectorShardsCount:
				result.ShardsCount = source.ShardsCount
			}
		}
	}
	return result
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*RecoveryStoreShardingInfo_ShardingSpec))
}

func (fieldMask *RecoveryStoreShardingInfo_ShardingSpec_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
