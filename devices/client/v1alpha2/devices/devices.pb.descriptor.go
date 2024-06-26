// Code generated by protoc-gen-goten-client
// Service: Devices
// DO NOT EDIT!!!

package devices_client

import (
	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	broker_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha2/broker"
	customized_image_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha2/customized_image"
	device_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha2/device"
	device_type_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha2/device_type"
	os_image_profile_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha2/os_image_profile"
	os_version_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha2/os_version"
	project_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha2/project"
	provisioning_approval_request_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha2/provisioning_approval_request"
	provisioning_policy_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha2/provisioning_policy"
	public_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha2/public"
	customized_image "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/customized_image"
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/device"
	device_type "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/device_type"
	os_image_profile "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/os_image_profile"
	os_version "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/os_version"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/project"
	provisioning_approval_request "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/provisioning_approval_request"
	provisioning_policy "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/provisioning_policy"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = new(gotenclient.MethodDescriptor)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &customized_image.CustomizedImage{}
	_ = &customized_image_client.GetCustomizedImageRequest{}
	_ = &device.Device{}
	_ = &device_client.GetDeviceRequest{}
	_ = &device_type.DeviceType{}
	_ = &device_type_client.GetDeviceTypeRequest{}
	_ = &os_image_profile.OsImageProfile{}
	_ = &os_image_profile_client.GetOsImageProfileRequest{}
	_ = &os_version.OsVersion{}
	_ = &os_version_client.GetOsVersionRequest{}
	_ = &project.Project{}
	_ = &project_client.GetProjectRequest{}
	_ = &provisioning_approval_request.ProvisioningApprovalRequest{}
	_ = &provisioning_approval_request_client.GetProvisioningApprovalRequestRequest{}
	_ = &provisioning_policy.ProvisioningPolicy{}
	_ = &provisioning_policy_client.GetProvisioningPolicyRequest{}
)

var (
	descriptorInitialized bool
	devicesDescriptor     *DevicesDescriptor
)

type DevicesDescriptor struct{}

func (d *DevicesDescriptor) GetServiceName() string {
	return "devices"
}

func (d *DevicesDescriptor) GetServiceDomain() string {
	return "devices.edgelq.com"
}

func (d *DevicesDescriptor) GetVersion() string {
	return "v1alpha2"
}

func (d *DevicesDescriptor) GetNextVersion() string {

	return "v1"
}

func (d *DevicesDescriptor) AllResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		customized_image.GetDescriptor(),
		device.GetDescriptor(),
		device_type.GetDescriptor(),
		os_image_profile.GetDescriptor(),
		os_version.GetDescriptor(),
		project.GetDescriptor(),
		provisioning_approval_request.GetDescriptor(),
		provisioning_policy.GetDescriptor(),
	}
}

func (d *DevicesDescriptor) AllApiDescriptors() []gotenclient.ApiDescriptor {
	return []gotenclient.ApiDescriptor{
		broker_client.GetBrokerServiceDescriptor(),
		customized_image_client.GetCustomizedImageServiceDescriptor(),
		device_client.GetDeviceServiceDescriptor(),
		device_type_client.GetDeviceTypeServiceDescriptor(),
		os_image_profile_client.GetOsImageProfileServiceDescriptor(),
		os_version_client.GetOsVersionServiceDescriptor(),
		project_client.GetProjectServiceDescriptor(),
		provisioning_approval_request_client.GetProvisioningApprovalRequestServiceDescriptor(),
		provisioning_policy_client.GetProvisioningPolicyServiceDescriptor(),
		public_client.GetPublicServiceDescriptor(),
	}
}

func (d *DevicesDescriptor) AllImportedServiceInfos() []gotenclient.ServiceImportInfo {
	return []gotenclient.ServiceImportInfo{
		{
			Domain:  "iam.edgelq.com",
			Version: "v1alpha2",
		},
	}
}

func GetDevicesDescriptor() *DevicesDescriptor {
	return devicesDescriptor
}

func initDescriptor() {
	devicesDescriptor = &DevicesDescriptor{}
	gotenclient.GetRegistry().RegisterSvcDescriptor(devicesDescriptor)
}

func init() {
	if !descriptorInitialized {
		initDescriptor()
		descriptorInitialized = true
	}
}
