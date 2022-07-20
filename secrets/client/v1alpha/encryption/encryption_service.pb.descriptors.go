// Code generated by protoc-gen-goten-client
// API: EncryptionService
// DO NOT EDIT!!!

package encryption_client

import (
	"google.golang.org/protobuf/proto"

	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/secrets/resources/v1alpha/project"
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
)

var (
	descriptorsInitialized      bool
	encryptionServiceDescriptor *EncryptionServiceDescriptor
	encryptDataDescriptor       *EncryptDataDescriptor
	decryptDataDescriptor       *DecryptDataDescriptor
)

type EncryptDataDescriptor struct{}

type EncryptDataDescriptorClientMsgHandle struct{}

type EncryptDataDescriptorServerMsgHandle struct{}

func (d *EncryptDataDescriptor) NewEmptyClientMsg() proto.Message {
	return &EncryptDataRequest{}
}

func (d *EncryptDataDescriptor) NewEmptyServerMsg() proto.Message {
	return &EncryptDataResponse{}
}

func (d *EncryptDataDescriptor) IsUnary() bool {
	return true
}

func (d *EncryptDataDescriptor) IsClientStream() bool {
	return false
}

func (d *EncryptDataDescriptor) IsServerStream() bool {
	return false
}

func (d *EncryptDataDescriptor) IsCollection() bool {
	return false
}

func (d *EncryptDataDescriptor) IsPlural() bool {
	return false
}

func (d *EncryptDataDescriptor) HasResource() bool {
	return true
}

func (d *EncryptDataDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *EncryptDataDescriptor) GetVerb() string {
	return "encryptData"
}

func (d *EncryptDataDescriptor) GetMethodName() string {
	return "EncryptData"
}

func (d *EncryptDataDescriptor) GetFullMethodName() string {
	return "/ntt.secrets.v1alpha.EncryptionService/EncryptData"
}

func (d *EncryptDataDescriptor) GetProtoPkgName() string {
	return "ntt.secrets.v1alpha"
}

func (d *EncryptDataDescriptor) GetApiName() string {
	return "EncryptionService"
}

func (d *EncryptDataDescriptor) GetServiceDomain() string {
	return "secrets.edgelq.com"
}

func (d *EncryptDataDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *EncryptDataDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return encryptionServiceDescriptor
}

func (d *EncryptDataDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return project.GetDescriptor()
}

func (d *EncryptDataDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &EncryptDataDescriptorClientMsgHandle{}
}

func (d *EncryptDataDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &EncryptDataDescriptorServerMsgHandle{}
}

func (h *EncryptDataDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*EncryptDataRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*EncryptDataRequest) *project.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *EncryptDataDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*EncryptDataRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*EncryptDataRequest) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *EncryptDataDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *EncryptDataDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*EncryptDataResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*EncryptDataResponse) *project.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *EncryptDataDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*EncryptDataResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*EncryptDataResponse) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *EncryptDataDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetEncryptDataDescriptor() *EncryptDataDescriptor {
	return encryptDataDescriptor
}

type DecryptDataDescriptor struct{}

type DecryptDataDescriptorClientMsgHandle struct{}

type DecryptDataDescriptorServerMsgHandle struct{}

func (d *DecryptDataDescriptor) NewEmptyClientMsg() proto.Message {
	return &DecryptDataRequest{}
}

func (d *DecryptDataDescriptor) NewEmptyServerMsg() proto.Message {
	return &DecryptDataResponse{}
}

func (d *DecryptDataDescriptor) IsUnary() bool {
	return true
}

func (d *DecryptDataDescriptor) IsClientStream() bool {
	return false
}

func (d *DecryptDataDescriptor) IsServerStream() bool {
	return false
}

func (d *DecryptDataDescriptor) IsCollection() bool {
	return false
}

func (d *DecryptDataDescriptor) IsPlural() bool {
	return false
}

func (d *DecryptDataDescriptor) HasResource() bool {
	return true
}

func (d *DecryptDataDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *DecryptDataDescriptor) GetVerb() string {
	return "decryptData"
}

func (d *DecryptDataDescriptor) GetMethodName() string {
	return "DecryptData"
}

func (d *DecryptDataDescriptor) GetFullMethodName() string {
	return "/ntt.secrets.v1alpha.EncryptionService/DecryptData"
}

func (d *DecryptDataDescriptor) GetProtoPkgName() string {
	return "ntt.secrets.v1alpha"
}

func (d *DecryptDataDescriptor) GetApiName() string {
	return "EncryptionService"
}

func (d *DecryptDataDescriptor) GetServiceDomain() string {
	return "secrets.edgelq.com"
}

func (d *DecryptDataDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func (d *DecryptDataDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return encryptionServiceDescriptor
}

func (d *DecryptDataDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return project.GetDescriptor()
}

func (d *DecryptDataDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &DecryptDataDescriptorClientMsgHandle{}
}

func (d *DecryptDataDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &DecryptDataDescriptorServerMsgHandle{}
}

func (h *DecryptDataDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*DecryptDataRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*DecryptDataRequest) *project.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *DecryptDataDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*DecryptDataRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*DecryptDataRequest) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *DecryptDataDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *DecryptDataDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*DecryptDataResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*DecryptDataResponse) *project.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *DecryptDataDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*DecryptDataResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*DecryptDataResponse) []*project.Name
	})
	if ok {
		return project.ProjectNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *DecryptDataDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetDecryptDataDescriptor() *DecryptDataDescriptor {
	return decryptDataDescriptor
}

type EncryptionServiceDescriptor struct{}

func (d *EncryptionServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		encryptDataDescriptor,
		decryptDataDescriptor,
	}
}

func (d *EncryptionServiceDescriptor) GetFullAPIName() string {
	return "/ntt.secrets.v1alpha.EncryptionService"
}

func (d *EncryptionServiceDescriptor) GetProtoPkgName() string {
	return "ntt.secrets.v1alpha"
}

func (d *EncryptionServiceDescriptor) GetApiName() string {
	return "EncryptionService"
}

func (d *EncryptionServiceDescriptor) GetServiceDomain() string {
	return "secrets.edgelq.com"
}

func (d *EncryptionServiceDescriptor) GetServiceVersion() string {
	return "v1alpha"
}

func GetEncryptionServiceDescriptor() *EncryptionServiceDescriptor {
	return encryptionServiceDescriptor
}

func initDescriptors() {
	encryptionServiceDescriptor = &EncryptionServiceDescriptor{}
	encryptDataDescriptor = &EncryptDataDescriptor{}
	decryptDataDescriptor = &DecryptDataDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(encryptionServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(encryptDataDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(decryptDataDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}