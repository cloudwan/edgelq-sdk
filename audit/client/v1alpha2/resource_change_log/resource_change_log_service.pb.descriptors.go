// Code generated by protoc-gen-goten-client
// API: ResourceChangeLogService
// DO NOT EDIT!!!

package resource_change_log_client

import (
	"google.golang.org/protobuf/proto"

	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	resource_change_log "github.com/cloudwan/edgelq-sdk/audit/resources/v1alpha2/resource_change_log"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = new(proto.Message)
	_ = new(gotenclient.MethodDescriptor)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &resource_change_log.ResourceChangeLog{}
)

var (
	descriptorsInitialized                         bool
	resourceChangeLogServiceDescriptor             *ResourceChangeLogServiceDescriptor
	listResourceChangeLogsDescriptor               *ListResourceChangeLogsDescriptor
	createPreCommittedResourceChangeLogsDescriptor *CreatePreCommittedResourceChangeLogsDescriptor
	setResourceChangeLogsCommitStateDescriptor     *SetResourceChangeLogsCommitStateDescriptor
)

type ListResourceChangeLogsDescriptor struct{}

type ListResourceChangeLogsDescriptorClientMsgHandle struct{}

type ListResourceChangeLogsDescriptorServerMsgHandle struct{}

func (d *ListResourceChangeLogsDescriptor) NewEmptyClientMsg() proto.Message {
	return &ListResourceChangeLogsRequest{}
}

func (d *ListResourceChangeLogsDescriptor) NewEmptyServerMsg() proto.Message {
	return &ListResourceChangeLogsResponse{}
}

func (d *ListResourceChangeLogsDescriptor) IsUnary() bool {
	return true
}

func (d *ListResourceChangeLogsDescriptor) IsClientStream() bool {
	return false
}

func (d *ListResourceChangeLogsDescriptor) IsServerStream() bool {
	return false
}

func (d *ListResourceChangeLogsDescriptor) IsCollection() bool {
	return true
}

func (d *ListResourceChangeLogsDescriptor) IsPlural() bool {
	return true
}

func (d *ListResourceChangeLogsDescriptor) HasResource() bool {
	return true
}

func (d *ListResourceChangeLogsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *ListResourceChangeLogsDescriptor) GetVerb() string {
	return "query"
}

func (d *ListResourceChangeLogsDescriptor) GetMethodName() string {
	return "ListResourceChangeLogs"
}

func (d *ListResourceChangeLogsDescriptor) GetFullMethodName() string {
	return "/ntt.audit.v1alpha2.ResourceChangeLogService/ListResourceChangeLogs"
}

func (d *ListResourceChangeLogsDescriptor) GetProtoPkgName() string {
	return "ntt.audit.v1alpha2"
}

func (d *ListResourceChangeLogsDescriptor) GetApiName() string {
	return "ResourceChangeLogService"
}

func (d *ListResourceChangeLogsDescriptor) GetServiceDomain() string {
	return "audit.edgelq.com"
}

func (d *ListResourceChangeLogsDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *ListResourceChangeLogsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return resourceChangeLogServiceDescriptor
}

func (d *ListResourceChangeLogsDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return resource_change_log.GetDescriptor()
}

func (d *ListResourceChangeLogsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ListResourceChangeLogsDescriptorClientMsgHandle{}
}

func (d *ListResourceChangeLogsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ListResourceChangeLogsDescriptorServerMsgHandle{}
}

func (h *ListResourceChangeLogsDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListResourceChangeLogsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ListResourceChangeLogsRequest) *resource_change_log.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *ListResourceChangeLogsDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ListResourceChangeLogsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ListResourceChangeLogsRequest) []*resource_change_log.Name
	})
	if ok {
		return resource_change_log.ResourceChangeLogNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *ListResourceChangeLogsDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListResourceChangeLogsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*ListResourceChangeLogsRequest) *resource_change_log.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func (h *ListResourceChangeLogsDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListResourceChangeLogsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ListResourceChangeLogsResponse) *resource_change_log.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *ListResourceChangeLogsDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ListResourceChangeLogsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ListResourceChangeLogsResponse) []*resource_change_log.Name
	})
	if ok {
		return resource_change_log.ResourceChangeLogNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *ListResourceChangeLogsDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListResourceChangeLogsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*ListResourceChangeLogsResponse) *resource_change_log.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func GetListResourceChangeLogsDescriptor() *ListResourceChangeLogsDescriptor {
	return listResourceChangeLogsDescriptor
}

type CreatePreCommittedResourceChangeLogsDescriptor struct{}

type CreatePreCommittedResourceChangeLogsDescriptorClientMsgHandle struct{}

type CreatePreCommittedResourceChangeLogsDescriptorServerMsgHandle struct{}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) NewEmptyClientMsg() proto.Message {
	return &CreatePreCommittedResourceChangeLogsRequest{}
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) NewEmptyServerMsg() proto.Message {
	return &CreatePreCommittedResourceChangeLogsResponse{}
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) IsUnary() bool {
	return true
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) IsClientStream() bool {
	return false
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) IsServerStream() bool {
	return false
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) IsCollection() bool {
	return true
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) IsPlural() bool {
	return true
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) HasResource() bool {
	return true
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) GetVerb() string {
	return "batchCreate"
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) GetMethodName() string {
	return "CreatePreCommittedResourceChangeLogs"
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) GetFullMethodName() string {
	return "/ntt.audit.v1alpha2.ResourceChangeLogService/CreatePreCommittedResourceChangeLogs"
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) GetProtoPkgName() string {
	return "ntt.audit.v1alpha2"
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) GetApiName() string {
	return "ResourceChangeLogService"
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) GetServiceDomain() string {
	return "audit.edgelq.com"
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return resourceChangeLogServiceDescriptor
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return resource_change_log.GetDescriptor()
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CreatePreCommittedResourceChangeLogsDescriptorClientMsgHandle{}
}

func (d *CreatePreCommittedResourceChangeLogsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CreatePreCommittedResourceChangeLogsDescriptorServerMsgHandle{}
}

func (h *CreatePreCommittedResourceChangeLogsDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CreatePreCommittedResourceChangeLogsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*CreatePreCommittedResourceChangeLogsRequest) *resource_change_log.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *CreatePreCommittedResourceChangeLogsDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*CreatePreCommittedResourceChangeLogsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*CreatePreCommittedResourceChangeLogsRequest) []*resource_change_log.Name
	})
	if ok {
		return resource_change_log.ResourceChangeLogNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *CreatePreCommittedResourceChangeLogsDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CreatePreCommittedResourceChangeLogsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*CreatePreCommittedResourceChangeLogsRequest) *resource_change_log.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func (h *CreatePreCommittedResourceChangeLogsDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CreatePreCommittedResourceChangeLogsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*CreatePreCommittedResourceChangeLogsResponse) *resource_change_log.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *CreatePreCommittedResourceChangeLogsDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*CreatePreCommittedResourceChangeLogsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*CreatePreCommittedResourceChangeLogsResponse) []*resource_change_log.Name
	})
	if ok {
		return resource_change_log.ResourceChangeLogNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *CreatePreCommittedResourceChangeLogsDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CreatePreCommittedResourceChangeLogsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*CreatePreCommittedResourceChangeLogsResponse) *resource_change_log.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func GetCreatePreCommittedResourceChangeLogsDescriptor() *CreatePreCommittedResourceChangeLogsDescriptor {
	return createPreCommittedResourceChangeLogsDescriptor
}

type SetResourceChangeLogsCommitStateDescriptor struct{}

type SetResourceChangeLogsCommitStateDescriptorClientMsgHandle struct{}

type SetResourceChangeLogsCommitStateDescriptorServerMsgHandle struct{}

func (d *SetResourceChangeLogsCommitStateDescriptor) NewEmptyClientMsg() proto.Message {
	return &SetResourceChangeLogsCommitStateRequest{}
}

func (d *SetResourceChangeLogsCommitStateDescriptor) NewEmptyServerMsg() proto.Message {
	return &SetResourceChangeLogsCommitStateResponse{}
}

func (d *SetResourceChangeLogsCommitStateDescriptor) IsUnary() bool {
	return true
}

func (d *SetResourceChangeLogsCommitStateDescriptor) IsClientStream() bool {
	return false
}

func (d *SetResourceChangeLogsCommitStateDescriptor) IsServerStream() bool {
	return false
}

func (d *SetResourceChangeLogsCommitStateDescriptor) IsCollection() bool {
	return true
}

func (d *SetResourceChangeLogsCommitStateDescriptor) IsPlural() bool {
	return true
}

func (d *SetResourceChangeLogsCommitStateDescriptor) HasResource() bool {
	return true
}

func (d *SetResourceChangeLogsCommitStateDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *SetResourceChangeLogsCommitStateDescriptor) GetVerb() string {
	return "batchUpdate"
}

func (d *SetResourceChangeLogsCommitStateDescriptor) GetMethodName() string {
	return "SetResourceChangeLogsCommitState"
}

func (d *SetResourceChangeLogsCommitStateDescriptor) GetFullMethodName() string {
	return "/ntt.audit.v1alpha2.ResourceChangeLogService/SetResourceChangeLogsCommitState"
}

func (d *SetResourceChangeLogsCommitStateDescriptor) GetProtoPkgName() string {
	return "ntt.audit.v1alpha2"
}

func (d *SetResourceChangeLogsCommitStateDescriptor) GetApiName() string {
	return "ResourceChangeLogService"
}

func (d *SetResourceChangeLogsCommitStateDescriptor) GetServiceDomain() string {
	return "audit.edgelq.com"
}

func (d *SetResourceChangeLogsCommitStateDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *SetResourceChangeLogsCommitStateDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return resourceChangeLogServiceDescriptor
}

func (d *SetResourceChangeLogsCommitStateDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return resource_change_log.GetDescriptor()
}

func (d *SetResourceChangeLogsCommitStateDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &SetResourceChangeLogsCommitStateDescriptorClientMsgHandle{}
}

func (d *SetResourceChangeLogsCommitStateDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &SetResourceChangeLogsCommitStateDescriptorServerMsgHandle{}
}

func (h *SetResourceChangeLogsCommitStateDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*SetResourceChangeLogsCommitStateRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*SetResourceChangeLogsCommitStateRequest) *resource_change_log.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *SetResourceChangeLogsCommitStateDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*SetResourceChangeLogsCommitStateRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*SetResourceChangeLogsCommitStateRequest) []*resource_change_log.Name
	})
	if ok {
		return resource_change_log.ResourceChangeLogNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *SetResourceChangeLogsCommitStateDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*SetResourceChangeLogsCommitStateRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*SetResourceChangeLogsCommitStateRequest) *resource_change_log.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func (h *SetResourceChangeLogsCommitStateDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*SetResourceChangeLogsCommitStateResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*SetResourceChangeLogsCommitStateResponse) *resource_change_log.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *SetResourceChangeLogsCommitStateDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*SetResourceChangeLogsCommitStateResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*SetResourceChangeLogsCommitStateResponse) []*resource_change_log.Name
	})
	if ok {
		return resource_change_log.ResourceChangeLogNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *SetResourceChangeLogsCommitStateDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*SetResourceChangeLogsCommitStateResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*SetResourceChangeLogsCommitStateResponse) *resource_change_log.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func GetSetResourceChangeLogsCommitStateDescriptor() *SetResourceChangeLogsCommitStateDescriptor {
	return setResourceChangeLogsCommitStateDescriptor
}

type ResourceChangeLogServiceDescriptor struct{}

func (d *ResourceChangeLogServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		listResourceChangeLogsDescriptor,
		createPreCommittedResourceChangeLogsDescriptor,
		setResourceChangeLogsCommitStateDescriptor,
	}
}

func (d *ResourceChangeLogServiceDescriptor) GetFullAPIName() string {
	return "/ntt.audit.v1alpha2.ResourceChangeLogService"
}

func (d *ResourceChangeLogServiceDescriptor) GetProtoPkgName() string {
	return "ntt.audit.v1alpha2"
}

func (d *ResourceChangeLogServiceDescriptor) GetApiName() string {
	return "ResourceChangeLogService"
}

func (d *ResourceChangeLogServiceDescriptor) GetServiceDomain() string {
	return "audit.edgelq.com"
}

func (d *ResourceChangeLogServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func GetResourceChangeLogServiceDescriptor() *ResourceChangeLogServiceDescriptor {
	return resourceChangeLogServiceDescriptor
}

func initDescriptors() {
	resourceChangeLogServiceDescriptor = &ResourceChangeLogServiceDescriptor{}
	listResourceChangeLogsDescriptor = &ListResourceChangeLogsDescriptor{}
	createPreCommittedResourceChangeLogsDescriptor = &CreatePreCommittedResourceChangeLogsDescriptor{}
	setResourceChangeLogsCommitStateDescriptor = &SetResourceChangeLogsCommitStateDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(resourceChangeLogServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(listResourceChangeLogsDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(createPreCommittedResourceChangeLogsDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(setResourceChangeLogsCommitStateDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}
