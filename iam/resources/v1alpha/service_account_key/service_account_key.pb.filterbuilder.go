// Code generated by protoc-gen-goten-resource
// Resource: ServiceAccountKey
// DO NOT EDIT!!!

package service_account_key

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	gotenfilter "github.com/cloudwan/goten-sdk/runtime/resource/filter"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/service_account"
	syncing_meta "github.com/cloudwan/edgelq-sdk/meta/multi_region/proto/syncing_meta"
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
	_ = &service_account.ServiceAccount{}
	_ = &syncing_meta.SyncingMeta{}
	_ = &timestamp.Timestamp{}
)

type FilterBuilderOrCondition interface {
	_IsServiceAccountKeyFilterBuilderOrCondition()
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

func (b *FilterBuilder) _IsServiceAccountKeyFilterBuilderOrCondition() {}

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

func (b *FilterBuilder) WherePath(fp ServiceAccountKey_FieldPath, opts ...gotenfilter.FilterConditionOption) *filterCndBuilderAnyPath {
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
	fp      ServiceAccountKey_FieldPath
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
		ServiceAccountKey_FieldPathArrayOfValues: b.fp.WithIArrayOfValues(values),
	})
}

func (b *filterCndBuilderAnyPath) NotIn(values interface{}) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		ServiceAccountKey_FieldPathArrayOfValues: b.fp.WithIArrayOfValues(values),
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
		Operator:                         op,
		ServiceAccountKey_FieldPathValue: b.fp.WithIValue(value),
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

func (b *filterCndBuilder) PublicKeyData() *filterCndBuilderPublicKeyData {
	return &filterCndBuilderPublicKeyData{builder: b.builder}
}

func (b *filterCndBuilder) PrivateKeyData() *filterCndBuilderPrivateKeyData {
	return &filterCndBuilderPrivateKeyData{builder: b.builder}
}

func (b *filterCndBuilder) Algorithm() *filterCndBuilderAlgorithm {
	return &filterCndBuilderAlgorithm{builder: b.builder}
}

func (b *filterCndBuilder) ValidNotBefore() *filterCndBuilderValidNotBefore {
	return &filterCndBuilderValidNotBefore{builder: b.builder}
}

func (b *filterCndBuilder) ValidNotAfter() *filterCndBuilderValidNotAfter {
	return &filterCndBuilderValidNotAfter{builder: b.builder}
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
		ServiceAccountKey_FieldPathArrayOfValues: NewServiceAccountKeyFieldPathBuilder().Name().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderName) NotIn(values []*Name) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		ServiceAccountKey_FieldPathArrayOfValues: NewServiceAccountKeyFieldPathBuilder().Name().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderName) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewServiceAccountKeyFieldPathBuilder().Name().FieldPath(),
	})
}

func (b *filterCndBuilderName) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewServiceAccountKeyFieldPathBuilder().Name().FieldPath(),
	})
}

func (b *filterCndBuilderName) compare(op gotenfilter.CompareOperator, value *Name) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                         op,
		ServiceAccountKey_FieldPathValue: NewServiceAccountKeyFieldPathBuilder().Name().WithValue(value),
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
		ServiceAccountKey_FieldPathArrayOfValues: NewServiceAccountKeyFieldPathBuilder().DisplayName().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderDisplayName) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		ServiceAccountKey_FieldPathArrayOfValues: NewServiceAccountKeyFieldPathBuilder().DisplayName().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderDisplayName) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewServiceAccountKeyFieldPathBuilder().DisplayName().FieldPath(),
	})
}

func (b *filterCndBuilderDisplayName) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewServiceAccountKeyFieldPathBuilder().DisplayName().FieldPath(),
	})
}

func (b *filterCndBuilderDisplayName) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                         op,
		ServiceAccountKey_FieldPathValue: NewServiceAccountKeyFieldPathBuilder().DisplayName().WithValue(value),
	})
}

type filterCndBuilderPublicKeyData struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderPublicKeyData) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderPublicKeyData) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderPublicKeyData) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderPublicKeyData) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderPublicKeyData) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderPublicKeyData) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderPublicKeyData) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		ServiceAccountKey_FieldPathArrayOfValues: NewServiceAccountKeyFieldPathBuilder().PublicKeyData().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderPublicKeyData) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		ServiceAccountKey_FieldPathArrayOfValues: NewServiceAccountKeyFieldPathBuilder().PublicKeyData().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderPublicKeyData) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewServiceAccountKeyFieldPathBuilder().PublicKeyData().FieldPath(),
	})
}

func (b *filterCndBuilderPublicKeyData) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewServiceAccountKeyFieldPathBuilder().PublicKeyData().FieldPath(),
	})
}

func (b *filterCndBuilderPublicKeyData) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                         op,
		ServiceAccountKey_FieldPathValue: NewServiceAccountKeyFieldPathBuilder().PublicKeyData().WithValue(value),
	})
}

type filterCndBuilderPrivateKeyData struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderPrivateKeyData) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderPrivateKeyData) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderPrivateKeyData) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderPrivateKeyData) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderPrivateKeyData) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderPrivateKeyData) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderPrivateKeyData) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		ServiceAccountKey_FieldPathArrayOfValues: NewServiceAccountKeyFieldPathBuilder().PrivateKeyData().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderPrivateKeyData) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		ServiceAccountKey_FieldPathArrayOfValues: NewServiceAccountKeyFieldPathBuilder().PrivateKeyData().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderPrivateKeyData) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewServiceAccountKeyFieldPathBuilder().PrivateKeyData().FieldPath(),
	})
}

func (b *filterCndBuilderPrivateKeyData) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewServiceAccountKeyFieldPathBuilder().PrivateKeyData().FieldPath(),
	})
}

func (b *filterCndBuilderPrivateKeyData) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                         op,
		ServiceAccountKey_FieldPathValue: NewServiceAccountKeyFieldPathBuilder().PrivateKeyData().WithValue(value),
	})
}

type filterCndBuilderAlgorithm struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderAlgorithm) Eq(value ServiceAccountKey_Algorithm) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderAlgorithm) Neq(value ServiceAccountKey_Algorithm) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderAlgorithm) Gt(value ServiceAccountKey_Algorithm) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderAlgorithm) Gte(value ServiceAccountKey_Algorithm) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderAlgorithm) Lt(value ServiceAccountKey_Algorithm) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderAlgorithm) Lte(value ServiceAccountKey_Algorithm) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderAlgorithm) In(values []ServiceAccountKey_Algorithm) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		ServiceAccountKey_FieldPathArrayOfValues: NewServiceAccountKeyFieldPathBuilder().Algorithm().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderAlgorithm) NotIn(values []ServiceAccountKey_Algorithm) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		ServiceAccountKey_FieldPathArrayOfValues: NewServiceAccountKeyFieldPathBuilder().Algorithm().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderAlgorithm) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewServiceAccountKeyFieldPathBuilder().Algorithm().FieldPath(),
	})
}

func (b *filterCndBuilderAlgorithm) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewServiceAccountKeyFieldPathBuilder().Algorithm().FieldPath(),
	})
}

func (b *filterCndBuilderAlgorithm) compare(op gotenfilter.CompareOperator, value ServiceAccountKey_Algorithm) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                         op,
		ServiceAccountKey_FieldPathValue: NewServiceAccountKeyFieldPathBuilder().Algorithm().WithValue(value),
	})
}

type filterCndBuilderValidNotBefore struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderValidNotBefore) Eq(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderValidNotBefore) Neq(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderValidNotBefore) Gt(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderValidNotBefore) Gte(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderValidNotBefore) Lt(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderValidNotBefore) Lte(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderValidNotBefore) In(values []*timestamp.Timestamp) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		ServiceAccountKey_FieldPathArrayOfValues: NewServiceAccountKeyFieldPathBuilder().ValidNotBefore().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderValidNotBefore) NotIn(values []*timestamp.Timestamp) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		ServiceAccountKey_FieldPathArrayOfValues: NewServiceAccountKeyFieldPathBuilder().ValidNotBefore().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderValidNotBefore) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewServiceAccountKeyFieldPathBuilder().ValidNotBefore().FieldPath(),
	})
}

func (b *filterCndBuilderValidNotBefore) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewServiceAccountKeyFieldPathBuilder().ValidNotBefore().FieldPath(),
	})
}

func (b *filterCndBuilderValidNotBefore) compare(op gotenfilter.CompareOperator, value *timestamp.Timestamp) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                         op,
		ServiceAccountKey_FieldPathValue: NewServiceAccountKeyFieldPathBuilder().ValidNotBefore().WithValue(value),
	})
}

type filterCndBuilderValidNotAfter struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderValidNotAfter) Eq(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderValidNotAfter) Neq(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderValidNotAfter) Gt(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderValidNotAfter) Gte(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderValidNotAfter) Lt(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderValidNotAfter) Lte(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderValidNotAfter) In(values []*timestamp.Timestamp) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		ServiceAccountKey_FieldPathArrayOfValues: NewServiceAccountKeyFieldPathBuilder().ValidNotAfter().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderValidNotAfter) NotIn(values []*timestamp.Timestamp) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		ServiceAccountKey_FieldPathArrayOfValues: NewServiceAccountKeyFieldPathBuilder().ValidNotAfter().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderValidNotAfter) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewServiceAccountKeyFieldPathBuilder().ValidNotAfter().FieldPath(),
	})
}

func (b *filterCndBuilderValidNotAfter) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewServiceAccountKeyFieldPathBuilder().ValidNotAfter().FieldPath(),
	})
}

func (b *filterCndBuilderValidNotAfter) compare(op gotenfilter.CompareOperator, value *timestamp.Timestamp) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                         op,
		ServiceAccountKey_FieldPathValue: NewServiceAccountKeyFieldPathBuilder().ValidNotAfter().WithValue(value),
	})
}