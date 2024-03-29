// Code generated by protoc-gen-goten-object
// File: edgelq/monitoring/proto/v3/time_serie.proto
// DO NOT EDIT!!!

package time_serie

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
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/common"
	metric_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/metric_descriptor"
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
)

type Point_FieldMask struct {
	Paths []Point_FieldPath
}

func FullPoint_FieldMask() *Point_FieldMask {
	res := &Point_FieldMask{}
	res.Paths = append(res.Paths, &Point_FieldTerminalPath{selector: Point_FieldPathSelectorInterval})
	res.Paths = append(res.Paths, &Point_FieldTerminalPath{selector: Point_FieldPathSelectorValue})
	res.Paths = append(res.Paths, &Point_FieldTerminalPath{selector: Point_FieldPathSelectorAggregation})
	return res
}

func (fieldMask *Point_FieldMask) String() string {
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
func (fieldMask *Point_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *Point_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParsePoint_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Point_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 3)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*Point_FieldTerminalPath); ok {
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

func (fieldMask *Point_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParsePoint_FieldPath(raw)
	})
}

func (fieldMask *Point_FieldMask) ProtoMessage() {}

func (fieldMask *Point_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Point_FieldMask) Subtract(other *Point_FieldMask) *Point_FieldMask {
	result := &Point_FieldMask{}
	removedSelectors := make([]bool, 3)
	otherSubMasks := map[Point_FieldPathSelector]gotenobject.FieldMask{
		Point_FieldPathSelectorInterval:    &common.TimeInterval_FieldMask{},
		Point_FieldPathSelectorValue:       &common.TypedValue_FieldMask{},
		Point_FieldPathSelectorAggregation: &common.Aggregation_FieldMask{},
	}
	mySubMasks := map[Point_FieldPathSelector]gotenobject.FieldMask{
		Point_FieldPathSelectorInterval:    &common.TimeInterval_FieldMask{},
		Point_FieldPathSelectorValue:       &common.TypedValue_FieldMask{},
		Point_FieldPathSelectorAggregation: &common.Aggregation_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *Point_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *Point_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*Point_FieldTerminalPath); ok {
					switch tp.selector {
					case Point_FieldPathSelectorInterval:
						mySubMasks[Point_FieldPathSelectorInterval] = common.FullTimeInterval_FieldMask()
					case Point_FieldPathSelectorValue:
						mySubMasks[Point_FieldPathSelectorValue] = common.FullTypedValue_FieldMask()
					case Point_FieldPathSelectorAggregation:
						mySubMasks[Point_FieldPathSelectorAggregation] = common.FullAggregation_FieldMask()
					}
				} else if tp, ok := path.(*Point_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &Point_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *Point_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Point_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Point_FieldMask) FilterInputFields() *Point_FieldMask {
	result := &Point_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Point_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Point_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]Point_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParsePoint_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Point_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Point_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Point_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Point_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Point_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Point_FieldMask) AppendPath(path Point_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Point_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(Point_FieldPath))
}

func (fieldMask *Point_FieldMask) GetPaths() []Point_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Point_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Point_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParsePoint_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Point_FieldMask) Set(target, source *Point) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Point_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Point), source.(*Point))
}

func (fieldMask *Point_FieldMask) Project(source *Point) *Point {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Point{}
	intervalMask := &common.TimeInterval_FieldMask{}
	wholeIntervalAccepted := false
	valueMask := &common.TypedValue_FieldMask{}
	wholeValueAccepted := false
	aggregationMask := &common.Aggregation_FieldMask{}
	wholeAggregationAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *Point_FieldTerminalPath:
			switch tp.selector {
			case Point_FieldPathSelectorInterval:
				result.Interval = source.Interval
				wholeIntervalAccepted = true
			case Point_FieldPathSelectorValue:
				result.Value = source.Value
				wholeValueAccepted = true
			case Point_FieldPathSelectorAggregation:
				result.Aggregation = source.Aggregation
				wholeAggregationAccepted = true
			}
		case *Point_FieldSubPath:
			switch tp.selector {
			case Point_FieldPathSelectorInterval:
				intervalMask.AppendPath(tp.subPath.(common.TimeInterval_FieldPath))
			case Point_FieldPathSelectorValue:
				valueMask.AppendPath(tp.subPath.(common.TypedValue_FieldPath))
			case Point_FieldPathSelectorAggregation:
				aggregationMask.AppendPath(tp.subPath.(common.Aggregation_FieldPath))
			}
		}
	}
	if wholeIntervalAccepted == false && len(intervalMask.Paths) > 0 {
		result.Interval = intervalMask.Project(source.GetInterval())
	}
	if wholeValueAccepted == false && len(valueMask.Paths) > 0 {
		result.Value = valueMask.Project(source.GetValue())
	}
	if wholeAggregationAccepted == false && len(aggregationMask.Paths) > 0 {
		result.Aggregation = aggregationMask.Project(source.GetAggregation())
	}
	return result
}

func (fieldMask *Point_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Point))
}

func (fieldMask *Point_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type TimeSerie_FieldMask struct {
	Paths []TimeSerie_FieldPath
}

func FullTimeSerie_FieldMask() *TimeSerie_FieldMask {
	res := &TimeSerie_FieldMask{}
	res.Paths = append(res.Paths, &TimeSerie_FieldTerminalPath{selector: TimeSerie_FieldPathSelectorKey})
	res.Paths = append(res.Paths, &TimeSerie_FieldTerminalPath{selector: TimeSerie_FieldPathSelectorProject})
	res.Paths = append(res.Paths, &TimeSerie_FieldTerminalPath{selector: TimeSerie_FieldPathSelectorRegion})
	res.Paths = append(res.Paths, &TimeSerie_FieldTerminalPath{selector: TimeSerie_FieldPathSelectorMetric})
	res.Paths = append(res.Paths, &TimeSerie_FieldTerminalPath{selector: TimeSerie_FieldPathSelectorResource})
	res.Paths = append(res.Paths, &TimeSerie_FieldTerminalPath{selector: TimeSerie_FieldPathSelectorMetricKind})
	res.Paths = append(res.Paths, &TimeSerie_FieldTerminalPath{selector: TimeSerie_FieldPathSelectorValueType})
	res.Paths = append(res.Paths, &TimeSerie_FieldTerminalPath{selector: TimeSerie_FieldPathSelectorPoints})
	return res
}

func (fieldMask *TimeSerie_FieldMask) String() string {
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
func (fieldMask *TimeSerie_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *TimeSerie_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseTimeSerie_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *TimeSerie_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 8)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*TimeSerie_FieldTerminalPath); ok {
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

func (fieldMask *TimeSerie_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseTimeSerie_FieldPath(raw)
	})
}

func (fieldMask *TimeSerie_FieldMask) ProtoMessage() {}

func (fieldMask *TimeSerie_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *TimeSerie_FieldMask) Subtract(other *TimeSerie_FieldMask) *TimeSerie_FieldMask {
	result := &TimeSerie_FieldMask{}
	removedSelectors := make([]bool, 8)
	otherSubMasks := map[TimeSerie_FieldPathSelector]gotenobject.FieldMask{
		TimeSerie_FieldPathSelectorMetric:   &common.Metric_FieldMask{},
		TimeSerie_FieldPathSelectorResource: &common.MonitoredResource_FieldMask{},
		TimeSerie_FieldPathSelectorPoints:   &Point_FieldMask{},
	}
	mySubMasks := map[TimeSerie_FieldPathSelector]gotenobject.FieldMask{
		TimeSerie_FieldPathSelectorMetric:   &common.Metric_FieldMask{},
		TimeSerie_FieldPathSelectorResource: &common.MonitoredResource_FieldMask{},
		TimeSerie_FieldPathSelectorPoints:   &Point_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *TimeSerie_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *TimeSerie_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*TimeSerie_FieldTerminalPath); ok {
					switch tp.selector {
					case TimeSerie_FieldPathSelectorMetric:
						mySubMasks[TimeSerie_FieldPathSelectorMetric] = common.FullMetric_FieldMask()
					case TimeSerie_FieldPathSelectorResource:
						mySubMasks[TimeSerie_FieldPathSelectorResource] = common.FullMonitoredResource_FieldMask()
					case TimeSerie_FieldPathSelectorPoints:
						mySubMasks[TimeSerie_FieldPathSelectorPoints] = FullPoint_FieldMask()
					}
				} else if tp, ok := path.(*TimeSerie_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &TimeSerie_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *TimeSerie_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*TimeSerie_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *TimeSerie_FieldMask) FilterInputFields() *TimeSerie_FieldMask {
	result := &TimeSerie_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *TimeSerie_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *TimeSerie_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]TimeSerie_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseTimeSerie_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask TimeSerie_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *TimeSerie_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *TimeSerie_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask TimeSerie_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *TimeSerie_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *TimeSerie_FieldMask) AppendPath(path TimeSerie_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *TimeSerie_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(TimeSerie_FieldPath))
}

func (fieldMask *TimeSerie_FieldMask) GetPaths() []TimeSerie_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *TimeSerie_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *TimeSerie_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseTimeSerie_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *TimeSerie_FieldMask) Set(target, source *TimeSerie) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *TimeSerie_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*TimeSerie), source.(*TimeSerie))
}

func (fieldMask *TimeSerie_FieldMask) Project(source *TimeSerie) *TimeSerie {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &TimeSerie{}
	metricMask := &common.Metric_FieldMask{}
	wholeMetricAccepted := false
	resourceMask := &common.MonitoredResource_FieldMask{}
	wholeResourceAccepted := false
	pointsMask := &Point_FieldMask{}
	wholePointsAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *TimeSerie_FieldTerminalPath:
			switch tp.selector {
			case TimeSerie_FieldPathSelectorKey:
				result.Key = source.Key
			case TimeSerie_FieldPathSelectorProject:
				result.Project = source.Project
			case TimeSerie_FieldPathSelectorRegion:
				result.Region = source.Region
			case TimeSerie_FieldPathSelectorMetric:
				result.Metric = source.Metric
				wholeMetricAccepted = true
			case TimeSerie_FieldPathSelectorResource:
				result.Resource = source.Resource
				wholeResourceAccepted = true
			case TimeSerie_FieldPathSelectorMetricKind:
				result.MetricKind = source.MetricKind
			case TimeSerie_FieldPathSelectorValueType:
				result.ValueType = source.ValueType
			case TimeSerie_FieldPathSelectorPoints:
				result.Points = source.Points
				wholePointsAccepted = true
			}
		case *TimeSerie_FieldSubPath:
			switch tp.selector {
			case TimeSerie_FieldPathSelectorMetric:
				metricMask.AppendPath(tp.subPath.(common.Metric_FieldPath))
			case TimeSerie_FieldPathSelectorResource:
				resourceMask.AppendPath(tp.subPath.(common.MonitoredResource_FieldPath))
			case TimeSerie_FieldPathSelectorPoints:
				pointsMask.AppendPath(tp.subPath.(Point_FieldPath))
			}
		}
	}
	if wholeMetricAccepted == false && len(metricMask.Paths) > 0 {
		result.Metric = metricMask.Project(source.GetMetric())
	}
	if wholeResourceAccepted == false && len(resourceMask.Paths) > 0 {
		result.Resource = resourceMask.Project(source.GetResource())
	}
	if wholePointsAccepted == false && len(pointsMask.Paths) > 0 {
		for _, sourceItem := range source.GetPoints() {
			result.Points = append(result.Points, pointsMask.Project(sourceItem))
		}
	}
	return result
}

func (fieldMask *TimeSerie_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*TimeSerie))
}

func (fieldMask *TimeSerie_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type BulkTimeSeries_FieldMask struct {
	Paths []BulkTimeSeries_FieldPath
}

func FullBulkTimeSeries_FieldMask() *BulkTimeSeries_FieldMask {
	res := &BulkTimeSeries_FieldMask{}
	res.Paths = append(res.Paths, &BulkTimeSeries_FieldTerminalPath{selector: BulkTimeSeries_FieldPathSelectorTimeSeries})
	res.Paths = append(res.Paths, &BulkTimeSeries_FieldTerminalPath{selector: BulkTimeSeries_FieldPathSelectorPhantomFlag})
	return res
}

func (fieldMask *BulkTimeSeries_FieldMask) String() string {
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
func (fieldMask *BulkTimeSeries_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *BulkTimeSeries_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseBulkTimeSeries_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *BulkTimeSeries_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*BulkTimeSeries_FieldTerminalPath); ok {
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

func (fieldMask *BulkTimeSeries_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseBulkTimeSeries_FieldPath(raw)
	})
}

func (fieldMask *BulkTimeSeries_FieldMask) ProtoMessage() {}

func (fieldMask *BulkTimeSeries_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *BulkTimeSeries_FieldMask) Subtract(other *BulkTimeSeries_FieldMask) *BulkTimeSeries_FieldMask {
	result := &BulkTimeSeries_FieldMask{}
	removedSelectors := make([]bool, 2)
	otherSubMasks := map[BulkTimeSeries_FieldPathSelector]gotenobject.FieldMask{
		BulkTimeSeries_FieldPathSelectorTimeSeries: &TimeSerie_FieldMask{},
	}
	mySubMasks := map[BulkTimeSeries_FieldPathSelector]gotenobject.FieldMask{
		BulkTimeSeries_FieldPathSelectorTimeSeries: &TimeSerie_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *BulkTimeSeries_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *BulkTimeSeries_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*BulkTimeSeries_FieldTerminalPath); ok {
					switch tp.selector {
					case BulkTimeSeries_FieldPathSelectorTimeSeries:
						mySubMasks[BulkTimeSeries_FieldPathSelectorTimeSeries] = FullTimeSerie_FieldMask()
					}
				} else if tp, ok := path.(*BulkTimeSeries_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &BulkTimeSeries_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *BulkTimeSeries_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*BulkTimeSeries_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *BulkTimeSeries_FieldMask) FilterInputFields() *BulkTimeSeries_FieldMask {
	result := &BulkTimeSeries_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *BulkTimeSeries_FieldMask) ToProtoFieldMask() *googlefieldmaskpb.FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *BulkTimeSeries_FieldMask) FromProtoFieldMask(protoFieldMask *googlefieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]BulkTimeSeries_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseBulkTimeSeries_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask BulkTimeSeries_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *BulkTimeSeries_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *BulkTimeSeries_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask BulkTimeSeries_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *BulkTimeSeries_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *BulkTimeSeries_FieldMask) AppendPath(path BulkTimeSeries_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *BulkTimeSeries_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(BulkTimeSeries_FieldPath))
}

func (fieldMask *BulkTimeSeries_FieldMask) GetPaths() []BulkTimeSeries_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *BulkTimeSeries_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *BulkTimeSeries_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseBulkTimeSeries_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *BulkTimeSeries_FieldMask) Set(target, source *BulkTimeSeries) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *BulkTimeSeries_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*BulkTimeSeries), source.(*BulkTimeSeries))
}

func (fieldMask *BulkTimeSeries_FieldMask) Project(source *BulkTimeSeries) *BulkTimeSeries {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &BulkTimeSeries{}
	timeSeriesMask := &TimeSerie_FieldMask{}
	wholeTimeSeriesAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *BulkTimeSeries_FieldTerminalPath:
			switch tp.selector {
			case BulkTimeSeries_FieldPathSelectorTimeSeries:
				result.TimeSeries = source.TimeSeries
				wholeTimeSeriesAccepted = true
			case BulkTimeSeries_FieldPathSelectorPhantomFlag:
				result.PhantomFlag = source.PhantomFlag
			}
		case *BulkTimeSeries_FieldSubPath:
			switch tp.selector {
			case BulkTimeSeries_FieldPathSelectorTimeSeries:
				timeSeriesMask.AppendPath(tp.subPath.(TimeSerie_FieldPath))
			}
		}
	}
	if wholeTimeSeriesAccepted == false && len(timeSeriesMask.Paths) > 0 {
		for _, sourceItem := range source.GetTimeSeries() {
			result.TimeSeries = append(result.TimeSeries, timeSeriesMask.Project(sourceItem))
		}
	}
	return result
}

func (fieldMask *BulkTimeSeries_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*BulkTimeSeries))
}

func (fieldMask *BulkTimeSeries_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
