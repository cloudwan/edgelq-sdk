// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1/group.proto
// DO NOT EDIT!!!

package group

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
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
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
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type Group_FieldPath interface {
	gotenobject.FieldPath
	Selector() Group_FieldPathSelector
	Get(source *Group) []interface{}
	GetSingle(source *Group) (interface{}, bool)
	ClearValue(item *Group)

	// Those methods build corresponding Group_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) Group_FieldPathValue
	WithIArrayOfValues(values interface{}) Group_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) Group_FieldPathArrayItemValue
}

type Group_FieldPathSelector int32

const (
	Group_FieldPathSelectorName        Group_FieldPathSelector = 0
	Group_FieldPathSelectorMetadata    Group_FieldPathSelector = 1
	Group_FieldPathSelectorDisplayName Group_FieldPathSelector = 2
	Group_FieldPathSelectorDescription Group_FieldPathSelector = 3
	Group_FieldPathSelectorEmail       Group_FieldPathSelector = 4
)

func (s Group_FieldPathSelector) String() string {
	switch s {
	case Group_FieldPathSelectorName:
		return "name"
	case Group_FieldPathSelectorMetadata:
		return "metadata"
	case Group_FieldPathSelectorDisplayName:
		return "display_name"
	case Group_FieldPathSelectorDescription:
		return "description"
	case Group_FieldPathSelectorEmail:
		return "email"
	default:
		panic(fmt.Sprintf("Invalid selector for Group: %d", s))
	}
}

func BuildGroup_FieldPath(fp gotenobject.RawFieldPath) (Group_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object Group")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &Group_FieldTerminalPath{selector: Group_FieldPathSelectorName}, nil
		case "metadata":
			return &Group_FieldTerminalPath{selector: Group_FieldPathSelectorMetadata}, nil
		case "display_name", "displayName", "display-name":
			return &Group_FieldTerminalPath{selector: Group_FieldPathSelectorDisplayName}, nil
		case "description":
			return &Group_FieldTerminalPath{selector: Group_FieldPathSelectorDescription}, nil
		case "email":
			return &Group_FieldTerminalPath{selector: Group_FieldPathSelectorEmail}, nil
		}
	} else {
		switch fp[0] {
		case "metadata":
			if subpath, err := meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &Group_FieldSubPath{selector: Group_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object Group", fp)
}

func ParseGroup_FieldPath(rawField string) (Group_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildGroup_FieldPath(fp)
}

func MustParseGroup_FieldPath(rawField string) Group_FieldPath {
	fp, err := ParseGroup_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type Group_FieldTerminalPath struct {
	selector Group_FieldPathSelector
}

var _ Group_FieldPath = (*Group_FieldTerminalPath)(nil)

func (fp *Group_FieldTerminalPath) Selector() Group_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *Group_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *Group_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source Group
func (fp *Group_FieldTerminalPath) Get(source *Group) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case Group_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case Group_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		case Group_FieldPathSelectorDisplayName:
			values = append(values, source.DisplayName)
		case Group_FieldPathSelectorDescription:
			values = append(values, source.Description)
		case Group_FieldPathSelectorEmail:
			values = append(values, source.Email)
		default:
			panic(fmt.Sprintf("Invalid selector for Group: %d", fp.selector))
		}
	}
	return
}

func (fp *Group_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*Group))
}

// GetSingle returns value pointed by specific field of from source Group
func (fp *Group_FieldTerminalPath) GetSingle(source *Group) (interface{}, bool) {
	switch fp.selector {
	case Group_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case Group_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	case Group_FieldPathSelectorDisplayName:
		return source.GetDisplayName(), source != nil
	case Group_FieldPathSelectorDescription:
		return source.GetDescription(), source != nil
	case Group_FieldPathSelectorEmail:
		return source.GetEmail(), source != nil
	default:
		panic(fmt.Sprintf("Invalid selector for Group: %d", fp.selector))
	}
}

func (fp *Group_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*Group))
}

// GetDefault returns a default value of the field type
func (fp *Group_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case Group_FieldPathSelectorName:
		return (*Name)(nil)
	case Group_FieldPathSelectorMetadata:
		return (*meta.Meta)(nil)
	case Group_FieldPathSelectorDisplayName:
		return ""
	case Group_FieldPathSelectorDescription:
		return ""
	case Group_FieldPathSelectorEmail:
		return ""
	default:
		panic(fmt.Sprintf("Invalid selector for Group: %d", fp.selector))
	}
}

func (fp *Group_FieldTerminalPath) ClearValue(item *Group) {
	if item != nil {
		switch fp.selector {
		case Group_FieldPathSelectorName:
			item.Name = nil
		case Group_FieldPathSelectorMetadata:
			item.Metadata = nil
		case Group_FieldPathSelectorDisplayName:
			item.DisplayName = ""
		case Group_FieldPathSelectorDescription:
			item.Description = ""
		case Group_FieldPathSelectorEmail:
			item.Email = ""
		default:
			panic(fmt.Sprintf("Invalid selector for Group: %d", fp.selector))
		}
	}
}

func (fp *Group_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*Group))
}

// IsLeaf - whether field path is holds simple value
func (fp *Group_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == Group_FieldPathSelectorName ||
		fp.selector == Group_FieldPathSelectorDisplayName ||
		fp.selector == Group_FieldPathSelectorDescription ||
		fp.selector == Group_FieldPathSelectorEmail
}

func (fp *Group_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *Group_FieldTerminalPath) WithIValue(value interface{}) Group_FieldPathValue {
	switch fp.selector {
	case Group_FieldPathSelectorName:
		return &Group_FieldTerminalPathValue{Group_FieldTerminalPath: *fp, value: value.(*Name)}
	case Group_FieldPathSelectorMetadata:
		return &Group_FieldTerminalPathValue{Group_FieldTerminalPath: *fp, value: value.(*meta.Meta)}
	case Group_FieldPathSelectorDisplayName:
		return &Group_FieldTerminalPathValue{Group_FieldTerminalPath: *fp, value: value.(string)}
	case Group_FieldPathSelectorDescription:
		return &Group_FieldTerminalPathValue{Group_FieldTerminalPath: *fp, value: value.(string)}
	case Group_FieldPathSelectorEmail:
		return &Group_FieldTerminalPathValue{Group_FieldTerminalPath: *fp, value: value.(string)}
	default:
		panic(fmt.Sprintf("Invalid selector for Group: %d", fp.selector))
	}
}

func (fp *Group_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *Group_FieldTerminalPath) WithIArrayOfValues(values interface{}) Group_FieldPathArrayOfValues {
	fpaov := &Group_FieldTerminalPathArrayOfValues{Group_FieldTerminalPath: *fp}
	switch fp.selector {
	case Group_FieldPathSelectorName:
		return &Group_FieldTerminalPathArrayOfValues{Group_FieldTerminalPath: *fp, values: values.([]*Name)}
	case Group_FieldPathSelectorMetadata:
		return &Group_FieldTerminalPathArrayOfValues{Group_FieldTerminalPath: *fp, values: values.([]*meta.Meta)}
	case Group_FieldPathSelectorDisplayName:
		return &Group_FieldTerminalPathArrayOfValues{Group_FieldTerminalPath: *fp, values: values.([]string)}
	case Group_FieldPathSelectorDescription:
		return &Group_FieldTerminalPathArrayOfValues{Group_FieldTerminalPath: *fp, values: values.([]string)}
	case Group_FieldPathSelectorEmail:
		return &Group_FieldTerminalPathArrayOfValues{Group_FieldTerminalPath: *fp, values: values.([]string)}
	default:
		panic(fmt.Sprintf("Invalid selector for Group: %d", fp.selector))
	}
	return fpaov
}

func (fp *Group_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *Group_FieldTerminalPath) WithIArrayItemValue(value interface{}) Group_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for Group: %d", fp.selector))
	}
}

func (fp *Group_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

type Group_FieldSubPath struct {
	selector Group_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ Group_FieldPath = (*Group_FieldSubPath)(nil)

func (fps *Group_FieldSubPath) Selector() Group_FieldPathSelector {
	return fps.selector
}
func (fps *Group_FieldSubPath) AsMetadataSubPath() (meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(meta.Meta_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *Group_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *Group_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source Group
func (fps *Group_FieldSubPath) Get(source *Group) (values []interface{}) {
	switch fps.selector {
	case Group_FieldPathSelectorMetadata:
		values = append(values, fps.subPath.GetRaw(source.GetMetadata())...)
	default:
		panic(fmt.Sprintf("Invalid selector for Group: %d", fps.selector))
	}
	return
}

func (fps *Group_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*Group))
}

// GetSingle returns value of selected field from source Group
func (fps *Group_FieldSubPath) GetSingle(source *Group) (interface{}, bool) {
	switch fps.selector {
	case Group_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Group: %d", fps.selector))
	}
}

func (fps *Group_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*Group))
}

// GetDefault returns a default value of the field type
func (fps *Group_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *Group_FieldSubPath) ClearValue(item *Group) {
	if item != nil {
		switch fps.selector {
		case Group_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		default:
			panic(fmt.Sprintf("Invalid selector for Group: %d", fps.selector))
		}
	}
}

func (fps *Group_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*Group))
}

// IsLeaf - whether field path is holds simple value
func (fps *Group_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *Group_FieldSubPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	iPaths := []gotenobject.FieldPath{&Group_FieldTerminalPath{selector: fps.selector}}
	iPaths = append(iPaths, fps.subPath.SplitIntoTerminalIPaths()...)
	return iPaths
}

func (fps *Group_FieldSubPath) WithIValue(value interface{}) Group_FieldPathValue {
	return &Group_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *Group_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *Group_FieldSubPath) WithIArrayOfValues(values interface{}) Group_FieldPathArrayOfValues {
	return &Group_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *Group_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *Group_FieldSubPath) WithIArrayItemValue(value interface{}) Group_FieldPathArrayItemValue {
	return &Group_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *Group_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// Group_FieldPathValue allows storing values for Group fields according to their type
type Group_FieldPathValue interface {
	Group_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **Group)
	CompareWith(*Group) (cmp int, comparable bool)
}

func ParseGroup_FieldPathValue(pathStr, valueStr string) (Group_FieldPathValue, error) {
	fp, err := ParseGroup_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Group field path value from %s: %v", valueStr, err)
	}
	return fpv.(Group_FieldPathValue), nil
}

func MustParseGroup_FieldPathValue(pathStr, valueStr string) Group_FieldPathValue {
	fpv, err := ParseGroup_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type Group_FieldTerminalPathValue struct {
	Group_FieldTerminalPath
	value interface{}
}

var _ Group_FieldPathValue = (*Group_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'Group' as interface{}
func (fpv *Group_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *Group_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *Group_FieldTerminalPathValue) AsMetadataValue() (*meta.Meta, bool) {
	res, ok := fpv.value.(*meta.Meta)
	return res, ok
}
func (fpv *Group_FieldTerminalPathValue) AsDisplayNameValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *Group_FieldTerminalPathValue) AsDescriptionValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *Group_FieldTerminalPathValue) AsEmailValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}

// SetTo stores value for selected field for object Group
func (fpv *Group_FieldTerminalPathValue) SetTo(target **Group) {
	if *target == nil {
		*target = new(Group)
	}
	switch fpv.selector {
	case Group_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case Group_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*meta.Meta)
	case Group_FieldPathSelectorDisplayName:
		(*target).DisplayName = fpv.value.(string)
	case Group_FieldPathSelectorDescription:
		(*target).Description = fpv.value.(string)
	case Group_FieldPathSelectorEmail:
		(*target).Email = fpv.value.(string)
	default:
		panic(fmt.Sprintf("Invalid selector for Group: %d", fpv.selector))
	}
}

func (fpv *Group_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Group)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'Group_FieldTerminalPathValue' with the value under path in 'Group'.
func (fpv *Group_FieldTerminalPathValue) CompareWith(source *Group) (int, bool) {
	switch fpv.selector {
	case Group_FieldPathSelectorName:
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
	case Group_FieldPathSelectorMetadata:
		return 0, false
	case Group_FieldPathSelectorDisplayName:
		leftValue := fpv.value.(string)
		rightValue := source.GetDisplayName()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case Group_FieldPathSelectorDescription:
		leftValue := fpv.value.(string)
		rightValue := source.GetDescription()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case Group_FieldPathSelectorEmail:
		leftValue := fpv.value.(string)
		rightValue := source.GetEmail()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	default:
		panic(fmt.Sprintf("Invalid selector for Group: %d", fpv.selector))
	}
}

func (fpv *Group_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*Group))
}

type Group_FieldSubPathValue struct {
	Group_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ Group_FieldPathValue = (*Group_FieldSubPathValue)(nil)

func (fpvs *Group_FieldSubPathValue) AsMetadataPathValue() (meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(meta.Meta_FieldPathValue)
	return res, ok
}

func (fpvs *Group_FieldSubPathValue) SetTo(target **Group) {
	if *target == nil {
		*target = new(Group)
	}
	switch fpvs.Selector() {
	case Group_FieldPathSelectorMetadata:
		fpvs.subPathValue.(meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	default:
		panic(fmt.Sprintf("Invalid selector for Group: %d", fpvs.Selector()))
	}
}

func (fpvs *Group_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Group)
	fpvs.SetTo(&typedObject)
}

func (fpvs *Group_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *Group_FieldSubPathValue) CompareWith(source *Group) (int, bool) {
	switch fpvs.Selector() {
	case Group_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Group: %d", fpvs.Selector()))
	}
}

func (fpvs *Group_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*Group))
}

// Group_FieldPathArrayItemValue allows storing single item in Path-specific values for Group according to their type
// Present only for array (repeated) types.
type Group_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	Group_FieldPath
	ContainsValue(*Group) bool
}

// ParseGroup_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseGroup_FieldPathArrayItemValue(pathStr, valueStr string) (Group_FieldPathArrayItemValue, error) {
	fp, err := ParseGroup_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Group field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(Group_FieldPathArrayItemValue), nil
}

func MustParseGroup_FieldPathArrayItemValue(pathStr, valueStr string) Group_FieldPathArrayItemValue {
	fpaiv, err := ParseGroup_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type Group_FieldTerminalPathArrayItemValue struct {
	Group_FieldTerminalPath
	value interface{}
}

var _ Group_FieldPathArrayItemValue = (*Group_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object Group as interface{}
func (fpaiv *Group_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *Group_FieldTerminalPathArrayItemValue) GetSingle(source *Group) (interface{}, bool) {
	return nil, false
}

func (fpaiv *Group_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*Group))
}

// Contains returns a boolean indicating if value that is being held is present in given 'Group'
func (fpaiv *Group_FieldTerminalPathArrayItemValue) ContainsValue(source *Group) bool {
	slice := fpaiv.Group_FieldTerminalPath.Get(source)
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

type Group_FieldSubPathArrayItemValue struct {
	Group_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *Group_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *Group_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'Group'
func (fpaivs *Group_FieldSubPathArrayItemValue) ContainsValue(source *Group) bool {
	switch fpaivs.Selector() {
	case Group_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Group: %d", fpaivs.Selector()))
	}
}

// Group_FieldPathArrayOfValues allows storing slice of values for Group fields according to their type
type Group_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	Group_FieldPath
}

func ParseGroup_FieldPathArrayOfValues(pathStr, valuesStr string) (Group_FieldPathArrayOfValues, error) {
	fp, err := ParseGroup_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Group field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(Group_FieldPathArrayOfValues), nil
}

func MustParseGroup_FieldPathArrayOfValues(pathStr, valuesStr string) Group_FieldPathArrayOfValues {
	fpaov, err := ParseGroup_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type Group_FieldTerminalPathArrayOfValues struct {
	Group_FieldTerminalPath
	values interface{}
}

var _ Group_FieldPathArrayOfValues = (*Group_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *Group_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case Group_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case Group_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*meta.Meta) {
			values = append(values, v)
		}
	case Group_FieldPathSelectorDisplayName:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case Group_FieldPathSelectorDescription:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case Group_FieldPathSelectorEmail:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *Group_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *Group_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*meta.Meta, bool) {
	res, ok := fpaov.values.([]*meta.Meta)
	return res, ok
}
func (fpaov *Group_FieldTerminalPathArrayOfValues) AsDisplayNameArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *Group_FieldTerminalPathArrayOfValues) AsDescriptionArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *Group_FieldTerminalPathArrayOfValues) AsEmailArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}

type Group_FieldSubPathArrayOfValues struct {
	Group_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ Group_FieldPathArrayOfValues = (*Group_FieldSubPathArrayOfValues)(nil)

func (fpsaov *Group_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *Group_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
