// Code generated by protoc-gen-goten-object
// File: edgelq/devices/proto/v1alpha2/os_version.proto
// DO NOT EDIT!!!

package os_version

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
	device_type "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/device_type"
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
	_ = &device_type.DeviceType{}
	_ = &meta.Meta{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type OsVersion_FieldPath interface {
	gotenobject.FieldPath
	Selector() OsVersion_FieldPathSelector
	Get(source *OsVersion) []interface{}
	GetSingle(source *OsVersion) (interface{}, bool)
	ClearValue(item *OsVersion)

	// Those methods build corresponding OsVersion_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) OsVersion_FieldPathValue
	WithIArrayOfValues(values interface{}) OsVersion_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) OsVersion_FieldPathArrayItemValue
}

type OsVersion_FieldPathSelector int32

const (
	OsVersion_FieldPathSelectorName                   OsVersion_FieldPathSelector = 0
	OsVersion_FieldPathSelectorMetadata               OsVersion_FieldPathSelector = 1
	OsVersion_FieldPathSelectorVersion                OsVersion_FieldPathSelector = 2
	OsVersion_FieldPathSelectorDeviceType             OsVersion_FieldPathSelector = 3
	OsVersion_FieldPathSelectorMinimumPreviousVersion OsVersion_FieldPathSelector = 4
	OsVersion_FieldPathSelectorChannel                OsVersion_FieldPathSelector = 5
)

func (s OsVersion_FieldPathSelector) String() string {
	switch s {
	case OsVersion_FieldPathSelectorName:
		return "name"
	case OsVersion_FieldPathSelectorMetadata:
		return "metadata"
	case OsVersion_FieldPathSelectorVersion:
		return "version"
	case OsVersion_FieldPathSelectorDeviceType:
		return "device_type"
	case OsVersion_FieldPathSelectorMinimumPreviousVersion:
		return "minimum_previous_version"
	case OsVersion_FieldPathSelectorChannel:
		return "channel"
	default:
		panic(fmt.Sprintf("Invalid selector for OsVersion: %d", s))
	}
}

func BuildOsVersion_FieldPath(fp gotenobject.RawFieldPath) (OsVersion_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object OsVersion")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &OsVersion_FieldTerminalPath{selector: OsVersion_FieldPathSelectorName}, nil
		case "metadata":
			return &OsVersion_FieldTerminalPath{selector: OsVersion_FieldPathSelectorMetadata}, nil
		case "version":
			return &OsVersion_FieldTerminalPath{selector: OsVersion_FieldPathSelectorVersion}, nil
		case "device_type", "deviceType", "device-type":
			return &OsVersion_FieldTerminalPath{selector: OsVersion_FieldPathSelectorDeviceType}, nil
		case "minimum_previous_version", "minimumPreviousVersion", "minimum-previous-version":
			return &OsVersion_FieldTerminalPath{selector: OsVersion_FieldPathSelectorMinimumPreviousVersion}, nil
		case "channel":
			return &OsVersion_FieldTerminalPath{selector: OsVersion_FieldPathSelectorChannel}, nil
		}
	} else {
		switch fp[0] {
		case "metadata":
			if subpath, err := meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &OsVersion_FieldSubPath{selector: OsVersion_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object OsVersion", fp)
}

func ParseOsVersion_FieldPath(rawField string) (OsVersion_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildOsVersion_FieldPath(fp)
}

func MustParseOsVersion_FieldPath(rawField string) OsVersion_FieldPath {
	fp, err := ParseOsVersion_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type OsVersion_FieldTerminalPath struct {
	selector OsVersion_FieldPathSelector
}

var _ OsVersion_FieldPath = (*OsVersion_FieldTerminalPath)(nil)

func (fp *OsVersion_FieldTerminalPath) Selector() OsVersion_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *OsVersion_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *OsVersion_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source OsVersion
func (fp *OsVersion_FieldTerminalPath) Get(source *OsVersion) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case OsVersion_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case OsVersion_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		case OsVersion_FieldPathSelectorVersion:
			values = append(values, source.Version)
		case OsVersion_FieldPathSelectorDeviceType:
			if source.DeviceType != nil {
				values = append(values, source.DeviceType)
			}
		case OsVersion_FieldPathSelectorMinimumPreviousVersion:
			values = append(values, source.MinimumPreviousVersion)
		case OsVersion_FieldPathSelectorChannel:
			values = append(values, source.Channel)
		default:
			panic(fmt.Sprintf("Invalid selector for OsVersion: %d", fp.selector))
		}
	}
	return
}

func (fp *OsVersion_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*OsVersion))
}

// GetSingle returns value pointed by specific field of from source OsVersion
func (fp *OsVersion_FieldTerminalPath) GetSingle(source *OsVersion) (interface{}, bool) {
	switch fp.selector {
	case OsVersion_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case OsVersion_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	case OsVersion_FieldPathSelectorVersion:
		return source.GetVersion(), source != nil
	case OsVersion_FieldPathSelectorDeviceType:
		res := source.GetDeviceType()
		return res, res != nil
	case OsVersion_FieldPathSelectorMinimumPreviousVersion:
		return source.GetMinimumPreviousVersion(), source != nil
	case OsVersion_FieldPathSelectorChannel:
		return source.GetChannel(), source != nil
	default:
		panic(fmt.Sprintf("Invalid selector for OsVersion: %d", fp.selector))
	}
}

func (fp *OsVersion_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*OsVersion))
}

// GetDefault returns a default value of the field type
func (fp *OsVersion_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case OsVersion_FieldPathSelectorName:
		return (*Name)(nil)
	case OsVersion_FieldPathSelectorMetadata:
		return (*meta.Meta)(nil)
	case OsVersion_FieldPathSelectorVersion:
		return ""
	case OsVersion_FieldPathSelectorDeviceType:
		return (*device_type.Reference)(nil)
	case OsVersion_FieldPathSelectorMinimumPreviousVersion:
		return ""
	case OsVersion_FieldPathSelectorChannel:
		return OsVersion_CHANNEL_UNSPECIFIED
	default:
		panic(fmt.Sprintf("Invalid selector for OsVersion: %d", fp.selector))
	}
}

func (fp *OsVersion_FieldTerminalPath) ClearValue(item *OsVersion) {
	if item != nil {
		switch fp.selector {
		case OsVersion_FieldPathSelectorName:
			item.Name = nil
		case OsVersion_FieldPathSelectorMetadata:
			item.Metadata = nil
		case OsVersion_FieldPathSelectorVersion:
			item.Version = ""
		case OsVersion_FieldPathSelectorDeviceType:
			item.DeviceType = nil
		case OsVersion_FieldPathSelectorMinimumPreviousVersion:
			item.MinimumPreviousVersion = ""
		case OsVersion_FieldPathSelectorChannel:
			item.Channel = OsVersion_CHANNEL_UNSPECIFIED
		default:
			panic(fmt.Sprintf("Invalid selector for OsVersion: %d", fp.selector))
		}
	}
}

func (fp *OsVersion_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*OsVersion))
}

// IsLeaf - whether field path is holds simple value
func (fp *OsVersion_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == OsVersion_FieldPathSelectorName ||
		fp.selector == OsVersion_FieldPathSelectorVersion ||
		fp.selector == OsVersion_FieldPathSelectorDeviceType ||
		fp.selector == OsVersion_FieldPathSelectorMinimumPreviousVersion ||
		fp.selector == OsVersion_FieldPathSelectorChannel
}

func (fp *OsVersion_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *OsVersion_FieldTerminalPath) WithIValue(value interface{}) OsVersion_FieldPathValue {
	switch fp.selector {
	case OsVersion_FieldPathSelectorName:
		return &OsVersion_FieldTerminalPathValue{OsVersion_FieldTerminalPath: *fp, value: value.(*Name)}
	case OsVersion_FieldPathSelectorMetadata:
		return &OsVersion_FieldTerminalPathValue{OsVersion_FieldTerminalPath: *fp, value: value.(*meta.Meta)}
	case OsVersion_FieldPathSelectorVersion:
		return &OsVersion_FieldTerminalPathValue{OsVersion_FieldTerminalPath: *fp, value: value.(string)}
	case OsVersion_FieldPathSelectorDeviceType:
		return &OsVersion_FieldTerminalPathValue{OsVersion_FieldTerminalPath: *fp, value: value.(*device_type.Reference)}
	case OsVersion_FieldPathSelectorMinimumPreviousVersion:
		return &OsVersion_FieldTerminalPathValue{OsVersion_FieldTerminalPath: *fp, value: value.(string)}
	case OsVersion_FieldPathSelectorChannel:
		return &OsVersion_FieldTerminalPathValue{OsVersion_FieldTerminalPath: *fp, value: value.(OsVersion_Channel)}
	default:
		panic(fmt.Sprintf("Invalid selector for OsVersion: %d", fp.selector))
	}
}

func (fp *OsVersion_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *OsVersion_FieldTerminalPath) WithIArrayOfValues(values interface{}) OsVersion_FieldPathArrayOfValues {
	fpaov := &OsVersion_FieldTerminalPathArrayOfValues{OsVersion_FieldTerminalPath: *fp}
	switch fp.selector {
	case OsVersion_FieldPathSelectorName:
		return &OsVersion_FieldTerminalPathArrayOfValues{OsVersion_FieldTerminalPath: *fp, values: values.([]*Name)}
	case OsVersion_FieldPathSelectorMetadata:
		return &OsVersion_FieldTerminalPathArrayOfValues{OsVersion_FieldTerminalPath: *fp, values: values.([]*meta.Meta)}
	case OsVersion_FieldPathSelectorVersion:
		return &OsVersion_FieldTerminalPathArrayOfValues{OsVersion_FieldTerminalPath: *fp, values: values.([]string)}
	case OsVersion_FieldPathSelectorDeviceType:
		return &OsVersion_FieldTerminalPathArrayOfValues{OsVersion_FieldTerminalPath: *fp, values: values.([]*device_type.Reference)}
	case OsVersion_FieldPathSelectorMinimumPreviousVersion:
		return &OsVersion_FieldTerminalPathArrayOfValues{OsVersion_FieldTerminalPath: *fp, values: values.([]string)}
	case OsVersion_FieldPathSelectorChannel:
		return &OsVersion_FieldTerminalPathArrayOfValues{OsVersion_FieldTerminalPath: *fp, values: values.([]OsVersion_Channel)}
	default:
		panic(fmt.Sprintf("Invalid selector for OsVersion: %d", fp.selector))
	}
	return fpaov
}

func (fp *OsVersion_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *OsVersion_FieldTerminalPath) WithIArrayItemValue(value interface{}) OsVersion_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for OsVersion: %d", fp.selector))
	}
}

func (fp *OsVersion_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

type OsVersion_FieldSubPath struct {
	selector OsVersion_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ OsVersion_FieldPath = (*OsVersion_FieldSubPath)(nil)

func (fps *OsVersion_FieldSubPath) Selector() OsVersion_FieldPathSelector {
	return fps.selector
}
func (fps *OsVersion_FieldSubPath) AsMetadataSubPath() (meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(meta.Meta_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *OsVersion_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *OsVersion_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source OsVersion
func (fps *OsVersion_FieldSubPath) Get(source *OsVersion) (values []interface{}) {
	switch fps.selector {
	case OsVersion_FieldPathSelectorMetadata:
		values = append(values, fps.subPath.GetRaw(source.GetMetadata())...)
	default:
		panic(fmt.Sprintf("Invalid selector for OsVersion: %d", fps.selector))
	}
	return
}

func (fps *OsVersion_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*OsVersion))
}

// GetSingle returns value of selected field from source OsVersion
func (fps *OsVersion_FieldSubPath) GetSingle(source *OsVersion) (interface{}, bool) {
	switch fps.selector {
	case OsVersion_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for OsVersion: %d", fps.selector))
	}
}

func (fps *OsVersion_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*OsVersion))
}

// GetDefault returns a default value of the field type
func (fps *OsVersion_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *OsVersion_FieldSubPath) ClearValue(item *OsVersion) {
	if item != nil {
		switch fps.selector {
		case OsVersion_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		default:
			panic(fmt.Sprintf("Invalid selector for OsVersion: %d", fps.selector))
		}
	}
}

func (fps *OsVersion_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*OsVersion))
}

// IsLeaf - whether field path is holds simple value
func (fps *OsVersion_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *OsVersion_FieldSubPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	iPaths := []gotenobject.FieldPath{&OsVersion_FieldTerminalPath{selector: fps.selector}}
	iPaths = append(iPaths, fps.subPath.SplitIntoTerminalIPaths()...)
	return iPaths
}

func (fps *OsVersion_FieldSubPath) WithIValue(value interface{}) OsVersion_FieldPathValue {
	return &OsVersion_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *OsVersion_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *OsVersion_FieldSubPath) WithIArrayOfValues(values interface{}) OsVersion_FieldPathArrayOfValues {
	return &OsVersion_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *OsVersion_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *OsVersion_FieldSubPath) WithIArrayItemValue(value interface{}) OsVersion_FieldPathArrayItemValue {
	return &OsVersion_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *OsVersion_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// OsVersion_FieldPathValue allows storing values for OsVersion fields according to their type
type OsVersion_FieldPathValue interface {
	OsVersion_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **OsVersion)
	CompareWith(*OsVersion) (cmp int, comparable bool)
}

func ParseOsVersion_FieldPathValue(pathStr, valueStr string) (OsVersion_FieldPathValue, error) {
	fp, err := ParseOsVersion_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing OsVersion field path value from %s: %v", valueStr, err)
	}
	return fpv.(OsVersion_FieldPathValue), nil
}

func MustParseOsVersion_FieldPathValue(pathStr, valueStr string) OsVersion_FieldPathValue {
	fpv, err := ParseOsVersion_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type OsVersion_FieldTerminalPathValue struct {
	OsVersion_FieldTerminalPath
	value interface{}
}

var _ OsVersion_FieldPathValue = (*OsVersion_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'OsVersion' as interface{}
func (fpv *OsVersion_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *OsVersion_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *OsVersion_FieldTerminalPathValue) AsMetadataValue() (*meta.Meta, bool) {
	res, ok := fpv.value.(*meta.Meta)
	return res, ok
}
func (fpv *OsVersion_FieldTerminalPathValue) AsVersionValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *OsVersion_FieldTerminalPathValue) AsDeviceTypeValue() (*device_type.Reference, bool) {
	res, ok := fpv.value.(*device_type.Reference)
	return res, ok
}
func (fpv *OsVersion_FieldTerminalPathValue) AsMinimumPreviousVersionValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *OsVersion_FieldTerminalPathValue) AsChannelValue() (OsVersion_Channel, bool) {
	res, ok := fpv.value.(OsVersion_Channel)
	return res, ok
}

// SetTo stores value for selected field for object OsVersion
func (fpv *OsVersion_FieldTerminalPathValue) SetTo(target **OsVersion) {
	if *target == nil {
		*target = new(OsVersion)
	}
	switch fpv.selector {
	case OsVersion_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case OsVersion_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*meta.Meta)
	case OsVersion_FieldPathSelectorVersion:
		(*target).Version = fpv.value.(string)
	case OsVersion_FieldPathSelectorDeviceType:
		(*target).DeviceType = fpv.value.(*device_type.Reference)
	case OsVersion_FieldPathSelectorMinimumPreviousVersion:
		(*target).MinimumPreviousVersion = fpv.value.(string)
	case OsVersion_FieldPathSelectorChannel:
		(*target).Channel = fpv.value.(OsVersion_Channel)
	default:
		panic(fmt.Sprintf("Invalid selector for OsVersion: %d", fpv.selector))
	}
}

func (fpv *OsVersion_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*OsVersion)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'OsVersion_FieldTerminalPathValue' with the value under path in 'OsVersion'.
func (fpv *OsVersion_FieldTerminalPathValue) CompareWith(source *OsVersion) (int, bool) {
	switch fpv.selector {
	case OsVersion_FieldPathSelectorName:
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
	case OsVersion_FieldPathSelectorMetadata:
		return 0, false
	case OsVersion_FieldPathSelectorVersion:
		leftValue := fpv.value.(string)
		rightValue := source.GetVersion()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case OsVersion_FieldPathSelectorDeviceType:
		leftValue := fpv.value.(*device_type.Reference)
		rightValue := source.GetDeviceType()
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
	case OsVersion_FieldPathSelectorMinimumPreviousVersion:
		leftValue := fpv.value.(string)
		rightValue := source.GetMinimumPreviousVersion()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case OsVersion_FieldPathSelectorChannel:
		leftValue := fpv.value.(OsVersion_Channel)
		rightValue := source.GetChannel()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	default:
		panic(fmt.Sprintf("Invalid selector for OsVersion: %d", fpv.selector))
	}
}

func (fpv *OsVersion_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*OsVersion))
}

type OsVersion_FieldSubPathValue struct {
	OsVersion_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ OsVersion_FieldPathValue = (*OsVersion_FieldSubPathValue)(nil)

func (fpvs *OsVersion_FieldSubPathValue) AsMetadataPathValue() (meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(meta.Meta_FieldPathValue)
	return res, ok
}

func (fpvs *OsVersion_FieldSubPathValue) SetTo(target **OsVersion) {
	if *target == nil {
		*target = new(OsVersion)
	}
	switch fpvs.Selector() {
	case OsVersion_FieldPathSelectorMetadata:
		fpvs.subPathValue.(meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	default:
		panic(fmt.Sprintf("Invalid selector for OsVersion: %d", fpvs.Selector()))
	}
}

func (fpvs *OsVersion_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*OsVersion)
	fpvs.SetTo(&typedObject)
}

func (fpvs *OsVersion_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *OsVersion_FieldSubPathValue) CompareWith(source *OsVersion) (int, bool) {
	switch fpvs.Selector() {
	case OsVersion_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for OsVersion: %d", fpvs.Selector()))
	}
}

func (fpvs *OsVersion_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*OsVersion))
}

// OsVersion_FieldPathArrayItemValue allows storing single item in Path-specific values for OsVersion according to their type
// Present only for array (repeated) types.
type OsVersion_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	OsVersion_FieldPath
	ContainsValue(*OsVersion) bool
}

// ParseOsVersion_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseOsVersion_FieldPathArrayItemValue(pathStr, valueStr string) (OsVersion_FieldPathArrayItemValue, error) {
	fp, err := ParseOsVersion_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing OsVersion field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(OsVersion_FieldPathArrayItemValue), nil
}

func MustParseOsVersion_FieldPathArrayItemValue(pathStr, valueStr string) OsVersion_FieldPathArrayItemValue {
	fpaiv, err := ParseOsVersion_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type OsVersion_FieldTerminalPathArrayItemValue struct {
	OsVersion_FieldTerminalPath
	value interface{}
}

var _ OsVersion_FieldPathArrayItemValue = (*OsVersion_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object OsVersion as interface{}
func (fpaiv *OsVersion_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *OsVersion_FieldTerminalPathArrayItemValue) GetSingle(source *OsVersion) (interface{}, bool) {
	return nil, false
}

func (fpaiv *OsVersion_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*OsVersion))
}

// Contains returns a boolean indicating if value that is being held is present in given 'OsVersion'
func (fpaiv *OsVersion_FieldTerminalPathArrayItemValue) ContainsValue(source *OsVersion) bool {
	slice := fpaiv.OsVersion_FieldTerminalPath.Get(source)
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

type OsVersion_FieldSubPathArrayItemValue struct {
	OsVersion_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *OsVersion_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *OsVersion_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'OsVersion'
func (fpaivs *OsVersion_FieldSubPathArrayItemValue) ContainsValue(source *OsVersion) bool {
	switch fpaivs.Selector() {
	case OsVersion_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for OsVersion: %d", fpaivs.Selector()))
	}
}

// OsVersion_FieldPathArrayOfValues allows storing slice of values for OsVersion fields according to their type
type OsVersion_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	OsVersion_FieldPath
}

func ParseOsVersion_FieldPathArrayOfValues(pathStr, valuesStr string) (OsVersion_FieldPathArrayOfValues, error) {
	fp, err := ParseOsVersion_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing OsVersion field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(OsVersion_FieldPathArrayOfValues), nil
}

func MustParseOsVersion_FieldPathArrayOfValues(pathStr, valuesStr string) OsVersion_FieldPathArrayOfValues {
	fpaov, err := ParseOsVersion_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type OsVersion_FieldTerminalPathArrayOfValues struct {
	OsVersion_FieldTerminalPath
	values interface{}
}

var _ OsVersion_FieldPathArrayOfValues = (*OsVersion_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *OsVersion_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case OsVersion_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case OsVersion_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*meta.Meta) {
			values = append(values, v)
		}
	case OsVersion_FieldPathSelectorVersion:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case OsVersion_FieldPathSelectorDeviceType:
		for _, v := range fpaov.values.([]*device_type.Reference) {
			values = append(values, v)
		}
	case OsVersion_FieldPathSelectorMinimumPreviousVersion:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case OsVersion_FieldPathSelectorChannel:
		for _, v := range fpaov.values.([]OsVersion_Channel) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *OsVersion_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *OsVersion_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*meta.Meta, bool) {
	res, ok := fpaov.values.([]*meta.Meta)
	return res, ok
}
func (fpaov *OsVersion_FieldTerminalPathArrayOfValues) AsVersionArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *OsVersion_FieldTerminalPathArrayOfValues) AsDeviceTypeArrayOfValues() ([]*device_type.Reference, bool) {
	res, ok := fpaov.values.([]*device_type.Reference)
	return res, ok
}
func (fpaov *OsVersion_FieldTerminalPathArrayOfValues) AsMinimumPreviousVersionArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *OsVersion_FieldTerminalPathArrayOfValues) AsChannelArrayOfValues() ([]OsVersion_Channel, bool) {
	res, ok := fpaov.values.([]OsVersion_Channel)
	return res, ok
}

type OsVersion_FieldSubPathArrayOfValues struct {
	OsVersion_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ OsVersion_FieldPathArrayOfValues = (*OsVersion_FieldSubPathArrayOfValues)(nil)

func (fpsaov *OsVersion_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *OsVersion_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
