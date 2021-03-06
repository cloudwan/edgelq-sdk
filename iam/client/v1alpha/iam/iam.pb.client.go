// Code generated by protoc-gen-goten-client
// Service: IAM
// DO NOT EDIT!!!

package iam_client

import (
	"google.golang.org/grpc"
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
	_ = grpc.ClientConn{}
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

type IAMClient interface {
	authorization_client.AuthorizationServiceClient
	condition_client.ConditionServiceClient
	group_member_client.GroupMemberServiceClient
	group_client.GroupServiceClient
	organization_invitation_client.OrganizationInvitationServiceClient
	organization_client.OrganizationServiceClient
	permission_client.PermissionServiceClient
	project_invitation_client.ProjectInvitationServiceClient
	project_client.ProjectServiceClient
	role_binding_client.RoleBindingServiceClient
	role_client.RoleServiceClient
	service_account_key_client.ServiceAccountKeyServiceClient
	service_account_client.ServiceAccountServiceClient
	user_client.UserServiceClient
}

type iamClient struct {
	authorization_client.AuthorizationServiceClient
	condition_client.ConditionServiceClient
	group_member_client.GroupMemberServiceClient
	group_client.GroupServiceClient
	organization_invitation_client.OrganizationInvitationServiceClient
	organization_client.OrganizationServiceClient
	permission_client.PermissionServiceClient
	project_invitation_client.ProjectInvitationServiceClient
	project_client.ProjectServiceClient
	role_binding_client.RoleBindingServiceClient
	role_client.RoleServiceClient
	service_account_key_client.ServiceAccountKeyServiceClient
	service_account_client.ServiceAccountServiceClient
	user_client.UserServiceClient
}

func NewIAMClient(cc grpc.ClientConnInterface) IAMClient {
	return &iamClient{
		AuthorizationServiceClient:          authorization_client.NewAuthorizationServiceClient(cc),
		ConditionServiceClient:              condition_client.NewConditionServiceClient(cc),
		GroupMemberServiceClient:            group_member_client.NewGroupMemberServiceClient(cc),
		GroupServiceClient:                  group_client.NewGroupServiceClient(cc),
		OrganizationInvitationServiceClient: organization_invitation_client.NewOrganizationInvitationServiceClient(cc),
		OrganizationServiceClient:           organization_client.NewOrganizationServiceClient(cc),
		PermissionServiceClient:             permission_client.NewPermissionServiceClient(cc),
		ProjectInvitationServiceClient:      project_invitation_client.NewProjectInvitationServiceClient(cc),
		ProjectServiceClient:                project_client.NewProjectServiceClient(cc),
		RoleBindingServiceClient:            role_binding_client.NewRoleBindingServiceClient(cc),
		RoleServiceClient:                   role_client.NewRoleServiceClient(cc),
		ServiceAccountKeyServiceClient:      service_account_key_client.NewServiceAccountKeyServiceClient(cc),
		ServiceAccountServiceClient:         service_account_client.NewServiceAccountServiceClient(cc),
		UserServiceClient:                   user_client.NewUserServiceClient(cc),
	}
}
