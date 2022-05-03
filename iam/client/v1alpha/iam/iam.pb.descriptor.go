// Code generated by protoc-gen-goten-client
// Service: IAM
// DO NOT EDIT!!!

package iam_client

import (
	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	authorization_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha/authorization"
	condition_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha/condition"
	group_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha/group"
	group_member_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha/group_member"
	organization_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha/organization"
	organization_invitation_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha/organization_invitation"
	permission_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha/permission"
	project_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha/project"
	project_invitation_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha/project_invitation"
	role_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha/role"
	role_binding_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha/role_binding"
	service_account_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha/service_account"
	service_account_key_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha/service_account_key"
	user_client "github.com/cloudwan/edgelq-sdk/iam/client/v1alpha/user"
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/condition"
	group "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/group"
	group_member "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/group_member"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	organization_invitation "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization_invitation"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/permission"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	project_invitation "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project_invitation"
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/role"
	role_binding "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/role_binding"
	service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/service_account"
	service_account_key "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/service_account_key"
	user "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/user"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = gotenclient.MethodDescriptor(nil)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &condition.Condition{}
	_ = &condition_client.GetConditionRequest{}
	_ = &group.Group{}
	_ = &group_member.GroupMember{}
	_ = &group_member_client.GetGroupMemberRequest{}
	_ = &group_client.GetGroupRequest{}
	_ = &organization.Organization{}
	_ = &organization_invitation.OrganizationInvitation{}
	_ = &organization_invitation_client.GetOrganizationInvitationRequest{}
	_ = &organization_client.GetOrganizationRequest{}
	_ = &permission.Permission{}
	_ = &permission_client.GetPermissionRequest{}
	_ = &project.Project{}
	_ = &project_invitation.ProjectInvitation{}
	_ = &project_invitation_client.GetProjectInvitationRequest{}
	_ = &project_client.GetProjectRequest{}
	_ = &role.Role{}
	_ = &role_binding.RoleBinding{}
	_ = &role_binding_client.GetRoleBindingRequest{}
	_ = &role_client.GetRoleRequest{}
	_ = &service_account.ServiceAccount{}
	_ = &service_account_key.ServiceAccountKey{}
	_ = &service_account_key_client.GetServiceAccountKeyRequest{}
	_ = &service_account_client.GetServiceAccountRequest{}
	_ = &user.User{}
	_ = &user_client.GetUserRequest{}
)

var (
	descriptorInitialized bool
	iamDescriptor         *IAMDescriptor
)

type IAMDescriptor struct{}

func (d *IAMDescriptor) GetServiceName() string {
	return "iam"
}

func (d *IAMDescriptor) GetServiceDomain() string {
	return "iam.edgelq.com"
}

func (d *IAMDescriptor) GetVersion() string {
	return "v1alpha"
}

func (d *IAMDescriptor) GetNextVersion() string {

	return "v1alpha2"
}

func (d *IAMDescriptor) AllResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		condition.GetDescriptor(),
		group.GetDescriptor(),
		group_member.GetDescriptor(),
		organization.GetDescriptor(),
		organization_invitation.GetDescriptor(),
		permission.GetDescriptor(),
		project.GetDescriptor(),
		project_invitation.GetDescriptor(),
		role.GetDescriptor(),
		role_binding.GetDescriptor(),
		service_account.GetDescriptor(),
		service_account_key.GetDescriptor(),
		user.GetDescriptor(),
	}
}

func (d *IAMDescriptor) AllApiDescriptors() []gotenclient.ApiDescriptor {
	return []gotenclient.ApiDescriptor{
		authorization_client.GetAuthorizationServiceDescriptor(),
		condition_client.GetConditionServiceDescriptor(),
		group_member_client.GetGroupMemberServiceDescriptor(),
		group_client.GetGroupServiceDescriptor(),
		organization_invitation_client.GetOrganizationInvitationServiceDescriptor(),
		organization_client.GetOrganizationServiceDescriptor(),
		permission_client.GetPermissionServiceDescriptor(),
		project_invitation_client.GetProjectInvitationServiceDescriptor(),
		project_client.GetProjectServiceDescriptor(),
		role_binding_client.GetRoleBindingServiceDescriptor(),
		role_client.GetRoleServiceDescriptor(),
		service_account_key_client.GetServiceAccountKeyServiceDescriptor(),
		service_account_client.GetServiceAccountServiceDescriptor(),
		user_client.GetUserServiceDescriptor(),
	}
}

func GetIAMDescriptor() *IAMDescriptor {
	return iamDescriptor
}

func initDescriptor() {
	iamDescriptor = &IAMDescriptor{}
	gotenclient.GetRegistry().RegisterSvcDescriptor(iamDescriptor)
}

func init() {
	if !descriptorInitialized {
		initDescriptor()
		descriptorInitialized = true
	}
}
