// Code generated by protoc-gen-goten-resource
// Resource: Condition
// DO NOT EDIT!!!

package condition

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	gotenfilter "github.com/cloudwan/goten-sdk/runtime/resource/filter"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// ensure the imports are used
var (
	_ = gotenresource.ConditionContainsTypeAll
	_ = gotenfilter.AND
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &timestamp.Timestamp{}
)

type FilterBuilderOrCondition interface {
	_IsConditionFilterBuilderOrCondition()
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

func (b *FilterBuilder) _IsConditionFilterBuilderOrCondition() {}

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

func (b *FilterBuilder) WherePath(fp Condition_FieldPath, opts ...gotenfilter.FilterConditionOption) *filterCndBuilderAnyPath {
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
	fp      Condition_FieldPath
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
		Condition_FieldPathArrayOfValues: b.fp.WithIArrayOfValues(values),
	})
}

func (b *filterCndBuilderAnyPath) NotIn(values interface{}) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Condition_FieldPathArrayOfValues: b.fp.WithIArrayOfValues(values),
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

func (b *filterCndBuilderAnyPath) compare(op gotenfilter.CompareOperator, value interface{}) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                 op,
		Condition_FieldPathValue: b.fp.WithIValue(value),
	})
}

type filterCndBuilder struct {
	builder *FilterBuilder
}

func (b *filterCndBuilder) Name() *filterCndBuilderName {
	return &filterCndBuilderName{builder: b.builder}
}

func (b *filterCndBuilder) DisplayName() *filterCndBuilderDisplayName {
	return &filterCndBuilderDisplayName{builder: b.builder}
}

func (b *filterCndBuilder) Description() *filterCndBuilderDescription {
	return &filterCndBuilderDescription{builder: b.builder}
}

func (b *filterCndBuilder) Expression() *filterCndBuilderExpression {
	return &filterCndBuilderExpression{builder: b.builder}
}

func (b *filterCndBuilder) ParameterDeclarations() *filterCndBuilderParameterDeclarations {
	return &filterCndBuilderParameterDeclarations{builder: b.builder}
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
		Condition_FieldPathArrayOfValues: NewConditionFieldPathBuilder().Name().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderName) NotIn(values []*Name) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Condition_FieldPathArrayOfValues: NewConditionFieldPathBuilder().Name().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderName) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewConditionFieldPathBuilder().Name().FieldPath(),
	})
}

func (b *filterCndBuilderName) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewConditionFieldPathBuilder().Name().FieldPath(),
	})
}

func (b *filterCndBuilderName) compare(op gotenfilter.CompareOperator, value *Name) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                 op,
		Condition_FieldPathValue: NewConditionFieldPathBuilder().Name().WithValue(value),
	})
}

type filterCndBuilderDisplayName struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderDisplayName) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderDisplayName) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderDisplayName) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderDisplayName) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderDisplayName) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderDisplayName) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderDisplayName) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Condition_FieldPathArrayOfValues: NewConditionFieldPathBuilder().DisplayName().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderDisplayName) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Condition_FieldPathArrayOfValues: NewConditionFieldPathBuilder().DisplayName().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderDisplayName) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewConditionFieldPathBuilder().DisplayName().FieldPath(),
	})
}

func (b *filterCndBuilderDisplayName) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewConditionFieldPathBuilder().DisplayName().FieldPath(),
	})
}

func (b *filterCndBuilderDisplayName) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                 op,
		Condition_FieldPathValue: NewConditionFieldPathBuilder().DisplayName().WithValue(value),
	})
}

type filterCndBuilderDescription struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderDescription) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderDescription) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderDescription) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderDescription) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderDescription) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderDescription) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderDescription) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Condition_FieldPathArrayOfValues: NewConditionFieldPathBuilder().Description().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderDescription) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Condition_FieldPathArrayOfValues: NewConditionFieldPathBuilder().Description().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderDescription) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewConditionFieldPathBuilder().Description().FieldPath(),
	})
}

func (b *filterCndBuilderDescription) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewConditionFieldPathBuilder().Description().FieldPath(),
	})
}

func (b *filterCndBuilderDescription) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                 op,
		Condition_FieldPathValue: NewConditionFieldPathBuilder().Description().WithValue(value),
	})
}

type filterCndBuilderExpression struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderExpression) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderExpression) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderExpression) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderExpression) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderExpression) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderExpression) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderExpression) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Condition_FieldPathArrayOfValues: NewConditionFieldPathBuilder().Expression().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderExpression) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Condition_FieldPathArrayOfValues: NewConditionFieldPathBuilder().Expression().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderExpression) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewConditionFieldPathBuilder().Expression().FieldPath(),
	})
}

func (b *filterCndBuilderExpression) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewConditionFieldPathBuilder().Expression().FieldPath(),
	})
}

func (b *filterCndBuilderExpression) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                 op,
		Condition_FieldPathValue: NewConditionFieldPathBuilder().Expression().WithValue(value),
	})
}

type filterCndBuilderParameterDeclarations struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderParameterDeclarations) Eq(value []*Condition_ParameterDeclaration) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderParameterDeclarations) Neq(value []*Condition_ParameterDeclaration) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderParameterDeclarations) Gt(value []*Condition_ParameterDeclaration) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderParameterDeclarations) Gte(value []*Condition_ParameterDeclaration) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderParameterDeclarations) Lt(value []*Condition_ParameterDeclaration) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderParameterDeclarations) Lte(value []*Condition_ParameterDeclaration) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderParameterDeclarations) In(values [][]*Condition_ParameterDeclaration) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Condition_FieldPathArrayOfValues: NewConditionFieldPathBuilder().ParameterDeclarations().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderParameterDeclarations) NotIn(values [][]*Condition_ParameterDeclaration) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Condition_FieldPathArrayOfValues: NewConditionFieldPathBuilder().ParameterDeclarations().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderParameterDeclarations) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewConditionFieldPathBuilder().ParameterDeclarations().FieldPath(),
	})
}

func (b *filterCndBuilderParameterDeclarations) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewConditionFieldPathBuilder().ParameterDeclarations().FieldPath(),
	})
}

func (b *filterCndBuilderParameterDeclarations) Contains(value *Condition_ParameterDeclaration) *FilterBuilder {
	return b.builder.addCond(&FilterConditionContains{
		Type:      gotenresource.ConditionContainsTypeValue,
		FieldPath: NewConditionFieldPathBuilder().ParameterDeclarations().FieldPath(),
		Value:     NewConditionFieldPathBuilder().ParameterDeclarations().WithItemValue(value),
	})
}

func (b *filterCndBuilderParameterDeclarations) ContainsAnyOf(values []*Condition_ParameterDeclaration) *FilterBuilder {
	pathSelector := NewConditionFieldPathBuilder().ParameterDeclarations()
	itemValues := make([]Condition_FieldPathArrayItemValue, 0, len(values))
	for _, value := range values {
		itemValues = append(itemValues, pathSelector.WithItemValue(value))
	}
	return b.builder.addCond(&FilterConditionContains{
		Type:      gotenresource.ConditionContainsTypeAny,
		FieldPath: NewConditionFieldPathBuilder().ParameterDeclarations().FieldPath(),
		Values:    itemValues,
	})
}

func (b *filterCndBuilderParameterDeclarations) ContainsAll(values []*Condition_ParameterDeclaration) *FilterBuilder {
	pathSelector := NewConditionFieldPathBuilder().ParameterDeclarations()
	itemValues := make([]Condition_FieldPathArrayItemValue, 0, len(values))
	for _, value := range values {
		itemValues = append(itemValues, pathSelector.WithItemValue(value))
	}
	return b.builder.addCond(&FilterConditionContains{
		Type:      gotenresource.ConditionContainsTypeAll,
		FieldPath: NewConditionFieldPathBuilder().ParameterDeclarations().FieldPath(),
		Values:    itemValues,
	})
}

func (b *filterCndBuilderParameterDeclarations) compare(op gotenfilter.CompareOperator, value []*Condition_ParameterDeclaration) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                 op,
		Condition_FieldPathValue: NewConditionFieldPathBuilder().ParameterDeclarations().WithValue(value),
	})
}

func (b *filterCndBuilderParameterDeclarations) Key() *filterCndBuilderParameterDeclarationsKey {
	return &filterCndBuilderParameterDeclarationsKey{builder: b.builder}
}

func (b *filterCndBuilderParameterDeclarations) Type() *filterCndBuilderParameterDeclarationsType {
	return &filterCndBuilderParameterDeclarationsType{builder: b.builder}
}

type filterCndBuilderParameterDeclarationsKey struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderParameterDeclarationsKey) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderParameterDeclarationsKey) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderParameterDeclarationsKey) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderParameterDeclarationsKey) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderParameterDeclarationsKey) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderParameterDeclarationsKey) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderParameterDeclarationsKey) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Condition_FieldPathArrayOfValues: NewConditionFieldPathBuilder().ParameterDeclarations().Key().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderParameterDeclarationsKey) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Condition_FieldPathArrayOfValues: NewConditionFieldPathBuilder().ParameterDeclarations().Key().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderParameterDeclarationsKey) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewConditionFieldPathBuilder().ParameterDeclarations().Key().FieldPath(),
	})
}

func (b *filterCndBuilderParameterDeclarationsKey) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewConditionFieldPathBuilder().ParameterDeclarations().Key().FieldPath(),
	})
}

func (b *filterCndBuilderParameterDeclarationsKey) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                 op,
		Condition_FieldPathValue: NewConditionFieldPathBuilder().ParameterDeclarations().Key().WithValue(value),
	})
}

type filterCndBuilderParameterDeclarationsType struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderParameterDeclarationsType) Eq(value Condition_ParameterType) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderParameterDeclarationsType) Neq(value Condition_ParameterType) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderParameterDeclarationsType) Gt(value Condition_ParameterType) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderParameterDeclarationsType) Gte(value Condition_ParameterType) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderParameterDeclarationsType) Lt(value Condition_ParameterType) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderParameterDeclarationsType) Lte(value Condition_ParameterType) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderParameterDeclarationsType) In(values []Condition_ParameterType) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Condition_FieldPathArrayOfValues: NewConditionFieldPathBuilder().ParameterDeclarations().Type().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderParameterDeclarationsType) NotIn(values []Condition_ParameterType) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Condition_FieldPathArrayOfValues: NewConditionFieldPathBuilder().ParameterDeclarations().Type().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderParameterDeclarationsType) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewConditionFieldPathBuilder().ParameterDeclarations().Type().FieldPath(),
	})
}

func (b *filterCndBuilderParameterDeclarationsType) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewConditionFieldPathBuilder().ParameterDeclarations().Type().FieldPath(),
	})
}

func (b *filterCndBuilderParameterDeclarationsType) compare(op gotenfilter.CompareOperator, value Condition_ParameterType) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                 op,
		Condition_FieldPathValue: NewConditionFieldPathBuilder().ParameterDeclarations().Type().WithValue(value),
	})
}
