// Code generated by protoc-gen-goten-validate
// File: edgelq/limits/proto/v1alpha2/accepted_plan.proto
// DO NOT EDIT!!!

package accepted_plan

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
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	common "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/common"
	plan "github.com/cloudwan/edgelq-sdk/limits/resources/v1alpha2/plan"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
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
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &common.Allowance{}
	_ = &plan.Plan{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

func (obj *AcceptedPlan) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Extensions {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("AcceptedPlan", "extensions", obj.Extensions[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.RegionalDistributions {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("AcceptedPlan", "regionalDistributions", obj.RegionalDistributions[idx], "nested object validation failed", err)
			}
		}
	}
	if subobj, ok := interface{}(obj.Assignee).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("AcceptedPlan", "assignee", obj.Assignee, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("AcceptedPlan", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *AcceptedPlan_Assignee) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Assignee.(type) {
	case *AcceptedPlan_Assignee_ProjectAssignee:
	case *AcceptedPlan_Assignee_OrganizationAssignee:
	case *AcceptedPlan_Assignee_SystemAssignee:
	default:
		_ = opt
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
