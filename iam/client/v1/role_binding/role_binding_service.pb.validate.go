// Code generated by protoc-gen-goten-validate
// File: edgelq/iam/proto/v1/role_binding_service.proto
// DO NOT EDIT!!!

package role_binding_client

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
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	role_binding "github.com/cloudwan/edgelq-sdk/iam/resources/v1/role_binding"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	view "github.com/cloudwan/goten-sdk/types/view"
	watch_type "github.com/cloudwan/goten-sdk/types/watch_type"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &role_binding.RoleBinding{}
	_ = &emptypb.Empty{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &timestamppb.Timestamp{}
	_ = &meta_service.Service{}
	_ = view.View(0)
	_ = watch_type.WatchType(0)
)

func (obj *GetRoleBindingRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *BatchGetRoleBindingsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *BatchGetRoleBindingsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.RoleBindings {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("BatchGetRoleBindingsResponse", "roleBindings", obj.RoleBindings[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListRoleBindingsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if !(obj.PageSize >= 0) {
		return gotenvalidate.NewValidationError("ListRoleBindingsRequest", "pageSize", obj.PageSize, "field must be greater or equal to 0", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListRoleBindingsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.RoleBindings {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListRoleBindingsResponse", "roleBindings", obj.RoleBindings[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchRoleBindingRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchRoleBindingResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Change).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("WatchRoleBindingResponse", "change", obj.Change, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchRoleBindingsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if !(obj.PageSize >= 0) {
		return gotenvalidate.NewValidationError("WatchRoleBindingsRequest", "pageSize", obj.PageSize, "field must be greater or equal to 0", nil)
	}
	if !(obj.MaxChunkSize >= 0 && obj.MaxChunkSize <= 100) {
		return gotenvalidate.NewValidationError("WatchRoleBindingsRequest", "maxChunkSize", obj.MaxChunkSize, "field must be in range [0, 100]", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchRoleBindingsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.RoleBindingChanges {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("WatchRoleBindingsResponse", "roleBindingChanges", obj.RoleBindingChanges[idx], "nested object validation failed", err)
			}
		}
	}
	if subobj, ok := interface{}(obj.PageTokenChange).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("WatchRoleBindingsResponse", "pageTokenChange", obj.PageTokenChange, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchRoleBindingsResponse_PageTokenChange) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CreateRoleBindingRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.RoleBinding == nil {
		return gotenvalidate.NewValidationError("CreateRoleBindingRequest", "roleBinding", obj.RoleBinding, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.RoleBinding).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CreateRoleBindingRequest", "roleBinding", obj.RoleBinding, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.ResponseMask).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CreateRoleBindingRequest", "responseMask", obj.ResponseMask, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CreateRoleBindingRequest_ResponseMask) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Masking.(type) {
	case *CreateRoleBindingRequest_ResponseMask_SkipEntireResponseBody:
	case *CreateRoleBindingRequest_ResponseMask_BodyMask:
	default:
		_ = opt
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateRoleBindingRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.RoleBinding == nil {
		return gotenvalidate.NewValidationError("UpdateRoleBindingRequest", "roleBinding", obj.RoleBinding, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.RoleBinding).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateRoleBindingRequest", "roleBinding", obj.RoleBinding, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Cas).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateRoleBindingRequest", "cas", obj.Cas, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.ResponseMask).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateRoleBindingRequest", "responseMask", obj.ResponseMask, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateRoleBindingRequest_CAS) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateRoleBindingRequest_ResponseMask) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Masking.(type) {
	case *UpdateRoleBindingRequest_ResponseMask_SkipEntireResponseBody:
	case *UpdateRoleBindingRequest_ResponseMask_UpdatedFieldsOnly:
	case *UpdateRoleBindingRequest_ResponseMask_BodyMask:
	default:
		_ = opt
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *DeleteRoleBindingRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
