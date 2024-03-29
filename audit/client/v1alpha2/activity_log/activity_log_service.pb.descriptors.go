// Code generated by protoc-gen-goten-client
// API: ActivityLogService
// DO NOT EDIT!!!

package activity_log_client

import (
	"google.golang.org/protobuf/proto"

	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	activity_log "github.com/cloudwan/edgelq-sdk/audit/resources/v1alpha2/activity_log"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = new(proto.Message)
	_ = new(gotenclient.MethodDescriptor)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &activity_log.ActivityLog{}
)

var (
	descriptorsInitialized       bool
	activityLogServiceDescriptor *ActivityLogServiceDescriptor
	listActivityLogsDescriptor   *ListActivityLogsDescriptor
	createActivityLogsDescriptor *CreateActivityLogsDescriptor
)

type ListActivityLogsDescriptor struct{}

type ListActivityLogsDescriptorClientMsgHandle struct{}

type ListActivityLogsDescriptorServerMsgHandle struct{}

func (d *ListActivityLogsDescriptor) NewEmptyClientMsg() proto.Message {
	return &ListActivityLogsRequest{}
}

func (d *ListActivityLogsDescriptor) NewEmptyServerMsg() proto.Message {
	return &ListActivityLogsResponse{}
}

func (d *ListActivityLogsDescriptor) IsUnary() bool {
	return true
}

func (d *ListActivityLogsDescriptor) IsClientStream() bool {
	return false
}

func (d *ListActivityLogsDescriptor) IsServerStream() bool {
	return false
}

func (d *ListActivityLogsDescriptor) IsCollection() bool {
	return true
}

func (d *ListActivityLogsDescriptor) IsPlural() bool {
	return true
}

func (d *ListActivityLogsDescriptor) HasResource() bool {
	return true
}

func (d *ListActivityLogsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *ListActivityLogsDescriptor) GetVerb() string {
	return "query"
}

func (d *ListActivityLogsDescriptor) GetMethodName() string {
	return "ListActivityLogs"
}

func (d *ListActivityLogsDescriptor) GetFullMethodName() string {
	return "/ntt.audit.v1alpha2.ActivityLogService/ListActivityLogs"
}

func (d *ListActivityLogsDescriptor) GetProtoPkgName() string {
	return "ntt.audit.v1alpha2"
}

func (d *ListActivityLogsDescriptor) GetApiName() string {
	return "ActivityLogService"
}

func (d *ListActivityLogsDescriptor) GetServiceDomain() string {
	return "audit.edgelq.com"
}

func (d *ListActivityLogsDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *ListActivityLogsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return activityLogServiceDescriptor
}

func (d *ListActivityLogsDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return activity_log.GetDescriptor()
}

func (d *ListActivityLogsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ListActivityLogsDescriptorClientMsgHandle{}
}

func (d *ListActivityLogsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ListActivityLogsDescriptorServerMsgHandle{}
}

func (h *ListActivityLogsDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListActivityLogsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ListActivityLogsRequest) *activity_log.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *ListActivityLogsDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ListActivityLogsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ListActivityLogsRequest) []*activity_log.Name
	})
	if ok {
		return activity_log.ActivityLogNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *ListActivityLogsDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListActivityLogsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*ListActivityLogsRequest) *activity_log.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func (h *ListActivityLogsDescriptorClientMsgHandle) ExtractResourceBody(msg proto.Message) gotenresource.Resource {
	typedMsg := msg.(*ListActivityLogsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBody(*ListActivityLogsRequest) *activity_log.ActivityLog
	})
	if ok {
		return override.OverrideExtractResourceBody(typedMsg)
	}
	return nil
}

func (h *ListActivityLogsDescriptorClientMsgHandle) ExtractResourceBodies(msg proto.Message) gotenresource.ResourceList {
	typedMsg := msg.(*ListActivityLogsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBodies(*ListActivityLogsRequest) []*activity_log.ActivityLog
	})
	if ok {
		return activity_log.ActivityLogList(override.OverrideExtractResourceBodies(typedMsg))
	}
	return nil
}

func (h *ListActivityLogsDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListActivityLogsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ListActivityLogsResponse) *activity_log.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *ListActivityLogsDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ListActivityLogsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ListActivityLogsResponse) []*activity_log.Name
	})
	if ok {
		return activity_log.ActivityLogNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *ListActivityLogsDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListActivityLogsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*ListActivityLogsResponse) *activity_log.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func (h *ListActivityLogsDescriptorServerMsgHandle) ExtractResourceBody(msg proto.Message) gotenresource.Resource {
	typedMsg := msg.(*ListActivityLogsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBody(*ListActivityLogsResponse) *activity_log.ActivityLog
	})
	if ok {
		return override.OverrideExtractResourceBody(typedMsg)
	}
	return nil
}

func (h *ListActivityLogsDescriptorServerMsgHandle) ExtractResourceBodies(msg proto.Message) gotenresource.ResourceList {
	typedMsg := msg.(*ListActivityLogsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBodies(*ListActivityLogsResponse) []*activity_log.ActivityLog
	})
	if ok {
		return activity_log.ActivityLogList(override.OverrideExtractResourceBodies(typedMsg))
	}
	return nil
}

func GetListActivityLogsDescriptor() *ListActivityLogsDescriptor {
	return listActivityLogsDescriptor
}

type CreateActivityLogsDescriptor struct{}

type CreateActivityLogsDescriptorClientMsgHandle struct{}

type CreateActivityLogsDescriptorServerMsgHandle struct{}

func (d *CreateActivityLogsDescriptor) NewEmptyClientMsg() proto.Message {
	return &CreateActivityLogsRequest{}
}

func (d *CreateActivityLogsDescriptor) NewEmptyServerMsg() proto.Message {
	return &CreateActivityLogsResponse{}
}

func (d *CreateActivityLogsDescriptor) IsUnary() bool {
	return true
}

func (d *CreateActivityLogsDescriptor) IsClientStream() bool {
	return false
}

func (d *CreateActivityLogsDescriptor) IsServerStream() bool {
	return false
}

func (d *CreateActivityLogsDescriptor) IsCollection() bool {
	return true
}

func (d *CreateActivityLogsDescriptor) IsPlural() bool {
	return true
}

func (d *CreateActivityLogsDescriptor) HasResource() bool {
	return true
}

func (d *CreateActivityLogsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *CreateActivityLogsDescriptor) GetVerb() string {
	return "batchCreate"
}

func (d *CreateActivityLogsDescriptor) GetMethodName() string {
	return "CreateActivityLogs"
}

func (d *CreateActivityLogsDescriptor) GetFullMethodName() string {
	return "/ntt.audit.v1alpha2.ActivityLogService/CreateActivityLogs"
}

func (d *CreateActivityLogsDescriptor) GetProtoPkgName() string {
	return "ntt.audit.v1alpha2"
}

func (d *CreateActivityLogsDescriptor) GetApiName() string {
	return "ActivityLogService"
}

func (d *CreateActivityLogsDescriptor) GetServiceDomain() string {
	return "audit.edgelq.com"
}

func (d *CreateActivityLogsDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *CreateActivityLogsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return activityLogServiceDescriptor
}

func (d *CreateActivityLogsDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return activity_log.GetDescriptor()
}

func (d *CreateActivityLogsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CreateActivityLogsDescriptorClientMsgHandle{}
}

func (d *CreateActivityLogsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CreateActivityLogsDescriptorServerMsgHandle{}
}

func (h *CreateActivityLogsDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CreateActivityLogsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*CreateActivityLogsRequest) *activity_log.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *CreateActivityLogsDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*CreateActivityLogsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*CreateActivityLogsRequest) []*activity_log.Name
	})
	if ok {
		return activity_log.ActivityLogNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *CreateActivityLogsDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CreateActivityLogsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*CreateActivityLogsRequest) *activity_log.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func (h *CreateActivityLogsDescriptorClientMsgHandle) ExtractResourceBody(msg proto.Message) gotenresource.Resource {
	typedMsg := msg.(*CreateActivityLogsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBody(*CreateActivityLogsRequest) *activity_log.ActivityLog
	})
	if ok {
		return override.OverrideExtractResourceBody(typedMsg)
	}
	return nil
}

func (h *CreateActivityLogsDescriptorClientMsgHandle) ExtractResourceBodies(msg proto.Message) gotenresource.ResourceList {
	typedMsg := msg.(*CreateActivityLogsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBodies(*CreateActivityLogsRequest) []*activity_log.ActivityLog
	})
	if ok {
		return activity_log.ActivityLogList(override.OverrideExtractResourceBodies(typedMsg))
	}
	return nil
}

func (h *CreateActivityLogsDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CreateActivityLogsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*CreateActivityLogsResponse) *activity_log.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *CreateActivityLogsDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*CreateActivityLogsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*CreateActivityLogsResponse) []*activity_log.Name
	})
	if ok {
		return activity_log.ActivityLogNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *CreateActivityLogsDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CreateActivityLogsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*CreateActivityLogsResponse) *activity_log.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func (h *CreateActivityLogsDescriptorServerMsgHandle) ExtractResourceBody(msg proto.Message) gotenresource.Resource {
	typedMsg := msg.(*CreateActivityLogsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBody(*CreateActivityLogsResponse) *activity_log.ActivityLog
	})
	if ok {
		return override.OverrideExtractResourceBody(typedMsg)
	}
	return nil
}

func (h *CreateActivityLogsDescriptorServerMsgHandle) ExtractResourceBodies(msg proto.Message) gotenresource.ResourceList {
	typedMsg := msg.(*CreateActivityLogsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBodies(*CreateActivityLogsResponse) []*activity_log.ActivityLog
	})
	if ok {
		return activity_log.ActivityLogList(override.OverrideExtractResourceBodies(typedMsg))
	}
	return nil
}

func GetCreateActivityLogsDescriptor() *CreateActivityLogsDescriptor {
	return createActivityLogsDescriptor
}

type ActivityLogServiceDescriptor struct{}

func (d *ActivityLogServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		listActivityLogsDescriptor,
		createActivityLogsDescriptor,
	}
}

func (d *ActivityLogServiceDescriptor) GetFullAPIName() string {
	return "/ntt.audit.v1alpha2.ActivityLogService"
}

func (d *ActivityLogServiceDescriptor) GetProtoPkgName() string {
	return "ntt.audit.v1alpha2"
}

func (d *ActivityLogServiceDescriptor) GetApiName() string {
	return "ActivityLogService"
}

func (d *ActivityLogServiceDescriptor) GetServiceDomain() string {
	return "audit.edgelq.com"
}

func (d *ActivityLogServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func GetActivityLogServiceDescriptor() *ActivityLogServiceDescriptor {
	return activityLogServiceDescriptor
}

func initDescriptors() {
	activityLogServiceDescriptor = &ActivityLogServiceDescriptor{}
	listActivityLogsDescriptor = &ListActivityLogsDescriptor{}
	createActivityLogsDescriptor = &CreateActivityLogsDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(activityLogServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(listActivityLogsDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(createActivityLogsDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}
