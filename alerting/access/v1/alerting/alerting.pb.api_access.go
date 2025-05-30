// Code generated by protoc-gen-goten-access
// Service: Alerting
// DO NOT EDIT!!!

package alerting_access

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	alert_access "github.com/cloudwan/edgelq-sdk/alerting/access/v1/alert"
	document_access "github.com/cloudwan/edgelq-sdk/alerting/access/v1/document"
	log_condition_access "github.com/cloudwan/edgelq-sdk/alerting/access/v1/log_condition"
	log_condition_template_access "github.com/cloudwan/edgelq-sdk/alerting/access/v1/log_condition_template"
	notification_channel_access "github.com/cloudwan/edgelq-sdk/alerting/access/v1/notification_channel"
	policy_access "github.com/cloudwan/edgelq-sdk/alerting/access/v1/policy"
	policy_template_access "github.com/cloudwan/edgelq-sdk/alerting/access/v1/policy_template"
	ts_condition_access "github.com/cloudwan/edgelq-sdk/alerting/access/v1/ts_condition"
	ts_condition_template_access "github.com/cloudwan/edgelq-sdk/alerting/access/v1/ts_condition_template"
	ts_entry_access "github.com/cloudwan/edgelq-sdk/alerting/access/v1/ts_entry"
	alerting_client "github.com/cloudwan/edgelq-sdk/alerting/client/v1/alerting"
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

type AlertingApiAccess interface {
	gotenresource.Access

	alert.AlertAccess
	document.DocumentAccess
	log_condition.LogConditionAccess
	log_condition_template.LogConditionTemplateAccess
	notification_channel.NotificationChannelAccess
	policy.PolicyAccess
	policy_template.PolicyTemplateAccess
	ts_condition.TsConditionAccess
	ts_condition_template.TsConditionTemplateAccess
	ts_entry.TsEntryAccess
}

type apiAlertingAccess struct {
	gotenresource.Access

	alert.AlertAccess
	document.DocumentAccess
	log_condition.LogConditionAccess
	log_condition_template.LogConditionTemplateAccess
	notification_channel.NotificationChannelAccess
	policy.PolicyAccess
	policy_template.PolicyTemplateAccess
	ts_condition.TsConditionAccess
	ts_condition_template.TsConditionTemplateAccess
	ts_entry.TsEntryAccess
}

func NewApiAccess(client alerting_client.AlertingClient) AlertingApiAccess {

	alertAccess := alert_access.NewApiAlertAccess(client)
	documentAccess := document_access.NewApiDocumentAccess(client)
	logConditionAccess := log_condition_access.NewApiLogConditionAccess(client)
	logConditionTemplateAccess := log_condition_template_access.NewApiLogConditionTemplateAccess(client)
	notificationChannelAccess := notification_channel_access.NewApiNotificationChannelAccess(client)
	policyAccess := policy_access.NewApiPolicyAccess(client)
	policyTemplateAccess := policy_template_access.NewApiPolicyTemplateAccess(client)
	tsConditionAccess := ts_condition_access.NewApiTsConditionAccess(client)
	tsConditionTemplateAccess := ts_condition_template_access.NewApiTsConditionTemplateAccess(client)
	tsEntryAccess := ts_entry_access.NewApiTsEntryAccess(client)

	return &apiAlertingAccess{
		Access: gotenresource.NewCompositeAccess(

			alert.AsAnyCastAccess(alertAccess),
			document.AsAnyCastAccess(documentAccess),
			log_condition.AsAnyCastAccess(logConditionAccess),
			log_condition_template.AsAnyCastAccess(logConditionTemplateAccess),
			notification_channel.AsAnyCastAccess(notificationChannelAccess),
			policy.AsAnyCastAccess(policyAccess),
			policy_template.AsAnyCastAccess(policyTemplateAccess),
			ts_condition.AsAnyCastAccess(tsConditionAccess),
			ts_condition_template.AsAnyCastAccess(tsConditionTemplateAccess),
			ts_entry.AsAnyCastAccess(tsEntryAccess),
		),

		AlertAccess:                alertAccess,
		DocumentAccess:             documentAccess,
		LogConditionAccess:         logConditionAccess,
		LogConditionTemplateAccess: logConditionTemplateAccess,
		NotificationChannelAccess:  notificationChannelAccess,
		PolicyAccess:               policyAccess,
		PolicyTemplateAccess:       policyTemplateAccess,
		TsConditionAccess:          tsConditionAccess,
		TsConditionTemplateAccess:  tsConditionTemplateAccess,
		TsEntryAccess:              tsEntryAccess,
	}
}
