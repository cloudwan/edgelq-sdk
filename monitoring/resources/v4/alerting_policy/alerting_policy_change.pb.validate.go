// Code generated by protoc-gen-goten-validate
// File: edgelq/monitoring/proto/v4/alerting_policy_change.proto
// DO NOT EDIT!!!

package alerting_policy

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
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
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
	_ = &project.Project{}
	_ = &fieldmaskpb.FieldMask{}
)

func (obj *AlertingPolicyChange) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.ChangeType.(type) {
	case *AlertingPolicyChange_Added_:
		if subobj, ok := interface{}(opt.Added).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("AlertingPolicyChange", "added", opt.Added, "nested object validation failed", err)
			}
		}
	case *AlertingPolicyChange_Modified_:
		if subobj, ok := interface{}(opt.Modified).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("AlertingPolicyChange", "modified", opt.Modified, "nested object validation failed", err)
			}
		}
	case *AlertingPolicyChange_Current_:
		if subobj, ok := interface{}(opt.Current).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("AlertingPolicyChange", "current", opt.Current, "nested object validation failed", err)
			}
		}
	case *AlertingPolicyChange_Removed_:
		if subobj, ok := interface{}(opt.Removed).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("AlertingPolicyChange", "removed", opt.Removed, "nested object validation failed", err)
			}
		}
	default:
		_ = opt
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *AlertingPolicyChange_Added) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.AlertingPolicy).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Added", "alertingPolicy", obj.AlertingPolicy, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *AlertingPolicyChange_Modified) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.AlertingPolicy).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Modified", "alertingPolicy", obj.AlertingPolicy, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *AlertingPolicyChange_Current) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.AlertingPolicy).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Current", "alertingPolicy", obj.AlertingPolicy, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *AlertingPolicyChange_Removed) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}