// Code generated by protoc-gen-goten-validate
// File: edgelq/devices/proto/v1/device_service.proto
// DO NOT EDIT!!!

package device_client

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
	device "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1/project"
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
	_ = &device.Device{}
	_ = &project.Project{}
	_ = &emptypb.Empty{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &timestamppb.Timestamp{}
	_ = view.View(0)
	_ = watch_type.WatchType(0)
)

func (obj *GetDeviceRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *BatchGetDevicesRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *BatchGetDevicesResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Devices {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("BatchGetDevicesResponse", "devices", obj.Devices[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListDevicesRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if !(obj.PageSize >= 0) {
		return gotenvalidate.NewValidationError("ListDevicesRequest", "pageSize", obj.PageSize, "field must be greater or equal to 0", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListDevicesResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Devices {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListDevicesResponse", "devices", obj.Devices[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchDeviceRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchDeviceResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Change).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("WatchDeviceResponse", "change", obj.Change, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchDevicesRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if !(obj.PageSize >= 0) {
		return gotenvalidate.NewValidationError("WatchDevicesRequest", "pageSize", obj.PageSize, "field must be greater or equal to 0", nil)
	}
	if !(obj.MaxChunkSize >= 0 && obj.MaxChunkSize <= 100) {
		return gotenvalidate.NewValidationError("WatchDevicesRequest", "maxChunkSize", obj.MaxChunkSize, "field must be in range [0, 100]", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchDevicesResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.DeviceChanges {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("WatchDevicesResponse", "deviceChanges", obj.DeviceChanges[idx], "nested object validation failed", err)
			}
		}
	}
	if subobj, ok := interface{}(obj.PageTokenChange).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("WatchDevicesResponse", "pageTokenChange", obj.PageTokenChange, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchDevicesResponse_PageTokenChange) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CreateDeviceRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.Device == nil {
		return gotenvalidate.NewValidationError("CreateDeviceRequest", "device", obj.Device, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Device).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CreateDeviceRequest", "device", obj.Device, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.ResponseMask).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CreateDeviceRequest", "responseMask", obj.ResponseMask, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CreateDeviceRequest_ResponseMask) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Masking.(type) {
	case *CreateDeviceRequest_ResponseMask_SkipEntireResponseBody:
	case *CreateDeviceRequest_ResponseMask_BodyMask:
	default:
		_ = opt
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateDeviceRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.Device == nil {
		return gotenvalidate.NewValidationError("UpdateDeviceRequest", "device", obj.Device, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Device).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateDeviceRequest", "device", obj.Device, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Cas).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateDeviceRequest", "cas", obj.Cas, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.ResponseMask).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateDeviceRequest", "responseMask", obj.ResponseMask, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateDeviceRequest_CAS) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateDeviceRequest_ResponseMask) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Masking.(type) {
	case *UpdateDeviceRequest_ResponseMask_SkipEntireResponseBody:
	case *UpdateDeviceRequest_ResponseMask_UpdatedFieldsOnly:
	case *UpdateDeviceRequest_ResponseMask_BodyMask:
	default:
		_ = opt
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *DeleteDeviceRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
