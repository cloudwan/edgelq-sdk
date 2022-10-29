// Code generated by protoc-gen-goten-object
// File: edgelq/devices/proto/v1alpha/provisioning_approval_request.proto
// DO NOT EDIT!!!

package provisioning_approval_request

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
	provisioning_policy "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha/provisioning_policy"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/service_account"
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
	_ = &provisioning_policy.ProvisioningPolicy{}
	_ = &iam_service_account.ServiceAccount{}
)

type ProvisioningApprovalRequest_FieldMask struct {
	Paths []ProvisioningApprovalRequest_FieldPath
}

func FullProvisioningApprovalRequest_FieldMask() *ProvisioningApprovalRequest_FieldMask {
	res := &ProvisioningApprovalRequest_FieldMask{}
	res.Paths = append(res.Paths, &ProvisioningApprovalRequest_FieldTerminalPath{selector: ProvisioningApprovalRequest_FieldPathSelectorName})
	res.Paths = append(res.Paths, &ProvisioningApprovalRequest_FieldTerminalPath{selector: ProvisioningApprovalRequest_FieldPathSelectorSpec})
	res.Paths = append(res.Paths, &ProvisioningApprovalRequest_FieldTerminalPath{selector: ProvisioningApprovalRequest_FieldPathSelectorStatus})
	res.Paths = append(res.Paths, &ProvisioningApprovalRequest_FieldTerminalPath{selector: ProvisioningApprovalRequest_FieldPathSelectorMetadata})
	return res
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) String() string {
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
func (fieldMask *ProvisioningApprovalRequest_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ProvisioningApprovalRequest_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseProvisioningApprovalRequest_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 4)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ProvisioningApprovalRequest_FieldTerminalPath); ok {
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

func (fieldMask *ProvisioningApprovalRequest_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseProvisioningApprovalRequest_FieldPath(raw)
	})
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) ProtoMessage() {}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) Subtract(other *ProvisioningApprovalRequest_FieldMask) *ProvisioningApprovalRequest_FieldMask {
	result := &ProvisioningApprovalRequest_FieldMask{}
	removedSelectors := make([]bool, 4)
	otherSubMasks := map[ProvisioningApprovalRequest_FieldPathSelector]gotenobject.FieldMask{
		ProvisioningApprovalRequest_FieldPathSelectorSpec:     &ProvisioningApprovalRequest_Spec_FieldMask{},
		ProvisioningApprovalRequest_FieldPathSelectorStatus:   &ProvisioningApprovalRequest_Status_FieldMask{},
		ProvisioningApprovalRequest_FieldPathSelectorMetadata: &ntt_meta.Meta_FieldMask{},
	}
	mySubMasks := map[ProvisioningApprovalRequest_FieldPathSelector]gotenobject.FieldMask{
		ProvisioningApprovalRequest_FieldPathSelectorSpec:     &ProvisioningApprovalRequest_Spec_FieldMask{},
		ProvisioningApprovalRequest_FieldPathSelectorStatus:   &ProvisioningApprovalRequest_Status_FieldMask{},
		ProvisioningApprovalRequest_FieldPathSelectorMetadata: &ntt_meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ProvisioningApprovalRequest_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *ProvisioningApprovalRequest_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*ProvisioningApprovalRequest_FieldTerminalPath); ok {
					switch tp.selector {
					case ProvisioningApprovalRequest_FieldPathSelectorSpec:
						mySubMasks[ProvisioningApprovalRequest_FieldPathSelectorSpec] = FullProvisioningApprovalRequest_Spec_FieldMask()
					case ProvisioningApprovalRequest_FieldPathSelectorStatus:
						mySubMasks[ProvisioningApprovalRequest_FieldPathSelectorStatus] = FullProvisioningApprovalRequest_Status_FieldMask()
					case ProvisioningApprovalRequest_FieldPathSelectorMetadata:
						mySubMasks[ProvisioningApprovalRequest_FieldPathSelectorMetadata] = ntt_meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*ProvisioningApprovalRequest_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &ProvisioningApprovalRequest_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ProvisioningApprovalRequest_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ProvisioningApprovalRequest_FieldMask) FilterInputFields() *ProvisioningApprovalRequest_FieldMask {
	result := &ProvisioningApprovalRequest_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case ProvisioningApprovalRequest_FieldPathSelectorMetadata:
			if _, ok := path.(*ProvisioningApprovalRequest_FieldTerminalPath); ok {
				for _, subpath := range ntt_meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ProvisioningApprovalRequest_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*ProvisioningApprovalRequest_FieldSubPath); ok {
				selectedMask := &ntt_meta.Meta_FieldMask{
					Paths: []ntt_meta.Meta_FieldPath{sub.subPath.(ntt_meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ProvisioningApprovalRequest_FieldSubPath{selector: ProvisioningApprovalRequest_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ProvisioningApprovalRequest_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ProvisioningApprovalRequest_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseProvisioningApprovalRequest_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ProvisioningApprovalRequest_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ProvisioningApprovalRequest_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) AppendPath(path ProvisioningApprovalRequest_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ProvisioningApprovalRequest_FieldPath))
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) GetPaths() []ProvisioningApprovalRequest_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseProvisioningApprovalRequest_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) Set(target, source *ProvisioningApprovalRequest) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ProvisioningApprovalRequest), source.(*ProvisioningApprovalRequest))
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) Project(source *ProvisioningApprovalRequest) *ProvisioningApprovalRequest {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ProvisioningApprovalRequest{}
	specMask := &ProvisioningApprovalRequest_Spec_FieldMask{}
	wholeSpecAccepted := false
	statusMask := &ProvisioningApprovalRequest_Status_FieldMask{}
	wholeStatusAccepted := false
	metadataMask := &ntt_meta.Meta_FieldMask{}
	wholeMetadataAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ProvisioningApprovalRequest_FieldTerminalPath:
			switch tp.selector {
			case ProvisioningApprovalRequest_FieldPathSelectorName:
				result.Name = source.Name
			case ProvisioningApprovalRequest_FieldPathSelectorSpec:
				result.Spec = source.Spec
				wholeSpecAccepted = true
			case ProvisioningApprovalRequest_FieldPathSelectorStatus:
				result.Status = source.Status
				wholeStatusAccepted = true
			case ProvisioningApprovalRequest_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			}
		case *ProvisioningApprovalRequest_FieldSubPath:
			switch tp.selector {
			case ProvisioningApprovalRequest_FieldPathSelectorSpec:
				specMask.AppendPath(tp.subPath.(ProvisioningApprovalRequestSpec_FieldPath))
			case ProvisioningApprovalRequest_FieldPathSelectorStatus:
				statusMask.AppendPath(tp.subPath.(ProvisioningApprovalRequestStatus_FieldPath))
			case ProvisioningApprovalRequest_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(ntt_meta.Meta_FieldPath))
			}
		}
	}
	if wholeSpecAccepted == false && len(specMask.Paths) > 0 {
		result.Spec = specMask.Project(source.GetSpec())
	}
	if wholeStatusAccepted == false && len(statusMask.Paths) > 0 {
		result.Status = statusMask.Project(source.GetStatus())
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	return result
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ProvisioningApprovalRequest))
}

func (fieldMask *ProvisioningApprovalRequest_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type ProvisioningApprovalRequest_Spec_FieldMask struct {
	Paths []ProvisioningApprovalRequestSpec_FieldPath
}

func FullProvisioningApprovalRequest_Spec_FieldMask() *ProvisioningApprovalRequest_Spec_FieldMask {
	res := &ProvisioningApprovalRequest_Spec_FieldMask{}
	res.Paths = append(res.Paths, &ProvisioningApprovalRequestSpec_FieldTerminalPath{selector: ProvisioningApprovalRequestSpec_FieldPathSelectorConclusion})
	res.Paths = append(res.Paths, &ProvisioningApprovalRequestSpec_FieldTerminalPath{selector: ProvisioningApprovalRequestSpec_FieldPathSelectorServiceAccount})
	return res
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) String() string {
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
func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseProvisioningApprovalRequestSpec_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ProvisioningApprovalRequestSpec_FieldTerminalPath); ok {
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

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseProvisioningApprovalRequestSpec_FieldPath(raw)
	})
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) ProtoMessage() {}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) Subtract(other *ProvisioningApprovalRequest_Spec_FieldMask) *ProvisioningApprovalRequest_Spec_FieldMask {
	result := &ProvisioningApprovalRequest_Spec_FieldMask{}
	removedSelectors := make([]bool, 2)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ProvisioningApprovalRequestSpec_FieldTerminalPath:
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

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ProvisioningApprovalRequest_Spec_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) FilterInputFields() *ProvisioningApprovalRequest_Spec_FieldMask {
	result := &ProvisioningApprovalRequest_Spec_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ProvisioningApprovalRequestSpec_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseProvisioningApprovalRequestSpec_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ProvisioningApprovalRequest_Spec_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ProvisioningApprovalRequest_Spec_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) AppendPath(path ProvisioningApprovalRequestSpec_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ProvisioningApprovalRequestSpec_FieldPath))
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) GetPaths() []ProvisioningApprovalRequestSpec_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseProvisioningApprovalRequestSpec_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) Set(target, source *ProvisioningApprovalRequest_Spec) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ProvisioningApprovalRequest_Spec), source.(*ProvisioningApprovalRequest_Spec))
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) Project(source *ProvisioningApprovalRequest_Spec) *ProvisioningApprovalRequest_Spec {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ProvisioningApprovalRequest_Spec{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ProvisioningApprovalRequestSpec_FieldTerminalPath:
			switch tp.selector {
			case ProvisioningApprovalRequestSpec_FieldPathSelectorConclusion:
				result.Conclusion = source.Conclusion
			case ProvisioningApprovalRequestSpec_FieldPathSelectorServiceAccount:
				result.ServiceAccount = source.ServiceAccount
			}
		}
	}
	return result
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ProvisioningApprovalRequest_Spec))
}

func (fieldMask *ProvisioningApprovalRequest_Spec_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type ProvisioningApprovalRequest_Status_FieldMask struct {
	Paths []ProvisioningApprovalRequestStatus_FieldPath
}

func FullProvisioningApprovalRequest_Status_FieldMask() *ProvisioningApprovalRequest_Status_FieldMask {
	res := &ProvisioningApprovalRequest_Status_FieldMask{}
	return res
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) String() string {
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
func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseProvisioningApprovalRequestStatus_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 0)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ProvisioningApprovalRequestStatus_FieldTerminalPath); ok {
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

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseProvisioningApprovalRequestStatus_FieldPath(raw)
	})
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) ProtoMessage() {}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) Subtract(other *ProvisioningApprovalRequest_Status_FieldMask) *ProvisioningApprovalRequest_Status_FieldMask {
	result := &ProvisioningApprovalRequest_Status_FieldMask{}
	removedSelectors := make([]bool, 0)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ProvisioningApprovalRequestStatus_FieldTerminalPath:
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

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ProvisioningApprovalRequest_Status_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) FilterInputFields() *ProvisioningApprovalRequest_Status_FieldMask {
	result := &ProvisioningApprovalRequest_Status_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ProvisioningApprovalRequestStatus_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseProvisioningApprovalRequestStatus_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ProvisioningApprovalRequest_Status_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ProvisioningApprovalRequest_Status_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) AppendPath(path ProvisioningApprovalRequestStatus_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ProvisioningApprovalRequestStatus_FieldPath))
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) GetPaths() []ProvisioningApprovalRequestStatus_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseProvisioningApprovalRequestStatus_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) Set(target, source *ProvisioningApprovalRequest_Status) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ProvisioningApprovalRequest_Status), source.(*ProvisioningApprovalRequest_Status))
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) Project(source *ProvisioningApprovalRequest_Status) *ProvisioningApprovalRequest_Status {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ProvisioningApprovalRequest_Status{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ProvisioningApprovalRequestStatus_FieldTerminalPath:
			switch tp.selector {
			}
		}
	}
	return result
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ProvisioningApprovalRequest_Status))
}

func (fieldMask *ProvisioningApprovalRequest_Status_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
