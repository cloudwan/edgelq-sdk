// Code generated by protoc-gen-goten-access
// Service: Monitoring
// DO NOT EDIT!!!

package monitoring_access

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	alert_access "github.com/cloudwan/edgelq-sdk/monitoring/access/v3/alert"
	alerting_condition_access "github.com/cloudwan/edgelq-sdk/monitoring/access/v3/alerting_condition"
	alerting_policy_access "github.com/cloudwan/edgelq-sdk/monitoring/access/v3/alerting_policy"
	metric_descriptor_access "github.com/cloudwan/edgelq-sdk/monitoring/access/v3/metric_descriptor"
	monitored_resource_descriptor_access "github.com/cloudwan/edgelq-sdk/monitoring/access/v3/monitored_resource_descriptor"
	phantom_time_serie_access "github.com/cloudwan/edgelq-sdk/monitoring/access/v3/phantom_time_serie"
	project_access "github.com/cloudwan/edgelq-sdk/monitoring/access/v3/project"
	monitoring_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/monitoring"
	alert "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alert"
	alerting_condition "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_condition"
	alerting_policy "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_policy"
	metric_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/metric_descriptor"
	monitored_resource_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/monitored_resource_descriptor"
	phantom_time_serie "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/phantom_time_serie"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/project"
)

type MonitoringApiAccess interface {
	gotenresource.Access

	alert.AlertAccess
	alerting_condition.AlertingConditionAccess
	alerting_policy.AlertingPolicyAccess
	metric_descriptor.MetricDescriptorAccess
	monitored_resource_descriptor.MonitoredResourceDescriptorAccess
	phantom_time_serie.PhantomTimeSerieAccess
	project.ProjectAccess
}

type apiMonitoringAccess struct {
	gotenresource.Access

	alert.AlertAccess
	alerting_condition.AlertingConditionAccess
	alerting_policy.AlertingPolicyAccess
	metric_descriptor.MetricDescriptorAccess
	monitored_resource_descriptor.MonitoredResourceDescriptorAccess
	phantom_time_serie.PhantomTimeSerieAccess
	project.ProjectAccess
}

func NewApiAccess(client monitoring_client.MonitoringClient) MonitoringApiAccess {

	alertAccess := alert_access.NewApiAlertAccess(client)
	alertingConditionAccess := alerting_condition_access.NewApiAlertingConditionAccess(client)
	alertingPolicyAccess := alerting_policy_access.NewApiAlertingPolicyAccess(client)
	metricDescriptorAccess := metric_descriptor_access.NewApiMetricDescriptorAccess(client)
	monitoredResourceDescriptorAccess := monitored_resource_descriptor_access.NewApiMonitoredResourceDescriptorAccess(client)
	phantomTimeSerieAccess := phantom_time_serie_access.NewApiPhantomTimeSerieAccess(client)
	projectAccess := project_access.NewApiProjectAccess(client)

	return &apiMonitoringAccess{
		Access: gotenresource.NewCompositeAccess(

			alert.AsAnyCastAccess(alertAccess),
			alerting_condition.AsAnyCastAccess(alertingConditionAccess),
			alerting_policy.AsAnyCastAccess(alertingPolicyAccess),
			metric_descriptor.AsAnyCastAccess(metricDescriptorAccess),
			monitored_resource_descriptor.AsAnyCastAccess(monitoredResourceDescriptorAccess),
			phantom_time_serie.AsAnyCastAccess(phantomTimeSerieAccess),
			project.AsAnyCastAccess(projectAccess),
		),

		AlertAccess:                       alertAccess,
		AlertingConditionAccess:           alertingConditionAccess,
		AlertingPolicyAccess:              alertingPolicyAccess,
		MetricDescriptorAccess:            metricDescriptorAccess,
		MonitoredResourceDescriptorAccess: monitoredResourceDescriptorAccess,
		PhantomTimeSerieAccess:            phantomTimeSerieAccess,
		ProjectAccess:                     projectAccess,
	}
}
