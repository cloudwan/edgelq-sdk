// Code generated by protoc-gen-goten-object
// File: edgelq/monitoring/proto/v3/project.proto
// DO NOT EDIT!!!

package project

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
	multi_region_policy "github.com/cloudwan/goten-sdk/types/multi_region_policy"
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
	_ = &multi_region_policy.MultiRegionPolicy{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type Project_FieldPath interface {
	gotenobject.FieldPath
	Selector() Project_FieldPathSelector
	Get(source *Project) []interface{}
	GetSingle(source *Project) (interface{}, bool)
	ClearValue(item *Project)

	// Those methods build corresponding Project_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) Project_FieldPathValue
	WithIArrayOfValues(values interface{}) Project_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) Project_FieldPathArrayItemValue
}

type Project_FieldPathSelector int32

const (
	Project_FieldPathSelectorName              Project_FieldPathSelector = 0
	Project_FieldPathSelectorTitle             Project_FieldPathSelector = 1
	Project_FieldPathSelectorMetadata          Project_FieldPathSelector = 2
	Project_FieldPathSelectorMultiRegionPolicy Project_FieldPathSelector = 3
)

func (s Project_FieldPathSelector) String() string {
	switch s {
	case Project_FieldPathSelectorName:
		return "name"
	case Project_FieldPathSelectorTitle:
		return "title"
	case Project_FieldPathSelectorMetadata:
		return "metadata"
	case Project_FieldPathSelectorMultiRegionPolicy:
		return "multi_region_policy"
	default:
		panic(fmt.Sprintf("Invalid selector for Project: %d", s))
	}
}

func BuildProject_FieldPath(fp gotenobject.RawFieldPath) (Project_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object Project")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &Project_FieldTerminalPath{selector: Project_FieldPathSelectorName}, nil
		case "title":
			return &Project_FieldTerminalPath{selector: Project_FieldPathSelectorTitle}, nil
		case "metadata":
			return &Project_FieldTerminalPath{selector: Project_FieldPathSelectorMetadata}, nil
		case "multi_region_policy", "multiRegionPolicy", "multi-region-policy":
			return &Project_FieldTerminalPath{selector: Project_FieldPathSelectorMultiRegionPolicy}, nil
		}
	} else {
		switch fp[0] {
		case "metadata":
			if subpath, err := meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &Project_FieldSubPath{selector: Project_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		case "multi_region_policy", "multiRegionPolicy", "multi-region-policy":
			if subpath, err := multi_region_policy.BuildMultiRegionPolicy_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &Project_FieldSubPath{selector: Project_FieldPathSelectorMultiRegionPolicy, subPath: subpath}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object Project", fp)
}

func ParseProject_FieldPath(rawField string) (Project_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildProject_FieldPath(fp)
}

func MustParseProject_FieldPath(rawField string) Project_FieldPath {
	fp, err := ParseProject_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type Project_FieldTerminalPath struct {
	selector Project_FieldPathSelector
}

var _ Project_FieldPath = (*Project_FieldTerminalPath)(nil)

func (fp *Project_FieldTerminalPath) Selector() Project_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *Project_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *Project_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source Project
func (fp *Project_FieldTerminalPath) Get(source *Project) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case Project_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case Project_FieldPathSelectorTitle:
			values = append(values, source.Title)
		case Project_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		case Project_FieldPathSelectorMultiRegionPolicy:
			if source.MultiRegionPolicy != nil {
				values = append(values, source.MultiRegionPolicy)
			}
		default:
			panic(fmt.Sprintf("Invalid selector for Project: %d", fp.selector))
		}
	}
	return
}

func (fp *Project_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*Project))
}

// GetSingle returns value pointed by specific field of from source Project
func (fp *Project_FieldTerminalPath) GetSingle(source *Project) (interface{}, bool) {
	switch fp.selector {
	case Project_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case Project_FieldPathSelectorTitle:
		return source.GetTitle(), source != nil
	case Project_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	case Project_FieldPathSelectorMultiRegionPolicy:
		res := source.GetMultiRegionPolicy()
		return res, res != nil
	default:
		panic(fmt.Sprintf("Invalid selector for Project: %d", fp.selector))
	}
}

func (fp *Project_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*Project))
}

// GetDefault returns a default value of the field type
func (fp *Project_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case Project_FieldPathSelectorName:
		return (*Name)(nil)
	case Project_FieldPathSelectorTitle:
		return ""
	case Project_FieldPathSelectorMetadata:
		return (*meta.Meta)(nil)
	case Project_FieldPathSelectorMultiRegionPolicy:
		return (*multi_region_policy.MultiRegionPolicy)(nil)
	default:
		panic(fmt.Sprintf("Invalid selector for Project: %d", fp.selector))
	}
}

func (fp *Project_FieldTerminalPath) ClearValue(item *Project) {
	if item != nil {
		switch fp.selector {
		case Project_FieldPathSelectorName:
			item.Name = nil
		case Project_FieldPathSelectorTitle:
			item.Title = ""
		case Project_FieldPathSelectorMetadata:
			item.Metadata = nil
		case Project_FieldPathSelectorMultiRegionPolicy:
			item.MultiRegionPolicy = nil
		default:
			panic(fmt.Sprintf("Invalid selector for Project: %d", fp.selector))
		}
	}
}

func (fp *Project_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*Project))
}

// IsLeaf - whether field path is holds simple value
func (fp *Project_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == Project_FieldPathSelectorName ||
		fp.selector == Project_FieldPathSelectorTitle
}

func (fp *Project_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *Project_FieldTerminalPath) WithIValue(value interface{}) Project_FieldPathValue {
	switch fp.selector {
	case Project_FieldPathSelectorName:
		return &Project_FieldTerminalPathValue{Project_FieldTerminalPath: *fp, value: value.(*Name)}
	case Project_FieldPathSelectorTitle:
		return &Project_FieldTerminalPathValue{Project_FieldTerminalPath: *fp, value: value.(string)}
	case Project_FieldPathSelectorMetadata:
		return &Project_FieldTerminalPathValue{Project_FieldTerminalPath: *fp, value: value.(*meta.Meta)}
	case Project_FieldPathSelectorMultiRegionPolicy:
		return &Project_FieldTerminalPathValue{Project_FieldTerminalPath: *fp, value: value.(*multi_region_policy.MultiRegionPolicy)}
	default:
		panic(fmt.Sprintf("Invalid selector for Project: %d", fp.selector))
	}
}

func (fp *Project_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *Project_FieldTerminalPath) WithIArrayOfValues(values interface{}) Project_FieldPathArrayOfValues {
	fpaov := &Project_FieldTerminalPathArrayOfValues{Project_FieldTerminalPath: *fp}
	switch fp.selector {
	case Project_FieldPathSelectorName:
		return &Project_FieldTerminalPathArrayOfValues{Project_FieldTerminalPath: *fp, values: values.([]*Name)}
	case Project_FieldPathSelectorTitle:
		return &Project_FieldTerminalPathArrayOfValues{Project_FieldTerminalPath: *fp, values: values.([]string)}
	case Project_FieldPathSelectorMetadata:
		return &Project_FieldTerminalPathArrayOfValues{Project_FieldTerminalPath: *fp, values: values.([]*meta.Meta)}
	case Project_FieldPathSelectorMultiRegionPolicy:
		return &Project_FieldTerminalPathArrayOfValues{Project_FieldTerminalPath: *fp, values: values.([]*multi_region_policy.MultiRegionPolicy)}
	default:
		panic(fmt.Sprintf("Invalid selector for Project: %d", fp.selector))
	}
	return fpaov
}

func (fp *Project_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *Project_FieldTerminalPath) WithIArrayItemValue(value interface{}) Project_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for Project: %d", fp.selector))
	}
}

func (fp *Project_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

type Project_FieldSubPath struct {
	selector Project_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ Project_FieldPath = (*Project_FieldSubPath)(nil)

func (fps *Project_FieldSubPath) Selector() Project_FieldPathSelector {
	return fps.selector
}
func (fps *Project_FieldSubPath) AsMetadataSubPath() (meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(meta.Meta_FieldPath)
	return res, ok
}
func (fps *Project_FieldSubPath) AsMultiRegionPolicySubPath() (multi_region_policy.MultiRegionPolicy_FieldPath, bool) {
	res, ok := fps.subPath.(multi_region_policy.MultiRegionPolicy_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *Project_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *Project_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source Project
func (fps *Project_FieldSubPath) Get(source *Project) (values []interface{}) {
	switch fps.selector {
	case Project_FieldPathSelectorMetadata:
		values = append(values, fps.subPath.GetRaw(source.GetMetadata())...)
	case Project_FieldPathSelectorMultiRegionPolicy:
		values = append(values, fps.subPath.GetRaw(source.GetMultiRegionPolicy())...)
	default:
		panic(fmt.Sprintf("Invalid selector for Project: %d", fps.selector))
	}
	return
}

func (fps *Project_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*Project))
}

// GetSingle returns value of selected field from source Project
func (fps *Project_FieldSubPath) GetSingle(source *Project) (interface{}, bool) {
	switch fps.selector {
	case Project_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	case Project_FieldPathSelectorMultiRegionPolicy:
		if source.GetMultiRegionPolicy() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMultiRegionPolicy())
	default:
		panic(fmt.Sprintf("Invalid selector for Project: %d", fps.selector))
	}
}

func (fps *Project_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*Project))
}

// GetDefault returns a default value of the field type
func (fps *Project_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *Project_FieldSubPath) ClearValue(item *Project) {
	if item != nil {
		switch fps.selector {
		case Project_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		case Project_FieldPathSelectorMultiRegionPolicy:
			fps.subPath.ClearValueRaw(item.MultiRegionPolicy)
		default:
			panic(fmt.Sprintf("Invalid selector for Project: %d", fps.selector))
		}
	}
}

func (fps *Project_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*Project))
}

// IsLeaf - whether field path is holds simple value
func (fps *Project_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *Project_FieldSubPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	iPaths := []gotenobject.FieldPath{&Project_FieldTerminalPath{selector: fps.selector}}
	iPaths = append(iPaths, fps.subPath.SplitIntoTerminalIPaths()...)
	return iPaths
}

func (fps *Project_FieldSubPath) WithIValue(value interface{}) Project_FieldPathValue {
	return &Project_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *Project_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *Project_FieldSubPath) WithIArrayOfValues(values interface{}) Project_FieldPathArrayOfValues {
	return &Project_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *Project_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *Project_FieldSubPath) WithIArrayItemValue(value interface{}) Project_FieldPathArrayItemValue {
	return &Project_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *Project_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// Project_FieldPathValue allows storing values for Project fields according to their type
type Project_FieldPathValue interface {
	Project_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **Project)
	CompareWith(*Project) (cmp int, comparable bool)
}

func ParseProject_FieldPathValue(pathStr, valueStr string) (Project_FieldPathValue, error) {
	fp, err := ParseProject_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Project field path value from %s: %v", valueStr, err)
	}
	return fpv.(Project_FieldPathValue), nil
}

func MustParseProject_FieldPathValue(pathStr, valueStr string) Project_FieldPathValue {
	fpv, err := ParseProject_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type Project_FieldTerminalPathValue struct {
	Project_FieldTerminalPath
	value interface{}
}

var _ Project_FieldPathValue = (*Project_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'Project' as interface{}
func (fpv *Project_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *Project_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *Project_FieldTerminalPathValue) AsTitleValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *Project_FieldTerminalPathValue) AsMetadataValue() (*meta.Meta, bool) {
	res, ok := fpv.value.(*meta.Meta)
	return res, ok
}
func (fpv *Project_FieldTerminalPathValue) AsMultiRegionPolicyValue() (*multi_region_policy.MultiRegionPolicy, bool) {
	res, ok := fpv.value.(*multi_region_policy.MultiRegionPolicy)
	return res, ok
}

// SetTo stores value for selected field for object Project
func (fpv *Project_FieldTerminalPathValue) SetTo(target **Project) {
	if *target == nil {
		*target = new(Project)
	}
	switch fpv.selector {
	case Project_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case Project_FieldPathSelectorTitle:
		(*target).Title = fpv.value.(string)
	case Project_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*meta.Meta)
	case Project_FieldPathSelectorMultiRegionPolicy:
		(*target).MultiRegionPolicy = fpv.value.(*multi_region_policy.MultiRegionPolicy)
	default:
		panic(fmt.Sprintf("Invalid selector for Project: %d", fpv.selector))
	}
}

func (fpv *Project_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Project)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'Project_FieldTerminalPathValue' with the value under path in 'Project'.
func (fpv *Project_FieldTerminalPathValue) CompareWith(source *Project) (int, bool) {
	switch fpv.selector {
	case Project_FieldPathSelectorName:
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
	case Project_FieldPathSelectorTitle:
		leftValue := fpv.value.(string)
		rightValue := source.GetTitle()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case Project_FieldPathSelectorMetadata:
		return 0, false
	case Project_FieldPathSelectorMultiRegionPolicy:
		return 0, false
	default:
		panic(fmt.Sprintf("Invalid selector for Project: %d", fpv.selector))
	}
}

func (fpv *Project_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*Project))
}

type Project_FieldSubPathValue struct {
	Project_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ Project_FieldPathValue = (*Project_FieldSubPathValue)(nil)

func (fpvs *Project_FieldSubPathValue) AsMetadataPathValue() (meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(meta.Meta_FieldPathValue)
	return res, ok
}
func (fpvs *Project_FieldSubPathValue) AsMultiRegionPolicyPathValue() (multi_region_policy.MultiRegionPolicy_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(multi_region_policy.MultiRegionPolicy_FieldPathValue)
	return res, ok
}

func (fpvs *Project_FieldSubPathValue) SetTo(target **Project) {
	if *target == nil {
		*target = new(Project)
	}
	switch fpvs.Selector() {
	case Project_FieldPathSelectorMetadata:
		fpvs.subPathValue.(meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	case Project_FieldPathSelectorMultiRegionPolicy:
		fpvs.subPathValue.(multi_region_policy.MultiRegionPolicy_FieldPathValue).SetTo(&(*target).MultiRegionPolicy)
	default:
		panic(fmt.Sprintf("Invalid selector for Project: %d", fpvs.Selector()))
	}
}

func (fpvs *Project_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Project)
	fpvs.SetTo(&typedObject)
}

func (fpvs *Project_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *Project_FieldSubPathValue) CompareWith(source *Project) (int, bool) {
	switch fpvs.Selector() {
	case Project_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	case Project_FieldPathSelectorMultiRegionPolicy:
		return fpvs.subPathValue.(multi_region_policy.MultiRegionPolicy_FieldPathValue).CompareWith(source.GetMultiRegionPolicy())
	default:
		panic(fmt.Sprintf("Invalid selector for Project: %d", fpvs.Selector()))
	}
}

func (fpvs *Project_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*Project))
}

// Project_FieldPathArrayItemValue allows storing single item in Path-specific values for Project according to their type
// Present only for array (repeated) types.
type Project_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	Project_FieldPath
	ContainsValue(*Project) bool
}

// ParseProject_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseProject_FieldPathArrayItemValue(pathStr, valueStr string) (Project_FieldPathArrayItemValue, error) {
	fp, err := ParseProject_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Project field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(Project_FieldPathArrayItemValue), nil
}

func MustParseProject_FieldPathArrayItemValue(pathStr, valueStr string) Project_FieldPathArrayItemValue {
	fpaiv, err := ParseProject_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type Project_FieldTerminalPathArrayItemValue struct {
	Project_FieldTerminalPath
	value interface{}
}

var _ Project_FieldPathArrayItemValue = (*Project_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object Project as interface{}
func (fpaiv *Project_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *Project_FieldTerminalPathArrayItemValue) GetSingle(source *Project) (interface{}, bool) {
	return nil, false
}

func (fpaiv *Project_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*Project))
}

// Contains returns a boolean indicating if value that is being held is present in given 'Project'
func (fpaiv *Project_FieldTerminalPathArrayItemValue) ContainsValue(source *Project) bool {
	slice := fpaiv.Project_FieldTerminalPath.Get(source)
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

type Project_FieldSubPathArrayItemValue struct {
	Project_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *Project_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *Project_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue)
	return res, ok
}
func (fpaivs *Project_FieldSubPathArrayItemValue) AsMultiRegionPolicyPathItemValue() (multi_region_policy.MultiRegionPolicy_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(multi_region_policy.MultiRegionPolicy_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'Project'
func (fpaivs *Project_FieldSubPathArrayItemValue) ContainsValue(source *Project) bool {
	switch fpaivs.Selector() {
	case Project_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	case Project_FieldPathSelectorMultiRegionPolicy:
		return fpaivs.subPathItemValue.(multi_region_policy.MultiRegionPolicy_FieldPathArrayItemValue).ContainsValue(source.GetMultiRegionPolicy())
	default:
		panic(fmt.Sprintf("Invalid selector for Project: %d", fpaivs.Selector()))
	}
}

// Project_FieldPathArrayOfValues allows storing slice of values for Project fields according to their type
type Project_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	Project_FieldPath
}

func ParseProject_FieldPathArrayOfValues(pathStr, valuesStr string) (Project_FieldPathArrayOfValues, error) {
	fp, err := ParseProject_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Project field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(Project_FieldPathArrayOfValues), nil
}

func MustParseProject_FieldPathArrayOfValues(pathStr, valuesStr string) Project_FieldPathArrayOfValues {
	fpaov, err := ParseProject_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type Project_FieldTerminalPathArrayOfValues struct {
	Project_FieldTerminalPath
	values interface{}
}

var _ Project_FieldPathArrayOfValues = (*Project_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *Project_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case Project_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case Project_FieldPathSelectorTitle:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case Project_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*meta.Meta) {
			values = append(values, v)
		}
	case Project_FieldPathSelectorMultiRegionPolicy:
		for _, v := range fpaov.values.([]*multi_region_policy.MultiRegionPolicy) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *Project_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *Project_FieldTerminalPathArrayOfValues) AsTitleArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *Project_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*meta.Meta, bool) {
	res, ok := fpaov.values.([]*meta.Meta)
	return res, ok
}
func (fpaov *Project_FieldTerminalPathArrayOfValues) AsMultiRegionPolicyArrayOfValues() ([]*multi_region_policy.MultiRegionPolicy, bool) {
	res, ok := fpaov.values.([]*multi_region_policy.MultiRegionPolicy)
	return res, ok
}

type Project_FieldSubPathArrayOfValues struct {
	Project_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ Project_FieldPathArrayOfValues = (*Project_FieldSubPathArrayOfValues)(nil)

func (fpsaov *Project_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *Project_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
func (fpsaov *Project_FieldSubPathArrayOfValues) AsMultiRegionPolicyPathArrayOfValues() (multi_region_policy.MultiRegionPolicy_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(multi_region_policy.MultiRegionPolicy_FieldPathArrayOfValues)
	return res, ok
}
