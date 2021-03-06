// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha/organization_invitation.proto
// DO NOT EDIT!!!

package organization_invitation

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

type OrganizationInvitationFieldPathBuilder struct{}

func NewOrganizationInvitationFieldPathBuilder() OrganizationInvitationFieldPathBuilder {
	return OrganizationInvitationFieldPathBuilder{}
}
func (OrganizationInvitationFieldPathBuilder) Name() OrganizationInvitationPathSelectorName {
	return OrganizationInvitationPathSelectorName{}
}
func (OrganizationInvitationFieldPathBuilder) Invitation() OrganizationInvitationPathSelectorInvitation {
	return OrganizationInvitationPathSelectorInvitation{}
}

type OrganizationInvitationPathSelectorName struct{}

func (OrganizationInvitationPathSelectorName) FieldPath() *OrganizationInvitation_FieldTerminalPath {
	return &OrganizationInvitation_FieldTerminalPath{selector: OrganizationInvitation_FieldPathSelectorName}
}

func (s OrganizationInvitationPathSelectorName) WithValue(value *Name) *OrganizationInvitation_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*OrganizationInvitation_FieldTerminalPathValue)
}

func (s OrganizationInvitationPathSelectorName) WithArrayOfValues(values []*Name) *OrganizationInvitation_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OrganizationInvitation_FieldTerminalPathArrayOfValues)
}

type OrganizationInvitationPathSelectorInvitation struct{}

func (OrganizationInvitationPathSelectorInvitation) FieldPath() *OrganizationInvitation_FieldTerminalPath {
	return &OrganizationInvitation_FieldTerminalPath{selector: OrganizationInvitation_FieldPathSelectorInvitation}
}

func (s OrganizationInvitationPathSelectorInvitation) WithValue(value *iam_common.Invitation) *OrganizationInvitation_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*OrganizationInvitation_FieldTerminalPathValue)
}

func (s OrganizationInvitationPathSelectorInvitation) WithArrayOfValues(values []*iam_common.Invitation) *OrganizationInvitation_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OrganizationInvitation_FieldTerminalPathArrayOfValues)
}

func (OrganizationInvitationPathSelectorInvitation) WithSubPath(subPath iam_common.Invitation_FieldPath) *OrganizationInvitation_FieldSubPath {
	return &OrganizationInvitation_FieldSubPath{selector: OrganizationInvitation_FieldPathSelectorInvitation, subPath: subPath}
}

func (s OrganizationInvitationPathSelectorInvitation) WithSubValue(subPathValue iam_common.Invitation_FieldPathValue) *OrganizationInvitation_FieldSubPathValue {
	return &OrganizationInvitation_FieldSubPathValue{OrganizationInvitation_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s OrganizationInvitationPathSelectorInvitation) WithSubArrayOfValues(subPathArrayOfValues iam_common.Invitation_FieldPathArrayOfValues) *OrganizationInvitation_FieldSubPathArrayOfValues {
	return &OrganizationInvitation_FieldSubPathArrayOfValues{OrganizationInvitation_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s OrganizationInvitationPathSelectorInvitation) WithSubArrayItemValue(subPathArrayItemValue iam_common.Invitation_FieldPathArrayItemValue) *OrganizationInvitation_FieldSubPathArrayItemValue {
	return &OrganizationInvitation_FieldSubPathArrayItemValue{OrganizationInvitation_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (OrganizationInvitationPathSelectorInvitation) InviteeEmail() OrganizationInvitationPathSelectorInvitationInviteeEmail {
	return OrganizationInvitationPathSelectorInvitationInviteeEmail{}
}

func (OrganizationInvitationPathSelectorInvitation) InviterActor() OrganizationInvitationPathSelectorInvitationInviterActor {
	return OrganizationInvitationPathSelectorInvitationInviterActor{}
}

func (OrganizationInvitationPathSelectorInvitation) InviterFullName() OrganizationInvitationPathSelectorInvitationInviterFullName {
	return OrganizationInvitationPathSelectorInvitationInviterFullName{}
}

func (OrganizationInvitationPathSelectorInvitation) InviterEmail() OrganizationInvitationPathSelectorInvitationInviterEmail {
	return OrganizationInvitationPathSelectorInvitationInviterEmail{}
}

func (OrganizationInvitationPathSelectorInvitation) Roles() OrganizationInvitationPathSelectorInvitationRoles {
	return OrganizationInvitationPathSelectorInvitationRoles{}
}

func (OrganizationInvitationPathSelectorInvitation) ExpirationDate() OrganizationInvitationPathSelectorInvitationExpirationDate {
	return OrganizationInvitationPathSelectorInvitationExpirationDate{}
}

func (OrganizationInvitationPathSelectorInvitation) State() OrganizationInvitationPathSelectorInvitationState {
	return OrganizationInvitationPathSelectorInvitationState{}
}

type OrganizationInvitationPathSelectorInvitationInviteeEmail struct{}

func (OrganizationInvitationPathSelectorInvitationInviteeEmail) FieldPath() *OrganizationInvitation_FieldSubPath {
	return &OrganizationInvitation_FieldSubPath{
		selector: OrganizationInvitation_FieldPathSelectorInvitation,
		subPath:  iam_common.NewInvitationFieldPathBuilder().InviteeEmail().FieldPath(),
	}
}

func (s OrganizationInvitationPathSelectorInvitationInviteeEmail) WithValue(value string) *OrganizationInvitation_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OrganizationInvitation_FieldSubPathValue)
}

func (s OrganizationInvitationPathSelectorInvitationInviteeEmail) WithArrayOfValues(values []string) *OrganizationInvitation_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OrganizationInvitation_FieldSubPathArrayOfValues)
}

type OrganizationInvitationPathSelectorInvitationInviterActor struct{}

func (OrganizationInvitationPathSelectorInvitationInviterActor) FieldPath() *OrganizationInvitation_FieldSubPath {
	return &OrganizationInvitation_FieldSubPath{
		selector: OrganizationInvitation_FieldPathSelectorInvitation,
		subPath:  iam_common.NewInvitationFieldPathBuilder().InviterActor().FieldPath(),
	}
}

func (s OrganizationInvitationPathSelectorInvitationInviterActor) WithValue(value *iam_common.Actor) *OrganizationInvitation_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OrganizationInvitation_FieldSubPathValue)
}

func (s OrganizationInvitationPathSelectorInvitationInviterActor) WithArrayOfValues(values []*iam_common.Actor) *OrganizationInvitation_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OrganizationInvitation_FieldSubPathArrayOfValues)
}

func (OrganizationInvitationPathSelectorInvitationInviterActor) User() OrganizationInvitationPathSelectorInvitationInviterActorUser {
	return OrganizationInvitationPathSelectorInvitationInviterActorUser{}
}

func (OrganizationInvitationPathSelectorInvitationInviterActor) ServiceAccount() OrganizationInvitationPathSelectorInvitationInviterActorServiceAccount {
	return OrganizationInvitationPathSelectorInvitationInviterActorServiceAccount{}
}

type OrganizationInvitationPathSelectorInvitationInviterActorUser struct{}

func (OrganizationInvitationPathSelectorInvitationInviterActorUser) FieldPath() *OrganizationInvitation_FieldSubPath {
	return &OrganizationInvitation_FieldSubPath{
		selector: OrganizationInvitation_FieldPathSelectorInvitation,
		subPath:  iam_common.NewInvitationFieldPathBuilder().InviterActor().User().FieldPath(),
	}
}

func (s OrganizationInvitationPathSelectorInvitationInviterActorUser) WithValue(value *user.Reference) *OrganizationInvitation_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OrganizationInvitation_FieldSubPathValue)
}

func (s OrganizationInvitationPathSelectorInvitationInviterActorUser) WithArrayOfValues(values []*user.Reference) *OrganizationInvitation_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OrganizationInvitation_FieldSubPathArrayOfValues)
}

type OrganizationInvitationPathSelectorInvitationInviterActorServiceAccount struct{}

func (OrganizationInvitationPathSelectorInvitationInviterActorServiceAccount) FieldPath() *OrganizationInvitation_FieldSubPath {
	return &OrganizationInvitation_FieldSubPath{
		selector: OrganizationInvitation_FieldPathSelectorInvitation,
		subPath:  iam_common.NewInvitationFieldPathBuilder().InviterActor().ServiceAccount().FieldPath(),
	}
}

func (s OrganizationInvitationPathSelectorInvitationInviterActorServiceAccount) WithValue(value *service_account.Reference) *OrganizationInvitation_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OrganizationInvitation_FieldSubPathValue)
}

func (s OrganizationInvitationPathSelectorInvitationInviterActorServiceAccount) WithArrayOfValues(values []*service_account.Reference) *OrganizationInvitation_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OrganizationInvitation_FieldSubPathArrayOfValues)
}

type OrganizationInvitationPathSelectorInvitationInviterFullName struct{}

func (OrganizationInvitationPathSelectorInvitationInviterFullName) FieldPath() *OrganizationInvitation_FieldSubPath {
	return &OrganizationInvitation_FieldSubPath{
		selector: OrganizationInvitation_FieldPathSelectorInvitation,
		subPath:  iam_common.NewInvitationFieldPathBuilder().InviterFullName().FieldPath(),
	}
}

func (s OrganizationInvitationPathSelectorInvitationInviterFullName) WithValue(value string) *OrganizationInvitation_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OrganizationInvitation_FieldSubPathValue)
}

func (s OrganizationInvitationPathSelectorInvitationInviterFullName) WithArrayOfValues(values []string) *OrganizationInvitation_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OrganizationInvitation_FieldSubPathArrayOfValues)
}

type OrganizationInvitationPathSelectorInvitationInviterEmail struct{}

func (OrganizationInvitationPathSelectorInvitationInviterEmail) FieldPath() *OrganizationInvitation_FieldSubPath {
	return &OrganizationInvitation_FieldSubPath{
		selector: OrganizationInvitation_FieldPathSelectorInvitation,
		subPath:  iam_common.NewInvitationFieldPathBuilder().InviterEmail().FieldPath(),
	}
}

func (s OrganizationInvitationPathSelectorInvitationInviterEmail) WithValue(value string) *OrganizationInvitation_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OrganizationInvitation_FieldSubPathValue)
}

func (s OrganizationInvitationPathSelectorInvitationInviterEmail) WithArrayOfValues(values []string) *OrganizationInvitation_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OrganizationInvitation_FieldSubPathArrayOfValues)
}

type OrganizationInvitationPathSelectorInvitationRoles struct{}

func (OrganizationInvitationPathSelectorInvitationRoles) FieldPath() *OrganizationInvitation_FieldSubPath {
	return &OrganizationInvitation_FieldSubPath{
		selector: OrganizationInvitation_FieldPathSelectorInvitation,
		subPath:  iam_common.NewInvitationFieldPathBuilder().Roles().FieldPath(),
	}
}

func (s OrganizationInvitationPathSelectorInvitationRoles) WithValue(value []*role.Reference) *OrganizationInvitation_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OrganizationInvitation_FieldSubPathValue)
}

func (s OrganizationInvitationPathSelectorInvitationRoles) WithArrayOfValues(values [][]*role.Reference) *OrganizationInvitation_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OrganizationInvitation_FieldSubPathArrayOfValues)
}

func (s OrganizationInvitationPathSelectorInvitationRoles) WithItemValue(value *role.Reference) *OrganizationInvitation_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*OrganizationInvitation_FieldSubPathArrayItemValue)
}

type OrganizationInvitationPathSelectorInvitationExpirationDate struct{}

func (OrganizationInvitationPathSelectorInvitationExpirationDate) FieldPath() *OrganizationInvitation_FieldSubPath {
	return &OrganizationInvitation_FieldSubPath{
		selector: OrganizationInvitation_FieldPathSelectorInvitation,
		subPath:  iam_common.NewInvitationFieldPathBuilder().ExpirationDate().FieldPath(),
	}
}

func (s OrganizationInvitationPathSelectorInvitationExpirationDate) WithValue(value *timestamp.Timestamp) *OrganizationInvitation_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OrganizationInvitation_FieldSubPathValue)
}

func (s OrganizationInvitationPathSelectorInvitationExpirationDate) WithArrayOfValues(values []*timestamp.Timestamp) *OrganizationInvitation_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OrganizationInvitation_FieldSubPathArrayOfValues)
}

type OrganizationInvitationPathSelectorInvitationState struct{}

func (OrganizationInvitationPathSelectorInvitationState) FieldPath() *OrganizationInvitation_FieldSubPath {
	return &OrganizationInvitation_FieldSubPath{
		selector: OrganizationInvitation_FieldPathSelectorInvitation,
		subPath:  iam_common.NewInvitationFieldPathBuilder().State().FieldPath(),
	}
}

func (s OrganizationInvitationPathSelectorInvitationState) WithValue(value iam_common.Invitation_State) *OrganizationInvitation_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*OrganizationInvitation_FieldSubPathValue)
}

func (s OrganizationInvitationPathSelectorInvitationState) WithArrayOfValues(values []iam_common.Invitation_State) *OrganizationInvitation_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*OrganizationInvitation_FieldSubPathArrayOfValues)
}
