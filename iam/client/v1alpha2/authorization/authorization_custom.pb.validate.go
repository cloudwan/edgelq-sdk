// Code generated by protoc-gen-goten-validate
// File: edgelq/iam/proto/v1alpha2/authorization_custom.proto
// DO NOT EDIT!!!

package authorization_client

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
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/condition"
	permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/permission"
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/role"
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
	_ = &condition.Condition{}
	_ = &permission.Permission{}
	_ = &role.Role{}
)

func (obj *Check) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ConditionalGrant) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.ConditionBindings {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ConditionalGrant", "conditionBindings", obj.ConditionBindings[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CheckResult) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.ConditionallyGrantedPermissions {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("CheckResult", "conditionallyGrantedPermissions", obj.ConditionallyGrantedPermissions[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CheckPermissionsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Checks {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("CheckPermissionsRequest", "checks", obj.Checks[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CheckPermissionsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.CheckResults {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("CheckPermissionsResponse", "checkResults", obj.CheckResults[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CheckMyPermissionsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Checks {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("CheckMyPermissionsRequest", "checks", obj.Checks[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CheckMyPermissionsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.CheckResults {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("CheckMyPermissionsResponse", "checkResults", obj.CheckResults[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CheckMyRolesRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CheckMyRolesResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.ConditionallyGrantedRoles {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("CheckMyRolesResponse", "conditionallyGrantedRoles", obj.ConditionallyGrantedRoles[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CheckMyRolesResponse_ConditionalGrant) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.ConditionBindings {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ConditionalGrant", "conditionBindings", obj.ConditionBindings[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
