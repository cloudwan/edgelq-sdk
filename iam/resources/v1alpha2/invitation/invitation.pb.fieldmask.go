// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha2/invitation.proto
// DO NOT EDIT!!!

package iam_invitation

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
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/role"
	service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/service_account"
	user "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/user"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	_ = &role.Role{}
	_ = &service_account.ServiceAccount{}
	_ = &user.User{}
	_ = &timestamppb.Timestamp{}
)

type Actor_FieldMask struct {
	Paths []Actor_FieldPath
}

func FullActor_FieldMask() *Actor_FieldMask {
	res := &Actor_FieldMask{}
	res.Paths = append(res.Paths, &Actor_FieldTerminalPath{selector: Actor_FieldPathSelectorUser})
	res.Paths = append(res.Paths, &Actor_FieldTerminalPath{selector: Actor_FieldPathSelectorServiceAccount})
	return res
}

func (fieldMask *Actor_FieldMask) String() string {
	if fieldMask == nil {
		return "<nil>"
	}
	pathsStr := make([]string, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		pathsStr = append(pathsStr, path.String())
	}
	return strings.Join(pathsStr, ", ")
}

func (fieldMask *Actor_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*Actor_FieldTerminalPath); ok {
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

func (fieldMask *Actor_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseActor_FieldPath(raw)
	})
}

func (fieldMask *Actor_FieldMask) ProtoMessage() {}

func (fieldMask *Actor_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Actor_FieldMask) Subtract(other *Actor_FieldMask) *Actor_FieldMask {
	result := &Actor_FieldMask{}
	removedSelectors := make([]bool, 2)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *Actor_FieldTerminalPath:
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

func (fieldMask *Actor_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Actor_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Actor_FieldMask) FilterInputFields() *Actor_FieldMask {
	result := &Actor_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Actor_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Actor_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]Actor_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseActor_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Actor_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Actor_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Actor_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Actor_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Actor_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Actor_FieldMask) AppendPath(path Actor_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Actor_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(Actor_FieldPath))
}

func (fieldMask *Actor_FieldMask) GetPaths() []Actor_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Actor_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Actor_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseActor_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Actor_FieldMask) Set(target, source *Actor) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Actor_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Actor), source.(*Actor))
}

func (fieldMask *Actor_FieldMask) Project(source *Actor) *Actor {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Actor{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *Actor_FieldTerminalPath:
			switch tp.selector {
			case Actor_FieldPathSelectorUser:
				result.User = source.User
			case Actor_FieldPathSelectorServiceAccount:
				result.ServiceAccount = source.ServiceAccount
			}
		}
	}
	return result
}

func (fieldMask *Actor_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Actor))
}

func (fieldMask *Actor_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type Invitation_FieldMask struct {
	Paths []Invitation_FieldPath
}

func FullInvitation_FieldMask() *Invitation_FieldMask {
	res := &Invitation_FieldMask{}
	res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorInviteeEmail})
	res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorInviterActor})
	res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorInviterFullName})
	res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorInviterEmail})
	res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorLanguageCode})
	res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorRoles})
	res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorExpirationDate})
	res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorExtras})
	res.Paths = append(res.Paths, &Invitation_FieldTerminalPath{selector: Invitation_FieldPathSelectorState})
	return res
}

func (fieldMask *Invitation_FieldMask) String() string {
	if fieldMask == nil {
		return "<nil>"
	}
	pathsStr := make([]string, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		pathsStr = append(pathsStr, path.String())
	}
	return strings.Join(pathsStr, ", ")
}

func (fieldMask *Invitation_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 9)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*Invitation_FieldTerminalPath); ok {
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

func (fieldMask *Invitation_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseInvitation_FieldPath(raw)
	})
}

func (fieldMask *Invitation_FieldMask) ProtoMessage() {}

func (fieldMask *Invitation_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Invitation_FieldMask) Subtract(other *Invitation_FieldMask) *Invitation_FieldMask {
	result := &Invitation_FieldMask{}
	removedSelectors := make([]bool, 9)
	otherSubMasks := map[Invitation_FieldPathSelector]gotenobject.FieldMask{
		Invitation_FieldPathSelectorInviterActor: &Actor_FieldMask{},
	}
	mySubMasks := map[Invitation_FieldPathSelector]gotenobject.FieldMask{
		Invitation_FieldPathSelectorInviterActor: &Actor_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *Invitation_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *Invitation_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*Invitation_FieldTerminalPath); ok {
					switch tp.selector {
					case Invitation_FieldPathSelectorInviterActor:
						mySubMasks[Invitation_FieldPathSelectorInviterActor] = FullActor_FieldMask()
					}
				} else if tp, ok := path.(*Invitation_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &Invitation_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *Invitation_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Invitation_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Invitation_FieldMask) FilterInputFields() *Invitation_FieldMask {
	result := &Invitation_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Invitation_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Invitation_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]Invitation_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseInvitation_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Invitation_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Invitation_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Invitation_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Invitation_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Invitation_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Invitation_FieldMask) AppendPath(path Invitation_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Invitation_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(Invitation_FieldPath))
}

func (fieldMask *Invitation_FieldMask) GetPaths() []Invitation_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Invitation_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Invitation_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseInvitation_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Invitation_FieldMask) Set(target, source *Invitation) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Invitation_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Invitation), source.(*Invitation))
}

func (fieldMask *Invitation_FieldMask) Project(source *Invitation) *Invitation {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Invitation{}
	inviterActorMask := &Actor_FieldMask{}
	wholeInviterActorAccepted := false
	var extrasMapKeys []string
	wholeExtrasAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *Invitation_FieldTerminalPath:
			switch tp.selector {
			case Invitation_FieldPathSelectorInviteeEmail:
				result.InviteeEmail = source.InviteeEmail
			case Invitation_FieldPathSelectorInviterActor:
				result.InviterActor = source.InviterActor
				wholeInviterActorAccepted = true
			case Invitation_FieldPathSelectorInviterFullName:
				result.InviterFullName = source.InviterFullName
			case Invitation_FieldPathSelectorInviterEmail:
				result.InviterEmail = source.InviterEmail
			case Invitation_FieldPathSelectorLanguageCode:
				result.LanguageCode = source.LanguageCode
			case Invitation_FieldPathSelectorRoles:
				result.Roles = source.Roles
			case Invitation_FieldPathSelectorExpirationDate:
				result.ExpirationDate = source.ExpirationDate
			case Invitation_FieldPathSelectorExtras:
				result.Extras = source.Extras
				wholeExtrasAccepted = true
			case Invitation_FieldPathSelectorState:
				result.State = source.State
			}
		case *Invitation_FieldSubPath:
			switch tp.selector {
			case Invitation_FieldPathSelectorInviterActor:
				inviterActorMask.AppendPath(tp.subPath.(Actor_FieldPath))
			}
		case *Invitation_FieldPathMap:
			switch tp.selector {
			case Invitation_FieldPathSelectorExtras:
				extrasMapKeys = append(extrasMapKeys, tp.key)
			}
		}
	}
	if wholeInviterActorAccepted == false && len(inviterActorMask.Paths) > 0 {
		result.InviterActor = inviterActorMask.Project(source.GetInviterActor())
	}
	if wholeExtrasAccepted == false && len(extrasMapKeys) > 0 && source.GetExtras() != nil {
		copiedMap := map[string]string{}
		sourceMap := source.GetExtras()
		for _, key := range extrasMapKeys {
			copiedMap[key] = sourceMap[key]
		}
		result.Extras = copiedMap
	}
	return result
}

func (fieldMask *Invitation_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Invitation))
}

func (fieldMask *Invitation_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
