// Code generated by protoc-gen-goten-object
// File: edgelq/common/rpc/status.proto
// DO NOT EDIT!!!

package rpc

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
	anypb "google.golang.org/protobuf/types/known/anypb"
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
	_ = &anypb.Any{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type Status_FieldPath interface {
	gotenobject.FieldPath
	Selector() Status_FieldPathSelector
	Get(source *Status) []interface{}
	GetSingle(source *Status) (interface{}, bool)
	ClearValue(item *Status)

	// Those methods build corresponding Status_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) Status_FieldPathValue
	WithIArrayOfValues(values interface{}) Status_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) Status_FieldPathArrayItemValue
}

type Status_FieldPathSelector int32

const (
	Status_FieldPathSelectorCode    Status_FieldPathSelector = 0
	Status_FieldPathSelectorMessage Status_FieldPathSelector = 1
	Status_FieldPathSelectorDetails Status_FieldPathSelector = 2
)

func (s Status_FieldPathSelector) String() string {
	switch s {
	case Status_FieldPathSelectorCode:
		return "code"
	case Status_FieldPathSelectorMessage:
		return "message"
	case Status_FieldPathSelectorDetails:
		return "details"
	default:
		panic(fmt.Sprintf("Invalid selector for Status: %d", s))
	}
}

func BuildStatus_FieldPath(fp gotenobject.RawFieldPath) (Status_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object Status")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "code":
			return &Status_FieldTerminalPath{selector: Status_FieldPathSelectorCode}, nil
		case "message":
			return &Status_FieldTerminalPath{selector: Status_FieldPathSelectorMessage}, nil
		case "details":
			return &Status_FieldTerminalPath{selector: Status_FieldPathSelectorDetails}, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object Status", fp)
}

func ParseStatus_FieldPath(rawField string) (Status_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildStatus_FieldPath(fp)
}

func MustParseStatus_FieldPath(rawField string) Status_FieldPath {
	fp, err := ParseStatus_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type Status_FieldTerminalPath struct {
	selector Status_FieldPathSelector
}

var _ Status_FieldPath = (*Status_FieldTerminalPath)(nil)

func (fp *Status_FieldTerminalPath) Selector() Status_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *Status_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *Status_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source Status
func (fp *Status_FieldTerminalPath) Get(source *Status) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case Status_FieldPathSelectorCode:
			values = append(values, source.Code)
		case Status_FieldPathSelectorMessage:
			values = append(values, source.Message)
		case Status_FieldPathSelectorDetails:
			for _, value := range source.GetDetails() {
				values = append(values, value)
			}
		default:
			panic(fmt.Sprintf("Invalid selector for Status: %d", fp.selector))
		}
	}
	return
}

func (fp *Status_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*Status))
}

// GetSingle returns value pointed by specific field of from source Status
func (fp *Status_FieldTerminalPath) GetSingle(source *Status) (interface{}, bool) {
	switch fp.selector {
	case Status_FieldPathSelectorCode:
		return source.GetCode(), source != nil
	case Status_FieldPathSelectorMessage:
		return source.GetMessage(), source != nil
	case Status_FieldPathSelectorDetails:
		res := source.GetDetails()
		return res, res != nil
	default:
		panic(fmt.Sprintf("Invalid selector for Status: %d", fp.selector))
	}
}

func (fp *Status_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*Status))
}

// GetDefault returns a default value of the field type
func (fp *Status_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case Status_FieldPathSelectorCode:
		return int32(0)
	case Status_FieldPathSelectorMessage:
		return ""
	case Status_FieldPathSelectorDetails:
		return ([]*anypb.Any)(nil)
	default:
		panic(fmt.Sprintf("Invalid selector for Status: %d", fp.selector))
	}
}

func (fp *Status_FieldTerminalPath) ClearValue(item *Status) {
	if item != nil {
		switch fp.selector {
		case Status_FieldPathSelectorCode:
			item.Code = int32(0)
		case Status_FieldPathSelectorMessage:
			item.Message = ""
		case Status_FieldPathSelectorDetails:
			item.Details = nil
		default:
			panic(fmt.Sprintf("Invalid selector for Status: %d", fp.selector))
		}
	}
}

func (fp *Status_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*Status))
}

// IsLeaf - whether field path is holds simple value
func (fp *Status_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == Status_FieldPathSelectorCode ||
		fp.selector == Status_FieldPathSelectorMessage ||
		fp.selector == Status_FieldPathSelectorDetails
}

func (fp *Status_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *Status_FieldTerminalPath) WithIValue(value interface{}) Status_FieldPathValue {
	switch fp.selector {
	case Status_FieldPathSelectorCode:
		return &Status_FieldTerminalPathValue{Status_FieldTerminalPath: *fp, value: value.(int32)}
	case Status_FieldPathSelectorMessage:
		return &Status_FieldTerminalPathValue{Status_FieldTerminalPath: *fp, value: value.(string)}
	case Status_FieldPathSelectorDetails:
		return &Status_FieldTerminalPathValue{Status_FieldTerminalPath: *fp, value: value.([]*anypb.Any)}
	default:
		panic(fmt.Sprintf("Invalid selector for Status: %d", fp.selector))
	}
}

func (fp *Status_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *Status_FieldTerminalPath) WithIArrayOfValues(values interface{}) Status_FieldPathArrayOfValues {
	fpaov := &Status_FieldTerminalPathArrayOfValues{Status_FieldTerminalPath: *fp}
	switch fp.selector {
	case Status_FieldPathSelectorCode:
		return &Status_FieldTerminalPathArrayOfValues{Status_FieldTerminalPath: *fp, values: values.([]int32)}
	case Status_FieldPathSelectorMessage:
		return &Status_FieldTerminalPathArrayOfValues{Status_FieldTerminalPath: *fp, values: values.([]string)}
	case Status_FieldPathSelectorDetails:
		return &Status_FieldTerminalPathArrayOfValues{Status_FieldTerminalPath: *fp, values: values.([][]*anypb.Any)}
	default:
		panic(fmt.Sprintf("Invalid selector for Status: %d", fp.selector))
	}
	return fpaov
}

func (fp *Status_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *Status_FieldTerminalPath) WithIArrayItemValue(value interface{}) Status_FieldPathArrayItemValue {
	switch fp.selector {
	case Status_FieldPathSelectorDetails:
		return &Status_FieldTerminalPathArrayItemValue{Status_FieldTerminalPath: *fp, value: value.(*anypb.Any)}
	default:
		panic(fmt.Sprintf("Invalid selector for Status: %d", fp.selector))
	}
}

func (fp *Status_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

// Status_FieldPathValue allows storing values for Status fields according to their type
type Status_FieldPathValue interface {
	Status_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **Status)
	CompareWith(*Status) (cmp int, comparable bool)
}

func ParseStatus_FieldPathValue(pathStr, valueStr string) (Status_FieldPathValue, error) {
	fp, err := ParseStatus_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Status field path value from %s: %v", valueStr, err)
	}
	return fpv.(Status_FieldPathValue), nil
}

func MustParseStatus_FieldPathValue(pathStr, valueStr string) Status_FieldPathValue {
	fpv, err := ParseStatus_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type Status_FieldTerminalPathValue struct {
	Status_FieldTerminalPath
	value interface{}
}

var _ Status_FieldPathValue = (*Status_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'Status' as interface{}
func (fpv *Status_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *Status_FieldTerminalPathValue) AsCodeValue() (int32, bool) {
	res, ok := fpv.value.(int32)
	return res, ok
}
func (fpv *Status_FieldTerminalPathValue) AsMessageValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *Status_FieldTerminalPathValue) AsDetailsValue() ([]*anypb.Any, bool) {
	res, ok := fpv.value.([]*anypb.Any)
	return res, ok
}

// SetTo stores value for selected field for object Status
func (fpv *Status_FieldTerminalPathValue) SetTo(target **Status) {
	if *target == nil {
		*target = new(Status)
	}
	switch fpv.selector {
	case Status_FieldPathSelectorCode:
		(*target).Code = fpv.value.(int32)
	case Status_FieldPathSelectorMessage:
		(*target).Message = fpv.value.(string)
	case Status_FieldPathSelectorDetails:
		(*target).Details = fpv.value.([]*anypb.Any)
	default:
		panic(fmt.Sprintf("Invalid selector for Status: %d", fpv.selector))
	}
}

func (fpv *Status_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Status)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'Status_FieldTerminalPathValue' with the value under path in 'Status'.
func (fpv *Status_FieldTerminalPathValue) CompareWith(source *Status) (int, bool) {
	switch fpv.selector {
	case Status_FieldPathSelectorCode:
		leftValue := fpv.value.(int32)
		rightValue := source.GetCode()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case Status_FieldPathSelectorMessage:
		leftValue := fpv.value.(string)
		rightValue := source.GetMessage()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case Status_FieldPathSelectorDetails:
		return 0, false
	default:
		panic(fmt.Sprintf("Invalid selector for Status: %d", fpv.selector))
	}
}

func (fpv *Status_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*Status))
}

// Status_FieldPathArrayItemValue allows storing single item in Path-specific values for Status according to their type
// Present only for array (repeated) types.
type Status_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	Status_FieldPath
	ContainsValue(*Status) bool
}

// ParseStatus_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseStatus_FieldPathArrayItemValue(pathStr, valueStr string) (Status_FieldPathArrayItemValue, error) {
	fp, err := ParseStatus_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Status field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(Status_FieldPathArrayItemValue), nil
}

func MustParseStatus_FieldPathArrayItemValue(pathStr, valueStr string) Status_FieldPathArrayItemValue {
	fpaiv, err := ParseStatus_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type Status_FieldTerminalPathArrayItemValue struct {
	Status_FieldTerminalPath
	value interface{}
}

var _ Status_FieldPathArrayItemValue = (*Status_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object Status as interface{}
func (fpaiv *Status_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}
func (fpaiv *Status_FieldTerminalPathArrayItemValue) AsDetailsItemValue() (*anypb.Any, bool) {
	res, ok := fpaiv.value.(*anypb.Any)
	return res, ok
}

func (fpaiv *Status_FieldTerminalPathArrayItemValue) GetSingle(source *Status) (interface{}, bool) {
	return nil, false
}

func (fpaiv *Status_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*Status))
}

// Contains returns a boolean indicating if value that is being held is present in given 'Status'
func (fpaiv *Status_FieldTerminalPathArrayItemValue) ContainsValue(source *Status) bool {
	slice := fpaiv.Status_FieldTerminalPath.Get(source)
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

// Status_FieldPathArrayOfValues allows storing slice of values for Status fields according to their type
type Status_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	Status_FieldPath
}

func ParseStatus_FieldPathArrayOfValues(pathStr, valuesStr string) (Status_FieldPathArrayOfValues, error) {
	fp, err := ParseStatus_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Status field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(Status_FieldPathArrayOfValues), nil
}

func MustParseStatus_FieldPathArrayOfValues(pathStr, valuesStr string) Status_FieldPathArrayOfValues {
	fpaov, err := ParseStatus_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type Status_FieldTerminalPathArrayOfValues struct {
	Status_FieldTerminalPath
	values interface{}
}

var _ Status_FieldPathArrayOfValues = (*Status_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *Status_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case Status_FieldPathSelectorCode:
		for _, v := range fpaov.values.([]int32) {
			values = append(values, v)
		}
	case Status_FieldPathSelectorMessage:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case Status_FieldPathSelectorDetails:
		for _, v := range fpaov.values.([][]*anypb.Any) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *Status_FieldTerminalPathArrayOfValues) AsCodeArrayOfValues() ([]int32, bool) {
	res, ok := fpaov.values.([]int32)
	return res, ok
}
func (fpaov *Status_FieldTerminalPathArrayOfValues) AsMessageArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *Status_FieldTerminalPathArrayOfValues) AsDetailsArrayOfValues() ([][]*anypb.Any, bool) {
	res, ok := fpaov.values.([][]*anypb.Any)
	return res, ok
}
