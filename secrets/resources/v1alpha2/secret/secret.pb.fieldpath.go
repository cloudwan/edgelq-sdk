// Code generated by protoc-gen-goten-object
// File: edgelq/secrets/proto/v1alpha2/secret.proto
// DO NOT EDIT!!!

package secret

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
	project "github.com/cloudwan/edgelq-sdk/secrets/resources/v1alpha2/project"
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
type Secret_FieldPath interface {
	gotenobject.FieldPath
	Selector() Secret_FieldPathSelector
	Get(source *Secret) []interface{}
	GetSingle(source *Secret) (interface{}, bool)
	ClearValue(item *Secret)

	// Those methods build corresponding Secret_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) Secret_FieldPathValue
	WithIArrayOfValues(values interface{}) Secret_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) Secret_FieldPathArrayItemValue
}

type Secret_FieldPathSelector int32

const (
	Secret_FieldPathSelectorName     Secret_FieldPathSelector = 0
	Secret_FieldPathSelectorEncData  Secret_FieldPathSelector = 1
	Secret_FieldPathSelectorData     Secret_FieldPathSelector = 2
	Secret_FieldPathSelectorMetadata Secret_FieldPathSelector = 3
)

func (s Secret_FieldPathSelector) String() string {
	switch s {
	case Secret_FieldPathSelectorName:
		return "name"
	case Secret_FieldPathSelectorEncData:
		return "enc_data"
	case Secret_FieldPathSelectorData:
		return "data"
	case Secret_FieldPathSelectorMetadata:
		return "metadata"
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", s))
	}
}

func BuildSecret_FieldPath(fp gotenobject.RawFieldPath) (Secret_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object Secret")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &Secret_FieldTerminalPath{selector: Secret_FieldPathSelectorName}, nil
		case "enc_data", "encData", "enc-data":
			return &Secret_FieldTerminalPath{selector: Secret_FieldPathSelectorEncData}, nil
		case "data":
			return &Secret_FieldTerminalPath{selector: Secret_FieldPathSelectorData}, nil
		case "metadata":
			return &Secret_FieldTerminalPath{selector: Secret_FieldPathSelectorMetadata}, nil
		}
	} else {
		switch fp[0] {
		case "metadata":
			if subpath, err := meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &Secret_FieldSubPath{selector: Secret_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		case "data":
			if len(fp) > 2 {
				return nil, status.Errorf(codes.InvalidArgument, "sub path for maps ('%s') are not supported (object Secret)", fp)
			}
			return &Secret_FieldPathMap{selector: Secret_FieldPathSelectorData, key: fp[1]}, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object Secret", fp)
}

func ParseSecret_FieldPath(rawField string) (Secret_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildSecret_FieldPath(fp)
}

func MustParseSecret_FieldPath(rawField string) Secret_FieldPath {
	fp, err := ParseSecret_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type Secret_FieldTerminalPath struct {
	selector Secret_FieldPathSelector
}

var _ Secret_FieldPath = (*Secret_FieldTerminalPath)(nil)

func (fp *Secret_FieldTerminalPath) Selector() Secret_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *Secret_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *Secret_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source Secret
func (fp *Secret_FieldTerminalPath) Get(source *Secret) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case Secret_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case Secret_FieldPathSelectorEncData:
			values = append(values, source.EncData)
		case Secret_FieldPathSelectorData:
			values = append(values, source.Data)
		case Secret_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		default:
			panic(fmt.Sprintf("Invalid selector for Secret: %d", fp.selector))
		}
	}
	return
}

func (fp *Secret_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*Secret))
}

// GetSingle returns value pointed by specific field of from source Secret
func (fp *Secret_FieldTerminalPath) GetSingle(source *Secret) (interface{}, bool) {
	switch fp.selector {
	case Secret_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case Secret_FieldPathSelectorEncData:
		res := source.GetEncData()
		return res, res != nil
	case Secret_FieldPathSelectorData:
		res := source.GetData()
		return res, res != nil
	case Secret_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fp.selector))
	}
}

func (fp *Secret_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*Secret))
}

// GetDefault returns a default value of the field type
func (fp *Secret_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case Secret_FieldPathSelectorName:
		return (*Name)(nil)
	case Secret_FieldPathSelectorEncData:
		return ([]byte)(nil)
	case Secret_FieldPathSelectorData:
		return (map[string]string)(nil)
	case Secret_FieldPathSelectorMetadata:
		return (*meta.Meta)(nil)
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fp.selector))
	}
}

func (fp *Secret_FieldTerminalPath) ClearValue(item *Secret) {
	if item != nil {
		switch fp.selector {
		case Secret_FieldPathSelectorName:
			item.Name = nil
		case Secret_FieldPathSelectorEncData:
			item.EncData = nil
		case Secret_FieldPathSelectorData:
			item.Data = nil
		case Secret_FieldPathSelectorMetadata:
			item.Metadata = nil
		default:
			panic(fmt.Sprintf("Invalid selector for Secret: %d", fp.selector))
		}
	}
}

func (fp *Secret_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*Secret))
}

// IsLeaf - whether field path is holds simple value
func (fp *Secret_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == Secret_FieldPathSelectorName ||
		fp.selector == Secret_FieldPathSelectorEncData ||
		fp.selector == Secret_FieldPathSelectorData
}

func (fp *Secret_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *Secret_FieldTerminalPath) WithIValue(value interface{}) Secret_FieldPathValue {
	switch fp.selector {
	case Secret_FieldPathSelectorName:
		return &Secret_FieldTerminalPathValue{Secret_FieldTerminalPath: *fp, value: value.(*Name)}
	case Secret_FieldPathSelectorEncData:
		return &Secret_FieldTerminalPathValue{Secret_FieldTerminalPath: *fp, value: value.([]byte)}
	case Secret_FieldPathSelectorData:
		return &Secret_FieldTerminalPathValue{Secret_FieldTerminalPath: *fp, value: value.(map[string]string)}
	case Secret_FieldPathSelectorMetadata:
		return &Secret_FieldTerminalPathValue{Secret_FieldTerminalPath: *fp, value: value.(*meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fp.selector))
	}
}

func (fp *Secret_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *Secret_FieldTerminalPath) WithIArrayOfValues(values interface{}) Secret_FieldPathArrayOfValues {
	fpaov := &Secret_FieldTerminalPathArrayOfValues{Secret_FieldTerminalPath: *fp}
	switch fp.selector {
	case Secret_FieldPathSelectorName:
		return &Secret_FieldTerminalPathArrayOfValues{Secret_FieldTerminalPath: *fp, values: values.([]*Name)}
	case Secret_FieldPathSelectorEncData:
		return &Secret_FieldTerminalPathArrayOfValues{Secret_FieldTerminalPath: *fp, values: values.([][]byte)}
	case Secret_FieldPathSelectorData:
		return &Secret_FieldTerminalPathArrayOfValues{Secret_FieldTerminalPath: *fp, values: values.([]map[string]string)}
	case Secret_FieldPathSelectorMetadata:
		return &Secret_FieldTerminalPathArrayOfValues{Secret_FieldTerminalPath: *fp, values: values.([]*meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fp.selector))
	}
	return fpaov
}

func (fp *Secret_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *Secret_FieldTerminalPath) WithIArrayItemValue(value interface{}) Secret_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fp.selector))
	}
}

func (fp *Secret_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

// FieldPath for map type with additional Key information
type Secret_FieldPathMap struct {
	key      string
	selector Secret_FieldPathSelector
}

var _ Secret_FieldPath = (*Secret_FieldPathMap)(nil)

func (fpm *Secret_FieldPathMap) Selector() Secret_FieldPathSelector {
	return fpm.selector
}

func (fpm *Secret_FieldPathMap) Key() string {
	return fpm.key
}

// String returns path representation in proto convention
func (fpm *Secret_FieldPathMap) String() string {
	return fpm.selector.String() + "." + fpm.key
}

// JSONString returns path representation is JSON convention. Note that map keys are not transformed
func (fpm *Secret_FieldPathMap) JSONString() string {
	return strcase.ToLowerCamel(fpm.selector.String()) + "." + fpm.key
}

// Get returns all values pointed by selected field map key from source Secret
func (fpm *Secret_FieldPathMap) Get(source *Secret) (values []interface{}) {
	switch fpm.selector {
	case Secret_FieldPathSelectorData:
		if value, ok := source.GetData()[fpm.key]; ok {
			values = append(values, value)
		}
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fpm.selector))
	}
	return
}

func (fpm *Secret_FieldPathMap) GetRaw(source proto.Message) []interface{} {
	return fpm.Get(source.(*Secret))
}

// GetSingle returns value by selected field map key from source Secret
func (fpm *Secret_FieldPathMap) GetSingle(source *Secret) (interface{}, bool) {
	switch fpm.selector {
	case Secret_FieldPathSelectorData:
		res, ok := source.GetData()[fpm.key]
		return res, ok
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fpm.selector))
	}
}

func (fpm *Secret_FieldPathMap) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpm.GetSingle(source.(*Secret))
}

// GetDefault returns a default value of the field type
func (fpm *Secret_FieldPathMap) GetDefault() interface{} {
	switch fpm.selector {
	case Secret_FieldPathSelectorData:
		var v string
		return v
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fpm.selector))
	}
}

func (fpm *Secret_FieldPathMap) ClearValue(item *Secret) {
	if item != nil {
		switch fpm.selector {
		case Secret_FieldPathSelectorData:
			delete(item.Data, fpm.key)
		default:
			panic(fmt.Sprintf("Invalid selector for Secret: %d", fpm.selector))
		}
	}
}

func (fpm *Secret_FieldPathMap) ClearValueRaw(item proto.Message) {
	fpm.ClearValue(item.(*Secret))
}

// IsLeaf - whether field path is holds simple value
func (fpm *Secret_FieldPathMap) IsLeaf() bool {
	switch fpm.selector {
	case Secret_FieldPathSelectorData:
		return true
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fpm.selector))
	}
}

func (fpm *Secret_FieldPathMap) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fpm}
}

func (fpm *Secret_FieldPathMap) WithIValue(value interface{}) Secret_FieldPathValue {
	switch fpm.selector {
	case Secret_FieldPathSelectorData:
		return &Secret_FieldPathMapValue{Secret_FieldPathMap: *fpm, value: value.(string)}
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fpm.selector))
	}
}

func (fpm *Secret_FieldPathMap) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fpm.WithIValue(value)
}

func (fpm *Secret_FieldPathMap) WithIArrayOfValues(values interface{}) Secret_FieldPathArrayOfValues {
	switch fpm.selector {
	case Secret_FieldPathSelectorData:
		return &Secret_FieldPathMapArrayOfValues{Secret_FieldPathMap: *fpm, values: values.([]string)}
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fpm.selector))
	}
}

func (fpm *Secret_FieldPathMap) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fpm.WithIArrayOfValues(values)
}

func (fpm *Secret_FieldPathMap) WithIArrayItemValue(value interface{}) Secret_FieldPathArrayItemValue {
	panic("Cannot create array item value from map fieldpath")
}

func (fpm *Secret_FieldPathMap) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fpm.WithIArrayItemValue(value)
}

type Secret_FieldSubPath struct {
	selector Secret_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ Secret_FieldPath = (*Secret_FieldSubPath)(nil)

func (fps *Secret_FieldSubPath) Selector() Secret_FieldPathSelector {
	return fps.selector
}
func (fps *Secret_FieldSubPath) AsMetadataSubPath() (meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(meta.Meta_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *Secret_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *Secret_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source Secret
func (fps *Secret_FieldSubPath) Get(source *Secret) (values []interface{}) {
	switch fps.selector {
	case Secret_FieldPathSelectorMetadata:
		values = append(values, fps.subPath.GetRaw(source.GetMetadata())...)
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fps.selector))
	}
	return
}

func (fps *Secret_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*Secret))
}

// GetSingle returns value of selected field from source Secret
func (fps *Secret_FieldSubPath) GetSingle(source *Secret) (interface{}, bool) {
	switch fps.selector {
	case Secret_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fps.selector))
	}
}

func (fps *Secret_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*Secret))
}

// GetDefault returns a default value of the field type
func (fps *Secret_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *Secret_FieldSubPath) ClearValue(item *Secret) {
	if item != nil {
		switch fps.selector {
		case Secret_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		default:
			panic(fmt.Sprintf("Invalid selector for Secret: %d", fps.selector))
		}
	}
}

func (fps *Secret_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*Secret))
}

// IsLeaf - whether field path is holds simple value
func (fps *Secret_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *Secret_FieldSubPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	iPaths := []gotenobject.FieldPath{&Secret_FieldTerminalPath{selector: fps.selector}}
	iPaths = append(iPaths, fps.subPath.SplitIntoTerminalIPaths()...)
	return iPaths
}

func (fps *Secret_FieldSubPath) WithIValue(value interface{}) Secret_FieldPathValue {
	return &Secret_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *Secret_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *Secret_FieldSubPath) WithIArrayOfValues(values interface{}) Secret_FieldPathArrayOfValues {
	return &Secret_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *Secret_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *Secret_FieldSubPath) WithIArrayItemValue(value interface{}) Secret_FieldPathArrayItemValue {
	return &Secret_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *Secret_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// Secret_FieldPathValue allows storing values for Secret fields according to their type
type Secret_FieldPathValue interface {
	Secret_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **Secret)
	CompareWith(*Secret) (cmp int, comparable bool)
}

func ParseSecret_FieldPathValue(pathStr, valueStr string) (Secret_FieldPathValue, error) {
	fp, err := ParseSecret_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Secret field path value from %s: %v", valueStr, err)
	}
	return fpv.(Secret_FieldPathValue), nil
}

func MustParseSecret_FieldPathValue(pathStr, valueStr string) Secret_FieldPathValue {
	fpv, err := ParseSecret_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type Secret_FieldTerminalPathValue struct {
	Secret_FieldTerminalPath
	value interface{}
}

var _ Secret_FieldPathValue = (*Secret_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'Secret' as interface{}
func (fpv *Secret_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *Secret_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *Secret_FieldTerminalPathValue) AsEncDataValue() ([]byte, bool) {
	res, ok := fpv.value.([]byte)
	return res, ok
}
func (fpv *Secret_FieldTerminalPathValue) AsDataValue() (map[string]string, bool) {
	res, ok := fpv.value.(map[string]string)
	return res, ok
}
func (fpv *Secret_FieldTerminalPathValue) AsMetadataValue() (*meta.Meta, bool) {
	res, ok := fpv.value.(*meta.Meta)
	return res, ok
}

// SetTo stores value for selected field for object Secret
func (fpv *Secret_FieldTerminalPathValue) SetTo(target **Secret) {
	if *target == nil {
		*target = new(Secret)
	}
	switch fpv.selector {
	case Secret_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case Secret_FieldPathSelectorEncData:
		(*target).EncData = fpv.value.([]byte)
	case Secret_FieldPathSelectorData:
		(*target).Data = fpv.value.(map[string]string)
	case Secret_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*meta.Meta)
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fpv.selector))
	}
}

func (fpv *Secret_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Secret)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'Secret_FieldTerminalPathValue' with the value under path in 'Secret'.
func (fpv *Secret_FieldTerminalPathValue) CompareWith(source *Secret) (int, bool) {
	switch fpv.selector {
	case Secret_FieldPathSelectorName:
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
	case Secret_FieldPathSelectorEncData:
		return 0, false
	case Secret_FieldPathSelectorData:
		return 0, false
	case Secret_FieldPathSelectorMetadata:
		return 0, false
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fpv.selector))
	}
}

func (fpv *Secret_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*Secret))
}

type Secret_FieldPathMapValue struct {
	Secret_FieldPathMap
	value interface{}
}

var _ Secret_FieldPathValue = (*Secret_FieldPathMapValue)(nil)

// GetValue returns value stored under selected field in Secret as interface{}
func (fpmv *Secret_FieldPathMapValue) GetRawValue() interface{} {
	return fpmv.value
}
func (fpmv *Secret_FieldPathMapValue) AsDataElementValue() (string, bool) {
	res, ok := fpmv.value.(string)
	return res, ok
}

// SetTo stores value for selected field in Secret
func (fpmv *Secret_FieldPathMapValue) SetTo(target **Secret) {
	if *target == nil {
		*target = new(Secret)
	}
	switch fpmv.selector {
	case Secret_FieldPathSelectorData:
		if (*target).Data == nil {
			(*target).Data = make(map[string]string)
		}
		(*target).Data[fpmv.key] = fpmv.value.(string)
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fpmv.selector))
	}
}

func (fpmv *Secret_FieldPathMapValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Secret)
	fpmv.SetTo(&typedObject)
}

// CompareWith compares value in the 'Secret_FieldPathMapValue' with the value under path in 'Secret'.
func (fpmv *Secret_FieldPathMapValue) CompareWith(source *Secret) (int, bool) {
	switch fpmv.selector {
	case Secret_FieldPathSelectorData:
		leftValue := fpmv.value.(string)
		rightValue := source.GetData()[fpmv.key]
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fpmv.selector))
	}
}

func (fpmv *Secret_FieldPathMapValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpmv.CompareWith(source.(*Secret))
}

type Secret_FieldSubPathValue struct {
	Secret_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ Secret_FieldPathValue = (*Secret_FieldSubPathValue)(nil)

func (fpvs *Secret_FieldSubPathValue) AsMetadataPathValue() (meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(meta.Meta_FieldPathValue)
	return res, ok
}

func (fpvs *Secret_FieldSubPathValue) SetTo(target **Secret) {
	if *target == nil {
		*target = new(Secret)
	}
	switch fpvs.Selector() {
	case Secret_FieldPathSelectorMetadata:
		fpvs.subPathValue.(meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fpvs.Selector()))
	}
}

func (fpvs *Secret_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Secret)
	fpvs.SetTo(&typedObject)
}

func (fpvs *Secret_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *Secret_FieldSubPathValue) CompareWith(source *Secret) (int, bool) {
	switch fpvs.Selector() {
	case Secret_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fpvs.Selector()))
	}
}

func (fpvs *Secret_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*Secret))
}

// Secret_FieldPathArrayItemValue allows storing single item in Path-specific values for Secret according to their type
// Present only for array (repeated) types.
type Secret_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	Secret_FieldPath
	ContainsValue(*Secret) bool
}

// ParseSecret_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseSecret_FieldPathArrayItemValue(pathStr, valueStr string) (Secret_FieldPathArrayItemValue, error) {
	fp, err := ParseSecret_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Secret field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(Secret_FieldPathArrayItemValue), nil
}

func MustParseSecret_FieldPathArrayItemValue(pathStr, valueStr string) Secret_FieldPathArrayItemValue {
	fpaiv, err := ParseSecret_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type Secret_FieldTerminalPathArrayItemValue struct {
	Secret_FieldTerminalPath
	value interface{}
}

var _ Secret_FieldPathArrayItemValue = (*Secret_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object Secret as interface{}
func (fpaiv *Secret_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *Secret_FieldTerminalPathArrayItemValue) GetSingle(source *Secret) (interface{}, bool) {
	return nil, false
}

func (fpaiv *Secret_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*Secret))
}

// Contains returns a boolean indicating if value that is being held is present in given 'Secret'
func (fpaiv *Secret_FieldTerminalPathArrayItemValue) ContainsValue(source *Secret) bool {
	slice := fpaiv.Secret_FieldTerminalPath.Get(source)
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

type Secret_FieldSubPathArrayItemValue struct {
	Secret_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *Secret_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *Secret_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'Secret'
func (fpaivs *Secret_FieldSubPathArrayItemValue) ContainsValue(source *Secret) bool {
	switch fpaivs.Selector() {
	case Secret_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Secret: %d", fpaivs.Selector()))
	}
}

// Secret_FieldPathArrayOfValues allows storing slice of values for Secret fields according to their type
type Secret_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	Secret_FieldPath
}

func ParseSecret_FieldPathArrayOfValues(pathStr, valuesStr string) (Secret_FieldPathArrayOfValues, error) {
	fp, err := ParseSecret_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Secret field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(Secret_FieldPathArrayOfValues), nil
}

func MustParseSecret_FieldPathArrayOfValues(pathStr, valuesStr string) Secret_FieldPathArrayOfValues {
	fpaov, err := ParseSecret_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type Secret_FieldTerminalPathArrayOfValues struct {
	Secret_FieldTerminalPath
	values interface{}
}

var _ Secret_FieldPathArrayOfValues = (*Secret_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *Secret_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case Secret_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case Secret_FieldPathSelectorEncData:
		for _, v := range fpaov.values.([][]byte) {
			values = append(values, v)
		}
	case Secret_FieldPathSelectorData:
		for _, v := range fpaov.values.([]map[string]string) {
			values = append(values, v)
		}
	case Secret_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*meta.Meta) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *Secret_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *Secret_FieldTerminalPathArrayOfValues) AsEncDataArrayOfValues() ([][]byte, bool) {
	res, ok := fpaov.values.([][]byte)
	return res, ok
}
func (fpaov *Secret_FieldTerminalPathArrayOfValues) AsDataArrayOfValues() ([]map[string]string, bool) {
	res, ok := fpaov.values.([]map[string]string)
	return res, ok
}
func (fpaov *Secret_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*meta.Meta, bool) {
	res, ok := fpaov.values.([]*meta.Meta)
	return res, ok
}

type Secret_FieldPathMapArrayOfValues struct {
	Secret_FieldPathMap
	values interface{}
}

var _ Secret_FieldPathArrayOfValues = (*Secret_FieldPathMapArrayOfValues)(nil)

func (fpmaov *Secret_FieldPathMapArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpmaov.selector {
	case Secret_FieldPathSelectorData:
		for _, v := range fpmaov.values.([]string) {
			values = append(values, v)
		}
	}
	return
}
func (fpmaov *Secret_FieldPathMapArrayOfValues) AsDataArrayOfElementValues() ([]string, bool) {
	res, ok := fpmaov.values.([]string)
	return res, ok
}

type Secret_FieldSubPathArrayOfValues struct {
	Secret_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ Secret_FieldPathArrayOfValues = (*Secret_FieldSubPathArrayOfValues)(nil)

func (fpsaov *Secret_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *Secret_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
