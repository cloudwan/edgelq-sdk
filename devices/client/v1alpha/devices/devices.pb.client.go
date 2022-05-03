// Code generated by protoc-gen-goten-client
// Service: Devices
// DO NOT EDIT!!!

package devices_client

import (
	"google.golang.org/grpc"
)

// proto imports
import (
	broker_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha/broker"
	device_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha/device"
	project_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha/project"
	provisioning_approval_request_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha/provisioning_approval_request"
	provisioning_policy_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha/provisioning_policy"
	public_client "github.com/cloudwan/edgelq-sdk/devices/client/v1alpha/public"
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha/device"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha/project"
	provisioning_approval_request "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha/provisioning_approval_request"
	provisioning_policy "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha/provisioning_policy"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &device.Device{}
	_ = &device_client.GetDeviceRequest{}
	_ = &project.Project{}
	_ = &project_client.GetProjectRequest{}
	_ = &provisioning_approval_request.ProvisioningApprovalRequest{}
	_ = &provisioning_approval_request_client.GetProvisioningApprovalRequestRequest{}
	_ = &provisioning_policy.ProvisioningPolicy{}
	_ = &provisioning_policy_client.GetProvisioningPolicyRequest{}
)

type DevicesClient interface {
	broker_client.BrokerServiceClient
	device_client.DeviceServiceClient
	project_client.ProjectServiceClient
	provisioning_approval_request_client.ProvisioningApprovalRequestServiceClient
	provisioning_policy_client.ProvisioningPolicyServiceClient
	public_client.PublicServiceClient
}

type devicesClient struct {
	broker_client.BrokerServiceClient
	device_client.DeviceServiceClient
	project_client.ProjectServiceClient
	provisioning_approval_request_client.ProvisioningApprovalRequestServiceClient
	provisioning_policy_client.ProvisioningPolicyServiceClient
	public_client.PublicServiceClient
}

func NewDevicesClient(cc grpc.ClientConnInterface) DevicesClient {
	return &devicesClient{
		BrokerServiceClient:                      broker_client.NewBrokerServiceClient(cc),
		DeviceServiceClient:                      device_client.NewDeviceServiceClient(cc),
		ProjectServiceClient:                     project_client.NewProjectServiceClient(cc),
		ProvisioningApprovalRequestServiceClient: provisioning_approval_request_client.NewProvisioningApprovalRequestServiceClient(cc),
		ProvisioningPolicyServiceClient:          provisioning_policy_client.NewProvisioningPolicyServiceClient(cc),
		PublicServiceClient:                      public_client.NewPublicServiceClient(cc),
	}
}
