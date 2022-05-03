// Code generated by protoc-gen-goten-client
// Service: Proxies
// DO NOT EDIT!!!

package proxies_client

import (
	"google.golang.org/grpc"
)

// proto imports
import (
	broker_client "github.com/cloudwan/edgelq-sdk/proxies/client/v1alpha2/broker"
	project_client "github.com/cloudwan/edgelq-sdk/proxies/client/v1alpha2/project"
	project "github.com/cloudwan/edgelq-sdk/proxies/resources/v1alpha2/project"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &project_client.GetProjectRequest{}
)

type ProxiesClient interface {
	broker_client.BrokerServiceClient
	project_client.ProjectServiceClient
}

type proxiesClient struct {
	broker_client.BrokerServiceClient
	project_client.ProjectServiceClient
}

func NewProxiesClient(cc grpc.ClientConnInterface) ProxiesClient {
	return &proxiesClient{
		BrokerServiceClient:  broker_client.NewBrokerServiceClient(cc),
		ProjectServiceClient: project_client.NewProjectServiceClient(cc),
	}
}
