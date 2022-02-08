// Code generated by protoc-gen-goten-access
// Service: Applications
// DO NOT EDIT!!!

package applications_access

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	config_map_access "github.com/cloudwan/edgelq-sdk/applications/access/v1alpha/config_map"
	distribution_access "github.com/cloudwan/edgelq-sdk/applications/access/v1alpha/distribution"
	pod_access "github.com/cloudwan/edgelq-sdk/applications/access/v1alpha/pod"
	project_access "github.com/cloudwan/edgelq-sdk/applications/access/v1alpha/project"
	applications_client "github.com/cloudwan/edgelq-sdk/applications/client/v1alpha/applications"
	config_map "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha/config_map"
	distribution "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha/distribution"
	pod "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha/pod"
	project "github.com/cloudwan/edgelq-sdk/applications/resources/v1alpha/project"
)

type ApplicationsApiAccess interface {
	gotenresource.Access

	config_map.ConfigMapAccess
	distribution.DistributionAccess
	pod.PodAccess
	project.ProjectAccess
}

type apiApplicationsAccess struct {
	gotenresource.Access

	config_map.ConfigMapAccess
	distribution.DistributionAccess
	pod.PodAccess
	project.ProjectAccess
}

func NewApiAccess(client applications_client.ApplicationsClient) ApplicationsApiAccess {

	configMapAccess := config_map_access.NewApiConfigMapAccess(client)
	distributionAccess := distribution_access.NewApiDistributionAccess(client)
	podAccess := pod_access.NewApiPodAccess(client)
	projectAccess := project_access.NewApiProjectAccess(client)

	return &apiApplicationsAccess{
		Access: gotenresource.NewCompositeAccess(

			config_map.AsAnyCastAccess(configMapAccess),
			distribution.AsAnyCastAccess(distributionAccess),
			pod.AsAnyCastAccess(podAccess),
			project.AsAnyCastAccess(projectAccess),
		),

		ConfigMapAccess:    configMapAccess,
		DistributionAccess: distributionAccess,
		PodAccess:          podAccess,
		ProjectAccess:      projectAccess,
	}
}