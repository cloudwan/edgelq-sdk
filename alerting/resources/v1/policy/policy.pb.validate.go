// Code generated by protoc-gen-goten-validate
// File: edgelq/alerting/proto/v1/policy.proto
// DO NOT EDIT!!!

package policy

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	gotenvalidate "github.com/cloudwan/goten-sdk/runtime/validate"
)

// proto imports
import (
	rcommon "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/common"
	document "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/document"
	notification_channel "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/notification_channel"
	policy_template "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/policy_template"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Errorf
	_ = net.ParseIP
	_ = regexp.Match
	_ = strings.Split
	_ = time.Now
	_ = utf8.RuneCountInString
	_ = url.Parse
	_ = gotenvalidate.NewValidationError
)

// make sure we're using proto imports
var (
	_ = &document.Document{}
	_ = &notification_channel.NotificationChannel{}
	_ = &policy_template.PolicyTemplate{}
	_ = &rcommon.LogCndSpec{}
	_ = &iam_project.Project{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &meta.Meta{}
)

func (obj *Policy) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Policy", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	{
		rlen := utf8.RuneCountInString(obj.DisplayName)
		if rlen > 256 {
			return gotenvalidate.NewValidationError("Policy", "displayName", obj.DisplayName, "field must contain at most 256 characters", nil)
		}
	}
	{
		rlen := utf8.RuneCountInString(obj.Description)
		if rlen > 512 {
			return gotenvalidate.NewValidationError("Policy", "description", obj.Description, "field must contain at most 512 characters", nil)
		}
	}
	if obj.Spec == nil {
		return gotenvalidate.NewValidationError("Policy", "spec", obj.Spec, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Spec).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Policy", "spec", obj.Spec, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.TemplateSource).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Policy", "templateSource", obj.TemplateSource, "nested object validation failed", err)
		}
	}
	if len(obj.NotificationChannels) > 8 {
		return gotenvalidate.NewValidationError("Policy", "notificationChannels", obj.NotificationChannels, "field must have at most 8 items", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Policy_TemplateSource) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
