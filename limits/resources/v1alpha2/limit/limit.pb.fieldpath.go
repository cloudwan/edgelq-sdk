// Code generated by protoc-gen-goten-object
// File: edgelq/limits/proto/v1alpha2/limit.proto
// DO NOT EDIT!!!

package limit

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
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	limit_pool "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/limit_pool"
	meta_resource "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/resource"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
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
	_ = &iam_project.Project{}
	_ = &limit_pool.LimitPool{}
	_ = &meta_resource.Resource{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type Limit_FieldPath interface {
	gotenobject.FieldPath
	Selector() Limit_FieldPathSelector
	Get(source *Limit) []interface{}
	GetSingle(source *Limit) (interface{}, bool)
	ClearValue(item *Limit)

	// Those methods build corresponding Limit_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) Limit_FieldPathValue
	WithIArrayOfValues(values interface{}) Limit_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) Limit_FieldPathArrayItemValue
}

type Limit_FieldPathSelector int32

const (
	Limit_FieldPathSelectorName            Limit_FieldPathSelector = 0
	Limit_FieldPathSelectorService         Limit_FieldPathSelector = 1
	Limit_FieldPathSelectorResource        Limit_FieldPathSelector = 2
	Limit_FieldPathSelectorRegion          Limit_FieldPathSelector = 3
	Limit_FieldPathSelectorConfiguredLimit Limit_FieldPathSelector = 4
	Limit_FieldPathSelectorActiveLimit     Limit_FieldPathSelector = 5
	Limit_FieldPathSelectorUsage           Limit_FieldPathSelector = 6
	Limit_FieldPathSelectorSource          Limit_FieldPathSelector = 7
	Limit_FieldPathSelectorMetadata        Limit_FieldPathSelector = 8
)

func (s Limit_FieldPathSelector) String() string {
	switch s {
	case Limit_FieldPathSelectorName:
		return "name"
	case Limit_FieldPathSelectorService:
		return "service"
	case Limit_FieldPathSelectorResource:
		return "resource"
	case Limit_FieldPathSelectorRegion:
		return "region"
	case Limit_FieldPathSelectorConfiguredLimit:
		return "configured_limit"
	case Limit_FieldPathSelectorActiveLimit:
		return "active_limit"
	case Limit_FieldPathSelectorUsage:
		return "usage"
	case Limit_FieldPathSelectorSource:
		return "source"
	case Limit_FieldPathSelectorMetadata:
		return "metadata"
	default:
		panic(fmt.Sprintf("Invalid selector for Limit: %d", s))
	}
}

func BuildLimit_FieldPath(fp gotenobject.RawFieldPath) (Limit_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object Limit")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorName}, nil
		case "service":
			return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorService}, nil
		case "resource":
			return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorResource}, nil
		case "region":
			return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorRegion}, nil
		case "configured_limit", "configuredLimit", "configured-limit":
			return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorConfiguredLimit}, nil
		case "active_limit", "activeLimit", "active-limit":
			return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorActiveLimit}, nil
		case "usage":
			return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorUsage}, nil
		case "source":
			return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorSource}, nil
		case "metadata":
			return &Limit_FieldTerminalPath{selector: Limit_FieldPathSelectorMetadata}, nil
		}
	} else {
		switch fp[0] {
		case "metadata":
			if subpath, err := meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &Limit_FieldSubPath{selector: Limit_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object Limit", fp)
}

func ParseLimit_FieldPath(rawField string) (Limit_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildLimit_FieldPath(fp)
}

func MustParseLimit_FieldPath(rawField string) Limit_FieldPath {
	fp, err := ParseLimit_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type Limit_FieldTerminalPath struct {
	selector Limit_FieldPathSelector
}

var _ Limit_FieldPath = (*Limit_FieldTerminalPath)(nil)

func (fp *Limit_FieldTerminalPath) Selector() Limit_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *Limit_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *Limit_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source Limit
func (fp *Limit_FieldTerminalPath) Get(source *Limit) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case Limit_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case Limit_FieldPathSelectorService:
			if source.Service != nil {
				values = append(values, source.Service)
			}
		case Limit_FieldPathSelectorResource:
			if source.Resource != nil {
				values = append(values, source.Resource)
			}
		case Limit_FieldPathSelectorRegion:
			values = append(values, source.Region)
		case Limit_FieldPathSelectorConfiguredLimit:
			values = append(values, source.ConfiguredLimit)
		case Limit_FieldPathSelectorActiveLimit:
			values = append(values, source.ActiveLimit)
		case Limit_FieldPathSelectorUsage:
			values = append(values, source.Usage)
		case Limit_FieldPathSelectorSource:
			if source.Source != nil {
				values = append(values, source.Source)
			}
		case Limit_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		default:
			panic(fmt.Sprintf("Invalid selector for Limit: %d", fp.selector))
		}
	}
	return
}

func (fp *Limit_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*Limit))
}

// GetSingle returns value pointed by specific field of from source Limit
func (fp *Limit_FieldTerminalPath) GetSingle(source *Limit) (interface{}, bool) {
	switch fp.selector {
	case Limit_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case Limit_FieldPathSelectorService:
		res := source.GetService()
		return res, res != nil
	case Limit_FieldPathSelectorResource:
		res := source.GetResource()
		return res, res != nil
	case Limit_FieldPathSelectorRegion:
		return source.GetRegion(), source != nil
	case Limit_FieldPathSelectorConfiguredLimit:
		return source.GetConfiguredLimit(), source != nil
	case Limit_FieldPathSelectorActiveLimit:
		return source.GetActiveLimit(), source != nil
	case Limit_FieldPathSelectorUsage:
		return source.GetUsage(), source != nil
	case Limit_FieldPathSelectorSource:
		res := source.GetSource()
		return res, res != nil
	case Limit_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	default:
		panic(fmt.Sprintf("Invalid selector for Limit: %d", fp.selector))
	}
}

func (fp *Limit_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*Limit))
}

// GetDefault returns a default value of the field type
func (fp *Limit_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case Limit_FieldPathSelectorName:
		return (*Name)(nil)
	case Limit_FieldPathSelectorService:
		return (*meta_service.Reference)(nil)
	case Limit_FieldPathSelectorResource:
		return (*meta_resource.Reference)(nil)
	case Limit_FieldPathSelectorRegion:
		return ""
	case Limit_FieldPathSelectorConfiguredLimit:
		return int64(0)
	case Limit_FieldPathSelectorActiveLimit:
		return int64(0)
	case Limit_FieldPathSelectorUsage:
		return int64(0)
	case Limit_FieldPathSelectorSource:
		return (*limit_pool.Reference)(nil)
	case Limit_FieldPathSelectorMetadata:
		return (*meta.Meta)(nil)
	default:
		panic(fmt.Sprintf("Invalid selector for Limit: %d", fp.selector))
	}
}

func (fp *Limit_FieldTerminalPath) ClearValue(item *Limit) {
	if item != nil {
		switch fp.selector {
		case Limit_FieldPathSelectorName:
			item.Name = nil
		case Limit_FieldPathSelectorService:
			item.Service = nil
		case Limit_FieldPathSelectorResource:
			item.Resource = nil
		case Limit_FieldPathSelectorRegion:
			item.Region = ""
		case Limit_FieldPathSelectorConfiguredLimit:
			item.ConfiguredLimit = int64(0)
		case Limit_FieldPathSelectorActiveLimit:
			item.ActiveLimit = int64(0)
		case Limit_FieldPathSelectorUsage:
			item.Usage = int64(0)
		case Limit_FieldPathSelectorSource:
			item.Source = nil
		case Limit_FieldPathSelectorMetadata:
			item.Metadata = nil
		default:
			panic(fmt.Sprintf("Invalid selector for Limit: %d", fp.selector))
		}
	}
}

func (fp *Limit_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*Limit))
}

// IsLeaf - whether field path is holds simple value
func (fp *Limit_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == Limit_FieldPathSelectorName ||
		fp.selector == Limit_FieldPathSelectorService ||
		fp.selector == Limit_FieldPathSelectorResource ||
		fp.selector == Limit_FieldPathSelectorRegion ||
		fp.selector == Limit_FieldPathSelectorConfiguredLimit ||
		fp.selector == Limit_FieldPathSelectorActiveLimit ||
		fp.selector == Limit_FieldPathSelectorUsage ||
		fp.selector == Limit_FieldPathSelectorSource
}

func (fp *Limit_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *Limit_FieldTerminalPath) WithIValue(value interface{}) Limit_FieldPathValue {
	switch fp.selector {
	case Limit_FieldPathSelectorName:
		return &Limit_FieldTerminalPathValue{Limit_FieldTerminalPath: *fp, value: value.(*Name)}
	case Limit_FieldPathSelectorService:
		return &Limit_FieldTerminalPathValue{Limit_FieldTerminalPath: *fp, value: value.(*meta_service.Reference)}
	case Limit_FieldPathSelectorResource:
		return &Limit_FieldTerminalPathValue{Limit_FieldTerminalPath: *fp, value: value.(*meta_resource.Reference)}
	case Limit_FieldPathSelectorRegion:
		return &Limit_FieldTerminalPathValue{Limit_FieldTerminalPath: *fp, value: value.(string)}
	case Limit_FieldPathSelectorConfiguredLimit:
		return &Limit_FieldTerminalPathValue{Limit_FieldTerminalPath: *fp, value: value.(int64)}
	case Limit_FieldPathSelectorActiveLimit:
		return &Limit_FieldTerminalPathValue{Limit_FieldTerminalPath: *fp, value: value.(int64)}
	case Limit_FieldPathSelectorUsage:
		return &Limit_FieldTerminalPathValue{Limit_FieldTerminalPath: *fp, value: value.(int64)}
	case Limit_FieldPathSelectorSource:
		return &Limit_FieldTerminalPathValue{Limit_FieldTerminalPath: *fp, value: value.(*limit_pool.Reference)}
	case Limit_FieldPathSelectorMetadata:
		return &Limit_FieldTerminalPathValue{Limit_FieldTerminalPath: *fp, value: value.(*meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for Limit: %d", fp.selector))
	}
}

func (fp *Limit_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *Limit_FieldTerminalPath) WithIArrayOfValues(values interface{}) Limit_FieldPathArrayOfValues {
	fpaov := &Limit_FieldTerminalPathArrayOfValues{Limit_FieldTerminalPath: *fp}
	switch fp.selector {
	case Limit_FieldPathSelectorName:
		return &Limit_FieldTerminalPathArrayOfValues{Limit_FieldTerminalPath: *fp, values: values.([]*Name)}
	case Limit_FieldPathSelectorService:
		return &Limit_FieldTerminalPathArrayOfValues{Limit_FieldTerminalPath: *fp, values: values.([]*meta_service.Reference)}
	case Limit_FieldPathSelectorResource:
		return &Limit_FieldTerminalPathArrayOfValues{Limit_FieldTerminalPath: *fp, values: values.([]*meta_resource.Reference)}
	case Limit_FieldPathSelectorRegion:
		return &Limit_FieldTerminalPathArrayOfValues{Limit_FieldTerminalPath: *fp, values: values.([]string)}
	case Limit_FieldPathSelectorConfiguredLimit:
		return &Limit_FieldTerminalPathArrayOfValues{Limit_FieldTerminalPath: *fp, values: values.([]int64)}
	case Limit_FieldPathSelectorActiveLimit:
		return &Limit_FieldTerminalPathArrayOfValues{Limit_FieldTerminalPath: *fp, values: values.([]int64)}
	case Limit_FieldPathSelectorUsage:
		return &Limit_FieldTerminalPathArrayOfValues{Limit_FieldTerminalPath: *fp, values: values.([]int64)}
	case Limit_FieldPathSelectorSource:
		return &Limit_FieldTerminalPathArrayOfValues{Limit_FieldTerminalPath: *fp, values: values.([]*limit_pool.Reference)}
	case Limit_FieldPathSelectorMetadata:
		return &Limit_FieldTerminalPathArrayOfValues{Limit_FieldTerminalPath: *fp, values: values.([]*meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for Limit: %d", fp.selector))
	}
	return fpaov
}

func (fp *Limit_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *Limit_FieldTerminalPath) WithIArrayItemValue(value interface{}) Limit_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for Limit: %d", fp.selector))
	}
}

func (fp *Limit_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

type Limit_FieldSubPath struct {
	selector Limit_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ Limit_FieldPath = (*Limit_FieldSubPath)(nil)

func (fps *Limit_FieldSubPath) Selector() Limit_FieldPathSelector {
	return fps.selector
}
func (fps *Limit_FieldSubPath) AsMetadataSubPath() (meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(meta.Meta_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *Limit_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *Limit_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source Limit
func (fps *Limit_FieldSubPath) Get(source *Limit) (values []interface{}) {
	switch fps.selector {
	case Limit_FieldPathSelectorMetadata:
		values = append(values, fps.subPath.GetRaw(source.GetMetadata())...)
	default:
		panic(fmt.Sprintf("Invalid selector for Limit: %d", fps.selector))
	}
	return
}

func (fps *Limit_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*Limit))
}

// GetSingle returns value of selected field from source Limit
func (fps *Limit_FieldSubPath) GetSingle(source *Limit) (interface{}, bool) {
	switch fps.selector {
	case Limit_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Limit: %d", fps.selector))
	}
}

func (fps *Limit_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*Limit))
}

// GetDefault returns a default value of the field type
func (fps *Limit_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *Limit_FieldSubPath) ClearValue(item *Limit) {
	if item != nil {
		switch fps.selector {
		case Limit_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		default:
			panic(fmt.Sprintf("Invalid selector for Limit: %d", fps.selector))
		}
	}
}

func (fps *Limit_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*Limit))
}

// IsLeaf - whether field path is holds simple value
func (fps *Limit_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *Limit_FieldSubPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	iPaths := []gotenobject.FieldPath{&Limit_FieldTerminalPath{selector: fps.selector}}
	iPaths = append(iPaths, fps.subPath.SplitIntoTerminalIPaths()...)
	return iPaths
}

func (fps *Limit_FieldSubPath) WithIValue(value interface{}) Limit_FieldPathValue {
	return &Limit_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *Limit_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *Limit_FieldSubPath) WithIArrayOfValues(values interface{}) Limit_FieldPathArrayOfValues {
	return &Limit_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *Limit_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *Limit_FieldSubPath) WithIArrayItemValue(value interface{}) Limit_FieldPathArrayItemValue {
	return &Limit_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *Limit_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// Limit_FieldPathValue allows storing values for Limit fields according to their type
type Limit_FieldPathValue interface {
	Limit_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **Limit)
	CompareWith(*Limit) (cmp int, comparable bool)
}

func ParseLimit_FieldPathValue(pathStr, valueStr string) (Limit_FieldPathValue, error) {
	fp, err := ParseLimit_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Limit field path value from %s: %v", valueStr, err)
	}
	return fpv.(Limit_FieldPathValue), nil
}

func MustParseLimit_FieldPathValue(pathStr, valueStr string) Limit_FieldPathValue {
	fpv, err := ParseLimit_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type Limit_FieldTerminalPathValue struct {
	Limit_FieldTerminalPath
	value interface{}
}

var _ Limit_FieldPathValue = (*Limit_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'Limit' as interface{}
func (fpv *Limit_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *Limit_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *Limit_FieldTerminalPathValue) AsServiceValue() (*meta_service.Reference, bool) {
	res, ok := fpv.value.(*meta_service.Reference)
	return res, ok
}
func (fpv *Limit_FieldTerminalPathValue) AsResourceValue() (*meta_resource.Reference, bool) {
	res, ok := fpv.value.(*meta_resource.Reference)
	return res, ok
}
func (fpv *Limit_FieldTerminalPathValue) AsRegionValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *Limit_FieldTerminalPathValue) AsConfiguredLimitValue() (int64, bool) {
	res, ok := fpv.value.(int64)
	return res, ok
}
func (fpv *Limit_FieldTerminalPathValue) AsActiveLimitValue() (int64, bool) {
	res, ok := fpv.value.(int64)
	return res, ok
}
func (fpv *Limit_FieldTerminalPathValue) AsUsageValue() (int64, bool) {
	res, ok := fpv.value.(int64)
	return res, ok
}
func (fpv *Limit_FieldTerminalPathValue) AsSourceValue() (*limit_pool.Reference, bool) {
	res, ok := fpv.value.(*limit_pool.Reference)
	return res, ok
}
func (fpv *Limit_FieldTerminalPathValue) AsMetadataValue() (*meta.Meta, bool) {
	res, ok := fpv.value.(*meta.Meta)
	return res, ok
}

// SetTo stores value for selected field for object Limit
func (fpv *Limit_FieldTerminalPathValue) SetTo(target **Limit) {
	if *target == nil {
		*target = new(Limit)
	}
	switch fpv.selector {
	case Limit_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case Limit_FieldPathSelectorService:
		(*target).Service = fpv.value.(*meta_service.Reference)
	case Limit_FieldPathSelectorResource:
		(*target).Resource = fpv.value.(*meta_resource.Reference)
	case Limit_FieldPathSelectorRegion:
		(*target).Region = fpv.value.(string)
	case Limit_FieldPathSelectorConfiguredLimit:
		(*target).ConfiguredLimit = fpv.value.(int64)
	case Limit_FieldPathSelectorActiveLimit:
		(*target).ActiveLimit = fpv.value.(int64)
	case Limit_FieldPathSelectorUsage:
		(*target).Usage = fpv.value.(int64)
	case Limit_FieldPathSelectorSource:
		(*target).Source = fpv.value.(*limit_pool.Reference)
	case Limit_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*meta.Meta)
	default:
		panic(fmt.Sprintf("Invalid selector for Limit: %d", fpv.selector))
	}
}

func (fpv *Limit_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Limit)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'Limit_FieldTerminalPathValue' with the value under path in 'Limit'.
func (fpv *Limit_FieldTerminalPathValue) CompareWith(source *Limit) (int, bool) {
	switch fpv.selector {
	case Limit_FieldPathSelectorName:
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
	case Limit_FieldPathSelectorService:
		leftValue := fpv.value.(*meta_service.Reference)
		rightValue := source.GetService()
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
	case Limit_FieldPathSelectorResource:
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
	case Limit_FieldPathSelectorRegion:
		leftValue := fpv.value.(string)
		rightValue := source.GetRegion()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case Limit_FieldPathSelectorConfiguredLimit:
		leftValue := fpv.value.(int64)
		rightValue := source.GetConfiguredLimit()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case Limit_FieldPathSelectorActiveLimit:
		leftValue := fpv.value.(int64)
		rightValue := source.GetActiveLimit()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case Limit_FieldPathSelectorUsage:
		leftValue := fpv.value.(int64)
		rightValue := source.GetUsage()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case Limit_FieldPathSelectorSource:
		leftValue := fpv.value.(*limit_pool.Reference)
		rightValue := source.GetSource()
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
	case Limit_FieldPathSelectorMetadata:
		return 0, false
	default:
		panic(fmt.Sprintf("Invalid selector for Limit: %d", fpv.selector))
	}
}

func (fpv *Limit_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*Limit))
}

type Limit_FieldSubPathValue struct {
	Limit_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ Limit_FieldPathValue = (*Limit_FieldSubPathValue)(nil)

func (fpvs *Limit_FieldSubPathValue) AsMetadataPathValue() (meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(meta.Meta_FieldPathValue)
	return res, ok
}

func (fpvs *Limit_FieldSubPathValue) SetTo(target **Limit) {
	if *target == nil {
		*target = new(Limit)
	}
	switch fpvs.Selector() {
	case Limit_FieldPathSelectorMetadata:
		fpvs.subPathValue.(meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	default:
		panic(fmt.Sprintf("Invalid selector for Limit: %d", fpvs.Selector()))
	}
}

func (fpvs *Limit_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Limit)
	fpvs.SetTo(&typedObject)
}

func (fpvs *Limit_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *Limit_FieldSubPathValue) CompareWith(source *Limit) (int, bool) {
	switch fpvs.Selector() {
	case Limit_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Limit: %d", fpvs.Selector()))
	}
}

func (fpvs *Limit_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*Limit))
}

// Limit_FieldPathArrayItemValue allows storing single item in Path-specific values for Limit according to their type
// Present only for array (repeated) types.
type Limit_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	Limit_FieldPath
	ContainsValue(*Limit) bool
}

// ParseLimit_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseLimit_FieldPathArrayItemValue(pathStr, valueStr string) (Limit_FieldPathArrayItemValue, error) {
	fp, err := ParseLimit_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Limit field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(Limit_FieldPathArrayItemValue), nil
}

func MustParseLimit_FieldPathArrayItemValue(pathStr, valueStr string) Limit_FieldPathArrayItemValue {
	fpaiv, err := ParseLimit_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type Limit_FieldTerminalPathArrayItemValue struct {
	Limit_FieldTerminalPath
	value interface{}
}

var _ Limit_FieldPathArrayItemValue = (*Limit_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object Limit as interface{}
func (fpaiv *Limit_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *Limit_FieldTerminalPathArrayItemValue) GetSingle(source *Limit) (interface{}, bool) {
	return nil, false
}

func (fpaiv *Limit_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*Limit))
}

// Contains returns a boolean indicating if value that is being held is present in given 'Limit'
func (fpaiv *Limit_FieldTerminalPathArrayItemValue) ContainsValue(source *Limit) bool {
	slice := fpaiv.Limit_FieldTerminalPath.Get(source)
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

type Limit_FieldSubPathArrayItemValue struct {
	Limit_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *Limit_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *Limit_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'Limit'
func (fpaivs *Limit_FieldSubPathArrayItemValue) ContainsValue(source *Limit) bool {
	switch fpaivs.Selector() {
	case Limit_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Limit: %d", fpaivs.Selector()))
	}
}

// Limit_FieldPathArrayOfValues allows storing slice of values for Limit fields according to their type
type Limit_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	Limit_FieldPath
}

func ParseLimit_FieldPathArrayOfValues(pathStr, valuesStr string) (Limit_FieldPathArrayOfValues, error) {
	fp, err := ParseLimit_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Limit field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(Limit_FieldPathArrayOfValues), nil
}

func MustParseLimit_FieldPathArrayOfValues(pathStr, valuesStr string) Limit_FieldPathArrayOfValues {
	fpaov, err := ParseLimit_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type Limit_FieldTerminalPathArrayOfValues struct {
	Limit_FieldTerminalPath
	values interface{}
}

var _ Limit_FieldPathArrayOfValues = (*Limit_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *Limit_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case Limit_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case Limit_FieldPathSelectorService:
		for _, v := range fpaov.values.([]*meta_service.Reference) {
			values = append(values, v)
		}
	case Limit_FieldPathSelectorResource:
		for _, v := range fpaov.values.([]*meta_resource.Reference) {
			values = append(values, v)
		}
	case Limit_FieldPathSelectorRegion:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case Limit_FieldPathSelectorConfiguredLimit:
		for _, v := range fpaov.values.([]int64) {
			values = append(values, v)
		}
	case Limit_FieldPathSelectorActiveLimit:
		for _, v := range fpaov.values.([]int64) {
			values = append(values, v)
		}
	case Limit_FieldPathSelectorUsage:
		for _, v := range fpaov.values.([]int64) {
			values = append(values, v)
		}
	case Limit_FieldPathSelectorSource:
		for _, v := range fpaov.values.([]*limit_pool.Reference) {
			values = append(values, v)
		}
	case Limit_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*meta.Meta) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *Limit_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *Limit_FieldTerminalPathArrayOfValues) AsServiceArrayOfValues() ([]*meta_service.Reference, bool) {
	res, ok := fpaov.values.([]*meta_service.Reference)
	return res, ok
}
func (fpaov *Limit_FieldTerminalPathArrayOfValues) AsResourceArrayOfValues() ([]*meta_resource.Reference, bool) {
	res, ok := fpaov.values.([]*meta_resource.Reference)
	return res, ok
}
func (fpaov *Limit_FieldTerminalPathArrayOfValues) AsRegionArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *Limit_FieldTerminalPathArrayOfValues) AsConfiguredLimitArrayOfValues() ([]int64, bool) {
	res, ok := fpaov.values.([]int64)
	return res, ok
}
func (fpaov *Limit_FieldTerminalPathArrayOfValues) AsActiveLimitArrayOfValues() ([]int64, bool) {
	res, ok := fpaov.values.([]int64)
	return res, ok
}
func (fpaov *Limit_FieldTerminalPathArrayOfValues) AsUsageArrayOfValues() ([]int64, bool) {
	res, ok := fpaov.values.([]int64)
	return res, ok
}
func (fpaov *Limit_FieldTerminalPathArrayOfValues) AsSourceArrayOfValues() ([]*limit_pool.Reference, bool) {
	res, ok := fpaov.values.([]*limit_pool.Reference)
	return res, ok
}
func (fpaov *Limit_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*meta.Meta, bool) {
	res, ok := fpaov.values.([]*meta.Meta)
	return res, ok
}

type Limit_FieldSubPathArrayOfValues struct {
	Limit_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ Limit_FieldPathArrayOfValues = (*Limit_FieldSubPathArrayOfValues)(nil)

func (fpsaov *Limit_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *Limit_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
