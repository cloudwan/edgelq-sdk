// Code generated by protoc-gen-goten-resource
// Resource: OrganizationInvitation
// DO NOT EDIT!!!

package organization_invitation

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	gotenfilter "github.com/cloudwan/goten-sdk/runtime/resource/filter"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/common"
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/condition"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/permission"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/role"
	service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/service_account"
	user "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/user"
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
	_ = &iam_common.Actor{}
	_ = &condition.Condition{}
	_ = &organization.Organization{}
	_ = &permission.Permission{}
	_ = &project.Project{}
	_ = &role.Role{}
	_ = &service_account.ServiceAccount{}
	_ = &user.User{}
	_ = &timestamp.Timestamp{}
)

type FilterBuilderOrCondition interface {
	_IsOrganizationInvitationFilterBuilderOrCondition()
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

func (b *FilterBuilder) _IsOrganizationInvitationFilterBuilderOrCondition() {}

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

func (b *FilterBuilder) WherePath(fp OrganizationInvitation_FieldPath, opts ...gotenfilter.FilterConditionOption) *filterCndBuilderAnyPath {
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
	fp      OrganizationInvitation_FieldPath
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
		OrganizationInvitation_FieldPathArrayOfValues: b.fp.WithIArrayOfValues(values),
	})
}

func (b *filterCndBuilderAnyPath) NotIn(values interface{}) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		OrganizationInvitation_FieldPathArrayOfValues: b.fp.WithIArrayOfValues(values),
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
		Operator:                              op,
		OrganizationInvitation_FieldPathValue: b.fp.WithIValue(value),
	})
}

type filterCndBuilder struct {
	builder *FilterBuilder
}

func (b *filterCndBuilder) Name() *filterCndBuilderName {
	return &filterCndBuilderName{builder: b.builder}
}

func (b *filterCndBuilder) Invitation() *filterCndBuilderInvitation {
	return &filterCndBuilderInvitation{builder: b.builder}
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
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Name().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderName) NotIn(values []*Name) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Name().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderName) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Name().FieldPath(),
	})
}

func (b *filterCndBuilderName) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Name().FieldPath(),
	})
}

func (b *filterCndBuilderName) compare(op gotenfilter.CompareOperator, value *Name) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                              op,
		OrganizationInvitation_FieldPathValue: NewOrganizationInvitationFieldPathBuilder().Name().WithValue(value),
	})
}

type filterCndBuilderInvitation struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderInvitation) Eq(value *iam_common.Invitation) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderInvitation) Neq(value *iam_common.Invitation) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderInvitation) Gt(value *iam_common.Invitation) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderInvitation) Gte(value *iam_common.Invitation) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderInvitation) Lt(value *iam_common.Invitation) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderInvitation) Lte(value *iam_common.Invitation) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderInvitation) In(values []*iam_common.Invitation) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitation) NotIn(values []*iam_common.Invitation) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitation) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().FieldPath(),
	})
}

func (b *filterCndBuilderInvitation) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().FieldPath(),
	})
}

func (b *filterCndBuilderInvitation) compare(op gotenfilter.CompareOperator, value *iam_common.Invitation) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                              op,
		OrganizationInvitation_FieldPathValue: NewOrganizationInvitationFieldPathBuilder().Invitation().WithValue(value),
	})
}

func (b *filterCndBuilderInvitation) InviteeEmail() *filterCndBuilderInvitationInviteeEmail {
	return &filterCndBuilderInvitationInviteeEmail{builder: b.builder}
}

func (b *filterCndBuilderInvitation) InviterActor() *filterCndBuilderInvitationInviterActor {
	return &filterCndBuilderInvitationInviterActor{builder: b.builder}
}

func (b *filterCndBuilderInvitation) InviterFullName() *filterCndBuilderInvitationInviterFullName {
	return &filterCndBuilderInvitationInviterFullName{builder: b.builder}
}

func (b *filterCndBuilderInvitation) InviterEmail() *filterCndBuilderInvitationInviterEmail {
	return &filterCndBuilderInvitationInviterEmail{builder: b.builder}
}

func (b *filterCndBuilderInvitation) Roles() *filterCndBuilderInvitationRoles {
	return &filterCndBuilderInvitationRoles{builder: b.builder}
}

func (b *filterCndBuilderInvitation) ExpirationDate() *filterCndBuilderInvitationExpirationDate {
	return &filterCndBuilderInvitationExpirationDate{builder: b.builder}
}

func (b *filterCndBuilderInvitation) State() *filterCndBuilderInvitationState {
	return &filterCndBuilderInvitationState{builder: b.builder}
}

type filterCndBuilderInvitationInviteeEmail struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderInvitationInviteeEmail) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderInvitationInviteeEmail) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderInvitationInviteeEmail) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderInvitationInviteeEmail) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderInvitationInviteeEmail) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderInvitationInviteeEmail) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderInvitationInviteeEmail) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().InviteeEmail().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationInviteeEmail) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().InviteeEmail().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationInviteeEmail) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().InviteeEmail().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationInviteeEmail) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().InviteeEmail().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationInviteeEmail) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                              op,
		OrganizationInvitation_FieldPathValue: NewOrganizationInvitationFieldPathBuilder().Invitation().InviteeEmail().WithValue(value),
	})
}

type filterCndBuilderInvitationInviterActor struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderInvitationInviterActor) Eq(value *iam_common.Actor) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderInvitationInviterActor) Neq(value *iam_common.Actor) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderInvitationInviterActor) Gt(value *iam_common.Actor) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderInvitationInviterActor) Gte(value *iam_common.Actor) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderInvitationInviterActor) Lt(value *iam_common.Actor) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderInvitationInviterActor) Lte(value *iam_common.Actor) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderInvitationInviterActor) In(values []*iam_common.Actor) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationInviterActor) NotIn(values []*iam_common.Actor) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationInviterActor) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationInviterActor) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationInviterActor) compare(op gotenfilter.CompareOperator, value *iam_common.Actor) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                              op,
		OrganizationInvitation_FieldPathValue: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().WithValue(value),
	})
}

func (b *filterCndBuilderInvitationInviterActor) User() *filterCndBuilderInvitationInviterActorUser {
	return &filterCndBuilderInvitationInviterActorUser{builder: b.builder}
}

func (b *filterCndBuilderInvitationInviterActor) ServiceAccount() *filterCndBuilderInvitationInviterActorServiceAccount {
	return &filterCndBuilderInvitationInviterActorServiceAccount{builder: b.builder}
}

type filterCndBuilderInvitationInviterActorUser struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderInvitationInviterActorUser) Eq(value *user.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderInvitationInviterActorUser) Neq(value *user.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderInvitationInviterActorUser) Gt(value *user.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderInvitationInviterActorUser) Gte(value *user.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderInvitationInviterActorUser) Lt(value *user.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderInvitationInviterActorUser) Lte(value *user.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderInvitationInviterActorUser) In(values []*user.Reference) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().User().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationInviterActorUser) NotIn(values []*user.Reference) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().User().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationInviterActorUser) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().User().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationInviterActorUser) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().User().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationInviterActorUser) compare(op gotenfilter.CompareOperator, value *user.Reference) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                              op,
		OrganizationInvitation_FieldPathValue: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().User().WithValue(value),
	})
}

type filterCndBuilderInvitationInviterActorServiceAccount struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderInvitationInviterActorServiceAccount) Eq(value *service_account.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderInvitationInviterActorServiceAccount) Neq(value *service_account.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderInvitationInviterActorServiceAccount) Gt(value *service_account.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderInvitationInviterActorServiceAccount) Gte(value *service_account.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderInvitationInviterActorServiceAccount) Lt(value *service_account.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderInvitationInviterActorServiceAccount) Lte(value *service_account.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderInvitationInviterActorServiceAccount) In(values []*service_account.Reference) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().ServiceAccount().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationInviterActorServiceAccount) NotIn(values []*service_account.Reference) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().ServiceAccount().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationInviterActorServiceAccount) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().ServiceAccount().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationInviterActorServiceAccount) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().ServiceAccount().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationInviterActorServiceAccount) compare(op gotenfilter.CompareOperator, value *service_account.Reference) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                              op,
		OrganizationInvitation_FieldPathValue: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().ServiceAccount().WithValue(value),
	})
}

type filterCndBuilderInvitationInviterFullName struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderInvitationInviterFullName) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderInvitationInviterFullName) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderInvitationInviterFullName) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderInvitationInviterFullName) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderInvitationInviterFullName) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderInvitationInviterFullName) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderInvitationInviterFullName) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterFullName().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationInviterFullName) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterFullName().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationInviterFullName) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterFullName().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationInviterFullName) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterFullName().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationInviterFullName) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                              op,
		OrganizationInvitation_FieldPathValue: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterFullName().WithValue(value),
	})
}

type filterCndBuilderInvitationInviterEmail struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderInvitationInviterEmail) Eq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderInvitationInviterEmail) Neq(value string) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderInvitationInviterEmail) Gt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderInvitationInviterEmail) Gte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderInvitationInviterEmail) Lt(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderInvitationInviterEmail) Lte(value string) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderInvitationInviterEmail) In(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterEmail().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationInviterEmail) NotIn(values []string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterEmail().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationInviterEmail) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterEmail().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationInviterEmail) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterEmail().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationInviterEmail) compare(op gotenfilter.CompareOperator, value string) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                              op,
		OrganizationInvitation_FieldPathValue: NewOrganizationInvitationFieldPathBuilder().Invitation().InviterEmail().WithValue(value),
	})
}

type filterCndBuilderInvitationRoles struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderInvitationRoles) Eq(value []*role.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderInvitationRoles) Neq(value []*role.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderInvitationRoles) Gt(value []*role.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderInvitationRoles) Gte(value []*role.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderInvitationRoles) Lt(value []*role.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderInvitationRoles) Lte(value []*role.Reference) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderInvitationRoles) In(values [][]*role.Reference) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().Roles().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationRoles) NotIn(values [][]*role.Reference) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().Roles().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationRoles) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().Roles().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationRoles) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().Roles().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationRoles) Contains(value *role.Reference) *FilterBuilder {
	return b.builder.addCond(&FilterConditionContains{
		Type:      gotenresource.ConditionContainsTypeValue,
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().Roles().FieldPath(),
		Value:     NewOrganizationInvitationFieldPathBuilder().Invitation().Roles().WithItemValue(value),
	})
}

func (b *filterCndBuilderInvitationRoles) ContainsAnyOf(values []*role.Reference) *FilterBuilder {
	pathSelector := NewOrganizationInvitationFieldPathBuilder().Invitation().Roles()
	itemValues := make([]OrganizationInvitation_FieldPathArrayItemValue, 0, len(values))
	for _, value := range values {
		itemValues = append(itemValues, pathSelector.WithItemValue(value))
	}
	return b.builder.addCond(&FilterConditionContains{
		Type:      gotenresource.ConditionContainsTypeAny,
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().Roles().FieldPath(),
		Values:    itemValues,
	})
}

func (b *filterCndBuilderInvitationRoles) ContainsAll(values []*role.Reference) *FilterBuilder {
	pathSelector := NewOrganizationInvitationFieldPathBuilder().Invitation().Roles()
	itemValues := make([]OrganizationInvitation_FieldPathArrayItemValue, 0, len(values))
	for _, value := range values {
		itemValues = append(itemValues, pathSelector.WithItemValue(value))
	}
	return b.builder.addCond(&FilterConditionContains{
		Type:      gotenresource.ConditionContainsTypeAll,
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().Roles().FieldPath(),
		Values:    itemValues,
	})
}

func (b *filterCndBuilderInvitationRoles) compare(op gotenfilter.CompareOperator, value []*role.Reference) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                              op,
		OrganizationInvitation_FieldPathValue: NewOrganizationInvitationFieldPathBuilder().Invitation().Roles().WithValue(value),
	})
}

type filterCndBuilderInvitationExpirationDate struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderInvitationExpirationDate) Eq(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderInvitationExpirationDate) Neq(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderInvitationExpirationDate) Gt(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderInvitationExpirationDate) Gte(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderInvitationExpirationDate) Lt(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderInvitationExpirationDate) Lte(value *timestamp.Timestamp) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderInvitationExpirationDate) In(values []*timestamp.Timestamp) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().ExpirationDate().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationExpirationDate) NotIn(values []*timestamp.Timestamp) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().ExpirationDate().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationExpirationDate) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().ExpirationDate().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationExpirationDate) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().ExpirationDate().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationExpirationDate) compare(op gotenfilter.CompareOperator, value *timestamp.Timestamp) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                              op,
		OrganizationInvitation_FieldPathValue: NewOrganizationInvitationFieldPathBuilder().Invitation().ExpirationDate().WithValue(value),
	})
}

type filterCndBuilderInvitationState struct {
	builder *FilterBuilder
}

func (b *filterCndBuilderInvitationState) Eq(value iam_common.Invitation_State) *FilterBuilder {
	return b.compare(gotenfilter.Eq, value)
}

func (b *filterCndBuilderInvitationState) Neq(value iam_common.Invitation_State) *FilterBuilder {
	return b.compare(gotenfilter.Neq, value)
}

func (b *filterCndBuilderInvitationState) Gt(value iam_common.Invitation_State) *FilterBuilder {
	return b.compare(gotenfilter.Gt, value)
}

func (b *filterCndBuilderInvitationState) Gte(value iam_common.Invitation_State) *FilterBuilder {
	return b.compare(gotenfilter.Gte, value)
}

func (b *filterCndBuilderInvitationState) Lt(value iam_common.Invitation_State) *FilterBuilder {
	return b.compare(gotenfilter.Lt, value)
}

func (b *filterCndBuilderInvitationState) Lte(value iam_common.Invitation_State) *FilterBuilder {
	return b.compare(gotenfilter.Lte, value)
}

func (b *filterCndBuilderInvitationState) In(values []iam_common.Invitation_State) *FilterBuilder {
	return b.builder.addCond(&FilterConditionIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().State().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationState) NotIn(values []iam_common.Invitation_State) *FilterBuilder {
	return b.builder.addCond(&FilterConditionNotIn{
		OrganizationInvitation_FieldPathArrayOfValues: NewOrganizationInvitationFieldPathBuilder().Invitation().State().WithArrayOfValues(values),
	})
}

func (b *filterCndBuilderInvitationState) IsNull() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNull{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().State().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationState) IsNan() *FilterBuilder {
	return b.builder.addCond(&FilterConditionIsNaN{
		FieldPath: NewOrganizationInvitationFieldPathBuilder().Invitation().State().FieldPath(),
	})
}

func (b *filterCndBuilderInvitationState) compare(op gotenfilter.CompareOperator, value iam_common.Invitation_State) *FilterBuilder {
	return b.builder.addCond(&FilterConditionCompare{
		Operator:                              op,
		OrganizationInvitation_FieldPathValue: NewOrganizationInvitationFieldPathBuilder().Invitation().State().WithValue(value),
	})
}
