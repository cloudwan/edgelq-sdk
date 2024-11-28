// Code generated by protoc-gen-goten-object
// File: edgelq/logging/proto/v1alpha2/log_descriptor.proto
// DO NOT EDIT!!!

package log_descriptor

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	"github.com/cloudwan/goten-sdk/runtime/strcase"
)

// proto imports
import (
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	common "github.com/cloudwan/edgelq-sdk/logging/resources/v1alpha2/common"
	meta "github.com/cloudwan/goten-sdk/types/meta"
)

// ensure the imports are used
var (
	_ = new(json.Marshaler)
	_ = new(fmt.Stringer)
	_ = reflect.DeepEqual
	_ = strings.Builder{}
	_ = time.Second

	_ = strcase.ToLowerCamel
	_ = codes.NotFound
	_ = status.Status{}
	_ = protojson.UnmarshalOptions{}
	_ = new(proto.Message)
	_ = protoregistry.GlobalTypes

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &common.LabelDescriptor{}
	_ = &meta.Meta{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type LogDescriptor_FieldPath interface {
	gotenobject.FieldPath
	Selector() LogDescriptor_FieldPathSelector
	Get(source *LogDescriptor) []interface{}
	GetSingle(source *LogDescriptor) (interface{}, bool)
	ClearValue(item *LogDescriptor)

	// Those methods build corresponding LogDescriptor_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) LogDescriptor_FieldPathValue
	WithIArrayOfValues(values interface{}) LogDescriptor_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) LogDescriptor_FieldPathArrayItemValue
}

type LogDescriptor_FieldPathSelector int32

const (
	LogDescriptor_FieldPathSelectorName                 LogDescriptor_FieldPathSelector = 0
	LogDescriptor_FieldPathSelectorDisplayName          LogDescriptor_FieldPathSelector = 1
	LogDescriptor_FieldPathSelectorDescription          LogDescriptor_FieldPathSelector = 2
	LogDescriptor_FieldPathSelectorLabels               LogDescriptor_FieldPathSelector = 3
	LogDescriptor_FieldPathSelectorPromotedLabelKeySets LogDescriptor_FieldPathSelector = 4
	LogDescriptor_FieldPathSelectorMetadata             LogDescriptor_FieldPathSelector = 5
)

func (s LogDescriptor_FieldPathSelector) String() string {
	switch s {
	case LogDescriptor_FieldPathSelectorName:
		return "name"
	case LogDescriptor_FieldPathSelectorDisplayName:
		return "display_name"
	case LogDescriptor_FieldPathSelectorDescription:
		return "description"
	case LogDescriptor_FieldPathSelectorLabels:
		return "labels"
	case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
		return "promoted_label_key_sets"
	case LogDescriptor_FieldPathSelectorMetadata:
		return "metadata"
	default:
		panic(fmt.Sprintf("Invalid selector for LogDescriptor: %d", s))
	}
}

func BuildLogDescriptor_FieldPath(fp gotenobject.RawFieldPath) (LogDescriptor_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object LogDescriptor")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorName}, nil
		case "display_name", "displayName", "display-name":
			return &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorDisplayName}, nil
		case "description":
			return &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorDescription}, nil
		case "labels":
			return &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorLabels}, nil
		case "promoted_label_key_sets", "promotedLabelKeySets", "promoted-label-key-sets":
			return &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorPromotedLabelKeySets}, nil
		case "metadata":
			return &LogDescriptor_FieldTerminalPath{selector: LogDescriptor_FieldPathSelectorMetadata}, nil
		}
	} else {
		switch fp[0] {
		case "labels":
			if subpath, err := common.BuildLabelDescriptor_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &LogDescriptor_FieldSubPath{selector: LogDescriptor_FieldPathSelectorLabels, subPath: subpath}, nil
			}
		case "promoted_label_key_sets", "promotedLabelKeySets", "promoted-label-key-sets":
			if subpath, err := common.BuildLabelKeySet_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &LogDescriptor_FieldSubPath{selector: LogDescriptor_FieldPathSelectorPromotedLabelKeySets, subPath: subpath}, nil
			}
		case "metadata":
			if subpath, err := meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &LogDescriptor_FieldSubPath{selector: LogDescriptor_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object LogDescriptor", fp)
}

func ParseLogDescriptor_FieldPath(rawField string) (LogDescriptor_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildLogDescriptor_FieldPath(fp)
}

func MustParseLogDescriptor_FieldPath(rawField string) LogDescriptor_FieldPath {
	fp, err := ParseLogDescriptor_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type LogDescriptor_FieldTerminalPath struct {
	selector LogDescriptor_FieldPathSelector
}

var _ LogDescriptor_FieldPath = (*LogDescriptor_FieldTerminalPath)(nil)

func (fp *LogDescriptor_FieldTerminalPath) Selector() LogDescriptor_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *LogDescriptor_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *LogDescriptor_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source LogDescriptor
func (fp *LogDescriptor_FieldTerminalPath) Get(source *LogDescriptor) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case LogDescriptor_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case LogDescriptor_FieldPathSelectorDisplayName:
			values = append(values, source.DisplayName)
		case LogDescriptor_FieldPathSelectorDescription:
			values = append(values, source.Description)
		case LogDescriptor_FieldPathSelectorLabels:
			for _, value := range source.GetLabels() {
				values = append(values, value)
			}
		case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
			for _, value := range source.GetPromotedLabelKeySets() {
				values = append(values, value)
			}
		case LogDescriptor_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		default:
			panic(fmt.Sprintf("Invalid selector for LogDescriptor: %d", fp.selector))
		}
	}
	return
}

func (fp *LogDescriptor_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*LogDescriptor))
}

// GetSingle returns value pointed by specific field of from source LogDescriptor
func (fp *LogDescriptor_FieldTerminalPath) GetSingle(source *LogDescriptor) (interface{}, bool) {
	switch fp.selector {
	case LogDescriptor_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case LogDescriptor_FieldPathSelectorDisplayName:
		return source.GetDisplayName(), source != nil
	case LogDescriptor_FieldPathSelectorDescription:
		return source.GetDescription(), source != nil
	case LogDescriptor_FieldPathSelectorLabels:
		res := source.GetLabels()
		return res, res != nil
	case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
		res := source.GetPromotedLabelKeySets()
		return res, res != nil
	case LogDescriptor_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	default:
		panic(fmt.Sprintf("Invalid selector for LogDescriptor: %d", fp.selector))
	}
}

func (fp *LogDescriptor_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*LogDescriptor))
}

// GetDefault returns a default value of the field type
func (fp *LogDescriptor_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case LogDescriptor_FieldPathSelectorName:
		return (*Name)(nil)
	case LogDescriptor_FieldPathSelectorDisplayName:
		return ""
	case LogDescriptor_FieldPathSelectorDescription:
		return ""
	case LogDescriptor_FieldPathSelectorLabels:
		return ([]*common.LabelDescriptor)(nil)
	case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
		return ([]*common.LabelKeySet)(nil)
	case LogDescriptor_FieldPathSelectorMetadata:
		return (*meta.Meta)(nil)
	default:
		panic(fmt.Sprintf("Invalid selector for LogDescriptor: %d", fp.selector))
	}
}

func (fp *LogDescriptor_FieldTerminalPath) ClearValue(item *LogDescriptor) {
	if item != nil {
		switch fp.selector {
		case LogDescriptor_FieldPathSelectorName:
			item.Name = nil
		case LogDescriptor_FieldPathSelectorDisplayName:
			item.DisplayName = ""
		case LogDescriptor_FieldPathSelectorDescription:
			item.Description = ""
		case LogDescriptor_FieldPathSelectorLabels:
			item.Labels = nil
		case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
			item.PromotedLabelKeySets = nil
		case LogDescriptor_FieldPathSelectorMetadata:
			item.Metadata = nil
		default:
			panic(fmt.Sprintf("Invalid selector for LogDescriptor: %d", fp.selector))
		}
	}
}

func (fp *LogDescriptor_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*LogDescriptor))
}

// IsLeaf - whether field path is holds simple value
func (fp *LogDescriptor_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == LogDescriptor_FieldPathSelectorName ||
		fp.selector == LogDescriptor_FieldPathSelectorDisplayName ||
		fp.selector == LogDescriptor_FieldPathSelectorDescription
}

func (fp *LogDescriptor_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *LogDescriptor_FieldTerminalPath) WithIValue(value interface{}) LogDescriptor_FieldPathValue {
	switch fp.selector {
	case LogDescriptor_FieldPathSelectorName:
		return &LogDescriptor_FieldTerminalPathValue{LogDescriptor_FieldTerminalPath: *fp, value: value.(*Name)}
	case LogDescriptor_FieldPathSelectorDisplayName:
		return &LogDescriptor_FieldTerminalPathValue{LogDescriptor_FieldTerminalPath: *fp, value: value.(string)}
	case LogDescriptor_FieldPathSelectorDescription:
		return &LogDescriptor_FieldTerminalPathValue{LogDescriptor_FieldTerminalPath: *fp, value: value.(string)}
	case LogDescriptor_FieldPathSelectorLabels:
		return &LogDescriptor_FieldTerminalPathValue{LogDescriptor_FieldTerminalPath: *fp, value: value.([]*common.LabelDescriptor)}
	case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
		return &LogDescriptor_FieldTerminalPathValue{LogDescriptor_FieldTerminalPath: *fp, value: value.([]*common.LabelKeySet)}
	case LogDescriptor_FieldPathSelectorMetadata:
		return &LogDescriptor_FieldTerminalPathValue{LogDescriptor_FieldTerminalPath: *fp, value: value.(*meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for LogDescriptor: %d", fp.selector))
	}
}

func (fp *LogDescriptor_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *LogDescriptor_FieldTerminalPath) WithIArrayOfValues(values interface{}) LogDescriptor_FieldPathArrayOfValues {
	fpaov := &LogDescriptor_FieldTerminalPathArrayOfValues{LogDescriptor_FieldTerminalPath: *fp}
	switch fp.selector {
	case LogDescriptor_FieldPathSelectorName:
		return &LogDescriptor_FieldTerminalPathArrayOfValues{LogDescriptor_FieldTerminalPath: *fp, values: values.([]*Name)}
	case LogDescriptor_FieldPathSelectorDisplayName:
		return &LogDescriptor_FieldTerminalPathArrayOfValues{LogDescriptor_FieldTerminalPath: *fp, values: values.([]string)}
	case LogDescriptor_FieldPathSelectorDescription:
		return &LogDescriptor_FieldTerminalPathArrayOfValues{LogDescriptor_FieldTerminalPath: *fp, values: values.([]string)}
	case LogDescriptor_FieldPathSelectorLabels:
		return &LogDescriptor_FieldTerminalPathArrayOfValues{LogDescriptor_FieldTerminalPath: *fp, values: values.([][]*common.LabelDescriptor)}
	case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
		return &LogDescriptor_FieldTerminalPathArrayOfValues{LogDescriptor_FieldTerminalPath: *fp, values: values.([][]*common.LabelKeySet)}
	case LogDescriptor_FieldPathSelectorMetadata:
		return &LogDescriptor_FieldTerminalPathArrayOfValues{LogDescriptor_FieldTerminalPath: *fp, values: values.([]*meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for LogDescriptor: %d", fp.selector))
	}
	return fpaov
}

func (fp *LogDescriptor_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *LogDescriptor_FieldTerminalPath) WithIArrayItemValue(value interface{}) LogDescriptor_FieldPathArrayItemValue {
	switch fp.selector {
	case LogDescriptor_FieldPathSelectorLabels:
		return &LogDescriptor_FieldTerminalPathArrayItemValue{LogDescriptor_FieldTerminalPath: *fp, value: value.(*common.LabelDescriptor)}
	case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
		return &LogDescriptor_FieldTerminalPathArrayItemValue{LogDescriptor_FieldTerminalPath: *fp, value: value.(*common.LabelKeySet)}
	default:
		panic(fmt.Sprintf("Invalid selector for LogDescriptor: %d", fp.selector))
	}
}

func (fp *LogDescriptor_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

type LogDescriptor_FieldSubPath struct {
	selector LogDescriptor_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ LogDescriptor_FieldPath = (*LogDescriptor_FieldSubPath)(nil)

func (fps *LogDescriptor_FieldSubPath) Selector() LogDescriptor_FieldPathSelector {
	return fps.selector
}
func (fps *LogDescriptor_FieldSubPath) AsLabelsSubPath() (common.LabelDescriptor_FieldPath, bool) {
	res, ok := fps.subPath.(common.LabelDescriptor_FieldPath)
	return res, ok
}
func (fps *LogDescriptor_FieldSubPath) AsPromotedLabelKeySetsSubPath() (common.LabelKeySet_FieldPath, bool) {
	res, ok := fps.subPath.(common.LabelKeySet_FieldPath)
	return res, ok
}
func (fps *LogDescriptor_FieldSubPath) AsMetadataSubPath() (meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(meta.Meta_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *LogDescriptor_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *LogDescriptor_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source LogDescriptor
func (fps *LogDescriptor_FieldSubPath) Get(source *LogDescriptor) (values []interface{}) {
	switch fps.selector {
	case LogDescriptor_FieldPathSelectorLabels:
		for _, item := range source.GetLabels() {
			values = append(values, fps.subPath.GetRaw(item)...)
		}
	case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
		for _, item := range source.GetPromotedLabelKeySets() {
			values = append(values, fps.subPath.GetRaw(item)...)
		}
	case LogDescriptor_FieldPathSelectorMetadata:
		values = append(values, fps.subPath.GetRaw(source.GetMetadata())...)
	default:
		panic(fmt.Sprintf("Invalid selector for LogDescriptor: %d", fps.selector))
	}
	return
}

func (fps *LogDescriptor_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*LogDescriptor))
}

// GetSingle returns value of selected field from source LogDescriptor
func (fps *LogDescriptor_FieldSubPath) GetSingle(source *LogDescriptor) (interface{}, bool) {
	switch fps.selector {
	case LogDescriptor_FieldPathSelectorLabels:
		if len(source.GetLabels()) == 0 {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetLabels()[0])
	case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
		if len(source.GetPromotedLabelKeySets()) == 0 {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetPromotedLabelKeySets()[0])
	case LogDescriptor_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for LogDescriptor: %d", fps.selector))
	}
}

func (fps *LogDescriptor_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*LogDescriptor))
}

// GetDefault returns a default value of the field type
func (fps *LogDescriptor_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *LogDescriptor_FieldSubPath) ClearValue(item *LogDescriptor) {
	if item != nil {
		switch fps.selector {
		case LogDescriptor_FieldPathSelectorLabels:
			for _, subItem := range item.Labels {
				fps.subPath.ClearValueRaw(subItem)
			}
		case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
			for _, subItem := range item.PromotedLabelKeySets {
				fps.subPath.ClearValueRaw(subItem)
			}
		case LogDescriptor_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		default:
			panic(fmt.Sprintf("Invalid selector for LogDescriptor: %d", fps.selector))
		}
	}
}

func (fps *LogDescriptor_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*LogDescriptor))
}

// IsLeaf - whether field path is holds simple value
func (fps *LogDescriptor_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *LogDescriptor_FieldSubPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	iPaths := []gotenobject.FieldPath{&LogDescriptor_FieldTerminalPath{selector: fps.selector}}
	iPaths = append(iPaths, fps.subPath.SplitIntoTerminalIPaths()...)
	return iPaths
}

func (fps *LogDescriptor_FieldSubPath) WithIValue(value interface{}) LogDescriptor_FieldPathValue {
	return &LogDescriptor_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *LogDescriptor_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *LogDescriptor_FieldSubPath) WithIArrayOfValues(values interface{}) LogDescriptor_FieldPathArrayOfValues {
	return &LogDescriptor_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *LogDescriptor_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *LogDescriptor_FieldSubPath) WithIArrayItemValue(value interface{}) LogDescriptor_FieldPathArrayItemValue {
	return &LogDescriptor_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *LogDescriptor_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// LogDescriptor_FieldPathValue allows storing values for LogDescriptor fields according to their type
type LogDescriptor_FieldPathValue interface {
	LogDescriptor_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **LogDescriptor)
	CompareWith(*LogDescriptor) (cmp int, comparable bool)
}

func ParseLogDescriptor_FieldPathValue(pathStr, valueStr string) (LogDescriptor_FieldPathValue, error) {
	fp, err := ParseLogDescriptor_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing LogDescriptor field path value from %s: %v", valueStr, err)
	}
	return fpv.(LogDescriptor_FieldPathValue), nil
}

func MustParseLogDescriptor_FieldPathValue(pathStr, valueStr string) LogDescriptor_FieldPathValue {
	fpv, err := ParseLogDescriptor_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type LogDescriptor_FieldTerminalPathValue struct {
	LogDescriptor_FieldTerminalPath
	value interface{}
}

var _ LogDescriptor_FieldPathValue = (*LogDescriptor_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'LogDescriptor' as interface{}
func (fpv *LogDescriptor_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *LogDescriptor_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *LogDescriptor_FieldTerminalPathValue) AsDisplayNameValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *LogDescriptor_FieldTerminalPathValue) AsDescriptionValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *LogDescriptor_FieldTerminalPathValue) AsLabelsValue() ([]*common.LabelDescriptor, bool) {
	res, ok := fpv.value.([]*common.LabelDescriptor)
	return res, ok
}
func (fpv *LogDescriptor_FieldTerminalPathValue) AsPromotedLabelKeySetsValue() ([]*common.LabelKeySet, bool) {
	res, ok := fpv.value.([]*common.LabelKeySet)
	return res, ok
}
func (fpv *LogDescriptor_FieldTerminalPathValue) AsMetadataValue() (*meta.Meta, bool) {
	res, ok := fpv.value.(*meta.Meta)
	return res, ok
}

// SetTo stores value for selected field for object LogDescriptor
func (fpv *LogDescriptor_FieldTerminalPathValue) SetTo(target **LogDescriptor) {
	if *target == nil {
		*target = new(LogDescriptor)
	}
	switch fpv.selector {
	case LogDescriptor_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case LogDescriptor_FieldPathSelectorDisplayName:
		(*target).DisplayName = fpv.value.(string)
	case LogDescriptor_FieldPathSelectorDescription:
		(*target).Description = fpv.value.(string)
	case LogDescriptor_FieldPathSelectorLabels:
		(*target).Labels = fpv.value.([]*common.LabelDescriptor)
	case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
		(*target).PromotedLabelKeySets = fpv.value.([]*common.LabelKeySet)
	case LogDescriptor_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*meta.Meta)
	default:
		panic(fmt.Sprintf("Invalid selector for LogDescriptor: %d", fpv.selector))
	}
}

func (fpv *LogDescriptor_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*LogDescriptor)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'LogDescriptor_FieldTerminalPathValue' with the value under path in 'LogDescriptor'.
func (fpv *LogDescriptor_FieldTerminalPathValue) CompareWith(source *LogDescriptor) (int, bool) {
	switch fpv.selector {
	case LogDescriptor_FieldPathSelectorName:
		leftValue := fpv.value.(*Name)
		rightValue := source.GetName()
		if leftValue == nil {
			if rightValue != nil {
				return -1, true
			}
			return 0, true
		}
		if rightValue == nil {
			return 1, true
		}
		if leftValue.String() == rightValue.String() {
			return 0, true
		} else if leftValue.String() < rightValue.String() {
			return -1, true
		} else {
			return 1, true
		}
	case LogDescriptor_FieldPathSelectorDisplayName:
		leftValue := fpv.value.(string)
		rightValue := source.GetDisplayName()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case LogDescriptor_FieldPathSelectorDescription:
		leftValue := fpv.value.(string)
		rightValue := source.GetDescription()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case LogDescriptor_FieldPathSelectorLabels:
		return 0, false
	case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
		return 0, false
	case LogDescriptor_FieldPathSelectorMetadata:
		return 0, false
	default:
		panic(fmt.Sprintf("Invalid selector for LogDescriptor: %d", fpv.selector))
	}
}

func (fpv *LogDescriptor_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*LogDescriptor))
}

type LogDescriptor_FieldSubPathValue struct {
	LogDescriptor_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ LogDescriptor_FieldPathValue = (*LogDescriptor_FieldSubPathValue)(nil)

func (fpvs *LogDescriptor_FieldSubPathValue) AsLabelsPathValue() (common.LabelDescriptor_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(common.LabelDescriptor_FieldPathValue)
	return res, ok
}
func (fpvs *LogDescriptor_FieldSubPathValue) AsPromotedLabelKeySetsPathValue() (common.LabelKeySet_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(common.LabelKeySet_FieldPathValue)
	return res, ok
}
func (fpvs *LogDescriptor_FieldSubPathValue) AsMetadataPathValue() (meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(meta.Meta_FieldPathValue)
	return res, ok
}

func (fpvs *LogDescriptor_FieldSubPathValue) SetTo(target **LogDescriptor) {
	if *target == nil {
		*target = new(LogDescriptor)
	}
	switch fpvs.Selector() {
	case LogDescriptor_FieldPathSelectorLabels:
		panic("FieldPath setter is unsupported for array subpaths")
	case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
		panic("FieldPath setter is unsupported for array subpaths")
	case LogDescriptor_FieldPathSelectorMetadata:
		fpvs.subPathValue.(meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	default:
		panic(fmt.Sprintf("Invalid selector for LogDescriptor: %d", fpvs.Selector()))
	}
}

func (fpvs *LogDescriptor_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*LogDescriptor)
	fpvs.SetTo(&typedObject)
}

func (fpvs *LogDescriptor_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *LogDescriptor_FieldSubPathValue) CompareWith(source *LogDescriptor) (int, bool) {
	switch fpvs.Selector() {
	case LogDescriptor_FieldPathSelectorLabels:
		return 0, false // repeated field
	case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
		return 0, false // repeated field
	case LogDescriptor_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for LogDescriptor: %d", fpvs.Selector()))
	}
}

func (fpvs *LogDescriptor_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*LogDescriptor))
}

// LogDescriptor_FieldPathArrayItemValue allows storing single item in Path-specific values for LogDescriptor according to their type
// Present only for array (repeated) types.
type LogDescriptor_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	LogDescriptor_FieldPath
	ContainsValue(*LogDescriptor) bool
}

// ParseLogDescriptor_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseLogDescriptor_FieldPathArrayItemValue(pathStr, valueStr string) (LogDescriptor_FieldPathArrayItemValue, error) {
	fp, err := ParseLogDescriptor_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing LogDescriptor field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(LogDescriptor_FieldPathArrayItemValue), nil
}

func MustParseLogDescriptor_FieldPathArrayItemValue(pathStr, valueStr string) LogDescriptor_FieldPathArrayItemValue {
	fpaiv, err := ParseLogDescriptor_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type LogDescriptor_FieldTerminalPathArrayItemValue struct {
	LogDescriptor_FieldTerminalPath
	value interface{}
}

var _ LogDescriptor_FieldPathArrayItemValue = (*LogDescriptor_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object LogDescriptor as interface{}
func (fpaiv *LogDescriptor_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}
func (fpaiv *LogDescriptor_FieldTerminalPathArrayItemValue) AsLabelsItemValue() (*common.LabelDescriptor, bool) {
	res, ok := fpaiv.value.(*common.LabelDescriptor)
	return res, ok
}
func (fpaiv *LogDescriptor_FieldTerminalPathArrayItemValue) AsPromotedLabelKeySetsItemValue() (*common.LabelKeySet, bool) {
	res, ok := fpaiv.value.(*common.LabelKeySet)
	return res, ok
}

func (fpaiv *LogDescriptor_FieldTerminalPathArrayItemValue) GetSingle(source *LogDescriptor) (interface{}, bool) {
	return nil, false
}

func (fpaiv *LogDescriptor_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*LogDescriptor))
}

// Contains returns a boolean indicating if value that is being held is present in given 'LogDescriptor'
func (fpaiv *LogDescriptor_FieldTerminalPathArrayItemValue) ContainsValue(source *LogDescriptor) bool {
	slice := fpaiv.LogDescriptor_FieldTerminalPath.Get(source)
	for _, v := range slice {
		if asProtoMsg, ok := fpaiv.value.(proto.Message); ok {
			if proto.Equal(asProtoMsg, v.(proto.Message)) {
				return true
			}
		} else if reflect.DeepEqual(v, fpaiv.value) {
			return true
		}
	}
	return false
}

type LogDescriptor_FieldSubPathArrayItemValue struct {
	LogDescriptor_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *LogDescriptor_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *LogDescriptor_FieldSubPathArrayItemValue) AsLabelsPathItemValue() (common.LabelDescriptor_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(common.LabelDescriptor_FieldPathArrayItemValue)
	return res, ok
}
func (fpaivs *LogDescriptor_FieldSubPathArrayItemValue) AsPromotedLabelKeySetsPathItemValue() (common.LabelKeySet_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(common.LabelKeySet_FieldPathArrayItemValue)
	return res, ok
}
func (fpaivs *LogDescriptor_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'LogDescriptor'
func (fpaivs *LogDescriptor_FieldSubPathArrayItemValue) ContainsValue(source *LogDescriptor) bool {
	switch fpaivs.Selector() {
	case LogDescriptor_FieldPathSelectorLabels:
		return false // repeated/map field
	case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
		return false // repeated/map field
	case LogDescriptor_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for LogDescriptor: %d", fpaivs.Selector()))
	}
}

// LogDescriptor_FieldPathArrayOfValues allows storing slice of values for LogDescriptor fields according to their type
type LogDescriptor_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	LogDescriptor_FieldPath
}

func ParseLogDescriptor_FieldPathArrayOfValues(pathStr, valuesStr string) (LogDescriptor_FieldPathArrayOfValues, error) {
	fp, err := ParseLogDescriptor_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing LogDescriptor field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(LogDescriptor_FieldPathArrayOfValues), nil
}

func MustParseLogDescriptor_FieldPathArrayOfValues(pathStr, valuesStr string) LogDescriptor_FieldPathArrayOfValues {
	fpaov, err := ParseLogDescriptor_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type LogDescriptor_FieldTerminalPathArrayOfValues struct {
	LogDescriptor_FieldTerminalPath
	values interface{}
}

var _ LogDescriptor_FieldPathArrayOfValues = (*LogDescriptor_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *LogDescriptor_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case LogDescriptor_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case LogDescriptor_FieldPathSelectorDisplayName:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case LogDescriptor_FieldPathSelectorDescription:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case LogDescriptor_FieldPathSelectorLabels:
		for _, v := range fpaov.values.([][]*common.LabelDescriptor) {
			values = append(values, v)
		}
	case LogDescriptor_FieldPathSelectorPromotedLabelKeySets:
		for _, v := range fpaov.values.([][]*common.LabelKeySet) {
			values = append(values, v)
		}
	case LogDescriptor_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*meta.Meta) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *LogDescriptor_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *LogDescriptor_FieldTerminalPathArrayOfValues) AsDisplayNameArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *LogDescriptor_FieldTerminalPathArrayOfValues) AsDescriptionArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *LogDescriptor_FieldTerminalPathArrayOfValues) AsLabelsArrayOfValues() ([][]*common.LabelDescriptor, bool) {
	res, ok := fpaov.values.([][]*common.LabelDescriptor)
	return res, ok
}
func (fpaov *LogDescriptor_FieldTerminalPathArrayOfValues) AsPromotedLabelKeySetsArrayOfValues() ([][]*common.LabelKeySet, bool) {
	res, ok := fpaov.values.([][]*common.LabelKeySet)
	return res, ok
}
func (fpaov *LogDescriptor_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*meta.Meta, bool) {
	res, ok := fpaov.values.([]*meta.Meta)
	return res, ok
}

type LogDescriptor_FieldSubPathArrayOfValues struct {
	LogDescriptor_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ LogDescriptor_FieldPathArrayOfValues = (*LogDescriptor_FieldSubPathArrayOfValues)(nil)

func (fpsaov *LogDescriptor_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *LogDescriptor_FieldSubPathArrayOfValues) AsLabelsPathArrayOfValues() (common.LabelDescriptor_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(common.LabelDescriptor_FieldPathArrayOfValues)
	return res, ok
}
func (fpsaov *LogDescriptor_FieldSubPathArrayOfValues) AsPromotedLabelKeySetsPathArrayOfValues() (common.LabelKeySet_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(common.LabelKeySet_FieldPathArrayOfValues)
	return res, ok
}
func (fpsaov *LogDescriptor_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
