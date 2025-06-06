// Code generated by protoc-gen-goten-object
// File: edgelq/secrets/proto/v1/crypto_key.proto
// DO NOT EDIT!!!

package crypto_key

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
	project "github.com/cloudwan/edgelq-sdk/secrets/resources/v1/project"
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
type CryptoKey_FieldPath interface {
	gotenobject.FieldPath
	Selector() CryptoKey_FieldPathSelector
	Get(source *CryptoKey) []interface{}
	GetSingle(source *CryptoKey) (interface{}, bool)
	ClearValue(item *CryptoKey)

	// Those methods build corresponding CryptoKey_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) CryptoKey_FieldPathValue
	WithIArrayOfValues(values interface{}) CryptoKey_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) CryptoKey_FieldPathArrayItemValue
}

type CryptoKey_FieldPathSelector int32

const (
	CryptoKey_FieldPathSelectorName     CryptoKey_FieldPathSelector = 0
	CryptoKey_FieldPathSelectorMetadata CryptoKey_FieldPathSelector = 1
	CryptoKey_FieldPathSelectorKey      CryptoKey_FieldPathSelector = 2
)

func (s CryptoKey_FieldPathSelector) String() string {
	switch s {
	case CryptoKey_FieldPathSelectorName:
		return "name"
	case CryptoKey_FieldPathSelectorMetadata:
		return "metadata"
	case CryptoKey_FieldPathSelectorKey:
		return "key"
	default:
		panic(fmt.Sprintf("Invalid selector for CryptoKey: %d", s))
	}
}

func BuildCryptoKey_FieldPath(fp gotenobject.RawFieldPath) (CryptoKey_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object CryptoKey")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &CryptoKey_FieldTerminalPath{selector: CryptoKey_FieldPathSelectorName}, nil
		case "metadata":
			return &CryptoKey_FieldTerminalPath{selector: CryptoKey_FieldPathSelectorMetadata}, nil
		case "key":
			return &CryptoKey_FieldTerminalPath{selector: CryptoKey_FieldPathSelectorKey}, nil
		}
	} else {
		switch fp[0] {
		case "metadata":
			if subpath, err := meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &CryptoKey_FieldSubPath{selector: CryptoKey_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object CryptoKey", fp)
}

func ParseCryptoKey_FieldPath(rawField string) (CryptoKey_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildCryptoKey_FieldPath(fp)
}

func MustParseCryptoKey_FieldPath(rawField string) CryptoKey_FieldPath {
	fp, err := ParseCryptoKey_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type CryptoKey_FieldTerminalPath struct {
	selector CryptoKey_FieldPathSelector
}

var _ CryptoKey_FieldPath = (*CryptoKey_FieldTerminalPath)(nil)

func (fp *CryptoKey_FieldTerminalPath) Selector() CryptoKey_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *CryptoKey_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *CryptoKey_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source CryptoKey
func (fp *CryptoKey_FieldTerminalPath) Get(source *CryptoKey) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case CryptoKey_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case CryptoKey_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		case CryptoKey_FieldPathSelectorKey:
			values = append(values, source.Key)
		default:
			panic(fmt.Sprintf("Invalid selector for CryptoKey: %d", fp.selector))
		}
	}
	return
}

func (fp *CryptoKey_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*CryptoKey))
}

// GetSingle returns value pointed by specific field of from source CryptoKey
func (fp *CryptoKey_FieldTerminalPath) GetSingle(source *CryptoKey) (interface{}, bool) {
	switch fp.selector {
	case CryptoKey_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case CryptoKey_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	case CryptoKey_FieldPathSelectorKey:
		return source.GetKey(), source != nil
	default:
		panic(fmt.Sprintf("Invalid selector for CryptoKey: %d", fp.selector))
	}
}

func (fp *CryptoKey_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*CryptoKey))
}

// GetDefault returns a default value of the field type
func (fp *CryptoKey_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case CryptoKey_FieldPathSelectorName:
		return (*Name)(nil)
	case CryptoKey_FieldPathSelectorMetadata:
		return (*meta.Meta)(nil)
	case CryptoKey_FieldPathSelectorKey:
		return ""
	default:
		panic(fmt.Sprintf("Invalid selector for CryptoKey: %d", fp.selector))
	}
}

func (fp *CryptoKey_FieldTerminalPath) ClearValue(item *CryptoKey) {
	if item != nil {
		switch fp.selector {
		case CryptoKey_FieldPathSelectorName:
			item.Name = nil
		case CryptoKey_FieldPathSelectorMetadata:
			item.Metadata = nil
		case CryptoKey_FieldPathSelectorKey:
			item.Key = ""
		default:
			panic(fmt.Sprintf("Invalid selector for CryptoKey: %d", fp.selector))
		}
	}
}

func (fp *CryptoKey_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*CryptoKey))
}

// IsLeaf - whether field path is holds simple value
func (fp *CryptoKey_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == CryptoKey_FieldPathSelectorName ||
		fp.selector == CryptoKey_FieldPathSelectorKey
}

func (fp *CryptoKey_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *CryptoKey_FieldTerminalPath) WithIValue(value interface{}) CryptoKey_FieldPathValue {
	switch fp.selector {
	case CryptoKey_FieldPathSelectorName:
		return &CryptoKey_FieldTerminalPathValue{CryptoKey_FieldTerminalPath: *fp, value: value.(*Name)}
	case CryptoKey_FieldPathSelectorMetadata:
		return &CryptoKey_FieldTerminalPathValue{CryptoKey_FieldTerminalPath: *fp, value: value.(*meta.Meta)}
	case CryptoKey_FieldPathSelectorKey:
		return &CryptoKey_FieldTerminalPathValue{CryptoKey_FieldTerminalPath: *fp, value: value.(string)}
	default:
		panic(fmt.Sprintf("Invalid selector for CryptoKey: %d", fp.selector))
	}
}

func (fp *CryptoKey_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *CryptoKey_FieldTerminalPath) WithIArrayOfValues(values interface{}) CryptoKey_FieldPathArrayOfValues {
	fpaov := &CryptoKey_FieldTerminalPathArrayOfValues{CryptoKey_FieldTerminalPath: *fp}
	switch fp.selector {
	case CryptoKey_FieldPathSelectorName:
		return &CryptoKey_FieldTerminalPathArrayOfValues{CryptoKey_FieldTerminalPath: *fp, values: values.([]*Name)}
	case CryptoKey_FieldPathSelectorMetadata:
		return &CryptoKey_FieldTerminalPathArrayOfValues{CryptoKey_FieldTerminalPath: *fp, values: values.([]*meta.Meta)}
	case CryptoKey_FieldPathSelectorKey:
		return &CryptoKey_FieldTerminalPathArrayOfValues{CryptoKey_FieldTerminalPath: *fp, values: values.([]string)}
	default:
		panic(fmt.Sprintf("Invalid selector for CryptoKey: %d", fp.selector))
	}
	return fpaov
}

func (fp *CryptoKey_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *CryptoKey_FieldTerminalPath) WithIArrayItemValue(value interface{}) CryptoKey_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for CryptoKey: %d", fp.selector))
	}
}

func (fp *CryptoKey_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

type CryptoKey_FieldSubPath struct {
	selector CryptoKey_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ CryptoKey_FieldPath = (*CryptoKey_FieldSubPath)(nil)

func (fps *CryptoKey_FieldSubPath) Selector() CryptoKey_FieldPathSelector {
	return fps.selector
}
func (fps *CryptoKey_FieldSubPath) AsMetadataSubPath() (meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(meta.Meta_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *CryptoKey_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *CryptoKey_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source CryptoKey
func (fps *CryptoKey_FieldSubPath) Get(source *CryptoKey) (values []interface{}) {
	switch fps.selector {
	case CryptoKey_FieldPathSelectorMetadata:
		values = append(values, fps.subPath.GetRaw(source.GetMetadata())...)
	default:
		panic(fmt.Sprintf("Invalid selector for CryptoKey: %d", fps.selector))
	}
	return
}

func (fps *CryptoKey_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*CryptoKey))
}

// GetSingle returns value of selected field from source CryptoKey
func (fps *CryptoKey_FieldSubPath) GetSingle(source *CryptoKey) (interface{}, bool) {
	switch fps.selector {
	case CryptoKey_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for CryptoKey: %d", fps.selector))
	}
}

func (fps *CryptoKey_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*CryptoKey))
}

// GetDefault returns a default value of the field type
func (fps *CryptoKey_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *CryptoKey_FieldSubPath) ClearValue(item *CryptoKey) {
	if item != nil {
		switch fps.selector {
		case CryptoKey_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		default:
			panic(fmt.Sprintf("Invalid selector for CryptoKey: %d", fps.selector))
		}
	}
}

func (fps *CryptoKey_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*CryptoKey))
}

// IsLeaf - whether field path is holds simple value
func (fps *CryptoKey_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *CryptoKey_FieldSubPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	iPaths := []gotenobject.FieldPath{&CryptoKey_FieldTerminalPath{selector: fps.selector}}
	iPaths = append(iPaths, fps.subPath.SplitIntoTerminalIPaths()...)
	return iPaths
}

func (fps *CryptoKey_FieldSubPath) WithIValue(value interface{}) CryptoKey_FieldPathValue {
	return &CryptoKey_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *CryptoKey_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *CryptoKey_FieldSubPath) WithIArrayOfValues(values interface{}) CryptoKey_FieldPathArrayOfValues {
	return &CryptoKey_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *CryptoKey_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *CryptoKey_FieldSubPath) WithIArrayItemValue(value interface{}) CryptoKey_FieldPathArrayItemValue {
	return &CryptoKey_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *CryptoKey_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// CryptoKey_FieldPathValue allows storing values for CryptoKey fields according to their type
type CryptoKey_FieldPathValue interface {
	CryptoKey_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **CryptoKey)
	CompareWith(*CryptoKey) (cmp int, comparable bool)
}

func ParseCryptoKey_FieldPathValue(pathStr, valueStr string) (CryptoKey_FieldPathValue, error) {
	fp, err := ParseCryptoKey_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing CryptoKey field path value from %s: %v", valueStr, err)
	}
	return fpv.(CryptoKey_FieldPathValue), nil
}

func MustParseCryptoKey_FieldPathValue(pathStr, valueStr string) CryptoKey_FieldPathValue {
	fpv, err := ParseCryptoKey_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type CryptoKey_FieldTerminalPathValue struct {
	CryptoKey_FieldTerminalPath
	value interface{}
}

var _ CryptoKey_FieldPathValue = (*CryptoKey_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'CryptoKey' as interface{}
func (fpv *CryptoKey_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *CryptoKey_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *CryptoKey_FieldTerminalPathValue) AsMetadataValue() (*meta.Meta, bool) {
	res, ok := fpv.value.(*meta.Meta)
	return res, ok
}
func (fpv *CryptoKey_FieldTerminalPathValue) AsKeyValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}

// SetTo stores value for selected field for object CryptoKey
func (fpv *CryptoKey_FieldTerminalPathValue) SetTo(target **CryptoKey) {
	if *target == nil {
		*target = new(CryptoKey)
	}
	switch fpv.selector {
	case CryptoKey_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case CryptoKey_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*meta.Meta)
	case CryptoKey_FieldPathSelectorKey:
		(*target).Key = fpv.value.(string)
	default:
		panic(fmt.Sprintf("Invalid selector for CryptoKey: %d", fpv.selector))
	}
}

func (fpv *CryptoKey_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*CryptoKey)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'CryptoKey_FieldTerminalPathValue' with the value under path in 'CryptoKey'.
func (fpv *CryptoKey_FieldTerminalPathValue) CompareWith(source *CryptoKey) (int, bool) {
	switch fpv.selector {
	case CryptoKey_FieldPathSelectorName:
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
	case CryptoKey_FieldPathSelectorMetadata:
		return 0, false
	case CryptoKey_FieldPathSelectorKey:
		leftValue := fpv.value.(string)
		rightValue := source.GetKey()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	default:
		panic(fmt.Sprintf("Invalid selector for CryptoKey: %d", fpv.selector))
	}
}

func (fpv *CryptoKey_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*CryptoKey))
}

type CryptoKey_FieldSubPathValue struct {
	CryptoKey_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ CryptoKey_FieldPathValue = (*CryptoKey_FieldSubPathValue)(nil)

func (fpvs *CryptoKey_FieldSubPathValue) AsMetadataPathValue() (meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(meta.Meta_FieldPathValue)
	return res, ok
}

func (fpvs *CryptoKey_FieldSubPathValue) SetTo(target **CryptoKey) {
	if *target == nil {
		*target = new(CryptoKey)
	}
	switch fpvs.Selector() {
	case CryptoKey_FieldPathSelectorMetadata:
		fpvs.subPathValue.(meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	default:
		panic(fmt.Sprintf("Invalid selector for CryptoKey: %d", fpvs.Selector()))
	}
}

func (fpvs *CryptoKey_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*CryptoKey)
	fpvs.SetTo(&typedObject)
}

func (fpvs *CryptoKey_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *CryptoKey_FieldSubPathValue) CompareWith(source *CryptoKey) (int, bool) {
	switch fpvs.Selector() {
	case CryptoKey_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for CryptoKey: %d", fpvs.Selector()))
	}
}

func (fpvs *CryptoKey_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*CryptoKey))
}

// CryptoKey_FieldPathArrayItemValue allows storing single item in Path-specific values for CryptoKey according to their type
// Present only for array (repeated) types.
type CryptoKey_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	CryptoKey_FieldPath
	ContainsValue(*CryptoKey) bool
}

// ParseCryptoKey_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseCryptoKey_FieldPathArrayItemValue(pathStr, valueStr string) (CryptoKey_FieldPathArrayItemValue, error) {
	fp, err := ParseCryptoKey_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing CryptoKey field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(CryptoKey_FieldPathArrayItemValue), nil
}

func MustParseCryptoKey_FieldPathArrayItemValue(pathStr, valueStr string) CryptoKey_FieldPathArrayItemValue {
	fpaiv, err := ParseCryptoKey_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type CryptoKey_FieldTerminalPathArrayItemValue struct {
	CryptoKey_FieldTerminalPath
	value interface{}
}

var _ CryptoKey_FieldPathArrayItemValue = (*CryptoKey_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object CryptoKey as interface{}
func (fpaiv *CryptoKey_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *CryptoKey_FieldTerminalPathArrayItemValue) GetSingle(source *CryptoKey) (interface{}, bool) {
	return nil, false
}

func (fpaiv *CryptoKey_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*CryptoKey))
}

// Contains returns a boolean indicating if value that is being held is present in given 'CryptoKey'
func (fpaiv *CryptoKey_FieldTerminalPathArrayItemValue) ContainsValue(source *CryptoKey) bool {
	slice := fpaiv.CryptoKey_FieldTerminalPath.Get(source)
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

type CryptoKey_FieldSubPathArrayItemValue struct {
	CryptoKey_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *CryptoKey_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *CryptoKey_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'CryptoKey'
func (fpaivs *CryptoKey_FieldSubPathArrayItemValue) ContainsValue(source *CryptoKey) bool {
	switch fpaivs.Selector() {
	case CryptoKey_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for CryptoKey: %d", fpaivs.Selector()))
	}
}

// CryptoKey_FieldPathArrayOfValues allows storing slice of values for CryptoKey fields according to their type
type CryptoKey_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	CryptoKey_FieldPath
}

func ParseCryptoKey_FieldPathArrayOfValues(pathStr, valuesStr string) (CryptoKey_FieldPathArrayOfValues, error) {
	fp, err := ParseCryptoKey_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing CryptoKey field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(CryptoKey_FieldPathArrayOfValues), nil
}

func MustParseCryptoKey_FieldPathArrayOfValues(pathStr, valuesStr string) CryptoKey_FieldPathArrayOfValues {
	fpaov, err := ParseCryptoKey_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type CryptoKey_FieldTerminalPathArrayOfValues struct {
	CryptoKey_FieldTerminalPath
	values interface{}
}

var _ CryptoKey_FieldPathArrayOfValues = (*CryptoKey_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *CryptoKey_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case CryptoKey_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case CryptoKey_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*meta.Meta) {
			values = append(values, v)
		}
	case CryptoKey_FieldPathSelectorKey:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *CryptoKey_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *CryptoKey_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*meta.Meta, bool) {
	res, ok := fpaov.values.([]*meta.Meta)
	return res, ok
}
func (fpaov *CryptoKey_FieldTerminalPathArrayOfValues) AsKeyArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}

type CryptoKey_FieldSubPathArrayOfValues struct {
	CryptoKey_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ CryptoKey_FieldPathArrayOfValues = (*CryptoKey_FieldSubPathArrayOfValues)(nil)

func (fpsaov *CryptoKey_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *CryptoKey_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
