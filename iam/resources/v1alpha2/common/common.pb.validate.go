// Code generated by protoc-gen-goten-validate
// File: edgelq/iam/proto/v1alpha2/common.proto
// DO NOT EDIT!!!

package iam_common

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
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
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
	_ = &meta_service.Service{}
)

func (obj *PCR) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ServiceBusinessTier) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if _, ok := BusinessTier_name[int32(obj.BusinessTier)]; !ok {
		return gotenvalidate.NewValidationError("ServiceBusinessTier", "businessTier", obj.BusinessTier, "field must be a defined enum value", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ServiceErrors) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Errors {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ServiceErrors", "errors", obj.Errors[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ServiceErrors_Error) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
