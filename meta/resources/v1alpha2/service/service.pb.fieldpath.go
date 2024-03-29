// Code generated by protoc-gen-goten-object
// File: edgelq/meta/proto/v1alpha2/service.proto
// DO NOT EDIT!!!

package service

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
	_ = &meta.Meta{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type Service_FieldPath interface {
	gotenobject.FieldPath
	Selector() Service_FieldPathSelector
	Get(source *Service) []interface{}
	GetSingle(source *Service) (interface{}, bool)
	ClearValue(item *Service)

	// Those methods build corresponding Service_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) Service_FieldPathValue
	WithIArrayOfValues(values interface{}) Service_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) Service_FieldPathArrayItemValue
}

type Service_FieldPathSelector int32

const (
	Service_FieldPathSelectorName           Service_FieldPathSelector = 0
	Service_FieldPathSelectorDisplayName    Service_FieldPathSelector = 1
	Service_FieldPathSelectorCurrentVersion Service_FieldPathSelector = 2
	Service_FieldPathSelectorAllVersions    Service_FieldPathSelector = 3
	Service_FieldPathSelectorMetadata       Service_FieldPathSelector = 4
)

func (s Service_FieldPathSelector) String() string {
	switch s {
	case Service_FieldPathSelectorName:
		return "name"
	case Service_FieldPathSelectorDisplayName:
		return "display_name"
	case Service_FieldPathSelectorCurrentVersion:
		return "current_version"
	case Service_FieldPathSelectorAllVersions:
		return "all_versions"
	case Service_FieldPathSelectorMetadata:
		return "metadata"
	default:
		panic(fmt.Sprintf("Invalid selector for Service: %d", s))
	}
}

func BuildService_FieldPath(fp gotenobject.RawFieldPath) (Service_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object Service")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &Service_FieldTerminalPath{selector: Service_FieldPathSelectorName}, nil
		case "display_name", "displayName", "display-name":
			return &Service_FieldTerminalPath{selector: Service_FieldPathSelectorDisplayName}, nil
		case "current_version", "currentVersion", "current-version":
			return &Service_FieldTerminalPath{selector: Service_FieldPathSelectorCurrentVersion}, nil
		case "all_versions", "allVersions", "all-versions":
			return &Service_FieldTerminalPath{selector: Service_FieldPathSelectorAllVersions}, nil
		case "metadata":
			return &Service_FieldTerminalPath{selector: Service_FieldPathSelectorMetadata}, nil
		}
	} else {
		switch fp[0] {
		case "metadata":
			if subpath, err := meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &Service_FieldSubPath{selector: Service_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object Service", fp)
}

func ParseService_FieldPath(rawField string) (Service_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildService_FieldPath(fp)
}

func MustParseService_FieldPath(rawField string) Service_FieldPath {
	fp, err := ParseService_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type Service_FieldTerminalPath struct {
	selector Service_FieldPathSelector
}

var _ Service_FieldPath = (*Service_FieldTerminalPath)(nil)

func (fp *Service_FieldTerminalPath) Selector() Service_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *Service_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *Service_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source Service
func (fp *Service_FieldTerminalPath) Get(source *Service) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case Service_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case Service_FieldPathSelectorDisplayName:
			values = append(values, source.DisplayName)
		case Service_FieldPathSelectorCurrentVersion:
			values = append(values, source.CurrentVersion)
		case Service_FieldPathSelectorAllVersions:
			for _, value := range source.GetAllVersions() {
				values = append(values, value)
			}
		case Service_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		default:
			panic(fmt.Sprintf("Invalid selector for Service: %d", fp.selector))
		}
	}
	return
}

func (fp *Service_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*Service))
}

// GetSingle returns value pointed by specific field of from source Service
func (fp *Service_FieldTerminalPath) GetSingle(source *Service) (interface{}, bool) {
	switch fp.selector {
	case Service_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case Service_FieldPathSelectorDisplayName:
		return source.GetDisplayName(), source != nil
	case Service_FieldPathSelectorCurrentVersion:
		return source.GetCurrentVersion(), source != nil
	case Service_FieldPathSelectorAllVersions:
		res := source.GetAllVersions()
		return res, res != nil
	case Service_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	default:
		panic(fmt.Sprintf("Invalid selector for Service: %d", fp.selector))
	}
}

func (fp *Service_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*Service))
}

// GetDefault returns a default value of the field type
func (fp *Service_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case Service_FieldPathSelectorName:
		return (*Name)(nil)
	case Service_FieldPathSelectorDisplayName:
		return ""
	case Service_FieldPathSelectorCurrentVersion:
		return ""
	case Service_FieldPathSelectorAllVersions:
		return ([]string)(nil)
	case Service_FieldPathSelectorMetadata:
		return (*meta.Meta)(nil)
	default:
		panic(fmt.Sprintf("Invalid selector for Service: %d", fp.selector))
	}
}

func (fp *Service_FieldTerminalPath) ClearValue(item *Service) {
	if item != nil {
		switch fp.selector {
		case Service_FieldPathSelectorName:
			item.Name = nil
		case Service_FieldPathSelectorDisplayName:
			item.DisplayName = ""
		case Service_FieldPathSelectorCurrentVersion:
			item.CurrentVersion = ""
		case Service_FieldPathSelectorAllVersions:
			item.AllVersions = nil
		case Service_FieldPathSelectorMetadata:
			item.Metadata = nil
		default:
			panic(fmt.Sprintf("Invalid selector for Service: %d", fp.selector))
		}
	}
}

func (fp *Service_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*Service))
}

// IsLeaf - whether field path is holds simple value
func (fp *Service_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == Service_FieldPathSelectorName ||
		fp.selector == Service_FieldPathSelectorDisplayName ||
		fp.selector == Service_FieldPathSelectorCurrentVersion ||
		fp.selector == Service_FieldPathSelectorAllVersions
}

func (fp *Service_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *Service_FieldTerminalPath) WithIValue(value interface{}) Service_FieldPathValue {
	switch fp.selector {
	case Service_FieldPathSelectorName:
		return &Service_FieldTerminalPathValue{Service_FieldTerminalPath: *fp, value: value.(*Name)}
	case Service_FieldPathSelectorDisplayName:
		return &Service_FieldTerminalPathValue{Service_FieldTerminalPath: *fp, value: value.(string)}
	case Service_FieldPathSelectorCurrentVersion:
		return &Service_FieldTerminalPathValue{Service_FieldTerminalPath: *fp, value: value.(string)}
	case Service_FieldPathSelectorAllVersions:
		return &Service_FieldTerminalPathValue{Service_FieldTerminalPath: *fp, value: value.([]string)}
	case Service_FieldPathSelectorMetadata:
		return &Service_FieldTerminalPathValue{Service_FieldTerminalPath: *fp, value: value.(*meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for Service: %d", fp.selector))
	}
}

func (fp *Service_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *Service_FieldTerminalPath) WithIArrayOfValues(values interface{}) Service_FieldPathArrayOfValues {
	fpaov := &Service_FieldTerminalPathArrayOfValues{Service_FieldTerminalPath: *fp}
	switch fp.selector {
	case Service_FieldPathSelectorName:
		return &Service_FieldTerminalPathArrayOfValues{Service_FieldTerminalPath: *fp, values: values.([]*Name)}
	case Service_FieldPathSelectorDisplayName:
		return &Service_FieldTerminalPathArrayOfValues{Service_FieldTerminalPath: *fp, values: values.([]string)}
	case Service_FieldPathSelectorCurrentVersion:
		return &Service_FieldTerminalPathArrayOfValues{Service_FieldTerminalPath: *fp, values: values.([]string)}
	case Service_FieldPathSelectorAllVersions:
		return &Service_FieldTerminalPathArrayOfValues{Service_FieldTerminalPath: *fp, values: values.([][]string)}
	case Service_FieldPathSelectorMetadata:
		return &Service_FieldTerminalPathArrayOfValues{Service_FieldTerminalPath: *fp, values: values.([]*meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for Service: %d", fp.selector))
	}
	return fpaov
}

func (fp *Service_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *Service_FieldTerminalPath) WithIArrayItemValue(value interface{}) Service_FieldPathArrayItemValue {
	switch fp.selector {
	case Service_FieldPathSelectorAllVersions:
		return &Service_FieldTerminalPathArrayItemValue{Service_FieldTerminalPath: *fp, value: value.(string)}
	default:
		panic(fmt.Sprintf("Invalid selector for Service: %d", fp.selector))
	}
}

func (fp *Service_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

type Service_FieldSubPath struct {
	selector Service_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ Service_FieldPath = (*Service_FieldSubPath)(nil)

func (fps *Service_FieldSubPath) Selector() Service_FieldPathSelector {
	return fps.selector
}
func (fps *Service_FieldSubPath) AsMetadataSubPath() (meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(meta.Meta_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *Service_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *Service_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source Service
func (fps *Service_FieldSubPath) Get(source *Service) (values []interface{}) {
	switch fps.selector {
	case Service_FieldPathSelectorMetadata:
		values = append(values, fps.subPath.GetRaw(source.GetMetadata())...)
	default:
		panic(fmt.Sprintf("Invalid selector for Service: %d", fps.selector))
	}
	return
}

func (fps *Service_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*Service))
}

// GetSingle returns value of selected field from source Service
func (fps *Service_FieldSubPath) GetSingle(source *Service) (interface{}, bool) {
	switch fps.selector {
	case Service_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Service: %d", fps.selector))
	}
}

func (fps *Service_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*Service))
}

// GetDefault returns a default value of the field type
func (fps *Service_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *Service_FieldSubPath) ClearValue(item *Service) {
	if item != nil {
		switch fps.selector {
		case Service_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		default:
			panic(fmt.Sprintf("Invalid selector for Service: %d", fps.selector))
		}
	}
}

func (fps *Service_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*Service))
}

// IsLeaf - whether field path is holds simple value
func (fps *Service_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *Service_FieldSubPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	iPaths := []gotenobject.FieldPath{&Service_FieldTerminalPath{selector: fps.selector}}
	iPaths = append(iPaths, fps.subPath.SplitIntoTerminalIPaths()...)
	return iPaths
}

func (fps *Service_FieldSubPath) WithIValue(value interface{}) Service_FieldPathValue {
	return &Service_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *Service_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *Service_FieldSubPath) WithIArrayOfValues(values interface{}) Service_FieldPathArrayOfValues {
	return &Service_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *Service_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *Service_FieldSubPath) WithIArrayItemValue(value interface{}) Service_FieldPathArrayItemValue {
	return &Service_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *Service_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// Service_FieldPathValue allows storing values for Service fields according to their type
type Service_FieldPathValue interface {
	Service_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **Service)
	CompareWith(*Service) (cmp int, comparable bool)
}

func ParseService_FieldPathValue(pathStr, valueStr string) (Service_FieldPathValue, error) {
	fp, err := ParseService_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Service field path value from %s: %v", valueStr, err)
	}
	return fpv.(Service_FieldPathValue), nil
}

func MustParseService_FieldPathValue(pathStr, valueStr string) Service_FieldPathValue {
	fpv, err := ParseService_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type Service_FieldTerminalPathValue struct {
	Service_FieldTerminalPath
	value interface{}
}

var _ Service_FieldPathValue = (*Service_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'Service' as interface{}
func (fpv *Service_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *Service_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *Service_FieldTerminalPathValue) AsDisplayNameValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *Service_FieldTerminalPathValue) AsCurrentVersionValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *Service_FieldTerminalPathValue) AsAllVersionsValue() ([]string, bool) {
	res, ok := fpv.value.([]string)
	return res, ok
}
func (fpv *Service_FieldTerminalPathValue) AsMetadataValue() (*meta.Meta, bool) {
	res, ok := fpv.value.(*meta.Meta)
	return res, ok
}

// SetTo stores value for selected field for object Service
func (fpv *Service_FieldTerminalPathValue) SetTo(target **Service) {
	if *target == nil {
		*target = new(Service)
	}
	switch fpv.selector {
	case Service_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case Service_FieldPathSelectorDisplayName:
		(*target).DisplayName = fpv.value.(string)
	case Service_FieldPathSelectorCurrentVersion:
		(*target).CurrentVersion = fpv.value.(string)
	case Service_FieldPathSelectorAllVersions:
		(*target).AllVersions = fpv.value.([]string)
	case Service_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*meta.Meta)
	default:
		panic(fmt.Sprintf("Invalid selector for Service: %d", fpv.selector))
	}
}

func (fpv *Service_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Service)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'Service_FieldTerminalPathValue' with the value under path in 'Service'.
func (fpv *Service_FieldTerminalPathValue) CompareWith(source *Service) (int, bool) {
	switch fpv.selector {
	case Service_FieldPathSelectorName:
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
	case Service_FieldPathSelectorDisplayName:
		leftValue := fpv.value.(string)
		rightValue := source.GetDisplayName()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case Service_FieldPathSelectorCurrentVersion:
		leftValue := fpv.value.(string)
		rightValue := source.GetCurrentVersion()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case Service_FieldPathSelectorAllVersions:
		return 0, false
	case Service_FieldPathSelectorMetadata:
		return 0, false
	default:
		panic(fmt.Sprintf("Invalid selector for Service: %d", fpv.selector))
	}
}

func (fpv *Service_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*Service))
}

type Service_FieldSubPathValue struct {
	Service_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ Service_FieldPathValue = (*Service_FieldSubPathValue)(nil)

func (fpvs *Service_FieldSubPathValue) AsMetadataPathValue() (meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(meta.Meta_FieldPathValue)
	return res, ok
}

func (fpvs *Service_FieldSubPathValue) SetTo(target **Service) {
	if *target == nil {
		*target = new(Service)
	}
	switch fpvs.Selector() {
	case Service_FieldPathSelectorMetadata:
		fpvs.subPathValue.(meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	default:
		panic(fmt.Sprintf("Invalid selector for Service: %d", fpvs.Selector()))
	}
}

func (fpvs *Service_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Service)
	fpvs.SetTo(&typedObject)
}

func (fpvs *Service_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *Service_FieldSubPathValue) CompareWith(source *Service) (int, bool) {
	switch fpvs.Selector() {
	case Service_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Service: %d", fpvs.Selector()))
	}
}

func (fpvs *Service_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*Service))
}

// Service_FieldPathArrayItemValue allows storing single item in Path-specific values for Service according to their type
// Present only for array (repeated) types.
type Service_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	Service_FieldPath
	ContainsValue(*Service) bool
}

// ParseService_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseService_FieldPathArrayItemValue(pathStr, valueStr string) (Service_FieldPathArrayItemValue, error) {
	fp, err := ParseService_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Service field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(Service_FieldPathArrayItemValue), nil
}

func MustParseService_FieldPathArrayItemValue(pathStr, valueStr string) Service_FieldPathArrayItemValue {
	fpaiv, err := ParseService_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type Service_FieldTerminalPathArrayItemValue struct {
	Service_FieldTerminalPath
	value interface{}
}

var _ Service_FieldPathArrayItemValue = (*Service_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object Service as interface{}
func (fpaiv *Service_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}
func (fpaiv *Service_FieldTerminalPathArrayItemValue) AsAllVersionsItemValue() (string, bool) {
	res, ok := fpaiv.value.(string)
	return res, ok
}

func (fpaiv *Service_FieldTerminalPathArrayItemValue) GetSingle(source *Service) (interface{}, bool) {
	return nil, false
}

func (fpaiv *Service_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*Service))
}

// Contains returns a boolean indicating if value that is being held is present in given 'Service'
func (fpaiv *Service_FieldTerminalPathArrayItemValue) ContainsValue(source *Service) bool {
	slice := fpaiv.Service_FieldTerminalPath.Get(source)
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

type Service_FieldSubPathArrayItemValue struct {
	Service_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *Service_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *Service_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'Service'
func (fpaivs *Service_FieldSubPathArrayItemValue) ContainsValue(source *Service) bool {
	switch fpaivs.Selector() {
	case Service_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Service: %d", fpaivs.Selector()))
	}
}

// Service_FieldPathArrayOfValues allows storing slice of values for Service fields according to their type
type Service_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	Service_FieldPath
}

func ParseService_FieldPathArrayOfValues(pathStr, valuesStr string) (Service_FieldPathArrayOfValues, error) {
	fp, err := ParseService_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Service field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(Service_FieldPathArrayOfValues), nil
}

func MustParseService_FieldPathArrayOfValues(pathStr, valuesStr string) Service_FieldPathArrayOfValues {
	fpaov, err := ParseService_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type Service_FieldTerminalPathArrayOfValues struct {
	Service_FieldTerminalPath
	values interface{}
}

var _ Service_FieldPathArrayOfValues = (*Service_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *Service_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case Service_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case Service_FieldPathSelectorDisplayName:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case Service_FieldPathSelectorCurrentVersion:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case Service_FieldPathSelectorAllVersions:
		for _, v := range fpaov.values.([][]string) {
			values = append(values, v)
		}
	case Service_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*meta.Meta) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *Service_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *Service_FieldTerminalPathArrayOfValues) AsDisplayNameArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *Service_FieldTerminalPathArrayOfValues) AsCurrentVersionArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *Service_FieldTerminalPathArrayOfValues) AsAllVersionsArrayOfValues() ([][]string, bool) {
	res, ok := fpaov.values.([][]string)
	return res, ok
}
func (fpaov *Service_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*meta.Meta, bool) {
	res, ok := fpaov.values.([]*meta.Meta)
	return res, ok
}

type Service_FieldSubPathArrayOfValues struct {
	Service_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ Service_FieldPathArrayOfValues = (*Service_FieldSubPathArrayOfValues)(nil)

func (fpsaov *Service_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *Service_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
