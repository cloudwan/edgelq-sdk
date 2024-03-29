// Code generated by protoc-gen-goten-object
// File: edgelq/ztp/proto/v1/tpm_attestation_cert.proto
// DO NOT EDIT!!!

package tpm_attestation_cert

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
	project "github.com/cloudwan/edgelq-sdk/ztp/resources/v1/project"
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
	_ = &project.Project{}
	_ = &meta.Meta{}
)

type TpmAttestationCert_FieldMask struct {
	Paths []TpmAttestationCert_FieldPath
}

func FullTpmAttestationCert_FieldMask() *TpmAttestationCert_FieldMask {
	res := &TpmAttestationCert_FieldMask{}
	res.Paths = append(res.Paths, &TpmAttestationCert_FieldTerminalPath{selector: TpmAttestationCert_FieldPathSelectorName})
	res.Paths = append(res.Paths, &TpmAttestationCert_FieldTerminalPath{selector: TpmAttestationCert_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &TpmAttestationCert_FieldTerminalPath{selector: TpmAttestationCert_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &TpmAttestationCert_FieldTerminalPath{selector: TpmAttestationCert_FieldPathSelectorManufacturer})
	res.Paths = append(res.Paths, &TpmAttestationCert_FieldTerminalPath{selector: TpmAttestationCert_FieldPathSelectorProductName})
	res.Paths = append(res.Paths, &TpmAttestationCert_FieldTerminalPath{selector: TpmAttestationCert_FieldPathSelectorTpmManufacturerCaCert})
	res.Paths = append(res.Paths, &TpmAttestationCert_FieldTerminalPath{selector: TpmAttestationCert_FieldPathSelectorIdevidIssuerCaCert})
	res.Paths = append(res.Paths, &TpmAttestationCert_FieldTerminalPath{selector: TpmAttestationCert_FieldPathSelectorLdevidIssuerCaCert})
	return res
}

func (fieldMask *TpmAttestationCert_FieldMask) String() string {
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
func (fieldMask *TpmAttestationCert_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *TpmAttestationCert_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseTpmAttestationCert_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *TpmAttestationCert_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 8)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*TpmAttestationCert_FieldTerminalPath); ok {
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

func (fieldMask *TpmAttestationCert_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseTpmAttestationCert_FieldPath(raw)
	})
}

func (fieldMask *TpmAttestationCert_FieldMask) ProtoMessage() {}

func (fieldMask *TpmAttestationCert_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *TpmAttestationCert_FieldMask) Subtract(other *TpmAttestationCert_FieldMask) *TpmAttestationCert_FieldMask {
	result := &TpmAttestationCert_FieldMask{}
	removedSelectors := make([]bool, 8)
	otherSubMasks := map[TpmAttestationCert_FieldPathSelector]gotenobject.FieldMask{
		TpmAttestationCert_FieldPathSelectorMetadata: &meta.Meta_FieldMask{},
	}
	mySubMasks := map[TpmAttestationCert_FieldPathSelector]gotenobject.FieldMask{
		TpmAttestationCert_FieldPathSelectorMetadata: &meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *TpmAttestationCert_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *TpmAttestationCert_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*TpmAttestationCert_FieldTerminalPath); ok {
					switch tp.selector {
					case TpmAttestationCert_FieldPathSelectorMetadata:
						mySubMasks[TpmAttestationCert_FieldPathSelectorMetadata] = meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*TpmAttestationCert_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &TpmAttestationCert_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *TpmAttestationCert_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*TpmAttestationCert_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *TpmAttestationCert_FieldMask) FilterInputFields() *TpmAttestationCert_FieldMask {
	result := &TpmAttestationCert_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case TpmAttestationCert_FieldPathSelectorMetadata:
			if _, ok := path.(*TpmAttestationCert_FieldTerminalPath); ok {
				for _, subpath := range meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &TpmAttestationCert_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*TpmAttestationCert_FieldSubPath); ok {
				selectedMask := &meta.Meta_FieldMask{
					Paths: []meta.Meta_FieldPath{sub.subPath.(meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &TpmAttestationCert_FieldSubPath{selector: TpmAttestationCert_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *TpmAttestationCert_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *TpmAttestationCert_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]TpmAttestationCert_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseTpmAttestationCert_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask TpmAttestationCert_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *TpmAttestationCert_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *TpmAttestationCert_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask TpmAttestationCert_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *TpmAttestationCert_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *TpmAttestationCert_FieldMask) AppendPath(path TpmAttestationCert_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *TpmAttestationCert_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(TpmAttestationCert_FieldPath))
}

func (fieldMask *TpmAttestationCert_FieldMask) GetPaths() []TpmAttestationCert_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *TpmAttestationCert_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *TpmAttestationCert_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseTpmAttestationCert_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *TpmAttestationCert_FieldMask) Set(target, source *TpmAttestationCert) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *TpmAttestationCert_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*TpmAttestationCert), source.(*TpmAttestationCert))
}

func (fieldMask *TpmAttestationCert_FieldMask) Project(source *TpmAttestationCert) *TpmAttestationCert {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &TpmAttestationCert{}
	metadataMask := &meta.Meta_FieldMask{}
	wholeMetadataAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *TpmAttestationCert_FieldTerminalPath:
			switch tp.selector {
			case TpmAttestationCert_FieldPathSelectorName:
				result.Name = source.Name
			case TpmAttestationCert_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case TpmAttestationCert_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case TpmAttestationCert_FieldPathSelectorManufacturer:
				result.Manufacturer = source.Manufacturer
			case TpmAttestationCert_FieldPathSelectorProductName:
				result.ProductName = source.ProductName
			case TpmAttestationCert_FieldPathSelectorTpmManufacturerCaCert:
				result.TpmManufacturerCaCert = source.TpmManufacturerCaCert
			case TpmAttestationCert_FieldPathSelectorIdevidIssuerCaCert:
				result.IdevidIssuerCaCert = source.IdevidIssuerCaCert
			case TpmAttestationCert_FieldPathSelectorLdevidIssuerCaCert:
				result.LdevidIssuerCaCert = source.LdevidIssuerCaCert
			}
		case *TpmAttestationCert_FieldSubPath:
			switch tp.selector {
			case TpmAttestationCert_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(meta.Meta_FieldPath))
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	return result
}

func (fieldMask *TpmAttestationCert_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*TpmAttestationCert))
}

func (fieldMask *TpmAttestationCert_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
