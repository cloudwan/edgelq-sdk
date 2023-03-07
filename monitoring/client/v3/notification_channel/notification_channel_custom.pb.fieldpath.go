// Code generated by protoc-gen-goten-object
// File: edgelq/monitoring/proto/v3/notification_channel_custom.proto
// DO NOT EDIT!!!

package notification_channel_client

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
	notification_channel "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/notification_channel"
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
	_ = fieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &notification_channel.NotificationChannel{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type TestNotificationChannelRequest_FieldPath interface {
	gotenobject.FieldPath
	Selector() TestNotificationChannelRequest_FieldPathSelector
	Get(source *TestNotificationChannelRequest) []interface{}
	GetSingle(source *TestNotificationChannelRequest) (interface{}, bool)
	ClearValue(item *TestNotificationChannelRequest)

	// Those methods build corresponding TestNotificationChannelRequest_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) TestNotificationChannelRequest_FieldPathValue
	WithIArrayOfValues(values interface{}) TestNotificationChannelRequest_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) TestNotificationChannelRequest_FieldPathArrayItemValue
}

type TestNotificationChannelRequest_FieldPathSelector int32

const (
	TestNotificationChannelRequest_FieldPathSelectorName TestNotificationChannelRequest_FieldPathSelector = 0
)

func (s TestNotificationChannelRequest_FieldPathSelector) String() string {
	switch s {
	case TestNotificationChannelRequest_FieldPathSelectorName:
		return "name"
	default:
		panic(fmt.Sprintf("Invalid selector for TestNotificationChannelRequest: %d", s))
	}
}

func BuildTestNotificationChannelRequest_FieldPath(fp gotenobject.RawFieldPath) (TestNotificationChannelRequest_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object TestNotificationChannelRequest")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &TestNotificationChannelRequest_FieldTerminalPath{selector: TestNotificationChannelRequest_FieldPathSelectorName}, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object TestNotificationChannelRequest", fp)
}

func ParseTestNotificationChannelRequest_FieldPath(rawField string) (TestNotificationChannelRequest_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildTestNotificationChannelRequest_FieldPath(fp)
}

func MustParseTestNotificationChannelRequest_FieldPath(rawField string) TestNotificationChannelRequest_FieldPath {
	fp, err := ParseTestNotificationChannelRequest_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type TestNotificationChannelRequest_FieldTerminalPath struct {
	selector TestNotificationChannelRequest_FieldPathSelector
}

var _ TestNotificationChannelRequest_FieldPath = (*TestNotificationChannelRequest_FieldTerminalPath)(nil)

func (fp *TestNotificationChannelRequest_FieldTerminalPath) Selector() TestNotificationChannelRequest_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *TestNotificationChannelRequest_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *TestNotificationChannelRequest_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source TestNotificationChannelRequest
func (fp *TestNotificationChannelRequest_FieldTerminalPath) Get(source *TestNotificationChannelRequest) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case TestNotificationChannelRequest_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		default:
			panic(fmt.Sprintf("Invalid selector for TestNotificationChannelRequest: %d", fp.selector))
		}
	}
	return
}

func (fp *TestNotificationChannelRequest_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*TestNotificationChannelRequest))
}

// GetSingle returns value pointed by specific field of from source TestNotificationChannelRequest
func (fp *TestNotificationChannelRequest_FieldTerminalPath) GetSingle(source *TestNotificationChannelRequest) (interface{}, bool) {
	switch fp.selector {
	case TestNotificationChannelRequest_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	default:
		panic(fmt.Sprintf("Invalid selector for TestNotificationChannelRequest: %d", fp.selector))
	}
}

func (fp *TestNotificationChannelRequest_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*TestNotificationChannelRequest))
}

// GetDefault returns a default value of the field type
func (fp *TestNotificationChannelRequest_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case TestNotificationChannelRequest_FieldPathSelectorName:
		return (*notification_channel.Reference)(nil)
	default:
		panic(fmt.Sprintf("Invalid selector for TestNotificationChannelRequest: %d", fp.selector))
	}
}

func (fp *TestNotificationChannelRequest_FieldTerminalPath) ClearValue(item *TestNotificationChannelRequest) {
	if item != nil {
		switch fp.selector {
		case TestNotificationChannelRequest_FieldPathSelectorName:
			item.Name = nil
		default:
			panic(fmt.Sprintf("Invalid selector for TestNotificationChannelRequest: %d", fp.selector))
		}
	}
}

func (fp *TestNotificationChannelRequest_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*TestNotificationChannelRequest))
}

// IsLeaf - whether field path is holds simple value
func (fp *TestNotificationChannelRequest_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == TestNotificationChannelRequest_FieldPathSelectorName
}

func (fp *TestNotificationChannelRequest_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *TestNotificationChannelRequest_FieldTerminalPath) WithIValue(value interface{}) TestNotificationChannelRequest_FieldPathValue {
	switch fp.selector {
	case TestNotificationChannelRequest_FieldPathSelectorName:
		return &TestNotificationChannelRequest_FieldTerminalPathValue{TestNotificationChannelRequest_FieldTerminalPath: *fp, value: value.(*notification_channel.Reference)}
	default:
		panic(fmt.Sprintf("Invalid selector for TestNotificationChannelRequest: %d", fp.selector))
	}
}

func (fp *TestNotificationChannelRequest_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *TestNotificationChannelRequest_FieldTerminalPath) WithIArrayOfValues(values interface{}) TestNotificationChannelRequest_FieldPathArrayOfValues {
	fpaov := &TestNotificationChannelRequest_FieldTerminalPathArrayOfValues{TestNotificationChannelRequest_FieldTerminalPath: *fp}
	switch fp.selector {
	case TestNotificationChannelRequest_FieldPathSelectorName:
		return &TestNotificationChannelRequest_FieldTerminalPathArrayOfValues{TestNotificationChannelRequest_FieldTerminalPath: *fp, values: values.([]*notification_channel.Reference)}
	default:
		panic(fmt.Sprintf("Invalid selector for TestNotificationChannelRequest: %d", fp.selector))
	}
	return fpaov
}

func (fp *TestNotificationChannelRequest_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *TestNotificationChannelRequest_FieldTerminalPath) WithIArrayItemValue(value interface{}) TestNotificationChannelRequest_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for TestNotificationChannelRequest: %d", fp.selector))
	}
}

func (fp *TestNotificationChannelRequest_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

// TestNotificationChannelRequest_FieldPathValue allows storing values for TestNotificationChannelRequest fields according to their type
type TestNotificationChannelRequest_FieldPathValue interface {
	TestNotificationChannelRequest_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **TestNotificationChannelRequest)
	CompareWith(*TestNotificationChannelRequest) (cmp int, comparable bool)
}

func ParseTestNotificationChannelRequest_FieldPathValue(pathStr, valueStr string) (TestNotificationChannelRequest_FieldPathValue, error) {
	fp, err := ParseTestNotificationChannelRequest_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing TestNotificationChannelRequest field path value from %s: %v", valueStr, err)
	}
	return fpv.(TestNotificationChannelRequest_FieldPathValue), nil
}

func MustParseTestNotificationChannelRequest_FieldPathValue(pathStr, valueStr string) TestNotificationChannelRequest_FieldPathValue {
	fpv, err := ParseTestNotificationChannelRequest_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type TestNotificationChannelRequest_FieldTerminalPathValue struct {
	TestNotificationChannelRequest_FieldTerminalPath
	value interface{}
}

var _ TestNotificationChannelRequest_FieldPathValue = (*TestNotificationChannelRequest_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'TestNotificationChannelRequest' as interface{}
func (fpv *TestNotificationChannelRequest_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *TestNotificationChannelRequest_FieldTerminalPathValue) AsNameValue() (*notification_channel.Reference, bool) {
	res, ok := fpv.value.(*notification_channel.Reference)
	return res, ok
}

// SetTo stores value for selected field for object TestNotificationChannelRequest
func (fpv *TestNotificationChannelRequest_FieldTerminalPathValue) SetTo(target **TestNotificationChannelRequest) {
	if *target == nil {
		*target = new(TestNotificationChannelRequest)
	}
	switch fpv.selector {
	case TestNotificationChannelRequest_FieldPathSelectorName:
		(*target).Name = fpv.value.(*notification_channel.Reference)
	default:
		panic(fmt.Sprintf("Invalid selector for TestNotificationChannelRequest: %d", fpv.selector))
	}
}

func (fpv *TestNotificationChannelRequest_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*TestNotificationChannelRequest)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'TestNotificationChannelRequest_FieldTerminalPathValue' with the value under path in 'TestNotificationChannelRequest'.
func (fpv *TestNotificationChannelRequest_FieldTerminalPathValue) CompareWith(source *TestNotificationChannelRequest) (int, bool) {
	switch fpv.selector {
	case TestNotificationChannelRequest_FieldPathSelectorName:
		leftValue := fpv.value.(*notification_channel.Reference)
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
	default:
		panic(fmt.Sprintf("Invalid selector for TestNotificationChannelRequest: %d", fpv.selector))
	}
}

func (fpv *TestNotificationChannelRequest_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*TestNotificationChannelRequest))
}

// TestNotificationChannelRequest_FieldPathArrayItemValue allows storing single item in Path-specific values for TestNotificationChannelRequest according to their type
// Present only for array (repeated) types.
type TestNotificationChannelRequest_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	TestNotificationChannelRequest_FieldPath
	ContainsValue(*TestNotificationChannelRequest) bool
}

// ParseTestNotificationChannelRequest_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseTestNotificationChannelRequest_FieldPathArrayItemValue(pathStr, valueStr string) (TestNotificationChannelRequest_FieldPathArrayItemValue, error) {
	fp, err := ParseTestNotificationChannelRequest_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing TestNotificationChannelRequest field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(TestNotificationChannelRequest_FieldPathArrayItemValue), nil
}

func MustParseTestNotificationChannelRequest_FieldPathArrayItemValue(pathStr, valueStr string) TestNotificationChannelRequest_FieldPathArrayItemValue {
	fpaiv, err := ParseTestNotificationChannelRequest_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type TestNotificationChannelRequest_FieldTerminalPathArrayItemValue struct {
	TestNotificationChannelRequest_FieldTerminalPath
	value interface{}
}

var _ TestNotificationChannelRequest_FieldPathArrayItemValue = (*TestNotificationChannelRequest_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object TestNotificationChannelRequest as interface{}
func (fpaiv *TestNotificationChannelRequest_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *TestNotificationChannelRequest_FieldTerminalPathArrayItemValue) GetSingle(source *TestNotificationChannelRequest) (interface{}, bool) {
	return nil, false
}

func (fpaiv *TestNotificationChannelRequest_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*TestNotificationChannelRequest))
}

// Contains returns a boolean indicating if value that is being held is present in given 'TestNotificationChannelRequest'
func (fpaiv *TestNotificationChannelRequest_FieldTerminalPathArrayItemValue) ContainsValue(source *TestNotificationChannelRequest) bool {
	slice := fpaiv.TestNotificationChannelRequest_FieldTerminalPath.Get(source)
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

// TestNotificationChannelRequest_FieldPathArrayOfValues allows storing slice of values for TestNotificationChannelRequest fields according to their type
type TestNotificationChannelRequest_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	TestNotificationChannelRequest_FieldPath
}

func ParseTestNotificationChannelRequest_FieldPathArrayOfValues(pathStr, valuesStr string) (TestNotificationChannelRequest_FieldPathArrayOfValues, error) {
	fp, err := ParseTestNotificationChannelRequest_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing TestNotificationChannelRequest field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(TestNotificationChannelRequest_FieldPathArrayOfValues), nil
}

func MustParseTestNotificationChannelRequest_FieldPathArrayOfValues(pathStr, valuesStr string) TestNotificationChannelRequest_FieldPathArrayOfValues {
	fpaov, err := ParseTestNotificationChannelRequest_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type TestNotificationChannelRequest_FieldTerminalPathArrayOfValues struct {
	TestNotificationChannelRequest_FieldTerminalPath
	values interface{}
}

var _ TestNotificationChannelRequest_FieldPathArrayOfValues = (*TestNotificationChannelRequest_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *TestNotificationChannelRequest_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case TestNotificationChannelRequest_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*notification_channel.Reference) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *TestNotificationChannelRequest_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*notification_channel.Reference, bool) {
	res, ok := fpaov.values.([]*notification_channel.Reference)
	return res, ok
}