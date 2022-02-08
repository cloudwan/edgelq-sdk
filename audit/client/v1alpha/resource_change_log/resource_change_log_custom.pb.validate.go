// Code generated by protoc-gen-goten-validate
// File: edgelq/audit/proto/v1alpha/resource_change_log_custom.proto
// DO NOT EDIT!!!

package resource_change_log_client

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
	resource_change_log "github.com/cloudwan/edgelq-sdk/audit/resources/v1alpha/resource_change_log"
	rpc "github.com/cloudwan/edgelq-sdk/common/rpc"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	_ = &audit_common.Authentication{}
	_ = &resource_change_log.ResourceChangeLog{}
	_ = &rpc.Status{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &timestamp.Timestamp{}
)

func (obj *ListResourceChangeLogsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.Interval == nil {
		return gotenvalidate.NewValidationError("ListResourceChangeLogsRequest", "interval", obj.Interval, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Interval).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ListResourceChangeLogsRequest", "interval", obj.Interval, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListResourceChangeLogsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.ResourceChangeLogs {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListResourceChangeLogsResponse", "resourceChangeLogs", obj.ResourceChangeLogs[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.ExecutionErrors {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListResourceChangeLogsResponse", "executionErrors", obj.ExecutionErrors[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CreatePreCommittedResourceChangeLogsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Authentication).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CreatePreCommittedResourceChangeLogsRequest", "authentication", obj.Authentication, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Service).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CreatePreCommittedResourceChangeLogsRequest", "service", obj.Service, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Transaction).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CreatePreCommittedResourceChangeLogsRequest", "transaction", obj.Transaction, "nested object validation failed", err)
		}
	}
	for idx, elem := range obj.Changes {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("CreatePreCommittedResourceChangeLogsRequest", "changes", obj.Changes[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CreatePreCommittedResourceChangeLogsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *SetResourceChangeLogsCommitStateRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *SetResourceChangeLogsCommitStateResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
