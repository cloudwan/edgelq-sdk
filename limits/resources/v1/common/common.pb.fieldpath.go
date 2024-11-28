// Code generated by protoc-gen-goten-object
// File: edgelq/limits/proto/v1/common.proto
// DO NOT EDIT!!!

package common

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
	plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1/plan"
	meta_resource "github.com/cloudwan/goten-sdk/meta-service/resources/v1/resource"
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
	_ = &plan.Plan{}
	_ = &meta_resource.Resource{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type RegionalPlanAssignment_FieldPath interface {
	gotenobject.FieldPath
	Selector() RegionalPlanAssignment_FieldPathSelector
	Get(source *RegionalPlanAssignment) []interface{}
	GetSingle(source *RegionalPlanAssignment) (interface{}, bool)
	ClearValue(item *RegionalPlanAssignment)

	// Those methods build corresponding RegionalPlanAssignment_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) RegionalPlanAssignment_FieldPathValue
	WithIArrayOfValues(values interface{}) RegionalPlanAssignment_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) RegionalPlanAssignment_FieldPathArrayItemValue
}

type RegionalPlanAssignment_FieldPathSelector int32

const (
	RegionalPlanAssignment_FieldPathSelectorPlan           RegionalPlanAssignment_FieldPathSelector = 0
	RegionalPlanAssignment_FieldPathSelectorRegion         RegionalPlanAssignment_FieldPathSelector = 1
	RegionalPlanAssignment_FieldPathSelectorPlanGeneration RegionalPlanAssignment_FieldPathSelector = 2
)

func (s RegionalPlanAssignment_FieldPathSelector) String() string {
	switch s {
	case RegionalPlanAssignment_FieldPathSelectorPlan:
		return "plan"
	case RegionalPlanAssignment_FieldPathSelectorRegion:
		return "region"
	case RegionalPlanAssignment_FieldPathSelectorPlanGeneration:
		return "plan_generation"
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalPlanAssignment: %d", s))
	}
}

func BuildRegionalPlanAssignment_FieldPath(fp gotenobject.RawFieldPath) (RegionalPlanAssignment_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object RegionalPlanAssignment")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "plan":
			return &RegionalPlanAssignment_FieldTerminalPath{selector: RegionalPlanAssignment_FieldPathSelectorPlan}, nil
		case "region":
			return &RegionalPlanAssignment_FieldTerminalPath{selector: RegionalPlanAssignment_FieldPathSelectorRegion}, nil
		case "plan_generation", "planGeneration", "plan-generation":
			return &RegionalPlanAssignment_FieldTerminalPath{selector: RegionalPlanAssignment_FieldPathSelectorPlanGeneration}, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object RegionalPlanAssignment", fp)
}

func ParseRegionalPlanAssignment_FieldPath(rawField string) (RegionalPlanAssignment_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildRegionalPlanAssignment_FieldPath(fp)
}

func MustParseRegionalPlanAssignment_FieldPath(rawField string) RegionalPlanAssignment_FieldPath {
	fp, err := ParseRegionalPlanAssignment_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type RegionalPlanAssignment_FieldTerminalPath struct {
	selector RegionalPlanAssignment_FieldPathSelector
}

var _ RegionalPlanAssignment_FieldPath = (*RegionalPlanAssignment_FieldTerminalPath)(nil)

func (fp *RegionalPlanAssignment_FieldTerminalPath) Selector() RegionalPlanAssignment_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *RegionalPlanAssignment_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *RegionalPlanAssignment_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source RegionalPlanAssignment
func (fp *RegionalPlanAssignment_FieldTerminalPath) Get(source *RegionalPlanAssignment) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case RegionalPlanAssignment_FieldPathSelectorPlan:
			if source.Plan != nil {
				values = append(values, source.Plan)
			}
		case RegionalPlanAssignment_FieldPathSelectorRegion:
			values = append(values, source.Region)
		case RegionalPlanAssignment_FieldPathSelectorPlanGeneration:
			values = append(values, source.PlanGeneration)
		default:
			panic(fmt.Sprintf("Invalid selector for RegionalPlanAssignment: %d", fp.selector))
		}
	}
	return
}

func (fp *RegionalPlanAssignment_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*RegionalPlanAssignment))
}

// GetSingle returns value pointed by specific field of from source RegionalPlanAssignment
func (fp *RegionalPlanAssignment_FieldTerminalPath) GetSingle(source *RegionalPlanAssignment) (interface{}, bool) {
	switch fp.selector {
	case RegionalPlanAssignment_FieldPathSelectorPlan:
		res := source.GetPlan()
		return res, res != nil
	case RegionalPlanAssignment_FieldPathSelectorRegion:
		return source.GetRegion(), source != nil
	case RegionalPlanAssignment_FieldPathSelectorPlanGeneration:
		return source.GetPlanGeneration(), source != nil
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalPlanAssignment: %d", fp.selector))
	}
}

func (fp *RegionalPlanAssignment_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*RegionalPlanAssignment))
}

// GetDefault returns a default value of the field type
func (fp *RegionalPlanAssignment_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case RegionalPlanAssignment_FieldPathSelectorPlan:
		return (*plan.Reference)(nil)
	case RegionalPlanAssignment_FieldPathSelectorRegion:
		return ""
	case RegionalPlanAssignment_FieldPathSelectorPlanGeneration:
		return int64(0)
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalPlanAssignment: %d", fp.selector))
	}
}

func (fp *RegionalPlanAssignment_FieldTerminalPath) ClearValue(item *RegionalPlanAssignment) {
	if item != nil {
		switch fp.selector {
		case RegionalPlanAssignment_FieldPathSelectorPlan:
			item.Plan = nil
		case RegionalPlanAssignment_FieldPathSelectorRegion:
			item.Region = ""
		case RegionalPlanAssignment_FieldPathSelectorPlanGeneration:
			item.PlanGeneration = int64(0)
		default:
			panic(fmt.Sprintf("Invalid selector for RegionalPlanAssignment: %d", fp.selector))
		}
	}
}

func (fp *RegionalPlanAssignment_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*RegionalPlanAssignment))
}

// IsLeaf - whether field path is holds simple value
func (fp *RegionalPlanAssignment_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == RegionalPlanAssignment_FieldPathSelectorPlan ||
		fp.selector == RegionalPlanAssignment_FieldPathSelectorRegion ||
		fp.selector == RegionalPlanAssignment_FieldPathSelectorPlanGeneration
}

func (fp *RegionalPlanAssignment_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *RegionalPlanAssignment_FieldTerminalPath) WithIValue(value interface{}) RegionalPlanAssignment_FieldPathValue {
	switch fp.selector {
	case RegionalPlanAssignment_FieldPathSelectorPlan:
		return &RegionalPlanAssignment_FieldTerminalPathValue{RegionalPlanAssignment_FieldTerminalPath: *fp, value: value.(*plan.Reference)}
	case RegionalPlanAssignment_FieldPathSelectorRegion:
		return &RegionalPlanAssignment_FieldTerminalPathValue{RegionalPlanAssignment_FieldTerminalPath: *fp, value: value.(string)}
	case RegionalPlanAssignment_FieldPathSelectorPlanGeneration:
		return &RegionalPlanAssignment_FieldTerminalPathValue{RegionalPlanAssignment_FieldTerminalPath: *fp, value: value.(int64)}
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalPlanAssignment: %d", fp.selector))
	}
}

func (fp *RegionalPlanAssignment_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *RegionalPlanAssignment_FieldTerminalPath) WithIArrayOfValues(values interface{}) RegionalPlanAssignment_FieldPathArrayOfValues {
	fpaov := &RegionalPlanAssignment_FieldTerminalPathArrayOfValues{RegionalPlanAssignment_FieldTerminalPath: *fp}
	switch fp.selector {
	case RegionalPlanAssignment_FieldPathSelectorPlan:
		return &RegionalPlanAssignment_FieldTerminalPathArrayOfValues{RegionalPlanAssignment_FieldTerminalPath: *fp, values: values.([]*plan.Reference)}
	case RegionalPlanAssignment_FieldPathSelectorRegion:
		return &RegionalPlanAssignment_FieldTerminalPathArrayOfValues{RegionalPlanAssignment_FieldTerminalPath: *fp, values: values.([]string)}
	case RegionalPlanAssignment_FieldPathSelectorPlanGeneration:
		return &RegionalPlanAssignment_FieldTerminalPathArrayOfValues{RegionalPlanAssignment_FieldTerminalPath: *fp, values: values.([]int64)}
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalPlanAssignment: %d", fp.selector))
	}
	return fpaov
}

func (fp *RegionalPlanAssignment_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *RegionalPlanAssignment_FieldTerminalPath) WithIArrayItemValue(value interface{}) RegionalPlanAssignment_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalPlanAssignment: %d", fp.selector))
	}
}

func (fp *RegionalPlanAssignment_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

// RegionalPlanAssignment_FieldPathValue allows storing values for RegionalPlanAssignment fields according to their type
type RegionalPlanAssignment_FieldPathValue interface {
	RegionalPlanAssignment_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **RegionalPlanAssignment)
	CompareWith(*RegionalPlanAssignment) (cmp int, comparable bool)
}

func ParseRegionalPlanAssignment_FieldPathValue(pathStr, valueStr string) (RegionalPlanAssignment_FieldPathValue, error) {
	fp, err := ParseRegionalPlanAssignment_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing RegionalPlanAssignment field path value from %s: %v", valueStr, err)
	}
	return fpv.(RegionalPlanAssignment_FieldPathValue), nil
}

func MustParseRegionalPlanAssignment_FieldPathValue(pathStr, valueStr string) RegionalPlanAssignment_FieldPathValue {
	fpv, err := ParseRegionalPlanAssignment_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type RegionalPlanAssignment_FieldTerminalPathValue struct {
	RegionalPlanAssignment_FieldTerminalPath
	value interface{}
}

var _ RegionalPlanAssignment_FieldPathValue = (*RegionalPlanAssignment_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'RegionalPlanAssignment' as interface{}
func (fpv *RegionalPlanAssignment_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *RegionalPlanAssignment_FieldTerminalPathValue) AsPlanValue() (*plan.Reference, bool) {
	res, ok := fpv.value.(*plan.Reference)
	return res, ok
}
func (fpv *RegionalPlanAssignment_FieldTerminalPathValue) AsRegionValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *RegionalPlanAssignment_FieldTerminalPathValue) AsPlanGenerationValue() (int64, bool) {
	res, ok := fpv.value.(int64)
	return res, ok
}

// SetTo stores value for selected field for object RegionalPlanAssignment
func (fpv *RegionalPlanAssignment_FieldTerminalPathValue) SetTo(target **RegionalPlanAssignment) {
	if *target == nil {
		*target = new(RegionalPlanAssignment)
	}
	switch fpv.selector {
	case RegionalPlanAssignment_FieldPathSelectorPlan:
		(*target).Plan = fpv.value.(*plan.Reference)
	case RegionalPlanAssignment_FieldPathSelectorRegion:
		(*target).Region = fpv.value.(string)
	case RegionalPlanAssignment_FieldPathSelectorPlanGeneration:
		(*target).PlanGeneration = fpv.value.(int64)
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalPlanAssignment: %d", fpv.selector))
	}
}

func (fpv *RegionalPlanAssignment_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*RegionalPlanAssignment)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'RegionalPlanAssignment_FieldTerminalPathValue' with the value under path in 'RegionalPlanAssignment'.
func (fpv *RegionalPlanAssignment_FieldTerminalPathValue) CompareWith(source *RegionalPlanAssignment) (int, bool) {
	switch fpv.selector {
	case RegionalPlanAssignment_FieldPathSelectorPlan:
		leftValue := fpv.value.(*plan.Reference)
		rightValue := source.GetPlan()
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
	case RegionalPlanAssignment_FieldPathSelectorRegion:
		leftValue := fpv.value.(string)
		rightValue := source.GetRegion()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case RegionalPlanAssignment_FieldPathSelectorPlanGeneration:
		leftValue := fpv.value.(int64)
		rightValue := source.GetPlanGeneration()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	default:
		panic(fmt.Sprintf("Invalid selector for RegionalPlanAssignment: %d", fpv.selector))
	}
}

func (fpv *RegionalPlanAssignment_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*RegionalPlanAssignment))
}

// RegionalPlanAssignment_FieldPathArrayItemValue allows storing single item in Path-specific values for RegionalPlanAssignment according to their type
// Present only for array (repeated) types.
type RegionalPlanAssignment_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	RegionalPlanAssignment_FieldPath
	ContainsValue(*RegionalPlanAssignment) bool
}

// ParseRegionalPlanAssignment_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseRegionalPlanAssignment_FieldPathArrayItemValue(pathStr, valueStr string) (RegionalPlanAssignment_FieldPathArrayItemValue, error) {
	fp, err := ParseRegionalPlanAssignment_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing RegionalPlanAssignment field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(RegionalPlanAssignment_FieldPathArrayItemValue), nil
}

func MustParseRegionalPlanAssignment_FieldPathArrayItemValue(pathStr, valueStr string) RegionalPlanAssignment_FieldPathArrayItemValue {
	fpaiv, err := ParseRegionalPlanAssignment_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type RegionalPlanAssignment_FieldTerminalPathArrayItemValue struct {
	RegionalPlanAssignment_FieldTerminalPath
	value interface{}
}

var _ RegionalPlanAssignment_FieldPathArrayItemValue = (*RegionalPlanAssignment_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object RegionalPlanAssignment as interface{}
func (fpaiv *RegionalPlanAssignment_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *RegionalPlanAssignment_FieldTerminalPathArrayItemValue) GetSingle(source *RegionalPlanAssignment) (interface{}, bool) {
	return nil, false
}

func (fpaiv *RegionalPlanAssignment_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*RegionalPlanAssignment))
}

// Contains returns a boolean indicating if value that is being held is present in given 'RegionalPlanAssignment'
func (fpaiv *RegionalPlanAssignment_FieldTerminalPathArrayItemValue) ContainsValue(source *RegionalPlanAssignment) bool {
	slice := fpaiv.RegionalPlanAssignment_FieldTerminalPath.Get(source)
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

// RegionalPlanAssignment_FieldPathArrayOfValues allows storing slice of values for RegionalPlanAssignment fields according to their type
type RegionalPlanAssignment_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	RegionalPlanAssignment_FieldPath
}

func ParseRegionalPlanAssignment_FieldPathArrayOfValues(pathStr, valuesStr string) (RegionalPlanAssignment_FieldPathArrayOfValues, error) {
	fp, err := ParseRegionalPlanAssignment_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing RegionalPlanAssignment field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(RegionalPlanAssignment_FieldPathArrayOfValues), nil
}

func MustParseRegionalPlanAssignment_FieldPathArrayOfValues(pathStr, valuesStr string) RegionalPlanAssignment_FieldPathArrayOfValues {
	fpaov, err := ParseRegionalPlanAssignment_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type RegionalPlanAssignment_FieldTerminalPathArrayOfValues struct {
	RegionalPlanAssignment_FieldTerminalPath
	values interface{}
}

var _ RegionalPlanAssignment_FieldPathArrayOfValues = (*RegionalPlanAssignment_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *RegionalPlanAssignment_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case RegionalPlanAssignment_FieldPathSelectorPlan:
		for _, v := range fpaov.values.([]*plan.Reference) {
			values = append(values, v)
		}
	case RegionalPlanAssignment_FieldPathSelectorRegion:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case RegionalPlanAssignment_FieldPathSelectorPlanGeneration:
		for _, v := range fpaov.values.([]int64) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *RegionalPlanAssignment_FieldTerminalPathArrayOfValues) AsPlanArrayOfValues() ([]*plan.Reference, bool) {
	res, ok := fpaov.values.([]*plan.Reference)
	return res, ok
}
func (fpaov *RegionalPlanAssignment_FieldTerminalPathArrayOfValues) AsRegionArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *RegionalPlanAssignment_FieldTerminalPathArrayOfValues) AsPlanGenerationArrayOfValues() ([]int64, bool) {
	res, ok := fpaov.values.([]int64)
	return res, ok
}

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
	Allowance_FieldPathSelectorRegion   Allowance_FieldPathSelector = 2
)

func (s Allowance_FieldPathSelector) String() string {
	switch s {
	case Allowance_FieldPathSelectorResource:
		return "resource"
	case Allowance_FieldPathSelectorValue:
		return "value"
	case Allowance_FieldPathSelectorRegion:
		return "region"
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
		case "region":
			return &Allowance_FieldTerminalPath{selector: Allowance_FieldPathSelectorRegion}, nil
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
		case Allowance_FieldPathSelectorRegion:
			values = append(values, source.Region)
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
	case Allowance_FieldPathSelectorRegion:
		return source.GetRegion(), source != nil
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
	case Allowance_FieldPathSelectorRegion:
		return ""
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
		case Allowance_FieldPathSelectorRegion:
			item.Region = ""
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
		fp.selector == Allowance_FieldPathSelectorValue ||
		fp.selector == Allowance_FieldPathSelectorRegion
}

func (fp *Allowance_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *Allowance_FieldTerminalPath) WithIValue(value interface{}) Allowance_FieldPathValue {
	switch fp.selector {
	case Allowance_FieldPathSelectorResource:
		return &Allowance_FieldTerminalPathValue{Allowance_FieldTerminalPath: *fp, value: value.(*meta_resource.Reference)}
	case Allowance_FieldPathSelectorValue:
		return &Allowance_FieldTerminalPathValue{Allowance_FieldTerminalPath: *fp, value: value.(int64)}
	case Allowance_FieldPathSelectorRegion:
		return &Allowance_FieldTerminalPathValue{Allowance_FieldTerminalPath: *fp, value: value.(string)}
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
	case Allowance_FieldPathSelectorRegion:
		return &Allowance_FieldTerminalPathArrayOfValues{Allowance_FieldTerminalPath: *fp, values: values.([]string)}
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
func (fpv *Allowance_FieldTerminalPathValue) AsRegionValue() (string, bool) {
	res, ok := fpv.value.(string)
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
	case Allowance_FieldPathSelectorRegion:
		(*target).Region = fpv.value.(string)
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
	case Allowance_FieldPathSelectorRegion:
		leftValue := fpv.value.(string)
		rightValue := source.GetRegion()
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
	case Allowance_FieldPathSelectorRegion:
		for _, v := range fpaov.values.([]string) {
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
func (fpaov *Allowance_FieldTerminalPathArrayOfValues) AsRegionArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
