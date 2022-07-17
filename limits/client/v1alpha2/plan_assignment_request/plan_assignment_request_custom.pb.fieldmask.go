// Code generated by protoc-gen-goten-object
// File: edgelq/limits/proto/v1alpha2/plan_assignment_request_custom.proto
// DO NOT EDIT!!!

package plan_assignment_request_client

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
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	accepted_plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/accepted_plan"
	plan_assignment_request "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/plan_assignment_request"
)

// ensure the imports are used
var (
	_ = json.Marshaler(nil)
	_ = strings.Builder{}

	_ = firestorepb.Value{}
	_ = codes.NotFound
	_ = status.Status{}
	_ = proto.Message(nil)
	_ = preflect.Message(nil)
	_ = fieldmaskpb.FieldMask{}

	_ = gotenobject.FieldMask(nil)
)

// make sure we're using proto imports
var (
	_ = &iam_organization.Organization{}
	_ = &accepted_plan.AcceptedPlan{}
	_ = &plan_assignment_request.PlanAssignmentRequest{}
)

type AcceptPlanAssignmentRequest_FieldMask struct {
	Paths []AcceptPlanAssignmentRequest_FieldPath
}

func FullAcceptPlanAssignmentRequest_FieldMask() *AcceptPlanAssignmentRequest_FieldMask {
	res := &AcceptPlanAssignmentRequest_FieldMask{}
	res.Paths = append(res.Paths, &AcceptPlanAssignmentRequest_FieldTerminalPath{selector: AcceptPlanAssignmentRequest_FieldPathSelectorName})
	res.Paths = append(res.Paths, &AcceptPlanAssignmentRequest_FieldTerminalPath{selector: AcceptPlanAssignmentRequest_FieldPathSelectorApprover})
	return res
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) String() string {
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
func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseAcceptPlanAssignmentRequest_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*AcceptPlanAssignmentRequest_FieldTerminalPath); ok {
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

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseAcceptPlanAssignmentRequest_FieldPath(raw)
	})
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) ProtoMessage() {}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) Subtract(other *AcceptPlanAssignmentRequest_FieldMask) *AcceptPlanAssignmentRequest_FieldMask {
	result := &AcceptPlanAssignmentRequest_FieldMask{}
	removedSelectors := make([]bool, 2)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *AcceptPlanAssignmentRequest_FieldTerminalPath:
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

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*AcceptPlanAssignmentRequest_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) FilterInputFields() *AcceptPlanAssignmentRequest_FieldMask {
	result := &AcceptPlanAssignmentRequest_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]AcceptPlanAssignmentRequest_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseAcceptPlanAssignmentRequest_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask AcceptPlanAssignmentRequest_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask AcceptPlanAssignmentRequest_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) AppendPath(path AcceptPlanAssignmentRequest_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(AcceptPlanAssignmentRequest_FieldPath))
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) GetPaths() []AcceptPlanAssignmentRequest_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseAcceptPlanAssignmentRequest_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) Set(target, source *AcceptPlanAssignmentRequest) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*AcceptPlanAssignmentRequest), source.(*AcceptPlanAssignmentRequest))
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) Project(source *AcceptPlanAssignmentRequest) *AcceptPlanAssignmentRequest {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &AcceptPlanAssignmentRequest{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *AcceptPlanAssignmentRequest_FieldTerminalPath:
			switch tp.selector {
			case AcceptPlanAssignmentRequest_FieldPathSelectorName:
				result.Name = source.Name
			case AcceptPlanAssignmentRequest_FieldPathSelectorApprover:
				result.Approver = source.Approver
			}
		}
	}
	return result
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*AcceptPlanAssignmentRequest))
}

func (fieldMask *AcceptPlanAssignmentRequest_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type AcceptPlanAssignmentResponse_FieldMask struct {
	Paths []AcceptPlanAssignmentResponse_FieldPath
}

func FullAcceptPlanAssignmentResponse_FieldMask() *AcceptPlanAssignmentResponse_FieldMask {
	res := &AcceptPlanAssignmentResponse_FieldMask{}
	res.Paths = append(res.Paths, &AcceptPlanAssignmentResponse_FieldTerminalPath{selector: AcceptPlanAssignmentResponse_FieldPathSelectorAcceptedPlan})
	return res
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) String() string {
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
func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseAcceptPlanAssignmentResponse_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 1)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*AcceptPlanAssignmentResponse_FieldTerminalPath); ok {
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

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseAcceptPlanAssignmentResponse_FieldPath(raw)
	})
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) ProtoMessage() {}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) Subtract(other *AcceptPlanAssignmentResponse_FieldMask) *AcceptPlanAssignmentResponse_FieldMask {
	result := &AcceptPlanAssignmentResponse_FieldMask{}
	removedSelectors := make([]bool, 1)
	otherSubMasks := map[AcceptPlanAssignmentResponse_FieldPathSelector]gotenobject.FieldMask{
		AcceptPlanAssignmentResponse_FieldPathSelectorAcceptedPlan: &accepted_plan.AcceptedPlan_FieldMask{},
	}
	mySubMasks := map[AcceptPlanAssignmentResponse_FieldPathSelector]gotenobject.FieldMask{
		AcceptPlanAssignmentResponse_FieldPathSelectorAcceptedPlan: &accepted_plan.AcceptedPlan_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *AcceptPlanAssignmentResponse_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *AcceptPlanAssignmentResponse_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*AcceptPlanAssignmentResponse_FieldTerminalPath); ok {
					switch tp.selector {
					case AcceptPlanAssignmentResponse_FieldPathSelectorAcceptedPlan:
						mySubMasks[AcceptPlanAssignmentResponse_FieldPathSelectorAcceptedPlan] = accepted_plan.FullAcceptedPlan_FieldMask()
					}
				} else if tp, ok := path.(*AcceptPlanAssignmentResponse_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &AcceptPlanAssignmentResponse_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*AcceptPlanAssignmentResponse_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) FilterInputFields() *AcceptPlanAssignmentResponse_FieldMask {
	result := &AcceptPlanAssignmentResponse_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case AcceptPlanAssignmentResponse_FieldPathSelectorAcceptedPlan:
			if _, ok := path.(*AcceptPlanAssignmentResponse_FieldTerminalPath); ok {
				for _, subpath := range accepted_plan.FullAcceptedPlan_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &AcceptPlanAssignmentResponse_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*AcceptPlanAssignmentResponse_FieldSubPath); ok {
				selectedMask := &accepted_plan.AcceptedPlan_FieldMask{
					Paths: []accepted_plan.AcceptedPlan_FieldPath{sub.subPath.(accepted_plan.AcceptedPlan_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &AcceptPlanAssignmentResponse_FieldSubPath{selector: AcceptPlanAssignmentResponse_FieldPathSelectorAcceptedPlan, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]AcceptPlanAssignmentResponse_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseAcceptPlanAssignmentResponse_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask AcceptPlanAssignmentResponse_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask AcceptPlanAssignmentResponse_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) AppendPath(path AcceptPlanAssignmentResponse_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(AcceptPlanAssignmentResponse_FieldPath))
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) GetPaths() []AcceptPlanAssignmentResponse_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseAcceptPlanAssignmentResponse_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) Set(target, source *AcceptPlanAssignmentResponse) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*AcceptPlanAssignmentResponse), source.(*AcceptPlanAssignmentResponse))
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) Project(source *AcceptPlanAssignmentResponse) *AcceptPlanAssignmentResponse {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &AcceptPlanAssignmentResponse{}
	acceptedPlanMask := &accepted_plan.AcceptedPlan_FieldMask{}
	wholeAcceptedPlanAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *AcceptPlanAssignmentResponse_FieldTerminalPath:
			switch tp.selector {
			case AcceptPlanAssignmentResponse_FieldPathSelectorAcceptedPlan:
				result.AcceptedPlan = source.AcceptedPlan
				wholeAcceptedPlanAccepted = true
			}
		case *AcceptPlanAssignmentResponse_FieldSubPath:
			switch tp.selector {
			case AcceptPlanAssignmentResponse_FieldPathSelectorAcceptedPlan:
				acceptedPlanMask.AppendPath(tp.subPath.(accepted_plan.AcceptedPlan_FieldPath))
			}
		}
	}
	if wholeAcceptedPlanAccepted == false && len(acceptedPlanMask.Paths) > 0 {
		result.AcceptedPlan = acceptedPlanMask.Project(source.GetAcceptedPlan())
	}
	return result
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*AcceptPlanAssignmentResponse))
}

func (fieldMask *AcceptPlanAssignmentResponse_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type DeclinePlanAssignmentRequest_FieldMask struct {
	Paths []DeclinePlanAssignmentRequest_FieldPath
}

func FullDeclinePlanAssignmentRequest_FieldMask() *DeclinePlanAssignmentRequest_FieldMask {
	res := &DeclinePlanAssignmentRequest_FieldMask{}
	res.Paths = append(res.Paths, &DeclinePlanAssignmentRequest_FieldTerminalPath{selector: DeclinePlanAssignmentRequest_FieldPathSelectorName})
	res.Paths = append(res.Paths, &DeclinePlanAssignmentRequest_FieldTerminalPath{selector: DeclinePlanAssignmentRequest_FieldPathSelectorApprover})
	res.Paths = append(res.Paths, &DeclinePlanAssignmentRequest_FieldTerminalPath{selector: DeclinePlanAssignmentRequest_FieldPathSelectorReason})
	return res
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) String() string {
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
func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseDeclinePlanAssignmentRequest_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 3)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*DeclinePlanAssignmentRequest_FieldTerminalPath); ok {
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

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseDeclinePlanAssignmentRequest_FieldPath(raw)
	})
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) ProtoMessage() {}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) Subtract(other *DeclinePlanAssignmentRequest_FieldMask) *DeclinePlanAssignmentRequest_FieldMask {
	result := &DeclinePlanAssignmentRequest_FieldMask{}
	removedSelectors := make([]bool, 3)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *DeclinePlanAssignmentRequest_FieldTerminalPath:
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

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*DeclinePlanAssignmentRequest_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) FilterInputFields() *DeclinePlanAssignmentRequest_FieldMask {
	result := &DeclinePlanAssignmentRequest_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]DeclinePlanAssignmentRequest_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseDeclinePlanAssignmentRequest_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask DeclinePlanAssignmentRequest_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask DeclinePlanAssignmentRequest_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) AppendPath(path DeclinePlanAssignmentRequest_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(DeclinePlanAssignmentRequest_FieldPath))
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) GetPaths() []DeclinePlanAssignmentRequest_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseDeclinePlanAssignmentRequest_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) Set(target, source *DeclinePlanAssignmentRequest) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*DeclinePlanAssignmentRequest), source.(*DeclinePlanAssignmentRequest))
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) Project(source *DeclinePlanAssignmentRequest) *DeclinePlanAssignmentRequest {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &DeclinePlanAssignmentRequest{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *DeclinePlanAssignmentRequest_FieldTerminalPath:
			switch tp.selector {
			case DeclinePlanAssignmentRequest_FieldPathSelectorName:
				result.Name = source.Name
			case DeclinePlanAssignmentRequest_FieldPathSelectorApprover:
				result.Approver = source.Approver
			case DeclinePlanAssignmentRequest_FieldPathSelectorReason:
				result.Reason = source.Reason
			}
		}
	}
	return result
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*DeclinePlanAssignmentRequest))
}

func (fieldMask *DeclinePlanAssignmentRequest_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type DeclinePlanAssignmentResponse_FieldMask struct {
	Paths []DeclinePlanAssignmentResponse_FieldPath
}

func FullDeclinePlanAssignmentResponse_FieldMask() *DeclinePlanAssignmentResponse_FieldMask {
	res := &DeclinePlanAssignmentResponse_FieldMask{}
	return res
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) String() string {
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
func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseDeclinePlanAssignmentResponse_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 0)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*DeclinePlanAssignmentResponse_FieldTerminalPath); ok {
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

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseDeclinePlanAssignmentResponse_FieldPath(raw)
	})
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) ProtoMessage() {}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) Subtract(other *DeclinePlanAssignmentResponse_FieldMask) *DeclinePlanAssignmentResponse_FieldMask {
	result := &DeclinePlanAssignmentResponse_FieldMask{}
	removedSelectors := make([]bool, 0)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *DeclinePlanAssignmentResponse_FieldTerminalPath:
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

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*DeclinePlanAssignmentResponse_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) FilterInputFields() *DeclinePlanAssignmentResponse_FieldMask {
	result := &DeclinePlanAssignmentResponse_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]DeclinePlanAssignmentResponse_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseDeclinePlanAssignmentResponse_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask DeclinePlanAssignmentResponse_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask DeclinePlanAssignmentResponse_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) AppendPath(path DeclinePlanAssignmentResponse_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(DeclinePlanAssignmentResponse_FieldPath))
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) GetPaths() []DeclinePlanAssignmentResponse_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseDeclinePlanAssignmentResponse_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) Set(target, source *DeclinePlanAssignmentResponse) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*DeclinePlanAssignmentResponse), source.(*DeclinePlanAssignmentResponse))
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) Project(source *DeclinePlanAssignmentResponse) *DeclinePlanAssignmentResponse {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &DeclinePlanAssignmentResponse{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *DeclinePlanAssignmentResponse_FieldTerminalPath:
			switch tp.selector {
			}
		}
	}
	return result
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*DeclinePlanAssignmentResponse))
}

func (fieldMask *DeclinePlanAssignmentResponse_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
