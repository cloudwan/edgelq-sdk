// Code generated by protoc-gen-goten-validate
// File: edgelq/iam/proto/v1/group_service.proto
// DO NOT EDIT!!!

package group_client

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
	group "github.com/cloudwan/edgelq-sdk/iam/resources/v1/group"
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
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
	_ = &group.Group{}
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &emptypb.Empty{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &timestamppb.Timestamp{}
	_ = &meta_service.Service{}
	_ = view.View(0)
	_ = watch_type.WatchType(0)
)

func (obj *GetGroupRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *BatchGetGroupsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *BatchGetGroupsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Groups {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("BatchGetGroupsResponse", "groups", obj.Groups[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListGroupsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if !(obj.PageSize >= 0) {
		return gotenvalidate.NewValidationError("ListGroupsRequest", "pageSize", obj.PageSize, "field must be greater or equal to 0", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListGroupsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Groups {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListGroupsResponse", "groups", obj.Groups[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchGroupRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchGroupResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Change).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("WatchGroupResponse", "change", obj.Change, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchGroupsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if !(obj.PageSize >= 0) {
		return gotenvalidate.NewValidationError("WatchGroupsRequest", "pageSize", obj.PageSize, "field must be greater or equal to 0", nil)
	}
	if !(obj.MaxChunkSize >= 0 && obj.MaxChunkSize <= 100) {
		return gotenvalidate.NewValidationError("WatchGroupsRequest", "maxChunkSize", obj.MaxChunkSize, "field must be in range [0, 100]", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchGroupsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.GroupChanges {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("WatchGroupsResponse", "groupChanges", obj.GroupChanges[idx], "nested object validation failed", err)
			}
		}
	}
	if subobj, ok := interface{}(obj.PageTokenChange).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("WatchGroupsResponse", "pageTokenChange", obj.PageTokenChange, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchGroupsResponse_PageTokenChange) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CreateGroupRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.Group == nil {
		return gotenvalidate.NewValidationError("CreateGroupRequest", "group", obj.Group, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Group).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CreateGroupRequest", "group", obj.Group, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.ResponseMask).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CreateGroupRequest", "responseMask", obj.ResponseMask, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CreateGroupRequest_ResponseMask) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Masking.(type) {
	case *CreateGroupRequest_ResponseMask_SkipEntireResponseBody:
	case *CreateGroupRequest_ResponseMask_BodyMask:
	default:
		_ = opt
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateGroupRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.Group == nil {
		return gotenvalidate.NewValidationError("UpdateGroupRequest", "group", obj.Group, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Group).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateGroupRequest", "group", obj.Group, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Cas).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateGroupRequest", "cas", obj.Cas, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.ResponseMask).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateGroupRequest", "responseMask", obj.ResponseMask, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateGroupRequest_CAS) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateGroupRequest_ResponseMask) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Masking.(type) {
	case *UpdateGroupRequest_ResponseMask_SkipEntireResponseBody:
	case *UpdateGroupRequest_ResponseMask_UpdatedFieldsOnly:
	case *UpdateGroupRequest_ResponseMask_BodyMask:
	default:
		_ = opt
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *DeleteGroupRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
