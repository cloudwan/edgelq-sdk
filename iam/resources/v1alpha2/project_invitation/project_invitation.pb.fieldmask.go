// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha2/project_invitation.proto
// DO NOT EDIT!!!

package project_invitation

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
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/common"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
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
	_ = &ntt_meta.Meta{}
	_ = &iam_common.Actor{}
	_ = &project.Project{}
)

type ProjectInvitation_FieldMask struct {
	Paths []ProjectInvitation_FieldPath
}

func FullProjectInvitation_FieldMask() *ProjectInvitation_FieldMask {
	res := &ProjectInvitation_FieldMask{}
	res.Paths = append(res.Paths, &ProjectInvitation_FieldTerminalPath{selector: ProjectInvitation_FieldPathSelectorName})
	res.Paths = append(res.Paths, &ProjectInvitation_FieldTerminalPath{selector: ProjectInvitation_FieldPathSelectorProjectDisplayName})
	res.Paths = append(res.Paths, &ProjectInvitation_FieldTerminalPath{selector: ProjectInvitation_FieldPathSelectorInvitation})
	res.Paths = append(res.Paths, &ProjectInvitation_FieldTerminalPath{selector: ProjectInvitation_FieldPathSelectorMetadata})
	return res
}

func (fieldMask *ProjectInvitation_FieldMask) String() string {
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
func (fieldMask *ProjectInvitation_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ProjectInvitation_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseProjectInvitation_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ProjectInvitation_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 4)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ProjectInvitation_FieldTerminalPath); ok {
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

func (fieldMask *ProjectInvitation_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseProjectInvitation_FieldPath(raw)
	})
}

func (fieldMask *ProjectInvitation_FieldMask) ProtoMessage() {}

func (fieldMask *ProjectInvitation_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ProjectInvitation_FieldMask) Subtract(other *ProjectInvitation_FieldMask) *ProjectInvitation_FieldMask {
	result := &ProjectInvitation_FieldMask{}
	removedSelectors := make([]bool, 4)
	otherSubMasks := map[ProjectInvitation_FieldPathSelector]gotenobject.FieldMask{
		ProjectInvitation_FieldPathSelectorInvitation: &iam_common.Invitation_FieldMask{},
		ProjectInvitation_FieldPathSelectorMetadata:   &ntt_meta.Meta_FieldMask{},
	}
	mySubMasks := map[ProjectInvitation_FieldPathSelector]gotenobject.FieldMask{
		ProjectInvitation_FieldPathSelectorInvitation: &iam_common.Invitation_FieldMask{},
		ProjectInvitation_FieldPathSelectorMetadata:   &ntt_meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ProjectInvitation_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *ProjectInvitation_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*ProjectInvitation_FieldTerminalPath); ok {
					switch tp.selector {
					case ProjectInvitation_FieldPathSelectorInvitation:
						mySubMasks[ProjectInvitation_FieldPathSelectorInvitation] = iam_common.FullInvitation_FieldMask()
					case ProjectInvitation_FieldPathSelectorMetadata:
						mySubMasks[ProjectInvitation_FieldPathSelectorMetadata] = ntt_meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*ProjectInvitation_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &ProjectInvitation_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *ProjectInvitation_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ProjectInvitation_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ProjectInvitation_FieldMask) FilterInputFields() *ProjectInvitation_FieldMask {
	result := &ProjectInvitation_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case ProjectInvitation_FieldPathSelectorMetadata:
			if _, ok := path.(*ProjectInvitation_FieldTerminalPath); ok {
				for _, subpath := range ntt_meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ProjectInvitation_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*ProjectInvitation_FieldSubPath); ok {
				selectedMask := &ntt_meta.Meta_FieldMask{
					Paths: []ntt_meta.Meta_FieldPath{sub.subPath.(ntt_meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ProjectInvitation_FieldSubPath{selector: ProjectInvitation_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ProjectInvitation_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ProjectInvitation_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ProjectInvitation_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseProjectInvitation_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ProjectInvitation_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ProjectInvitation_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProjectInvitation_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ProjectInvitation_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ProjectInvitation_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProjectInvitation_FieldMask) AppendPath(path ProjectInvitation_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ProjectInvitation_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ProjectInvitation_FieldPath))
}

func (fieldMask *ProjectInvitation_FieldMask) GetPaths() []ProjectInvitation_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ProjectInvitation_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ProjectInvitation_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseProjectInvitation_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ProjectInvitation_FieldMask) Set(target, source *ProjectInvitation) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ProjectInvitation_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ProjectInvitation), source.(*ProjectInvitation))
}

func (fieldMask *ProjectInvitation_FieldMask) Project(source *ProjectInvitation) *ProjectInvitation {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ProjectInvitation{}
	invitationMask := &iam_common.Invitation_FieldMask{}
	wholeInvitationAccepted := false
	metadataMask := &ntt_meta.Meta_FieldMask{}
	wholeMetadataAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ProjectInvitation_FieldTerminalPath:
			switch tp.selector {
			case ProjectInvitation_FieldPathSelectorName:
				result.Name = source.Name
			case ProjectInvitation_FieldPathSelectorProjectDisplayName:
				result.ProjectDisplayName = source.ProjectDisplayName
			case ProjectInvitation_FieldPathSelectorInvitation:
				result.Invitation = source.Invitation
				wholeInvitationAccepted = true
			case ProjectInvitation_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			}
		case *ProjectInvitation_FieldSubPath:
			switch tp.selector {
			case ProjectInvitation_FieldPathSelectorInvitation:
				invitationMask.AppendPath(tp.subPath.(iam_common.Invitation_FieldPath))
			case ProjectInvitation_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(ntt_meta.Meta_FieldPath))
			}
		}
	}
	if wholeInvitationAccepted == false && len(invitationMask.Paths) > 0 {
		result.Invitation = invitationMask.Project(source.GetInvitation())
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	return result
}

func (fieldMask *ProjectInvitation_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ProjectInvitation))
}

func (fieldMask *ProjectInvitation_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
