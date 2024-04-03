// Code generated by protoc-gen-goten-validate
// File: edgelq/devices/proto/v1/device_hardware_register_session.proto
// DO NOT EDIT!!!

package device_hardware_register_session

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
	device_hardware "github.com/cloudwan/edgelq-sdk/devices/resources/v1/device_hardware"
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1/project"
	provisioning_policy "github.com/cloudwan/edgelq-sdk/devices/resources/v1/provisioning_policy"
	meta "github.com/cloudwan/goten-sdk/types/meta"
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
	_ = &device_hardware.DeviceHardware{}
	_ = &project.Project{}
	_ = &provisioning_policy.ProvisioningPolicy{}
	_ = &timestamppb.Timestamp{}
	_ = &meta.Meta{}
)

func (obj *DeviceHardwareRegisterSession) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("DeviceHardwareRegisterSession", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	if err := gotenvalidate.ValidateEmail(string(obj.UserEmail)); err != nil {
		return gotenvalidate.NewValidationError("DeviceHardwareRegisterSession", "userEmail", obj.UserEmail, "field must contain a valid email address", err)
	}
	if obj.UserEmail == "" {
		return gotenvalidate.NewValidationError("DeviceHardwareRegisterSession", "userEmail", obj.UserEmail, "field is required", nil)
	}
	if obj.LanguageCode != "" && obj.LanguageCode != "en-us" && obj.LanguageCode != "ja-jp" {
		return gotenvalidate.NewValidationError("DeviceHardwareRegisterSession", "languageCode", obj.LanguageCode, "field must be equal to exactly one of the following values: , en-us, ja-jp", nil)
	}
	if subobj, ok := interface{}(obj.Status).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("DeviceHardwareRegisterSession", "status", obj.Status, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *DeviceHardwareRegisterSession_Status) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}