// Code generated by protoc-gen-goten-resource
// Resource: Log
// DO NOT EDIT!!!

package log

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	gotenfilter "github.com/cloudwan/goten-sdk/runtime/resource/filter"
)

// proto imports
import (
	iam_iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1/common"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	bucket "github.com/cloudwan/edgelq-sdk/logging/resources/v1/bucket"
	common "github.com/cloudwan/edgelq-sdk/logging/resources/v1/common"
	log_descriptor "github.com/cloudwan/edgelq-sdk/logging/resources/v1/log_descriptor"
	meta_common "github.com/cloudwan/goten-sdk/meta-service/resources/v1/common"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	multi_region_policy "github.com/cloudwan/goten-sdk/types/multi_region_policy"
	anypb "google.golang.org/protobuf/types/known/anypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// ensure the imports are used
var (
	_ = gotenresource.ConditionContainsTypeAll
	_ = gotenfilter.AND
)

// make sure we're using proto imports
var (
	_ = &iam_iam_common.PCR{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &bucket.Bucket{}
	_ = &common.LabelDescriptor{}
	_ = &log_descriptor.LogDescriptor{}
	_ = &anypb.Any{}
	_ = &structpb.Struct{}
	_ = &timestamppb.Timestamp{}
	_ = &meta_common.LabelledDomain{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

type FilterBuilderOrCondition interface {
	_IsLogFilterBuilderOrCondition()
}

type FilterBuilder struct {
	conds  []FilterCondition
	useNot bool
	op     gotenfilter.CompositeOperator
}

func NewFilterBuilder() *FilterBuilder {
	return NewAndFilterBuilder()
}

func NewAndFilterBuilder() *FilterBuilder {
	return &FilterBuilder{
		op: gotenfilter.AND,
	}
}

func NewOrFilterBuilder() *FilterBuilder {
	return &FilterBuilder{
		op: gotenfilter.OR,
	}
}

func (b *FilterBuilder) _IsLogFilterBuilderOrCondition() {}

func (b *FilterBuilder) With(condOrBuilder FilterBuilderOrCondition, opts ...gotenfilter.FilterConditionOption) *FilterBuilder {
	var cond FilterCondition
	switch typedObj := condOrBuilder.(type) {
	case *Filter:
		cond = typedObj.GetCondition()
	case *FilterBuilder:
		cond = &FilterConditionComposite{Operator: typedObj.op, Conditions: typedObj.conds}
	case FilterCondition:
		cond = typedObj
	default:
		panic("Unknown condition or builder type")
	}
	cfg := gotenfilter.MakeFilterCondOptions(opts)
	if cfg.IsNot() {
		cond = &FilterConditionNot{cond}
	}
	b.conds = append(b.conds, cond)
	return b
}

func (b *FilterBuilder) Where(opts ...gotenfilter.FilterConditionOption) *filterCndBuilder {
	cfg := gotenfilter.MakeFilterCondOptions(opts)
	b.useNot = cfg.IsNot()
	return &filterCndBuilder{builder: b}
}

func (b *FilterBuilder) WherePath(fp Log_FieldPath, opts ...gotenfilter.FilterConditionOption) *filterCndBuilderAnyPath {
	cfg := gotenfilter.MakeFilterCondOptions(opts)
	b.useNot = cfg.IsNot()
	return &filterCndBuilderAnyPath{builder: b, fp: fp}
}

func (b *FilterBuilder) Filter() *Filter {
	return &Filter{
		FilterCondition: &FilterConditionComposite{Operator: b.op, Conditions: b.conds},
	}
}

func (b *FilterBuilder) addCond(cond FilterCondition) *FilterBuilder {
	if b.useNot {
		cond = &FilterConditionNot{cond}
		b.useNot = false
	}
	b.conds = append(b.conds, cond)
	return b
}

type filterCndBuilderAnyPath struct {
	builder *FilterBuilder
	fp      Log_FieldPath
}

func (b *filterCndBuilderAnyPath) Eq(value interface{}) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderAnyPath) Neq(value interface{}) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderAnyPath) Gt(value interface{}) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderAnyPath) Gte(value interface{}) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderAnyPath) Lt(value interface{}) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderAnyPath) Lte(value interface{}) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderAnyPath) In(values interface{}) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Log_FieldPathArrayOfValues: b.fp.WithIArrayOfValues(values),
	})
}

func (b *filterCndBuilderAnyPath) NotIn(values interface{}) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Log_FieldPathArrayOfValues: b.fp.WithIArrayOfValues(values),
	})
}

func (b *filterCndBuilderAnyPath) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: b.fp,
	})
}

func (b *filterCndBuilderAnyPath) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: b.fp,
	})
}

func (b *filterCndBuilderAnyPath) Contains(value interface{}) *FilterBuilder {
	return b.builder.addCond(&FilterConditionContains{
		Type:      gotenresource.ConditionContainsTypeValue,
		FieldPath: b.fp,
		Value:     b.fp.WithIArrayItemValue(value),
	})
}

func (b *filterCndBuilderAnyPath) ContainsAnyOf(values []interface{}) *FilterBuilder {
	itemValues := make([]Log_FieldPathArrayItemValue, 0, len(values))
	for _, value := range values {
		itemValues = append(itemValues, b.fp.WithIArrayItemValue(value))
	}
	return b.builder.addCond(&FilterConditionContains{
		Type:      gotenresource.ConditionContainsTypeAny,
		FieldPath: b.fp,
		Values:    itemValues,
	})
}

func (b *filterCndBuilderAnyPath) ContainsAll(values []interface{}) *FilterBuilder {
	itemValues := make([]Log_FieldPathArrayItemValue, 0, len(values))
	for _, value := range values {
		itemValues = append(itemValues, b.fp.WithIArrayItemValue(value))
	}
	return b.builder.addCond(&FilterConditionContains{
		Type:      gotenresource.ConditionContainsTypeAll,
		FieldPath: b.fp,
		Values:    itemValues,
	})
}

func (b *filterCndBuilderAnyPath) compare(op gotenfilter.CompareOperator, value interface{}) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:           op,
		Log_FieldPathValue: b.fp.WithIValue(value),
	})
}

type filterCndBuilder struct {
	builder *FilterBuilder
}

func (b *filterCndBuilder) Name() *filterCndBuilderName {
	return &filterCndBuilderName{builder: b.builder}
}

func (b *filterCndBuilder) Scope() *filterCndBuilderScope {
	return &filterCndBuilderScope{builder: b.builder}
}

func (b *filterCndBuilder) Service() *filterCndBuilderService {
	return &filterCndBuilderService{builder: b.builder}
}

func (b *filterCndBuilder) Region() *filterCndBuilderRegion {
	return &filterCndBuilderRegion{builder: b.builder}
}

func (b *filterCndBuilder) Version() *filterCndBuilderVersion {
	return &filterCndBuilderVersion{builder: b.builder}
}

func (b *filterCndBuilder) LogDescriptor() *filterCndBuilderLogDescriptor {
	return &filterCndBuilderLogDescriptor{builder: b.builder}
}

func (b *filterCndBuilder) Labels() *filterCndBuilderLabels {
	return &filterCndBuilderLabels{builder: b.builder}
}

func (b *filterCndBuilder) Time() *filterCndBuilderTime {
	return &filterCndBuilderTime{builder: b.builder}
}

func (b *filterCndBuilder) JsonPayload() *filterCndBuilderJsonPayload {
	return &filterCndBuilderJsonPayload{builder: b.builder}
}

func (b *filterCndBuilder) PbPayload() *filterCndBuilderPbPayload {
	return &filterCndBuilderPbPayload{builder: b.builder}
}

func (b *filterCndBuilder) StringPayload() *filterCndBuilderStringPayload {
	return &filterCndBuilderStringPayload{builder: b.builder}
}

func (b *filterCndBuilder) BytesPayload() *filterCndBuilderBytesPayload {
	return &filterCndBuilderBytesPayload{builder: b.builder}
}

func (b *filterCndBuilder) BinKey() *filterCndBuilderBinKey {
	return &filterCndBuilderBinKey{builder: b.builder}
}

type filterCndBuilderName struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderName) Eq(value *Name) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderName) Neq(value *Name) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderName) Gt(value *Name) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderName) Gte(value *Name) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderName) Lt(value *Name) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderName) Lte(value *Name) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderName) In(values []*Name) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().Name().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderName) NotIn(values []*Name) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().Name().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderName) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewLogFieldPathBuilder().Name().FieldPath(),
	})
}

func (b *filterCndBuilderName) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewLogFieldPathBuilder().Name().FieldPath(),
	})
}

func (b *filterCndBuilderName) compare(op gotenfilter.CompareOperator, value *Name) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:           op,
		Log_FieldPathValue: NewLogFieldPathBuilder().Name().WithValue(value),
	})
}

type filterCndBuilderScope struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderScope) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderScope) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderScope) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderScope) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderScope) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderScope) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderScope) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().Scope().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderScope) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().Scope().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderScope) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewLogFieldPathBuilder().Scope().FieldPath(),
	})
}

func (b *filterCndBuilderScope) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewLogFieldPathBuilder().Scope().FieldPath(),
	})
}

func (b *filterCndBuilderScope) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:           op,
		Log_FieldPathValue: NewLogFieldPathBuilder().Scope().WithValue(value),
	})
}

type filterCndBuilderService struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderService) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderService) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderService) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderService) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderService) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderService) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderService) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().Service().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderService) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().Service().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderService) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewLogFieldPathBuilder().Service().FieldPath(),
	})
}

func (b *filterCndBuilderService) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewLogFieldPathBuilder().Service().FieldPath(),
	})
}

func (b *filterCndBuilderService) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:           op,
		Log_FieldPathValue: NewLogFieldPathBuilder().Service().WithValue(value),
	})
}

type filterCndBuilderRegion struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderRegion) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderRegion) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderRegion) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderRegion) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderRegion) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderRegion) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderRegion) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().Region().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderRegion) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().Region().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderRegion) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewLogFieldPathBuilder().Region().FieldPath(),
	})
}

func (b *filterCndBuilderRegion) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewLogFieldPathBuilder().Region().FieldPath(),
	})
}

func (b *filterCndBuilderRegion) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:           op,
		Log_FieldPathValue: NewLogFieldPathBuilder().Region().WithValue(value),
	})
}

type filterCndBuilderVersion struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderVersion) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderVersion) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderVersion) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderVersion) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderVersion) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderVersion) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderVersion) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().Version().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderVersion) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().Version().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderVersion) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewLogFieldPathBuilder().Version().FieldPath(),
	})
}

func (b *filterCndBuilderVersion) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewLogFieldPathBuilder().Version().FieldPath(),
	})
}

func (b *filterCndBuilderVersion) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:           op,
		Log_FieldPathValue: NewLogFieldPathBuilder().Version().WithValue(value),
	})
}

type filterCndBuilderLogDescriptor struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderLogDescriptor) Eq(value *log_descriptor.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderLogDescriptor) Neq(value *log_descriptor.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderLogDescriptor) Gt(value *log_descriptor.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderLogDescriptor) Gte(value *log_descriptor.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderLogDescriptor) Lt(value *log_descriptor.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderLogDescriptor) Lte(value *log_descriptor.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderLogDescriptor) In(values []*log_descriptor.Reference) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().LogDescriptor().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderLogDescriptor) NotIn(values []*log_descriptor.Reference) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().LogDescriptor().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderLogDescriptor) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewLogFieldPathBuilder().LogDescriptor().FieldPath(),
	})
}

func (b *filterCndBuilderLogDescriptor) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewLogFieldPathBuilder().LogDescriptor().FieldPath(),
	})
}

func (b *filterCndBuilderLogDescriptor) compare(op gotenfilter.CompareOperator, value *log_descriptor.Reference) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:           op,
		Log_FieldPathValue: NewLogFieldPathBuilder().LogDescriptor().WithValue(value),
	})
}

type filterCndBuilderLabels struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderLabels) Eq(value map[string]string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderLabels) Neq(value map[string]string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderLabels) Gt(value map[string]string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderLabels) Gte(value map[string]string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderLabels) Lt(value map[string]string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderLabels) Lte(value map[string]string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderLabels) In(values []map[string]string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().Labels().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderLabels) NotIn(values []map[string]string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().Labels().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderLabels) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewLogFieldPathBuilder().Labels().FieldPath(),
	})
}

func (b *filterCndBuilderLabels) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewLogFieldPathBuilder().Labels().FieldPath(),
	})
}

func (b *filterCndBuilderLabels) compare(op gotenfilter.CompareOperator, value map[string]string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:           op,
		Log_FieldPathValue: NewLogFieldPathBuilder().Labels().WithValue(value),
	})
}

func (b *filterCndBuilderLabels) WithKey(key string) *mapFilterCndBuilderLabels {
	return &mapFilterCndBuilderLabels{builder: b.builder, key: key}
}

type mapFilterCndBuilderLabels struct {
	builder *FilterBuilder
	key     string
}

func (b *mapFilterCndBuilderLabels) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *mapFilterCndBuilderLabels) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *mapFilterCndBuilderLabels) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *mapFilterCndBuilderLabels) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *mapFilterCndBuilderLabels) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *mapFilterCndBuilderLabels) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *mapFilterCndBuilderLabels) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().Labels().WithKey(b.key).WithArrayOfValues(values),
	})
}

func (b *mapFilterCndBuilderLabels) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().Labels().WithKey(b.key).WithArrayOfValues(values),
	})
}

func (b *mapFilterCndBuilderLabels) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewLogFieldPathBuilder().Labels().WithKey(b.key).FieldPath(),
	})
}

func (b *mapFilterCndBuilderLabels) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewLogFieldPathBuilder().Labels().WithKey(b.key).FieldPath(),
	})
}

func (b *mapFilterCndBuilderLabels) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:           op,
		Log_FieldPathValue: NewLogFieldPathBuilder().Labels().WithKey(b.key).WithValue(value),
	})
}

type filterCndBuilderTime struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderTime) Eq(value *timestamppb.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderTime) Neq(value *timestamppb.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderTime) Gt(value *timestamppb.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderTime) Gte(value *timestamppb.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderTime) Lt(value *timestamppb.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderTime) Lte(value *timestamppb.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderTime) In(values []*timestamppb.Timestamp) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().Time().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderTime) NotIn(values []*timestamppb.Timestamp) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().Time().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderTime) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewLogFieldPathBuilder().Time().FieldPath(),
	})
}

func (b *filterCndBuilderTime) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewLogFieldPathBuilder().Time().FieldPath(),
	})
}

func (b *filterCndBuilderTime) compare(op gotenfilter.CompareOperator, value *timestamppb.Timestamp) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:           op,
		Log_FieldPathValue: NewLogFieldPathBuilder().Time().WithValue(value),
	})
}

type filterCndBuilderJsonPayload struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderJsonPayload) Eq(value *structpb.Struct) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderJsonPayload) Neq(value *structpb.Struct) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderJsonPayload) Gt(value *structpb.Struct) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderJsonPayload) Gte(value *structpb.Struct) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderJsonPayload) Lt(value *structpb.Struct) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderJsonPayload) Lte(value *structpb.Struct) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderJsonPayload) In(values []*structpb.Struct) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().JsonPayload().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderJsonPayload) NotIn(values []*structpb.Struct) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().JsonPayload().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderJsonPayload) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewLogFieldPathBuilder().JsonPayload().FieldPath(),
	})
}

func (b *filterCndBuilderJsonPayload) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewLogFieldPathBuilder().JsonPayload().FieldPath(),
	})
}

func (b *filterCndBuilderJsonPayload) compare(op gotenfilter.CompareOperator, value *structpb.Struct) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:           op,
		Log_FieldPathValue: NewLogFieldPathBuilder().JsonPayload().WithValue(value),
	})
}

type filterCndBuilderPbPayload struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderPbPayload) Eq(value *anypb.Any) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderPbPayload) Neq(value *anypb.Any) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderPbPayload) Gt(value *anypb.Any) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderPbPayload) Gte(value *anypb.Any) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderPbPayload) Lt(value *anypb.Any) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderPbPayload) Lte(value *anypb.Any) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderPbPayload) In(values []*anypb.Any) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().PbPayload().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderPbPayload) NotIn(values []*anypb.Any) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().PbPayload().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderPbPayload) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewLogFieldPathBuilder().PbPayload().FieldPath(),
	})
}

func (b *filterCndBuilderPbPayload) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewLogFieldPathBuilder().PbPayload().FieldPath(),
	})
}

func (b *filterCndBuilderPbPayload) compare(op gotenfilter.CompareOperator, value *anypb.Any) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:           op,
		Log_FieldPathValue: NewLogFieldPathBuilder().PbPayload().WithValue(value),
	})
}

type filterCndBuilderStringPayload struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderStringPayload) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderStringPayload) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderStringPayload) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderStringPayload) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderStringPayload) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderStringPayload) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderStringPayload) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().StringPayload().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderStringPayload) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().StringPayload().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderStringPayload) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewLogFieldPathBuilder().StringPayload().FieldPath(),
	})
}

func (b *filterCndBuilderStringPayload) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewLogFieldPathBuilder().StringPayload().FieldPath(),
	})
}

func (b *filterCndBuilderStringPayload) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:           op,
		Log_FieldPathValue: NewLogFieldPathBuilder().StringPayload().WithValue(value),
	})
}

type filterCndBuilderBytesPayload struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderBytesPayload) Eq(value []byte) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderBytesPayload) Neq(value []byte) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderBytesPayload) Gt(value []byte) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderBytesPayload) Gte(value []byte) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderBytesPayload) Lt(value []byte) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderBytesPayload) Lte(value []byte) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderBytesPayload) In(values [][]byte) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().BytesPayload().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderBytesPayload) NotIn(values [][]byte) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().BytesPayload().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderBytesPayload) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewLogFieldPathBuilder().BytesPayload().FieldPath(),
	})
}

func (b *filterCndBuilderBytesPayload) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewLogFieldPathBuilder().BytesPayload().FieldPath(),
	})
}

func (b *filterCndBuilderBytesPayload) compare(op gotenfilter.CompareOperator, value []byte) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:           op,
		Log_FieldPathValue: NewLogFieldPathBuilder().BytesPayload().WithValue(value),
	})
}

type filterCndBuilderBinKey struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderBinKey) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderBinKey) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderBinKey) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderBinKey) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderBinKey) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderBinKey) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderBinKey) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().BinKey().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderBinKey) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Log_FieldPathArrayOfValues: NewLogFieldPathBuilder().BinKey().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderBinKey) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewLogFieldPathBuilder().BinKey().FieldPath(),
	})
}

func (b *filterCndBuilderBinKey) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewLogFieldPathBuilder().BinKey().FieldPath(),
	})
}

func (b *filterCndBuilderBinKey) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:           op,
		Log_FieldPathValue: NewLogFieldPathBuilder().BinKey().WithValue(value),
	})
}
