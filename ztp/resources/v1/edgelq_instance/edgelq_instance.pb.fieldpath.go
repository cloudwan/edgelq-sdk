// Code generated by protoc-gen-goten-object
// File: edgelq/ztp/proto/v1/edgelq_instance.proto
// DO NOT EDIT!!!

package edgelq_instance

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
	_ = &project.Project{}
	_ = &meta.Meta{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type EdgelqInstance_FieldPath interface {
	gotenobject.FieldPath
	Selector() EdgelqInstance_FieldPathSelector
	Get(source *EdgelqInstance) []interface{}
	GetSingle(source *EdgelqInstance) (interface{}, bool)
	ClearValue(item *EdgelqInstance)

	// Those methods build corresponding EdgelqInstance_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) EdgelqInstance_FieldPathValue
	WithIArrayOfValues(values interface{}) EdgelqInstance_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) EdgelqInstance_FieldPathArrayItemValue
}

type EdgelqInstance_FieldPathSelector int32

const (
	EdgelqInstance_FieldPathSelectorName             EdgelqInstance_FieldPathSelector = 0
	EdgelqInstance_FieldPathSelectorMetadata         EdgelqInstance_FieldPathSelector = 1
	EdgelqInstance_FieldPathSelectorDisplayName      EdgelqInstance_FieldPathSelector = 2
	EdgelqInstance_FieldPathSelectorControllerDomain EdgelqInstance_FieldPathSelector = 3
	EdgelqInstance_FieldPathSelectorEndpoints        EdgelqInstance_FieldPathSelector = 4
)

func (s EdgelqInstance_FieldPathSelector) String() string {
	switch s {
	case EdgelqInstance_FieldPathSelectorName:
		return "name"
	case EdgelqInstance_FieldPathSelectorMetadata:
		return "metadata"
	case EdgelqInstance_FieldPathSelectorDisplayName:
		return "display_name"
	case EdgelqInstance_FieldPathSelectorControllerDomain:
		return "controller_domain"
	case EdgelqInstance_FieldPathSelectorEndpoints:
		return "endpoints"
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", s))
	}
}

func BuildEdgelqInstance_FieldPath(fp gotenobject.RawFieldPath) (EdgelqInstance_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object EdgelqInstance")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &EdgelqInstance_FieldTerminalPath{selector: EdgelqInstance_FieldPathSelectorName}, nil
		case "metadata":
			return &EdgelqInstance_FieldTerminalPath{selector: EdgelqInstance_FieldPathSelectorMetadata}, nil
		case "display_name", "displayName", "display-name":
			return &EdgelqInstance_FieldTerminalPath{selector: EdgelqInstance_FieldPathSelectorDisplayName}, nil
		case "controller_domain", "controllerDomain", "controller-domain":
			return &EdgelqInstance_FieldTerminalPath{selector: EdgelqInstance_FieldPathSelectorControllerDomain}, nil
		case "endpoints":
			return &EdgelqInstance_FieldTerminalPath{selector: EdgelqInstance_FieldPathSelectorEndpoints}, nil
		}
	} else {
		switch fp[0] {
		case "metadata":
			if subpath, err := meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &EdgelqInstance_FieldSubPath{selector: EdgelqInstance_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		case "endpoints":
			if len(fp) > 2 {
				return nil, status.Errorf(codes.InvalidArgument, "sub path for maps ('%s') are not supported (object EdgelqInstance)", fp)
			}
			return &EdgelqInstance_FieldPathMap{selector: EdgelqInstance_FieldPathSelectorEndpoints, key: fp[1]}, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object EdgelqInstance", fp)
}

func ParseEdgelqInstance_FieldPath(rawField string) (EdgelqInstance_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildEdgelqInstance_FieldPath(fp)
}

func MustParseEdgelqInstance_FieldPath(rawField string) EdgelqInstance_FieldPath {
	fp, err := ParseEdgelqInstance_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type EdgelqInstance_FieldTerminalPath struct {
	selector EdgelqInstance_FieldPathSelector
}

var _ EdgelqInstance_FieldPath = (*EdgelqInstance_FieldTerminalPath)(nil)

func (fp *EdgelqInstance_FieldTerminalPath) Selector() EdgelqInstance_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *EdgelqInstance_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *EdgelqInstance_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source EdgelqInstance
func (fp *EdgelqInstance_FieldTerminalPath) Get(source *EdgelqInstance) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case EdgelqInstance_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case EdgelqInstance_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		case EdgelqInstance_FieldPathSelectorDisplayName:
			values = append(values, source.DisplayName)
		case EdgelqInstance_FieldPathSelectorControllerDomain:
			values = append(values, source.ControllerDomain)
		case EdgelqInstance_FieldPathSelectorEndpoints:
			values = append(values, source.Endpoints)
		default:
			panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fp.selector))
		}
	}
	return
}

func (fp *EdgelqInstance_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*EdgelqInstance))
}

// GetSingle returns value pointed by specific field of from source EdgelqInstance
func (fp *EdgelqInstance_FieldTerminalPath) GetSingle(source *EdgelqInstance) (interface{}, bool) {
	switch fp.selector {
	case EdgelqInstance_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case EdgelqInstance_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	case EdgelqInstance_FieldPathSelectorDisplayName:
		return source.GetDisplayName(), source != nil
	case EdgelqInstance_FieldPathSelectorControllerDomain:
		return source.GetControllerDomain(), source != nil
	case EdgelqInstance_FieldPathSelectorEndpoints:
		res := source.GetEndpoints()
		return res, res != nil
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fp.selector))
	}
}

func (fp *EdgelqInstance_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*EdgelqInstance))
}

// GetDefault returns a default value of the field type
func (fp *EdgelqInstance_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case EdgelqInstance_FieldPathSelectorName:
		return (*Name)(nil)
	case EdgelqInstance_FieldPathSelectorMetadata:
		return (*meta.Meta)(nil)
	case EdgelqInstance_FieldPathSelectorDisplayName:
		return ""
	case EdgelqInstance_FieldPathSelectorControllerDomain:
		return ""
	case EdgelqInstance_FieldPathSelectorEndpoints:
		return (map[string]string)(nil)
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fp.selector))
	}
}

func (fp *EdgelqInstance_FieldTerminalPath) ClearValue(item *EdgelqInstance) {
	if item != nil {
		switch fp.selector {
		case EdgelqInstance_FieldPathSelectorName:
			item.Name = nil
		case EdgelqInstance_FieldPathSelectorMetadata:
			item.Metadata = nil
		case EdgelqInstance_FieldPathSelectorDisplayName:
			item.DisplayName = ""
		case EdgelqInstance_FieldPathSelectorControllerDomain:
			item.ControllerDomain = ""
		case EdgelqInstance_FieldPathSelectorEndpoints:
			item.Endpoints = nil
		default:
			panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fp.selector))
		}
	}
}

func (fp *EdgelqInstance_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*EdgelqInstance))
}

// IsLeaf - whether field path is holds simple value
func (fp *EdgelqInstance_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == EdgelqInstance_FieldPathSelectorName ||
		fp.selector == EdgelqInstance_FieldPathSelectorDisplayName ||
		fp.selector == EdgelqInstance_FieldPathSelectorControllerDomain ||
		fp.selector == EdgelqInstance_FieldPathSelectorEndpoints
}

func (fp *EdgelqInstance_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *EdgelqInstance_FieldTerminalPath) WithIValue(value interface{}) EdgelqInstance_FieldPathValue {
	switch fp.selector {
	case EdgelqInstance_FieldPathSelectorName:
		return &EdgelqInstance_FieldTerminalPathValue{EdgelqInstance_FieldTerminalPath: *fp, value: value.(*Name)}
	case EdgelqInstance_FieldPathSelectorMetadata:
		return &EdgelqInstance_FieldTerminalPathValue{EdgelqInstance_FieldTerminalPath: *fp, value: value.(*meta.Meta)}
	case EdgelqInstance_FieldPathSelectorDisplayName:
		return &EdgelqInstance_FieldTerminalPathValue{EdgelqInstance_FieldTerminalPath: *fp, value: value.(string)}
	case EdgelqInstance_FieldPathSelectorControllerDomain:
		return &EdgelqInstance_FieldTerminalPathValue{EdgelqInstance_FieldTerminalPath: *fp, value: value.(string)}
	case EdgelqInstance_FieldPathSelectorEndpoints:
		return &EdgelqInstance_FieldTerminalPathValue{EdgelqInstance_FieldTerminalPath: *fp, value: value.(map[string]string)}
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fp.selector))
	}
}

func (fp *EdgelqInstance_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *EdgelqInstance_FieldTerminalPath) WithIArrayOfValues(values interface{}) EdgelqInstance_FieldPathArrayOfValues {
	fpaov := &EdgelqInstance_FieldTerminalPathArrayOfValues{EdgelqInstance_FieldTerminalPath: *fp}
	switch fp.selector {
	case EdgelqInstance_FieldPathSelectorName:
		return &EdgelqInstance_FieldTerminalPathArrayOfValues{EdgelqInstance_FieldTerminalPath: *fp, values: values.([]*Name)}
	case EdgelqInstance_FieldPathSelectorMetadata:
		return &EdgelqInstance_FieldTerminalPathArrayOfValues{EdgelqInstance_FieldTerminalPath: *fp, values: values.([]*meta.Meta)}
	case EdgelqInstance_FieldPathSelectorDisplayName:
		return &EdgelqInstance_FieldTerminalPathArrayOfValues{EdgelqInstance_FieldTerminalPath: *fp, values: values.([]string)}
	case EdgelqInstance_FieldPathSelectorControllerDomain:
		return &EdgelqInstance_FieldTerminalPathArrayOfValues{EdgelqInstance_FieldTerminalPath: *fp, values: values.([]string)}
	case EdgelqInstance_FieldPathSelectorEndpoints:
		return &EdgelqInstance_FieldTerminalPathArrayOfValues{EdgelqInstance_FieldTerminalPath: *fp, values: values.([]map[string]string)}
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fp.selector))
	}
	return fpaov
}

func (fp *EdgelqInstance_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *EdgelqInstance_FieldTerminalPath) WithIArrayItemValue(value interface{}) EdgelqInstance_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fp.selector))
	}
}

func (fp *EdgelqInstance_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

// FieldPath for map type with additional Key information
type EdgelqInstance_FieldPathMap struct {
	key      string
	selector EdgelqInstance_FieldPathSelector
}

var _ EdgelqInstance_FieldPath = (*EdgelqInstance_FieldPathMap)(nil)

func (fpm *EdgelqInstance_FieldPathMap) Selector() EdgelqInstance_FieldPathSelector {
	return fpm.selector
}

func (fpm *EdgelqInstance_FieldPathMap) Key() string {
	return fpm.key
}

// String returns path representation in proto convention
func (fpm *EdgelqInstance_FieldPathMap) String() string {
	return fpm.selector.String() + "." + fpm.key
}

// JSONString returns path representation is JSON convention. Note that map keys are not transformed
func (fpm *EdgelqInstance_FieldPathMap) JSONString() string {
	return strcase.ToLowerCamel(fpm.selector.String()) + "." + fpm.key
}

// Get returns all values pointed by selected field map key from source EdgelqInstance
func (fpm *EdgelqInstance_FieldPathMap) Get(source *EdgelqInstance) (values []interface{}) {
	switch fpm.selector {
	case EdgelqInstance_FieldPathSelectorEndpoints:
		if value, ok := source.GetEndpoints()[fpm.key]; ok {
			values = append(values, value)
		}
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fpm.selector))
	}
	return
}

func (fpm *EdgelqInstance_FieldPathMap) GetRaw(source proto.Message) []interface{} {
	return fpm.Get(source.(*EdgelqInstance))
}

// GetSingle returns value by selected field map key from source EdgelqInstance
func (fpm *EdgelqInstance_FieldPathMap) GetSingle(source *EdgelqInstance) (interface{}, bool) {
	switch fpm.selector {
	case EdgelqInstance_FieldPathSelectorEndpoints:
		res, ok := source.GetEndpoints()[fpm.key]
		return res, ok
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fpm.selector))
	}
}

func (fpm *EdgelqInstance_FieldPathMap) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpm.GetSingle(source.(*EdgelqInstance))
}

// GetDefault returns a default value of the field type
func (fpm *EdgelqInstance_FieldPathMap) GetDefault() interface{} {
	switch fpm.selector {
	case EdgelqInstance_FieldPathSelectorEndpoints:
		var v string
		return v
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fpm.selector))
	}
}

func (fpm *EdgelqInstance_FieldPathMap) ClearValue(item *EdgelqInstance) {
	if item != nil {
		switch fpm.selector {
		case EdgelqInstance_FieldPathSelectorEndpoints:
			delete(item.Endpoints, fpm.key)
		default:
			panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fpm.selector))
		}
	}
}

func (fpm *EdgelqInstance_FieldPathMap) ClearValueRaw(item proto.Message) {
	fpm.ClearValue(item.(*EdgelqInstance))
}

// IsLeaf - whether field path is holds simple value
func (fpm *EdgelqInstance_FieldPathMap) IsLeaf() bool {
	switch fpm.selector {
	case EdgelqInstance_FieldPathSelectorEndpoints:
		return true
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fpm.selector))
	}
}

func (fpm *EdgelqInstance_FieldPathMap) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fpm}
}

func (fpm *EdgelqInstance_FieldPathMap) WithIValue(value interface{}) EdgelqInstance_FieldPathValue {
	switch fpm.selector {
	case EdgelqInstance_FieldPathSelectorEndpoints:
		return &EdgelqInstance_FieldPathMapValue{EdgelqInstance_FieldPathMap: *fpm, value: value.(string)}
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fpm.selector))
	}
}

func (fpm *EdgelqInstance_FieldPathMap) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fpm.WithIValue(value)
}

func (fpm *EdgelqInstance_FieldPathMap) WithIArrayOfValues(values interface{}) EdgelqInstance_FieldPathArrayOfValues {
	switch fpm.selector {
	case EdgelqInstance_FieldPathSelectorEndpoints:
		return &EdgelqInstance_FieldPathMapArrayOfValues{EdgelqInstance_FieldPathMap: *fpm, values: values.([]string)}
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fpm.selector))
	}
}

func (fpm *EdgelqInstance_FieldPathMap) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fpm.WithIArrayOfValues(values)
}

func (fpm *EdgelqInstance_FieldPathMap) WithIArrayItemValue(value interface{}) EdgelqInstance_FieldPathArrayItemValue {
	panic("Cannot create array item value from map fieldpath")
}

func (fpm *EdgelqInstance_FieldPathMap) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fpm.WithIArrayItemValue(value)
}

type EdgelqInstance_FieldSubPath struct {
	selector EdgelqInstance_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ EdgelqInstance_FieldPath = (*EdgelqInstance_FieldSubPath)(nil)

func (fps *EdgelqInstance_FieldSubPath) Selector() EdgelqInstance_FieldPathSelector {
	return fps.selector
}
func (fps *EdgelqInstance_FieldSubPath) AsMetadataSubPath() (meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(meta.Meta_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *EdgelqInstance_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *EdgelqInstance_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source EdgelqInstance
func (fps *EdgelqInstance_FieldSubPath) Get(source *EdgelqInstance) (values []interface{}) {
	switch fps.selector {
	case EdgelqInstance_FieldPathSelectorMetadata:
		values = append(values, fps.subPath.GetRaw(source.GetMetadata())...)
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fps.selector))
	}
	return
}

func (fps *EdgelqInstance_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*EdgelqInstance))
}

// GetSingle returns value of selected field from source EdgelqInstance
func (fps *EdgelqInstance_FieldSubPath) GetSingle(source *EdgelqInstance) (interface{}, bool) {
	switch fps.selector {
	case EdgelqInstance_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fps.selector))
	}
}

func (fps *EdgelqInstance_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*EdgelqInstance))
}

// GetDefault returns a default value of the field type
func (fps *EdgelqInstance_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *EdgelqInstance_FieldSubPath) ClearValue(item *EdgelqInstance) {
	if item != nil {
		switch fps.selector {
		case EdgelqInstance_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		default:
			panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fps.selector))
		}
	}
}

func (fps *EdgelqInstance_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*EdgelqInstance))
}

// IsLeaf - whether field path is holds simple value
func (fps *EdgelqInstance_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *EdgelqInstance_FieldSubPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	iPaths := []gotenobject.FieldPath{&EdgelqInstance_FieldTerminalPath{selector: fps.selector}}
	iPaths = append(iPaths, fps.subPath.SplitIntoTerminalIPaths()...)
	return iPaths
}

func (fps *EdgelqInstance_FieldSubPath) WithIValue(value interface{}) EdgelqInstance_FieldPathValue {
	return &EdgelqInstance_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *EdgelqInstance_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *EdgelqInstance_FieldSubPath) WithIArrayOfValues(values interface{}) EdgelqInstance_FieldPathArrayOfValues {
	return &EdgelqInstance_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *EdgelqInstance_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *EdgelqInstance_FieldSubPath) WithIArrayItemValue(value interface{}) EdgelqInstance_FieldPathArrayItemValue {
	return &EdgelqInstance_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *EdgelqInstance_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// EdgelqInstance_FieldPathValue allows storing values for EdgelqInstance fields according to their type
type EdgelqInstance_FieldPathValue interface {
	EdgelqInstance_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **EdgelqInstance)
	CompareWith(*EdgelqInstance) (cmp int, comparable bool)
}

func ParseEdgelqInstance_FieldPathValue(pathStr, valueStr string) (EdgelqInstance_FieldPathValue, error) {
	fp, err := ParseEdgelqInstance_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing EdgelqInstance field path value from %s: %v", valueStr, err)
	}
	return fpv.(EdgelqInstance_FieldPathValue), nil
}

func MustParseEdgelqInstance_FieldPathValue(pathStr, valueStr string) EdgelqInstance_FieldPathValue {
	fpv, err := ParseEdgelqInstance_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type EdgelqInstance_FieldTerminalPathValue struct {
	EdgelqInstance_FieldTerminalPath
	value interface{}
}

var _ EdgelqInstance_FieldPathValue = (*EdgelqInstance_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'EdgelqInstance' as interface{}
func (fpv *EdgelqInstance_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *EdgelqInstance_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *EdgelqInstance_FieldTerminalPathValue) AsMetadataValue() (*meta.Meta, bool) {
	res, ok := fpv.value.(*meta.Meta)
	return res, ok
}
func (fpv *EdgelqInstance_FieldTerminalPathValue) AsDisplayNameValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *EdgelqInstance_FieldTerminalPathValue) AsControllerDomainValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *EdgelqInstance_FieldTerminalPathValue) AsEndpointsValue() (map[string]string, bool) {
	res, ok := fpv.value.(map[string]string)
	return res, ok
}

// SetTo stores value for selected field for object EdgelqInstance
func (fpv *EdgelqInstance_FieldTerminalPathValue) SetTo(target **EdgelqInstance) {
	if *target == nil {
		*target = new(EdgelqInstance)
	}
	switch fpv.selector {
	case EdgelqInstance_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case EdgelqInstance_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*meta.Meta)
	case EdgelqInstance_FieldPathSelectorDisplayName:
		(*target).DisplayName = fpv.value.(string)
	case EdgelqInstance_FieldPathSelectorControllerDomain:
		(*target).ControllerDomain = fpv.value.(string)
	case EdgelqInstance_FieldPathSelectorEndpoints:
		(*target).Endpoints = fpv.value.(map[string]string)
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fpv.selector))
	}
}

func (fpv *EdgelqInstance_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*EdgelqInstance)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'EdgelqInstance_FieldTerminalPathValue' with the value under path in 'EdgelqInstance'.
func (fpv *EdgelqInstance_FieldTerminalPathValue) CompareWith(source *EdgelqInstance) (int, bool) {
	switch fpv.selector {
	case EdgelqInstance_FieldPathSelectorName:
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
	case EdgelqInstance_FieldPathSelectorMetadata:
		return 0, false
	case EdgelqInstance_FieldPathSelectorDisplayName:
		leftValue := fpv.value.(string)
		rightValue := source.GetDisplayName()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case EdgelqInstance_FieldPathSelectorControllerDomain:
		leftValue := fpv.value.(string)
		rightValue := source.GetControllerDomain()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case EdgelqInstance_FieldPathSelectorEndpoints:
		return 0, false
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fpv.selector))
	}
}

func (fpv *EdgelqInstance_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*EdgelqInstance))
}

type EdgelqInstance_FieldPathMapValue struct {
	EdgelqInstance_FieldPathMap
	value interface{}
}

var _ EdgelqInstance_FieldPathValue = (*EdgelqInstance_FieldPathMapValue)(nil)

// GetValue returns value stored under selected field in EdgelqInstance as interface{}
func (fpmv *EdgelqInstance_FieldPathMapValue) GetRawValue() interface{} {
	return fpmv.value
}
func (fpmv *EdgelqInstance_FieldPathMapValue) AsEndpointsElementValue() (string, bool) {
	res, ok := fpmv.value.(string)
	return res, ok
}

// SetTo stores value for selected field in EdgelqInstance
func (fpmv *EdgelqInstance_FieldPathMapValue) SetTo(target **EdgelqInstance) {
	if *target == nil {
		*target = new(EdgelqInstance)
	}
	switch fpmv.selector {
	case EdgelqInstance_FieldPathSelectorEndpoints:
		if (*target).Endpoints == nil {
			(*target).Endpoints = make(map[string]string)
		}
		(*target).Endpoints[fpmv.key] = fpmv.value.(string)
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fpmv.selector))
	}
}

func (fpmv *EdgelqInstance_FieldPathMapValue) SetToRaw(target proto.Message) {
	typedObject := target.(*EdgelqInstance)
	fpmv.SetTo(&typedObject)
}

// CompareWith compares value in the 'EdgelqInstance_FieldPathMapValue' with the value under path in 'EdgelqInstance'.
func (fpmv *EdgelqInstance_FieldPathMapValue) CompareWith(source *EdgelqInstance) (int, bool) {
	switch fpmv.selector {
	case EdgelqInstance_FieldPathSelectorEndpoints:
		leftValue := fpmv.value.(string)
		rightValue := source.GetEndpoints()[fpmv.key]
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fpmv.selector))
	}
}

func (fpmv *EdgelqInstance_FieldPathMapValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpmv.CompareWith(source.(*EdgelqInstance))
}

type EdgelqInstance_FieldSubPathValue struct {
	EdgelqInstance_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ EdgelqInstance_FieldPathValue = (*EdgelqInstance_FieldSubPathValue)(nil)

func (fpvs *EdgelqInstance_FieldSubPathValue) AsMetadataPathValue() (meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(meta.Meta_FieldPathValue)
	return res, ok
}

func (fpvs *EdgelqInstance_FieldSubPathValue) SetTo(target **EdgelqInstance) {
	if *target == nil {
		*target = new(EdgelqInstance)
	}
	switch fpvs.Selector() {
	case EdgelqInstance_FieldPathSelectorMetadata:
		fpvs.subPathValue.(meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fpvs.Selector()))
	}
}

func (fpvs *EdgelqInstance_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*EdgelqInstance)
	fpvs.SetTo(&typedObject)
}

func (fpvs *EdgelqInstance_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *EdgelqInstance_FieldSubPathValue) CompareWith(source *EdgelqInstance) (int, bool) {
	switch fpvs.Selector() {
	case EdgelqInstance_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fpvs.Selector()))
	}
}

func (fpvs *EdgelqInstance_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*EdgelqInstance))
}

// EdgelqInstance_FieldPathArrayItemValue allows storing single item in Path-specific values for EdgelqInstance according to their type
// Present only for array (repeated) types.
type EdgelqInstance_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	EdgelqInstance_FieldPath
	ContainsValue(*EdgelqInstance) bool
}

// ParseEdgelqInstance_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseEdgelqInstance_FieldPathArrayItemValue(pathStr, valueStr string) (EdgelqInstance_FieldPathArrayItemValue, error) {
	fp, err := ParseEdgelqInstance_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing EdgelqInstance field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(EdgelqInstance_FieldPathArrayItemValue), nil
}

func MustParseEdgelqInstance_FieldPathArrayItemValue(pathStr, valueStr string) EdgelqInstance_FieldPathArrayItemValue {
	fpaiv, err := ParseEdgelqInstance_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type EdgelqInstance_FieldTerminalPathArrayItemValue struct {
	EdgelqInstance_FieldTerminalPath
	value interface{}
}

var _ EdgelqInstance_FieldPathArrayItemValue = (*EdgelqInstance_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object EdgelqInstance as interface{}
func (fpaiv *EdgelqInstance_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *EdgelqInstance_FieldTerminalPathArrayItemValue) GetSingle(source *EdgelqInstance) (interface{}, bool) {
	return nil, false
}

func (fpaiv *EdgelqInstance_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*EdgelqInstance))
}

// Contains returns a boolean indicating if value that is being held is present in given 'EdgelqInstance'
func (fpaiv *EdgelqInstance_FieldTerminalPathArrayItemValue) ContainsValue(source *EdgelqInstance) bool {
	slice := fpaiv.EdgelqInstance_FieldTerminalPath.Get(source)
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

type EdgelqInstance_FieldSubPathArrayItemValue struct {
	EdgelqInstance_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *EdgelqInstance_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *EdgelqInstance_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'EdgelqInstance'
func (fpaivs *EdgelqInstance_FieldSubPathArrayItemValue) ContainsValue(source *EdgelqInstance) bool {
	switch fpaivs.Selector() {
	case EdgelqInstance_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for EdgelqInstance: %d", fpaivs.Selector()))
	}
}

// EdgelqInstance_FieldPathArrayOfValues allows storing slice of values for EdgelqInstance fields according to their type
type EdgelqInstance_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	EdgelqInstance_FieldPath
}

func ParseEdgelqInstance_FieldPathArrayOfValues(pathStr, valuesStr string) (EdgelqInstance_FieldPathArrayOfValues, error) {
	fp, err := ParseEdgelqInstance_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing EdgelqInstance field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(EdgelqInstance_FieldPathArrayOfValues), nil
}

func MustParseEdgelqInstance_FieldPathArrayOfValues(pathStr, valuesStr string) EdgelqInstance_FieldPathArrayOfValues {
	fpaov, err := ParseEdgelqInstance_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type EdgelqInstance_FieldTerminalPathArrayOfValues struct {
	EdgelqInstance_FieldTerminalPath
	values interface{}
}

var _ EdgelqInstance_FieldPathArrayOfValues = (*EdgelqInstance_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *EdgelqInstance_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case EdgelqInstance_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case EdgelqInstance_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*meta.Meta) {
			values = append(values, v)
		}
	case EdgelqInstance_FieldPathSelectorDisplayName:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case EdgelqInstance_FieldPathSelectorControllerDomain:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case EdgelqInstance_FieldPathSelectorEndpoints:
		for _, v := range fpaov.values.([]map[string]string) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *EdgelqInstance_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *EdgelqInstance_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*meta.Meta, bool) {
	res, ok := fpaov.values.([]*meta.Meta)
	return res, ok
}
func (fpaov *EdgelqInstance_FieldTerminalPathArrayOfValues) AsDisplayNameArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *EdgelqInstance_FieldTerminalPathArrayOfValues) AsControllerDomainArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *EdgelqInstance_FieldTerminalPathArrayOfValues) AsEndpointsArrayOfValues() ([]map[string]string, bool) {
	res, ok := fpaov.values.([]map[string]string)
	return res, ok
}

type EdgelqInstance_FieldPathMapArrayOfValues struct {
	EdgelqInstance_FieldPathMap
	values interface{}
}

var _ EdgelqInstance_FieldPathArrayOfValues = (*EdgelqInstance_FieldPathMapArrayOfValues)(nil)

func (fpmaov *EdgelqInstance_FieldPathMapArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpmaov.selector {
	case EdgelqInstance_FieldPathSelectorEndpoints:
		for _, v := range fpmaov.values.([]string) {
			values = append(values, v)
		}
	}
	return
}
func (fpmaov *EdgelqInstance_FieldPathMapArrayOfValues) AsEndpointsArrayOfElementValues() ([]string, bool) {
	res, ok := fpmaov.values.([]string)
	return res, ok
}

type EdgelqInstance_FieldSubPathArrayOfValues struct {
	EdgelqInstance_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ EdgelqInstance_FieldPathArrayOfValues = (*EdgelqInstance_FieldSubPathArrayOfValues)(nil)

func (fpsaov *EdgelqInstance_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *EdgelqInstance_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
