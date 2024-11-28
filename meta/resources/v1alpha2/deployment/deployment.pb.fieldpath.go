// Code generated by protoc-gen-goten-object
// File: edgelq/meta/proto/v1alpha2/deployment.proto
// DO NOT EDIT!!!

package deployment

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
	region "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/region"
	service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
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
	_ = &region.Region{}
	_ = &service.Service{}
	_ = &meta.Meta{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type Deployment_FieldPath interface {
	gotenobject.FieldPath
	Selector() Deployment_FieldPathSelector
	Get(source *Deployment) []interface{}
	GetSingle(source *Deployment) (interface{}, bool)
	ClearValue(item *Deployment)

	// Those methods build corresponding Deployment_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) Deployment_FieldPathValue
	WithIArrayOfValues(values interface{}) Deployment_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) Deployment_FieldPathArrayItemValue
}

type Deployment_FieldPathSelector int32

const (
	Deployment_FieldPathSelectorName     Deployment_FieldPathSelector = 0
	Deployment_FieldPathSelectorService  Deployment_FieldPathSelector = 1
	Deployment_FieldPathSelectorMetadata Deployment_FieldPathSelector = 2
)

func (s Deployment_FieldPathSelector) String() string {
	switch s {
	case Deployment_FieldPathSelectorName:
		return "name"
	case Deployment_FieldPathSelectorService:
		return "service"
	case Deployment_FieldPathSelectorMetadata:
		return "metadata"
	default:
		panic(fmt.Sprintf("Invalid selector for Deployment: %d", s))
	}
}

func BuildDeployment_FieldPath(fp gotenobject.RawFieldPath) (Deployment_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object Deployment")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorName}, nil
		case "service":
			return &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorService}, nil
		case "metadata":
			return &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorMetadata}, nil
		}
	} else {
		switch fp[0] {
		case "metadata":
			if subpath, err := meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &Deployment_FieldSubPath{selector: Deployment_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object Deployment", fp)
}

func ParseDeployment_FieldPath(rawField string) (Deployment_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildDeployment_FieldPath(fp)
}

func MustParseDeployment_FieldPath(rawField string) Deployment_FieldPath {
	fp, err := ParseDeployment_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type Deployment_FieldTerminalPath struct {
	selector Deployment_FieldPathSelector
}

var _ Deployment_FieldPath = (*Deployment_FieldTerminalPath)(nil)

func (fp *Deployment_FieldTerminalPath) Selector() Deployment_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *Deployment_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *Deployment_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source Deployment
func (fp *Deployment_FieldTerminalPath) Get(source *Deployment) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case Deployment_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case Deployment_FieldPathSelectorService:
			if source.Service != nil {
				values = append(values, source.Service)
			}
		case Deployment_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		default:
			panic(fmt.Sprintf("Invalid selector for Deployment: %d", fp.selector))
		}
	}
	return
}

func (fp *Deployment_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*Deployment))
}

// GetSingle returns value pointed by specific field of from source Deployment
func (fp *Deployment_FieldTerminalPath) GetSingle(source *Deployment) (interface{}, bool) {
	switch fp.selector {
	case Deployment_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case Deployment_FieldPathSelectorService:
		res := source.GetService()
		return res, res != nil
	case Deployment_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	default:
		panic(fmt.Sprintf("Invalid selector for Deployment: %d", fp.selector))
	}
}

func (fp *Deployment_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*Deployment))
}

// GetDefault returns a default value of the field type
func (fp *Deployment_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case Deployment_FieldPathSelectorName:
		return (*Name)(nil)
	case Deployment_FieldPathSelectorService:
		return (*service.Reference)(nil)
	case Deployment_FieldPathSelectorMetadata:
		return (*meta.Meta)(nil)
	default:
		panic(fmt.Sprintf("Invalid selector for Deployment: %d", fp.selector))
	}
}

func (fp *Deployment_FieldTerminalPath) ClearValue(item *Deployment) {
	if item != nil {
		switch fp.selector {
		case Deployment_FieldPathSelectorName:
			item.Name = nil
		case Deployment_FieldPathSelectorService:
			item.Service = nil
		case Deployment_FieldPathSelectorMetadata:
			item.Metadata = nil
		default:
			panic(fmt.Sprintf("Invalid selector for Deployment: %d", fp.selector))
		}
	}
}

func (fp *Deployment_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*Deployment))
}

// IsLeaf - whether field path is holds simple value
func (fp *Deployment_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == Deployment_FieldPathSelectorName ||
		fp.selector == Deployment_FieldPathSelectorService
}

func (fp *Deployment_FieldTerminalPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	return []gotenobject.FieldPath{fp}
}

func (fp *Deployment_FieldTerminalPath) WithIValue(value interface{}) Deployment_FieldPathValue {
	switch fp.selector {
	case Deployment_FieldPathSelectorName:
		return &Deployment_FieldTerminalPathValue{Deployment_FieldTerminalPath: *fp, value: value.(*Name)}
	case Deployment_FieldPathSelectorService:
		return &Deployment_FieldTerminalPathValue{Deployment_FieldTerminalPath: *fp, value: value.(*service.Reference)}
	case Deployment_FieldPathSelectorMetadata:
		return &Deployment_FieldTerminalPathValue{Deployment_FieldTerminalPath: *fp, value: value.(*meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for Deployment: %d", fp.selector))
	}
}

func (fp *Deployment_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *Deployment_FieldTerminalPath) WithIArrayOfValues(values interface{}) Deployment_FieldPathArrayOfValues {
	fpaov := &Deployment_FieldTerminalPathArrayOfValues{Deployment_FieldTerminalPath: *fp}
	switch fp.selector {
	case Deployment_FieldPathSelectorName:
		return &Deployment_FieldTerminalPathArrayOfValues{Deployment_FieldTerminalPath: *fp, values: values.([]*Name)}
	case Deployment_FieldPathSelectorService:
		return &Deployment_FieldTerminalPathArrayOfValues{Deployment_FieldTerminalPath: *fp, values: values.([]*service.Reference)}
	case Deployment_FieldPathSelectorMetadata:
		return &Deployment_FieldTerminalPathArrayOfValues{Deployment_FieldTerminalPath: *fp, values: values.([]*meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for Deployment: %d", fp.selector))
	}
	return fpaov
}

func (fp *Deployment_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *Deployment_FieldTerminalPath) WithIArrayItemValue(value interface{}) Deployment_FieldPathArrayItemValue {
	switch fp.selector {
	default:
		panic(fmt.Sprintf("Invalid selector for Deployment: %d", fp.selector))
	}
}

func (fp *Deployment_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

type Deployment_FieldSubPath struct {
	selector Deployment_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ Deployment_FieldPath = (*Deployment_FieldSubPath)(nil)

func (fps *Deployment_FieldSubPath) Selector() Deployment_FieldPathSelector {
	return fps.selector
}
func (fps *Deployment_FieldSubPath) AsMetadataSubPath() (meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(meta.Meta_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *Deployment_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *Deployment_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source Deployment
func (fps *Deployment_FieldSubPath) Get(source *Deployment) (values []interface{}) {
	switch fps.selector {
	case Deployment_FieldPathSelectorMetadata:
		values = append(values, fps.subPath.GetRaw(source.GetMetadata())...)
	default:
		panic(fmt.Sprintf("Invalid selector for Deployment: %d", fps.selector))
	}
	return
}

func (fps *Deployment_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*Deployment))
}

// GetSingle returns value of selected field from source Deployment
func (fps *Deployment_FieldSubPath) GetSingle(source *Deployment) (interface{}, bool) {
	switch fps.selector {
	case Deployment_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Deployment: %d", fps.selector))
	}
}

func (fps *Deployment_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*Deployment))
}

// GetDefault returns a default value of the field type
func (fps *Deployment_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *Deployment_FieldSubPath) ClearValue(item *Deployment) {
	if item != nil {
		switch fps.selector {
		case Deployment_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		default:
			panic(fmt.Sprintf("Invalid selector for Deployment: %d", fps.selector))
		}
	}
}

func (fps *Deployment_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*Deployment))
}

// IsLeaf - whether field path is holds simple value
func (fps *Deployment_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *Deployment_FieldSubPath) SplitIntoTerminalIPaths() []gotenobject.FieldPath {
	iPaths := []gotenobject.FieldPath{&Deployment_FieldTerminalPath{selector: fps.selector}}
	iPaths = append(iPaths, fps.subPath.SplitIntoTerminalIPaths()...)
	return iPaths
}

func (fps *Deployment_FieldSubPath) WithIValue(value interface{}) Deployment_FieldPathValue {
	return &Deployment_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *Deployment_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *Deployment_FieldSubPath) WithIArrayOfValues(values interface{}) Deployment_FieldPathArrayOfValues {
	return &Deployment_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *Deployment_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *Deployment_FieldSubPath) WithIArrayItemValue(value interface{}) Deployment_FieldPathArrayItemValue {
	return &Deployment_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *Deployment_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// Deployment_FieldPathValue allows storing values for Deployment fields according to their type
type Deployment_FieldPathValue interface {
	Deployment_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **Deployment)
	CompareWith(*Deployment) (cmp int, comparable bool)
}

func ParseDeployment_FieldPathValue(pathStr, valueStr string) (Deployment_FieldPathValue, error) {
	fp, err := ParseDeployment_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Deployment field path value from %s: %v", valueStr, err)
	}
	return fpv.(Deployment_FieldPathValue), nil
}

func MustParseDeployment_FieldPathValue(pathStr, valueStr string) Deployment_FieldPathValue {
	fpv, err := ParseDeployment_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type Deployment_FieldTerminalPathValue struct {
	Deployment_FieldTerminalPath
	value interface{}
}

var _ Deployment_FieldPathValue = (*Deployment_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'Deployment' as interface{}
func (fpv *Deployment_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *Deployment_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *Deployment_FieldTerminalPathValue) AsServiceValue() (*service.Reference, bool) {
	res, ok := fpv.value.(*service.Reference)
	return res, ok
}
func (fpv *Deployment_FieldTerminalPathValue) AsMetadataValue() (*meta.Meta, bool) {
	res, ok := fpv.value.(*meta.Meta)
	return res, ok
}

// SetTo stores value for selected field for object Deployment
func (fpv *Deployment_FieldTerminalPathValue) SetTo(target **Deployment) {
	if *target == nil {
		*target = new(Deployment)
	}
	switch fpv.selector {
	case Deployment_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case Deployment_FieldPathSelectorService:
		(*target).Service = fpv.value.(*service.Reference)
	case Deployment_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*meta.Meta)
	default:
		panic(fmt.Sprintf("Invalid selector for Deployment: %d", fpv.selector))
	}
}

func (fpv *Deployment_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Deployment)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'Deployment_FieldTerminalPathValue' with the value under path in 'Deployment'.
func (fpv *Deployment_FieldTerminalPathValue) CompareWith(source *Deployment) (int, bool) {
	switch fpv.selector {
	case Deployment_FieldPathSelectorName:
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
	case Deployment_FieldPathSelectorService:
		leftValue := fpv.value.(*service.Reference)
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
	case Deployment_FieldPathSelectorMetadata:
		return 0, false
	default:
		panic(fmt.Sprintf("Invalid selector for Deployment: %d", fpv.selector))
	}
}

func (fpv *Deployment_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*Deployment))
}

type Deployment_FieldSubPathValue struct {
	Deployment_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ Deployment_FieldPathValue = (*Deployment_FieldSubPathValue)(nil)

func (fpvs *Deployment_FieldSubPathValue) AsMetadataPathValue() (meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(meta.Meta_FieldPathValue)
	return res, ok
}

func (fpvs *Deployment_FieldSubPathValue) SetTo(target **Deployment) {
	if *target == nil {
		*target = new(Deployment)
	}
	switch fpvs.Selector() {
	case Deployment_FieldPathSelectorMetadata:
		fpvs.subPathValue.(meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	default:
		panic(fmt.Sprintf("Invalid selector for Deployment: %d", fpvs.Selector()))
	}
}

func (fpvs *Deployment_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Deployment)
	fpvs.SetTo(&typedObject)
}

func (fpvs *Deployment_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *Deployment_FieldSubPathValue) CompareWith(source *Deployment) (int, bool) {
	switch fpvs.Selector() {
	case Deployment_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Deployment: %d", fpvs.Selector()))
	}
}

func (fpvs *Deployment_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*Deployment))
}

// Deployment_FieldPathArrayItemValue allows storing single item in Path-specific values for Deployment according to their type
// Present only for array (repeated) types.
type Deployment_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	Deployment_FieldPath
	ContainsValue(*Deployment) bool
}

// ParseDeployment_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseDeployment_FieldPathArrayItemValue(pathStr, valueStr string) (Deployment_FieldPathArrayItemValue, error) {
	fp, err := ParseDeployment_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Deployment field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(Deployment_FieldPathArrayItemValue), nil
}

func MustParseDeployment_FieldPathArrayItemValue(pathStr, valueStr string) Deployment_FieldPathArrayItemValue {
	fpaiv, err := ParseDeployment_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type Deployment_FieldTerminalPathArrayItemValue struct {
	Deployment_FieldTerminalPath
	value interface{}
}

var _ Deployment_FieldPathArrayItemValue = (*Deployment_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object Deployment as interface{}
func (fpaiv *Deployment_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}

func (fpaiv *Deployment_FieldTerminalPathArrayItemValue) GetSingle(source *Deployment) (interface{}, bool) {
	return nil, false
}

func (fpaiv *Deployment_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*Deployment))
}

// Contains returns a boolean indicating if value that is being held is present in given 'Deployment'
func (fpaiv *Deployment_FieldTerminalPathArrayItemValue) ContainsValue(source *Deployment) bool {
	slice := fpaiv.Deployment_FieldTerminalPath.Get(source)
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

type Deployment_FieldSubPathArrayItemValue struct {
	Deployment_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *Deployment_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *Deployment_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'Deployment'
func (fpaivs *Deployment_FieldSubPathArrayItemValue) ContainsValue(source *Deployment) bool {
	switch fpaivs.Selector() {
	case Deployment_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for Deployment: %d", fpaivs.Selector()))
	}
}

// Deployment_FieldPathArrayOfValues allows storing slice of values for Deployment fields according to their type
type Deployment_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	Deployment_FieldPath
}

func ParseDeployment_FieldPathArrayOfValues(pathStr, valuesStr string) (Deployment_FieldPathArrayOfValues, error) {
	fp, err := ParseDeployment_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Deployment field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(Deployment_FieldPathArrayOfValues), nil
}

func MustParseDeployment_FieldPathArrayOfValues(pathStr, valuesStr string) Deployment_FieldPathArrayOfValues {
	fpaov, err := ParseDeployment_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type Deployment_FieldTerminalPathArrayOfValues struct {
	Deployment_FieldTerminalPath
	values interface{}
}

var _ Deployment_FieldPathArrayOfValues = (*Deployment_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *Deployment_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case Deployment_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case Deployment_FieldPathSelectorService:
		for _, v := range fpaov.values.([]*service.Reference) {
			values = append(values, v)
		}
	case Deployment_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*meta.Meta) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *Deployment_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *Deployment_FieldTerminalPathArrayOfValues) AsServiceArrayOfValues() ([]*service.Reference, bool) {
	res, ok := fpaov.values.([]*service.Reference)
	return res, ok
}
func (fpaov *Deployment_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*meta.Meta, bool) {
	res, ok := fpaov.values.([]*meta.Meta)
	return res, ok
}

type Deployment_FieldSubPathArrayOfValues struct {
	Deployment_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ Deployment_FieldPathArrayOfValues = (*Deployment_FieldSubPathArrayOfValues)(nil)

func (fpsaov *Deployment_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *Deployment_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
