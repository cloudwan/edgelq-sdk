// Code generated by protoc-gen-goten-client
// API: ServiceService
// DO NOT EDIT!!!

package service_client

import (
	"google.golang.org/protobuf/proto"

	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
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
	_ = &service.Service{}
	_ = &empty.Empty{}
)

var (
	descriptorsInitialized     bool
	serviceServiceDescriptor   *ServiceServiceDescriptor
	getServiceDescriptor       *GetServiceDescriptor
	batchGetServicesDescriptor *BatchGetServicesDescriptor
	listServicesDescriptor     *ListServicesDescriptor
	watchServiceDescriptor     *WatchServiceDescriptor
	watchServicesDescriptor    *WatchServicesDescriptor
	createServiceDescriptor    *CreateServiceDescriptor
	updateServiceDescriptor    *UpdateServiceDescriptor
	deleteServiceDescriptor    *DeleteServiceDescriptor
)

type GetServiceDescriptor struct{}

type GetServiceDescriptorClientMsgHandle struct{}

type GetServiceDescriptorServerMsgHandle struct{}

func (d *GetServiceDescriptor) NewEmptyClientMsg() proto.Message {
	return &GetServiceRequest{}
}

func (d *GetServiceDescriptor) NewEmptyServerMsg() proto.Message {
	return &service.Service{}
}

func (d *GetServiceDescriptor) IsUnary() bool {
	return true
}

func (d *GetServiceDescriptor) IsClientStream() bool {
	return false
}

func (d *GetServiceDescriptor) IsServerStream() bool {
	return false
}

func (d *GetServiceDescriptor) IsCollectionSubject() bool {
	return false
}

func (d *GetServiceDescriptor) IsPluralSubject() bool {
	return false
}

func (d *GetServiceDescriptor) HasSubjectResource() bool {
	return true
}

func (d *GetServiceDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *GetServiceDescriptor) GetVerb() string {
	return "get"
}

func (d *GetServiceDescriptor) GetMethodName() string {
	return "GetService"
}

func (d *GetServiceDescriptor) GetFullMethodName() string {
	return "/ntt.meta.v1alpha2.ServiceService/GetService"
}

func (d *GetServiceDescriptor) GetProtoPkgName() string {
	return "ntt.meta.v1alpha2"
}

func (d *GetServiceDescriptor) GetApiName() string {
	return "ServiceService"
}

func (d *GetServiceDescriptor) GetServiceDomain() string {
	return "meta.edgelq.com"
}

func (d *GetServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *GetServiceDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return serviceServiceDescriptor
}

func (d *GetServiceDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return service.GetDescriptor()
}

func (d *GetServiceDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetServiceDescriptorClientMsgHandle{}
}

func (d *GetServiceDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetServiceDescriptorServerMsgHandle{}
}

func (h *GetServiceDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*GetServiceRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*GetServiceRequest) *service.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	if ref := typedMsg.GetName(); ref != nil {
		return &ref.Name
	}
	return (*service.Name)(nil)
}

func (h *GetServiceDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*GetServiceRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*GetServiceRequest) []*service.Name
	})
	if ok {
		return service.ServiceNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *GetServiceDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *GetServiceDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*service.Service)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*service.Service) *service.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return typedMsg.GetName()
}

func (h *GetServiceDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*service.Service)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*service.Service) []*service.Name
	})
	if ok {
		return service.ServiceNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *GetServiceDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetGetServiceDescriptor() *GetServiceDescriptor {
	return getServiceDescriptor
}

type BatchGetServicesDescriptor struct{}

type BatchGetServicesDescriptorClientMsgHandle struct{}

type BatchGetServicesDescriptorServerMsgHandle struct{}

func (d *BatchGetServicesDescriptor) NewEmptyClientMsg() proto.Message {
	return &BatchGetServicesRequest{}
}

func (d *BatchGetServicesDescriptor) NewEmptyServerMsg() proto.Message {
	return &BatchGetServicesResponse{}
}

func (d *BatchGetServicesDescriptor) IsUnary() bool {
	return true
}

func (d *BatchGetServicesDescriptor) IsClientStream() bool {
	return false
}

func (d *BatchGetServicesDescriptor) IsServerStream() bool {
	return false
}

func (d *BatchGetServicesDescriptor) IsCollectionSubject() bool {
	return true
}

func (d *BatchGetServicesDescriptor) IsPluralSubject() bool {
	return true
}

func (d *BatchGetServicesDescriptor) HasSubjectResource() bool {
	return true
}

func (d *BatchGetServicesDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *BatchGetServicesDescriptor) GetVerb() string {
	return "batchGet"
}

func (d *BatchGetServicesDescriptor) GetMethodName() string {
	return "BatchGetServices"
}

func (d *BatchGetServicesDescriptor) GetFullMethodName() string {
	return "/ntt.meta.v1alpha2.ServiceService/BatchGetServices"
}

func (d *BatchGetServicesDescriptor) GetProtoPkgName() string {
	return "ntt.meta.v1alpha2"
}

func (d *BatchGetServicesDescriptor) GetApiName() string {
	return "ServiceService"
}

func (d *BatchGetServicesDescriptor) GetServiceDomain() string {
	return "meta.edgelq.com"
}

func (d *BatchGetServicesDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *BatchGetServicesDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return serviceServiceDescriptor
}

func (d *BatchGetServicesDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return service.GetDescriptor()
}

func (d *BatchGetServicesDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &BatchGetServicesDescriptorClientMsgHandle{}
}

func (d *BatchGetServicesDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &BatchGetServicesDescriptorServerMsgHandle{}
}

func (h *BatchGetServicesDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*BatchGetServicesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*BatchGetServicesRequest) *service.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *BatchGetServicesDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*BatchGetServicesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*BatchGetServicesRequest) []*service.Name
	})
	if ok {
		return service.ServiceNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	if refs := typedMsg.GetNames(); len(refs) > 0 {
		list := make(service.ServiceNameList, 0, len(refs))
		for _, ref := range refs {
			list = append(list, &ref.Name)
		}
		return list
	}
	return (service.ServiceNameList)(nil)
}

func (h *BatchGetServicesDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *BatchGetServicesDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*BatchGetServicesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*BatchGetServicesResponse) *service.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *BatchGetServicesDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*BatchGetServicesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*BatchGetServicesResponse) []*service.Name
	})
	if ok {
		return service.ServiceNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	resources := typedMsg.GetServices()
	list := make(service.ServiceNameList, 0, len(resources))
	for _, res := range resources {
		list = append(list, res.GetName())
	}
	return list
}

func (h *BatchGetServicesDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetBatchGetServicesDescriptor() *BatchGetServicesDescriptor {
	return batchGetServicesDescriptor
}

type ListServicesDescriptor struct{}

type ListServicesDescriptorClientMsgHandle struct{}

type ListServicesDescriptorServerMsgHandle struct{}

func (d *ListServicesDescriptor) NewEmptyClientMsg() proto.Message {
	return &ListServicesRequest{}
}

func (d *ListServicesDescriptor) NewEmptyServerMsg() proto.Message {
	return &ListServicesResponse{}
}

func (d *ListServicesDescriptor) IsUnary() bool {
	return true
}

func (d *ListServicesDescriptor) IsClientStream() bool {
	return false
}

func (d *ListServicesDescriptor) IsServerStream() bool {
	return false
}

func (d *ListServicesDescriptor) IsCollectionSubject() bool {
	return true
}

func (d *ListServicesDescriptor) IsPluralSubject() bool {
	return true
}

func (d *ListServicesDescriptor) HasSubjectResource() bool {
	return true
}

func (d *ListServicesDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *ListServicesDescriptor) GetVerb() string {
	return "list"
}

func (d *ListServicesDescriptor) GetMethodName() string {
	return "ListServices"
}

func (d *ListServicesDescriptor) GetFullMethodName() string {
	return "/ntt.meta.v1alpha2.ServiceService/ListServices"
}

func (d *ListServicesDescriptor) GetProtoPkgName() string {
	return "ntt.meta.v1alpha2"
}

func (d *ListServicesDescriptor) GetApiName() string {
	return "ServiceService"
}

func (d *ListServicesDescriptor) GetServiceDomain() string {
	return "meta.edgelq.com"
}

func (d *ListServicesDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *ListServicesDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return serviceServiceDescriptor
}

func (d *ListServicesDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return service.GetDescriptor()
}

func (d *ListServicesDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ListServicesDescriptorClientMsgHandle{}
}

func (d *ListServicesDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ListServicesDescriptorServerMsgHandle{}
}

func (h *ListServicesDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListServicesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*ListServicesRequest) *service.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *ListServicesDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ListServicesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*ListServicesRequest) []*service.Name
	})
	if ok {
		return service.ServiceNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *ListServicesDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *ListServicesDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListServicesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*ListServicesResponse) *service.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *ListServicesDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ListServicesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*ListServicesResponse) []*service.Name
	})
	if ok {
		return service.ServiceNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	resources := typedMsg.GetServices()
	list := make(service.ServiceNameList, 0, len(resources))
	for _, res := range resources {
		list = append(list, res.GetName())
	}
	return list
}

func (h *ListServicesDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetListServicesDescriptor() *ListServicesDescriptor {
	return listServicesDescriptor
}

type WatchServiceDescriptor struct{}

type WatchServiceDescriptorClientMsgHandle struct{}

type WatchServiceDescriptorServerMsgHandle struct{}

func (d *WatchServiceDescriptor) NewEmptyClientMsg() proto.Message {
	return &WatchServiceRequest{}
}

func (d *WatchServiceDescriptor) NewEmptyServerMsg() proto.Message {
	return &WatchServiceResponse{}
}

func (d *WatchServiceDescriptor) IsUnary() bool {
	return false
}

func (d *WatchServiceDescriptor) IsClientStream() bool {
	return false
}

func (d *WatchServiceDescriptor) IsServerStream() bool {
	return true
}

func (d *WatchServiceDescriptor) IsCollectionSubject() bool {
	return false
}

func (d *WatchServiceDescriptor) IsPluralSubject() bool {
	return false
}

func (d *WatchServiceDescriptor) HasSubjectResource() bool {
	return true
}

func (d *WatchServiceDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *WatchServiceDescriptor) GetVerb() string {
	return "watch"
}

func (d *WatchServiceDescriptor) GetMethodName() string {
	return "WatchService"
}

func (d *WatchServiceDescriptor) GetFullMethodName() string {
	return "/ntt.meta.v1alpha2.ServiceService/WatchService"
}

func (d *WatchServiceDescriptor) GetProtoPkgName() string {
	return "ntt.meta.v1alpha2"
}

func (d *WatchServiceDescriptor) GetApiName() string {
	return "ServiceService"
}

func (d *WatchServiceDescriptor) GetServiceDomain() string {
	return "meta.edgelq.com"
}

func (d *WatchServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *WatchServiceDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return serviceServiceDescriptor
}

func (d *WatchServiceDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return service.GetDescriptor()
}

func (d *WatchServiceDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &WatchServiceDescriptorClientMsgHandle{}
}

func (d *WatchServiceDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &WatchServiceDescriptorServerMsgHandle{}
}

func (h *WatchServiceDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*WatchServiceRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*WatchServiceRequest) *service.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	if ref := typedMsg.GetName(); ref != nil {
		return &ref.Name
	}
	return (*service.Name)(nil)
}

func (h *WatchServiceDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*WatchServiceRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*WatchServiceRequest) []*service.Name
	})
	if ok {
		return service.ServiceNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *WatchServiceDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *WatchServiceDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*WatchServiceResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*WatchServiceResponse) *service.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	if typedMsg.GetChange() != nil {
		switch tResChange := typedMsg.GetChange().ChangeType.(type) {
		case *service.ServiceChange_Added_:
			return tResChange.Added.GetService().GetName()
		case *service.ServiceChange_Modified_:
			return tResChange.Modified.GetName()
		case *service.ServiceChange_Removed_:
			return tResChange.Removed.GetName()
		case *service.ServiceChange_Current_:
			return tResChange.Current.GetService().GetName()
		}
	}
	return (*service.Name)(nil)
}

func (h *WatchServiceDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*WatchServiceResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*WatchServiceResponse) []*service.Name
	})
	if ok {
		return service.ServiceNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *WatchServiceDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetWatchServiceDescriptor() *WatchServiceDescriptor {
	return watchServiceDescriptor
}

type WatchServicesDescriptor struct{}

type WatchServicesDescriptorClientMsgHandle struct{}

type WatchServicesDescriptorServerMsgHandle struct{}

func (d *WatchServicesDescriptor) NewEmptyClientMsg() proto.Message {
	return &WatchServicesRequest{}
}

func (d *WatchServicesDescriptor) NewEmptyServerMsg() proto.Message {
	return &WatchServicesResponse{}
}

func (d *WatchServicesDescriptor) IsUnary() bool {
	return false
}

func (d *WatchServicesDescriptor) IsClientStream() bool {
	return false
}

func (d *WatchServicesDescriptor) IsServerStream() bool {
	return true
}

func (d *WatchServicesDescriptor) IsCollectionSubject() bool {
	return true
}

func (d *WatchServicesDescriptor) IsPluralSubject() bool {
	return true
}

func (d *WatchServicesDescriptor) HasSubjectResource() bool {
	return true
}

func (d *WatchServicesDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *WatchServicesDescriptor) GetVerb() string {
	return "watch"
}

func (d *WatchServicesDescriptor) GetMethodName() string {
	return "WatchServices"
}

func (d *WatchServicesDescriptor) GetFullMethodName() string {
	return "/ntt.meta.v1alpha2.ServiceService/WatchServices"
}

func (d *WatchServicesDescriptor) GetProtoPkgName() string {
	return "ntt.meta.v1alpha2"
}

func (d *WatchServicesDescriptor) GetApiName() string {
	return "ServiceService"
}

func (d *WatchServicesDescriptor) GetServiceDomain() string {
	return "meta.edgelq.com"
}

func (d *WatchServicesDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *WatchServicesDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return serviceServiceDescriptor
}

func (d *WatchServicesDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return service.GetDescriptor()
}

func (d *WatchServicesDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &WatchServicesDescriptorClientMsgHandle{}
}

func (d *WatchServicesDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &WatchServicesDescriptorServerMsgHandle{}
}

func (h *WatchServicesDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*WatchServicesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*WatchServicesRequest) *service.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *WatchServicesDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*WatchServicesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*WatchServicesRequest) []*service.Name
	})
	if ok {
		return service.ServiceNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *WatchServicesDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *WatchServicesDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*WatchServicesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*WatchServicesResponse) *service.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *WatchServicesDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*WatchServicesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*WatchServicesResponse) []*service.Name
	})
	if ok {
		return service.ServiceNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	resourceChanges := typedMsg.GetServiceChanges()
	list := make(service.ServiceNameList, 0, len(resourceChanges))
	for _, resChange := range resourceChanges {
		switch tResChange := resChange.ChangeType.(type) {
		case *service.ServiceChange_Added_:
			list = append(list, tResChange.Added.GetService().GetName())
		case *service.ServiceChange_Modified_:
			list = append(list, tResChange.Modified.GetName())
		case *service.ServiceChange_Removed_:
			list = append(list, tResChange.Removed.GetName())
		case *service.ServiceChange_Current_:
			list = append(list, tResChange.Current.GetService().GetName())
		}
	}
	return list
}

func (h *WatchServicesDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetWatchServicesDescriptor() *WatchServicesDescriptor {
	return watchServicesDescriptor
}

type CreateServiceDescriptor struct{}

type CreateServiceDescriptorClientMsgHandle struct{}

type CreateServiceDescriptorServerMsgHandle struct{}

func (d *CreateServiceDescriptor) NewEmptyClientMsg() proto.Message {
	return &CreateServiceRequest{}
}

func (d *CreateServiceDescriptor) NewEmptyServerMsg() proto.Message {
	return &service.Service{}
}

func (d *CreateServiceDescriptor) IsUnary() bool {
	return true
}

func (d *CreateServiceDescriptor) IsClientStream() bool {
	return false
}

func (d *CreateServiceDescriptor) IsServerStream() bool {
	return false
}

func (d *CreateServiceDescriptor) IsCollectionSubject() bool {
	return true
}

func (d *CreateServiceDescriptor) IsPluralSubject() bool {
	return false
}

func (d *CreateServiceDescriptor) HasSubjectResource() bool {
	return true
}

func (d *CreateServiceDescriptor) RequestHasResourceBody() bool {
	return true
}

func (d *CreateServiceDescriptor) GetVerb() string {
	return "create"
}

func (d *CreateServiceDescriptor) GetMethodName() string {
	return "CreateService"
}

func (d *CreateServiceDescriptor) GetFullMethodName() string {
	return "/ntt.meta.v1alpha2.ServiceService/CreateService"
}

func (d *CreateServiceDescriptor) GetProtoPkgName() string {
	return "ntt.meta.v1alpha2"
}

func (d *CreateServiceDescriptor) GetApiName() string {
	return "ServiceService"
}

func (d *CreateServiceDescriptor) GetServiceDomain() string {
	return "meta.edgelq.com"
}

func (d *CreateServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *CreateServiceDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return serviceServiceDescriptor
}

func (d *CreateServiceDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return service.GetDescriptor()
}

func (d *CreateServiceDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CreateServiceDescriptorClientMsgHandle{}
}

func (d *CreateServiceDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CreateServiceDescriptorServerMsgHandle{}
}

func (h *CreateServiceDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CreateServiceRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*CreateServiceRequest) *service.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return typedMsg.GetService().GetName()
}

func (h *CreateServiceDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*CreateServiceRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*CreateServiceRequest) []*service.Name
	})
	if ok {
		return service.ServiceNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *CreateServiceDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *CreateServiceDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*service.Service)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*service.Service) *service.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return typedMsg.GetName()
}

func (h *CreateServiceDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*service.Service)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*service.Service) []*service.Name
	})
	if ok {
		return service.ServiceNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *CreateServiceDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetCreateServiceDescriptor() *CreateServiceDescriptor {
	return createServiceDescriptor
}

type UpdateServiceDescriptor struct{}

type UpdateServiceDescriptorClientMsgHandle struct{}

type UpdateServiceDescriptorServerMsgHandle struct{}

func (d *UpdateServiceDescriptor) NewEmptyClientMsg() proto.Message {
	return &UpdateServiceRequest{}
}

func (d *UpdateServiceDescriptor) NewEmptyServerMsg() proto.Message {
	return &service.Service{}
}

func (d *UpdateServiceDescriptor) IsUnary() bool {
	return true
}

func (d *UpdateServiceDescriptor) IsClientStream() bool {
	return false
}

func (d *UpdateServiceDescriptor) IsServerStream() bool {
	return false
}

func (d *UpdateServiceDescriptor) IsCollectionSubject() bool {
	return false
}

func (d *UpdateServiceDescriptor) IsPluralSubject() bool {
	return false
}

func (d *UpdateServiceDescriptor) HasSubjectResource() bool {
	return true
}

func (d *UpdateServiceDescriptor) RequestHasResourceBody() bool {
	return true
}

func (d *UpdateServiceDescriptor) GetVerb() string {
	return "update"
}

func (d *UpdateServiceDescriptor) GetMethodName() string {
	return "UpdateService"
}

func (d *UpdateServiceDescriptor) GetFullMethodName() string {
	return "/ntt.meta.v1alpha2.ServiceService/UpdateService"
}

func (d *UpdateServiceDescriptor) GetProtoPkgName() string {
	return "ntt.meta.v1alpha2"
}

func (d *UpdateServiceDescriptor) GetApiName() string {
	return "ServiceService"
}

func (d *UpdateServiceDescriptor) GetServiceDomain() string {
	return "meta.edgelq.com"
}

func (d *UpdateServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *UpdateServiceDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return serviceServiceDescriptor
}

func (d *UpdateServiceDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return service.GetDescriptor()
}

func (d *UpdateServiceDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &UpdateServiceDescriptorClientMsgHandle{}
}

func (d *UpdateServiceDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &UpdateServiceDescriptorServerMsgHandle{}
}

func (h *UpdateServiceDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*UpdateServiceRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*UpdateServiceRequest) *service.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return typedMsg.GetService().GetName()
}

func (h *UpdateServiceDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*UpdateServiceRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*UpdateServiceRequest) []*service.Name
	})
	if ok {
		return service.ServiceNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *UpdateServiceDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *UpdateServiceDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*service.Service)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*service.Service) *service.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return typedMsg.GetName()
}

func (h *UpdateServiceDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*service.Service)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*service.Service) []*service.Name
	})
	if ok {
		return service.ServiceNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *UpdateServiceDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetUpdateServiceDescriptor() *UpdateServiceDescriptor {
	return updateServiceDescriptor
}

type DeleteServiceDescriptor struct{}

type DeleteServiceDescriptorClientMsgHandle struct{}

type DeleteServiceDescriptorServerMsgHandle struct{}

func (d *DeleteServiceDescriptor) NewEmptyClientMsg() proto.Message {
	return &DeleteServiceRequest{}
}

func (d *DeleteServiceDescriptor) NewEmptyServerMsg() proto.Message {
	return &empty.Empty{}
}

func (d *DeleteServiceDescriptor) IsUnary() bool {
	return true
}

func (d *DeleteServiceDescriptor) IsClientStream() bool {
	return false
}

func (d *DeleteServiceDescriptor) IsServerStream() bool {
	return false
}

func (d *DeleteServiceDescriptor) IsCollectionSubject() bool {
	return false
}

func (d *DeleteServiceDescriptor) IsPluralSubject() bool {
	return false
}

func (d *DeleteServiceDescriptor) HasSubjectResource() bool {
	return true
}

func (d *DeleteServiceDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *DeleteServiceDescriptor) GetVerb() string {
	return "delete"
}

func (d *DeleteServiceDescriptor) GetMethodName() string {
	return "DeleteService"
}

func (d *DeleteServiceDescriptor) GetFullMethodName() string {
	return "/ntt.meta.v1alpha2.ServiceService/DeleteService"
}

func (d *DeleteServiceDescriptor) GetProtoPkgName() string {
	return "ntt.meta.v1alpha2"
}

func (d *DeleteServiceDescriptor) GetApiName() string {
	return "ServiceService"
}

func (d *DeleteServiceDescriptor) GetServiceDomain() string {
	return "meta.edgelq.com"
}

func (d *DeleteServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *DeleteServiceDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return serviceServiceDescriptor
}

func (d *DeleteServiceDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return service.GetDescriptor()
}

func (d *DeleteServiceDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &DeleteServiceDescriptorClientMsgHandle{}
}

func (d *DeleteServiceDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &DeleteServiceDescriptorServerMsgHandle{}
}

func (h *DeleteServiceDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*DeleteServiceRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*DeleteServiceRequest) *service.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	if ref := typedMsg.GetName(); ref != nil {
		return &ref.Name
	}
	return (*service.Name)(nil)
}

func (h *DeleteServiceDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*DeleteServiceRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*DeleteServiceRequest) []*service.Name
	})
	if ok {
		return service.ServiceNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *DeleteServiceDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *DeleteServiceDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*empty.Empty)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*empty.Empty) *service.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *DeleteServiceDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*empty.Empty)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*empty.Empty) []*service.Name
	})
	if ok {
		return service.ServiceNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *DeleteServiceDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetDeleteServiceDescriptor() *DeleteServiceDescriptor {
	return deleteServiceDescriptor
}

type ServiceServiceDescriptor struct{}

func (d *ServiceServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		getServiceDescriptor,
		batchGetServicesDescriptor,
		listServicesDescriptor,
		watchServiceDescriptor,
		watchServicesDescriptor,
		createServiceDescriptor,
		updateServiceDescriptor,
		deleteServiceDescriptor,
	}
}

func (d *ServiceServiceDescriptor) GetFullAPIName() string {
	return "/ntt.meta.v1alpha2.ServiceService"
}

func (d *ServiceServiceDescriptor) GetProtoPkgName() string {
	return "ntt.meta.v1alpha2"
}

func (d *ServiceServiceDescriptor) GetApiName() string {
	return "ServiceService"
}

func (d *ServiceServiceDescriptor) GetServiceDomain() string {
	return "meta.edgelq.com"
}

func (d *ServiceServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func GetServiceServiceDescriptor() *ServiceServiceDescriptor {
	return serviceServiceDescriptor
}

func initDescriptors() {
	serviceServiceDescriptor = &ServiceServiceDescriptor{}
	getServiceDescriptor = &GetServiceDescriptor{}
	batchGetServicesDescriptor = &BatchGetServicesDescriptor{}
	listServicesDescriptor = &ListServicesDescriptor{}
	watchServiceDescriptor = &WatchServiceDescriptor{}
	watchServicesDescriptor = &WatchServicesDescriptor{}
	createServiceDescriptor = &CreateServiceDescriptor{}
	updateServiceDescriptor = &UpdateServiceDescriptor{}
	deleteServiceDescriptor = &DeleteServiceDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(serviceServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(getServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(batchGetServicesDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(listServicesDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(watchServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(watchServicesDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(createServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(updateServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(deleteServiceDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}
