// Code generated by protoc-gen-goten-validate
// File: edgelq/alerting/proto/v1/policy_template_change.proto
// DO NOT EDIT!!!

package policy_template

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
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
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
	_ = &iam_project.Project{}
	_ = &fieldmaskpb.FieldMask{}
)

func (obj *PolicyTemplateChange) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.ChangeType.(type) {
	case *PolicyTemplateChange_Added_:
		if subobj, ok := interface{}(opt.Added).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("PolicyTemplateChange", "added", opt.Added, "nested object validation failed", err)
			}
		}
	case *PolicyTemplateChange_Modified_:
		if subobj, ok := interface{}(opt.Modified).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("PolicyTemplateChange", "modified", opt.Modified, "nested object validation failed", err)
			}
		}
	case *PolicyTemplateChange_Current_:
		if subobj, ok := interface{}(opt.Current).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("PolicyTemplateChange", "current", opt.Current, "nested object validation failed", err)
			}
		}
	case *PolicyTemplateChange_Removed_:
		if subobj, ok := interface{}(opt.Removed).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("PolicyTemplateChange", "removed", opt.Removed, "nested object validation failed", err)
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
func (obj *PolicyTemplateChange_Added) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.PolicyTemplate).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Added", "policyTemplate", obj.PolicyTemplate, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *PolicyTemplateChange_Modified) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.PolicyTemplate).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Modified", "policyTemplate", obj.PolicyTemplate, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *PolicyTemplateChange_Current) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.PolicyTemplate).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Current", "policyTemplate", obj.PolicyTemplate, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *PolicyTemplateChange_Removed) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
