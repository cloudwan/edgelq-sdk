// Code generated by protoc-gen-goten-resource
// Resource: Secret
// DO NOT EDIT!!!

package secret

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	gotenfilter "github.com/cloudwan/goten-sdk/runtime/resource/filter"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/secrets/resources/v1alpha/project"
)

// ensure the imports are used
var (
	_ = gotenresource.ConditionContainsTypeAll
	_ = gotenfilter.AND
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
)

type FilterBuilderOrCondition interface {
	_IsSecretFilterBuilderOrCondition()
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

func (b *FilterBuilder) _IsSecretFilterBuilderOrCondition() {}

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

func (b *FilterBuilder) WherePath(fp Secret_FieldPath, opts ...gotenfilter.FilterConditionOption) *filterCndBuilderAnyPath {
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
	fp      Secret_FieldPath
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
		Secret_FieldPathArrayOfValues: b.fp.WithIArrayOfValues(values),
	})
}

func (b *filterCndBuilderAnyPath) NotIn(values interface{}) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Secret_FieldPathArrayOfValues: b.fp.WithIArrayOfValues(values),
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
		Operator:              op,
		Secret_FieldPathValue: b.fp.WithIValue(value),
	})
}

type filterCndBuilder struct {
	builder *FilterBuilder
}

func (b *filterCndBuilder) Name() *filterCndBuilderName {
	return &filterCndBuilderName{builder: b.builder}
}

func (b *filterCndBuilder) EncData() *filterCndBuilderEncData {
	return &filterCndBuilderEncData{builder: b.builder}
}

func (b *filterCndBuilder) Data() *filterCndBuilderData {
	return &filterCndBuilderData{builder: b.builder}
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
		Secret_FieldPathArrayOfValues: NewSecretFieldPathBuilder().Name().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderName) NotIn(values []*Name) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Secret_FieldPathArrayOfValues: NewSecretFieldPathBuilder().Name().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderName) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewSecretFieldPathBuilder().Name().FieldPath(),
	})
}

func (b *filterCndBuilderName) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewSecretFieldPathBuilder().Name().FieldPath(),
	})
}

func (b *filterCndBuilderName) compare(op gotenfilter.CompareOperator, value *Name) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:              op,
		Secret_FieldPathValue: NewSecretFieldPathBuilder().Name().WithValue(value),
	})
}

type filterCndBuilderEncData struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderEncData) Eq(value []byte) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderEncData) Neq(value []byte) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderEncData) Gt(value []byte) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderEncData) Gte(value []byte) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderEncData) Lt(value []byte) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderEncData) Lte(value []byte) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderEncData) In(values [][]byte) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Secret_FieldPathArrayOfValues: NewSecretFieldPathBuilder().EncData().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderEncData) NotIn(values [][]byte) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Secret_FieldPathArrayOfValues: NewSecretFieldPathBuilder().EncData().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderEncData) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewSecretFieldPathBuilder().EncData().FieldPath(),
	})
}

func (b *filterCndBuilderEncData) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewSecretFieldPathBuilder().EncData().FieldPath(),
	})
}

func (b *filterCndBuilderEncData) compare(op gotenfilter.CompareOperator, value []byte) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:              op,
		Secret_FieldPathValue: NewSecretFieldPathBuilder().EncData().WithValue(value),
	})
}

type filterCndBuilderData struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderData) Eq(value map[string]string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderData) Neq(value map[string]string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderData) Gt(value map[string]string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderData) Gte(value map[string]string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderData) Lt(value map[string]string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderData) Lte(value map[string]string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderData) In(values []map[string]string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Secret_FieldPathArrayOfValues: NewSecretFieldPathBuilder().Data().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderData) NotIn(values []map[string]string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Secret_FieldPathArrayOfValues: NewSecretFieldPathBuilder().Data().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderData) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewSecretFieldPathBuilder().Data().FieldPath(),
	})
}

func (b *filterCndBuilderData) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewSecretFieldPathBuilder().Data().FieldPath(),
	})
}

func (b *filterCndBuilderData) compare(op gotenfilter.CompareOperator, value map[string]string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:              op,
		Secret_FieldPathValue: NewSecretFieldPathBuilder().Data().WithValue(value),
	})
}

func (b *filterCndBuilderData) WithKey(key string) *mapFilterCndBuilderData {
	return &mapFilterCndBuilderData{builder: b.builder, key: key}
}

type mapFilterCndBuilderData struct {
	builder *FilterBuilder
	key     string
}

func (b *mapFilterCndBuilderData) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *mapFilterCndBuilderData) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *mapFilterCndBuilderData) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *mapFilterCndBuilderData) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *mapFilterCndBuilderData) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *mapFilterCndBuilderData) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *mapFilterCndBuilderData) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		Secret_FieldPathArrayOfValues: NewSecretFieldPathBuilder().Data().WithKey(b.key).WithArrayOfValues(values),
	})
}

func (b *mapFilterCndBuilderData) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		Secret_FieldPathArrayOfValues: NewSecretFieldPathBuilder().Data().WithKey(b.key).WithArrayOfValues(values),
	})
}

func (b *mapFilterCndBuilderData) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewSecretFieldPathBuilder().Data().WithKey(b.key).FieldPath(),
	})
}

func (b *mapFilterCndBuilderData) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewSecretFieldPathBuilder().Data().WithKey(b.key).FieldPath(),
	})
}

func (b *mapFilterCndBuilderData) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:              op,
		Secret_FieldPathValue: NewSecretFieldPathBuilder().Data().WithKey(b.key).WithValue(value),
	})
}
