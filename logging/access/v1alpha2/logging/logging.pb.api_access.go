// Code generated by protoc-gen-goten-access
// Service: Logging
// DO NOT EDIT!!!

package logging_access

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	log_descriptor_access "github.com/cloudwan/edgelq-sdk/logging/access/v1alpha2/log_descriptor"
	logging_client "github.com/cloudwan/edgelq-sdk/logging/client/v1alpha2/logging"
	log_descriptor "github.com/cloudwan/edgelq-sdk/logging/resources/v1alpha2/log_descriptor"
)

type LoggingApiAccess interface {
	gotenresource.Access

	log_descriptor.LogDescriptorAccess
}

type apiLoggingAccess struct {
	gotenresource.Access

	log_descriptor.LogDescriptorAccess
}

func NewApiAccess(client logging_client.LoggingClient) LoggingApiAccess {

	logDescriptorAccess := log_descriptor_access.NewApiLogDescriptorAccess(client)

	return &apiLoggingAccess{
		Access: gotenresource.NewCompositeAccess(

			log_descriptor.AsAnyCastAccess(logDescriptorAccess),
		),

		LogDescriptorAccess: logDescriptorAccess,
	}
}
