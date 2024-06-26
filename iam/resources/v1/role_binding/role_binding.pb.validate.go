// Code generated by protoc-gen-goten-validate
// File: edgelq/iam/proto/v1/role_binding.proto
// DO NOT EDIT!!!

package role_binding

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
	condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1/condition"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	role "github.com/cloudwan/edgelq-sdk/iam/resources/v1/role"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
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
	_ = &condition.Condition{}
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &role.Role{}
	_ = &meta_service.Service{}
	_ = &meta.Meta{}
)

func (obj *RoleBinding) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("RoleBinding", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	if obj.Member == "" {
		return gotenvalidate.NewValidationError("RoleBinding", "member", obj.Member, "field is required", nil)
	}
	for idx, elem := range obj.ScopeParams {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("RoleBinding", "scopeParams", obj.ScopeParams[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.ExecutableConditions {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("RoleBinding", "executableConditions", obj.ExecutableConditions[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.AncestryPath {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("RoleBinding", "ancestryPath", obj.AncestryPath[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *RoleBinding_Parent) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if len(obj.Member) > 256 {
		return gotenvalidate.NewValidationError("Parent", "member", obj.Member, "field must contain at most 256 characters", nil)
	}
	if obj.Member == "" {
		return gotenvalidate.NewValidationError("Parent", "member", obj.Member, "field is required", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
