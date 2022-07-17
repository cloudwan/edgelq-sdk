// Code generated by protoc-gen-goten-object
// File: edgelq/limits/proto/v1alpha2/common.proto
// DO NOT EDIT!!!

package common

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
	meta_resource "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/resource"
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
	_ = &meta_resource.Resource{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type Allowance_FieldPath interface {
	gotenobject.FieldPath
	Selector() Allowance_FieldPathSelector
	Get(source *Allowance) []interface{}
	GetSingle(source *Allowance) (interface{}, bool)
	ClearValue(item *Allowance)

	// Those methods build corresponding Allowance_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) Allowance_FieldPathValue
	WithIArrayOfValues(values interface{}) Allowance_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) Allowance_FieldPathArrayItemValue
}

type Allowance_FieldPathSelector int32

const (
	Allowance_FieldPathSelectorResource Allowance_FieldPathSelector = 0
	Allowance_FieldPathSelectorValue    Allowance_FieldPathSelector = 1
)

func (s Allowance_FieldPathSelector) String() string {
	switch s {
	case Allowance_FieldPathSelectorResource:
		return "resource"
	case Allowance_FieldPathSelectorValue:
		return "value"
	default:
		panic(fmt.Sprintf("Invalid selector for Allowance: %d", s))
	}
}

func BuildAllowance_FieldPath(fp gotenobject.RawFieldPath) (Allowance_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object Allowance")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "resource":
			return &Allowance_FieldTerminalPath{selector: Allowance_FieldPathSelectorResource}, nil
		case "value":
			return &Allowance_FieldTerminalPath{selector: Allowance_FieldPathSelectorValue}, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object Allowance", fp)
}

func ParseAllowance_FieldPath(rawField string) (Allowance_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildAllowance_FieldPath(fp)
}

func MustParseAllowance_FieldPath(rawField string) Allowance_FieldPath {
	fp, err := ParseAllowance_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type Allowance_FieldTerminalPath struct {
	selector Allowance_FieldPathSelector
}

var _ Allowance_FieldPath = (*Allowance_FieldTerminalPath)(nil)

func (fp *Allowance_FieldTerminalPath) Selector() Allowance_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *Allowance_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *Allowance_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source Allowance
func (fp *Allowance_FieldTerminalPath) Get(source *Allowance) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case Allowance_FieldPathSelectorResource:
			if source.Resource != nil {
				values = append(values, source.Resource)
			}
		case Allowance_FieldPathSelectorValue:
			values = append(values, source.Value)
		default:
			panic(fmt.Sprintf("Invalid selector for Allowance: %d", fp.selector))
		}
	}
	return
}

func (fp *Allowance_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*Allowance))
}

// GetSingle returns value pointed by specific field of from source Allowance
func (fp *Allowance_FieldTerminalPath) GetSingle(source *Allowance) (interface{}, bool) {
	switch fp.selector {
	case Allowance_FieldPathSelectorResource:
		res := source.GetResource()
		return res, res != nil
	case Allowance_FieldPathSelectorValue:
		return source.GetValue(), source != nil
	default:
		panic(fmt.Sprintf("Invalid selector for Allowance: %d", fp.selector))
	}
}

func (fp *Allowance_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*Allowance))
}

// GetDefault returns a default value of the field type
func (fp *Allowance_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case Allowance_FieldPathSelectorResource:
		return (*meta_resource.Reference)(nil)
	case Allowance_FieldPathSelectorValue:
		return int64(0)
	default:
		panic(fmt.Sprintf("Invalid selector for Allowance: %d", fp.selector))
	}
}

func (fp *Allowance_FieldTerminalPath) ClearValue(item *Allowance) {
	if item != nil {
		switch fp.selector {
		case Allowance_FieldPathSelectorResource:
			item.Resource = nil
		case Allowance_FieldPathSelectorValue:
			item.Value = int64(0)
		default:
			panic(fmt.Sprintf("Invalid selector for Allowance: %d", fp.selector))
		}
	}
}

func (fp *Allowance_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*Allowance))
}

// IsLeaf - whether field path is holds simple value
func (fp *Allowance_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == Allowance_FieldPathSelectorResource ||
		fp.selector == Allowance_FieldPathSelectorValue
}

func (fp *Allowance_FieldTerminalPath) WithIValue(value interface{}) Allowance_FieldPathValue {
	switch fp.selector {
	case Allowance_FieldPathSelectorResource:
		return &Allowance_FieldTerminalPathValue{Allowance_FieldTerminalPath: *fp, value: value.(*meta_resource.Reference)}
	case Allowance_FieldPathSelectorValue:
		return &Allowance_FieldTerminalPathValue{Allowance_FieldTerminalPath: *fp, value: value.(int64)}
	default:
		panic(fmt.Sprintf("Invalid selector for Allowance: %d", fp.selector))
	}
}

func (fp *Allowance_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *Allowance_FieldTerminalPath) WithIArrayOfValues(values interface{}) Allowance_FieldPathArrayOfValues {
	fpaov := &Allowance_FieldTerminalPathArrayOfValues{Allowance_FieldTerminalPath: *fp}
	switch fp.selector {
	case Allowance_FieldPathSelectorResource:
		return &Allowance_FieldTerminalPathArrayOfValues{Allowance_FieldTerminalPath: *fp, values: values.([]*meta_resource.Reference)}
	case Allowance_FieldPathSelectorValue:
		return &Allowance_FieldTerminalPathArrayOfValues{Allowance_FieldTerminalPath: *fp, values: values.([]int64)}
	default:
		panic(fmt.Sprintf("Invalid selector for Allowance: %d", fp.selector))
	}
	return fpaov
}

func (fp *Allowance_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *Allowance_FieldTerminalPath) WithIArrayItemValue(value interface{}) Allowance_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for Allowance: %d", fp.selector))
	}
}

func (fp *Allowance_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

// Allowance_FieldPathValue allows storing values for Allowance fields according to their type
type Allowance_FieldPathValue interface {
	Allowance_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **Allowance)
	CompareWith(*Allowance) (cmp int, comparable bool)
}

func ParseAllowance_FieldPathValue(pathStr, valueStr string) (Allowance_FieldPathValue, error) {
	fp, err := ParseAllowance_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Allowance field path value from %s: %v", valueStr, err)
	}
	return fpv.(Allowance_FieldPathValue), nil
}

func MustParseAllowance_FieldPathValue(pathStr, valueStr string) Allowance_FieldPathValue {
	fpv, err := ParseAllowance_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type Allowance_FieldTerminalPathValue struct {
	Allowance_FieldTerminalPath
	value interface{}
}

var _ Allowance_FieldPathValue = (*Allowance_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'Allowance' as interface{}
func (fpv *Allowance_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *Allowance_FieldTerminalPathValue) AsResourceValue() (*meta_resource.Reference, bool) {
	res, ok := fpv.value.(*meta_resource.Reference)
	return res, ok
}
func (fpv *Allowance_FieldTerminalPathValue) AsValueValue() (int64, bool) {
	res, ok := fpv.value.(int64)
	return res, ok
}

// SetTo stores value for selected field for object Allowance
func (fpv *Allowance_FieldTerminalPathValue) SetTo(target **Allowance) {
	if *target == nil {
		*target = new(Allowance)
	}
	switch fpv.selector {
	case Allowance_FieldPathSelectorResource:
		(*target).Resource = fpv.value.(*meta_resource.Reference)
	case Allowance_FieldPathSelectorValue:
		(*target).Value = fpv.value.(int64)
	default:
		panic(fmt.Sprintf("Invalid selector for Allowance: %d", fpv.selector))
	}
}

func (fpv *Allowance_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Allowance)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'Allowance_FieldTerminalPathValue' with the value under path in 'Allowance'.
func (fpv *Allowance_FieldTerminalPathValue) CompareWith(source *Allowance) (int, bool) {
	switch fpv.selector {
	case Allowance_FieldPathSelectorResource:
		leftValue := fpv.value.(*meta_resource.Reference)
		rightValue := source.GetResource()
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
	case Allowance_FieldPathSelectorValue:
		leftValue := fpv.value.(int64)
		rightValue := source.GetValue()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	default:
		panic(fmt.Sprintf("Invalid selector for Allowance: %d", fpv.selector))
	}
}

func (fpv *Allowance_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*Allowance))
}

// Allowance_FieldPathArrayItemValue allows storing single item in Path-specific values for Allowance according to their type
// Present only for array (repeated) types.
type Allowance_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	Allowance_FieldPath
	ContainsValue(*Allowance) bool
}

// ParseAllowance_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseAllowance_FieldPathArrayItemValue(pathStr, valueStr string) (Allowance_FieldPathArrayItemValue, error) {
	fp, err := ParseAllowance_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Allowance field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(Allowance_FieldPathArrayItemValue), nil
}

func MustParseAllowance_FieldPathArrayItemValue(pathStr, valueStr string) Allowance_FieldPathArrayItemValue {
	fpaiv, err := ParseAllowance_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type Allowance_FieldTerminalPathArrayItemValue struct {
	Allowance_FieldTerminalPath
	value interface{}
}

var _ Allowance_FieldPathArrayItemValue = (*Allowance_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object Allowance as interface{}
func (fpaiv *Allowance_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *Allowance_FieldTerminalPathArrayItemValue) GetSingle(source *Allowance) (interface{}, bool) {
	return nil, false
}

func (fpaiv *Allowance_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*Allowance))
}

// Contains returns a boolean indicating if value that is being held is present in given 'Allowance'
func (fpaiv *Allowance_FieldTerminalPathArrayItemValue) ContainsValue(source *Allowance) bool {
	slice := fpaiv.Allowance_FieldTerminalPath.Get(source)
	for _, v := range slice {
		if reflect.DeepEqual(v, fpaiv.value) {
			return true
		}
	}
	return false
}

// Allowance_FieldPathArrayOfValues allows storing slice of values for Allowance fields according to their type
type Allowance_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	Allowance_FieldPath
}

func ParseAllowance_FieldPathArrayOfValues(pathStr, valuesStr string) (Allowance_FieldPathArrayOfValues, error) {
	fp, err := ParseAllowance_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Allowance field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(Allowance_FieldPathArrayOfValues), nil
}

func MustParseAllowance_FieldPathArrayOfValues(pathStr, valuesStr string) Allowance_FieldPathArrayOfValues {
	fpaov, err := ParseAllowance_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type Allowance_FieldTerminalPathArrayOfValues struct {
	Allowance_FieldTerminalPath
	values interface{}
}

var _ Allowance_FieldPathArrayOfValues = (*Allowance_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *Allowance_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case Allowance_FieldPathSelectorResource:
		for _, v := range fpaov.values.([]*meta_resource.Reference) {
			values = append(values, v)
		}
	case Allowance_FieldPathSelectorValue:
		for _, v := range fpaov.values.([]int64) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *Allowance_FieldTerminalPathArrayOfValues) AsResourceArrayOfValues() ([]*meta_resource.Reference, bool) {
	res, ok := fpaov.values.([]*meta_resource.Reference)
	return res, ok
}
func (fpaov *Allowance_FieldTerminalPathArrayOfValues) AsValueArrayOfValues() ([]int64, bool) {
	res, ok := fpaov.values.([]int64)
	return res, ok
}

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type RegionalDistribution_FieldPath interface {
	gotenobject.FieldPath
	Selector() RegionalDistribution_FieldPathSelector
	Get(source *RegionalDistribution) []interface{}
	GetSingle(source *RegionalDistribution) (interface{}, bool)
	ClearValue(item *RegionalDistribution)

	// Those methods build corresponding RegionalDistribution_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) RegionalDistribution_FieldPathValue
	WithIArrayOfValues(values interface{}) RegionalDistribution_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) RegionalDistribution_FieldPathArrayItemValue
}

type RegionalDistribution_FieldPathSelector int32

const (
	RegionalDistribution_FieldPathSelectorResource       RegionalDistribution_FieldPathSelector = 0
	RegionalDistribution_FieldPathSelectorLimitsByRegion RegionalDistribution_FieldPathSelector = 1
)

func (s RegionalDistribution_FieldPathSelector) String() string {
	switch s {
	case RegionalDistribution_FieldPathSelectorResource:
		return "resource"
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		return "limits_by_region"
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", s))
	}
}

func BuildRegionalDistribution_FieldPath(fp gotenobject.RawFieldPath) (RegionalDistribution_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object RegionalDistribution")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "resource":
			return &RegionalDistribution_FieldTerminalPath{selector: RegionalDistribution_FieldPathSelectorResource}, nil
		case "limits_by_region", "limitsByRegion", "limits-by-region":
			return &RegionalDistribution_FieldTerminalPath{selector: RegionalDistribution_FieldPathSelectorLimitsByRegion}, nil
		}
	} else {
		switch fp[0] {
		case "limits_by_region", "limitsByRegion", "limits-by-region":
			if len(fp) > 2 {
				return nil, status.Errorf(codes.InvalidArgument, "sub path for maps ('%s') are not supported (object RegionalDistribution)", fp)
			}
			return &RegionalDistribution_FieldPathMap{selector: RegionalDistribution_FieldPathSelectorLimitsByRegion, key: fp[1]}, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object RegionalDistribution", fp)
}

func ParseRegionalDistribution_FieldPath(rawField string) (RegionalDistribution_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildRegionalDistribution_FieldPath(fp)
}

func MustParseRegionalDistribution_FieldPath(rawField string) RegionalDistribution_FieldPath {
	fp, err := ParseRegionalDistribution_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type RegionalDistribution_FieldTerminalPath struct {
	selector RegionalDistribution_FieldPathSelector
}

var _ RegionalDistribution_FieldPath = (*RegionalDistribution_FieldTerminalPath)(nil)

func (fp *RegionalDistribution_FieldTerminalPath) Selector() RegionalDistribution_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *RegionalDistribution_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *RegionalDistribution_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source RegionalDistribution
func (fp *RegionalDistribution_FieldTerminalPath) Get(source *RegionalDistribution) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case RegionalDistribution_FieldPathSelectorResource:
			if source.Resource != nil {
				values = append(values, source.Resource)
			}
		case RegionalDistribution_FieldPathSelectorLimitsByRegion:
			values = append(values, source.LimitsByRegion)
		default:
			panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fp.selector))
		}
	}
	return
}

func (fp *RegionalDistribution_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*RegionalDistribution))
}

// GetSingle returns value pointed by specific field of from source RegionalDistribution
func (fp *RegionalDistribution_FieldTerminalPath) GetSingle(source *RegionalDistribution) (interface{}, bool) {
	switch fp.selector {
	case RegionalDistribution_FieldPathSelectorResource:
		res := source.GetResource()
		return res, res != nil
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		res := source.GetLimitsByRegion()
		return res, res != nil
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fp.selector))
	}
}

func (fp *RegionalDistribution_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*RegionalDistribution))
}

// GetDefault returns a default value of the field type
func (fp *RegionalDistribution_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case RegionalDistribution_FieldPathSelectorResource:
		return (*meta_resource.Reference)(nil)
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		return (map[string]int64)(nil)
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fp.selector))
	}
}

func (fp *RegionalDistribution_FieldTerminalPath) ClearValue(item *RegionalDistribution) {
	if item != nil {
		switch fp.selector {
		case RegionalDistribution_FieldPathSelectorResource:
			item.Resource = nil
		case RegionalDistribution_FieldPathSelectorLimitsByRegion:
			item.LimitsByRegion = nil
		default:
			panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fp.selector))
		}
	}
}

func (fp *RegionalDistribution_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*RegionalDistribution))
}

// IsLeaf - whether field path is holds simple value
func (fp *RegionalDistribution_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == RegionalDistribution_FieldPathSelectorResource ||
		fp.selector == RegionalDistribution_FieldPathSelectorLimitsByRegion
}

func (fp *RegionalDistribution_FieldTerminalPath) WithIValue(value interface{}) RegionalDistribution_FieldPathValue {
	switch fp.selector {
	case RegionalDistribution_FieldPathSelectorResource:
		return &RegionalDistribution_FieldTerminalPathValue{RegionalDistribution_FieldTerminalPath: *fp, value: value.(*meta_resource.Reference)}
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		return &RegionalDistribution_FieldTerminalPathValue{RegionalDistribution_FieldTerminalPath: *fp, value: value.(map[string]int64)}
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fp.selector))
	}
}

func (fp *RegionalDistribution_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *RegionalDistribution_FieldTerminalPath) WithIArrayOfValues(values interface{}) RegionalDistribution_FieldPathArrayOfValues {
	fpaov := &RegionalDistribution_FieldTerminalPathArrayOfValues{RegionalDistribution_FieldTerminalPath: *fp}
	switch fp.selector {
	case RegionalDistribution_FieldPathSelectorResource:
		return &RegionalDistribution_FieldTerminalPathArrayOfValues{RegionalDistribution_FieldTerminalPath: *fp, values: values.([]*meta_resource.Reference)}
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		return &RegionalDistribution_FieldTerminalPathArrayOfValues{RegionalDistribution_FieldTerminalPath: *fp, values: values.([]map[string]int64)}
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fp.selector))
	}
	return fpaov
}

func (fp *RegionalDistribution_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *RegionalDistribution_FieldTerminalPath) WithIArrayItemValue(value interface{}) RegionalDistribution_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fp.selector))
	}
}

func (fp *RegionalDistribution_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

// FieldPath for map type with additional Key information
type RegionalDistribution_FieldPathMap struct {
	key      string
	selector RegionalDistribution_FieldPathSelector
}

var _ RegionalDistribution_FieldPath = (*RegionalDistribution_FieldPathMap)(nil)

func (fpm *RegionalDistribution_FieldPathMap) Selector() RegionalDistribution_FieldPathSelector {
	return fpm.selector
}

func (fpm *RegionalDistribution_FieldPathMap) Key() string {
	return fpm.key
}

// String returns path representation in proto convention
func (fpm *RegionalDistribution_FieldPathMap) String() string {
	return fpm.selector.String() + "." + fpm.key
}

// JSONString returns path representation is JSON convention. Note that map keys are not transformed
func (fpm *RegionalDistribution_FieldPathMap) JSONString() string {
	return strcase.ToLowerCamel(fpm.selector.String()) + "." + fpm.key
}

// Get returns all values pointed by selected field map key from source RegionalDistribution
func (fpm *RegionalDistribution_FieldPathMap) Get(source *RegionalDistribution) (values []interface{}) {
	switch fpm.selector {
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		if value, ok := source.GetLimitsByRegion()[fpm.key]; ok {
			values = append(values, value)
		}
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fpm.selector))
	}
	return
}

func (fpm *RegionalDistribution_FieldPathMap) GetRaw(source proto.Message) []interface{} {
	return fpm.Get(source.(*RegionalDistribution))
}

// GetSingle returns value by selected field map key from source RegionalDistribution
func (fpm *RegionalDistribution_FieldPathMap) GetSingle(source *RegionalDistribution) (interface{}, bool) {
	switch fpm.selector {
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		res, ok := source.GetLimitsByRegion()[fpm.key]
		return res, ok
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fpm.selector))
	}
}

func (fpm *RegionalDistribution_FieldPathMap) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpm.GetSingle(source.(*RegionalDistribution))
}

// GetDefault returns a default value of the field type
func (fpm *RegionalDistribution_FieldPathMap) GetDefault() interface{} {
	switch fpm.selector {
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		var v int64
		return v
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fpm.selector))
	}
}

func (fpm *RegionalDistribution_FieldPathMap) ClearValue(item *RegionalDistribution) {
	if item != nil {
		switch fpm.selector {
		case RegionalDistribution_FieldPathSelectorLimitsByRegion:
			delete(item.LimitsByRegion, fpm.key)
		default:
			panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fpm.selector))
		}
	}
}

func (fpm *RegionalDistribution_FieldPathMap) ClearValueRaw(item proto.Message) {
	fpm.ClearValue(item.(*RegionalDistribution))
}

// IsLeaf - whether field path is holds simple value
func (fpm *RegionalDistribution_FieldPathMap) IsLeaf() bool {
	switch fpm.selector {
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		return true
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fpm.selector))
	}
}

func (fpm *RegionalDistribution_FieldPathMap) WithIValue(value interface{}) RegionalDistribution_FieldPathValue {
	switch fpm.selector {
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		return &RegionalDistribution_FieldPathMapValue{RegionalDistribution_FieldPathMap: *fpm, value: value.(int64)}
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fpm.selector))
	}
}

func (fpm *RegionalDistribution_FieldPathMap) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fpm.WithIValue(value)
}

func (fpm *RegionalDistribution_FieldPathMap) WithIArrayOfValues(values interface{}) RegionalDistribution_FieldPathArrayOfValues {
	switch fpm.selector {
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		return &RegionalDistribution_FieldPathMapArrayOfValues{RegionalDistribution_FieldPathMap: *fpm, values: values.([]int64)}
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fpm.selector))
	}
}

func (fpm *RegionalDistribution_FieldPathMap) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fpm.WithIArrayOfValues(values)
}

func (fpm *RegionalDistribution_FieldPathMap) WithIArrayItemValue(value interface{}) RegionalDistribution_FieldPathArrayItemValue {
	panic("Cannot create array item value from map fieldpath")
}

func (fpm *RegionalDistribution_FieldPathMap) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fpm.WithIArrayItemValue(value)
}

// RegionalDistribution_FieldPathValue allows storing values for RegionalDistribution fields according to their type
type RegionalDistribution_FieldPathValue interface {
	RegionalDistribution_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **RegionalDistribution)
	CompareWith(*RegionalDistribution) (cmp int, comparable bool)
}

func ParseRegionalDistribution_FieldPathValue(pathStr, valueStr string) (RegionalDistribution_FieldPathValue, error) {
	fp, err := ParseRegionalDistribution_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing RegionalDistribution field path value from %s: %v", valueStr, err)
	}
	return fpv.(RegionalDistribution_FieldPathValue), nil
}

func MustParseRegionalDistribution_FieldPathValue(pathStr, valueStr string) RegionalDistribution_FieldPathValue {
	fpv, err := ParseRegionalDistribution_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type RegionalDistribution_FieldTerminalPathValue struct {
	RegionalDistribution_FieldTerminalPath
	value interface{}
}

var _ RegionalDistribution_FieldPathValue = (*RegionalDistribution_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'RegionalDistribution' as interface{}
func (fpv *RegionalDistribution_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *RegionalDistribution_FieldTerminalPathValue) AsResourceValue() (*meta_resource.Reference, bool) {
	res, ok := fpv.value.(*meta_resource.Reference)
	return res, ok
}
func (fpv *RegionalDistribution_FieldTerminalPathValue) AsLimitsByRegionValue() (map[string]int64, bool) {
	res, ok := fpv.value.(map[string]int64)
	return res, ok
}

// SetTo stores value for selected field for object RegionalDistribution
func (fpv *RegionalDistribution_FieldTerminalPathValue) SetTo(target **RegionalDistribution) {
	if *target == nil {
		*target = new(RegionalDistribution)
	}
	switch fpv.selector {
	case RegionalDistribution_FieldPathSelectorResource:
		(*target).Resource = fpv.value.(*meta_resource.Reference)
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		(*target).LimitsByRegion = fpv.value.(map[string]int64)
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fpv.selector))
	}
}

func (fpv *RegionalDistribution_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*RegionalDistribution)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'RegionalDistribution_FieldTerminalPathValue' with the value under path in 'RegionalDistribution'.
func (fpv *RegionalDistribution_FieldTerminalPathValue) CompareWith(source *RegionalDistribution) (int, bool) {
	switch fpv.selector {
	case RegionalDistribution_FieldPathSelectorResource:
		leftValue := fpv.value.(*meta_resource.Reference)
		rightValue := source.GetResource()
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
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		return 0, false
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fpv.selector))
	}
}

func (fpv *RegionalDistribution_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*RegionalDistribution))
}

type RegionalDistribution_FieldPathMapValue struct {
	RegionalDistribution_FieldPathMap
	value interface{}
}

var _ RegionalDistribution_FieldPathValue = (*RegionalDistribution_FieldPathMapValue)(nil)

// GetValue returns value stored under selected field in RegionalDistribution as interface{}
func (fpmv *RegionalDistribution_FieldPathMapValue) GetRawValue() interface{} {
	return fpmv.value
}
func (fpmv *RegionalDistribution_FieldPathMapValue) AsLimitsByRegionElementValue() (int64, bool) {
	res, ok := fpmv.value.(int64)
	return res, ok
}

// SetTo stores value for selected field in RegionalDistribution
func (fpmv *RegionalDistribution_FieldPathMapValue) SetTo(target **RegionalDistribution) {
	if *target == nil {
		*target = new(RegionalDistribution)
	}
	switch fpmv.selector {
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		if (*target).LimitsByRegion == nil {
			(*target).LimitsByRegion = make(map[string]int64)
		}
		(*target).LimitsByRegion[fpmv.key] = fpmv.value.(int64)
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fpmv.selector))
	}
}

func (fpmv *RegionalDistribution_FieldPathMapValue) SetToRaw(target proto.Message) {
	typedObject := target.(*RegionalDistribution)
	fpmv.SetTo(&typedObject)
}

// CompareWith compares value in the 'RegionalDistribution_FieldPathMapValue' with the value under path in 'RegionalDistribution'.
func (fpmv *RegionalDistribution_FieldPathMapValue) CompareWith(source *RegionalDistribution) (int, bool) {
	switch fpmv.selector {
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		leftValue := fpmv.value.(int64)
		rightValue := source.GetLimitsByRegion()[fpmv.key]
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalDistribution: %d", fpmv.selector))
	}
}

func (fpmv *RegionalDistribution_FieldPathMapValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpmv.CompareWith(source.(*RegionalDistribution))
}

// RegionalDistribution_FieldPathArrayItemValue allows storing single item in Path-specific values for RegionalDistribution according to their type
// Present only for array (repeated) types.
type RegionalDistribution_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	RegionalDistribution_FieldPath
	ContainsValue(*RegionalDistribution) bool
}

// ParseRegionalDistribution_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseRegionalDistribution_FieldPathArrayItemValue(pathStr, valueStr string) (RegionalDistribution_FieldPathArrayItemValue, error) {
	fp, err := ParseRegionalDistribution_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing RegionalDistribution field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(RegionalDistribution_FieldPathArrayItemValue), nil
}

func MustParseRegionalDistribution_FieldPathArrayItemValue(pathStr, valueStr string) RegionalDistribution_FieldPathArrayItemValue {
	fpaiv, err := ParseRegionalDistribution_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type RegionalDistribution_FieldTerminalPathArrayItemValue struct {
	RegionalDistribution_FieldTerminalPath
	value interface{}
}

var _ RegionalDistribution_FieldPathArrayItemValue = (*RegionalDistribution_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object RegionalDistribution as interface{}
func (fpaiv *RegionalDistribution_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *RegionalDistribution_FieldTerminalPathArrayItemValue) GetSingle(source *RegionalDistribution) (interface{}, bool) {
	return nil, false
}

func (fpaiv *RegionalDistribution_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*RegionalDistribution))
}

// Contains returns a boolean indicating if value that is being held is present in given 'RegionalDistribution'
func (fpaiv *RegionalDistribution_FieldTerminalPathArrayItemValue) ContainsValue(source *RegionalDistribution) bool {
	slice := fpaiv.RegionalDistribution_FieldTerminalPath.Get(source)
	for _, v := range slice {
		if reflect.DeepEqual(v, fpaiv.value) {
			return true
		}
	}
	return false
}

// RegionalDistribution_FieldPathArrayOfValues allows storing slice of values for RegionalDistribution fields according to their type
type RegionalDistribution_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	RegionalDistribution_FieldPath
}

func ParseRegionalDistribution_FieldPathArrayOfValues(pathStr, valuesStr string) (RegionalDistribution_FieldPathArrayOfValues, error) {
	fp, err := ParseRegionalDistribution_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing RegionalDistribution field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(RegionalDistribution_FieldPathArrayOfValues), nil
}

func MustParseRegionalDistribution_FieldPathArrayOfValues(pathStr, valuesStr string) RegionalDistribution_FieldPathArrayOfValues {
	fpaov, err := ParseRegionalDistribution_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type RegionalDistribution_FieldTerminalPathArrayOfValues struct {
	RegionalDistribution_FieldTerminalPath
	values interface{}
}

var _ RegionalDistribution_FieldPathArrayOfValues = (*RegionalDistribution_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *RegionalDistribution_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case RegionalDistribution_FieldPathSelectorResource:
		for _, v := range fpaov.values.([]*meta_resource.Reference) {
			values = append(values, v)
		}
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		for _, v := range fpaov.values.([]map[string]int64) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *RegionalDistribution_FieldTerminalPathArrayOfValues) AsResourceArrayOfValues() ([]*meta_resource.Reference, bool) {
	res, ok := fpaov.values.([]*meta_resource.Reference)
	return res, ok
}
func (fpaov *RegionalDistribution_FieldTerminalPathArrayOfValues) AsLimitsByRegionArrayOfValues() ([]map[string]int64, bool) {
	res, ok := fpaov.values.([]map[string]int64)
	return res, ok
}

type RegionalDistribution_FieldPathMapArrayOfValues struct {
	RegionalDistribution_FieldPathMap
	values interface{}
}

var _ RegionalDistribution_FieldPathArrayOfValues = (*RegionalDistribution_FieldPathMapArrayOfValues)(nil)

func (fpmaov *RegionalDistribution_FieldPathMapArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpmaov.selector {
	case RegionalDistribution_FieldPathSelectorLimitsByRegion:
		for _, v := range fpmaov.values.([]int64) {
			values = append(values, v)
		}
	}
	return
}
func (fpmaov *RegionalDistribution_FieldPathMapArrayOfValues) AsLimitsByRegionArrayOfElementValues() ([]int64, bool) {
	res, ok := fpmaov.values.([]int64)
	return res, ok
}
