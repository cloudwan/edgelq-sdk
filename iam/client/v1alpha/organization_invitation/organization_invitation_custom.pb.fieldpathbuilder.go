// Code generated by protoc-gen-goten-object
// File: edgelq/iam/proto/v1alpha/organization_invitation_custom.proto
// DO NOT EDIT!!!

package organization_invitation_client

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/common"
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/condition"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	organization_invitation "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization_invitation"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/permission"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	project_invitation "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project_invitation"
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
	_ = &organization_invitation.OrganizationInvitation{}
	_ = &permission.Permission{}
	_ = &project.Project{}
	_ = &project_invitation.ProjectInvitation{}
	_ = &role.Role{}
	_ = &service_account.ServiceAccount{}
	_ = &user.User{}
	_ = &timestamp.Timestamp{}
)

type AcceptOrganizationInvitationRequestFieldPathBuilder struct{}

func NewAcceptOrganizationInvitationRequestFieldPathBuilder() AcceptOrganizationInvitationRequestFieldPathBuilder {
	return AcceptOrganizationInvitationRequestFieldPathBuilder{}
}
func (AcceptOrganizationInvitationRequestFieldPathBuilder) Name() AcceptOrganizationInvitationRequestPathSelectorName {
	return AcceptOrganizationInvitationRequestPathSelectorName{}
}

type AcceptOrganizationInvitationRequestPathSelectorName struct{}

func (AcceptOrganizationInvitationRequestPathSelectorName) FieldPath() *AcceptOrganizationInvitationRequest_FieldTerminalPath {
	return &AcceptOrganizationInvitationRequest_FieldTerminalPath{selector: AcceptOrganizationInvitationRequest_FieldPathSelectorName}
}

func (s AcceptOrganizationInvitationRequestPathSelectorName) WithValue(value *organization_invitation.Reference) *AcceptOrganizationInvitationRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*AcceptOrganizationInvitationRequest_FieldTerminalPathValue)
}

func (s AcceptOrganizationInvitationRequestPathSelectorName) WithArrayOfValues(values []*organization_invitation.Reference) *AcceptOrganizationInvitationRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*AcceptOrganizationInvitationRequest_FieldTerminalPathArrayOfValues)
}

type AcceptOrganizationInvitationResponseFieldPathBuilder struct{}

func NewAcceptOrganizationInvitationResponseFieldPathBuilder() AcceptOrganizationInvitationResponseFieldPathBuilder {
	return AcceptOrganizationInvitationResponseFieldPathBuilder{}
}

type DeclineOrganizationInvitationRequestFieldPathBuilder struct{}

func NewDeclineOrganizationInvitationRequestFieldPathBuilder() DeclineOrganizationInvitationRequestFieldPathBuilder {
	return DeclineOrganizationInvitationRequestFieldPathBuilder{}
}
func (DeclineOrganizationInvitationRequestFieldPathBuilder) Name() DeclineOrganizationInvitationRequestPathSelectorName {
	return DeclineOrganizationInvitationRequestPathSelectorName{}
}
func (DeclineOrganizationInvitationRequestFieldPathBuilder) Filter() DeclineOrganizationInvitationRequestPathSelectorFilter {
	return DeclineOrganizationInvitationRequestPathSelectorFilter{}
}

type DeclineOrganizationInvitationRequestPathSelectorName struct{}

func (DeclineOrganizationInvitationRequestPathSelectorName) FieldPath() *DeclineOrganizationInvitationRequest_FieldTerminalPath {
	return &DeclineOrganizationInvitationRequest_FieldTerminalPath{selector: DeclineOrganizationInvitationRequest_FieldPathSelectorName}
}

func (s DeclineOrganizationInvitationRequestPathSelectorName) WithValue(value *organization_invitation.Reference) *DeclineOrganizationInvitationRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*DeclineOrganizationInvitationRequest_FieldTerminalPathValue)
}

func (s DeclineOrganizationInvitationRequestPathSelectorName) WithArrayOfValues(values []*organization_invitation.Reference) *DeclineOrganizationInvitationRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeclineOrganizationInvitationRequest_FieldTerminalPathArrayOfValues)
}

type DeclineOrganizationInvitationRequestPathSelectorFilter struct{}

func (DeclineOrganizationInvitationRequestPathSelectorFilter) FieldPath() *DeclineOrganizationInvitationRequest_FieldTerminalPath {
	return &DeclineOrganizationInvitationRequest_FieldTerminalPath{selector: DeclineOrganizationInvitationRequest_FieldPathSelectorFilter}
}

func (s DeclineOrganizationInvitationRequestPathSelectorFilter) WithValue(value *project_invitation.Filter) *DeclineOrganizationInvitationRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*DeclineOrganizationInvitationRequest_FieldTerminalPathValue)
}

func (s DeclineOrganizationInvitationRequestPathSelectorFilter) WithArrayOfValues(values []*project_invitation.Filter) *DeclineOrganizationInvitationRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*DeclineOrganizationInvitationRequest_FieldTerminalPathArrayOfValues)
}

type DeclineOrganizationInvitationResponseFieldPathBuilder struct{}

func NewDeclineOrganizationInvitationResponseFieldPathBuilder() DeclineOrganizationInvitationResponseFieldPathBuilder {
	return DeclineOrganizationInvitationResponseFieldPathBuilder{}
}

type ListMyOrganizationInvitationsRequestFieldPathBuilder struct{}

func NewListMyOrganizationInvitationsRequestFieldPathBuilder() ListMyOrganizationInvitationsRequestFieldPathBuilder {
	return ListMyOrganizationInvitationsRequestFieldPathBuilder{}
}
func (ListMyOrganizationInvitationsRequestFieldPathBuilder) Parent() ListMyOrganizationInvitationsRequestPathSelectorParent {
	return ListMyOrganizationInvitationsRequestPathSelectorParent{}
}
func (ListMyOrganizationInvitationsRequestFieldPathBuilder) Filter() ListMyOrganizationInvitationsRequestPathSelectorFilter {
	return ListMyOrganizationInvitationsRequestPathSelectorFilter{}
}

type ListMyOrganizationInvitationsRequestPathSelectorParent struct{}

func (ListMyOrganizationInvitationsRequestPathSelectorParent) FieldPath() *ListMyOrganizationInvitationsRequest_FieldTerminalPath {
	return &ListMyOrganizationInvitationsRequest_FieldTerminalPath{selector: ListMyOrganizationInvitationsRequest_FieldPathSelectorParent}
}

func (s ListMyOrganizationInvitationsRequestPathSelectorParent) WithValue(value *organization_invitation.ParentReference) *ListMyOrganizationInvitationsRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ListMyOrganizationInvitationsRequest_FieldTerminalPathValue)
}

func (s ListMyOrganizationInvitationsRequestPathSelectorParent) WithArrayOfValues(values []*organization_invitation.ParentReference) *ListMyOrganizationInvitationsRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ListMyOrganizationInvitationsRequest_FieldTerminalPathArrayOfValues)
}

type ListMyOrganizationInvitationsRequestPathSelectorFilter struct{}

func (ListMyOrganizationInvitationsRequestPathSelectorFilter) FieldPath() *ListMyOrganizationInvitationsRequest_FieldTerminalPath {
	return &ListMyOrganizationInvitationsRequest_FieldTerminalPath{selector: ListMyOrganizationInvitationsRequest_FieldPathSelectorFilter}
}

func (s ListMyOrganizationInvitationsRequestPathSelectorFilter) WithValue(value *organization_invitation.Filter) *ListMyOrganizationInvitationsRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ListMyOrganizationInvitationsRequest_FieldTerminalPathValue)
}

func (s ListMyOrganizationInvitationsRequestPathSelectorFilter) WithArrayOfValues(values []*organization_invitation.Filter) *ListMyOrganizationInvitationsRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ListMyOrganizationInvitationsRequest_FieldTerminalPathArrayOfValues)
}

type ListMyOrganizationInvitationsResponseFieldPathBuilder struct{}

func NewListMyOrganizationInvitationsResponseFieldPathBuilder() ListMyOrganizationInvitationsResponseFieldPathBuilder {
	return ListMyOrganizationInvitationsResponseFieldPathBuilder{}
}
func (ListMyOrganizationInvitationsResponseFieldPathBuilder) OrganizationInvitations() ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitations {
	return ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitations{}
}

type ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitations struct{}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitations) FieldPath() *ListMyOrganizationInvitationsResponse_FieldTerminalPath {
	return &ListMyOrganizationInvitationsResponse_FieldTerminalPath{selector: ListMyOrganizationInvitationsResponse_FieldPathSelectorOrganizationInvitations}
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitations) WithValue(value []*organization_invitation.OrganizationInvitation) *ListMyOrganizationInvitationsResponse_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ListMyOrganizationInvitationsResponse_FieldTerminalPathValue)
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitations) WithArrayOfValues(values [][]*organization_invitation.OrganizationInvitation) *ListMyOrganizationInvitationsResponse_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ListMyOrganizationInvitationsResponse_FieldTerminalPathArrayOfValues)
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitations) WithItemValue(value *organization_invitation.OrganizationInvitation) *ListMyOrganizationInvitationsResponse_FieldTerminalPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*ListMyOrganizationInvitationsResponse_FieldTerminalPathArrayItemValue)
}
func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitations) WithSubPath(subPath organization_invitation.OrganizationInvitation_FieldPath) *ListMyOrganizationInvitationsResponse_FieldSubPath {
	return &ListMyOrganizationInvitationsResponse_FieldSubPath{selector: ListMyOrganizationInvitationsResponse_FieldPathSelectorOrganizationInvitations, subPath: subPath}
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitations) WithSubValue(subPathValue organization_invitation.OrganizationInvitation_FieldPathValue) *ListMyOrganizationInvitationsResponse_FieldSubPathValue {
	return &ListMyOrganizationInvitationsResponse_FieldSubPathValue{ListMyOrganizationInvitationsResponse_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitations) WithSubArrayOfValues(subPathArrayOfValues organization_invitation.OrganizationInvitation_FieldPathArrayOfValues) *ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues {
	return &ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues{ListMyOrganizationInvitationsResponse_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitations) WithSubArrayItemValue(subPathArrayItemValue organization_invitation.OrganizationInvitation_FieldPathArrayItemValue) *ListMyOrganizationInvitationsResponse_FieldSubPathArrayItemValue {
	return &ListMyOrganizationInvitationsResponse_FieldSubPathArrayItemValue{ListMyOrganizationInvitationsResponse_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitations) Name() ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsName {
	return ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsName{}
}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitations) Invitation() ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitation {
	return ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitation{}
}

type ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsName struct{}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsName) FieldPath() *ListMyOrganizationInvitationsResponse_FieldSubPath {
	return &ListMyOrganizationInvitationsResponse_FieldSubPath{
		selector: ListMyOrganizationInvitationsResponse_FieldPathSelectorOrganizationInvitations,
		subPath:  organization_invitation.NewOrganizationInvitationFieldPathBuilder().Name().FieldPath(),
	}
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsName) WithValue(value *organization_invitation.Name) *ListMyOrganizationInvitationsResponse_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ListMyOrganizationInvitationsResponse_FieldSubPathValue)
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsName) WithArrayOfValues(values []*organization_invitation.Name) *ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues)
}

type ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitation struct{}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitation) FieldPath() *ListMyOrganizationInvitationsResponse_FieldSubPath {
	return &ListMyOrganizationInvitationsResponse_FieldSubPath{
		selector: ListMyOrganizationInvitationsResponse_FieldPathSelectorOrganizationInvitations,
		subPath:  organization_invitation.NewOrganizationInvitationFieldPathBuilder().Invitation().FieldPath(),
	}
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitation) WithValue(value *iam_common.Invitation) *ListMyOrganizationInvitationsResponse_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ListMyOrganizationInvitationsResponse_FieldSubPathValue)
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitation) WithArrayOfValues(values []*iam_common.Invitation) *ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues)
}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitation) InviteeEmail() ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviteeEmail {
	return ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviteeEmail{}
}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitation) InviterActor() ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActor {
	return ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActor{}
}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitation) InviterFullName() ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterFullName {
	return ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterFullName{}
}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitation) InviterEmail() ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterEmail {
	return ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterEmail{}
}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitation) Roles() ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationRoles {
	return ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationRoles{}
}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitation) ExpirationDate() ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationExpirationDate {
	return ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationExpirationDate{}
}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitation) State() ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationState {
	return ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationState{}
}

type ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviteeEmail struct{}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviteeEmail) FieldPath() *ListMyOrganizationInvitationsResponse_FieldSubPath {
	return &ListMyOrganizationInvitationsResponse_FieldSubPath{
		selector: ListMyOrganizationInvitationsResponse_FieldPathSelectorOrganizationInvitations,
		subPath:  organization_invitation.NewOrganizationInvitationFieldPathBuilder().Invitation().InviteeEmail().FieldPath(),
	}
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviteeEmail) WithValue(value string) *ListMyOrganizationInvitationsResponse_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ListMyOrganizationInvitationsResponse_FieldSubPathValue)
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviteeEmail) WithArrayOfValues(values []string) *ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues)
}

type ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActor struct{}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActor) FieldPath() *ListMyOrganizationInvitationsResponse_FieldSubPath {
	return &ListMyOrganizationInvitationsResponse_FieldSubPath{
		selector: ListMyOrganizationInvitationsResponse_FieldPathSelectorOrganizationInvitations,
		subPath:  organization_invitation.NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().FieldPath(),
	}
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActor) WithValue(value *iam_common.Actor) *ListMyOrganizationInvitationsResponse_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ListMyOrganizationInvitationsResponse_FieldSubPathValue)
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActor) WithArrayOfValues(values []*iam_common.Actor) *ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues)
}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActor) User() ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActorUser {
	return ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActorUser{}
}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActor) ServiceAccount() ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActorServiceAccount {
	return ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActorServiceAccount{}
}

type ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActorUser struct{}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActorUser) FieldPath() *ListMyOrganizationInvitationsResponse_FieldSubPath {
	return &ListMyOrganizationInvitationsResponse_FieldSubPath{
		selector: ListMyOrganizationInvitationsResponse_FieldPathSelectorOrganizationInvitations,
		subPath:  organization_invitation.NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().User().FieldPath(),
	}
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActorUser) WithValue(value *user.Reference) *ListMyOrganizationInvitationsResponse_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ListMyOrganizationInvitationsResponse_FieldSubPathValue)
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActorUser) WithArrayOfValues(values []*user.Reference) *ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues)
}

type ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActorServiceAccount struct{}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActorServiceAccount) FieldPath() *ListMyOrganizationInvitationsResponse_FieldSubPath {
	return &ListMyOrganizationInvitationsResponse_FieldSubPath{
		selector: ListMyOrganizationInvitationsResponse_FieldPathSelectorOrganizationInvitations,
		subPath:  organization_invitation.NewOrganizationInvitationFieldPathBuilder().Invitation().InviterActor().ServiceAccount().FieldPath(),
	}
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActorServiceAccount) WithValue(value *service_account.Reference) *ListMyOrganizationInvitationsResponse_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ListMyOrganizationInvitationsResponse_FieldSubPathValue)
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterActorServiceAccount) WithArrayOfValues(values []*service_account.Reference) *ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues)
}

type ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterFullName struct{}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterFullName) FieldPath() *ListMyOrganizationInvitationsResponse_FieldSubPath {
	return &ListMyOrganizationInvitationsResponse_FieldSubPath{
		selector: ListMyOrganizationInvitationsResponse_FieldPathSelectorOrganizationInvitations,
		subPath:  organization_invitation.NewOrganizationInvitationFieldPathBuilder().Invitation().InviterFullName().FieldPath(),
	}
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterFullName) WithValue(value string) *ListMyOrganizationInvitationsResponse_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ListMyOrganizationInvitationsResponse_FieldSubPathValue)
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterFullName) WithArrayOfValues(values []string) *ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues)
}

type ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterEmail struct{}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterEmail) FieldPath() *ListMyOrganizationInvitationsResponse_FieldSubPath {
	return &ListMyOrganizationInvitationsResponse_FieldSubPath{
		selector: ListMyOrganizationInvitationsResponse_FieldPathSelectorOrganizationInvitations,
		subPath:  organization_invitation.NewOrganizationInvitationFieldPathBuilder().Invitation().InviterEmail().FieldPath(),
	}
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterEmail) WithValue(value string) *ListMyOrganizationInvitationsResponse_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ListMyOrganizationInvitationsResponse_FieldSubPathValue)
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationInviterEmail) WithArrayOfValues(values []string) *ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues)
}

type ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationRoles struct{}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationRoles) FieldPath() *ListMyOrganizationInvitationsResponse_FieldSubPath {
	return &ListMyOrganizationInvitationsResponse_FieldSubPath{
		selector: ListMyOrganizationInvitationsResponse_FieldPathSelectorOrganizationInvitations,
		subPath:  organization_invitation.NewOrganizationInvitationFieldPathBuilder().Invitation().Roles().FieldPath(),
	}
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationRoles) WithValue(value []*role.Reference) *ListMyOrganizationInvitationsResponse_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ListMyOrganizationInvitationsResponse_FieldSubPathValue)
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationRoles) WithArrayOfValues(values [][]*role.Reference) *ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues)
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationRoles) WithItemValue(value *role.Reference) *ListMyOrganizationInvitationsResponse_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*ListMyOrganizationInvitationsResponse_FieldSubPathArrayItemValue)
}

type ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationExpirationDate struct{}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationExpirationDate) FieldPath() *ListMyOrganizationInvitationsResponse_FieldSubPath {
	return &ListMyOrganizationInvitationsResponse_FieldSubPath{
		selector: ListMyOrganizationInvitationsResponse_FieldPathSelectorOrganizationInvitations,
		subPath:  organization_invitation.NewOrganizationInvitationFieldPathBuilder().Invitation().ExpirationDate().FieldPath(),
	}
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationExpirationDate) WithValue(value *timestamp.Timestamp) *ListMyOrganizationInvitationsResponse_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ListMyOrganizationInvitationsResponse_FieldSubPathValue)
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationExpirationDate) WithArrayOfValues(values []*timestamp.Timestamp) *ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues)
}

type ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationState struct{}

func (ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationState) FieldPath() *ListMyOrganizationInvitationsResponse_FieldSubPath {
	return &ListMyOrganizationInvitationsResponse_FieldSubPath{
		selector: ListMyOrganizationInvitationsResponse_FieldPathSelectorOrganizationInvitations,
		subPath:  organization_invitation.NewOrganizationInvitationFieldPathBuilder().Invitation().State().FieldPath(),
	}
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationState) WithValue(value iam_common.Invitation_State) *ListMyOrganizationInvitationsResponse_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ListMyOrganizationInvitationsResponse_FieldSubPathValue)
}

func (s ListMyOrganizationInvitationsResponsePathSelectorOrganizationInvitationsInvitationState) WithArrayOfValues(values []iam_common.Invitation_State) *ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ListMyOrganizationInvitationsResponse_FieldSubPathArrayOfValues)
}
