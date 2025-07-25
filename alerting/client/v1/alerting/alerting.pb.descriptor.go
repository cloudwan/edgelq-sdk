// Code generated by protoc-gen-goten-client
// Service: Alerting
// DO NOT EDIT!!!

package alerting_client

import (
	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	alert_client "github.com/cloudwan/edgelq-sdk/alerting/client/v1/alert"
	document_client "github.com/cloudwan/edgelq-sdk/alerting/client/v1/document"
	edge_watch_service_client "github.com/cloudwan/edgelq-sdk/alerting/client/v1/edge_watch_service"
	log_condition_client "github.com/cloudwan/edgelq-sdk/alerting/client/v1/log_condition"
	log_condition_template_client "github.com/cloudwan/edgelq-sdk/alerting/client/v1/log_condition_template"
	notification_channel_client "github.com/cloudwan/edgelq-sdk/alerting/client/v1/notification_channel"
	policy_client "github.com/cloudwan/edgelq-sdk/alerting/client/v1/policy"
	policy_template_client "github.com/cloudwan/edgelq-sdk/alerting/client/v1/policy_template"
	ts_condition_client "github.com/cloudwan/edgelq-sdk/alerting/client/v1/ts_condition"
	ts_condition_template_client "github.com/cloudwan/edgelq-sdk/alerting/client/v1/ts_condition_template"
	ts_entry_client "github.com/cloudwan/edgelq-sdk/alerting/client/v1/ts_entry"
	alert "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/alert"
	document "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/document"
	log_condition "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/log_condition"
	log_condition_template "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/log_condition_template"
	notification_channel "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/notification_channel"
	policy "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/policy"
	policy_template "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/policy_template"
	ts_condition "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/ts_condition"
	ts_condition_template "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/ts_condition_template"
	ts_entry "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/ts_entry"
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
	_ = &document.Document{}
	_ = &document_client.GetDocumentRequest{}
	_ = &log_condition.LogCondition{}
	_ = &log_condition_client.GetLogConditionRequest{}
	_ = &log_condition_template.LogConditionTemplate{}
	_ = &log_condition_template_client.GetLogConditionTemplateRequest{}
	_ = &notification_channel.NotificationChannel{}
	_ = &notification_channel_client.GetNotificationChannelRequest{}
	_ = &policy.Policy{}
	_ = &policy_client.GetPolicyRequest{}
	_ = &policy_template.PolicyTemplate{}
	_ = &policy_template_client.GetPolicyTemplateRequest{}
	_ = &ts_condition.TsCondition{}
	_ = &ts_condition_client.GetTsConditionRequest{}
	_ = &ts_condition_template.TsConditionTemplate{}
	_ = &ts_condition_template_client.GetTsConditionTemplateRequest{}
	_ = &ts_entry.TsEntry{}
	_ = &ts_entry_client.GetTsEntryRequest{}
)

var (
	descriptorInitialized bool
	alertingDescriptor    *AlertingDescriptor
)

type AlertingDescriptor struct{}

func (d *AlertingDescriptor) GetServiceName() string {
	return "alerting"
}

func (d *AlertingDescriptor) GetServiceDomain() string {
	return "alerting.edgelq.com"
}

func (d *AlertingDescriptor) GetVersion() string {
	return "v1"
}

func (d *AlertingDescriptor) GetNextVersion() string {

	return ""
}

func (d *AlertingDescriptor) AllResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		alert.GetDescriptor(),
		document.GetDescriptor(),
		log_condition.GetDescriptor(),
		log_condition_template.GetDescriptor(),
		notification_channel.GetDescriptor(),
		policy.GetDescriptor(),
		policy_template.GetDescriptor(),
		ts_condition.GetDescriptor(),
		ts_condition_template.GetDescriptor(),
		ts_entry.GetDescriptor(),
	}
}

func (d *AlertingDescriptor) AllApiDescriptors() []gotenclient.ApiDescriptor {
	return []gotenclient.ApiDescriptor{
		alert_client.GetAlertServiceDescriptor(),
		document_client.GetDocumentServiceDescriptor(),
		edge_watch_service_client.GetEdgeWatchServiceServiceDescriptor(),
		log_condition_client.GetLogConditionServiceDescriptor(),
		log_condition_template_client.GetLogConditionTemplateServiceDescriptor(),
		notification_channel_client.GetNotificationChannelServiceDescriptor(),
		policy_client.GetPolicyServiceDescriptor(),
		policy_template_client.GetPolicyTemplateServiceDescriptor(),
		ts_condition_client.GetTsConditionServiceDescriptor(),
		ts_condition_template_client.GetTsConditionTemplateServiceDescriptor(),
		ts_entry_client.GetTsEntryServiceDescriptor(),
	}
}

func (d *AlertingDescriptor) AllImportedServiceInfos() []gotenclient.ServiceImportInfo {
	return []gotenclient.ServiceImportInfo{
		{
			Domain:  "audit.edgelq.com",
			Version: "v1",
		},
		{
			Domain:  "iam.edgelq.com",
			Version: "v1",
		},
		{
			Domain:  "logging.edgelq.com",
			Version: "v1",
		},
		{
			Domain:  "meta.goten.com",
			Version: "v1",
		},
		{
			Domain:  "monitoring.edgelq.com",
			Version: "v4",
		},
		{
			Domain:  "secrets.edgelq.com",
			Version: "v1",
		},
	}
}

func GetAlertingDescriptor() *AlertingDescriptor {
	return alertingDescriptor
}

func initDescriptor() {
	alertingDescriptor = &AlertingDescriptor{}
	gotenclient.GetRegistry().RegisterSvcDescriptor(alertingDescriptor)
}

func init() {
	if !descriptorInitialized {
		initDescriptor()
		descriptorInitialized = true
	}
}
