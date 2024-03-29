// Code generated by protoc-gen-goten-access
// Service: IAM
// DO NOT EDIT!!!

package iam_access

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	attestation_domain_access "github.com/cloudwan/edgelq-sdk/iam/access/v1/attestation_domain"
	condition_access "github.com/cloudwan/edgelq-sdk/iam/access/v1/condition"
	group_access "github.com/cloudwan/edgelq-sdk/iam/access/v1/group"
	group_member_access "github.com/cloudwan/edgelq-sdk/iam/access/v1/group_member"
	organization_access "github.com/cloudwan/edgelq-sdk/iam/access/v1/organization"
	organization_invitation_access "github.com/cloudwan/edgelq-sdk/iam/access/v1/organization_invitation"
	permission_access "github.com/cloudwan/edgelq-sdk/iam/access/v1/permission"
	project_access "github.com/cloudwan/edgelq-sdk/iam/access/v1/project"
	project_invitation_access "github.com/cloudwan/edgelq-sdk/iam/access/v1/project_invitation"
	role_access "github.com/cloudwan/edgelq-sdk/iam/access/v1/role"
	role_binding_access "github.com/cloudwan/edgelq-sdk/iam/access/v1/role_binding"
	service_account_access "github.com/cloudwan/edgelq-sdk/iam/access/v1/service_account"
	service_account_key_access "github.com/cloudwan/edgelq-sdk/iam/access/v1/service_account_key"
	user_access "github.com/cloudwan/edgelq-sdk/iam/access/v1/user"
	iam_client "github.com/cloudwan/edgelq-sdk/iam/client/v1/iam"
	attestation_domain "github.com/cloudwan/edgelq-sdk/iam/resources/v1/attestation_domain"
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1/condition"
	group "github.com/cloudwan/edgelq-sdk/iam/resources/v1/group"
	group_member "github.com/cloudwan/edgelq-sdk/iam/resources/v1/group_member"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	organization_invitation "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization_invitation"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1/permission"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	project_invitation "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project_invitation"
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1/role"
	role_binding "github.com/cloudwan/edgelq-sdk/iam/resources/v1/role_binding"
	service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1/service_account"
	service_account_key "github.com/cloudwan/edgelq-sdk/iam/resources/v1/service_account_key"
	user "github.com/cloudwan/edgelq-sdk/iam/resources/v1/user"
)

type IAMApiAccess interface {
	gotenresource.Access

	attestation_domain.AttestationDomainAccess
	condition.ConditionAccess
	group.GroupAccess
	group_member.GroupMemberAccess
	organization.OrganizationAccess
	organization_invitation.OrganizationInvitationAccess
	permission.PermissionAccess
	project.ProjectAccess
	project_invitation.ProjectInvitationAccess
	role.RoleAccess
	role_binding.RoleBindingAccess
	service_account.ServiceAccountAccess
	service_account_key.ServiceAccountKeyAccess
	user.UserAccess
}

type apiIAMAccess struct {
	gotenresource.Access

	attestation_domain.AttestationDomainAccess
	condition.ConditionAccess
	group.GroupAccess
	group_member.GroupMemberAccess
	organization.OrganizationAccess
	organization_invitation.OrganizationInvitationAccess
	permission.PermissionAccess
	project.ProjectAccess
	project_invitation.ProjectInvitationAccess
	role.RoleAccess
	role_binding.RoleBindingAccess
	service_account.ServiceAccountAccess
	service_account_key.ServiceAccountKeyAccess
	user.UserAccess
}

func NewApiAccess(client iam_client.IAMClient) IAMApiAccess {

	attestationDomainAccess := attestation_domain_access.NewApiAttestationDomainAccess(client)
	conditionAccess := condition_access.NewApiConditionAccess(client)
	groupAccess := group_access.NewApiGroupAccess(client)
	groupMemberAccess := group_member_access.NewApiGroupMemberAccess(client)
	organizationAccess := organization_access.NewApiOrganizationAccess(client)
	organizationInvitationAccess := organization_invitation_access.NewApiOrganizationInvitationAccess(client)
	permissionAccess := permission_access.NewApiPermissionAccess(client)
	projectAccess := project_access.NewApiProjectAccess(client)
	projectInvitationAccess := project_invitation_access.NewApiProjectInvitationAccess(client)
	roleAccess := role_access.NewApiRoleAccess(client)
	roleBindingAccess := role_binding_access.NewApiRoleBindingAccess(client)
	serviceAccountAccess := service_account_access.NewApiServiceAccountAccess(client)
	serviceAccountKeyAccess := service_account_key_access.NewApiServiceAccountKeyAccess(client)
	userAccess := user_access.NewApiUserAccess(client)

	return &apiIAMAccess{
		Access: gotenresource.NewCompositeAccess(

			attestation_domain.AsAnyCastAccess(attestationDomainAccess),
			condition.AsAnyCastAccess(conditionAccess),
			group.AsAnyCastAccess(groupAccess),
			group_member.AsAnyCastAccess(groupMemberAccess),
			organization.AsAnyCastAccess(organizationAccess),
			organization_invitation.AsAnyCastAccess(organizationInvitationAccess),
			permission.AsAnyCastAccess(permissionAccess),
			project.AsAnyCastAccess(projectAccess),
			project_invitation.AsAnyCastAccess(projectInvitationAccess),
			role.AsAnyCastAccess(roleAccess),
			role_binding.AsAnyCastAccess(roleBindingAccess),
			service_account.AsAnyCastAccess(serviceAccountAccess),
			service_account_key.AsAnyCastAccess(serviceAccountKeyAccess),
			user.AsAnyCastAccess(userAccess),
		),

		AttestationDomainAccess:      attestationDomainAccess,
		ConditionAccess:              conditionAccess,
		GroupAccess:                  groupAccess,
		GroupMemberAccess:            groupMemberAccess,
		OrganizationAccess:           organizationAccess,
		OrganizationInvitationAccess: organizationInvitationAccess,
		PermissionAccess:             permissionAccess,
		ProjectAccess:                projectAccess,
		ProjectInvitationAccess:      projectInvitationAccess,
		RoleAccess:                   roleAccess,
		RoleBindingAccess:            roleBindingAccess,
		ServiceAccountAccess:         serviceAccountAccess,
		ServiceAccountKeyAccess:      serviceAccountKeyAccess,
		UserAccess:                   userAccess,
	}
}
