// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha/user.proto
// DO NOT EDIT!!!

package user

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
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	_ = &timestamp.Timestamp{}
)

type User_FieldMask struct {
	Paths []User_FieldPath
}

func FullUser_FieldMask() *User_FieldMask {
	res := &User_FieldMask{}
	res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorName})
	res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorFullName})
	res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorEmail})
	res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorEmailVerified})
	res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorAuthInfo})
	res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorSettings})
	res.Paths = append(res.Paths, &User_FieldTerminalPath{selector: User_FieldPathSelectorRefreshedTime})
	return res
}

func (fieldMask *User_FieldMask) String() string {
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
func (fieldMask *User_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *User_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseUser_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *User_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 8)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*User_FieldTerminalPath); ok {
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

func (fieldMask *User_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseUser_FieldPath(raw)
	})
}

func (fieldMask *User_FieldMask) ProtoMessage() {}

func (fieldMask *User_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *User_FieldMask) Subtract(other *User_FieldMask) *User_FieldMask {
	result := &User_FieldMask{}
	removedSelectors := make([]bool, 8)
	otherSubMasks := map[User_FieldPathSelector]gotenobject.FieldMask{
		User_FieldPathSelectorMetadata: &ntt_meta.Meta_FieldMask{},
		User_FieldPathSelectorAuthInfo: &User_AuthInfo_FieldMask{},
	}
	mySubMasks := map[User_FieldPathSelector]gotenobject.FieldMask{
		User_FieldPathSelectorMetadata: &ntt_meta.Meta_FieldMask{},
		User_FieldPathSelectorAuthInfo: &User_AuthInfo_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *User_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *User_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*User_FieldTerminalPath); ok {
					switch tp.selector {
					case User_FieldPathSelectorMetadata:
						mySubMasks[User_FieldPathSelectorMetadata] = ntt_meta.FullMeta_FieldMask()
					case User_FieldPathSelectorAuthInfo:
						mySubMasks[User_FieldPathSelectorAuthInfo] = FullUser_AuthInfo_FieldMask()
					}
				} else if tp, ok := path.(*User_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &User_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *User_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*User_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *User_FieldMask) FilterInputFields() *User_FieldMask {
	result := &User_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case User_FieldPathSelectorRefreshedTime:
		case User_FieldPathSelectorMetadata:
			if _, ok := path.(*User_FieldTerminalPath); ok {
				for _, subpath := range ntt_meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &User_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*User_FieldSubPath); ok {
				selectedMask := &ntt_meta.Meta_FieldMask{
					Paths: []ntt_meta.Meta_FieldPath{sub.subPath.(ntt_meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &User_FieldSubPath{selector: User_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *User_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *User_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]User_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseUser_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask User_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *User_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *User_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask User_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *User_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *User_FieldMask) AppendPath(path User_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *User_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(User_FieldPath))
}

func (fieldMask *User_FieldMask) GetPaths() []User_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *User_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *User_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseUser_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *User_FieldMask) Set(target, source *User) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *User_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*User), source.(*User))
}

func (fieldMask *User_FieldMask) Project(source *User) *User {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &User{}
	metadataMask := &ntt_meta.Meta_FieldMask{}
	wholeMetadataAccepted := false
	authInfoMask := &User_AuthInfo_FieldMask{}
	wholeAuthInfoAccepted := false
	var settingsMapKeys []string
	wholeSettingsAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *User_FieldTerminalPath:
			switch tp.selector {
			case User_FieldPathSelectorName:
				result.Name = source.Name
			case User_FieldPathSelectorFullName:
				result.FullName = source.FullName
			case User_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case User_FieldPathSelectorEmail:
				result.Email = source.Email
			case User_FieldPathSelectorEmailVerified:
				result.EmailVerified = source.EmailVerified
			case User_FieldPathSelectorAuthInfo:
				result.AuthInfo = source.AuthInfo
				wholeAuthInfoAccepted = true
			case User_FieldPathSelectorSettings:
				result.Settings = source.Settings
				wholeSettingsAccepted = true
			case User_FieldPathSelectorRefreshedTime:
				result.RefreshedTime = source.RefreshedTime
			}
		case *User_FieldSubPath:
			switch tp.selector {
			case User_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(ntt_meta.Meta_FieldPath))
			case User_FieldPathSelectorAuthInfo:
				authInfoMask.AppendPath(tp.subPath.(UserAuthInfo_FieldPath))
			}
		case *User_FieldPathMap:
			switch tp.selector {
			case User_FieldPathSelectorSettings:
				settingsMapKeys = append(settingsMapKeys, tp.key)
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	if wholeAuthInfoAccepted == false && len(authInfoMask.Paths) > 0 {
		result.AuthInfo = authInfoMask.Project(source.GetAuthInfo())
	}
	if wholeSettingsAccepted == false && len(settingsMapKeys) > 0 && source.GetSettings() != nil {
		copiedMap := map[string]string{}
		sourceMap := source.GetSettings()
		for _, key := range settingsMapKeys {
			copiedMap[key] = sourceMap[key]
		}
		result.Settings = copiedMap
	}
	return result
}

func (fieldMask *User_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*User))
}

func (fieldMask *User_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type User_AuthInfo_FieldMask struct {
	Paths []UserAuthInfo_FieldPath
}

func FullUser_AuthInfo_FieldMask() *User_AuthInfo_FieldMask {
	res := &User_AuthInfo_FieldMask{}
	res.Paths = append(res.Paths, &UserAuthInfo_FieldTerminalPath{selector: UserAuthInfo_FieldPathSelectorProvider})
	res.Paths = append(res.Paths, &UserAuthInfo_FieldTerminalPath{selector: UserAuthInfo_FieldPathSelectorId})
	return res
}

func (fieldMask *User_AuthInfo_FieldMask) String() string {
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
func (fieldMask *User_AuthInfo_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *User_AuthInfo_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseUserAuthInfo_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *User_AuthInfo_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*UserAuthInfo_FieldTerminalPath); ok {
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

func (fieldMask *User_AuthInfo_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseUserAuthInfo_FieldPath(raw)
	})
}

func (fieldMask *User_AuthInfo_FieldMask) ProtoMessage() {}

func (fieldMask *User_AuthInfo_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *User_AuthInfo_FieldMask) Subtract(other *User_AuthInfo_FieldMask) *User_AuthInfo_FieldMask {
	result := &User_AuthInfo_FieldMask{}
	removedSelectors := make([]bool, 2)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *UserAuthInfo_FieldTerminalPath:
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

func (fieldMask *User_AuthInfo_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*User_AuthInfo_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *User_AuthInfo_FieldMask) FilterInputFields() *User_AuthInfo_FieldMask {
	result := &User_AuthInfo_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *User_AuthInfo_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *User_AuthInfo_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]UserAuthInfo_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseUserAuthInfo_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask User_AuthInfo_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *User_AuthInfo_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *User_AuthInfo_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask User_AuthInfo_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *User_AuthInfo_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *User_AuthInfo_FieldMask) AppendPath(path UserAuthInfo_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *User_AuthInfo_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(UserAuthInfo_FieldPath))
}

func (fieldMask *User_AuthInfo_FieldMask) GetPaths() []UserAuthInfo_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *User_AuthInfo_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *User_AuthInfo_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseUserAuthInfo_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *User_AuthInfo_FieldMask) Set(target, source *User_AuthInfo) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *User_AuthInfo_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*User_AuthInfo), source.(*User_AuthInfo))
}

func (fieldMask *User_AuthInfo_FieldMask) Project(source *User_AuthInfo) *User_AuthInfo {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &User_AuthInfo{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *UserAuthInfo_FieldTerminalPath:
			switch tp.selector {
			case UserAuthInfo_FieldPathSelectorProvider:
				result.Provider = source.Provider
			case UserAuthInfo_FieldPathSelectorId:
				result.Id = source.Id
			}
		}
	}
	return result
}

func (fieldMask *User_AuthInfo_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*User_AuthInfo))
}

func (fieldMask *User_AuthInfo_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
