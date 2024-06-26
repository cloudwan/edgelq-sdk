// Code generated by protoc-gen-goten-client
// Service: Monitoring
// DO NOT EDIT!!!

package monitoring_client

import (
	"google.golang.org/grpc"
)

// proto imports
import (
	alert_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/alert"
	alerting_condition_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/alerting_condition"
	alerting_policy_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/alerting_policy"
	metric_descriptor_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/metric_descriptor"
	monitored_resource_descriptor_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/monitored_resource_descriptor"
	notification_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/notification"
	notification_channel_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/notification_channel"
	phantom_time_serie_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/phantom_time_serie"
	project_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/project"
	recovery_store_sharding_info_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/recovery_store_sharding_info"
	time_serie_client "github.com/cloudwan/edgelq-sdk/monitoring/client/v3/time_serie"
	alert "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alert"
	alerting_condition "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_condition"
	alerting_policy "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_policy"
	metric_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/metric_descriptor"
	monitored_resource_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/monitored_resource_descriptor"
	notification "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/notification"
	notification_channel "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/notification_channel"
	phantom_time_serie "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/phantom_time_serie"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/project"
	recovery_store_sharding_info "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/recovery_store_sharding_info"
	time_serie "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/time_serie"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &alert.Alert{}
	_ = &alert_client.GetAlertRequest{}
	_ = &alerting_condition.AlertingCondition{}
	_ = &alerting_condition_client.GetAlertingConditionRequest{}
	_ = &alerting_policy.AlertingPolicy{}
	_ = &alerting_policy_client.GetAlertingPolicyRequest{}
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

type MonitoringClient interface {
	alert_client.AlertServiceClient
	alerting_condition_client.AlertingConditionServiceClient
	alerting_policy_client.AlertingPolicyServiceClient
	metric_descriptor_client.MetricDescriptorServiceClient
	monitored_resource_descriptor_client.MonitoredResourceDescriptorServiceClient
	notification_channel_client.NotificationChannelServiceClient
	notification_client.NotificationServiceClient
	phantom_time_serie_client.PhantomTimeSerieServiceClient
	project_client.ProjectServiceClient
	recovery_store_sharding_info_client.RecoveryStoreShardingInfoServiceClient
	time_serie_client.TimeSerieServiceClient
}

type monitoringClient struct {
	alert_client.AlertServiceClient
	alerting_condition_client.AlertingConditionServiceClient
	alerting_policy_client.AlertingPolicyServiceClient
	metric_descriptor_client.MetricDescriptorServiceClient
	monitored_resource_descriptor_client.MonitoredResourceDescriptorServiceClient
	notification_channel_client.NotificationChannelServiceClient
	notification_client.NotificationServiceClient
	phantom_time_serie_client.PhantomTimeSerieServiceClient
	project_client.ProjectServiceClient
	recovery_store_sharding_info_client.RecoveryStoreShardingInfoServiceClient
	time_serie_client.TimeSerieServiceClient
}

func NewMonitoringClient(cc grpc.ClientConnInterface) MonitoringClient {
	return &monitoringClient{
		AlertServiceClient:                       alert_client.NewAlertServiceClient(cc),
		AlertingConditionServiceClient:           alerting_condition_client.NewAlertingConditionServiceClient(cc),
		AlertingPolicyServiceClient:              alerting_policy_client.NewAlertingPolicyServiceClient(cc),
		MetricDescriptorServiceClient:            metric_descriptor_client.NewMetricDescriptorServiceClient(cc),
		MonitoredResourceDescriptorServiceClient: monitored_resource_descriptor_client.NewMonitoredResourceDescriptorServiceClient(cc),
		NotificationChannelServiceClient:         notification_channel_client.NewNotificationChannelServiceClient(cc),
		NotificationServiceClient:                notification_client.NewNotificationServiceClient(cc),
		PhantomTimeSerieServiceClient:            phantom_time_serie_client.NewPhantomTimeSerieServiceClient(cc),
		ProjectServiceClient:                     project_client.NewProjectServiceClient(cc),
		RecoveryStoreShardingInfoServiceClient:   recovery_store_sharding_info_client.NewRecoveryStoreShardingInfoServiceClient(cc),
		TimeSerieServiceClient:                   time_serie_client.NewTimeSerieServiceClient(cc),
	}
}
