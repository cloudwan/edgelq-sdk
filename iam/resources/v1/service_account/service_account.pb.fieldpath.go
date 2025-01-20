// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1/service_account.proto
// DO NOT EDIT!!!

package service_account

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
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
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
type ServiceAccount_FieldPath interface {
	gotenobject.FieldPath
	Selector() ServiceAccount_FieldPathSelector
	Get(source *ServiceAccount) []interface{}
	GetSingle(source *ServiceAccount) (interface{}, bool)
	ClearValue(item *ServiceAccount)

	// Those methods build corresponding ServiceAccount_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) ServiceAccount_FieldPathValue
	WithIArrayOfValues(values interface{}) ServiceAccount_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) ServiceAccount_FieldPathArrayItemValue
}

type ServiceAccount_FieldPathSelector int32

const (
	ServiceAccount_FieldPathSelectorName        ServiceAccount_FieldPathSelector = 0
	ServiceAccount_FieldPathSelectorMetadata    ServiceAccount_FieldPathSelector = 1
	ServiceAccount_FieldPathSelectorDisplayName ServiceAccount_FieldPathSelector = 2
	ServiceAccount_FieldPathSelectorDescription ServiceAccount_FieldPathSelector = 3
	ServiceAccount_FieldPathSelectorEmail       ServiceAccount_FieldPathSelector = 4
	ServiceAccount_FieldPathSelectorKind        ServiceAccount_FieldPathSelector = 5
)

func (s ServiceAccount_FieldPathSelector) String() string {
	switch s {
	case ServiceAccount_FieldPathSelectorName:
		return "name"
	case ServiceAccount_FieldPathSelectorMetadata:
		return "metadata"
	case ServiceAccount_FieldPathSelectorDisplayName:
		return "display_name"
	case ServiceAccount_FieldPathSelectorDescription:
		return "description"
	case ServiceAccount_FieldPathSelectorEmail:
		return "email"
	case ServiceAccount_FieldPathSelectorKind:
		return "kind"
	default:
		panic(fmt.Sprintf("Invalid selector for ServiceAccount: %d", s))
	}
}

func BuildServiceAccount_FieldPath(fp gotenobject.RawFieldPath) (ServiceAccount_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object ServiceAccount")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorName}, nil
		case "metadata":
			return &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorMetadata}, nil
		case "display_name", "displayName", "display-name":
			return &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorDisplayName}, nil
		case "description":
			return &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorDescription}, nil
		case "email":
			return &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorEmail}, nil
		case "kind":
			return &ServiceAccount_FieldTerminalPath{selector: ServiceAccount_FieldPathSelectorKind}, nil
		}
	} else {
		switch fp[0] {
		case "metadata":
			if subpath, err := meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &ServiceAccount_FieldSubPath{selector: ServiceAccount_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object ServiceAccount", fp)
}

func ParseServiceAccount_FieldPath(rawField string) (ServiceAccount_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildServiceAccount_FieldPath(fp)
}

func MustParseServiceAccount_FieldPath(rawField string) ServiceAccount_FieldPath {
	fp, err := ParseServiceAccount_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type ServiceAccount_FieldTerminalPath struct {
	selector ServiceAccount_FieldPathSelector
}

var _ ServiceAccount_FieldPath = (*ServiceAccount_FieldTerminalPath)(nil)

func (fp *ServiceAccount_FieldTerminalPath) Selector() ServiceAccount_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *ServiceAccount_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *ServiceAccount_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source ServiceAccount
func (fp *ServiceAccount_FieldTerminalPath) Get(source *ServiceAccount) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case ServiceAccount_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case ServiceAccount_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		case ServiceAccount_FieldPathSelectorDisplayName:
			values = append(values, source.DisplayName)
		case ServiceAccount_FieldPathSelectorDescription:
			values = append(values, source.Description)
		case ServiceAccount_FieldPathSelectorEmail:
			values = append(values, source.Email)
		case ServiceAccount_FieldPathSelectorKind:
			values = append(values, source.Kind)
		default:
			panic(fmt.Sprintf("Invalid selector for ServiceAccount: %d", fp.selector))
		}
	}
	return
}

func (fp *ServiceAccount_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*ServiceAccount))
}

// GetSingle returns value pointed by specific field of from source ServiceAccount
func (fp *ServiceAccount_FieldTerminalPath) GetSingle(source *ServiceAccount) (interface{}, bool) {
	switch fp.selector {
	case ServiceAccount_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case ServiceAccount_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	case ServiceAccount_FieldPathSelectorDisplayName:
		return source.GetDisplayName(), source != nil
	case ServiceAccount_FieldPathSelectorDescription:
		return source.GetDescription(), source != nil
	case ServiceAccount_FieldPathSelectorEmail:
		return source.GetEmail(), source != nil
	case ServiceAccount_FieldPathSelectorKind:
		return source.GetKind(), source != nil
	default:
		panic(fmt.Sprintf("Invalid selector for ServiceAccount: %d", fp.selector))
	}
}

func (fp *ServiceAccount_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*ServiceAccount))
}

// GetDefault returns a default value of the field type
func (fp *ServiceAccount_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case ServiceAccount_FieldPathSelectorName:
		return (*Name)(nil)
	case ServiceAccount_FieldPathSelectorMetadata:
		return (*meta.Meta)(nil)
	case ServiceAccount_FieldPathSelectorDisplayName:
		return ""
	case ServiceAccount_FieldPathSelectorDescription:
		return ""
	case ServiceAccount_FieldPathSelectorEmail:
		return ""
	case ServiceAccount_FieldPathSelectorKind:
		return ServiceAccount_UNSPECIFIED
	default:
		panic(fmt.Sprintf("Invalid selector for ServiceAccount: %d", fp.selector))
	}
}

func (fp *ServiceAccount_FieldTerminalPath) ClearValue(item *ServiceAccount) {
	if item != nil {
		switch fp.selector {
		case ServiceAccount_FieldPathSelectorName:
			item.Name = nil
		case ServiceAccount_FieldPathSelectorMetadata:
			item.Metadata = nil
		case ServiceAccount_FieldPathSelectorDisplayName:
			item.DisplayName = ""
		case ServiceAccount_FieldPathSelectorDescription:
			item.Description = ""
		case ServiceAccount_FieldPathSelectorEmail:
			item.Email = ""
		case ServiceAccount_FieldPathSelectorKind:
			item.Kind = ServiceAccount_UNSPECIFIED
		default:
			panic(fmt.Sprintf("Invalid selector for ServiceAccount: %d", fp.selector))
		}
	}
}

func (fp *ServiceAccount_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*ServiceAccount))
}

// IsLeaf - whether field path is holds simple value
func (fp *ServiceAccount_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == ServiceAccount_FieldPathSelectorName ||
		fp.selector == ServiceAccount_FieldPathSelectorDisplayName ||
		fp.selector == ServiceAccount_FieldPathSelectorDescription ||
		fp.selector == ServiceAccount_FieldPathSelectorEmail ||
		fp.selector == ServiceAccount_FieldPathSelectorKind
}

func (fp *ServiceAccount_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *ServiceAccount_FieldTerminalPath) WithIValue(value interface{}) ServiceAccount_FieldPathValue {
	switch fp.selector {
	case ServiceAccount_FieldPathSelectorName:
		return &ServiceAccount_FieldTerminalPathValue{ServiceAccount_FieldTerminalPath: *fp, value: value.(*Name)}
	case ServiceAccount_FieldPathSelectorMetadata:
		return &ServiceAccount_FieldTerminalPathValue{ServiceAccount_FieldTerminalPath: *fp, value: value.(*meta.Meta)}
	case ServiceAccount_FieldPathSelectorDisplayName:
		return &ServiceAccount_FieldTerminalPathValue{ServiceAccount_FieldTerminalPath: *fp, value: value.(string)}
	case ServiceAccount_FieldPathSelectorDescription:
		return &ServiceAccount_FieldTerminalPathValue{ServiceAccount_FieldTerminalPath: *fp, value: value.(string)}
	case ServiceAccount_FieldPathSelectorEmail:
		return &ServiceAccount_FieldTerminalPathValue{ServiceAccount_FieldTerminalPath: *fp, value: value.(string)}
	case ServiceAccount_FieldPathSelectorKind:
		return &ServiceAccount_FieldTerminalPathValue{ServiceAccount_FieldTerminalPath: *fp, value: value.(ServiceAccount_Kind)}
	default:
		panic(fmt.Sprintf("Invalid selector for ServiceAccount: %d", fp.selector))
	}
}

func (fp *ServiceAccount_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *ServiceAccount_FieldTerminalPath) WithIArrayOfValues(values interface{}) ServiceAccount_FieldPathArrayOfValues {
	fpaov := &ServiceAccount_FieldTerminalPathArrayOfValues{ServiceAccount_FieldTerminalPath: *fp}
	switch fp.selector {
	case ServiceAccount_FieldPathSelectorName:
		return &ServiceAccount_FieldTerminalPathArrayOfValues{ServiceAccount_FieldTerminalPath: *fp, values: values.([]*Name)}
	case ServiceAccount_FieldPathSelectorMetadata:
		return &ServiceAccount_FieldTerminalPathArrayOfValues{ServiceAccount_FieldTerminalPath: *fp, values: values.([]*meta.Meta)}
	case ServiceAccount_FieldPathSelectorDisplayName:
		return &ServiceAccount_FieldTerminalPathArrayOfValues{ServiceAccount_FieldTerminalPath: *fp, values: values.([]string)}
	case ServiceAccount_FieldPathSelectorDescription:
		return &ServiceAccount_FieldTerminalPathArrayOfValues{ServiceAccount_FieldTerminalPath: *fp, values: values.([]string)}
	case ServiceAccount_FieldPathSelectorEmail:
		return &ServiceAccount_FieldTerminalPathArrayOfValues{ServiceAccount_FieldTerminalPath: *fp, values: values.([]string)}
	case ServiceAccount_FieldPathSelectorKind:
		return &ServiceAccount_FieldTerminalPathArrayOfValues{ServiceAccount_FieldTerminalPath: *fp, values: values.([]ServiceAccount_Kind)}
	default:
		panic(fmt.Sprintf("Invalid selector for ServiceAccount: %d", fp.selector))
	}
	return fpaov
}

func (fp *ServiceAccount_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *ServiceAccount_FieldTerminalPath) WithIArrayItemValue(value interface{}) ServiceAccount_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for ServiceAccount: %d", fp.selector))
	}
}

func (fp *ServiceAccount_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

type ServiceAccount_FieldSubPath struct {
	selector ServiceAccount_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ ServiceAccount_FieldPath = (*ServiceAccount_FieldSubPath)(nil)

func (fps *ServiceAccount_FieldSubPath) Selector() ServiceAccount_FieldPathSelector {
	return fps.selector
}
func (fps *ServiceAccount_FieldSubPath) AsMetadataSubPath() (meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(meta.Meta_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *ServiceAccount_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *ServiceAccount_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source ServiceAccount
func (fps *ServiceAccount_FieldSubPath) Get(source *ServiceAccount) (values []interface{}) {
	switch fps.selector {
	case ServiceAccount_FieldPathSelectorMetadata:
		values = append(values, fps.subPath.GetRaw(source.GetMetadata())...)
	default:
		panic(fmt.Sprintf("Invalid selector for ServiceAccount: %d", fps.selector))
	}
	return
}

func (fps *ServiceAccount_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*ServiceAccount))
}

// GetSingle returns value of selected field from source ServiceAccount
func (fps *ServiceAccount_FieldSubPath) GetSingle(source *ServiceAccount) (interface{}, bool) {
	switch fps.selector {
	case ServiceAccount_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for ServiceAccount: %d", fps.selector))
	}
}

func (fps *ServiceAccount_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*ServiceAccount))
}

// GetDefault returns a default value of the field type
func (fps *ServiceAccount_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *ServiceAccount_FieldSubPath) ClearValue(item *ServiceAccount) {
	if item != nil {
		switch fps.selector {
		case ServiceAccount_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		default:
			panic(fmt.Sprintf("Invalid selector for ServiceAccount: %d", fps.selector))
		}
	}
}

func (fps *ServiceAccount_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*ServiceAccount))
}

// IsLeaf - whether field path is holds simple value
func (fps *ServiceAccount_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *ServiceAccount_FieldSubPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	iPaths := []gotenobject.FieldPath{&ServiceAccount_FieldTerminalPath{selector: fps.selector}}
	iPaths = append(iPaths, fps.subPath.SplitIntoTerminalIPaths()...)
	return iPaths
}

func (fps *ServiceAccount_FieldSubPath) WithIValue(value interface{}) ServiceAccount_FieldPathValue {
	return &ServiceAccount_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *ServiceAccount_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *ServiceAccount_FieldSubPath) WithIArrayOfValues(values interface{}) ServiceAccount_FieldPathArrayOfValues {
	return &ServiceAccount_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *ServiceAccount_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *ServiceAccount_FieldSubPath) WithIArrayItemValue(value interface{}) ServiceAccount_FieldPathArrayItemValue {
	return &ServiceAccount_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *ServiceAccount_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// ServiceAccount_FieldPathValue allows storing values for ServiceAccount fields according to their type
type ServiceAccount_FieldPathValue interface {
	ServiceAccount_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **ServiceAccount)
	CompareWith(*ServiceAccount) (cmp int, comparable bool)
}

func ParseServiceAccount_FieldPathValue(pathStr, valueStr string) (ServiceAccount_FieldPathValue, error) {
	fp, err := ParseServiceAccount_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing ServiceAccount field path value from %s: %v", valueStr, err)
	}
	return fpv.(ServiceAccount_FieldPathValue), nil
}

func MustParseServiceAccount_FieldPathValue(pathStr, valueStr string) ServiceAccount_FieldPathValue {
	fpv, err := ParseServiceAccount_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type ServiceAccount_FieldTerminalPathValue struct {
	ServiceAccount_FieldTerminalPath
	value interface{}
}

var _ ServiceAccount_FieldPathValue = (*ServiceAccount_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'ServiceAccount' as interface{}
func (fpv *ServiceAccount_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *ServiceAccount_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *ServiceAccount_FieldTerminalPathValue) AsMetadataValue() (*meta.Meta, bool) {
	res, ok := fpv.value.(*meta.Meta)
	return res, ok
}
func (fpv *ServiceAccount_FieldTerminalPathValue) AsDisplayNameValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *ServiceAccount_FieldTerminalPathValue) AsDescriptionValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *ServiceAccount_FieldTerminalPathValue) AsEmailValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *ServiceAccount_FieldTerminalPathValue) AsKindValue() (ServiceAccount_Kind, bool) {
	res, ok := fpv.value.(ServiceAccount_Kind)
	return res, ok
}

// SetTo stores value for selected field for object ServiceAccount
func (fpv *ServiceAccount_FieldTerminalPathValue) SetTo(target **ServiceAccount) {
	if *target == nil {
		*target = new(ServiceAccount)
	}
	switch fpv.selector {
	case ServiceAccount_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case ServiceAccount_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*meta.Meta)
	case ServiceAccount_FieldPathSelectorDisplayName:
		(*target).DisplayName = fpv.value.(string)
	case ServiceAccount_FieldPathSelectorDescription:
		(*target).Description = fpv.value.(string)
	case ServiceAccount_FieldPathSelectorEmail:
		(*target).Email = fpv.value.(string)
	case ServiceAccount_FieldPathSelectorKind:
		(*target).Kind = fpv.value.(ServiceAccount_Kind)
	default:
		panic(fmt.Sprintf("Invalid selector for ServiceAccount: %d", fpv.selector))
	}
}

func (fpv *ServiceAccount_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*ServiceAccount)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'ServiceAccount_FieldTerminalPathValue' with the value under path in 'ServiceAccount'.
func (fpv *ServiceAccount_FieldTerminalPathValue) CompareWith(source *ServiceAccount) (int, bool) {
	switch fpv.selector {
	case ServiceAccount_FieldPathSelectorName:
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
	case ServiceAccount_FieldPathSelectorMetadata:
		return 0, false
	case ServiceAccount_FieldPathSelectorDisplayName:
		leftValue := fpv.value.(string)
		rightValue := source.GetDisplayName()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case ServiceAccount_FieldPathSelectorDescription:
		leftValue := fpv.value.(string)
		rightValue := source.GetDescription()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case ServiceAccount_FieldPathSelectorEmail:
		leftValue := fpv.value.(string)
		rightValue := source.GetEmail()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case ServiceAccount_FieldPathSelectorKind:
		leftValue := fpv.value.(ServiceAccount_Kind)
		rightValue := source.GetKind()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	default:
		panic(fmt.Sprintf("Invalid selector for ServiceAccount: %d", fpv.selector))
	}
}

func (fpv *ServiceAccount_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*ServiceAccount))
}

type ServiceAccount_FieldSubPathValue struct {
	ServiceAccount_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ ServiceAccount_FieldPathValue = (*ServiceAccount_FieldSubPathValue)(nil)

func (fpvs *ServiceAccount_FieldSubPathValue) AsMetadataPathValue() (meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(meta.Meta_FieldPathValue)
	return res, ok
}

func (fpvs *ServiceAccount_FieldSubPathValue) SetTo(target **ServiceAccount) {
	if *target == nil {
		*target = new(ServiceAccount)
	}
	switch fpvs.Selector() {
	case ServiceAccount_FieldPathSelectorMetadata:
		fpvs.subPathValue.(meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	default:
		panic(fmt.Sprintf("Invalid selector for ServiceAccount: %d", fpvs.Selector()))
	}
}

func (fpvs *ServiceAccount_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*ServiceAccount)
	fpvs.SetTo(&typedObject)
}

func (fpvs *ServiceAccount_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *ServiceAccount_FieldSubPathValue) CompareWith(source *ServiceAccount) (int, bool) {
	switch fpvs.Selector() {
	case ServiceAccount_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for ServiceAccount: %d", fpvs.Selector()))
	}
}

func (fpvs *ServiceAccount_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*ServiceAccount))
}

// ServiceAccount_FieldPathArrayItemValue allows storing single item in Path-specific values for ServiceAccount according to their type
// Present only for array (repeated) types.
type ServiceAccount_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	ServiceAccount_FieldPath
	ContainsValue(*ServiceAccount) bool
}

// ParseServiceAccount_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseServiceAccount_FieldPathArrayItemValue(pathStr, valueStr string) (ServiceAccount_FieldPathArrayItemValue, error) {
	fp, err := ParseServiceAccount_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing ServiceAccount field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(ServiceAccount_FieldPathArrayItemValue), nil
}

func MustParseServiceAccount_FieldPathArrayItemValue(pathStr, valueStr string) ServiceAccount_FieldPathArrayItemValue {
	fpaiv, err := ParseServiceAccount_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type ServiceAccount_FieldTerminalPathArrayItemValue struct {
	ServiceAccount_FieldTerminalPath
	value interface{}
}

var _ ServiceAccount_FieldPathArrayItemValue = (*ServiceAccount_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object ServiceAccount as interface{}
func (fpaiv *ServiceAccount_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *ServiceAccount_FieldTerminalPathArrayItemValue) GetSingle(source *ServiceAccount) (interface{}, bool) {
	return nil, false
}

func (fpaiv *ServiceAccount_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*ServiceAccount))
}

// Contains returns a boolean indicating if value that is being held is present in given 'ServiceAccount'
func (fpaiv *ServiceAccount_FieldTerminalPathArrayItemValue) ContainsValue(source *ServiceAccount) bool {
	slice := fpaiv.ServiceAccount_FieldTerminalPath.Get(source)
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

type ServiceAccount_FieldSubPathArrayItemValue struct {
	ServiceAccount_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *ServiceAccount_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *ServiceAccount_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'ServiceAccount'
func (fpaivs *ServiceAccount_FieldSubPathArrayItemValue) ContainsValue(source *ServiceAccount) bool {
	switch fpaivs.Selector() {
	case ServiceAccount_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for ServiceAccount: %d", fpaivs.Selector()))
	}
}

// ServiceAccount_FieldPathArrayOfValues allows storing slice of values for ServiceAccount fields according to their type
type ServiceAccount_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	ServiceAccount_FieldPath
}

func ParseServiceAccount_FieldPathArrayOfValues(pathStr, valuesStr string) (ServiceAccount_FieldPathArrayOfValues, error) {
	fp, err := ParseServiceAccount_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing ServiceAccount field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(ServiceAccount_FieldPathArrayOfValues), nil
}

func MustParseServiceAccount_FieldPathArrayOfValues(pathStr, valuesStr string) ServiceAccount_FieldPathArrayOfValues {
	fpaov, err := ParseServiceAccount_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type ServiceAccount_FieldTerminalPathArrayOfValues struct {
	ServiceAccount_FieldTerminalPath
	values interface{}
}

var _ ServiceAccount_FieldPathArrayOfValues = (*ServiceAccount_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *ServiceAccount_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case ServiceAccount_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case ServiceAccount_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*meta.Meta) {
			values = append(values, v)
		}
	case ServiceAccount_FieldPathSelectorDisplayName:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case ServiceAccount_FieldPathSelectorDescription:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case ServiceAccount_FieldPathSelectorEmail:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case ServiceAccount_FieldPathSelectorKind:
		for _, v := range fpaov.values.([]ServiceAccount_Kind) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *ServiceAccount_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *ServiceAccount_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*meta.Meta, bool) {
	res, ok := fpaov.values.([]*meta.Meta)
	return res, ok
}
func (fpaov *ServiceAccount_FieldTerminalPathArrayOfValues) AsDisplayNameArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *ServiceAccount_FieldTerminalPathArrayOfValues) AsDescriptionArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *ServiceAccount_FieldTerminalPathArrayOfValues) AsEmailArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *ServiceAccount_FieldTerminalPathArrayOfValues) AsKindArrayOfValues() ([]ServiceAccount_Kind, bool) {
	res, ok := fpaov.values.([]ServiceAccount_Kind)
	return res, ok
}

type ServiceAccount_FieldSubPathArrayOfValues struct {
	ServiceAccount_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ ServiceAccount_FieldPathArrayOfValues = (*ServiceAccount_FieldSubPathArrayOfValues)(nil)

func (fpsaov *ServiceAccount_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *ServiceAccount_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
