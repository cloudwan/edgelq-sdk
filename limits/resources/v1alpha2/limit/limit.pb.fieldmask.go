// Code generated by protoc-gen-goten-object
// File: edgelq/limits/proto/v1alpha2/limit.proto
// DO NOT EDIT!!!

package limit

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
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	limit_pool "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/limit_pool"
	meta_resource "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/resource"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
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
	_ = &iam_project.Project{}
	_ = &limit_pool.LimitPool{}
	_ = &meta_resource.Resource{}
	_ = &meta_service.Service{}
)

type Limit_FieldMask struct {
	Paths []Limit_FieldPath
}

func FullLimit_FieldMask() *Limit_FieldMask {
	res := &Limit_FieldMask{}
	res.Paths = append(res.Paths, &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorName})
	res.Paths = append(res.Paths, &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorService})
	res.Paths = append(res.Paths, &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorResource})
	res.Paths = append(res.Paths, &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorRegion})
	res.Paths = append(res.Paths, &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorConfiguredLimit})
	res.Paths = append(res.Paths, &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorActiveLimit})
	res.Paths = append(res.Paths, &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorUsage})
	res.Paths = append(res.Paths, &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorSource})
	res.Paths = append(res.Paths, &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorMetadata})
	return res
}

func (fieldMask *Limit_FieldMask) String() string {
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
func (fieldMask *Limit_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *Limit_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseLimit_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Limit_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 9)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*Limit_FieldTerminalPath); ok {
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

func (fieldMask *Limit_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseLimit_FieldPath(raw)
	})
}

func (fieldMask *Limit_FieldMask) ProtoMessage() {}

func (fieldMask *Limit_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Limit_FieldMask) Subtract(other *Limit_FieldMask) *Limit_FieldMask {
	result := &Limit_FieldMask{}
	removedSelectors := make([]bool, 9)
	otherSubMasks := map[Limit_FieldPathSelector]gotenobject.FieldMask{
		Limit_FieldPathSelectorMetadata: &ntt_meta.Meta_FieldMask{},
	}
	mySubMasks := map[Limit_FieldPathSelector]gotenobject.FieldMask{
		Limit_FieldPathSelectorMetadata: &ntt_meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *Limit_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *Limit_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*Limit_FieldTerminalPath); ok {
					switch tp.selector {
					case Limit_FieldPathSelectorMetadata:
						mySubMasks[Limit_FieldPathSelectorMetadata] = ntt_meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*Limit_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &Limit_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *Limit_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Limit_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Limit_FieldMask) FilterInputFields() *Limit_FieldMask {
	result := &Limit_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case Limit_FieldPathSelectorService:
		case Limit_FieldPathSelectorResource:
		case Limit_FieldPathSelectorRegion:
		case Limit_FieldPathSelectorConfiguredLimit:
		case Limit_FieldPathSelectorSource:
		case Limit_FieldPathSelectorMetadata:
			if _, ok := path.(*Limit_FieldTerminalPath); ok {
				for _, subpath := range ntt_meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &Limit_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*Limit_FieldSubPath); ok {
				selectedMask := &ntt_meta.Meta_FieldMask{
					Paths: []ntt_meta.Meta_FieldPath{sub.subPath.(ntt_meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &Limit_FieldSubPath{selector: Limit_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Limit_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Limit_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]Limit_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseLimit_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Limit_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Limit_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Limit_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Limit_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Limit_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Limit_FieldMask) AppendPath(path Limit_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Limit_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(Limit_FieldPath))
}

func (fieldMask *Limit_FieldMask) GetPaths() []Limit_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Limit_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Limit_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseLimit_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Limit_FieldMask) Set(target, source *Limit) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Limit_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Limit), source.(*Limit))
}

func (fieldMask *Limit_FieldMask) Project(source *Limit) *Limit {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Limit{}
	metadataMask := &ntt_meta.Meta_FieldMask{}
	wholeMetadataAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *Limit_FieldTerminalPath:
			switch tp.selector {
			case Limit_FieldPathSelectorName:
				result.Name = source.Name
			case Limit_FieldPathSelectorService:
				result.Service = source.Service
			case Limit_FieldPathSelectorResource:
				result.Resource = source.Resource
			case Limit_FieldPathSelectorRegion:
				result.Region = source.Region
			case Limit_FieldPathSelectorConfiguredLimit:
				result.ConfiguredLimit = source.ConfiguredLimit
			case Limit_FieldPathSelectorActiveLimit:
				result.ActiveLimit = source.ActiveLimit
			case Limit_FieldPathSelectorUsage:
				result.Usage = source.Usage
			case Limit_FieldPathSelectorSource:
				result.Source = source.Source
			case Limit_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			}
		case *Limit_FieldSubPath:
			switch tp.selector {
			case Limit_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(ntt_meta.Meta_FieldPath))
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	return result
}

func (fieldMask *Limit_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Limit))
}

func (fieldMask *Limit_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}