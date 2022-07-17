// Code generated by protoc-gen-goten-object
// File: edgelq/limits/proto/v1alpha2/plan_assignment.proto
// DO NOT EDIT!!!

package plan_assignment

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
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	accepted_plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/accepted_plan"
	common "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/common"
	plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/plan"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
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
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &accepted_plan.AcceptedPlan{}
	_ = &common.Allowance{}
	_ = &plan.Plan{}
	_ = &meta_service.Service{}
)

// FieldPath provides implementation to handle
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto
type PlanAssignment_FieldPath interface {
	gotenobject.FieldPath
	Selector() PlanAssignment_FieldPathSelector
	Get(source *PlanAssignment) []interface{}
	GetSingle(source *PlanAssignment) (interface{}, bool)
	ClearValue(item *PlanAssignment)

	// Those methods build corresponding PlanAssignment_FieldPathValue
	// (or array of values) and holds passed value. Panics if injected type is incorrect.
	WithIValue(value interface{}) PlanAssignment_FieldPathValue
	WithIArrayOfValues(values interface{}) PlanAssignment_FieldPathArrayOfValues
	WithIArrayItemValue(value interface{}) PlanAssignment_FieldPathArrayItemValue
}

type PlanAssignment_FieldPathSelector int32

const (
	PlanAssignment_FieldPathSelectorName                  PlanAssignment_FieldPathSelector = 0
	PlanAssignment_FieldPathSelectorPlan                  PlanAssignment_FieldPathSelector = 1
	PlanAssignment_FieldPathSelectorService               PlanAssignment_FieldPathSelector = 2
	PlanAssignment_FieldPathSelectorExtensions            PlanAssignment_FieldPathSelector = 3
	PlanAssignment_FieldPathSelectorRegionalDistributions PlanAssignment_FieldPathSelector = 4
	PlanAssignment_FieldPathSelectorSource                PlanAssignment_FieldPathSelector = 5
	PlanAssignment_FieldPathSelectorMetadata              PlanAssignment_FieldPathSelector = 6
)

func (s PlanAssignment_FieldPathSelector) String() string {
	switch s {
	case PlanAssignment_FieldPathSelectorName:
		return "name"
	case PlanAssignment_FieldPathSelectorPlan:
		return "plan"
	case PlanAssignment_FieldPathSelectorService:
		return "service"
	case PlanAssignment_FieldPathSelectorExtensions:
		return "extensions"
	case PlanAssignment_FieldPathSelectorRegionalDistributions:
		return "regional_distributions"
	case PlanAssignment_FieldPathSelectorSource:
		return "source"
	case PlanAssignment_FieldPathSelectorMetadata:
		return "metadata"
	default:
		panic(fmt.Sprintf("Invalid selector for PlanAssignment: %d", s))
	}
}

func BuildPlanAssignment_FieldPath(fp gotenobject.RawFieldPath) (PlanAssignment_FieldPath, error) {
	if len(fp) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty field path for object PlanAssignment")
	}
	if len(fp) == 1 {
		switch fp[0] {
		case "name":
			return &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorName}, nil
		case "plan":
			return &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorPlan}, nil
		case "service":
			return &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorService}, nil
		case "extensions":
			return &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorExtensions}, nil
		case "regional_distributions", "regionalDistributions", "regional-distributions":
			return &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorRegionalDistributions}, nil
		case "source":
			return &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorSource}, nil
		case "metadata":
			return &PlanAssignment_FieldTerminalPath{selector: PlanAssignment_FieldPathSelectorMetadata}, nil
		}
	} else {
		switch fp[0] {
		case "extensions":
			if subpath, err := common.BuildAllowance_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &PlanAssignment_FieldSubPath{selector: PlanAssignment_FieldPathSelectorExtensions, subPath: subpath}, nil
			}
		case "regional_distributions", "regionalDistributions", "regional-distributions":
			if subpath, err := common.BuildRegionalDistribution_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &PlanAssignment_FieldSubPath{selector: PlanAssignment_FieldPathSelectorRegionalDistributions, subPath: subpath}, nil
			}
		case "metadata":
			if subpath, err := ntt_meta.BuildMeta_FieldPath(fp[1:]); err != nil {
				return nil, err
			} else {
				return &PlanAssignment_FieldSubPath{selector: PlanAssignment_FieldPathSelectorMetadata, subPath: subpath}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "unknown field path '%s' for object PlanAssignment", fp)
}

func ParsePlanAssignment_FieldPath(rawField string) (PlanAssignment_FieldPath, error) {
	fp, err := gotenobject.ParseRawFieldPath(rawField)
	if err != nil {
		return nil, err
	}
	return BuildPlanAssignment_FieldPath(fp)
}

func MustParsePlanAssignment_FieldPath(rawField string) PlanAssignment_FieldPath {
	fp, err := ParsePlanAssignment_FieldPath(rawField)
	if err != nil {
		panic(err)
	}
	return fp
}

type PlanAssignment_FieldTerminalPath struct {
	selector PlanAssignment_FieldPathSelector
}

var _ PlanAssignment_FieldPath = (*PlanAssignment_FieldTerminalPath)(nil)

func (fp *PlanAssignment_FieldTerminalPath) Selector() PlanAssignment_FieldPathSelector {
	return fp.selector
}

// String returns path representation in proto convention
func (fp *PlanAssignment_FieldTerminalPath) String() string {
	return fp.selector.String()
}

// JSONString returns path representation is JSON convention
func (fp *PlanAssignment_FieldTerminalPath) JSONString() string {
	return strcase.ToLowerCamel(fp.String())
}

// Get returns all values pointed by specific field from source PlanAssignment
func (fp *PlanAssignment_FieldTerminalPath) Get(source *PlanAssignment) (values []interface{}) {
	if source != nil {
		switch fp.selector {
		case PlanAssignment_FieldPathSelectorName:
			if source.Name != nil {
				values = append(values, source.Name)
			}
		case PlanAssignment_FieldPathSelectorPlan:
			if source.Plan != nil {
				values = append(values, source.Plan)
			}
		case PlanAssignment_FieldPathSelectorService:
			if source.Service != nil {
				values = append(values, source.Service)
			}
		case PlanAssignment_FieldPathSelectorExtensions:
			for _, value := range source.GetExtensions() {
				values = append(values, value)
			}
		case PlanAssignment_FieldPathSelectorRegionalDistributions:
			for _, value := range source.GetRegionalDistributions() {
				values = append(values, value)
			}
		case PlanAssignment_FieldPathSelectorSource:
			if source.Source != nil {
				values = append(values, source.Source)
			}
		case PlanAssignment_FieldPathSelectorMetadata:
			if source.Metadata != nil {
				values = append(values, source.Metadata)
			}
		default:
			panic(fmt.Sprintf("Invalid selector for PlanAssignment: %d", fp.selector))
		}
	}
	return
}

func (fp *PlanAssignment_FieldTerminalPath) GetRaw(source proto.Message) []interface{} {
	return fp.Get(source.(*PlanAssignment))
}

// GetSingle returns value pointed by specific field of from source PlanAssignment
func (fp *PlanAssignment_FieldTerminalPath) GetSingle(source *PlanAssignment) (interface{}, bool) {
	switch fp.selector {
	case PlanAssignment_FieldPathSelectorName:
		res := source.GetName()
		return res, res != nil
	case PlanAssignment_FieldPathSelectorPlan:
		res := source.GetPlan()
		return res, res != nil
	case PlanAssignment_FieldPathSelectorService:
		res := source.GetService()
		return res, res != nil
	case PlanAssignment_FieldPathSelectorExtensions:
		res := source.GetExtensions()
		return res, res != nil
	case PlanAssignment_FieldPathSelectorRegionalDistributions:
		res := source.GetRegionalDistributions()
		return res, res != nil
	case PlanAssignment_FieldPathSelectorSource:
		res := source.GetSource()
		return res, res != nil
	case PlanAssignment_FieldPathSelectorMetadata:
		res := source.GetMetadata()
		return res, res != nil
	default:
		panic(fmt.Sprintf("Invalid selector for PlanAssignment: %d", fp.selector))
	}
}

func (fp *PlanAssignment_FieldTerminalPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fp.GetSingle(source.(*PlanAssignment))
}

// GetDefault returns a default value of the field type
func (fp *PlanAssignment_FieldTerminalPath) GetDefault() interface{} {
	switch fp.selector {
	case PlanAssignment_FieldPathSelectorName:
		return (*Name)(nil)
	case PlanAssignment_FieldPathSelectorPlan:
		return (*plan.Reference)(nil)
	case PlanAssignment_FieldPathSelectorService:
		return (*meta_service.Reference)(nil)
	case PlanAssignment_FieldPathSelectorExtensions:
		return ([]*common.Allowance)(nil)
	case PlanAssignment_FieldPathSelectorRegionalDistributions:
		return ([]*common.RegionalDistribution)(nil)
	case PlanAssignment_FieldPathSelectorSource:
		return (*accepted_plan.Reference)(nil)
	case PlanAssignment_FieldPathSelectorMetadata:
		return (*ntt_meta.Meta)(nil)
	default:
		panic(fmt.Sprintf("Invalid selector for PlanAssignment: %d", fp.selector))
	}
}

func (fp *PlanAssignment_FieldTerminalPath) ClearValue(item *PlanAssignment) {
	if item != nil {
		switch fp.selector {
		case PlanAssignment_FieldPathSelectorName:
			item.Name = nil
		case PlanAssignment_FieldPathSelectorPlan:
			item.Plan = nil
		case PlanAssignment_FieldPathSelectorService:
			item.Service = nil
		case PlanAssignment_FieldPathSelectorExtensions:
			item.Extensions = nil
		case PlanAssignment_FieldPathSelectorRegionalDistributions:
			item.RegionalDistributions = nil
		case PlanAssignment_FieldPathSelectorSource:
			item.Source = nil
		case PlanAssignment_FieldPathSelectorMetadata:
			item.Metadata = nil
		default:
			panic(fmt.Sprintf("Invalid selector for PlanAssignment: %d", fp.selector))
		}
	}
}

func (fp *PlanAssignment_FieldTerminalPath) ClearValueRaw(item proto.Message) {
	fp.ClearValue(item.(*PlanAssignment))
}

// IsLeaf - whether field path is holds simple value
func (fp *PlanAssignment_FieldTerminalPath) IsLeaf() bool {
	return fp.selector == PlanAssignment_FieldPathSelectorName ||
		fp.selector == PlanAssignment_FieldPathSelectorPlan ||
		fp.selector == PlanAssignment_FieldPathSelectorService ||
		fp.selector == PlanAssignment_FieldPathSelectorSource
}

func (fp *PlanAssignment_FieldTerminalPath) WithIValue(value interface{}) PlanAssignment_FieldPathValue {
	switch fp.selector {
	case PlanAssignment_FieldPathSelectorName:
		return &PlanAssignment_FieldTerminalPathValue{PlanAssignment_FieldTerminalPath: *fp, value: value.(*Name)}
	case PlanAssignment_FieldPathSelectorPlan:
		return &PlanAssignment_FieldTerminalPathValue{PlanAssignment_FieldTerminalPath: *fp, value: value.(*plan.Reference)}
	case PlanAssignment_FieldPathSelectorService:
		return &PlanAssignment_FieldTerminalPathValue{PlanAssignment_FieldTerminalPath: *fp, value: value.(*meta_service.Reference)}
	case PlanAssignment_FieldPathSelectorExtensions:
		return &PlanAssignment_FieldTerminalPathValue{PlanAssignment_FieldTerminalPath: *fp, value: value.([]*common.Allowance)}
	case PlanAssignment_FieldPathSelectorRegionalDistributions:
		return &PlanAssignment_FieldTerminalPathValue{PlanAssignment_FieldTerminalPath: *fp, value: value.([]*common.RegionalDistribution)}
	case PlanAssignment_FieldPathSelectorSource:
		return &PlanAssignment_FieldTerminalPathValue{PlanAssignment_FieldTerminalPath: *fp, value: value.(*accepted_plan.Reference)}
	case PlanAssignment_FieldPathSelectorMetadata:
		return &PlanAssignment_FieldTerminalPathValue{PlanAssignment_FieldTerminalPath: *fp, value: value.(*ntt_meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for PlanAssignment: %d", fp.selector))
	}
}

func (fp *PlanAssignment_FieldTerminalPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fp.WithIValue(value)
}

func (fp *PlanAssignment_FieldTerminalPath) WithIArrayOfValues(values interface{}) PlanAssignment_FieldPathArrayOfValues {
	fpaov := &PlanAssignment_FieldTerminalPathArrayOfValues{PlanAssignment_FieldTerminalPath: *fp}
	switch fp.selector {
	case PlanAssignment_FieldPathSelectorName:
		return &PlanAssignment_FieldTerminalPathArrayOfValues{PlanAssignment_FieldTerminalPath: *fp, values: values.([]*Name)}
	case PlanAssignment_FieldPathSelectorPlan:
		return &PlanAssignment_FieldTerminalPathArrayOfValues{PlanAssignment_FieldTerminalPath: *fp, values: values.([]*plan.Reference)}
	case PlanAssignment_FieldPathSelectorService:
		return &PlanAssignment_FieldTerminalPathArrayOfValues{PlanAssignment_FieldTerminalPath: *fp, values: values.([]*meta_service.Reference)}
	case PlanAssignment_FieldPathSelectorExtensions:
		return &PlanAssignment_FieldTerminalPathArrayOfValues{PlanAssignment_FieldTerminalPath: *fp, values: values.([][]*common.Allowance)}
	case PlanAssignment_FieldPathSelectorRegionalDistributions:
		return &PlanAssignment_FieldTerminalPathArrayOfValues{PlanAssignment_FieldTerminalPath: *fp, values: values.([][]*common.RegionalDistribution)}
	case PlanAssignment_FieldPathSelectorSource:
		return &PlanAssignment_FieldTerminalPathArrayOfValues{PlanAssignment_FieldTerminalPath: *fp, values: values.([]*accepted_plan.Reference)}
	case PlanAssignment_FieldPathSelectorMetadata:
		return &PlanAssignment_FieldTerminalPathArrayOfValues{PlanAssignment_FieldTerminalPath: *fp, values: values.([]*ntt_meta.Meta)}
	default:
		panic(fmt.Sprintf("Invalid selector for PlanAssignment: %d", fp.selector))
	}
	return fpaov
}

func (fp *PlanAssignment_FieldTerminalPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fp.WithIArrayOfValues(values)
}

func (fp *PlanAssignment_FieldTerminalPath) WithIArrayItemValue(value interface{}) PlanAssignment_FieldPathArrayItemValue {
	switch fp.selector {
	case PlanAssignment_FieldPathSelectorExtensions:
		return &PlanAssignment_FieldTerminalPathArrayItemValue{PlanAssignment_FieldTerminalPath: *fp, value: value.(*common.Allowance)}
	case PlanAssignment_FieldPathSelectorRegionalDistributions:
		return &PlanAssignment_FieldTerminalPathArrayItemValue{PlanAssignment_FieldTerminalPath: *fp, value: value.(*common.RegionalDistribution)}
	default:
		panic(fmt.Sprintf("Invalid selector for PlanAssignment: %d", fp.selector))
	}
}

func (fp *PlanAssignment_FieldTerminalPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fp.WithIArrayItemValue(value)
}

type PlanAssignment_FieldSubPath struct {
	selector PlanAssignment_FieldPathSelector
	subPath  gotenobject.FieldPath
}

var _ PlanAssignment_FieldPath = (*PlanAssignment_FieldSubPath)(nil)

func (fps *PlanAssignment_FieldSubPath) Selector() PlanAssignment_FieldPathSelector {
	return fps.selector
}
func (fps *PlanAssignment_FieldSubPath) AsExtensionsSubPath() (common.Allowance_FieldPath, bool) {
	res, ok := fps.subPath.(common.Allowance_FieldPath)
	return res, ok
}
func (fps *PlanAssignment_FieldSubPath) AsRegionalDistributionsSubPath() (common.RegionalDistribution_FieldPath, bool) {
	res, ok := fps.subPath.(common.RegionalDistribution_FieldPath)
	return res, ok
}
func (fps *PlanAssignment_FieldSubPath) AsMetadataSubPath() (ntt_meta.Meta_FieldPath, bool) {
	res, ok := fps.subPath.(ntt_meta.Meta_FieldPath)
	return res, ok
}

// String returns path representation in proto convention
func (fps *PlanAssignment_FieldSubPath) String() string {
	return fps.selector.String() + "." + fps.subPath.String()
}

// JSONString returns path representation is JSON convention
func (fps *PlanAssignment_FieldSubPath) JSONString() string {
	return strcase.ToLowerCamel(fps.selector.String()) + "." + fps.subPath.JSONString()
}

// Get returns all values pointed by selected field from source PlanAssignment
func (fps *PlanAssignment_FieldSubPath) Get(source *PlanAssignment) (values []interface{}) {
	if asAllowanceFieldPath, ok := fps.AsExtensionsSubPath(); ok {
		for _, item := range source.GetExtensions() {
			values = append(values, asAllowanceFieldPath.Get(item)...)
		}
	} else if asRegionalDistributionFieldPath, ok := fps.AsRegionalDistributionsSubPath(); ok {
		for _, item := range source.GetRegionalDistributions() {
			values = append(values, asRegionalDistributionFieldPath.Get(item)...)
		}
	} else if asMetaFieldPath, ok := fps.AsMetadataSubPath(); ok {
		values = append(values, asMetaFieldPath.Get(source.GetMetadata())...)
	} else {
		panic(fmt.Sprintf("Invalid selector for PlanAssignment: %d", fps.selector))
	}
	return
}

func (fps *PlanAssignment_FieldSubPath) GetRaw(source proto.Message) []interface{} {
	return fps.Get(source.(*PlanAssignment))
}

// GetSingle returns value of selected field from source PlanAssignment
func (fps *PlanAssignment_FieldSubPath) GetSingle(source *PlanAssignment) (interface{}, bool) {
	switch fps.selector {
	case PlanAssignment_FieldPathSelectorExtensions:
		if len(source.GetExtensions()) == 0 {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetExtensions()[0])
	case PlanAssignment_FieldPathSelectorRegionalDistributions:
		if len(source.GetRegionalDistributions()) == 0 {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetRegionalDistributions()[0])
	case PlanAssignment_FieldPathSelectorMetadata:
		if source.GetMetadata() == nil {
			return nil, false
		}
		return fps.subPath.GetSingleRaw(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for PlanAssignment: %d", fps.selector))
	}
}

func (fps *PlanAssignment_FieldSubPath) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fps.GetSingle(source.(*PlanAssignment))
}

// GetDefault returns a default value of the field type
func (fps *PlanAssignment_FieldSubPath) GetDefault() interface{} {
	return fps.subPath.GetDefault()
}

func (fps *PlanAssignment_FieldSubPath) ClearValue(item *PlanAssignment) {
	if item != nil {
		switch fps.selector {
		case PlanAssignment_FieldPathSelectorExtensions:
			for _, subItem := range item.Extensions {
				fps.subPath.ClearValueRaw(subItem)
			}
		case PlanAssignment_FieldPathSelectorRegionalDistributions:
			for _, subItem := range item.RegionalDistributions {
				fps.subPath.ClearValueRaw(subItem)
			}
		case PlanAssignment_FieldPathSelectorMetadata:
			fps.subPath.ClearValueRaw(item.Metadata)
		default:
			panic(fmt.Sprintf("Invalid selector for PlanAssignment: %d", fps.selector))
		}
	}
}

func (fps *PlanAssignment_FieldSubPath) ClearValueRaw(item proto.Message) {
	fps.ClearValue(item.(*PlanAssignment))
}

// IsLeaf - whether field path is holds simple value
func (fps *PlanAssignment_FieldSubPath) IsLeaf() bool {
	return fps.subPath.IsLeaf()
}

func (fps *PlanAssignment_FieldSubPath) WithIValue(value interface{}) PlanAssignment_FieldPathValue {
	return &PlanAssignment_FieldSubPathValue{fps, fps.subPath.WithRawIValue(value)}
}

func (fps *PlanAssignment_FieldSubPath) WithRawIValue(value interface{}) gotenobject.FieldPathValue {
	return fps.WithIValue(value)
}

func (fps *PlanAssignment_FieldSubPath) WithIArrayOfValues(values interface{}) PlanAssignment_FieldPathArrayOfValues {
	return &PlanAssignment_FieldSubPathArrayOfValues{fps, fps.subPath.WithRawIArrayOfValues(values)}
}

func (fps *PlanAssignment_FieldSubPath) WithRawIArrayOfValues(values interface{}) gotenobject.FieldPathArrayOfValues {
	return fps.WithIArrayOfValues(values)
}

func (fps *PlanAssignment_FieldSubPath) WithIArrayItemValue(value interface{}) PlanAssignment_FieldPathArrayItemValue {
	return &PlanAssignment_FieldSubPathArrayItemValue{fps, fps.subPath.WithRawIArrayItemValue(value)}
}

func (fps *PlanAssignment_FieldSubPath) WithRawIArrayItemValue(value interface{}) gotenobject.FieldPathArrayItemValue {
	return fps.WithIArrayItemValue(value)
}

// PlanAssignment_FieldPathValue allows storing values for PlanAssignment fields according to their type
type PlanAssignment_FieldPathValue interface {
	PlanAssignment_FieldPath
	gotenobject.FieldPathValue
	SetTo(target **PlanAssignment)
	CompareWith(*PlanAssignment) (cmp int, comparable bool)
}

func ParsePlanAssignment_FieldPathValue(pathStr, valueStr string) (PlanAssignment_FieldPathValue, error) {
	fp, err := ParsePlanAssignment_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpv, err := gotenobject.ParseFieldPathValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing PlanAssignment field path value from %s: %v", valueStr, err)
	}
	return fpv.(PlanAssignment_FieldPathValue), nil
}

func MustParsePlanAssignment_FieldPathValue(pathStr, valueStr string) PlanAssignment_FieldPathValue {
	fpv, err := ParsePlanAssignment_FieldPathValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpv
}

type PlanAssignment_FieldTerminalPathValue struct {
	PlanAssignment_FieldTerminalPath
	value interface{}
}

var _ PlanAssignment_FieldPathValue = (*PlanAssignment_FieldTerminalPathValue)(nil)

// GetRawValue returns raw value stored under selected path for 'PlanAssignment' as interface{}
func (fpv *PlanAssignment_FieldTerminalPathValue) GetRawValue() interface{} {
	return fpv.value
}
func (fpv *PlanAssignment_FieldTerminalPathValue) AsNameValue() (*Name, bool) {
	res, ok := fpv.value.(*Name)
	return res, ok
}
func (fpv *PlanAssignment_FieldTerminalPathValue) AsPlanValue() (*plan.Reference, bool) {
	res, ok := fpv.value.(*plan.Reference)
	return res, ok
}
func (fpv *PlanAssignment_FieldTerminalPathValue) AsServiceValue() (*meta_service.Reference, bool) {
	res, ok := fpv.value.(*meta_service.Reference)
	return res, ok
}
func (fpv *PlanAssignment_FieldTerminalPathValue) AsExtensionsValue() ([]*common.Allowance, bool) {
	res, ok := fpv.value.([]*common.Allowance)
	return res, ok
}
func (fpv *PlanAssignment_FieldTerminalPathValue) AsRegionalDistributionsValue() ([]*common.RegionalDistribution, bool) {
	res, ok := fpv.value.([]*common.RegionalDistribution)
	return res, ok
}
func (fpv *PlanAssignment_FieldTerminalPathValue) AsSourceValue() (*accepted_plan.Reference, bool) {
	res, ok := fpv.value.(*accepted_plan.Reference)
	return res, ok
}
func (fpv *PlanAssignment_FieldTerminalPathValue) AsMetadataValue() (*ntt_meta.Meta, bool) {
	res, ok := fpv.value.(*ntt_meta.Meta)
	return res, ok
}

// SetTo stores value for selected field for object PlanAssignment
func (fpv *PlanAssignment_FieldTerminalPathValue) SetTo(target **PlanAssignment) {
	if *target == nil {
		*target = new(PlanAssignment)
	}
	switch fpv.selector {
	case PlanAssignment_FieldPathSelectorName:
		(*target).Name = fpv.value.(*Name)
	case PlanAssignment_FieldPathSelectorPlan:
		(*target).Plan = fpv.value.(*plan.Reference)
	case PlanAssignment_FieldPathSelectorService:
		(*target).Service = fpv.value.(*meta_service.Reference)
	case PlanAssignment_FieldPathSelectorExtensions:
		(*target).Extensions = fpv.value.([]*common.Allowance)
	case PlanAssignment_FieldPathSelectorRegionalDistributions:
		(*target).RegionalDistributions = fpv.value.([]*common.RegionalDistribution)
	case PlanAssignment_FieldPathSelectorSource:
		(*target).Source = fpv.value.(*accepted_plan.Reference)
	case PlanAssignment_FieldPathSelectorMetadata:
		(*target).Metadata = fpv.value.(*ntt_meta.Meta)
	default:
		panic(fmt.Sprintf("Invalid selector for PlanAssignment: %d", fpv.selector))
	}
}

func (fpv *PlanAssignment_FieldTerminalPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*PlanAssignment)
	fpv.SetTo(&typedObject)
}

// CompareWith compares value in the 'PlanAssignment_FieldTerminalPathValue' with the value under path in 'PlanAssignment'.
func (fpv *PlanAssignment_FieldTerminalPathValue) CompareWith(source *PlanAssignment) (int, bool) {
	switch fpv.selector {
	case PlanAssignment_FieldPathSelectorName:
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
	case PlanAssignment_FieldPathSelectorPlan:
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
	case PlanAssignment_FieldPathSelectorService:
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
	case PlanAssignment_FieldPathSelectorExtensions:
		return 0, false
	case PlanAssignment_FieldPathSelectorRegionalDistributions:
		return 0, false
	case PlanAssignment_FieldPathSelectorSource:
		leftValue := fpv.value.(*accepted_plan.Reference)
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
	case PlanAssignment_FieldPathSelectorMetadata:
		return 0, false
	default:
		panic(fmt.Sprintf("Invalid selector for PlanAssignment: %d", fpv.selector))
	}
}

func (fpv *PlanAssignment_FieldTerminalPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpv.CompareWith(source.(*PlanAssignment))
}

type PlanAssignment_FieldSubPathValue struct {
	PlanAssignment_FieldPath
	subPathValue gotenobject.FieldPathValue
}

var _ PlanAssignment_FieldPathValue = (*PlanAssignment_FieldSubPathValue)(nil)

func (fpvs *PlanAssignment_FieldSubPathValue) AsExtensionsPathValue() (common.Allowance_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(common.Allowance_FieldPathValue)
	return res, ok
}
func (fpvs *PlanAssignment_FieldSubPathValue) AsRegionalDistributionsPathValue() (common.RegionalDistribution_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(common.RegionalDistribution_FieldPathValue)
	return res, ok
}
func (fpvs *PlanAssignment_FieldSubPathValue) AsMetadataPathValue() (ntt_meta.Meta_FieldPathValue, bool) {
	res, ok := fpvs.subPathValue.(ntt_meta.Meta_FieldPathValue)
	return res, ok
}

func (fpvs *PlanAssignment_FieldSubPathValue) SetTo(target **PlanAssignment) {
	if *target == nil {
		*target = new(PlanAssignment)
	}
	switch fpvs.Selector() {
	case PlanAssignment_FieldPathSelectorExtensions:
		panic("FieldPath setter is unsupported for array subpaths")
	case PlanAssignment_FieldPathSelectorRegionalDistributions:
		panic("FieldPath setter is unsupported for array subpaths")
	case PlanAssignment_FieldPathSelectorMetadata:
		fpvs.subPathValue.(ntt_meta.Meta_FieldPathValue).SetTo(&(*target).Metadata)
	default:
		panic(fmt.Sprintf("Invalid selector for PlanAssignment: %d", fpvs.Selector()))
	}
}

func (fpvs *PlanAssignment_FieldSubPathValue) SetToRaw(target proto.Message) {
	typedObject := target.(*PlanAssignment)
	fpvs.SetTo(&typedObject)
}

func (fpvs *PlanAssignment_FieldSubPathValue) GetRawValue() interface{} {
	return fpvs.subPathValue.GetRawValue()
}

func (fpvs *PlanAssignment_FieldSubPathValue) CompareWith(source *PlanAssignment) (int, bool) {
	switch fpvs.Selector() {
	case PlanAssignment_FieldPathSelectorExtensions:
		return 0, false // repeated field
	case PlanAssignment_FieldPathSelectorRegionalDistributions:
		return 0, false // repeated field
	case PlanAssignment_FieldPathSelectorMetadata:
		return fpvs.subPathValue.(ntt_meta.Meta_FieldPathValue).CompareWith(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for PlanAssignment: %d", fpvs.Selector()))
	}
}

func (fpvs *PlanAssignment_FieldSubPathValue) CompareWithRaw(source proto.Message) (int, bool) {
	return fpvs.CompareWith(source.(*PlanAssignment))
}

// PlanAssignment_FieldPathArrayItemValue allows storing single item in Path-specific values for PlanAssignment according to their type
// Present only for array (repeated) types.
type PlanAssignment_FieldPathArrayItemValue interface {
	gotenobject.FieldPathArrayItemValue
	PlanAssignment_FieldPath
	ContainsValue(*PlanAssignment) bool
}

// ParsePlanAssignment_FieldPathArrayItemValue parses string and JSON-encoded value to its Value
func ParsePlanAssignment_FieldPathArrayItemValue(pathStr, valueStr string) (PlanAssignment_FieldPathArrayItemValue, error) {
	fp, err := ParsePlanAssignment_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaiv, err := gotenobject.ParseFieldPathArrayItemValue(fp, valueStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing PlanAssignment field path array item value from %s: %v", valueStr, err)
	}
	return fpaiv.(PlanAssignment_FieldPathArrayItemValue), nil
}

func MustParsePlanAssignment_FieldPathArrayItemValue(pathStr, valueStr string) PlanAssignment_FieldPathArrayItemValue {
	fpaiv, err := ParsePlanAssignment_FieldPathArrayItemValue(pathStr, valueStr)
	if err != nil {
		panic(err)
	}
	return fpaiv
}

type PlanAssignment_FieldTerminalPathArrayItemValue struct {
	PlanAssignment_FieldTerminalPath
	value interface{}
}

var _ PlanAssignment_FieldPathArrayItemValue = (*PlanAssignment_FieldTerminalPathArrayItemValue)(nil)

// GetRawValue returns stored element value for array in object PlanAssignment as interface{}
func (fpaiv *PlanAssignment_FieldTerminalPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaiv.value
}
func (fpaiv *PlanAssignment_FieldTerminalPathArrayItemValue) AsExtensionsItemValue() (*common.Allowance, bool) {
	res, ok := fpaiv.value.(*common.Allowance)
	return res, ok
}
func (fpaiv *PlanAssignment_FieldTerminalPathArrayItemValue) AsRegionalDistributionsItemValue() (*common.RegionalDistribution, bool) {
	res, ok := fpaiv.value.(*common.RegionalDistribution)
	return res, ok
}

func (fpaiv *PlanAssignment_FieldTerminalPathArrayItemValue) GetSingle(source *PlanAssignment) (interface{}, bool) {
	return nil, false
}

func (fpaiv *PlanAssignment_FieldTerminalPathArrayItemValue) GetSingleRaw(source proto.Message) (interface{}, bool) {
	return fpaiv.GetSingle(source.(*PlanAssignment))
}

// Contains returns a boolean indicating if value that is being held is present in given 'PlanAssignment'
func (fpaiv *PlanAssignment_FieldTerminalPathArrayItemValue) ContainsValue(source *PlanAssignment) bool {
	slice := fpaiv.PlanAssignment_FieldTerminalPath.Get(source)
	for _, v := range slice {
		if reflect.DeepEqual(v, fpaiv.value) {
			return true
		}
	}
	return false
}

type PlanAssignment_FieldSubPathArrayItemValue struct {
	PlanAssignment_FieldPath
	subPathItemValue gotenobject.FieldPathArrayItemValue
}

// GetRawValue returns stored array item value
func (fpaivs *PlanAssignment_FieldSubPathArrayItemValue) GetRawItemValue() interface{} {
	return fpaivs.subPathItemValue.GetRawItemValue()
}
func (fpaivs *PlanAssignment_FieldSubPathArrayItemValue) AsExtensionsPathItemValue() (common.Allowance_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(common.Allowance_FieldPathArrayItemValue)
	return res, ok
}
func (fpaivs *PlanAssignment_FieldSubPathArrayItemValue) AsRegionalDistributionsPathItemValue() (common.RegionalDistribution_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(common.RegionalDistribution_FieldPathArrayItemValue)
	return res, ok
}
func (fpaivs *PlanAssignment_FieldSubPathArrayItemValue) AsMetadataPathItemValue() (ntt_meta.Meta_FieldPathArrayItemValue, bool) {
	res, ok := fpaivs.subPathItemValue.(ntt_meta.Meta_FieldPathArrayItemValue)
	return res, ok
}

// Contains returns a boolean indicating if value that is being held is present in given 'PlanAssignment'
func (fpaivs *PlanAssignment_FieldSubPathArrayItemValue) ContainsValue(source *PlanAssignment) bool {
	switch fpaivs.Selector() {
	case PlanAssignment_FieldPathSelectorExtensions:
		return false // repeated/map field
	case PlanAssignment_FieldPathSelectorRegionalDistributions:
		return false // repeated/map field
	case PlanAssignment_FieldPathSelectorMetadata:
		return fpaivs.subPathItemValue.(ntt_meta.Meta_FieldPathArrayItemValue).ContainsValue(source.GetMetadata())
	default:
		panic(fmt.Sprintf("Invalid selector for PlanAssignment: %d", fpaivs.Selector()))
	}
}

// PlanAssignment_FieldPathArrayOfValues allows storing slice of values for PlanAssignment fields according to their type
type PlanAssignment_FieldPathArrayOfValues interface {
	gotenobject.FieldPathArrayOfValues
	PlanAssignment_FieldPath
}

func ParsePlanAssignment_FieldPathArrayOfValues(pathStr, valuesStr string) (PlanAssignment_FieldPathArrayOfValues, error) {
	fp, err := ParsePlanAssignment_FieldPath(pathStr)
	if err != nil {
		return nil, err
	}
	fpaov, err := gotenobject.ParseFieldPathArrayOfValues(fp, valuesStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error parsing PlanAssignment field path array of values from %s: %v", valuesStr, err)
	}
	return fpaov.(PlanAssignment_FieldPathArrayOfValues), nil
}

func MustParsePlanAssignment_FieldPathArrayOfValues(pathStr, valuesStr string) PlanAssignment_FieldPathArrayOfValues {
	fpaov, err := ParsePlanAssignment_FieldPathArrayOfValues(pathStr, valuesStr)
	if err != nil {
		panic(err)
	}
	return fpaov
}

type PlanAssignment_FieldTerminalPathArrayOfValues struct {
	PlanAssignment_FieldTerminalPath
	values interface{}
}

var _ PlanAssignment_FieldPathArrayOfValues = (*PlanAssignment_FieldTerminalPathArrayOfValues)(nil)

func (fpaov *PlanAssignment_FieldTerminalPathArrayOfValues) GetRawValues() (values []interface{}) {
	switch fpaov.selector {
	case PlanAssignment_FieldPathSelectorName:
		for _, v := range fpaov.values.([]*Name) {
			values = append(values, v)
		}
	case PlanAssignment_FieldPathSelectorPlan:
		for _, v := range fpaov.values.([]*plan.Reference) {
			values = append(values, v)
		}
	case PlanAssignment_FieldPathSelectorService:
		for _, v := range fpaov.values.([]*meta_service.Reference) {
			values = append(values, v)
		}
	case PlanAssignment_FieldPathSelectorExtensions:
		for _, v := range fpaov.values.([][]*common.Allowance) {
			values = append(values, v)
		}
	case PlanAssignment_FieldPathSelectorRegionalDistributions:
		for _, v := range fpaov.values.([][]*common.RegionalDistribution) {
			values = append(values, v)
		}
	case PlanAssignment_FieldPathSelectorSource:
		for _, v := range fpaov.values.([]*accepted_plan.Reference) {
			values = append(values, v)
		}
	case PlanAssignment_FieldPathSelectorMetadata:
		for _, v := range fpaov.values.([]*ntt_meta.Meta) {
			values = append(values, v)
		}
	}
	return
}
func (fpaov *PlanAssignment_FieldTerminalPathArrayOfValues) AsNameArrayOfValues() ([]*Name, bool) {
	res, ok := fpaov.values.([]*Name)
	return res, ok
}
func (fpaov *PlanAssignment_FieldTerminalPathArrayOfValues) AsPlanArrayOfValues() ([]*plan.Reference, bool) {
	res, ok := fpaov.values.([]*plan.Reference)
	return res, ok
}
func (fpaov *PlanAssignment_FieldTerminalPathArrayOfValues) AsServiceArrayOfValues() ([]*meta_service.Reference, bool) {
	res, ok := fpaov.values.([]*meta_service.Reference)
	return res, ok
}
func (fpaov *PlanAssignment_FieldTerminalPathArrayOfValues) AsExtensionsArrayOfValues() ([][]*common.Allowance, bool) {
	res, ok := fpaov.values.([][]*common.Allowance)
	return res, ok
}
func (fpaov *PlanAssignment_FieldTerminalPathArrayOfValues) AsRegionalDistributionsArrayOfValues() ([][]*common.RegionalDistribution, bool) {
	res, ok := fpaov.values.([][]*common.RegionalDistribution)
	return res, ok
}
func (fpaov *PlanAssignment_FieldTerminalPathArrayOfValues) AsSourceArrayOfValues() ([]*accepted_plan.Reference, bool) {
	res, ok := fpaov.values.([]*accepted_plan.Reference)
	return res, ok
}
func (fpaov *PlanAssignment_FieldTerminalPathArrayOfValues) AsMetadataArrayOfValues() ([]*ntt_meta.Meta, bool) {
	res, ok := fpaov.values.([]*ntt_meta.Meta)
	return res, ok
}

type PlanAssignment_FieldSubPathArrayOfValues struct {
	PlanAssignment_FieldPath
	subPathArrayOfValues gotenobject.FieldPathArrayOfValues
}

var _ PlanAssignment_FieldPathArrayOfValues = (*PlanAssignment_FieldSubPathArrayOfValues)(nil)

func (fpsaov *PlanAssignment_FieldSubPathArrayOfValues) GetRawValues() []interface{} {
	return fpsaov.subPathArrayOfValues.GetRawValues()
}
func (fpsaov *PlanAssignment_FieldSubPathArrayOfValues) AsExtensionsPathArrayOfValues() (common.Allowance_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(common.Allowance_FieldPathArrayOfValues)
	return res, ok
}
func (fpsaov *PlanAssignment_FieldSubPathArrayOfValues) AsRegionalDistributionsPathArrayOfValues() (common.RegionalDistribution_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(common.RegionalDistribution_FieldPathArrayOfValues)
	return res, ok
}
func (fpsaov *PlanAssignment_FieldSubPathArrayOfValues) AsMetadataPathArrayOfValues() (ntt_meta.Meta_FieldPathArrayOfValues, bool) {
	res, ok := fpsaov.subPathArrayOfValues.(ntt_meta.Meta_FieldPathArrayOfValues)
	return res, ok
}
