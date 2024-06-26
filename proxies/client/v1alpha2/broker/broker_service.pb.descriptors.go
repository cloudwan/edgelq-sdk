// Code generated by protoc-gen-goten-client
// API: BrokerService
// DO NOT EDIT!!!

package broker_client

import (
	"google.golang.org/protobuf/proto"

	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/proxies/resources/v1alpha2/project"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = new(proto.Message)
	_ = new(gotenclient.MethodDescriptor)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
)

var (
	descriptorsInitialized  bool
	brokerServiceDescriptor *BrokerServiceDescriptor
	connectDescriptor       *ConnectDescriptor
	listenDescriptor        *ListenDescriptor
	acceptDescriptor        *AcceptDescriptor
)

type ConnectDescriptor struct{}

type ConnectDescriptorClientMsgHandle struct{}

type ConnectDescriptorServerMsgHandle struct{}

func (d *ConnectDescriptor) NewEmptyClientMsg() proto.Message {
	return &ConnectRequest{}
}

func (d *ConnectDescriptor) NewEmptyServerMsg() proto.Message {
	return &ConnectResponse{}
}

func (d *ConnectDescriptor) IsUnary() bool {
	return false
}

func (d *ConnectDescriptor) IsClientStream() bool {
	return true
}

func (d *ConnectDescriptor) IsServerStream() bool {
	return true
}

func (d *ConnectDescriptor) IsCollection() bool {
	return false
}

func (d *ConnectDescriptor) IsPlural() bool {
	return false
}

func (d *ConnectDescriptor) HasResource() bool {
	return true
}

func (d *ConnectDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *ConnectDescriptor) GetVerb() string {
	return "connect"
}

func (d *ConnectDescriptor) GetMethodName() string {
	return "Connect"
}

func (d *ConnectDescriptor) GetFullMethodName() string {
	return "/ntt.proxies.v1alpha2.BrokerService/Connect"
}

func (d *ConnectDescriptor) GetProtoPkgName() string {
	return "ntt.proxies.v1alpha2"
}

func (d *ConnectDescriptor) GetApiName() string {
	return "BrokerService"
}

func (d *ConnectDescriptor) GetServiceDomain() string {
	return "proxies.edgelq.com"
}

func (d *ConnectDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *ConnectDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return brokerServiceDescriptor
}

func (d *ConnectDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return project.GetDescriptor()
}

func (d *ConnectDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ConnectDescriptorClientMsgHandle{}
}

func (d *ConnectDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ConnectDescriptorServerMsgHandle{}
}

func (h *ConnectDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ConnectRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ConnectRequest) *project.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if name := typedMsg.GetOpenRequest().GetProject(); name != nil {
			return name
		}
	}
	{
		if name := typedMsg.GetResumeRequest().GetProject(); name != nil {
			return name
		}
	}
	return (*project.Name)(nil)
}

func (h *ConnectDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ConnectRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ConnectRequest) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *ConnectDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *ConnectDescriptorClientMsgHandle) ExtractResourceBody(msg proto.Message) gotenresource.Resource {
	typedMsg := msg.(*ConnectRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBody(*ConnectRequest) *project.Project
	})
	if ok {
		return override.OverrideExtractResourceBody(typedMsg)
	}
	return nil
}

func (h *ConnectDescriptorClientMsgHandle) ExtractResourceBodies(msg proto.Message) gotenresource.ResourceList {
	typedMsg := msg.(*ConnectRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBodies(*ConnectRequest) []*project.Project
	})
	if ok {
		return project.ProjectList(override.OverrideExtractResourceBodies(typedMsg))
	}
	return nil
}

func (h *ConnectDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ConnectResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ConnectResponse) *project.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *ConnectDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ConnectResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ConnectResponse) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *ConnectDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *ConnectDescriptorServerMsgHandle) ExtractResourceBody(msg proto.Message) gotenresource.Resource {
	typedMsg := msg.(*ConnectResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBody(*ConnectResponse) *project.Project
	})
	if ok {
		return override.OverrideExtractResourceBody(typedMsg)
	}
	return nil
}

func (h *ConnectDescriptorServerMsgHandle) ExtractResourceBodies(msg proto.Message) gotenresource.ResourceList {
	typedMsg := msg.(*ConnectResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBodies(*ConnectResponse) []*project.Project
	})
	if ok {
		return project.ProjectList(override.OverrideExtractResourceBodies(typedMsg))
	}
	return nil
}

func GetConnectDescriptor() *ConnectDescriptor {
	return connectDescriptor
}

type ListenDescriptor struct{}

type ListenDescriptorClientMsgHandle struct{}

type ListenDescriptorServerMsgHandle struct{}

func (d *ListenDescriptor) NewEmptyClientMsg() proto.Message {
	return &ListenRequest{}
}

func (d *ListenDescriptor) NewEmptyServerMsg() proto.Message {
	return &ListenResponse{}
}

func (d *ListenDescriptor) IsUnary() bool {
	return false
}

func (d *ListenDescriptor) IsClientStream() bool {
	return true
}

func (d *ListenDescriptor) IsServerStream() bool {
	return true
}

func (d *ListenDescriptor) IsCollection() bool {
	return false
}

func (d *ListenDescriptor) IsPlural() bool {
	return false
}

func (d *ListenDescriptor) HasResource() bool {
	return true
}

func (d *ListenDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *ListenDescriptor) GetVerb() string {
	return "listen"
}

func (d *ListenDescriptor) GetMethodName() string {
	return "Listen"
}

func (d *ListenDescriptor) GetFullMethodName() string {
	return "/ntt.proxies.v1alpha2.BrokerService/Listen"
}

func (d *ListenDescriptor) GetProtoPkgName() string {
	return "ntt.proxies.v1alpha2"
}

func (d *ListenDescriptor) GetApiName() string {
	return "BrokerService"
}

func (d *ListenDescriptor) GetServiceDomain() string {
	return "proxies.edgelq.com"
}

func (d *ListenDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *ListenDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return brokerServiceDescriptor
}

func (d *ListenDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return project.GetDescriptor()
}

func (d *ListenDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ListenDescriptorClientMsgHandle{}
}

func (d *ListenDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ListenDescriptorServerMsgHandle{}
}

func (h *ListenDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListenRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ListenRequest) *project.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if name := typedMsg.GetOpenRequest().GetProject(); name != nil {
			return name
		}
	}
	{
		if name := typedMsg.GetResumeRequest().GetProject(); name != nil {
			return name
		}
	}
	return (*project.Name)(nil)
}

func (h *ListenDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ListenRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ListenRequest) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *ListenDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *ListenDescriptorClientMsgHandle) ExtractResourceBody(msg proto.Message) gotenresource.Resource {
	typedMsg := msg.(*ListenRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBody(*ListenRequest) *project.Project
	})
	if ok {
		return override.OverrideExtractResourceBody(typedMsg)
	}
	return nil
}

func (h *ListenDescriptorClientMsgHandle) ExtractResourceBodies(msg proto.Message) gotenresource.ResourceList {
	typedMsg := msg.(*ListenRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBodies(*ListenRequest) []*project.Project
	})
	if ok {
		return project.ProjectList(override.OverrideExtractResourceBodies(typedMsg))
	}
	return nil
}

func (h *ListenDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListenResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ListenResponse) *project.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *ListenDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ListenResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ListenResponse) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *ListenDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *ListenDescriptorServerMsgHandle) ExtractResourceBody(msg proto.Message) gotenresource.Resource {
	typedMsg := msg.(*ListenResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBody(*ListenResponse) *project.Project
	})
	if ok {
		return override.OverrideExtractResourceBody(typedMsg)
	}
	return nil
}

func (h *ListenDescriptorServerMsgHandle) ExtractResourceBodies(msg proto.Message) gotenresource.ResourceList {
	typedMsg := msg.(*ListenResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBodies(*ListenResponse) []*project.Project
	})
	if ok {
		return project.ProjectList(override.OverrideExtractResourceBodies(typedMsg))
	}
	return nil
}

func GetListenDescriptor() *ListenDescriptor {
	return listenDescriptor
}

type AcceptDescriptor struct{}

type AcceptDescriptorClientMsgHandle struct{}

type AcceptDescriptorServerMsgHandle struct{}

func (d *AcceptDescriptor) NewEmptyClientMsg() proto.Message {
	return &AcceptRequest{}
}

func (d *AcceptDescriptor) NewEmptyServerMsg() proto.Message {
	return &AcceptResponse{}
}

func (d *AcceptDescriptor) IsUnary() bool {
	return false
}

func (d *AcceptDescriptor) IsClientStream() bool {
	return true
}

func (d *AcceptDescriptor) IsServerStream() bool {
	return true
}

func (d *AcceptDescriptor) IsCollection() bool {
	return false
}

func (d *AcceptDescriptor) IsPlural() bool {
	return false
}

func (d *AcceptDescriptor) HasResource() bool {
	return true
}

func (d *AcceptDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *AcceptDescriptor) GetVerb() string {
	return "accept"
}

func (d *AcceptDescriptor) GetMethodName() string {
	return "Accept"
}

func (d *AcceptDescriptor) GetFullMethodName() string {
	return "/ntt.proxies.v1alpha2.BrokerService/Accept"
}

func (d *AcceptDescriptor) GetProtoPkgName() string {
	return "ntt.proxies.v1alpha2"
}

func (d *AcceptDescriptor) GetApiName() string {
	return "BrokerService"
}

func (d *AcceptDescriptor) GetServiceDomain() string {
	return "proxies.edgelq.com"
}

func (d *AcceptDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *AcceptDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return brokerServiceDescriptor
}

func (d *AcceptDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return project.GetDescriptor()
}

func (d *AcceptDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &AcceptDescriptorClientMsgHandle{}
}

func (d *AcceptDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &AcceptDescriptorServerMsgHandle{}
}

func (h *AcceptDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*AcceptRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*AcceptRequest) *project.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if name := typedMsg.GetOpenRequest().GetProject(); name != nil {
			return name
		}
	}
	{
		if name := typedMsg.GetResumeRequest().GetProject(); name != nil {
			return name
		}
	}
	return (*project.Name)(nil)
}

func (h *AcceptDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*AcceptRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*AcceptRequest) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *AcceptDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *AcceptDescriptorClientMsgHandle) ExtractResourceBody(msg proto.Message) gotenresource.Resource {
	typedMsg := msg.(*AcceptRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBody(*AcceptRequest) *project.Project
	})
	if ok {
		return override.OverrideExtractResourceBody(typedMsg)
	}
	return nil
}

func (h *AcceptDescriptorClientMsgHandle) ExtractResourceBodies(msg proto.Message) gotenresource.ResourceList {
	typedMsg := msg.(*AcceptRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBodies(*AcceptRequest) []*project.Project
	})
	if ok {
		return project.ProjectList(override.OverrideExtractResourceBodies(typedMsg))
	}
	return nil
}

func (h *AcceptDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*AcceptResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*AcceptResponse) *project.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *AcceptDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*AcceptResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*AcceptResponse) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *AcceptDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *AcceptDescriptorServerMsgHandle) ExtractResourceBody(msg proto.Message) gotenresource.Resource {
	typedMsg := msg.(*AcceptResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBody(*AcceptResponse) *project.Project
	})
	if ok {
		return override.OverrideExtractResourceBody(typedMsg)
	}
	return nil
}

func (h *AcceptDescriptorServerMsgHandle) ExtractResourceBodies(msg proto.Message) gotenresource.ResourceList {
	typedMsg := msg.(*AcceptResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceBodies(*AcceptResponse) []*project.Project
	})
	if ok {
		return project.ProjectList(override.OverrideExtractResourceBodies(typedMsg))
	}
	return nil
}

func GetAcceptDescriptor() *AcceptDescriptor {
	return acceptDescriptor
}

type BrokerServiceDescriptor struct{}

func (d *BrokerServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		connectDescriptor,
		listenDescriptor,
		acceptDescriptor,
	}
}

func (d *BrokerServiceDescriptor) GetFullAPIName() string {
	return "/ntt.proxies.v1alpha2.BrokerService"
}

func (d *BrokerServiceDescriptor) GetProtoPkgName() string {
	return "ntt.proxies.v1alpha2"
}

func (d *BrokerServiceDescriptor) GetApiName() string {
	return "BrokerService"
}

func (d *BrokerServiceDescriptor) GetServiceDomain() string {
	return "proxies.edgelq.com"
}

func (d *BrokerServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func GetBrokerServiceDescriptor() *BrokerServiceDescriptor {
	return brokerServiceDescriptor
}

func initDescriptors() {
	brokerServiceDescriptor = &BrokerServiceDescriptor{}
	connectDescriptor = &ConnectDescriptor{}
	listenDescriptor = &ListenDescriptor{}
	acceptDescriptor = &AcceptDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(brokerServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(connectDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(listenDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(acceptDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}
