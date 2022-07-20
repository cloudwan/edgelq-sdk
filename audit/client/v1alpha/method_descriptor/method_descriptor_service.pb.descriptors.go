// Code generated by protoc-gen-goten-client
// API: MethodDescriptorService
// DO NOT EDIT!!!

package method_descriptor_client

import (
	"google.golang.org/protobuf/proto"

	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	method_descriptor "github.com/cloudwan/edgelq-sdk/audit/resources/v1alpha/method_descriptor"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Message(nil)
	_ = gotenclient.MethodDescriptor(nil)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &method_descriptor.MethodDescriptor{}
)

var (
	descriptorsInitialized              bool
	methodDescriptorServiceDescriptor   *MethodDescriptorServiceDescriptor
	getMethodDescriptorDescriptor       *GetMethodDescriptorDescriptor
	batchGetMethodDescriptorsDescriptor *BatchGetMethodDescriptorsDescriptor
	listMethodDescriptorsDescriptor     *ListMethodDescriptorsDescriptor
	watchMethodDescriptorDescriptor     *WatchMethodDescriptorDescriptor
	watchMethodDescriptorsDescriptor    *WatchMethodDescriptorsDescriptor
	createMethodDescriptorDescriptor    *CreateMethodDescriptorDescriptor
	updateMethodDescriptorDescriptor    *UpdateMethodDescriptorDescriptor
)

type GetMethodDescriptorDescriptor struct{}

type GetMethodDescriptorDescriptorClientMsgHandle struct{}

type GetMethodDescriptorDescriptorServerMsgHandle struct{}

func (d *GetMethodDescriptorDescriptor) NewEmptyClientMsg() proto.Message {
	return &GetMethodDescriptorRequest{}
}

func (d *GetMethodDescriptorDescriptor) NewEmptyServerMsg() proto.Message {
	return &method_descriptor.MethodDescriptor{}
}

func (d *GetMethodDescriptorDescriptor) IsUnary() bool {
	return true
}

func (d *GetMethodDescriptorDescriptor) IsClientStream() bool {
	return false
}

func (d *GetMethodDescriptorDescriptor) IsServerStream() bool {
	return false
}

func (d *GetMethodDescriptorDescriptor) IsCollection() bool {
	return false
}

func (d *GetMethodDescriptorDescriptor) IsPlural() bool {
	return false
}

func (d *GetMethodDescriptorDescriptor) HasResource() bool {
	return true
}

func (d *GetMethodDescriptorDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *GetMethodDescriptorDescriptor) GetVerb() string {
	return "get"
}

func (d *GetMethodDescriptorDescriptor) GetMethodName() string {
	return "GetMethodDescriptor"
}

func (d *GetMethodDescriptorDescriptor) GetFullMethodName() string {
	return "/ntt.audit.v1alpha.MethodDescriptorService/GetMethodDescriptor"
}

func (d *GetMethodDescriptorDescriptor) GetProtoPkgName() string {
	return "ntt.audit.v1alpha"
}

func (d *GetMethodDescriptorDescriptor) GetApiName() string {
	return "MethodDescriptorService"
}

func (d *GetMethodDescriptorDescriptor) GetServiceDomain() string {
	return "audit.edgelq.com"
}

func (d *GetMethodDescriptorDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *GetMethodDescriptorDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return methodDescriptorServiceDescriptor
}

func (d *GetMethodDescriptorDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return method_descriptor.GetDescriptor()
}

func (d *GetMethodDescriptorDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetMethodDescriptorDescriptorClientMsgHandle{}
}

func (d *GetMethodDescriptorDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetMethodDescriptorDescriptorServerMsgHandle{}
}

func (h *GetMethodDescriptorDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*GetMethodDescriptorRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*GetMethodDescriptorRequest) *method_descriptor.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if ref := typedMsg.GetName(); ref != nil {
			return &ref.Name
		}
	}
	return (*method_descriptor.Name)(nil)
}

func (h *GetMethodDescriptorDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*GetMethodDescriptorRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*GetMethodDescriptorRequest) []*method_descriptor.Name
	})
	if ok {
		return method_descriptor.MethodDescriptorNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *GetMethodDescriptorDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *GetMethodDescriptorDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*method_descriptor.MethodDescriptor)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*method_descriptor.MethodDescriptor) *method_descriptor.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if name := typedMsg.GetName(); name != nil {
			return name
		}
	}
	return (*method_descriptor.Name)(nil)
}

func (h *GetMethodDescriptorDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*method_descriptor.MethodDescriptor)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*method_descriptor.MethodDescriptor) []*method_descriptor.Name
	})
	if ok {
		return method_descriptor.MethodDescriptorNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *GetMethodDescriptorDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetGetMethodDescriptorDescriptor() *GetMethodDescriptorDescriptor {
	return getMethodDescriptorDescriptor
}

type BatchGetMethodDescriptorsDescriptor struct{}

type BatchGetMethodDescriptorsDescriptorClientMsgHandle struct{}

type BatchGetMethodDescriptorsDescriptorServerMsgHandle struct{}

func (d *BatchGetMethodDescriptorsDescriptor) NewEmptyClientMsg() proto.Message {
	return &BatchGetMethodDescriptorsRequest{}
}

func (d *BatchGetMethodDescriptorsDescriptor) NewEmptyServerMsg() proto.Message {
	return &BatchGetMethodDescriptorsResponse{}
}

func (d *BatchGetMethodDescriptorsDescriptor) IsUnary() bool {
	return true
}

func (d *BatchGetMethodDescriptorsDescriptor) IsClientStream() bool {
	return false
}

func (d *BatchGetMethodDescriptorsDescriptor) IsServerStream() bool {
	return false
}

func (d *BatchGetMethodDescriptorsDescriptor) IsCollection() bool {
	return false
}

func (d *BatchGetMethodDescriptorsDescriptor) IsPlural() bool {
	return true
}

func (d *BatchGetMethodDescriptorsDescriptor) HasResource() bool {
	return true
}

func (d *BatchGetMethodDescriptorsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *BatchGetMethodDescriptorsDescriptor) GetVerb() string {
	return "batchGet"
}

func (d *BatchGetMethodDescriptorsDescriptor) GetMethodName() string {
	return "BatchGetMethodDescriptors"
}

func (d *BatchGetMethodDescriptorsDescriptor) GetFullMethodName() string {
	return "/ntt.audit.v1alpha.MethodDescriptorService/BatchGetMethodDescriptors"
}

func (d *BatchGetMethodDescriptorsDescriptor) GetProtoPkgName() string {
	return "ntt.audit.v1alpha"
}

func (d *BatchGetMethodDescriptorsDescriptor) GetApiName() string {
	return "MethodDescriptorService"
}

func (d *BatchGetMethodDescriptorsDescriptor) GetServiceDomain() string {
	return "audit.edgelq.com"
}

func (d *BatchGetMethodDescriptorsDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *BatchGetMethodDescriptorsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return methodDescriptorServiceDescriptor
}

func (d *BatchGetMethodDescriptorsDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return method_descriptor.GetDescriptor()
}

func (d *BatchGetMethodDescriptorsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &BatchGetMethodDescriptorsDescriptorClientMsgHandle{}
}

func (d *BatchGetMethodDescriptorsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &BatchGetMethodDescriptorsDescriptorServerMsgHandle{}
}

func (h *BatchGetMethodDescriptorsDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*BatchGetMethodDescriptorsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*BatchGetMethodDescriptorsRequest) *method_descriptor.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *BatchGetMethodDescriptorsDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*BatchGetMethodDescriptorsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*BatchGetMethodDescriptorsRequest) []*method_descriptor.Name
	})
	if ok {
		return method_descriptor.MethodDescriptorNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	{
		if refs := typedMsg.GetNames(); len(refs) > 0 {
			list := make(method_descriptor.MethodDescriptorNameList, 0, len(refs))
			for _, ref := range refs {
				list = append(list, &ref.Name)
			}
			return list
		}
	}
	return (method_descriptor.MethodDescriptorNameList)(nil)
}

func (h *BatchGetMethodDescriptorsDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *BatchGetMethodDescriptorsDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*BatchGetMethodDescriptorsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*BatchGetMethodDescriptorsResponse) *method_descriptor.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *BatchGetMethodDescriptorsDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*BatchGetMethodDescriptorsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*BatchGetMethodDescriptorsResponse) []*method_descriptor.Name
	})
	if ok {
		return method_descriptor.MethodDescriptorNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	{
		if resources := typedMsg.GetMethodDescriptors(); len(resources) > 0 {
			list := make(method_descriptor.MethodDescriptorNameList, 0, len(resources))
			for _, res := range resources {
				list = append(list, res.GetName())
			}
			return list
		}
	}
	return (method_descriptor.MethodDescriptorNameList)(nil)
}

func (h *BatchGetMethodDescriptorsDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetBatchGetMethodDescriptorsDescriptor() *BatchGetMethodDescriptorsDescriptor {
	return batchGetMethodDescriptorsDescriptor
}

type ListMethodDescriptorsDescriptor struct{}

type ListMethodDescriptorsDescriptorClientMsgHandle struct{}

type ListMethodDescriptorsDescriptorServerMsgHandle struct{}

func (d *ListMethodDescriptorsDescriptor) NewEmptyClientMsg() proto.Message {
	return &ListMethodDescriptorsRequest{}
}

func (d *ListMethodDescriptorsDescriptor) NewEmptyServerMsg() proto.Message {
	return &ListMethodDescriptorsResponse{}
}

func (d *ListMethodDescriptorsDescriptor) IsUnary() bool {
	return true
}

func (d *ListMethodDescriptorsDescriptor) IsClientStream() bool {
	return false
}

func (d *ListMethodDescriptorsDescriptor) IsServerStream() bool {
	return false
}

func (d *ListMethodDescriptorsDescriptor) IsCollection() bool {
	return true
}

func (d *ListMethodDescriptorsDescriptor) IsPlural() bool {
	return true
}

func (d *ListMethodDescriptorsDescriptor) HasResource() bool {
	return true
}

func (d *ListMethodDescriptorsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *ListMethodDescriptorsDescriptor) GetVerb() string {
	return "list"
}

func (d *ListMethodDescriptorsDescriptor) GetMethodName() string {
	return "ListMethodDescriptors"
}

func (d *ListMethodDescriptorsDescriptor) GetFullMethodName() string {
	return "/ntt.audit.v1alpha.MethodDescriptorService/ListMethodDescriptors"
}

func (d *ListMethodDescriptorsDescriptor) GetProtoPkgName() string {
	return "ntt.audit.v1alpha"
}

func (d *ListMethodDescriptorsDescriptor) GetApiName() string {
	return "MethodDescriptorService"
}

func (d *ListMethodDescriptorsDescriptor) GetServiceDomain() string {
	return "audit.edgelq.com"
}

func (d *ListMethodDescriptorsDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *ListMethodDescriptorsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return methodDescriptorServiceDescriptor
}

func (d *ListMethodDescriptorsDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return method_descriptor.GetDescriptor()
}

func (d *ListMethodDescriptorsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ListMethodDescriptorsDescriptorClientMsgHandle{}
}

func (d *ListMethodDescriptorsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ListMethodDescriptorsDescriptorServerMsgHandle{}
}

func (h *ListMethodDescriptorsDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListMethodDescriptorsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ListMethodDescriptorsRequest) *method_descriptor.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *ListMethodDescriptorsDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ListMethodDescriptorsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ListMethodDescriptorsRequest) []*method_descriptor.Name
	})
	if ok {
		return method_descriptor.MethodDescriptorNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *ListMethodDescriptorsDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *ListMethodDescriptorsDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListMethodDescriptorsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ListMethodDescriptorsResponse) *method_descriptor.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *ListMethodDescriptorsDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ListMethodDescriptorsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ListMethodDescriptorsResponse) []*method_descriptor.Name
	})
	if ok {
		return method_descriptor.MethodDescriptorNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	{
		if resources := typedMsg.GetMethodDescriptors(); len(resources) > 0 {
			list := make(method_descriptor.MethodDescriptorNameList, 0, len(resources))
			for _, res := range resources {
				list = append(list, res.GetName())
			}
			return list
		}
	}
	return (method_descriptor.MethodDescriptorNameList)(nil)
}

func (h *ListMethodDescriptorsDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetListMethodDescriptorsDescriptor() *ListMethodDescriptorsDescriptor {
	return listMethodDescriptorsDescriptor
}

type WatchMethodDescriptorDescriptor struct{}

type WatchMethodDescriptorDescriptorClientMsgHandle struct{}

type WatchMethodDescriptorDescriptorServerMsgHandle struct{}

func (d *WatchMethodDescriptorDescriptor) NewEmptyClientMsg() proto.Message {
	return &WatchMethodDescriptorRequest{}
}

func (d *WatchMethodDescriptorDescriptor) NewEmptyServerMsg() proto.Message {
	return &WatchMethodDescriptorResponse{}
}

func (d *WatchMethodDescriptorDescriptor) IsUnary() bool {
	return false
}

func (d *WatchMethodDescriptorDescriptor) IsClientStream() bool {
	return false
}

func (d *WatchMethodDescriptorDescriptor) IsServerStream() bool {
	return true
}

func (d *WatchMethodDescriptorDescriptor) IsCollection() bool {
	return false
}

func (d *WatchMethodDescriptorDescriptor) IsPlural() bool {
	return false
}

func (d *WatchMethodDescriptorDescriptor) HasResource() bool {
	return true
}

func (d *WatchMethodDescriptorDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *WatchMethodDescriptorDescriptor) GetVerb() string {
	return "watch"
}

func (d *WatchMethodDescriptorDescriptor) GetMethodName() string {
	return "WatchMethodDescriptor"
}

func (d *WatchMethodDescriptorDescriptor) GetFullMethodName() string {
	return "/ntt.audit.v1alpha.MethodDescriptorService/WatchMethodDescriptor"
}

func (d *WatchMethodDescriptorDescriptor) GetProtoPkgName() string {
	return "ntt.audit.v1alpha"
}

func (d *WatchMethodDescriptorDescriptor) GetApiName() string {
	return "MethodDescriptorService"
}

func (d *WatchMethodDescriptorDescriptor) GetServiceDomain() string {
	return "audit.edgelq.com"
}

func (d *WatchMethodDescriptorDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *WatchMethodDescriptorDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return methodDescriptorServiceDescriptor
}

func (d *WatchMethodDescriptorDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return method_descriptor.GetDescriptor()
}

func (d *WatchMethodDescriptorDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &WatchMethodDescriptorDescriptorClientMsgHandle{}
}

func (d *WatchMethodDescriptorDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &WatchMethodDescriptorDescriptorServerMsgHandle{}
}

func (h *WatchMethodDescriptorDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*WatchMethodDescriptorRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*WatchMethodDescriptorRequest) *method_descriptor.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if ref := typedMsg.GetName(); ref != nil {
			return &ref.Name
		}
	}
	return (*method_descriptor.Name)(nil)
}

func (h *WatchMethodDescriptorDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*WatchMethodDescriptorRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*WatchMethodDescriptorRequest) []*method_descriptor.Name
	})
	if ok {
		return method_descriptor.MethodDescriptorNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *WatchMethodDescriptorDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *WatchMethodDescriptorDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*WatchMethodDescriptorResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*WatchMethodDescriptorResponse) *method_descriptor.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if resChange := typedMsg.GetChange(); resChange != nil {
			switch tResChange := resChange.ChangeType.(type) {
			case *method_descriptor.MethodDescriptorChange_Added_:
				return tResChange.Added.GetMethodDescriptor().GetName()
			case *method_descriptor.MethodDescriptorChange_Modified_:
				return tResChange.Modified.GetName()
			case *method_descriptor.MethodDescriptorChange_Removed_:
				return tResChange.Removed.GetName()
			case *method_descriptor.MethodDescriptorChange_Current_:
				return tResChange.Current.GetMethodDescriptor().GetName()
			}
		}
	}
	return (*method_descriptor.Name)(nil)
}

func (h *WatchMethodDescriptorDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*WatchMethodDescriptorResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*WatchMethodDescriptorResponse) []*method_descriptor.Name
	})
	if ok {
		return method_descriptor.MethodDescriptorNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *WatchMethodDescriptorDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetWatchMethodDescriptorDescriptor() *WatchMethodDescriptorDescriptor {
	return watchMethodDescriptorDescriptor
}

type WatchMethodDescriptorsDescriptor struct{}

type WatchMethodDescriptorsDescriptorClientMsgHandle struct{}

type WatchMethodDescriptorsDescriptorServerMsgHandle struct{}

func (d *WatchMethodDescriptorsDescriptor) NewEmptyClientMsg() proto.Message {
	return &WatchMethodDescriptorsRequest{}
}

func (d *WatchMethodDescriptorsDescriptor) NewEmptyServerMsg() proto.Message {
	return &WatchMethodDescriptorsResponse{}
}

func (d *WatchMethodDescriptorsDescriptor) IsUnary() bool {
	return false
}

func (d *WatchMethodDescriptorsDescriptor) IsClientStream() bool {
	return false
}

func (d *WatchMethodDescriptorsDescriptor) IsServerStream() bool {
	return true
}

func (d *WatchMethodDescriptorsDescriptor) IsCollection() bool {
	return true
}

func (d *WatchMethodDescriptorsDescriptor) IsPlural() bool {
	return true
}

func (d *WatchMethodDescriptorsDescriptor) HasResource() bool {
	return true
}

func (d *WatchMethodDescriptorsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *WatchMethodDescriptorsDescriptor) GetVerb() string {
	return "watch"
}

func (d *WatchMethodDescriptorsDescriptor) GetMethodName() string {
	return "WatchMethodDescriptors"
}

func (d *WatchMethodDescriptorsDescriptor) GetFullMethodName() string {
	return "/ntt.audit.v1alpha.MethodDescriptorService/WatchMethodDescriptors"
}

func (d *WatchMethodDescriptorsDescriptor) GetProtoPkgName() string {
	return "ntt.audit.v1alpha"
}

func (d *WatchMethodDescriptorsDescriptor) GetApiName() string {
	return "MethodDescriptorService"
}

func (d *WatchMethodDescriptorsDescriptor) GetServiceDomain() string {
	return "audit.edgelq.com"
}

func (d *WatchMethodDescriptorsDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *WatchMethodDescriptorsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return methodDescriptorServiceDescriptor
}

func (d *WatchMethodDescriptorsDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return method_descriptor.GetDescriptor()
}

func (d *WatchMethodDescriptorsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &WatchMethodDescriptorsDescriptorClientMsgHandle{}
}

func (d *WatchMethodDescriptorsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &WatchMethodDescriptorsDescriptorServerMsgHandle{}
}

func (h *WatchMethodDescriptorsDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*WatchMethodDescriptorsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*WatchMethodDescriptorsRequest) *method_descriptor.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *WatchMethodDescriptorsDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*WatchMethodDescriptorsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*WatchMethodDescriptorsRequest) []*method_descriptor.Name
	})
	if ok {
		return method_descriptor.MethodDescriptorNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *WatchMethodDescriptorsDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *WatchMethodDescriptorsDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*WatchMethodDescriptorsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*WatchMethodDescriptorsResponse) *method_descriptor.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *WatchMethodDescriptorsDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*WatchMethodDescriptorsResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*WatchMethodDescriptorsResponse) []*method_descriptor.Name
	})
	if ok {
		return method_descriptor.MethodDescriptorNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	{
		if resChanges := typedMsg.GetMethodDescriptorChanges(); len(resChanges) > 0 {
			list := make(method_descriptor.MethodDescriptorNameList, 0, len(resChanges))
			for _, resChange := range resChanges {
				switch tResChange := resChange.ChangeType.(type) {
				case *method_descriptor.MethodDescriptorChange_Added_:
					list = append(list, tResChange.Added.GetMethodDescriptor().GetName())
				case *method_descriptor.MethodDescriptorChange_Modified_:
					list = append(list, tResChange.Modified.GetName())
				case *method_descriptor.MethodDescriptorChange_Removed_:
					list = append(list, tResChange.Removed.GetName())
				case *method_descriptor.MethodDescriptorChange_Current_:
					list = append(list, tResChange.Current.GetMethodDescriptor().GetName())
				}
			}
			return list
		}
	}
	return (method_descriptor.MethodDescriptorNameList)(nil)
}

func (h *WatchMethodDescriptorsDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetWatchMethodDescriptorsDescriptor() *WatchMethodDescriptorsDescriptor {
	return watchMethodDescriptorsDescriptor
}

type CreateMethodDescriptorDescriptor struct{}

type CreateMethodDescriptorDescriptorClientMsgHandle struct{}

type CreateMethodDescriptorDescriptorServerMsgHandle struct{}

func (d *CreateMethodDescriptorDescriptor) NewEmptyClientMsg() proto.Message {
	return &CreateMethodDescriptorRequest{}
}

func (d *CreateMethodDescriptorDescriptor) NewEmptyServerMsg() proto.Message {
	return &method_descriptor.MethodDescriptor{}
}

func (d *CreateMethodDescriptorDescriptor) IsUnary() bool {
	return true
}

func (d *CreateMethodDescriptorDescriptor) IsClientStream() bool {
	return false
}

func (d *CreateMethodDescriptorDescriptor) IsServerStream() bool {
	return false
}

func (d *CreateMethodDescriptorDescriptor) IsCollection() bool {
	return true
}

func (d *CreateMethodDescriptorDescriptor) IsPlural() bool {
	return false
}

func (d *CreateMethodDescriptorDescriptor) HasResource() bool {
	return true
}

func (d *CreateMethodDescriptorDescriptor) RequestHasResourceBody() bool {
	return true
}

func (d *CreateMethodDescriptorDescriptor) GetVerb() string {
	return "create"
}

func (d *CreateMethodDescriptorDescriptor) GetMethodName() string {
	return "CreateMethodDescriptor"
}

func (d *CreateMethodDescriptorDescriptor) GetFullMethodName() string {
	return "/ntt.audit.v1alpha.MethodDescriptorService/CreateMethodDescriptor"
}

func (d *CreateMethodDescriptorDescriptor) GetProtoPkgName() string {
	return "ntt.audit.v1alpha"
}

func (d *CreateMethodDescriptorDescriptor) GetApiName() string {
	return "MethodDescriptorService"
}

func (d *CreateMethodDescriptorDescriptor) GetServiceDomain() string {
	return "audit.edgelq.com"
}

func (d *CreateMethodDescriptorDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *CreateMethodDescriptorDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return methodDescriptorServiceDescriptor
}

func (d *CreateMethodDescriptorDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return method_descriptor.GetDescriptor()
}

func (d *CreateMethodDescriptorDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CreateMethodDescriptorDescriptorClientMsgHandle{}
}

func (d *CreateMethodDescriptorDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &CreateMethodDescriptorDescriptorServerMsgHandle{}
}

func (h *CreateMethodDescriptorDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*CreateMethodDescriptorRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*CreateMethodDescriptorRequest) *method_descriptor.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		res := typedMsg.GetMethodDescriptor()
		if name := res.GetName(); name != nil {
			return name
		}
	}
	return (*method_descriptor.Name)(nil)
}

func (h *CreateMethodDescriptorDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*CreateMethodDescriptorRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*CreateMethodDescriptorRequest) []*method_descriptor.Name
	})
	if ok {
		return method_descriptor.MethodDescriptorNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *CreateMethodDescriptorDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *CreateMethodDescriptorDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*method_descriptor.MethodDescriptor)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*method_descriptor.MethodDescriptor) *method_descriptor.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if name := typedMsg.GetName(); name != nil {
			return name
		}
	}
	return (*method_descriptor.Name)(nil)
}

func (h *CreateMethodDescriptorDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*method_descriptor.MethodDescriptor)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*method_descriptor.MethodDescriptor) []*method_descriptor.Name
	})
	if ok {
		return method_descriptor.MethodDescriptorNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *CreateMethodDescriptorDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetCreateMethodDescriptorDescriptor() *CreateMethodDescriptorDescriptor {
	return createMethodDescriptorDescriptor
}

type UpdateMethodDescriptorDescriptor struct{}

type UpdateMethodDescriptorDescriptorClientMsgHandle struct{}

type UpdateMethodDescriptorDescriptorServerMsgHandle struct{}

func (d *UpdateMethodDescriptorDescriptor) NewEmptyClientMsg() proto.Message {
	return &UpdateMethodDescriptorRequest{}
}

func (d *UpdateMethodDescriptorDescriptor) NewEmptyServerMsg() proto.Message {
	return &method_descriptor.MethodDescriptor{}
}

func (d *UpdateMethodDescriptorDescriptor) IsUnary() bool {
	return true
}

func (d *UpdateMethodDescriptorDescriptor) IsClientStream() bool {
	return false
}

func (d *UpdateMethodDescriptorDescriptor) IsServerStream() bool {
	return false
}

func (d *UpdateMethodDescriptorDescriptor) IsCollection() bool {
	return false
}

func (d *UpdateMethodDescriptorDescriptor) IsPlural() bool {
	return false
}

func (d *UpdateMethodDescriptorDescriptor) HasResource() bool {
	return true
}

func (d *UpdateMethodDescriptorDescriptor) RequestHasResourceBody() bool {
	return true
}

func (d *UpdateMethodDescriptorDescriptor) GetVerb() string {
	return "update"
}

func (d *UpdateMethodDescriptorDescriptor) GetMethodName() string {
	return "UpdateMethodDescriptor"
}

func (d *UpdateMethodDescriptorDescriptor) GetFullMethodName() string {
	return "/ntt.audit.v1alpha.MethodDescriptorService/UpdateMethodDescriptor"
}

func (d *UpdateMethodDescriptorDescriptor) GetProtoPkgName() string {
	return "ntt.audit.v1alpha"
}

func (d *UpdateMethodDescriptorDescriptor) GetApiName() string {
	return "MethodDescriptorService"
}

func (d *UpdateMethodDescriptorDescriptor) GetServiceDomain() string {
	return "audit.edgelq.com"
}

func (d *UpdateMethodDescriptorDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *UpdateMethodDescriptorDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return methodDescriptorServiceDescriptor
}

func (d *UpdateMethodDescriptorDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return method_descriptor.GetDescriptor()
}

func (d *UpdateMethodDescriptorDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &UpdateMethodDescriptorDescriptorClientMsgHandle{}
}

func (d *UpdateMethodDescriptorDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &UpdateMethodDescriptorDescriptorServerMsgHandle{}
}

func (h *UpdateMethodDescriptorDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*UpdateMethodDescriptorRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*UpdateMethodDescriptorRequest) *method_descriptor.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		res := typedMsg.GetMethodDescriptor()
		if name := res.GetName(); name != nil {
			return name
		}
	}
	return (*method_descriptor.Name)(nil)
}

func (h *UpdateMethodDescriptorDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*UpdateMethodDescriptorRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*UpdateMethodDescriptorRequest) []*method_descriptor.Name
	})
	if ok {
		return method_descriptor.MethodDescriptorNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *UpdateMethodDescriptorDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *UpdateMethodDescriptorDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*method_descriptor.MethodDescriptor)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*method_descriptor.MethodDescriptor) *method_descriptor.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if name := typedMsg.GetName(); name != nil {
			return name
		}
	}
	return (*method_descriptor.Name)(nil)
}

func (h *UpdateMethodDescriptorDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*method_descriptor.MethodDescriptor)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*method_descriptor.MethodDescriptor) []*method_descriptor.Name
	})
	if ok {
		return method_descriptor.MethodDescriptorNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *UpdateMethodDescriptorDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetUpdateMethodDescriptorDescriptor() *UpdateMethodDescriptorDescriptor {
	return updateMethodDescriptorDescriptor
}

type MethodDescriptorServiceDescriptor struct{}

func (d *MethodDescriptorServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		getMethodDescriptorDescriptor,
		batchGetMethodDescriptorsDescriptor,
		listMethodDescriptorsDescriptor,
		watchMethodDescriptorDescriptor,
		watchMethodDescriptorsDescriptor,
		createMethodDescriptorDescriptor,
		updateMethodDescriptorDescriptor,
	}
}

func (d *MethodDescriptorServiceDescriptor) GetFullAPIName() string {
	return "/ntt.audit.v1alpha.MethodDescriptorService"
}

func (d *MethodDescriptorServiceDescriptor) GetProtoPkgName() string {
	return "ntt.audit.v1alpha"
}

func (d *MethodDescriptorServiceDescriptor) GetApiName() string {
	return "MethodDescriptorService"
}

func (d *MethodDescriptorServiceDescriptor) GetServiceDomain() string {
	return "audit.edgelq.com"
}

func (d *MethodDescriptorServiceDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func GetMethodDescriptorServiceDescriptor() *MethodDescriptorServiceDescriptor {
	return methodDescriptorServiceDescriptor
}

func initDescriptors() {
	methodDescriptorServiceDescriptor = &MethodDescriptorServiceDescriptor{}
	getMethodDescriptorDescriptor = &GetMethodDescriptorDescriptor{}
	batchGetMethodDescriptorsDescriptor = &BatchGetMethodDescriptorsDescriptor{}
	listMethodDescriptorsDescriptor = &ListMethodDescriptorsDescriptor{}
	watchMethodDescriptorDescriptor = &WatchMethodDescriptorDescriptor{}
	watchMethodDescriptorsDescriptor = &WatchMethodDescriptorsDescriptor{}
	createMethodDescriptorDescriptor = &CreateMethodDescriptorDescriptor{}
	updateMethodDescriptorDescriptor = &UpdateMethodDescriptorDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(methodDescriptorServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(getMethodDescriptorDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(batchGetMethodDescriptorsDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(listMethodDescriptorsDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(watchMethodDescriptorDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(watchMethodDescriptorsDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(createMethodDescriptorDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(updateMethodDescriptorDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}