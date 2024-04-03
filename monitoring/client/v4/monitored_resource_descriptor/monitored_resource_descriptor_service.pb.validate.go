// Code generated by protoc-gen-goten-validate
// File: edgelq/monitoring/proto/v4/monitored_resource_descriptor_service.proto
// DO NOT EDIT!!!

package monitored_resource_descriptor_client

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
	monitored_resource_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/monitored_resource_descriptor"
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
	_ = &monitored_resource_descriptor.MonitoredResourceDescriptor{}
	_ = &emptypb.Empty{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &timestamppb.Timestamp{}
	_ = &meta_service.Service{}
	_ = view.View(0)
	_ = watch_type.WatchType(0)
)

func (obj *BatchGetMonitoredResourceDescriptorsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *BatchGetMonitoredResourceDescriptorsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.MonitoredResourceDescriptors {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("BatchGetMonitoredResourceDescriptorsResponse", "monitoredResourceDescriptors", obj.MonitoredResourceDescriptors[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchMonitoredResourceDescriptorRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchMonitoredResourceDescriptorResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Change).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("WatchMonitoredResourceDescriptorResponse", "change", obj.Change, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchMonitoredResourceDescriptorsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if !(obj.PageSize >= 0) {
		return gotenvalidate.NewValidationError("WatchMonitoredResourceDescriptorsRequest", "pageSize", obj.PageSize, "field must be greater or equal to 0", nil)
	}
	if !(obj.MaxChunkSize >= 0 && obj.MaxChunkSize <= 100) {
		return gotenvalidate.NewValidationError("WatchMonitoredResourceDescriptorsRequest", "maxChunkSize", obj.MaxChunkSize, "field must be in range [0, 100]", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchMonitoredResourceDescriptorsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.MonitoredResourceDescriptorChanges {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("WatchMonitoredResourceDescriptorsResponse", "monitoredResourceDescriptorChanges", obj.MonitoredResourceDescriptorChanges[idx], "nested object validation failed", err)
			}
		}
	}
	if subobj, ok := interface{}(obj.PageTokenChange).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("WatchMonitoredResourceDescriptorsResponse", "pageTokenChange", obj.PageTokenChange, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchMonitoredResourceDescriptorsResponse_PageTokenChange) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CreateMonitoredResourceDescriptorRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.MonitoredResourceDescriptor == nil {
		return gotenvalidate.NewValidationError("CreateMonitoredResourceDescriptorRequest", "monitoredResourceDescriptor", obj.MonitoredResourceDescriptor, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.MonitoredResourceDescriptor).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CreateMonitoredResourceDescriptorRequest", "monitoredResourceDescriptor", obj.MonitoredResourceDescriptor, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateMonitoredResourceDescriptorRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.MonitoredResourceDescriptor == nil {
		return gotenvalidate.NewValidationError("UpdateMonitoredResourceDescriptorRequest", "monitoredResourceDescriptor", obj.MonitoredResourceDescriptor, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.MonitoredResourceDescriptor).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateMonitoredResourceDescriptorRequest", "monitoredResourceDescriptor", obj.MonitoredResourceDescriptor, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Cas).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateMonitoredResourceDescriptorRequest", "cas", obj.Cas, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateMonitoredResourceDescriptorRequest_CAS) GotenValidate() error {
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
func (obj *DeleteMonitoredResourceDescriptorRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}