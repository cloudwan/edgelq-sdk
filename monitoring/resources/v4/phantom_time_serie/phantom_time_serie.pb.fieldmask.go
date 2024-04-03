// Code generated by protoc-gen-goten-object
// File: edgelq/monitoring/proto/v4/phantom_time_serie.proto
// DO NOT EDIT!!!

package phantom_time_serie

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
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/common"
	metric_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/metric_descriptor"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
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
	_ = &common.LabelDescriptor{}
	_ = &metric_descriptor.MetricDescriptor{}
	_ = &project.Project{}
	_ = &meta.Meta{}
)

type PhantomTimeSerie_FieldMask struct {
	Paths []PhantomTimeSerie_FieldPath
}

func FullPhantomTimeSerie_FieldMask() *PhantomTimeSerie_FieldMask {
	res := &PhantomTimeSerie_FieldMask{}
	res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorName})
	res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorKey})
	res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorProject})
	res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorMetric})
	res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorResource})
	res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorMetricKind})
	res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorValueType})
	res.Paths = append(res.Paths, &PhantomTimeSerie_FieldTerminalPath{selector: PhantomTimeSerie_FieldPathSelectorValue})
	return res
}

func (fieldMask *PhantomTimeSerie_FieldMask) String() string {
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
func (fieldMask *PhantomTimeSerie_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *PhantomTimeSerie_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParsePhantomTimeSerie_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *PhantomTimeSerie_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 9)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*PhantomTimeSerie_FieldTerminalPath); ok {
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

func (fieldMask *PhantomTimeSerie_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParsePhantomTimeSerie_FieldPath(raw)
	})
}

func (fieldMask *PhantomTimeSerie_FieldMask) ProtoMessage() {}

func (fieldMask *PhantomTimeSerie_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *PhantomTimeSerie_FieldMask) Subtract(other *PhantomTimeSerie_FieldMask) *PhantomTimeSerie_FieldMask {
	result := &PhantomTimeSerie_FieldMask{}
	removedSelectors := make([]bool, 9)
	otherSubMasks := map[PhantomTimeSerie_FieldPathSelector]gotenobject.FieldMask{
		PhantomTimeSerie_FieldPathSelectorMetadata: &meta.Meta_FieldMask{},
		PhantomTimeSerie_FieldPathSelectorMetric:   &common.Metric_FieldMask{},
		PhantomTimeSerie_FieldPathSelectorResource: &common.MonitoredResource_FieldMask{},
		PhantomTimeSerie_FieldPathSelectorValue:    &common.TypedValue_FieldMask{},
	}
	mySubMasks := map[PhantomTimeSerie_FieldPathSelector]gotenobject.FieldMask{
		PhantomTimeSerie_FieldPathSelectorMetadata: &meta.Meta_FieldMask{},
		PhantomTimeSerie_FieldPathSelectorMetric:   &common.Metric_FieldMask{},
		PhantomTimeSerie_FieldPathSelectorResource: &common.MonitoredResource_FieldMask{},
		PhantomTimeSerie_FieldPathSelectorValue:    &common.TypedValue_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *PhantomTimeSerie_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *PhantomTimeSerie_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*PhantomTimeSerie_FieldTerminalPath); ok {
					switch tp.selector {
					case PhantomTimeSerie_FieldPathSelectorMetadata:
						mySubMasks[PhantomTimeSerie_FieldPathSelectorMetadata] = meta.FullMeta_FieldMask()
					case PhantomTimeSerie_FieldPathSelectorMetric:
						mySubMasks[PhantomTimeSerie_FieldPathSelectorMetric] = common.FullMetric_FieldMask()
					case PhantomTimeSerie_FieldPathSelectorResource:
						mySubMasks[PhantomTimeSerie_FieldPathSelectorResource] = common.FullMonitoredResource_FieldMask()
					case PhantomTimeSerie_FieldPathSelectorValue:
						mySubMasks[PhantomTimeSerie_FieldPathSelectorValue] = common.FullTypedValue_FieldMask()
					}
				} else if tp, ok := path.(*PhantomTimeSerie_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &PhantomTimeSerie_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *PhantomTimeSerie_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*PhantomTimeSerie_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *PhantomTimeSerie_FieldMask) FilterInputFields() *PhantomTimeSerie_FieldMask {
	result := &PhantomTimeSerie_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case PhantomTimeSerie_FieldPathSelectorKey:
		case PhantomTimeSerie_FieldPathSelectorProject:
		case PhantomTimeSerie_FieldPathSelectorMetricKind:
		case PhantomTimeSerie_FieldPathSelectorValueType:
		case PhantomTimeSerie_FieldPathSelectorMetadata:
			if _, ok := path.(*PhantomTimeSerie_FieldTerminalPath); ok {
				for _, subpath := range meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &PhantomTimeSerie_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*PhantomTimeSerie_FieldSubPath); ok {
				selectedMask := &meta.Meta_FieldMask{
					Paths: []meta.Meta_FieldPath{sub.subPath.(meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &PhantomTimeSerie_FieldSubPath{selector: PhantomTimeSerie_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *PhantomTimeSerie_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *PhantomTimeSerie_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]PhantomTimeSerie_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParsePhantomTimeSerie_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask PhantomTimeSerie_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *PhantomTimeSerie_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *PhantomTimeSerie_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask PhantomTimeSerie_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *PhantomTimeSerie_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *PhantomTimeSerie_FieldMask) AppendPath(path PhantomTimeSerie_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *PhantomTimeSerie_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(PhantomTimeSerie_FieldPath))
}

func (fieldMask *PhantomTimeSerie_FieldMask) GetPaths() []PhantomTimeSerie_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *PhantomTimeSerie_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *PhantomTimeSerie_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParsePhantomTimeSerie_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *PhantomTimeSerie_FieldMask) Set(target, source *PhantomTimeSerie) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *PhantomTimeSerie_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*PhantomTimeSerie), source.(*PhantomTimeSerie))
}

func (fieldMask *PhantomTimeSerie_FieldMask) Project(source *PhantomTimeSerie) *PhantomTimeSerie {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &PhantomTimeSerie{}
	metadataMask := &meta.Meta_FieldMask{}
	wholeMetadataAccepted := false
	metricMask := &common.Metric_FieldMask{}
	wholeMetricAccepted := false
	resourceMask := &common.MonitoredResource_FieldMask{}
	wholeResourceAccepted := false
	valueMask := &common.TypedValue_FieldMask{}
	wholeValueAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *PhantomTimeSerie_FieldTerminalPath:
			switch tp.selector {
			case PhantomTimeSerie_FieldPathSelectorName:
				result.Name = source.Name
			case PhantomTimeSerie_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case PhantomTimeSerie_FieldPathSelectorKey:
				result.Key = source.Key
			case PhantomTimeSerie_FieldPathSelectorProject:
				result.Project = source.Project
			case PhantomTimeSerie_FieldPathSelectorMetric:
				result.Metric = source.Metric
				wholeMetricAccepted = true
			case PhantomTimeSerie_FieldPathSelectorResource:
				result.Resource = source.Resource
				wholeResourceAccepted = true
			case PhantomTimeSerie_FieldPathSelectorMetricKind:
				result.MetricKind = source.MetricKind
			case PhantomTimeSerie_FieldPathSelectorValueType:
				result.ValueType = source.ValueType
			case PhantomTimeSerie_FieldPathSelectorValue:
				result.Value = source.Value
				wholeValueAccepted = true
			}
		case *PhantomTimeSerie_FieldSubPath:
			switch tp.selector {
			case PhantomTimeSerie_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(meta.Meta_FieldPath))
			case PhantomTimeSerie_FieldPathSelectorMetric:
				metricMask.AppendPath(tp.subPath.(common.Metric_FieldPath))
			case PhantomTimeSerie_FieldPathSelectorResource:
				resourceMask.AppendPath(tp.subPath.(common.MonitoredResource_FieldPath))
			case PhantomTimeSerie_FieldPathSelectorValue:
				valueMask.AppendPath(tp.subPath.(common.TypedValue_FieldPath))
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	if wholeMetricAccepted == false && len(metricMask.Paths) > 0 {
		result.Metric = metricMask.Project(source.GetMetric())
	}
	if wholeResourceAccepted == false && len(resourceMask.Paths) > 0 {
		result.Resource = resourceMask.Project(source.GetResource())
	}
	if wholeValueAccepted == false && len(valueMask.Paths) > 0 {
		result.Value = valueMask.Project(source.GetValue())
	}
	return result
}

func (fieldMask *PhantomTimeSerie_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*PhantomTimeSerie))
}

func (fieldMask *PhantomTimeSerie_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}