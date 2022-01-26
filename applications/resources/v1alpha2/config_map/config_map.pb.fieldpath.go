// Code generated by protoc-gen-goten-object
// File: edgelq/applications/proto/v1alpha2/config_map.proto
// DO NOT EDIT!!!

package config_map

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha2/project"
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
)

// ensure the imports are used
var (
	_ = json.Marshaler(nil)
	_ = fmt.Stringer(nil)
	_ = reflect.DeepEqual
	_ = strings.Builder{}
	_ = time.Second

	_ = strcase.ToLowerCamel
	_ = codes.NotFound
	_ = status.Status{}
	_ = protojson.UnmarshalOptions{}
	_ = proto.Message(nil)
	_ = protoregistry.GlobalTypes
	_ = fieldmaskpb.FieldMask{}

	_ = gotenobject.FieldPath(nil)
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &ntt_meta.Meta{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type ConfigMap_FieldPath interface {
	gotenobject.FieldPath
	Selector() ConfigMap_FieldPathSelector
	Get(source *ConfigMap) []interface{}
	GetSingle(source *ConfigMap) (interface{}, bool)
	ClearValue(item *ConfigMap)

	// Those methods build corresponding ConfigMap_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) ConfigMap_FieldPathValue
	WithIArrayOfValues(values interface{}) ConfigMap_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) ConfigMap_FieldPathArrayItemValue
}

type ConfigMap_FieldPathSelector int32

const (
	ConfigMap_FieldPathSelectorName        ConfigMap_FieldPathSelector = 0
	ConfigMap_FieldPathSelectorDisplayName ConfigMap_FieldPathSelector = 1
	ConfigMap_FieldPathSelectorMetadata    ConfigMap_FieldPathSelector = 2
	ConfigMap_FieldPathSelectorData        ConfigMap_FieldPathSelector = 3
	ConfigMap_FieldPathSelectorBinaryData  ConfigMap_FieldPathSelector = 4
)

func (s ConfigMap_FieldPathSelector) String() string {
	switch s {
	case ConfigMap_FieldPathSelectorName:
		return "name"
	case ConfigMap_FieldPathSelectorDisplayName:
		return "display_name"
	case ConfigMap_FieldPathSelectorMetadata:
		return "metadata"
	case ConfigMap_FieldPathSelectorData:
		return "data"
	case ConfigMap_FieldPathSelectorBinaryData:
		return "binary_data"
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", s))
	}
}

func BuildConfigMap_FieldPath(fp gotenobject.RawFieldPath) (ConfigMap_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object ConfigMap")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorName}, nil
		case "display_name", "displayName", "display-name":
			return &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorDisplayName}, nil
		case "metadata":
			return &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorMetadata}, nil
		case "data":
			return &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorData}, nil
		case "binary_data", "binaryData", "binary-data":
			return &ConfigMap_FieldTerminalPath{selector: ConfigMap_FieldPathSelectorBinaryData}, nil
		}
	} else {
		switch fp[0] {
		case "metadata":
			if subpath, err := ntt_meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &ConfigMap_FieldSubPath{selector: ConfigMap_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		case "data":
			if len(fp) > 2 {
				return nil, status.Errorf(codes.InvalidArgument, "sub path for maps ('%s') are not supported (object ConfigMap)", fp)
			}
			return &ConfigMap_FieldPathMap{selector: ConfigMap_FieldPathSelectorData, key: fp[1]}, nil
		case "binary_data", "binaryData", "binary-data":
			if len(fp) > 2 {
				return nil, status.Errorf(codes.InvalidArgument, "sub path for maps ('%s') are not supported (object ConfigMap)", fp)
			}
			return &ConfigMap_FieldPathMap{selector: ConfigMap_FieldPathSelectorBinaryData, key: fp[1]}, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object ConfigMap", fp)
}

func ParseConfigMap_FieldPath(rawField string) (ConfigMap_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildConfigMap_FieldPath(fp)
}

func MustParseConfigMap_FieldPath(rawField string) ConfigMap_FieldPath {
	fp, err := ParseConfigMap_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type ConfigMap_FieldTerminalPath struct {
	selector ConfigMap_FieldPathSelector
}

var _ ConfigMap_FieldPath = (*ConfigMap_FieldTerminalPath)(nil)

func (fp *ConfigMap_FieldTerminalPath) Selector() ConfigMap_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *ConfigMap_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *ConfigMap_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source ConfigMap
func (fp *ConfigMap_FieldTerminalPath) Get(source *ConfigMap) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case ConfigMap_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case ConfigMap_FieldPathSelectorDisplayName:
			values = append(values, source.DisplayName)
		case ConfigMap_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		case ConfigMap_FieldPathSelectorData:
			values = append(values, source.Data)
		case ConfigMap_FieldPathSelectorBinaryData:
			if source.BinaryData != nil {
				values = append(values, source.BinaryData)
			}
		default:
			panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fp.selector))
		}
	}
	return
}

func (fp *ConfigMap_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*ConfigMap))
}

// GetSingle returns value pointed by specific field of from source ConfigMap
func (fp *ConfigMap_FieldTerminalPath) GetSingle(source *ConfigMap) (interface{}, bool) {
	switch fp.selector {
	case ConfigMap_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case ConfigMap_FieldPathSelectorDisplayName:
		return source.GetDisplayName(), source != nil
	case ConfigMap_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	case ConfigMap_FieldPathSelectorData:
		res := source.GetData()
		return res, res != nil
	case ConfigMap_FieldPathSelectorBinaryData:
		res := source.GetBinaryData()
		return res, res != nil
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fp.selector))
	}
}

func (fp *ConfigMap_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*ConfigMap))
}

// GetDefault returns a default value of the field type
func (fp *ConfigMap_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case ConfigMap_FieldPathSelectorName:
		return (*Name)(nil)
	case ConfigMap_FieldPathSelectorDisplayName:
		return ""
	case ConfigMap_FieldPathSelectorMetadata:
		return (*ntt_meta.Meta)(nil)
	case ConfigMap_FieldPathSelectorData:
		return (map[string]string)(nil)
	case ConfigMap_FieldPathSelectorBinaryData:
		return (map[string][]byte)(nil)
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fp.selector))
	}
}

func (fp *ConfigMap_FieldTerminalPath) ClearValue(item *ConfigMap) {
	if item != nil {
		switch fp.selector {
		case ConfigMap_FieldPathSelectorName:
			item.Name = nil
		case ConfigMap_FieldPathSelectorDisplayName:
			item.DisplayName = ""
		case ConfigMap_FieldPathSelectorMetadata:
			item.Metadata = nil
		case ConfigMap_FieldPathSelectorData:
			item.Data = nil
		case ConfigMap_FieldPathSelectorBinaryData:
			item.BinaryData = nil
		default:
			panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fp.selector))
		}
	}
}

func (fp *ConfigMap_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*ConfigMap))
}

// IsLeaf - whether field path is holds simple value
func (fp *ConfigMap_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == ConfigMap_FieldPathSelectorName ||
		fp.selector == ConfigMap_FieldPathSelectorDisplayName ||
		fp.selector == ConfigMap_FieldPathSelectorData ||
		fp.selector == ConfigMap_FieldPathSelectorBinaryData
}

func (fp *ConfigMap_FieldTerminalPath) WithIValue(value interface{}) ConfigMap_FieldPathValue {
	switch fp.selector {
	case ConfigMap_FieldPathSelectorName:
		return &ConfigMap_FieldTerminalPathValue{ConfigMap_FieldTerminalPath: *fp, value: value.(*Name)}
	case ConfigMap_FieldPathSelectorDisplayName:
		return &ConfigMap_FieldTerminalPathValue{ConfigMap_FieldTerminalPath: *fp, value: value.(string)}
	case ConfigMap_FieldPathSelectorMetadata:
		return &ConfigMap_FieldTerminalPathValue{ConfigMap_FieldTerminalPath: *fp, value: value.(*ntt_meta.Meta)}
	case ConfigMap_FieldPathSelectorData:
		return &ConfigMap_FieldTerminalPathValue{ConfigMap_FieldTerminalPath: *fp, value: value.(map[string]string)}
	case ConfigMap_FieldPathSelectorBinaryData:
		return &ConfigMap_FieldTerminalPathValue{ConfigMap_FieldTerminalPath: *fp, value: value.(map[string][]byte)}
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fp.selector))
	}
}

func (fp *ConfigMap_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *ConfigMap_FieldTerminalPath) WithIArrayOfValues(values interface{}) ConfigMap_FieldPathArrayOfValues {
	fpaov := &ConfigMap_FieldTerminalPathArrayOfValues{ConfigMap_FieldTerminalPath: *fp}
	switch fp.selector {
	case ConfigMap_FieldPathSelectorName:
		return &ConfigMap_FieldTerminalPathArrayOfValues{ConfigMap_FieldTerminalPath: *fp, values: values.([]*Name)}
	case ConfigMap_FieldPathSelectorDisplayName:
		return &ConfigMap_FieldTerminalPathArrayOfValues{ConfigMap_FieldTerminalPath: *fp, values: values.([]string)}
	case ConfigMap_FieldPathSelectorMetadata:
		return &ConfigMap_FieldTerminalPathArrayOfValues{ConfigMap_FieldTerminalPath: *fp, values: values.([]*ntt_meta.Meta)}
	case ConfigMap_FieldPathSelectorData:
		return &ConfigMap_FieldTerminalPathArrayOfValues{ConfigMap_FieldTerminalPath: *fp, values: values.([]map[string]string)}
	case ConfigMap_FieldPathSelectorBinaryData:
		return &ConfigMap_FieldTerminalPathArrayOfValues{ConfigMap_FieldTerminalPath: *fp, values: values.([]map[string][]byte)}
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fp.selector))
	}
	return fpaov
}

func (fp *ConfigMap_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *ConfigMap_FieldTerminalPath) WithIArrayItemValue(value interface{}) ConfigMap_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fp.selector))
	}
}

func (fp *ConfigMap_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

// FieldPath for map type with additional Key information
type ConfigMap_FieldPathMap struct {
	key      string
	selector ConfigMap_FieldPathSelector
}

var _ ConfigMap_FieldPath = (*ConfigMap_FieldPathMap)(nil)

func (fpm *ConfigMap_FieldPathMap) Selector() ConfigMap_FieldPathSelector {
	return fpm.selector
}

func (fpm *ConfigMap_FieldPathMap) Key() string {
	return fpm.key
}

// String returns path representation in proto convention
func (fpm *ConfigMap_FieldPathMap) String() string {
	return fpm.selector.String() + "." + fpm.key
}

// JSONString returns path representation is JSON convention. Note that map keys are not transformed
func (fpm *ConfigMap_FieldPathMap) JSONString() string {
	return strcase.ToLowerCamel(fpm.selector.String()) + "." + fpm.key
}

// Get returns all values pointed by selected field map key from source ConfigMap
func (fpm *ConfigMap_FieldPathMap) Get(source *ConfigMap) (values []interface{}) {
	switch fpm.selector {
	case ConfigMap_FieldPathSelectorData:
		if value, ok := source.GetData()[fpm.key]; ok {
			values = append(values, value)
		}
	case ConfigMap_FieldPathSelectorBinaryData:
		if value, ok := source.GetBinaryData()[fpm.key]; ok {
			values = append(values, value)
		}
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fpm.selector))
	}
	return
}

func (fpm *ConfigMap_FieldPathMap) GetRaw(source proto.Message) []interface{} {
	return fpm.Get(source.(*ConfigMap))
}

// GetSingle returns value by selected field map key from source ConfigMap
func (fpm *ConfigMap_FieldPathMap) GetSingle(source *ConfigMap) (interface{}, bool) {
	switch fpm.selector {
	case ConfigMap_FieldPathSelectorData:
		res, ok := source.GetData()[fpm.key]
		return res, ok
	case ConfigMap_FieldPathSelectorBinaryData:
		res, ok := source.GetBinaryData()[fpm.key]
		return res, ok
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fpm.selector))
	}
}

func (fpm *ConfigMap_FieldPathMap) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpm.GetSingle(source.(*ConfigMap))
}

// GetDefault returns a default value of the field type
func (fpm *ConfigMap_FieldPathMap) GetDefault() interface{} {
	switch fpm.selector {
	case ConfigMap_FieldPathSelectorData:
		var v string
		return v
	case ConfigMap_FieldPathSelectorBinaryData:
		var v []byte
		return v
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fpm.selector))
	}
}

func (fpm *ConfigMap_FieldPathMap) ClearValue(item *ConfigMap) {
	if item != nil {
		switch fpm.selector {
		case ConfigMap_FieldPathSelectorData:
			delete(item.Data, fpm.key)
		case ConfigMap_FieldPathSelectorBinaryData:
			delete(item.BinaryData, fpm.key)
		default:
			panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fpm.selector))
		}
	}
}

func (fpm *ConfigMap_FieldPathMap) ClearValueRaw(item proto.Message) {
	fpm.ClearValue(item.(*ConfigMap))
}

// IsLeaf - whether field path is holds simple value
func (fpm *ConfigMap_FieldPathMap) IsLeaf() bool {
	switch fpm.selector {
	case ConfigMap_FieldPathSelectorData:
		return true
	case ConfigMap_FieldPathSelectorBinaryData:
		return true
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fpm.selector))
	}
}

func (fpm *ConfigMap_FieldPathMap) WithIValue(value interface{}) ConfigMap_FieldPathValue {
	switch fpm.selector {
	case ConfigMap_FieldPathSelectorData:
		return &ConfigMap_FieldPathMapValue{ConfigMap_FieldPathMap: *fpm, value: value.(string)}
	case ConfigMap_FieldPathSelectorBinaryData:
		return &ConfigMap_FieldPathMapValue{ConfigMap_FieldPathMap: *fpm, value: value.([]byte)}
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fpm.selector))
	}
}

func (fpm *ConfigMap_FieldPathMap) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fpm.WithIValue(value)
}

func (fpm *ConfigMap_FieldPathMap) WithIArrayOfValues(values interface{}) ConfigMap_FieldPathArrayOfValues {
	switch fpm.selector {
	case ConfigMap_FieldPathSelectorData:
		return &ConfigMap_FieldPathMapArrayOfValues{ConfigMap_FieldPathMap: *fpm, values: values.([]string)}
	case ConfigMap_FieldPathSelectorBinaryData:
		return &ConfigMap_FieldPathMapArrayOfValues{ConfigMap_FieldPathMap: *fpm, values: values.([][]byte)}
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fpm.selector))
	}
}

func (fpm *ConfigMap_FieldPathMap) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fpm.WithIArrayOfValues(values)
}

func (fpm *ConfigMap_FieldPathMap) WithIArrayItemValue(value interface{}) ConfigMap_FieldPathArrayItemValue {
	panic("Cannot create array item value from map fieldpath")
}

func (fpm *ConfigMap_FieldPathMap) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fpm.WithIArrayItemValue(value)
}

type ConfigMap_FieldSubPath struct {
	selector ConfigMap_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ ConfigMap_FieldPath = (*ConfigMap_FieldSubPath)(nil)

func (fps *ConfigMap_FieldSubPath) Selector() ConfigMap_FieldPathSelector {
	return fps.selector
}
func (fps *ConfigMap_FieldSubPath) AsMetadataSubPath() (ntt_meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(ntt_meta.Meta_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *ConfigMap_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *ConfigMap_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source ConfigMap
func (fps *ConfigMap_FieldSubPath) Get(source *ConfigMap) (values []interface{}) {
	if asMetaFieldPath, ok := fps.AsMetadataSubPath(); ok {
		values = append(values, asMetaFieldPath.Get(source.GetMetadata())...)
	} else {
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fps.selector))
	}
	return
}

func (fps *ConfigMap_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*ConfigMap))
}

// GetSingle returns value of selected field from source ConfigMap
func (fps *ConfigMap_FieldSubPath) GetSingle(source *ConfigMap) (interface{}, bool) {
	switch fps.selector {
	case ConfigMap_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fps.selector))
	}
}

func (fps *ConfigMap_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*ConfigMap))
}

// GetDefault returns a default value of the field type
func (fps *ConfigMap_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *ConfigMap_FieldSubPath) ClearValue(item *ConfigMap) {
	if item != nil {
		switch fps.selector {
		case ConfigMap_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		default:
			panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fps.selector))
		}
	}
}

func (fps *ConfigMap_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*ConfigMap))
}

// IsLeaf - whether field path is holds simple value
func (fps *ConfigMap_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *ConfigMap_FieldSubPath) WithIValue(value interface{}) ConfigMap_FieldPathValue {
	return &ConfigMap_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *ConfigMap_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *ConfigMap_FieldSubPath) WithIArrayOfValues(values interface{}) ConfigMap_FieldPathArrayOfValues {
	return &ConfigMap_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *ConfigMap_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *ConfigMap_FieldSubPath) WithIArrayItemValue(value interface{}) ConfigMap_FieldPathArrayItemValue {
	return &ConfigMap_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *ConfigMap_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// ConfigMap_FieldPathValue allows storing values for ConfigMap fields according to their type
type ConfigMap_FieldPathValue interface {
	ConfigMap_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **ConfigMap)
	CompareWith(*ConfigMap) (cmp int, comparable bool)
}

func ParseConfigMap_FieldPathValue(pathStr, valueStr string) (ConfigMap_FieldPathValue, error) {
	fp, err := ParseConfigMap_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing ConfigMap field path value from %s: %v", valueStr, err)
	}
	return fpv.(ConfigMap_FieldPathValue), nil
}

func MustParseConfigMap_FieldPathValue(pathStr, valueStr string) ConfigMap_FieldPathValue {
	fpv, err := ParseConfigMap_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type ConfigMap_FieldTerminalPathValue struct {
	ConfigMap_FieldTerminalPath
	value interface{}
}

var _ ConfigMap_FieldPathValue = (*ConfigMap_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'ConfigMap' as interface{}
func (fpv *ConfigMap_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *ConfigMap_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *ConfigMap_FieldTerminalPathValue) AsDisplayNameValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *ConfigMap_FieldTerminalPathValue) AsMetadataValue() (*ntt_meta.Meta, bool) {
	res, ok := fpv.value.(*ntt_meta.Meta)
	return res, ok
}
func (fpv *ConfigMap_FieldTerminalPathValue) AsDataValue() (map[string]string, bool) {
	res, ok := fpv.value.(map[string]string)
	return res, ok
}
func (fpv *ConfigMap_FieldTerminalPathValue) AsBinaryDataValue() (map[string][]byte, bool) {
	res, ok := fpv.value.(map[string][]byte)
	return res, ok
}

// SetTo stores value for selected field for object ConfigMap
func (fpv *ConfigMap_FieldTerminalPathValue) SetTo(target **ConfigMap) {
	if *target == nil {
		*target = new(ConfigMap)
	}
	switch fpv.selector {
	case ConfigMap_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case ConfigMap_FieldPathSelectorDisplayName:
		(*target).DisplayName = fpv.value.(string)
	case ConfigMap_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*ntt_meta.Meta)
	case ConfigMap_FieldPathSelectorData:
		(*target).Data = fpv.value.(map[string]string)
	case ConfigMap_FieldPathSelectorBinaryData:
		(*target).BinaryData = fpv.value.(map[string][]byte)
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fpv.selector))
	}
}

func (fpv *ConfigMap_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*ConfigMap)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'ConfigMap_FieldTerminalPathValue' with the value under path in 'ConfigMap'.
func (fpv *ConfigMap_FieldTerminalPathValue) CompareWith(source *ConfigMap) (int, bool) {
	switch fpv.selector {
	case ConfigMap_FieldPathSelectorName:
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
	case ConfigMap_FieldPathSelectorDisplayName:
		leftValue := fpv.value.(string)
		rightValue := source.GetDisplayName()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case ConfigMap_FieldPathSelectorMetadata:
		return 0, false
	case ConfigMap_FieldPathSelectorData:
		return 0, false
	case ConfigMap_FieldPathSelectorBinaryData:
		return 0, false
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fpv.selector))
	}
}

func (fpv *ConfigMap_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*ConfigMap))
}

type ConfigMap_FieldPathMapValue struct {
	ConfigMap_FieldPathMap
	value interface{}
}

var _ ConfigMap_FieldPathValue = (*ConfigMap_FieldPathMapValue)(nil)

// GetValue returns value stored under selected field in ConfigMap as interface{}
func (fpmv *ConfigMap_FieldPathMapValue) GetRawValue() interface{} {
	return fpmv.value
}
func (fpmv *ConfigMap_FieldPathMapValue) AsDataElementValue() (string, bool) {
	res, ok := fpmv.value.(string)
	return res, ok
}
func (fpmv *ConfigMap_FieldPathMapValue) AsBinaryDataElementValue() ([]byte, bool) {
	res, ok := fpmv.value.([]byte)
	return res, ok
}

// SetTo stores value for selected field in ConfigMap
func (fpmv *ConfigMap_FieldPathMapValue) SetTo(target **ConfigMap) {
	if *target == nil {
		*target = new(ConfigMap)
	}
	switch fpmv.selector {
	case ConfigMap_FieldPathSelectorData:
		if (*target).Data == nil {
			(*target).Data = make(map[string]string)
		}
		(*target).Data[fpmv.key] = fpmv.value.(string)
	case ConfigMap_FieldPathSelectorBinaryData:
		if (*target).BinaryData == nil {
			(*target).BinaryData = make(map[string][]byte)
		}
		(*target).BinaryData[fpmv.key] = fpmv.value.([]byte)
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fpmv.selector))
	}
}

func (fpmv *ConfigMap_FieldPathMapValue) SetToRaw(target proto.Message) {
	typedObject := target.(*ConfigMap)
	fpmv.SetTo(&typedObject)
}

// CompareWith compares value in the 'ConfigMap_FieldPathMapValue' with the value under path in 'ConfigMap'.
func (fpmv *ConfigMap_FieldPathMapValue) CompareWith(source *ConfigMap) (int, bool) {
	switch fpmv.selector {
	case ConfigMap_FieldPathSelectorData:
		leftValue := fpmv.value.(string)
		rightValue := source.GetData()[fpmv.key]
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case ConfigMap_FieldPathSelectorBinaryData:
		return 0, false
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fpmv.selector))
	}
}

func (fpmv *ConfigMap_FieldPathMapValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpmv.CompareWith(source.(*ConfigMap))
}

type ConfigMap_FieldSubPathValue struct {
	ConfigMap_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ ConfigMap_FieldPathValue = (*ConfigMap_FieldSubPathValue)(nil)

func (fpvs *ConfigMap_FieldSubPathValue) AsMetadataPathValue() (ntt_meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(ntt_meta.Meta_FieldPathValue)
	return res, ok
}

func (fpvs *ConfigMap_FieldSubPathValue) SetTo(target **ConfigMap) {
	if *target == nil {
		*target = new(ConfigMap)
	}
	switch fpvs.Selector() {
	case ConfigMap_FieldPathSelectorMetadata:
		fpvs.subPathValue.(ntt_meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fpvs.Selector()))
	}
}

func (fpvs *ConfigMap_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*ConfigMap)
	fpvs.SetTo(&typedObject)
}

func (fpvs *ConfigMap_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *ConfigMap_FieldSubPathValue) CompareWith(source *ConfigMap) (int, bool) {
	switch fpvs.Selector() {
	case ConfigMap_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(ntt_meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fpvs.Selector()))
	}
}

func (fpvs *ConfigMap_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*ConfigMap))
}

// ConfigMap_FieldPathArrayItemValue allows storing single item in Path-specific values for ConfigMap according to their type
// Present only for array (repeated) types.
type ConfigMap_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	ConfigMap_FieldPath
	ContainsValue(*ConfigMap) bool
}

// ParseConfigMap_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseConfigMap_FieldPathArrayItemValue(pathStr, valueStr string) (ConfigMap_FieldPathArrayItemValue, error) {
	fp, err := ParseConfigMap_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing ConfigMap field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(ConfigMap_FieldPathArrayItemValue), nil
}

func MustParseConfigMap_FieldPathArrayItemValue(pathStr, valueStr string) ConfigMap_FieldPathArrayItemValue {
	fpaiv, err := ParseConfigMap_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type ConfigMap_FieldTerminalPathArrayItemValue struct {
	ConfigMap_FieldTerminalPath
	value interface{}
}

var _ ConfigMap_FieldPathArrayItemValue = (*ConfigMap_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object ConfigMap as interface{}
func (fpaiv *ConfigMap_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *ConfigMap_FieldTerminalPathArrayItemValue) GetSingle(source *ConfigMap) (interface{}, bool) {
	return nil, false
}

func (fpaiv *ConfigMap_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*ConfigMap))
}

// Contains returns a boolean indicating if value that is being held is present in given 'ConfigMap'
func (fpaiv *ConfigMap_FieldTerminalPathArrayItemValue) ContainsValue(source *ConfigMap) bool {
	slice := fpaiv.ConfigMap_FieldTerminalPath.Get(source)
	for _, v := range slice {
		if reflect.DeepEqual(v, fpaiv.value) {
			return true
		}
	}
	return false
}

type ConfigMap_FieldSubPathArrayItemValue struct {
	ConfigMap_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *ConfigMap_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *ConfigMap_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (ntt_meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(ntt_meta.Meta_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'ConfigMap'
func (fpaivs *ConfigMap_FieldSubPathArrayItemValue) ContainsValue(source *ConfigMap) bool {
	switch fpaivs.Selector() {
	case ConfigMap_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(ntt_meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for ConfigMap: %d", fpaivs.Selector()))
	}
}

// ConfigMap_FieldPathArrayOfValues allows storing slice of values for ConfigMap fields according to their type
type ConfigMap_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	ConfigMap_FieldPath
}

func ParseConfigMap_FieldPathArrayOfValues(pathStr, valuesStr string) (ConfigMap_FieldPathArrayOfValues, error) {
	fp, err := ParseConfigMap_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing ConfigMap field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(ConfigMap_FieldPathArrayOfValues), nil
}

func MustParseConfigMap_FieldPathArrayOfValues(pathStr, valuesStr string) ConfigMap_FieldPathArrayOfValues {
	fpaov, err := ParseConfigMap_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type ConfigMap_FieldTerminalPathArrayOfValues struct {
	ConfigMap_FieldTerminalPath
	values interface{}
}

var _ ConfigMap_FieldPathArrayOfValues = (*ConfigMap_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *ConfigMap_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case ConfigMap_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case ConfigMap_FieldPathSelectorDisplayName:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case ConfigMap_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*ntt_meta.Meta) {
			values = append(values, v)
		}
	case ConfigMap_FieldPathSelectorData:
		for _, v := range fpaov.values.([]map[string]string) {
			values = append(values, v)
		}
	case ConfigMap_FieldPathSelectorBinaryData:
		for _, v := range fpaov.values.([]map[string][]byte) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *ConfigMap_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *ConfigMap_FieldTerminalPathArrayOfValues) AsDisplayNameArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *ConfigMap_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*ntt_meta.Meta, bool) {
	res, ok := fpaov.values.([]*ntt_meta.Meta)
	return res, ok
}
func (fpaov *ConfigMap_FieldTerminalPathArrayOfValues) AsDataArrayOfValues() ([]map[string]string, bool) {
	res, ok := fpaov.values.([]map[string]string)
	return res, ok
}
func (fpaov *ConfigMap_FieldTerminalPathArrayOfValues) AsBinaryDataArrayOfValues() ([]map[string][]byte, bool) {
	res, ok := fpaov.values.([]map[string][]byte)
	return res, ok
}

type ConfigMap_FieldPathMapArrayOfValues struct {
	ConfigMap_FieldPathMap
	values interface{}
}

var _ ConfigMap_FieldPathArrayOfValues = (*ConfigMap_FieldPathMapArrayOfValues)(nil)

func (fpmaov *ConfigMap_FieldPathMapArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpmaov.selector {
	case ConfigMap_FieldPathSelectorData:
		for _, v := range fpmaov.values.([]string) {
			values = append(values, v)
		}
	case ConfigMap_FieldPathSelectorBinaryData:
		for _, v := range fpmaov.values.([][]byte) {
			values = append(values, v)
		}
	}
	return
}
func (fpmaov *ConfigMap_FieldPathMapArrayOfValues) AsDataArrayOfElementValues() ([]string, bool) {
	res, ok := fpmaov.values.([]string)
	return res, ok
}
func (fpmaov *ConfigMap_FieldPathMapArrayOfValues) AsBinaryDataArrayOfElementValues() ([][]byte, bool) {
	res, ok := fpmaov.values.([][]byte)
	return res, ok
}

type ConfigMap_FieldSubPathArrayOfValues struct {
	ConfigMap_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ ConfigMap_FieldPathArrayOfValues = (*ConfigMap_FieldSubPathArrayOfValues)(nil)

func (fpsaov *ConfigMap_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *ConfigMap_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (ntt_meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(ntt_meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
