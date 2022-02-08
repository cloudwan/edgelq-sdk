// Code generated by protoc-gen-goten-validate
// File: edgelq/audit/proto/v1alpha/activity_log_custom.proto
// DO NOT EDIT!!!

package activity_log_client

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

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	gotenvalidate "github.com/cloudwan/goten-sdk/runtime/validate"
)

// proto imports
import (
	audit_common "github.com/cloudwan/edgelq-sdk/audit/common/v1alpha"
	activity_log "github.com/cloudwan/edgelq-sdk/audit/resources/v1alpha/activity_log"
	rpc "github.com/cloudwan/edgelq-sdk/common/rpc"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
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
	_ = durationpb.Duration{}
	_ = timestamppb.Timestamp{}
	_ = gotenvalidate.NewValidationError
)

// make sure we're using proto imports
var (
	_ = &activity_log.ActivityLog{}
	_ = &audit_common.Authentication{}
	_ = &rpc.Status{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
)

func (obj *ListActivityLogsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.Interval == nil {
		return gotenvalidate.NewValidationError("ListActivityLogsRequest", "interval", obj.Interval, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Interval).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ListActivityLogsRequest", "interval", obj.Interval, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListActivityLogsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.ActivityLogs {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListActivityLogsResponse", "activityLogs", obj.ActivityLogs[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.ExecutionErrors {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListActivityLogsResponse", "executionErrors", obj.ExecutionErrors[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CreateActivityLogsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.ActivityLogs {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("CreateActivityLogsRequest", "activityLogs", obj.ActivityLogs[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CreateActivityLogsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
