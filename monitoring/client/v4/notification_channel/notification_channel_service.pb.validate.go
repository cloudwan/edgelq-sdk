// Code generated by protoc-gen-goten-validate
// File: edgelq/monitoring/proto/v4/notification_channel_service.proto
// DO NOT EDIT!!!

package notification_channel_client

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
	notification_channel "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/notification_channel"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
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
	_ = &notification_channel.NotificationChannel{}
	_ = &project.Project{}
	_ = &emptypb.Empty{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &timestamppb.Timestamp{}
	_ = view.View(0)
	_ = watch_type.WatchType(0)
)

func (obj *GetNotificationChannelRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *BatchGetNotificationChannelsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *BatchGetNotificationChannelsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.NotificationChannels {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("BatchGetNotificationChannelsResponse", "notificationChannels", obj.NotificationChannels[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListNotificationChannelsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if !(obj.PageSize >= 0) {
		return gotenvalidate.NewValidationError("ListNotificationChannelsRequest", "pageSize", obj.PageSize, "field must be greater or equal to 0", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListNotificationChannelsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.NotificationChannels {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListNotificationChannelsResponse", "notificationChannels", obj.NotificationChannels[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchNotificationChannelRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchNotificationChannelResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Change).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("WatchNotificationChannelResponse", "change", obj.Change, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchNotificationChannelsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if !(obj.PageSize >= 0) {
		return gotenvalidate.NewValidationError("WatchNotificationChannelsRequest", "pageSize", obj.PageSize, "field must be greater or equal to 0", nil)
	}
	if !(obj.MaxChunkSize >= 0 && obj.MaxChunkSize <= 100) {
		return gotenvalidate.NewValidationError("WatchNotificationChannelsRequest", "maxChunkSize", obj.MaxChunkSize, "field must be in range [0, 100]", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchNotificationChannelsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.NotificationChannelChanges {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("WatchNotificationChannelsResponse", "notificationChannelChanges", obj.NotificationChannelChanges[idx], "nested object validation failed", err)
			}
		}
	}
	if subobj, ok := interface{}(obj.PageTokenChange).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("WatchNotificationChannelsResponse", "pageTokenChange", obj.PageTokenChange, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *WatchNotificationChannelsResponse_PageTokenChange) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CreateNotificationChannelRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.NotificationChannel == nil {
		return gotenvalidate.NewValidationError("CreateNotificationChannelRequest", "notificationChannel", obj.NotificationChannel, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.NotificationChannel).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CreateNotificationChannelRequest", "notificationChannel", obj.NotificationChannel, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.ResponseMask).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CreateNotificationChannelRequest", "responseMask", obj.ResponseMask, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CreateNotificationChannelRequest_ResponseMask) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Masking.(type) {
	case *CreateNotificationChannelRequest_ResponseMask_SkipEntireResponseBody:
	case *CreateNotificationChannelRequest_ResponseMask_BodyMask:
	default:
		_ = opt
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateNotificationChannelRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.NotificationChannel == nil {
		return gotenvalidate.NewValidationError("UpdateNotificationChannelRequest", "notificationChannel", obj.NotificationChannel, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.NotificationChannel).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateNotificationChannelRequest", "notificationChannel", obj.NotificationChannel, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Cas).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateNotificationChannelRequest", "cas", obj.Cas, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.ResponseMask).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateNotificationChannelRequest", "responseMask", obj.ResponseMask, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateNotificationChannelRequest_CAS) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateNotificationChannelRequest_ResponseMask) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Masking.(type) {
	case *UpdateNotificationChannelRequest_ResponseMask_SkipEntireResponseBody:
	case *UpdateNotificationChannelRequest_ResponseMask_UpdatedFieldsOnly:
	case *UpdateNotificationChannelRequest_ResponseMask_BodyMask:
	default:
		_ = opt
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *DeleteNotificationChannelRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
