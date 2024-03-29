// Code generated by protoc-gen-goten-validate
// File: edgelq/logging/proto/v1alpha2/log_descriptor.proto
// DO NOT EDIT!!!

package log_descriptor

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
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	common "github.com/cloudwan/edgelq-sdk/logging/resources/v1alpha2/common"
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
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &common.LabelDescriptor{}
	_ = &meta.Meta{}
)

func (obj *LogDescriptor) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Labels {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("LogDescriptor", "labels", obj.Labels[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.PromotedLabelKeySets {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("LogDescriptor", "promotedLabelKeySets", obj.PromotedLabelKeySets[idx], "nested object validation failed", err)
			}
		}
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("LogDescriptor", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
