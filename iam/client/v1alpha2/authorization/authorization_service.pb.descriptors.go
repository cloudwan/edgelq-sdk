// Code generated by protoc-gen-goten-client
// API: AuthorizationService
// DO NOT EDIT!!!

package authorization_client

import (
	"google.golang.org/protobuf/proto"

	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/permission"
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/role"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = new(proto.Message)
	_ = new(gotenclient.MethodDescriptor)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &permission.Permission{}
	_ = &role.Role{}
)

var (
	descriptorsInitialized         bool
	authorizationServiceDescriptor *AuthorizationServiceDescriptor
	checkPermissionsDescriptor     *CheckPermissionsDescriptor
	checkMyPermissionsDescriptor   *CheckMyPermissionsDescriptor
	checkMyRolesDescriptor         *CheckMyRolesDescriptor
)

type CheckPermissionsDescriptor struct{}

type CheckPermissionsDescriptorClientMsgHandle struct{}

type CheckPermissionsDescriptorServerMsgHandle struct{}

func (d *CheckPermissionsDescriptor) NewEmptyClientMsg() proto.Message {
	return &CheckPermissionsRequest{}
}

func (d *CheckPermissionsDescriptor) NewEmptyServerMsg() proto.Message {
	return &CheckPermissionsResponse{}
}

func (d *CheckPermissionsDescriptor) IsUnary() bool {
	return true
}

func (d *CheckPermissionsDescriptor) IsClientStream() bool {
	return false
}

func (d *CheckPermissionsDescriptor) IsServerStream() bool {
	return false
}

func (d *CheckPermissionsDescriptor) IsCollection() bool {
	return true
}

func (d *CheckPermissionsDescriptor) IsPlural() bool {
	return true
}

func (d *CheckPermissionsDescriptor) HasResource() bool {
	return true
}

func (d *CheckPermissionsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *CheckPermissionsDescriptor) GetVerb() string {
	return "check"
}

func (d *CheckPermissionsDescriptor) GetMethodName() string {
	return "CheckPermissions"
}

func (d *CheckPermissionsDescriptor) GetFullMethodName() string {
	return "/ntt.iam.v1alpha2.AuthorizationService/CheckPermissions"
}

func (d *CheckPermissionsDescriptor) GetProtoPkgName() string {
	return "ntt.iam.v1alpha2"
}

func (d *CheckPermissionsDescriptor) GetApiName() string {
	return "AuthorizationService"
}

func (d *CheckPermissionsDescriptor) GetServiceDomain() string {
	return "iam.edgelq.com"
}

func (d *CheckPermissionsDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *CheckPermissionsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return authorizationServiceDescriptor
}

func (d *CheckPermissionsDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return permission.GetDescriptor()
}

func (d *CheckPermissionsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CheckPermissionsDescriptorClientMsgHandle{}
}

func (d *CheckPermissionsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CheckPermissionsDescriptorServerMsgHandle{}
}

func (h *CheckPermissionsDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CheckPermissionsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*CheckPermissionsRequest) *permission.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *CheckPermissionsDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*CheckPermissionsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*CheckPermissionsRequest) []*permission.Name
	})
	if ok {
		return permission.PermissionNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *CheckPermissionsDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *CheckPermissionsDescriptorClientMsgHandle) ExtractResourceBody(msg proto.Message) gotenresource.Resource {
	typedMsg := msg.(*CheckPermissionsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBody(*CheckPermissionsRequest) *permission.Permission
	})
	if ok {
		return override.OverrideExtractResourceBody(typedMsg)
	}
	return nil
}

func (h *CheckPermissionsDescriptorClientMsgHandle) ExtractResourceBodies(msg proto.Message) gotenresource.ResourceList {
	typedMsg := msg.(*CheckPermissionsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBodies(*CheckPermissionsRequest) []*permission.Permission
	})
	if ok {
		return permission.PermissionList(override.OverrideExtractResourceBodies(typedMsg))
	}
	return nil
}

func (h *CheckPermissionsDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CheckPermissionsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*CheckPermissionsResponse) *permission.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *CheckPermissionsDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*CheckPermissionsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*CheckPermissionsResponse) []*permission.Name
	})
	if ok {
		return permission.PermissionNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *CheckPermissionsDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *CheckPermissionsDescriptorServerMsgHandle) ExtractResourceBody(msg proto.Message) gotenresource.Resource {
	typedMsg := msg.(*CheckPermissionsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBody(*CheckPermissionsResponse) *permission.Permission
	})
	if ok {
		return override.OverrideExtractResourceBody(typedMsg)
	}
	return nil
}

func (h *CheckPermissionsDescriptorServerMsgHandle) ExtractResourceBodies(msg proto.Message) gotenresource.ResourceList {
	typedMsg := msg.(*CheckPermissionsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBodies(*CheckPermissionsResponse) []*permission.Permission
	})
	if ok {
		return permission.PermissionList(override.OverrideExtractResourceBodies(typedMsg))
	}
	return nil
}

func GetCheckPermissionsDescriptor() *CheckPermissionsDescriptor {
	return checkPermissionsDescriptor
}

type CheckMyPermissionsDescriptor struct{}

type CheckMyPermissionsDescriptorClientMsgHandle struct{}

type CheckMyPermissionsDescriptorServerMsgHandle struct{}

func (d *CheckMyPermissionsDescriptor) NewEmptyClientMsg() proto.Message {
	return &CheckMyPermissionsRequest{}
}

func (d *CheckMyPermissionsDescriptor) NewEmptyServerMsg() proto.Message {
	return &CheckMyPermissionsResponse{}
}

func (d *CheckMyPermissionsDescriptor) IsUnary() bool {
	return true
}

func (d *CheckMyPermissionsDescriptor) IsClientStream() bool {
	return false
}

func (d *CheckMyPermissionsDescriptor) IsServerStream() bool {
	return false
}

func (d *CheckMyPermissionsDescriptor) IsCollection() bool {
	return true
}

func (d *CheckMyPermissionsDescriptor) IsPlural() bool {
	return true
}

func (d *CheckMyPermissionsDescriptor) HasResource() bool {
	return true
}

func (d *CheckMyPermissionsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *CheckMyPermissionsDescriptor) GetVerb() string {
	return "checkMy"
}

func (d *CheckMyPermissionsDescriptor) GetMethodName() string {
	return "CheckMyPermissions"
}

func (d *CheckMyPermissionsDescriptor) GetFullMethodName() string {
	return "/ntt.iam.v1alpha2.AuthorizationService/CheckMyPermissions"
}

func (d *CheckMyPermissionsDescriptor) GetProtoPkgName() string {
	return "ntt.iam.v1alpha2"
}

func (d *CheckMyPermissionsDescriptor) GetApiName() string {
	return "AuthorizationService"
}

func (d *CheckMyPermissionsDescriptor) GetServiceDomain() string {
	return "iam.edgelq.com"
}

func (d *CheckMyPermissionsDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *CheckMyPermissionsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return authorizationServiceDescriptor
}

func (d *CheckMyPermissionsDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return permission.GetDescriptor()
}

func (d *CheckMyPermissionsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CheckMyPermissionsDescriptorClientMsgHandle{}
}

func (d *CheckMyPermissionsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CheckMyPermissionsDescriptorServerMsgHandle{}
}

func (h *CheckMyPermissionsDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CheckMyPermissionsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*CheckMyPermissionsRequest) *permission.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *CheckMyPermissionsDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*CheckMyPermissionsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*CheckMyPermissionsRequest) []*permission.Name
	})
	if ok {
		return permission.PermissionNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *CheckMyPermissionsDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *CheckMyPermissionsDescriptorClientMsgHandle) ExtractResourceBody(msg proto.Message) gotenresource.Resource {
	typedMsg := msg.(*CheckMyPermissionsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBody(*CheckMyPermissionsRequest) *permission.Permission
	})
	if ok {
		return override.OverrideExtractResourceBody(typedMsg)
	}
	return nil
}

func (h *CheckMyPermissionsDescriptorClientMsgHandle) ExtractResourceBodies(msg proto.Message) gotenresource.ResourceList {
	typedMsg := msg.(*CheckMyPermissionsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBodies(*CheckMyPermissionsRequest) []*permission.Permission
	})
	if ok {
		return permission.PermissionList(override.OverrideExtractResourceBodies(typedMsg))
	}
	return nil
}

func (h *CheckMyPermissionsDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CheckMyPermissionsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*CheckMyPermissionsResponse) *permission.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *CheckMyPermissionsDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*CheckMyPermissionsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*CheckMyPermissionsResponse) []*permission.Name
	})
	if ok {
		return permission.PermissionNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *CheckMyPermissionsDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *CheckMyPermissionsDescriptorServerMsgHandle) ExtractResourceBody(msg proto.Message) gotenresource.Resource {
	typedMsg := msg.(*CheckMyPermissionsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBody(*CheckMyPermissionsResponse) *permission.Permission
	})
	if ok {
		return override.OverrideExtractResourceBody(typedMsg)
	}
	return nil
}

func (h *CheckMyPermissionsDescriptorServerMsgHandle) ExtractResourceBodies(msg proto.Message) gotenresource.ResourceList {
	typedMsg := msg.(*CheckMyPermissionsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBodies(*CheckMyPermissionsResponse) []*permission.Permission
	})
	if ok {
		return permission.PermissionList(override.OverrideExtractResourceBodies(typedMsg))
	}
	return nil
}

func GetCheckMyPermissionsDescriptor() *CheckMyPermissionsDescriptor {
	return checkMyPermissionsDescriptor
}

type CheckMyRolesDescriptor struct{}

type CheckMyRolesDescriptorClientMsgHandle struct{}

type CheckMyRolesDescriptorServerMsgHandle struct{}

func (d *CheckMyRolesDescriptor) NewEmptyClientMsg() proto.Message {
	return &CheckMyRolesRequest{}
}

func (d *CheckMyRolesDescriptor) NewEmptyServerMsg() proto.Message {
	return &CheckMyRolesResponse{}
}

func (d *CheckMyRolesDescriptor) IsUnary() bool {
	return true
}

func (d *CheckMyRolesDescriptor) IsClientStream() bool {
	return false
}

func (d *CheckMyRolesDescriptor) IsServerStream() bool {
	return false
}

func (d *CheckMyRolesDescriptor) IsCollection() bool {
	return true
}

func (d *CheckMyRolesDescriptor) IsPlural() bool {
	return true
}

func (d *CheckMyRolesDescriptor) HasResource() bool {
	return true
}

func (d *CheckMyRolesDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *CheckMyRolesDescriptor) GetVerb() string {
	return "checkMy"
}

func (d *CheckMyRolesDescriptor) GetMethodName() string {
	return "CheckMyRoles"
}

func (d *CheckMyRolesDescriptor) GetFullMethodName() string {
	return "/ntt.iam.v1alpha2.AuthorizationService/CheckMyRoles"
}

func (d *CheckMyRolesDescriptor) GetProtoPkgName() string {
	return "ntt.iam.v1alpha2"
}

func (d *CheckMyRolesDescriptor) GetApiName() string {
	return "AuthorizationService"
}

func (d *CheckMyRolesDescriptor) GetServiceDomain() string {
	return "iam.edgelq.com"
}

func (d *CheckMyRolesDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *CheckMyRolesDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return authorizationServiceDescriptor
}

func (d *CheckMyRolesDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return role.GetDescriptor()
}

func (d *CheckMyRolesDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CheckMyRolesDescriptorClientMsgHandle{}
}

func (d *CheckMyRolesDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CheckMyRolesDescriptorServerMsgHandle{}
}

func (h *CheckMyRolesDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CheckMyRolesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*CheckMyRolesRequest) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *CheckMyRolesDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*CheckMyRolesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*CheckMyRolesRequest) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *CheckMyRolesDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *CheckMyRolesDescriptorClientMsgHandle) ExtractResourceBody(msg proto.Message) gotenresource.Resource {
	typedMsg := msg.(*CheckMyRolesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBody(*CheckMyRolesRequest) *role.Role
	})
	if ok {
		return override.OverrideExtractResourceBody(typedMsg)
	}
	return nil
}

func (h *CheckMyRolesDescriptorClientMsgHandle) ExtractResourceBodies(msg proto.Message) gotenresource.ResourceList {
	typedMsg := msg.(*CheckMyRolesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBodies(*CheckMyRolesRequest) []*role.Role
	})
	if ok {
		return role.RoleList(override.OverrideExtractResourceBodies(typedMsg))
	}
	return nil
}

func (h *CheckMyRolesDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CheckMyRolesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*CheckMyRolesResponse) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *CheckMyRolesDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*CheckMyRolesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*CheckMyRolesResponse) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *CheckMyRolesDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *CheckMyRolesDescriptorServerMsgHandle) ExtractResourceBody(msg proto.Message) gotenresource.Resource {
	typedMsg := msg.(*CheckMyRolesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBody(*CheckMyRolesResponse) *role.Role
	})
	if ok {
		return override.OverrideExtractResourceBody(typedMsg)
	}
	return nil
}

func (h *CheckMyRolesDescriptorServerMsgHandle) ExtractResourceBodies(msg proto.Message) gotenresource.ResourceList {
	typedMsg := msg.(*CheckMyRolesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBodies(*CheckMyRolesResponse) []*role.Role
	})
	if ok {
		return role.RoleList(override.OverrideExtractResourceBodies(typedMsg))
	}
	return nil
}

func GetCheckMyRolesDescriptor() *CheckMyRolesDescriptor {
	return checkMyRolesDescriptor
}

type AuthorizationServiceDescriptor struct{}

func (d *AuthorizationServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		checkPermissionsDescriptor,
		checkMyPermissionsDescriptor,
		checkMyRolesDescriptor,
	}
}

func (d *AuthorizationServiceDescriptor) GetFullAPIName() string {
	return "/ntt.iam.v1alpha2.AuthorizationService"
}

func (d *AuthorizationServiceDescriptor) GetProtoPkgName() string {
	return "ntt.iam.v1alpha2"
}

func (d *AuthorizationServiceDescriptor) GetApiName() string {
	return "AuthorizationService"
}

func (d *AuthorizationServiceDescriptor) GetServiceDomain() string {
	return "iam.edgelq.com"
}

func (d *AuthorizationServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func GetAuthorizationServiceDescriptor() *AuthorizationServiceDescriptor {
	return authorizationServiceDescriptor
}

func initDescriptors() {
	authorizationServiceDescriptor = &AuthorizationServiceDescriptor{}
	checkPermissionsDescriptor = &CheckPermissionsDescriptor{}
	checkMyPermissionsDescriptor = &CheckMyPermissionsDescriptor{}
	checkMyRolesDescriptor = &CheckMyRolesDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(authorizationServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(checkPermissionsDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(checkMyPermissionsDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(checkMyRolesDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}
