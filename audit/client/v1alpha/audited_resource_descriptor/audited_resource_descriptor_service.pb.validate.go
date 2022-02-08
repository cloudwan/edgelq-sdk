// Code generated by protoc-gen-goten-validate
// File: edgelq/audit/proto/v1alpha/audited_resource_descriptor_service.proto
// DO NOT EDIT!!!

package audited_resource_descriptor_client

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
	audited_resource_descriptor "github.com/cloudwan/edgelq-sdk/audit/resources/v1alpha/audited_resource_descriptor"
	view "github.com/cloudwan/goten-sdk/runtime/api/view"
	watch_type "github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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
	_ = &audited_resource_descriptor.AuditedResourceDescriptor{}
	_ = &field_mask.FieldMask{}
	_ = view.View(0)
	_ = watch_type.WatchType(0)
)

func (obj *GetAuditedResourceDescriptorRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *BatchGetAuditedResourceDescriptorsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *BatchGetAuditedResourceDescriptorsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.AuditedResourceDescriptors {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("BatchGetAuditedResourceDescriptorsResponse", "auditedResourceDescriptors", obj.AuditedResourceDescriptors[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListAuditedResourceDescriptorsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if !(obj.PageSize >= 0) {
		return gotenvalidate.NewValidationError("ListAuditedResourceDescriptorsRequest", "pageSize", obj.PageSize, "field must be greater or equal to 0", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListAuditedResourceDescriptorsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.AuditedResourceDescriptors {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListAuditedResourceDescriptorsResponse", "auditedResourceDescriptors", obj.AuditedResourceDescriptors[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchAuditedResourceDescriptorRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchAuditedResourceDescriptorResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Change).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("WatchAuditedResourceDescriptorResponse", "change", obj.Change, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchAuditedResourceDescriptorsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if !(obj.PageSize >= 0) {
		return gotenvalidate.NewValidationError("WatchAuditedResourceDescriptorsRequest", "pageSize", obj.PageSize, "field must be greater or equal to 0", nil)
	}
	if !(obj.MaxChunkSize >= 0 && obj.MaxChunkSize <= 100) {
		return gotenvalidate.NewValidationError("WatchAuditedResourceDescriptorsRequest", "maxChunkSize", obj.MaxChunkSize, "field must be in range [0, 100]", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchAuditedResourceDescriptorsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.AuditedResourceDescriptorChanges {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("WatchAuditedResourceDescriptorsResponse", "auditedResourceDescriptorChanges", obj.AuditedResourceDescriptorChanges[idx], "nested object validation failed", err)
			}
		}
	}
	if subobj, ok := interface{}(obj.PageTokenChange).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("WatchAuditedResourceDescriptorsResponse", "pageTokenChange", obj.PageTokenChange, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchAuditedResourceDescriptorsResponse_PageTokenChange) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CreateAuditedResourceDescriptorRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.AuditedResourceDescriptor == nil {
		return gotenvalidate.NewValidationError("CreateAuditedResourceDescriptorRequest", "auditedResourceDescriptor", obj.AuditedResourceDescriptor, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.AuditedResourceDescriptor).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CreateAuditedResourceDescriptorRequest", "auditedResourceDescriptor", obj.AuditedResourceDescriptor, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateAuditedResourceDescriptorRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.AuditedResourceDescriptor == nil {
		return gotenvalidate.NewValidationError("UpdateAuditedResourceDescriptorRequest", "auditedResourceDescriptor", obj.AuditedResourceDescriptor, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.AuditedResourceDescriptor).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateAuditedResourceDescriptorRequest", "auditedResourceDescriptor", obj.AuditedResourceDescriptor, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Cas).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateAuditedResourceDescriptorRequest", "cas", obj.Cas, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateAuditedResourceDescriptorRequest_CAS) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.ConditionalState).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CAS", "conditionalState", obj.ConditionalState, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
