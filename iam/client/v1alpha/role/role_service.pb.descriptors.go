// Code generated by protoc-gen-goten-client
// API: RoleService
// DO NOT EDIT!!!

package role_client

import (
	"google.golang.org/protobuf/proto"

	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/role"
	empty "github.com/golang/protobuf/ptypes/empty"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = new(proto.Message)
	_ = new(gotenclient.MethodDescriptor)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &role.Role{}
	_ = &empty.Empty{}
)

var (
	descriptorsInitialized  bool
	roleServiceDescriptor   *RoleServiceDescriptor
	getRoleDescriptor       *GetRoleDescriptor
	batchGetRolesDescriptor *BatchGetRolesDescriptor
	listRolesDescriptor     *ListRolesDescriptor
	watchRoleDescriptor     *WatchRoleDescriptor
	watchRolesDescriptor    *WatchRolesDescriptor
	createRoleDescriptor    *CreateRoleDescriptor
	updateRoleDescriptor    *UpdateRoleDescriptor
	deleteRoleDescriptor    *DeleteRoleDescriptor
)

type GetRoleDescriptor struct{}

type GetRoleDescriptorClientMsgHandle struct{}

type GetRoleDescriptorServerMsgHandle struct{}

func (d *GetRoleDescriptor) NewEmptyClientMsg() proto.Message {
	return &GetRoleRequest{}
}

func (d *GetRoleDescriptor) NewEmptyServerMsg() proto.Message {
	return &role.Role{}
}

func (d *GetRoleDescriptor) IsUnary() bool {
	return true
}

func (d *GetRoleDescriptor) IsClientStream() bool {
	return false
}

func (d *GetRoleDescriptor) IsServerStream() bool {
	return false
}

func (d *GetRoleDescriptor) IsCollection() bool {
	return false
}

func (d *GetRoleDescriptor) IsPlural() bool {
	return false
}

func (d *GetRoleDescriptor) HasResource() bool {
	return true
}

func (d *GetRoleDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *GetRoleDescriptor) GetVerb() string {
	return "get"
}

func (d *GetRoleDescriptor) GetMethodName() string {
	return "GetRole"
}

func (d *GetRoleDescriptor) GetFullMethodName() string {
	return "/ntt.iam.v1alpha.RoleService/GetRole"
}

func (d *GetRoleDescriptor) GetProtoPkgName() string {
	return "ntt.iam.v1alpha"
}

func (d *GetRoleDescriptor) GetApiName() string {
	return "RoleService"
}

func (d *GetRoleDescriptor) GetServiceDomain() string {
	return "iam.edgelq.com"
}

func (d *GetRoleDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *GetRoleDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return roleServiceDescriptor
}

func (d *GetRoleDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return role.GetDescriptor()
}

func (d *GetRoleDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetRoleDescriptorClientMsgHandle{}
}

func (d *GetRoleDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetRoleDescriptorServerMsgHandle{}
}

func (h *GetRoleDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*GetRoleRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*GetRoleRequest) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if ref := typedMsg.GetName(); ref != nil {
			return &ref.Name
		}
	}
	return (*role.Name)(nil)
}

func (h *GetRoleDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*GetRoleRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*GetRoleRequest) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *GetRoleDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *GetRoleDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*role.Role)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*role.Role) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if name := typedMsg.GetName(); name != nil {
			return name
		}
	}
	return (*role.Name)(nil)
}

func (h *GetRoleDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*role.Role)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*role.Role) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *GetRoleDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetGetRoleDescriptor() *GetRoleDescriptor {
	return getRoleDescriptor
}

type BatchGetRolesDescriptor struct{}

type BatchGetRolesDescriptorClientMsgHandle struct{}

type BatchGetRolesDescriptorServerMsgHandle struct{}

func (d *BatchGetRolesDescriptor) NewEmptyClientMsg() proto.Message {
	return &BatchGetRolesRequest{}
}

func (d *BatchGetRolesDescriptor) NewEmptyServerMsg() proto.Message {
	return &BatchGetRolesResponse{}
}

func (d *BatchGetRolesDescriptor) IsUnary() bool {
	return true
}

func (d *BatchGetRolesDescriptor) IsClientStream() bool {
	return false
}

func (d *BatchGetRolesDescriptor) IsServerStream() bool {
	return false
}

func (d *BatchGetRolesDescriptor) IsCollection() bool {
	return false
}

func (d *BatchGetRolesDescriptor) IsPlural() bool {
	return true
}

func (d *BatchGetRolesDescriptor) HasResource() bool {
	return true
}

func (d *BatchGetRolesDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *BatchGetRolesDescriptor) GetVerb() string {
	return "batchGet"
}

func (d *BatchGetRolesDescriptor) GetMethodName() string {
	return "BatchGetRoles"
}

func (d *BatchGetRolesDescriptor) GetFullMethodName() string {
	return "/ntt.iam.v1alpha.RoleService/BatchGetRoles"
}

func (d *BatchGetRolesDescriptor) GetProtoPkgName() string {
	return "ntt.iam.v1alpha"
}

func (d *BatchGetRolesDescriptor) GetApiName() string {
	return "RoleService"
}

func (d *BatchGetRolesDescriptor) GetServiceDomain() string {
	return "iam.edgelq.com"
}

func (d *BatchGetRolesDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *BatchGetRolesDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return roleServiceDescriptor
}

func (d *BatchGetRolesDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return role.GetDescriptor()
}

func (d *BatchGetRolesDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &BatchGetRolesDescriptorClientMsgHandle{}
}

func (d *BatchGetRolesDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &BatchGetRolesDescriptorServerMsgHandle{}
}

func (h *BatchGetRolesDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*BatchGetRolesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*BatchGetRolesRequest) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *BatchGetRolesDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*BatchGetRolesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*BatchGetRolesRequest) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	{
		if refs := typedMsg.GetNames(); len(refs) > 0 {
			list := make(role.RoleNameList, 0, len(refs))
			for _, ref := range refs {
				list = append(list, &ref.Name)
			}
			return list
		}
	}
	return (role.RoleNameList)(nil)
}

func (h *BatchGetRolesDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *BatchGetRolesDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*BatchGetRolesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*BatchGetRolesResponse) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *BatchGetRolesDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*BatchGetRolesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*BatchGetRolesResponse) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	{
		if resources := typedMsg.GetRoles(); len(resources) > 0 {
			list := make(role.RoleNameList, 0, len(resources))
			for _, res := range resources {
				list = append(list, res.GetName())
			}
			return list
		}
	}
	return (role.RoleNameList)(nil)
}

func (h *BatchGetRolesDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetBatchGetRolesDescriptor() *BatchGetRolesDescriptor {
	return batchGetRolesDescriptor
}

type ListRolesDescriptor struct{}

type ListRolesDescriptorClientMsgHandle struct{}

type ListRolesDescriptorServerMsgHandle struct{}

func (d *ListRolesDescriptor) NewEmptyClientMsg() proto.Message {
	return &ListRolesRequest{}
}

func (d *ListRolesDescriptor) NewEmptyServerMsg() proto.Message {
	return &ListRolesResponse{}
}

func (d *ListRolesDescriptor) IsUnary() bool {
	return true
}

func (d *ListRolesDescriptor) IsClientStream() bool {
	return false
}

func (d *ListRolesDescriptor) IsServerStream() bool {
	return false
}

func (d *ListRolesDescriptor) IsCollection() bool {
	return true
}

func (d *ListRolesDescriptor) IsPlural() bool {
	return true
}

func (d *ListRolesDescriptor) HasResource() bool {
	return true
}

func (d *ListRolesDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *ListRolesDescriptor) GetVerb() string {
	return "list"
}

func (d *ListRolesDescriptor) GetMethodName() string {
	return "ListRoles"
}

func (d *ListRolesDescriptor) GetFullMethodName() string {
	return "/ntt.iam.v1alpha.RoleService/ListRoles"
}

func (d *ListRolesDescriptor) GetProtoPkgName() string {
	return "ntt.iam.v1alpha"
}

func (d *ListRolesDescriptor) GetApiName() string {
	return "RoleService"
}

func (d *ListRolesDescriptor) GetServiceDomain() string {
	return "iam.edgelq.com"
}

func (d *ListRolesDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *ListRolesDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return roleServiceDescriptor
}

func (d *ListRolesDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return role.GetDescriptor()
}

func (d *ListRolesDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ListRolesDescriptorClientMsgHandle{}
}

func (d *ListRolesDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ListRolesDescriptorServerMsgHandle{}
}

func (h *ListRolesDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListRolesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ListRolesRequest) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *ListRolesDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ListRolesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ListRolesRequest) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *ListRolesDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *ListRolesDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListRolesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ListRolesResponse) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *ListRolesDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ListRolesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ListRolesResponse) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	{
		if resources := typedMsg.GetRoles(); len(resources) > 0 {
			list := make(role.RoleNameList, 0, len(resources))
			for _, res := range resources {
				list = append(list, res.GetName())
			}
			return list
		}
	}
	return (role.RoleNameList)(nil)
}

func (h *ListRolesDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetListRolesDescriptor() *ListRolesDescriptor {
	return listRolesDescriptor
}

type WatchRoleDescriptor struct{}

type WatchRoleDescriptorClientMsgHandle struct{}

type WatchRoleDescriptorServerMsgHandle struct{}

func (d *WatchRoleDescriptor) NewEmptyClientMsg() proto.Message {
	return &WatchRoleRequest{}
}

func (d *WatchRoleDescriptor) NewEmptyServerMsg() proto.Message {
	return &WatchRoleResponse{}
}

func (d *WatchRoleDescriptor) IsUnary() bool {
	return false
}

func (d *WatchRoleDescriptor) IsClientStream() bool {
	return false
}

func (d *WatchRoleDescriptor) IsServerStream() bool {
	return true
}

func (d *WatchRoleDescriptor) IsCollection() bool {
	return false
}

func (d *WatchRoleDescriptor) IsPlural() bool {
	return false
}

func (d *WatchRoleDescriptor) HasResource() bool {
	return true
}

func (d *WatchRoleDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *WatchRoleDescriptor) GetVerb() string {
	return "watch"
}

func (d *WatchRoleDescriptor) GetMethodName() string {
	return "WatchRole"
}

func (d *WatchRoleDescriptor) GetFullMethodName() string {
	return "/ntt.iam.v1alpha.RoleService/WatchRole"
}

func (d *WatchRoleDescriptor) GetProtoPkgName() string {
	return "ntt.iam.v1alpha"
}

func (d *WatchRoleDescriptor) GetApiName() string {
	return "RoleService"
}

func (d *WatchRoleDescriptor) GetServiceDomain() string {
	return "iam.edgelq.com"
}

func (d *WatchRoleDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *WatchRoleDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return roleServiceDescriptor
}

func (d *WatchRoleDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return role.GetDescriptor()
}

func (d *WatchRoleDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &WatchRoleDescriptorClientMsgHandle{}
}

func (d *WatchRoleDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &WatchRoleDescriptorServerMsgHandle{}
}

func (h *WatchRoleDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*WatchRoleRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*WatchRoleRequest) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if ref := typedMsg.GetName(); ref != nil {
			return &ref.Name
		}
	}
	return (*role.Name)(nil)
}

func (h *WatchRoleDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*WatchRoleRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*WatchRoleRequest) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *WatchRoleDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *WatchRoleDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*WatchRoleResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*WatchRoleResponse) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if resChange := typedMsg.GetChange(); resChange != nil {
			switch tResChange := resChange.ChangeType.(type) {
			case *role.RoleChange_Added_:
				return tResChange.Added.GetRole().GetName()
			case *role.RoleChange_Modified_:
				return tResChange.Modified.GetName()
			case *role.RoleChange_Removed_:
				return tResChange.Removed.GetName()
			case *role.RoleChange_Current_:
				return tResChange.Current.GetRole().GetName()
			}
		}
	}
	return (*role.Name)(nil)
}

func (h *WatchRoleDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*WatchRoleResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*WatchRoleResponse) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *WatchRoleDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetWatchRoleDescriptor() *WatchRoleDescriptor {
	return watchRoleDescriptor
}

type WatchRolesDescriptor struct{}

type WatchRolesDescriptorClientMsgHandle struct{}

type WatchRolesDescriptorServerMsgHandle struct{}

func (d *WatchRolesDescriptor) NewEmptyClientMsg() proto.Message {
	return &WatchRolesRequest{}
}

func (d *WatchRolesDescriptor) NewEmptyServerMsg() proto.Message {
	return &WatchRolesResponse{}
}

func (d *WatchRolesDescriptor) IsUnary() bool {
	return false
}

func (d *WatchRolesDescriptor) IsClientStream() bool {
	return false
}

func (d *WatchRolesDescriptor) IsServerStream() bool {
	return true
}

func (d *WatchRolesDescriptor) IsCollection() bool {
	return true
}

func (d *WatchRolesDescriptor) IsPlural() bool {
	return true
}

func (d *WatchRolesDescriptor) HasResource() bool {
	return true
}

func (d *WatchRolesDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *WatchRolesDescriptor) GetVerb() string {
	return "watch"
}

func (d *WatchRolesDescriptor) GetMethodName() string {
	return "WatchRoles"
}

func (d *WatchRolesDescriptor) GetFullMethodName() string {
	return "/ntt.iam.v1alpha.RoleService/WatchRoles"
}

func (d *WatchRolesDescriptor) GetProtoPkgName() string {
	return "ntt.iam.v1alpha"
}

func (d *WatchRolesDescriptor) GetApiName() string {
	return "RoleService"
}

func (d *WatchRolesDescriptor) GetServiceDomain() string {
	return "iam.edgelq.com"
}

func (d *WatchRolesDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *WatchRolesDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return roleServiceDescriptor
}

func (d *WatchRolesDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return role.GetDescriptor()
}

func (d *WatchRolesDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &WatchRolesDescriptorClientMsgHandle{}
}

func (d *WatchRolesDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &WatchRolesDescriptorServerMsgHandle{}
}

func (h *WatchRolesDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*WatchRolesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*WatchRolesRequest) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *WatchRolesDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*WatchRolesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*WatchRolesRequest) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *WatchRolesDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *WatchRolesDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*WatchRolesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*WatchRolesResponse) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *WatchRolesDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*WatchRolesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*WatchRolesResponse) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	{
		if resChanges := typedMsg.GetRoleChanges(); len(resChanges) > 0 {
			list := make(role.RoleNameList, 0, len(resChanges))
			for _, resChange := range resChanges {
				switch tResChange := resChange.ChangeType.(type) {
				case *role.RoleChange_Added_:
					list = append(list, tResChange.Added.GetRole().GetName())
				case *role.RoleChange_Modified_:
					list = append(list, tResChange.Modified.GetName())
				case *role.RoleChange_Removed_:
					list = append(list, tResChange.Removed.GetName())
				case *role.RoleChange_Current_:
					list = append(list, tResChange.Current.GetRole().GetName())
				}
			}
			return list
		}
	}
	return (role.RoleNameList)(nil)
}

func (h *WatchRolesDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetWatchRolesDescriptor() *WatchRolesDescriptor {
	return watchRolesDescriptor
}

type CreateRoleDescriptor struct{}

type CreateRoleDescriptorClientMsgHandle struct{}

type CreateRoleDescriptorServerMsgHandle struct{}

func (d *CreateRoleDescriptor) NewEmptyClientMsg() proto.Message {
	return &CreateRoleRequest{}
}

func (d *CreateRoleDescriptor) NewEmptyServerMsg() proto.Message {
	return &role.Role{}
}

func (d *CreateRoleDescriptor) IsUnary() bool {
	return true
}

func (d *CreateRoleDescriptor) IsClientStream() bool {
	return false
}

func (d *CreateRoleDescriptor) IsServerStream() bool {
	return false
}

func (d *CreateRoleDescriptor) IsCollection() bool {
	return true
}

func (d *CreateRoleDescriptor) IsPlural() bool {
	return false
}

func (d *CreateRoleDescriptor) HasResource() bool {
	return true
}

func (d *CreateRoleDescriptor) RequestHasResourceBody() bool {
	return true
}

func (d *CreateRoleDescriptor) GetVerb() string {
	return "create"
}

func (d *CreateRoleDescriptor) GetMethodName() string {
	return "CreateRole"
}

func (d *CreateRoleDescriptor) GetFullMethodName() string {
	return "/ntt.iam.v1alpha.RoleService/CreateRole"
}

func (d *CreateRoleDescriptor) GetProtoPkgName() string {
	return "ntt.iam.v1alpha"
}

func (d *CreateRoleDescriptor) GetApiName() string {
	return "RoleService"
}

func (d *CreateRoleDescriptor) GetServiceDomain() string {
	return "iam.edgelq.com"
}

func (d *CreateRoleDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *CreateRoleDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return roleServiceDescriptor
}

func (d *CreateRoleDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return role.GetDescriptor()
}

func (d *CreateRoleDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CreateRoleDescriptorClientMsgHandle{}
}

func (d *CreateRoleDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CreateRoleDescriptorServerMsgHandle{}
}

func (h *CreateRoleDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CreateRoleRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*CreateRoleRequest) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		res := typedMsg.GetRole()
		if name := res.GetName(); name != nil {
			return name
		}
	}
	return (*role.Name)(nil)
}

func (h *CreateRoleDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*CreateRoleRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*CreateRoleRequest) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *CreateRoleDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *CreateRoleDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*role.Role)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*role.Role) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if name := typedMsg.GetName(); name != nil {
			return name
		}
	}
	return (*role.Name)(nil)
}

func (h *CreateRoleDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*role.Role)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*role.Role) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *CreateRoleDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetCreateRoleDescriptor() *CreateRoleDescriptor {
	return createRoleDescriptor
}

type UpdateRoleDescriptor struct{}

type UpdateRoleDescriptorClientMsgHandle struct{}

type UpdateRoleDescriptorServerMsgHandle struct{}

func (d *UpdateRoleDescriptor) NewEmptyClientMsg() proto.Message {
	return &UpdateRoleRequest{}
}

func (d *UpdateRoleDescriptor) NewEmptyServerMsg() proto.Message {
	return &role.Role{}
}

func (d *UpdateRoleDescriptor) IsUnary() bool {
	return true
}

func (d *UpdateRoleDescriptor) IsClientStream() bool {
	return false
}

func (d *UpdateRoleDescriptor) IsServerStream() bool {
	return false
}

func (d *UpdateRoleDescriptor) IsCollection() bool {
	return false
}

func (d *UpdateRoleDescriptor) IsPlural() bool {
	return false
}

func (d *UpdateRoleDescriptor) HasResource() bool {
	return true
}

func (d *UpdateRoleDescriptor) RequestHasResourceBody() bool {
	return true
}

func (d *UpdateRoleDescriptor) GetVerb() string {
	return "update"
}

func (d *UpdateRoleDescriptor) GetMethodName() string {
	return "UpdateRole"
}

func (d *UpdateRoleDescriptor) GetFullMethodName() string {
	return "/ntt.iam.v1alpha.RoleService/UpdateRole"
}

func (d *UpdateRoleDescriptor) GetProtoPkgName() string {
	return "ntt.iam.v1alpha"
}

func (d *UpdateRoleDescriptor) GetApiName() string {
	return "RoleService"
}

func (d *UpdateRoleDescriptor) GetServiceDomain() string {
	return "iam.edgelq.com"
}

func (d *UpdateRoleDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *UpdateRoleDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return roleServiceDescriptor
}

func (d *UpdateRoleDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return role.GetDescriptor()
}

func (d *UpdateRoleDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &UpdateRoleDescriptorClientMsgHandle{}
}

func (d *UpdateRoleDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &UpdateRoleDescriptorServerMsgHandle{}
}

func (h *UpdateRoleDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*UpdateRoleRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*UpdateRoleRequest) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		res := typedMsg.GetRole()
		if name := res.GetName(); name != nil {
			return name
		}
	}
	return (*role.Name)(nil)
}

func (h *UpdateRoleDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*UpdateRoleRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*UpdateRoleRequest) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *UpdateRoleDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *UpdateRoleDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*role.Role)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*role.Role) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if name := typedMsg.GetName(); name != nil {
			return name
		}
	}
	return (*role.Name)(nil)
}

func (h *UpdateRoleDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*role.Role)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*role.Role) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *UpdateRoleDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetUpdateRoleDescriptor() *UpdateRoleDescriptor {
	return updateRoleDescriptor
}

type DeleteRoleDescriptor struct{}

type DeleteRoleDescriptorClientMsgHandle struct{}

type DeleteRoleDescriptorServerMsgHandle struct{}

func (d *DeleteRoleDescriptor) NewEmptyClientMsg() proto.Message {
	return &DeleteRoleRequest{}
}

func (d *DeleteRoleDescriptor) NewEmptyServerMsg() proto.Message {
	return &empty.Empty{}
}

func (d *DeleteRoleDescriptor) IsUnary() bool {
	return true
}

func (d *DeleteRoleDescriptor) IsClientStream() bool {
	return false
}

func (d *DeleteRoleDescriptor) IsServerStream() bool {
	return false
}

func (d *DeleteRoleDescriptor) IsCollection() bool {
	return false
}

func (d *DeleteRoleDescriptor) IsPlural() bool {
	return false
}

func (d *DeleteRoleDescriptor) HasResource() bool {
	return true
}

func (d *DeleteRoleDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *DeleteRoleDescriptor) GetVerb() string {
	return "delete"
}

func (d *DeleteRoleDescriptor) GetMethodName() string {
	return "DeleteRole"
}

func (d *DeleteRoleDescriptor) GetFullMethodName() string {
	return "/ntt.iam.v1alpha.RoleService/DeleteRole"
}

func (d *DeleteRoleDescriptor) GetProtoPkgName() string {
	return "ntt.iam.v1alpha"
}

func (d *DeleteRoleDescriptor) GetApiName() string {
	return "RoleService"
}

func (d *DeleteRoleDescriptor) GetServiceDomain() string {
	return "iam.edgelq.com"
}

func (d *DeleteRoleDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *DeleteRoleDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return roleServiceDescriptor
}

func (d *DeleteRoleDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return role.GetDescriptor()
}

func (d *DeleteRoleDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &DeleteRoleDescriptorClientMsgHandle{}
}

func (d *DeleteRoleDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &DeleteRoleDescriptorServerMsgHandle{}
}

func (h *DeleteRoleDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*DeleteRoleRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*DeleteRoleRequest) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if ref := typedMsg.GetName(); ref != nil {
			return &ref.Name
		}
	}
	return (*role.Name)(nil)
}

func (h *DeleteRoleDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*DeleteRoleRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*DeleteRoleRequest) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *DeleteRoleDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *DeleteRoleDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*empty.Empty)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*empty.Empty) *role.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *DeleteRoleDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*empty.Empty)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*empty.Empty) []*role.Name
	})
	if ok {
		return role.RoleNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *DeleteRoleDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetDeleteRoleDescriptor() *DeleteRoleDescriptor {
	return deleteRoleDescriptor
}

type RoleServiceDescriptor struct{}

func (d *RoleServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		getRoleDescriptor,
		batchGetRolesDescriptor,
		listRolesDescriptor,
		watchRoleDescriptor,
		watchRolesDescriptor,
		createRoleDescriptor,
		updateRoleDescriptor,
		deleteRoleDescriptor,
	}
}

func (d *RoleServiceDescriptor) GetFullAPIName() string {
	return "/ntt.iam.v1alpha.RoleService"
}

func (d *RoleServiceDescriptor) GetProtoPkgName() string {
	return "ntt.iam.v1alpha"
}

func (d *RoleServiceDescriptor) GetApiName() string {
	return "RoleService"
}

func (d *RoleServiceDescriptor) GetServiceDomain() string {
	return "iam.edgelq.com"
}

func (d *RoleServiceDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func GetRoleServiceDescriptor() *RoleServiceDescriptor {
	return roleServiceDescriptor
}

func initDescriptors() {
	roleServiceDescriptor = &RoleServiceDescriptor{}
	getRoleDescriptor = &GetRoleDescriptor{}
	batchGetRolesDescriptor = &BatchGetRolesDescriptor{}
	listRolesDescriptor = &ListRolesDescriptor{}
	watchRoleDescriptor = &WatchRoleDescriptor{}
	watchRolesDescriptor = &WatchRolesDescriptor{}
	createRoleDescriptor = &CreateRoleDescriptor{}
	updateRoleDescriptor = &UpdateRoleDescriptor{}
	deleteRoleDescriptor = &DeleteRoleDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(roleServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(getRoleDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(batchGetRolesDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(listRolesDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(watchRoleDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(watchRolesDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(createRoleDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(updateRoleDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(deleteRoleDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}
