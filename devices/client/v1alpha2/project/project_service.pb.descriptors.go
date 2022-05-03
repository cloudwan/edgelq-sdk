// Code generated by protoc-gen-goten-client
// API: ProjectService
// DO NOT EDIT!!!

package project_client

import (
	"google.golang.org/protobuf/proto"

	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/project"
	empty "github.com/golang/protobuf/ptypes/empty"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Message(nil)
	_ = gotenclient.MethodDescriptor(nil)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &empty.Empty{}
)

var (
	descriptorsInitialized     bool
	projectServiceDescriptor   *ProjectServiceDescriptor
	getProjectDescriptor       *GetProjectDescriptor
	batchGetProjectsDescriptor *BatchGetProjectsDescriptor
	listProjectsDescriptor     *ListProjectsDescriptor
	watchProjectDescriptor     *WatchProjectDescriptor
	watchProjectsDescriptor    *WatchProjectsDescriptor
	createProjectDescriptor    *CreateProjectDescriptor
	updateProjectDescriptor    *UpdateProjectDescriptor
	deleteProjectDescriptor    *DeleteProjectDescriptor
)

type GetProjectDescriptor struct{}

type GetProjectDescriptorClientMsgHandle struct{}

type GetProjectDescriptorServerMsgHandle struct{}

func (d *GetProjectDescriptor) NewEmptyClientMsg() proto.Message {
	return &GetProjectRequest{}
}

func (d *GetProjectDescriptor) NewEmptyServerMsg() proto.Message {
	return &project.Project{}
}

func (d *GetProjectDescriptor) IsUnary() bool {
	return true
}

func (d *GetProjectDescriptor) IsClientStream() bool {
	return false
}

func (d *GetProjectDescriptor) IsServerStream() bool {
	return false
}

func (d *GetProjectDescriptor) IsCollectionSubject() bool {
	return false
}

func (d *GetProjectDescriptor) IsPluralSubject() bool {
	return false
}

func (d *GetProjectDescriptor) HasSubjectResource() bool {
	return true
}

func (d *GetProjectDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *GetProjectDescriptor) GetVerb() string {
	return "get"
}

func (d *GetProjectDescriptor) GetMethodName() string {
	return "GetProject"
}

func (d *GetProjectDescriptor) GetFullMethodName() string {
	return "/ntt.devices.v1alpha2.ProjectService/GetProject"
}

func (d *GetProjectDescriptor) GetProtoPkgName() string {
	return "ntt.devices.v1alpha2"
}

func (d *GetProjectDescriptor) GetApiName() string {
	return "ProjectService"
}

func (d *GetProjectDescriptor) GetServiceDomain() string {
	return "devices.edgelq.com"
}

func (d *GetProjectDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *GetProjectDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return projectServiceDescriptor
}

func (d *GetProjectDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return project.GetDescriptor()
}

func (d *GetProjectDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetProjectDescriptorClientMsgHandle{}
}

func (d *GetProjectDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetProjectDescriptorServerMsgHandle{}
}

func (h *GetProjectDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*GetProjectRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*GetProjectRequest) *project.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	if ref := typedMsg.GetName(); ref != nil {
		return &ref.Name
	}
	return (*project.Name)(nil)
}

func (h *GetProjectDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*GetProjectRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*GetProjectRequest) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *GetProjectDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *GetProjectDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*project.Project)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*project.Project) *project.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return typedMsg.GetName()
}

func (h *GetProjectDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*project.Project)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*project.Project) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *GetProjectDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetGetProjectDescriptor() *GetProjectDescriptor {
	return getProjectDescriptor
}

type BatchGetProjectsDescriptor struct{}

type BatchGetProjectsDescriptorClientMsgHandle struct{}

type BatchGetProjectsDescriptorServerMsgHandle struct{}

func (d *BatchGetProjectsDescriptor) NewEmptyClientMsg() proto.Message {
	return &BatchGetProjectsRequest{}
}

func (d *BatchGetProjectsDescriptor) NewEmptyServerMsg() proto.Message {
	return &BatchGetProjectsResponse{}
}

func (d *BatchGetProjectsDescriptor) IsUnary() bool {
	return true
}

func (d *BatchGetProjectsDescriptor) IsClientStream() bool {
	return false
}

func (d *BatchGetProjectsDescriptor) IsServerStream() bool {
	return false
}

func (d *BatchGetProjectsDescriptor) IsCollectionSubject() bool {
	return true
}

func (d *BatchGetProjectsDescriptor) IsPluralSubject() bool {
	return true
}

func (d *BatchGetProjectsDescriptor) HasSubjectResource() bool {
	return true
}

func (d *BatchGetProjectsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *BatchGetProjectsDescriptor) GetVerb() string {
	return "batchGet"
}

func (d *BatchGetProjectsDescriptor) GetMethodName() string {
	return "BatchGetProjects"
}

func (d *BatchGetProjectsDescriptor) GetFullMethodName() string {
	return "/ntt.devices.v1alpha2.ProjectService/BatchGetProjects"
}

func (d *BatchGetProjectsDescriptor) GetProtoPkgName() string {
	return "ntt.devices.v1alpha2"
}

func (d *BatchGetProjectsDescriptor) GetApiName() string {
	return "ProjectService"
}

func (d *BatchGetProjectsDescriptor) GetServiceDomain() string {
	return "devices.edgelq.com"
}

func (d *BatchGetProjectsDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *BatchGetProjectsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return projectServiceDescriptor
}

func (d *BatchGetProjectsDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return project.GetDescriptor()
}

func (d *BatchGetProjectsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &BatchGetProjectsDescriptorClientMsgHandle{}
}

func (d *BatchGetProjectsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &BatchGetProjectsDescriptorServerMsgHandle{}
}

func (h *BatchGetProjectsDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*BatchGetProjectsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*BatchGetProjectsRequest) *project.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *BatchGetProjectsDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*BatchGetProjectsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*BatchGetProjectsRequest) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	if refs := typedMsg.GetNames(); len(refs) > 0 {
		list := make(project.ProjectNameList, 0, len(refs))
		for _, ref := range refs {
			list = append(list, &ref.Name)
		}
		return list
	}
	return (project.ProjectNameList)(nil)
}

func (h *BatchGetProjectsDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *BatchGetProjectsDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*BatchGetProjectsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*BatchGetProjectsResponse) *project.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *BatchGetProjectsDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*BatchGetProjectsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*BatchGetProjectsResponse) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	resources := typedMsg.GetProjects()
	list := make(project.ProjectNameList, 0, len(resources))
	for _, res := range resources {
		list = append(list, res.GetName())
	}
	return list
}

func (h *BatchGetProjectsDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetBatchGetProjectsDescriptor() *BatchGetProjectsDescriptor {
	return batchGetProjectsDescriptor
}

type ListProjectsDescriptor struct{}

type ListProjectsDescriptorClientMsgHandle struct{}

type ListProjectsDescriptorServerMsgHandle struct{}

func (d *ListProjectsDescriptor) NewEmptyClientMsg() proto.Message {
	return &ListProjectsRequest{}
}

func (d *ListProjectsDescriptor) NewEmptyServerMsg() proto.Message {
	return &ListProjectsResponse{}
}

func (d *ListProjectsDescriptor) IsUnary() bool {
	return true
}

func (d *ListProjectsDescriptor) IsClientStream() bool {
	return false
}

func (d *ListProjectsDescriptor) IsServerStream() bool {
	return false
}

func (d *ListProjectsDescriptor) IsCollectionSubject() bool {
	return true
}

func (d *ListProjectsDescriptor) IsPluralSubject() bool {
	return true
}

func (d *ListProjectsDescriptor) HasSubjectResource() bool {
	return true
}

func (d *ListProjectsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *ListProjectsDescriptor) GetVerb() string {
	return "list"
}

func (d *ListProjectsDescriptor) GetMethodName() string {
	return "ListProjects"
}

func (d *ListProjectsDescriptor) GetFullMethodName() string {
	return "/ntt.devices.v1alpha2.ProjectService/ListProjects"
}

func (d *ListProjectsDescriptor) GetProtoPkgName() string {
	return "ntt.devices.v1alpha2"
}

func (d *ListProjectsDescriptor) GetApiName() string {
	return "ProjectService"
}

func (d *ListProjectsDescriptor) GetServiceDomain() string {
	return "devices.edgelq.com"
}

func (d *ListProjectsDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *ListProjectsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return projectServiceDescriptor
}

func (d *ListProjectsDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return project.GetDescriptor()
}

func (d *ListProjectsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ListProjectsDescriptorClientMsgHandle{}
}

func (d *ListProjectsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ListProjectsDescriptorServerMsgHandle{}
}

func (h *ListProjectsDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListProjectsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*ListProjectsRequest) *project.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *ListProjectsDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ListProjectsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*ListProjectsRequest) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *ListProjectsDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *ListProjectsDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListProjectsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*ListProjectsResponse) *project.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *ListProjectsDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ListProjectsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*ListProjectsResponse) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	resources := typedMsg.GetProjects()
	list := make(project.ProjectNameList, 0, len(resources))
	for _, res := range resources {
		list = append(list, res.GetName())
	}
	return list
}

func (h *ListProjectsDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetListProjectsDescriptor() *ListProjectsDescriptor {
	return listProjectsDescriptor
}

type WatchProjectDescriptor struct{}

type WatchProjectDescriptorClientMsgHandle struct{}

type WatchProjectDescriptorServerMsgHandle struct{}

func (d *WatchProjectDescriptor) NewEmptyClientMsg() proto.Message {
	return &WatchProjectRequest{}
}

func (d *WatchProjectDescriptor) NewEmptyServerMsg() proto.Message {
	return &WatchProjectResponse{}
}

func (d *WatchProjectDescriptor) IsUnary() bool {
	return false
}

func (d *WatchProjectDescriptor) IsClientStream() bool {
	return false
}

func (d *WatchProjectDescriptor) IsServerStream() bool {
	return true
}

func (d *WatchProjectDescriptor) IsCollectionSubject() bool {
	return false
}

func (d *WatchProjectDescriptor) IsPluralSubject() bool {
	return false
}

func (d *WatchProjectDescriptor) HasSubjectResource() bool {
	return true
}

func (d *WatchProjectDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *WatchProjectDescriptor) GetVerb() string {
	return "watch"
}

func (d *WatchProjectDescriptor) GetMethodName() string {
	return "WatchProject"
}

func (d *WatchProjectDescriptor) GetFullMethodName() string {
	return "/ntt.devices.v1alpha2.ProjectService/WatchProject"
}

func (d *WatchProjectDescriptor) GetProtoPkgName() string {
	return "ntt.devices.v1alpha2"
}

func (d *WatchProjectDescriptor) GetApiName() string {
	return "ProjectService"
}

func (d *WatchProjectDescriptor) GetServiceDomain() string {
	return "devices.edgelq.com"
}

func (d *WatchProjectDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *WatchProjectDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return projectServiceDescriptor
}

func (d *WatchProjectDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return project.GetDescriptor()
}

func (d *WatchProjectDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &WatchProjectDescriptorClientMsgHandle{}
}

func (d *WatchProjectDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &WatchProjectDescriptorServerMsgHandle{}
}

func (h *WatchProjectDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*WatchProjectRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*WatchProjectRequest) *project.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	if ref := typedMsg.GetName(); ref != nil {
		return &ref.Name
	}
	return (*project.Name)(nil)
}

func (h *WatchProjectDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*WatchProjectRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*WatchProjectRequest) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *WatchProjectDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *WatchProjectDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*WatchProjectResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*WatchProjectResponse) *project.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	if typedMsg.GetChange() != nil {
		switch tResChange := typedMsg.GetChange().ChangeType.(type) {
		case *project.ProjectChange_Added_:
			return tResChange.Added.GetProject().GetName()
		case *project.ProjectChange_Modified_:
			return tResChange.Modified.GetName()
		case *project.ProjectChange_Removed_:
			return tResChange.Removed.GetName()
		case *project.ProjectChange_Current_:
			return tResChange.Current.GetProject().GetName()
		}
	}
	return (*project.Name)(nil)
}

func (h *WatchProjectDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*WatchProjectResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*WatchProjectResponse) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *WatchProjectDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetWatchProjectDescriptor() *WatchProjectDescriptor {
	return watchProjectDescriptor
}

type WatchProjectsDescriptor struct{}

type WatchProjectsDescriptorClientMsgHandle struct{}

type WatchProjectsDescriptorServerMsgHandle struct{}

func (d *WatchProjectsDescriptor) NewEmptyClientMsg() proto.Message {
	return &WatchProjectsRequest{}
}

func (d *WatchProjectsDescriptor) NewEmptyServerMsg() proto.Message {
	return &WatchProjectsResponse{}
}

func (d *WatchProjectsDescriptor) IsUnary() bool {
	return false
}

func (d *WatchProjectsDescriptor) IsClientStream() bool {
	return false
}

func (d *WatchProjectsDescriptor) IsServerStream() bool {
	return true
}

func (d *WatchProjectsDescriptor) IsCollectionSubject() bool {
	return true
}

func (d *WatchProjectsDescriptor) IsPluralSubject() bool {
	return true
}

func (d *WatchProjectsDescriptor) HasSubjectResource() bool {
	return true
}

func (d *WatchProjectsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *WatchProjectsDescriptor) GetVerb() string {
	return "watch"
}

func (d *WatchProjectsDescriptor) GetMethodName() string {
	return "WatchProjects"
}

func (d *WatchProjectsDescriptor) GetFullMethodName() string {
	return "/ntt.devices.v1alpha2.ProjectService/WatchProjects"
}

func (d *WatchProjectsDescriptor) GetProtoPkgName() string {
	return "ntt.devices.v1alpha2"
}

func (d *WatchProjectsDescriptor) GetApiName() string {
	return "ProjectService"
}

func (d *WatchProjectsDescriptor) GetServiceDomain() string {
	return "devices.edgelq.com"
}

func (d *WatchProjectsDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *WatchProjectsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return projectServiceDescriptor
}

func (d *WatchProjectsDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return project.GetDescriptor()
}

func (d *WatchProjectsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &WatchProjectsDescriptorClientMsgHandle{}
}

func (d *WatchProjectsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &WatchProjectsDescriptorServerMsgHandle{}
}

func (h *WatchProjectsDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*WatchProjectsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*WatchProjectsRequest) *project.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *WatchProjectsDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*WatchProjectsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*WatchProjectsRequest) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *WatchProjectsDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *WatchProjectsDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*WatchProjectsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*WatchProjectsResponse) *project.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *WatchProjectsDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*WatchProjectsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*WatchProjectsResponse) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	resourceChanges := typedMsg.GetProjectChanges()
	list := make(project.ProjectNameList, 0, len(resourceChanges))
	for _, resChange := range resourceChanges {
		switch tResChange := resChange.ChangeType.(type) {
		case *project.ProjectChange_Added_:
			list = append(list, tResChange.Added.GetProject().GetName())
		case *project.ProjectChange_Modified_:
			list = append(list, tResChange.Modified.GetName())
		case *project.ProjectChange_Removed_:
			list = append(list, tResChange.Removed.GetName())
		case *project.ProjectChange_Current_:
			list = append(list, tResChange.Current.GetProject().GetName())
		}
	}
	return list
}

func (h *WatchProjectsDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetWatchProjectsDescriptor() *WatchProjectsDescriptor {
	return watchProjectsDescriptor
}

type CreateProjectDescriptor struct{}

type CreateProjectDescriptorClientMsgHandle struct{}

type CreateProjectDescriptorServerMsgHandle struct{}

func (d *CreateProjectDescriptor) NewEmptyClientMsg() proto.Message {
	return &CreateProjectRequest{}
}

func (d *CreateProjectDescriptor) NewEmptyServerMsg() proto.Message {
	return &project.Project{}
}

func (d *CreateProjectDescriptor) IsUnary() bool {
	return true
}

func (d *CreateProjectDescriptor) IsClientStream() bool {
	return false
}

func (d *CreateProjectDescriptor) IsServerStream() bool {
	return false
}

func (d *CreateProjectDescriptor) IsCollectionSubject() bool {
	return true
}

func (d *CreateProjectDescriptor) IsPluralSubject() bool {
	return false
}

func (d *CreateProjectDescriptor) HasSubjectResource() bool {
	return true
}

func (d *CreateProjectDescriptor) RequestHasResourceBody() bool {
	return true
}

func (d *CreateProjectDescriptor) GetVerb() string {
	return "create"
}

func (d *CreateProjectDescriptor) GetMethodName() string {
	return "CreateProject"
}

func (d *CreateProjectDescriptor) GetFullMethodName() string {
	return "/ntt.devices.v1alpha2.ProjectService/CreateProject"
}

func (d *CreateProjectDescriptor) GetProtoPkgName() string {
	return "ntt.devices.v1alpha2"
}

func (d *CreateProjectDescriptor) GetApiName() string {
	return "ProjectService"
}

func (d *CreateProjectDescriptor) GetServiceDomain() string {
	return "devices.edgelq.com"
}

func (d *CreateProjectDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *CreateProjectDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return projectServiceDescriptor
}

func (d *CreateProjectDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return project.GetDescriptor()
}

func (d *CreateProjectDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CreateProjectDescriptorClientMsgHandle{}
}

func (d *CreateProjectDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CreateProjectDescriptorServerMsgHandle{}
}

func (h *CreateProjectDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CreateProjectRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*CreateProjectRequest) *project.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return typedMsg.GetProject().GetName()
}

func (h *CreateProjectDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*CreateProjectRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*CreateProjectRequest) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *CreateProjectDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *CreateProjectDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*project.Project)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*project.Project) *project.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return typedMsg.GetName()
}

func (h *CreateProjectDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*project.Project)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*project.Project) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *CreateProjectDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetCreateProjectDescriptor() *CreateProjectDescriptor {
	return createProjectDescriptor
}

type UpdateProjectDescriptor struct{}

type UpdateProjectDescriptorClientMsgHandle struct{}

type UpdateProjectDescriptorServerMsgHandle struct{}

func (d *UpdateProjectDescriptor) NewEmptyClientMsg() proto.Message {
	return &UpdateProjectRequest{}
}

func (d *UpdateProjectDescriptor) NewEmptyServerMsg() proto.Message {
	return &project.Project{}
}

func (d *UpdateProjectDescriptor) IsUnary() bool {
	return true
}

func (d *UpdateProjectDescriptor) IsClientStream() bool {
	return false
}

func (d *UpdateProjectDescriptor) IsServerStream() bool {
	return false
}

func (d *UpdateProjectDescriptor) IsCollectionSubject() bool {
	return false
}

func (d *UpdateProjectDescriptor) IsPluralSubject() bool {
	return false
}

func (d *UpdateProjectDescriptor) HasSubjectResource() bool {
	return true
}

func (d *UpdateProjectDescriptor) RequestHasResourceBody() bool {
	return true
}

func (d *UpdateProjectDescriptor) GetVerb() string {
	return "update"
}

func (d *UpdateProjectDescriptor) GetMethodName() string {
	return "UpdateProject"
}

func (d *UpdateProjectDescriptor) GetFullMethodName() string {
	return "/ntt.devices.v1alpha2.ProjectService/UpdateProject"
}

func (d *UpdateProjectDescriptor) GetProtoPkgName() string {
	return "ntt.devices.v1alpha2"
}

func (d *UpdateProjectDescriptor) GetApiName() string {
	return "ProjectService"
}

func (d *UpdateProjectDescriptor) GetServiceDomain() string {
	return "devices.edgelq.com"
}

func (d *UpdateProjectDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *UpdateProjectDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return projectServiceDescriptor
}

func (d *UpdateProjectDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return project.GetDescriptor()
}

func (d *UpdateProjectDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &UpdateProjectDescriptorClientMsgHandle{}
}

func (d *UpdateProjectDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &UpdateProjectDescriptorServerMsgHandle{}
}

func (h *UpdateProjectDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*UpdateProjectRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*UpdateProjectRequest) *project.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return typedMsg.GetProject().GetName()
}

func (h *UpdateProjectDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*UpdateProjectRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*UpdateProjectRequest) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *UpdateProjectDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *UpdateProjectDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*project.Project)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*project.Project) *project.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return typedMsg.GetName()
}

func (h *UpdateProjectDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*project.Project)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*project.Project) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *UpdateProjectDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetUpdateProjectDescriptor() *UpdateProjectDescriptor {
	return updateProjectDescriptor
}

type DeleteProjectDescriptor struct{}

type DeleteProjectDescriptorClientMsgHandle struct{}

type DeleteProjectDescriptorServerMsgHandle struct{}

func (d *DeleteProjectDescriptor) NewEmptyClientMsg() proto.Message {
	return &DeleteProjectRequest{}
}

func (d *DeleteProjectDescriptor) NewEmptyServerMsg() proto.Message {
	return &empty.Empty{}
}

func (d *DeleteProjectDescriptor) IsUnary() bool {
	return true
}

func (d *DeleteProjectDescriptor) IsClientStream() bool {
	return false
}

func (d *DeleteProjectDescriptor) IsServerStream() bool {
	return false
}

func (d *DeleteProjectDescriptor) IsCollectionSubject() bool {
	return false
}

func (d *DeleteProjectDescriptor) IsPluralSubject() bool {
	return false
}

func (d *DeleteProjectDescriptor) HasSubjectResource() bool {
	return true
}

func (d *DeleteProjectDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *DeleteProjectDescriptor) GetVerb() string {
	return "delete"
}

func (d *DeleteProjectDescriptor) GetMethodName() string {
	return "DeleteProject"
}

func (d *DeleteProjectDescriptor) GetFullMethodName() string {
	return "/ntt.devices.v1alpha2.ProjectService/DeleteProject"
}

func (d *DeleteProjectDescriptor) GetProtoPkgName() string {
	return "ntt.devices.v1alpha2"
}

func (d *DeleteProjectDescriptor) GetApiName() string {
	return "ProjectService"
}

func (d *DeleteProjectDescriptor) GetServiceDomain() string {
	return "devices.edgelq.com"
}

func (d *DeleteProjectDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *DeleteProjectDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return projectServiceDescriptor
}

func (d *DeleteProjectDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return project.GetDescriptor()
}

func (d *DeleteProjectDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &DeleteProjectDescriptorClientMsgHandle{}
}

func (d *DeleteProjectDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &DeleteProjectDescriptorServerMsgHandle{}
}

func (h *DeleteProjectDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*DeleteProjectRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*DeleteProjectRequest) *project.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	if ref := typedMsg.GetName(); ref != nil {
		return &ref.Name
	}
	return (*project.Name)(nil)
}

func (h *DeleteProjectDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*DeleteProjectRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*DeleteProjectRequest) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *DeleteProjectDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *DeleteProjectDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*empty.Empty)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*empty.Empty) *project.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *DeleteProjectDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*empty.Empty)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*empty.Empty) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *DeleteProjectDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetDeleteProjectDescriptor() *DeleteProjectDescriptor {
	return deleteProjectDescriptor
}

type ProjectServiceDescriptor struct{}

func (d *ProjectServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		getProjectDescriptor,
		batchGetProjectsDescriptor,
		listProjectsDescriptor,
		watchProjectDescriptor,
		watchProjectsDescriptor,
		createProjectDescriptor,
		updateProjectDescriptor,
		deleteProjectDescriptor,
	}
}

func (d *ProjectServiceDescriptor) GetFullAPIName() string {
	return "/ntt.devices.v1alpha2.ProjectService"
}

func (d *ProjectServiceDescriptor) GetProtoPkgName() string {
	return "ntt.devices.v1alpha2"
}

func (d *ProjectServiceDescriptor) GetApiName() string {
	return "ProjectService"
}

func (d *ProjectServiceDescriptor) GetServiceDomain() string {
	return "devices.edgelq.com"
}

func (d *ProjectServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func GetProjectServiceDescriptor() *ProjectServiceDescriptor {
	return projectServiceDescriptor
}

func initDescriptors() {
	projectServiceDescriptor = &ProjectServiceDescriptor{}
	getProjectDescriptor = &GetProjectDescriptor{}
	batchGetProjectsDescriptor = &BatchGetProjectsDescriptor{}
	listProjectsDescriptor = &ListProjectsDescriptor{}
	watchProjectDescriptor = &WatchProjectDescriptor{}
	watchProjectsDescriptor = &WatchProjectsDescriptor{}
	createProjectDescriptor = &CreateProjectDescriptor{}
	updateProjectDescriptor = &UpdateProjectDescriptor{}
	deleteProjectDescriptor = &DeleteProjectDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(projectServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(getProjectDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(batchGetProjectsDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(listProjectsDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(watchProjectDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(watchProjectsDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(createProjectDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(updateProjectDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(deleteProjectDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}
