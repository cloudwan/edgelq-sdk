// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha2/organization.proto
// DO NOT EDIT!!!

package organization

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
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	policy "github.com/cloudwan/edgelq-sdk/meta/multi_region/proto/policy"
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
	_ = &ntt_meta.Meta{}
	_ = &policy.Policy{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type Organization_FieldPath interface {
	gotenobject.FieldPath
	Selector() Organization_FieldPathSelector
	Get(source *Organization) []interface{}
	GetSingle(source *Organization) (interface{}, bool)
	ClearValue(item *Organization)

	// Those methods build corresponding Organization_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) Organization_FieldPathValue
	WithIArrayOfValues(values interface{}) Organization_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) Organization_FieldPathArrayItemValue
}

type Organization_FieldPathSelector int32

const (
	Organization_FieldPathSelectorName               Organization_FieldPathSelector = 0
	Organization_FieldPathSelectorTitle              Organization_FieldPathSelector = 1
	Organization_FieldPathSelectorParentOrganization Organization_FieldPathSelector = 2
	Organization_FieldPathSelectorRootOrganization   Organization_FieldPathSelector = 3
	Organization_FieldPathSelectorAncestryPath       Organization_FieldPathSelector = 4
	Organization_FieldPathSelectorMetadata           Organization_FieldPathSelector = 5
	Organization_FieldPathSelectorMultiRegionPolicy  Organization_FieldPathSelector = 6
)

func (s Organization_FieldPathSelector) String() string {
	switch s {
	case Organization_FieldPathSelectorName:
		return "name"
	case Organization_FieldPathSelectorTitle:
		return "title"
	case Organization_FieldPathSelectorParentOrganization:
		return "parent_organization"
	case Organization_FieldPathSelectorRootOrganization:
		return "root_organization"
	case Organization_FieldPathSelectorAncestryPath:
		return "ancestry_path"
	case Organization_FieldPathSelectorMetadata:
		return "metadata"
	case Organization_FieldPathSelectorMultiRegionPolicy:
		return "multi_region_policy"
	default:
		panic(fmt.Sprintf("Invalid selector for Organization: %d", s))
	}
}

func BuildOrganization_FieldPath(fp gotenobject.RawFieldPath) (Organization_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object Organization")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorName}, nil
		case "title":
			return &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorTitle}, nil
		case "parent_organization", "parentOrganization", "parent-organization":
			return &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorParentOrganization}, nil
		case "root_organization", "rootOrganization", "root-organization":
			return &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorRootOrganization}, nil
		case "ancestry_path", "ancestryPath", "ancestry-path":
			return &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorAncestryPath}, nil
		case "metadata":
			return &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorMetadata}, nil
		case "multi_region_policy", "multiRegionPolicy", "multi-region-policy":
			return &Organization_FieldTerminalPath{selector: Organization_FieldPathSelectorMultiRegionPolicy}, nil
		}
	} else {
		switch fp[0] {
		case "metadata":
			if subpath, err := ntt_meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &Organization_FieldSubPath{selector: Organization_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		case "multi_region_policy", "multiRegionPolicy", "multi-region-policy":
			if subpath, err := policy.BuildPolicy_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &Organization_FieldSubPath{selector: Organization_FieldPathSelectorMultiRegionPolicy, subPath: subpath}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object Organization", fp)
}

func ParseOrganization_FieldPath(rawField string) (Organization_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildOrganization_FieldPath(fp)
}

func MustParseOrganization_FieldPath(rawField string) Organization_FieldPath {
	fp, err := ParseOrganization_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type Organization_FieldTerminalPath struct {
	selector Organization_FieldPathSelector
}

var _ Organization_FieldPath = (*Organization_FieldTerminalPath)(nil)

func (fp *Organization_FieldTerminalPath) Selector() Organization_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *Organization_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *Organization_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source Organization
func (fp *Organization_FieldTerminalPath) Get(source *Organization) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case Organization_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case Organization_FieldPathSelectorTitle:
			values = append(values, source.Title)
		case Organization_FieldPathSelectorParentOrganization:
			if source.ParentOrganization != nil {
				values = append(values, source.ParentOrganization)
			}
		case Organization_FieldPathSelectorRootOrganization:
			if source.RootOrganization != nil {
				values = append(values, source.RootOrganization)
			}
		case Organization_FieldPathSelectorAncestryPath:
			for _, value := range source.GetAncestryPath() {
				values = append(values, value)
			}
		case Organization_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		case Organization_FieldPathSelectorMultiRegionPolicy:
			if source.MultiRegionPolicy != nil {
				values = append(values, source.MultiRegionPolicy)
			}
		default:
			panic(fmt.Sprintf("Invalid selector for Organization: %d", fp.selector))
		}
	}
	return
}

func (fp *Organization_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*Organization))
}

// GetSingle returns value pointed by specific field of from source Organization
func (fp *Organization_FieldTerminalPath) GetSingle(source *Organization) (interface{}, bool) {
	switch fp.selector {
	case Organization_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case Organization_FieldPathSelectorTitle:
		return source.GetTitle(), source != nil
	case Organization_FieldPathSelectorParentOrganization:
		res := source.GetParentOrganization()
		return res, res != nil
	case Organization_FieldPathSelectorRootOrganization:
		res := source.GetRootOrganization()
		return res, res != nil
	case Organization_FieldPathSelectorAncestryPath:
		res := source.GetAncestryPath()
		return res, res != nil
	case Organization_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	case Organization_FieldPathSelectorMultiRegionPolicy:
		res := source.GetMultiRegionPolicy()
		return res, res != nil
	default:
		panic(fmt.Sprintf("Invalid selector for Organization: %d", fp.selector))
	}
}

func (fp *Organization_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*Organization))
}

// GetDefault returns a default value of the field type
func (fp *Organization_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case Organization_FieldPathSelectorName:
		return (*Name)(nil)
	case Organization_FieldPathSelectorTitle:
		return ""
	case Organization_FieldPathSelectorParentOrganization:
		return (*Reference)(nil)
	case Organization_FieldPathSelectorRootOrganization:
		return (*Reference)(nil)
	case Organization_FieldPathSelectorAncestryPath:
		return ([]*Reference)(nil)
	case Organization_FieldPathSelectorMetadata:
		return (*ntt_meta.Meta)(nil)
	case Organization_FieldPathSelectorMultiRegionPolicy:
		return (*policy.Policy)(nil)
	default:
		panic(fmt.Sprintf("Invalid selector for Organization: %d", fp.selector))
	}
}

func (fp *Organization_FieldTerminalPath) ClearValue(item *Organization) {
	if item != nil {
		switch fp.selector {
		case Organization_FieldPathSelectorName:
			item.Name = nil
		case Organization_FieldPathSelectorTitle:
			item.Title = ""
		case Organization_FieldPathSelectorParentOrganization:
			item.ParentOrganization = nil
		case Organization_FieldPathSelectorRootOrganization:
			item.RootOrganization = nil
		case Organization_FieldPathSelectorAncestryPath:
			item.AncestryPath = nil
		case Organization_FieldPathSelectorMetadata:
			item.Metadata = nil
		case Organization_FieldPathSelectorMultiRegionPolicy:
			item.MultiRegionPolicy = nil
		default:
			panic(fmt.Sprintf("Invalid selector for Organization: %d", fp.selector))
		}
	}
}

func (fp *Organization_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*Organization))
}

// IsLeaf - whether field path is holds simple value
func (fp *Organization_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == Organization_FieldPathSelectorName ||
		fp.selector == Organization_FieldPathSelectorTitle ||
		fp.selector == Organization_FieldPathSelectorParentOrganization ||
		fp.selector == Organization_FieldPathSelectorRootOrganization ||
		fp.selector == Organization_FieldPathSelectorAncestryPath
}

func (fp *Organization_FieldTerminalPath) WithIValue(value interface{}) Organization_FieldPathValue {
	switch fp.selector {
	case Organization_FieldPathSelectorName:
		return &Organization_FieldTerminalPathValue{Organization_FieldTerminalPath: *fp, value: value.(*Name)}
	case Organization_FieldPathSelectorTitle:
		return &Organization_FieldTerminalPathValue{Organization_FieldTerminalPath: *fp, value: value.(string)}
	case Organization_FieldPathSelectorParentOrganization:
		return &Organization_FieldTerminalPathValue{Organization_FieldTerminalPath: *fp, value: value.(*Reference)}
	case Organization_FieldPathSelectorRootOrganization:
		return &Organization_FieldTerminalPathValue{Organization_FieldTerminalPath: *fp, value: value.(*Reference)}
	case Organization_FieldPathSelectorAncestryPath:
		return &Organization_FieldTerminalPathValue{Organization_FieldTerminalPath: *fp, value: value.([]*Reference)}
	case Organization_FieldPathSelectorMetadata:
		return &Organization_FieldTerminalPathValue{Organization_FieldTerminalPath: *fp, value: value.(*ntt_meta.Meta)}
	case Organization_FieldPathSelectorMultiRegionPolicy:
		return &Organization_FieldTerminalPathValue{Organization_FieldTerminalPath: *fp, value: value.(*policy.Policy)}
	default:
		panic(fmt.Sprintf("Invalid selector for Organization: %d", fp.selector))
	}
}

func (fp *Organization_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *Organization_FieldTerminalPath) WithIArrayOfValues(values interface{}) Organization_FieldPathArrayOfValues {
	fpaov := &Organization_FieldTerminalPathArrayOfValues{Organization_FieldTerminalPath: *fp}
	switch fp.selector {
	case Organization_FieldPathSelectorName:
		return &Organization_FieldTerminalPathArrayOfValues{Organization_FieldTerminalPath: *fp, values: values.([]*Name)}
	case Organization_FieldPathSelectorTitle:
		return &Organization_FieldTerminalPathArrayOfValues{Organization_FieldTerminalPath: *fp, values: values.([]string)}
	case Organization_FieldPathSelectorParentOrganization:
		return &Organization_FieldTerminalPathArrayOfValues{Organization_FieldTerminalPath: *fp, values: values.([]*Reference)}
	case Organization_FieldPathSelectorRootOrganization:
		return &Organization_FieldTerminalPathArrayOfValues{Organization_FieldTerminalPath: *fp, values: values.([]*Reference)}
	case Organization_FieldPathSelectorAncestryPath:
		return &Organization_FieldTerminalPathArrayOfValues{Organization_FieldTerminalPath: *fp, values: values.([][]*Reference)}
	case Organization_FieldPathSelectorMetadata:
		return &Organization_FieldTerminalPathArrayOfValues{Organization_FieldTerminalPath: *fp, values: values.([]*ntt_meta.Meta)}
	case Organization_FieldPathSelectorMultiRegionPolicy:
		return &Organization_FieldTerminalPathArrayOfValues{Organization_FieldTerminalPath: *fp, values: values.([]*policy.Policy)}
	default:
		panic(fmt.Sprintf("Invalid selector for Organization: %d", fp.selector))
	}
	return fpaov
}

func (fp *Organization_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *Organization_FieldTerminalPath) WithIArrayItemValue(value interface{}) Organization_FieldPathArrayItemValue {
	switch fp.selector {
	case Organization_FieldPathSelectorAncestryPath:
		return &Organization_FieldTerminalPathArrayItemValue{Organization_FieldTerminalPath: *fp, value: value.(*Reference)}
	default:
		panic(fmt.Sprintf("Invalid selector for Organization: %d", fp.selector))
	}
}

func (fp *Organization_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

type Organization_FieldSubPath struct {
	selector Organization_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ Organization_FieldPath = (*Organization_FieldSubPath)(nil)

func (fps *Organization_FieldSubPath) Selector() Organization_FieldPathSelector {
	return fps.selector
}
func (fps *Organization_FieldSubPath) AsMetadataSubPath() (ntt_meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(ntt_meta.Meta_FieldPath)
	return res, ok
}
func (fps *Organization_FieldSubPath) AsMultiRegionPolicySubPath() (policy.Policy_FieldPath, bool) {
	res, ok := fps.subPath.(policy.Policy_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *Organization_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *Organization_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source Organization
func (fps *Organization_FieldSubPath) Get(source *Organization) (values []interface{}) {
	if asMetaFieldPath, ok := fps.AsMetadataSubPath(); ok {
		values = append(values, asMetaFieldPath.Get(source.GetMetadata())...)
	} else if asPolicyFieldPath, ok := fps.AsMultiRegionPolicySubPath(); ok {
		values = append(values, asPolicyFieldPath.Get(source.GetMultiRegionPolicy())...)
	} else {
		panic(fmt.Sprintf("Invalid selector for Organization: %d", fps.selector))
	}
	return
}

func (fps *Organization_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*Organization))
}

// GetSingle returns value of selected field from source Organization
func (fps *Organization_FieldSubPath) GetSingle(source *Organization) (interface{}, bool) {
	switch fps.selector {
	case Organization_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	case Organization_FieldPathSelectorMultiRegionPolicy:
		if source.GetMultiRegionPolicy() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMultiRegionPolicy())
	default:
		panic(fmt.Sprintf("Invalid selector for Organization: %d", fps.selector))
	}
}

func (fps *Organization_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*Organization))
}

// GetDefault returns a default value of the field type
func (fps *Organization_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *Organization_FieldSubPath) ClearValue(item *Organization) {
	if item != nil {
		switch fps.selector {
		case Organization_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		case Organization_FieldPathSelectorMultiRegionPolicy:
			fps.subPath.ClearValueRaw(item.MultiRegionPolicy)
		default:
			panic(fmt.Sprintf("Invalid selector for Organization: %d", fps.selector))
		}
	}
}

func (fps *Organization_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*Organization))
}

// IsLeaf - whether field path is holds simple value
func (fps *Organization_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *Organization_FieldSubPath) WithIValue(value interface{}) Organization_FieldPathValue {
	return &Organization_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *Organization_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *Organization_FieldSubPath) WithIArrayOfValues(values interface{}) Organization_FieldPathArrayOfValues {
	return &Organization_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *Organization_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *Organization_FieldSubPath) WithIArrayItemValue(value interface{}) Organization_FieldPathArrayItemValue {
	return &Organization_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *Organization_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// Organization_FieldPathValue allows storing values for Organization fields according to their type
type Organization_FieldPathValue interface {
	Organization_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **Organization)
	CompareWith(*Organization) (cmp int, comparable bool)
}

func ParseOrganization_FieldPathValue(pathStr, valueStr string) (Organization_FieldPathValue, error) {
	fp, err := ParseOrganization_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Organization field path value from %s: %v", valueStr, err)
	}
	return fpv.(Organization_FieldPathValue), nil
}

func MustParseOrganization_FieldPathValue(pathStr, valueStr string) Organization_FieldPathValue {
	fpv, err := ParseOrganization_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type Organization_FieldTerminalPathValue struct {
	Organization_FieldTerminalPath
	value interface{}
}

var _ Organization_FieldPathValue = (*Organization_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'Organization' as interface{}
func (fpv *Organization_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *Organization_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *Organization_FieldTerminalPathValue) AsTitleValue() (string, bool) {
	res, ok := fpv.value.(string)
	return res, ok
}
func (fpv *Organization_FieldTerminalPathValue) AsParentOrganizationValue() (*Reference, bool) {
	res, ok := fpv.value.(*Reference)
	return res, ok
}
func (fpv *Organization_FieldTerminalPathValue) AsRootOrganizationValue() (*Reference, bool) {
	res, ok := fpv.value.(*Reference)
	return res, ok
}
func (fpv *Organization_FieldTerminalPathValue) AsAncestryPathValue() ([]*Reference, bool) {
	res, ok := fpv.value.([]*Reference)
	return res, ok
}
func (fpv *Organization_FieldTerminalPathValue) AsMetadataValue() (*ntt_meta.Meta, bool) {
	res, ok := fpv.value.(*ntt_meta.Meta)
	return res, ok
}
func (fpv *Organization_FieldTerminalPathValue) AsMultiRegionPolicyValue() (*policy.Policy, bool) {
	res, ok := fpv.value.(*policy.Policy)
	return res, ok
}

// SetTo stores value for selected field for object Organization
func (fpv *Organization_FieldTerminalPathValue) SetTo(target **Organization) {
	if *target == nil {
		*target = new(Organization)
	}
	switch fpv.selector {
	case Organization_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case Organization_FieldPathSelectorTitle:
		(*target).Title = fpv.value.(string)
	case Organization_FieldPathSelectorParentOrganization:
		(*target).ParentOrganization = fpv.value.(*Reference)
	case Organization_FieldPathSelectorRootOrganization:
		(*target).RootOrganization = fpv.value.(*Reference)
	case Organization_FieldPathSelectorAncestryPath:
		(*target).AncestryPath = fpv.value.([]*Reference)
	case Organization_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*ntt_meta.Meta)
	case Organization_FieldPathSelectorMultiRegionPolicy:
		(*target).MultiRegionPolicy = fpv.value.(*policy.Policy)
	default:
		panic(fmt.Sprintf("Invalid selector for Organization: %d", fpv.selector))
	}
}

func (fpv *Organization_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Organization)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'Organization_FieldTerminalPathValue' with the value under path in 'Organization'.
func (fpv *Organization_FieldTerminalPathValue) CompareWith(source *Organization) (int, bool) {
	switch fpv.selector {
	case Organization_FieldPathSelectorName:
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
	case Organization_FieldPathSelectorTitle:
		leftValue := fpv.value.(string)
		rightValue := source.GetTitle()
		if (leftValue) == (rightValue) {
			return 0, true
		} else if (leftValue) < (rightValue) {
			return -1, true
		} else {
			return 1, true
		}
	case Organization_FieldPathSelectorParentOrganization:
		leftValue := fpv.value.(*Reference)
		rightValue := source.GetParentOrganization()
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
	case Organization_FieldPathSelectorRootOrganization:
		leftValue := fpv.value.(*Reference)
		rightValue := source.GetRootOrganization()
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
	case Organization_FieldPathSelectorAncestryPath:
		return 0, false
	case Organization_FieldPathSelectorMetadata:
		return 0, false
	case Organization_FieldPathSelectorMultiRegionPolicy:
		return 0, false
	default:
		panic(fmt.Sprintf("Invalid selector for Organization: %d", fpv.selector))
	}
}

func (fpv *Organization_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*Organization))
}

type Organization_FieldSubPathValue struct {
	Organization_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ Organization_FieldPathValue = (*Organization_FieldSubPathValue)(nil)

func (fpvs *Organization_FieldSubPathValue) AsMetadataPathValue() (ntt_meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(ntt_meta.Meta_FieldPathValue)
	return res, ok
}
func (fpvs *Organization_FieldSubPathValue) AsMultiRegionPolicyPathValue() (policy.Policy_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(policy.Policy_FieldPathValue)
	return res, ok
}

func (fpvs *Organization_FieldSubPathValue) SetTo(target **Organization) {
	if *target == nil {
		*target = new(Organization)
	}
	switch fpvs.Selector() {
	case Organization_FieldPathSelectorMetadata:
		fpvs.subPathValue.(ntt_meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	case Organization_FieldPathSelectorMultiRegionPolicy:
		fpvs.subPathValue.(policy.Policy_FieldPathValue).SetTo(&(*target).MultiRegionPolicy)
	default:
		panic(fmt.Sprintf("Invalid selector for Organization: %d", fpvs.Selector()))
	}
}

func (fpvs *Organization_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*Organization)
	fpvs.SetTo(&typedObject)
}

func (fpvs *Organization_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *Organization_FieldSubPathValue) CompareWith(source *Organization) (int, bool) {
	switch fpvs.Selector() {
	case Organization_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(ntt_meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	case Organization_FieldPathSelectorMultiRegionPolicy:
		return fpvs.subPathValue.(policy.Policy_FieldPathValue).CompareWith(source.GetMultiRegionPolicy())
	default:
		panic(fmt.Sprintf("Invalid selector for Organization: %d", fpvs.Selector()))
	}
}

func (fpvs *Organization_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*Organization))
}

// Organization_FieldPathArrayItemValue allows storing single item in Path-specific values for Organization according to their type
// Present only for array (repeated) types.
type Organization_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	Organization_FieldPath
	ContainsValue(*Organization) bool
}

// ParseOrganization_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParseOrganization_FieldPathArrayItemValue(pathStr, valueStr string) (Organization_FieldPathArrayItemValue, error) {
	fp, err := ParseOrganization_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Organization field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(Organization_FieldPathArrayItemValue), nil
}

func MustParseOrganization_FieldPathArrayItemValue(pathStr, valueStr string) Organization_FieldPathArrayItemValue {
	fpaiv, err := ParseOrganization_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type Organization_FieldTerminalPathArrayItemValue struct {
	Organization_FieldTerminalPath
	value interface{}
}

var _ Organization_FieldPathArrayItemValue = (*Organization_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object Organization as interface{}
func (fpaiv *Organization_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}
func (fpaiv *Organization_FieldTerminalPathArrayItemValue) AsAncestryPathItemValue() (*Reference, bool) {
	res, ok := fpaiv.value.(*Reference)
	return res, ok
}

func (fpaiv *Organization_FieldTerminalPathArrayItemValue) GetSingle(source *Organization) (interface{}, bool) {
	return nil, false
}

func (fpaiv *Organization_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*Organization))
}

// Contains returns a boolean indicating if value that is being held is present in given 'Organization'
func (fpaiv *Organization_FieldTerminalPathArrayItemValue) ContainsValue(source *Organization) bool {
	slice := fpaiv.Organization_FieldTerminalPath.Get(source)
	for _, v := range slice {
		if reflect.DeepEqual(v, fpaiv.value) {
			return true
		}
	}
	return false
}

type Organization_FieldSubPathArrayItemValue struct {
	Organization_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *Organization_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *Organization_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (ntt_meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(ntt_meta.Meta_FieldPathArrayItemValue)
	return res, ok
}
func (fpaivs *Organization_FieldSubPathArrayItemValue) AsMultiRegionPolicyPathItemValue() (policy.Policy_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(policy.Policy_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'Organization'
func (fpaivs *Organization_FieldSubPathArrayItemValue) ContainsValue(source *Organization) bool {
	switch fpaivs.Selector() {
	case Organization_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(ntt_meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	case Organization_FieldPathSelectorMultiRegionPolicy:
		return fpaivs.subPathItemValue.(policy.Policy_FieldPathArrayItemValue).ContainsValue(source.GetMultiRegionPolicy())
	default:
		panic(fmt.Sprintf("Invalid selector for Organization: %d", fpaivs.Selector()))
	}
}

// Organization_FieldPathArrayOfValues allows storing slice of values for Organization fields according to their type
type Organization_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	Organization_FieldPath
}

func ParseOrganization_FieldPathArrayOfValues(pathStr, valuesStr string) (Organization_FieldPathArrayOfValues, error) {
	fp, err := ParseOrganization_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing Organization field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(Organization_FieldPathArrayOfValues), nil
}

func MustParseOrganization_FieldPathArrayOfValues(pathStr, valuesStr string) Organization_FieldPathArrayOfValues {
	fpaov, err := ParseOrganization_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type Organization_FieldTerminalPathArrayOfValues struct {
	Organization_FieldTerminalPath
	values interface{}
}

var _ Organization_FieldPathArrayOfValues = (*Organization_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *Organization_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case Organization_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case Organization_FieldPathSelectorTitle:
		for _, v := range fpaov.values.([]string) {
			values = append(values, v)
		}
	case Organization_FieldPathSelectorParentOrganization:
		for _, v := range fpaov.values.([]*Reference) {
			values = append(values, v)
		}
	case Organization_FieldPathSelectorRootOrganization:
		for _, v := range fpaov.values.([]*Reference) {
			values = append(values, v)
		}
	case Organization_FieldPathSelectorAncestryPath:
		for _, v := range fpaov.values.([][]*Reference) {
			values = append(values, v)
		}
	case Organization_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*ntt_meta.Meta) {
			values = append(values, v)
		}
	case Organization_FieldPathSelectorMultiRegionPolicy:
		for _, v := range fpaov.values.([]*policy.Policy) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *Organization_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *Organization_FieldTerminalPathArrayOfValues) AsTitleArrayOfValues() ([]string, bool) {
	res, ok := fpaov.values.([]string)
	return res, ok
}
func (fpaov *Organization_FieldTerminalPathArrayOfValues) AsParentOrganizationArrayOfValues() ([]*Reference, bool) {
	res, ok := fpaov.values.([]*Reference)
	return res, ok
}
func (fpaov *Organization_FieldTerminalPathArrayOfValues) AsRootOrganizationArrayOfValues() ([]*Reference, bool) {
	res, ok := fpaov.values.([]*Reference)
	return res, ok
}
func (fpaov *Organization_FieldTerminalPathArrayOfValues) AsAncestryPathArrayOfValues() ([][]*Reference, bool) {
	res, ok := fpaov.values.([][]*Reference)
	return res, ok
}
func (fpaov *Organization_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*ntt_meta.Meta, bool) {
	res, ok := fpaov.values.([]*ntt_meta.Meta)
	return res, ok
}
func (fpaov *Organization_FieldTerminalPathArrayOfValues) AsMultiRegionPolicyArrayOfValues() ([]*policy.Policy, bool) {
	res, ok := fpaov.values.([]*policy.Policy)
	return res, ok
}

type Organization_FieldSubPathArrayOfValues struct {
	Organization_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ Organization_FieldPathArrayOfValues = (*Organization_FieldSubPathArrayOfValues)(nil)

func (fpsaov *Organization_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *Organization_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (ntt_meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(ntt_meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
func (fpsaov *Organization_FieldSubPathArrayOfValues) AsMultiRegionPolicyPathArrayOfValues() (policy.Policy_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(policy.Policy_FieldPathArrayOfValues)
	return res, ok
}
