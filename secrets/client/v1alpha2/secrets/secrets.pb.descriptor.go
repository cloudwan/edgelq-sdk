// Code generated by protoc-gen-goten-client
// Service: Secrets
// DO NOT EDIT!!!

package secrets_client

import (
	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project_client "github.com/cloudwan/edgelq-sdk/secrets/client/v1alpha2/project"
	secret_client "github.com/cloudwan/edgelq-sdk/secrets/client/v1alpha2/secret"
	project "github.com/cloudwan/edgelq-sdk/secrets/resources/v1alpha2/project"
	secret "github.com/cloudwan/edgelq-sdk/secrets/resources/v1alpha2/secret"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = new(gotenclient.MethodDescriptor)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &project_client.GetProjectRequest{}
	_ = &secret.Secret{}
	_ = &secret_client.GetSecretRequest{}
)

var (
	descriptorInitialized bool
	secretsDescriptor     *SecretsDescriptor
)

type SecretsDescriptor struct{}

func (d *SecretsDescriptor) GetServiceName() string {
	return "secrets"
}

func (d *SecretsDescriptor) GetServiceDomain() string {
	return "secrets.edgelq.com"
}

func (d *SecretsDescriptor) GetVersion() string {
	return "v1alpha2"
}

func (d *SecretsDescriptor) GetNextVersion() string {

	return "v1"
}

func (d *SecretsDescriptor) AllResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		project.GetDescriptor(),
		secret.GetDescriptor(),
	}
}

func (d *SecretsDescriptor) AllApiDescriptors() []gotenclient.ApiDescriptor {
	return []gotenclient.ApiDescriptor{
		project_client.GetProjectServiceDescriptor(),
		secret_client.GetSecretServiceDescriptor(),
	}
}

func (d *SecretsDescriptor) AllImportedServiceInfos() []gotenclient.ServiceImportInfo {
	return []gotenclient.ServiceImportInfo{}
}

func GetSecretsDescriptor() *SecretsDescriptor {
	return secretsDescriptor
}

func initDescriptor() {
	secretsDescriptor = &SecretsDescriptor{}
	gotenclient.GetRegistry().RegisterSvcDescriptor(secretsDescriptor)
}

func init() {
	if !descriptorInitialized {
		initDescriptor()
		descriptorInitialized = true
	}
}
