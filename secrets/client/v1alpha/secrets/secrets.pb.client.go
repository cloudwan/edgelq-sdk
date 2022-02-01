// Code generated by protoc-gen-goten-client
// Service: Secrets
// DO NOT EDIT!!!

package secrets_client

import (
	"google.golang.org/grpc"
)

// proto imports
import (
	encryption_client "github.com/cloudwan/edgelq-sdk/secrets/client/v1alpha/encryption"
	project_client "github.com/cloudwan/edgelq-sdk/secrets/client/v1alpha/project"
	secret_client "github.com/cloudwan/edgelq-sdk/secrets/client/v1alpha/secret"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &project_client.GetProjectRequest{}
	_ = &secret_client.GetSecretRequest{}
)

type SecretsClient interface {
	encryption_client.EncryptionServiceClient
	project_client.ProjectServiceClient
	secret_client.SecretServiceClient
}

type secretsClient struct {
	encryption_client.EncryptionServiceClient
	project_client.ProjectServiceClient
	secret_client.SecretServiceClient
}

func NewSecretsClient(cc grpc.ClientConnInterface) SecretsClient {
	return &secretsClient{
		EncryptionServiceClient: encryption_client.NewEncryptionServiceClient(cc),
		ProjectServiceClient:    project_client.NewProjectServiceClient(cc),
		SecretServiceClient:     secret_client.NewSecretServiceClient(cc),
	}
}
