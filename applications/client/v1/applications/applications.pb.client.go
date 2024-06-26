// Code generated by protoc-gen-goten-client
// Service: Applications
// DO NOT EDIT!!!

package applications_client

import (
	"google.golang.org/grpc"
)

// proto imports
import (
	config_map_client "github.com/cloudwan/edgelq-sdk/applications/client/v1/config_map"
	distribution_client "github.com/cloudwan/edgelq-sdk/applications/client/v1/distribution"
	pod_client "github.com/cloudwan/edgelq-sdk/applications/client/v1/pod"
	project_client "github.com/cloudwan/edgelq-sdk/applications/client/v1/project"
	config_map "github.com/cloudwan/edgelq-sdk/applications/resources/v1/config_map"
	distribution "github.com/cloudwan/edgelq-sdk/applications/resources/v1/distribution"
	pod "github.com/cloudwan/edgelq-sdk/applications/resources/v1/pod"
	project "github.com/cloudwan/edgelq-sdk/applications/resources/v1/project"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &config_map.ConfigMap{}
	_ = &config_map_client.GetConfigMapRequest{}
	_ = &distribution.Distribution{}
	_ = &distribution_client.GetDistributionRequest{}
	_ = &pod.Pod{}
	_ = &pod_client.GetPodRequest{}
	_ = &project.Project{}
	_ = &project_client.GetProjectRequest{}
)

type ApplicationsClient interface {
	config_map_client.ConfigMapServiceClient
	distribution_client.DistributionServiceClient
	pod_client.PodServiceClient
	project_client.ProjectServiceClient
}

type applicationsClient struct {
	config_map_client.ConfigMapServiceClient
	distribution_client.DistributionServiceClient
	pod_client.PodServiceClient
	project_client.ProjectServiceClient
}

func NewApplicationsClient(cc grpc.ClientConnInterface) ApplicationsClient {
	return &applicationsClient{
		ConfigMapServiceClient:    config_map_client.NewConfigMapServiceClient(cc),
		DistributionServiceClient: distribution_client.NewDistributionServiceClient(cc),
		PodServiceClient:          pod_client.NewPodServiceClient(cc),
		ProjectServiceClient:      project_client.NewProjectServiceClient(cc),
	}
}
