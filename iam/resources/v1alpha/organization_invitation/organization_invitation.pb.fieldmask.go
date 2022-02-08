// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha/organization_invitation.proto
// DO NOT EDIT!!!

package organization_invitation

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
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/common"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
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
	_ = &iam_common.Actor{}
	_ = &organization.Organization{}
)

type OrganizationInvitation_FieldMask struct {
	Paths []OrganizationInvitation_FieldPath
}

func FullOrganizationInvitation_FieldMask() *OrganizationInvitation_FieldMask {
	res := &OrganizationInvitation_FieldMask{}
	res.Paths = append(res.Paths, &OrganizationInvitation_FieldTerminalPath{selector: OrganizationInvitation_FieldPathSelectorName})
	res.Paths = append(res.Paths, &OrganizationInvitation_FieldTerminalPath{selector: OrganizationInvitation_FieldPathSelectorInvitation})
	return res
}

func (fieldMask *OrganizationInvitation_FieldMask) String() string {
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
func (fieldMask *OrganizationInvitation_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *OrganizationInvitation_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseOrganizationInvitation_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *OrganizationInvitation_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*OrganizationInvitation_FieldTerminalPath); ok {
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

func (fieldMask *OrganizationInvitation_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseOrganizationInvitation_FieldPath(raw)
	})
}

func (fieldMask *OrganizationInvitation_FieldMask) ProtoMessage() {}

func (fieldMask *OrganizationInvitation_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *OrganizationInvitation_FieldMask) Subtract(other *OrganizationInvitation_FieldMask) *OrganizationInvitation_FieldMask {
	result := &OrganizationInvitation_FieldMask{}
	removedSelectors := make([]bool, 2)
	otherSubMasks := map[OrganizationInvitation_FieldPathSelector]gotenobject.FieldMask{
		OrganizationInvitation_FieldPathSelectorInvitation: &iam_common.Invitation_FieldMask{},
	}
	mySubMasks := map[OrganizationInvitation_FieldPathSelector]gotenobject.FieldMask{
		OrganizationInvitation_FieldPathSelectorInvitation: &iam_common.Invitation_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *OrganizationInvitation_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *OrganizationInvitation_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*OrganizationInvitation_FieldTerminalPath); ok {
					switch tp.selector {
					case OrganizationInvitation_FieldPathSelectorInvitation:
						mySubMasks[OrganizationInvitation_FieldPathSelectorInvitation] = iam_common.FullInvitation_FieldMask()
					}
				} else if tp, ok := path.(*OrganizationInvitation_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &OrganizationInvitation_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *OrganizationInvitation_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*OrganizationInvitation_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *OrganizationInvitation_FieldMask) FilterInputFields() *OrganizationInvitation_FieldMask {
	result := &OrganizationInvitation_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *OrganizationInvitation_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *OrganizationInvitation_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]OrganizationInvitation_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseOrganizationInvitation_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask OrganizationInvitation_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *OrganizationInvitation_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *OrganizationInvitation_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask OrganizationInvitation_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *OrganizationInvitation_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *OrganizationInvitation_FieldMask) AppendPath(path OrganizationInvitation_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *OrganizationInvitation_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(OrganizationInvitation_FieldPath))
}

func (fieldMask *OrganizationInvitation_FieldMask) GetPaths() []OrganizationInvitation_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *OrganizationInvitation_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *OrganizationInvitation_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseOrganizationInvitation_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *OrganizationInvitation_FieldMask) Set(target, source *OrganizationInvitation) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *OrganizationInvitation_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*OrganizationInvitation), source.(*OrganizationInvitation))
}

func (fieldMask *OrganizationInvitation_FieldMask) Project(source *OrganizationInvitation) *OrganizationInvitation {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &OrganizationInvitation{}
	invitationMask := &iam_common.Invitation_FieldMask{}
	wholeInvitationAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *OrganizationInvitation_FieldTerminalPath:
			switch tp.selector {
			case OrganizationInvitation_FieldPathSelectorName:
				result.Name = source.Name
			case OrganizationInvitation_FieldPathSelectorInvitation:
				result.Invitation = source.Invitation
				wholeInvitationAccepted = true
			}
		case *OrganizationInvitation_FieldSubPath:
			switch tp.selector {
			case OrganizationInvitation_FieldPathSelectorInvitation:
				invitationMask.AppendPath(tp.subPath.(iam_common.Invitation_FieldPath))
			}
		}
	}
	if wholeInvitationAccepted == false && len(invitationMask.Paths) > 0 {
		result.Invitation = invitationMask.Project(source.GetInvitation())
	}
	return result
}

func (fieldMask *OrganizationInvitation_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*OrganizationInvitation))
}

func (fieldMask *OrganizationInvitation_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}