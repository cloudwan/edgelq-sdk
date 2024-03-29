// Code generated by protoc-gen-goten-client
// Service: Monitoring
// DO NOT EDIT!!!

package monitoring_client

import (
	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	alert_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v4/alert"
	alerting_condition_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v4/alerting_condition"
	alerting_policy_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v4/alerting_policy"
	bucket_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v4/bucket"
	metric_descriptor_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v4/metric_descriptor"
	monitored_resource_descriptor_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v4/monitored_resource_descriptor"
	notification_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v4/notification"
	notification_channel_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v4/notification_channel"
	phantom_time_serie_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v4/phantom_time_serie"
	project_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v4/project"
	recovery_store_sharding_info_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v4/recovery_store_sharding_info"
	time_serie_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v4/time_serie"
	alert "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/alert"
	alerting_condition "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/alerting_condition"
	alerting_policy "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/alerting_policy"
	bucket "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/bucket"
	metric_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/metric_descriptor"
	monitored_resource_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/monitored_resource_descriptor"
	notification "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/notification"
	notification_channel "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/notification_channel"
	phantom_time_serie "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/phantom_time_serie"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
	recovery_store_sharding_info "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/recovery_store_sharding_info"
	time_serie "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/time_serie"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = new(gotenclient.MethodDescriptor)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &alert.Alert{}
	_ = &alert_client.GetAlertRequest{}
	_ = &alerting_condition.AlertingCondition{}
	_ = &alerting_condition_client.GetAlertingConditionRequest{}
	_ = &alerting_policy.AlertingPolicy{}
	_ = &alerting_policy_client.GetAlertingPolicyRequest{}
	_ = &bucket.Bucket{}
	_ = &bucket_client.GetBucketRequest{}
	_ = &metric_descriptor.MetricDescriptor{}
	_ = &metric_descriptor_client.BatchGetMetricDescriptorsRequest{}
	_ = &monitored_resource_descriptor.MonitoredResourceDescriptor{}
	_ = &monitored_resource_descriptor_client.BatchGetMonitoredResourceDescriptorsRequest{}
	_ = &notification.Notification{}
	_ = &notification_channel.NotificationChannel{}
	_ = &notification_channel_client.GetNotificationChannelRequest{}
	_ = &notification_client.GetNotificationRequest{}
	_ = &phantom_time_serie.PhantomTimeSerie{}
	_ = &phantom_time_serie_client.GetPhantomTimeSerieRequest{}
	_ = &project.Project{}
	_ = &project_client.GetProjectRequest{}
	_ = &recovery_store_sharding_info.RecoveryStoreShardingInfo{}
	_ = &recovery_store_sharding_info_client.GetRecoveryStoreShardingInfoRequest{}
	_ = &time_serie.Point{}
)

var (
	descriptorInitialized bool
	monitoringDescriptor  *MonitoringDescriptor
)

type MonitoringDescriptor struct{}

func (d *MonitoringDescriptor) GetServiceName() string {
	return "monitoring"
}

func (d *MonitoringDescriptor) GetServiceDomain() string {
	return "monitoring.edgelq.com"
}

func (d *MonitoringDescriptor) GetVersion() string {
	return "v4"
}

func (d *MonitoringDescriptor) GetNextVersion() string {

	return ""
}

func (d *MonitoringDescriptor) AllResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		alert.GetDescriptor(),
		alerting_condition.GetDescriptor(),
		alerting_policy.GetDescriptor(),
		bucket.GetDescriptor(),
		metric_descriptor.GetDescriptor(),
		monitored_resource_descriptor.GetDescriptor(),
		notification.GetDescriptor(),
		notification_channel.GetDescriptor(),
		phantom_time_serie.GetDescriptor(),
		project.GetDescriptor(),
		recovery_store_sharding_info.GetDescriptor(),
	}
}

func (d *MonitoringDescriptor) AllApiDescriptors() []gotenclient.ApiDescriptor {
	return []gotenclient.ApiDescriptor{
		alert_client.GetAlertServiceDescriptor(),
		alerting_condition_client.GetAlertingConditionServiceDescriptor(),
		alerting_policy_client.GetAlertingPolicyServiceDescriptor(),
		bucket_client.GetBucketServiceDescriptor(),
		metric_descriptor_client.GetMetricDescriptorServiceDescriptor(),
		monitored_resource_descriptor_client.GetMonitoredResourceDescriptorServiceDescriptor(),
		notification_channel_client.GetNotificationChannelServiceDescriptor(),
		notification_client.GetNotificationServiceDescriptor(),
		phantom_time_serie_client.GetPhantomTimeSerieServiceDescriptor(),
		project_client.GetProjectServiceDescriptor(),
		recovery_store_sharding_info_client.GetRecoveryStoreShardingInfoServiceDescriptor(),
		time_serie_client.GetTimeSerieServiceDescriptor(),
	}
}

func (d *MonitoringDescriptor) AllImportedServiceInfos() []gotenclient.ServiceImportInfo {
	return []gotenclient.ServiceImportInfo{
		{
			Domain:  "meta.goten.com",
			Version: "v1",
		},
	}
}

func GetMonitoringDescriptor() *MonitoringDescriptor {
	return monitoringDescriptor
}

func initDescriptor() {
	monitoringDescriptor = &MonitoringDescriptor{}
	gotenclient.GetRegistry().RegisterSvcDescriptor(monitoringDescriptor)
}

func init() {
	if !descriptorInitialized {
		initDescriptor()
		descriptorInitialized = true
	}
}
