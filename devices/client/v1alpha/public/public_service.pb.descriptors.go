// Code generated by protoc-gen-goten-client
// API: PublicService
// DO NOT EDIT!!!

package public_client

import (
	"google.golang.org/protobuf/proto"

	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha/device"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = new(proto.Message)
	_ = new(gotenclient.MethodDescriptor)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &device.Device{}
)

var (
	descriptorsInitialized      bool
	publicServiceDescriptor     *PublicServiceDescriptor
	listPublicDevicesDescriptor *ListPublicDevicesDescriptor
)

type ListPublicDevicesDescriptor struct{}

type ListPublicDevicesDescriptorClientMsgHandle struct{}

type ListPublicDevicesDescriptorServerMsgHandle struct{}

func (d *ListPublicDevicesDescriptor) NewEmptyClientMsg() proto.Message {
	return &ListPublicDevicesRequest{}
}

func (d *ListPublicDevicesDescriptor) NewEmptyServerMsg() proto.Message {
	return &ListPublicDevicesResponse{}
}

func (d *ListPublicDevicesDescriptor) IsUnary() bool {
	return true
}

func (d *ListPublicDevicesDescriptor) IsClientStream() bool {
	return false
}

func (d *ListPublicDevicesDescriptor) IsServerStream() bool {
	return false
}

func (d *ListPublicDevicesDescriptor) IsCollection() bool {
	return true
}

func (d *ListPublicDevicesDescriptor) IsPlural() bool {
	return true
}

func (d *ListPublicDevicesDescriptor) HasResource() bool {
	return true
}

func (d *ListPublicDevicesDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *ListPublicDevicesDescriptor) GetVerb() string {
	return "listPublicDevices"
}

func (d *ListPublicDevicesDescriptor) GetMethodName() string {
	return "ListPublicDevices"
}

func (d *ListPublicDevicesDescriptor) GetFullMethodName() string {
	return "/ntt.devices.v1alpha.PublicService/ListPublicDevices"
}

func (d *ListPublicDevicesDescriptor) GetProtoPkgName() string {
	return "ntt.devices.v1alpha"
}

func (d *ListPublicDevicesDescriptor) GetApiName() string {
	return "PublicService"
}

func (d *ListPublicDevicesDescriptor) GetServiceDomain() string {
	return "devices.edgelq.com"
}

func (d *ListPublicDevicesDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *ListPublicDevicesDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return publicServiceDescriptor
}

func (d *ListPublicDevicesDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return device.GetDescriptor()
}

func (d *ListPublicDevicesDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ListPublicDevicesDescriptorClientMsgHandle{}
}

func (d *ListPublicDevicesDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ListPublicDevicesDescriptorServerMsgHandle{}
}

func (h *ListPublicDevicesDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListPublicDevicesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ListPublicDevicesRequest) *device.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *ListPublicDevicesDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ListPublicDevicesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ListPublicDevicesRequest) []*device.Name
	})
	if ok {
		return device.DeviceNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *ListPublicDevicesDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListPublicDevicesRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*ListPublicDevicesRequest) *device.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	{
		if parentName := typedMsg.GetParent(); parentName != nil {
			return parentName
		}
	}
	return (*device.ParentName)(nil)
}

func (h *ListPublicDevicesDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListPublicDevicesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ListPublicDevicesResponse) *device.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *ListPublicDevicesDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ListPublicDevicesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ListPublicDevicesResponse) []*device.Name
	})
	if ok {
		return device.DeviceNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	{
		if resources := typedMsg.GetDevices(); len(resources) > 0 {
			list := make(device.DeviceNameList, 0, len(resources))
			for _, res := range resources {
				list = append(list, res.GetName())
			}
			return list
		}
	}
	return (device.DeviceNameList)(nil)
}

func (h *ListPublicDevicesDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ListPublicDevicesResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*ListPublicDevicesResponse) *device.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func GetListPublicDevicesDescriptor() *ListPublicDevicesDescriptor {
	return listPublicDevicesDescriptor
}

type PublicServiceDescriptor struct{}

func (d *PublicServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		listPublicDevicesDescriptor,
	}
}

func (d *PublicServiceDescriptor) GetFullAPIName() string {
	return "/ntt.devices.v1alpha.PublicService"
}

func (d *PublicServiceDescriptor) GetProtoPkgName() string {
	return "ntt.devices.v1alpha"
}

func (d *PublicServiceDescriptor) GetApiName() string {
	return "PublicService"
}

func (d *PublicServiceDescriptor) GetServiceDomain() string {
	return "devices.edgelq.com"
}

func (d *PublicServiceDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func GetPublicServiceDescriptor() *PublicServiceDescriptor {
	return publicServiceDescriptor
}

func initDescriptors() {
	publicServiceDescriptor = &PublicServiceDescriptor{}
	listPublicDevicesDescriptor = &ListPublicDevicesDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(publicServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(listPublicDevicesDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}
