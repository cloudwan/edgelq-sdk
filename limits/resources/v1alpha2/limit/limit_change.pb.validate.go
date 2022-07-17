// Code generated by protoc-gen-goten-validate
// File: edgelq/limits/proto/v1alpha2/limit_change.proto
// DO NOT EDIT!!!

package limit

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
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
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
	_ = &iam_project.Project{}
	_ = &field_mask.FieldMask{}
)

func (obj *LimitChange) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.ChangeType.(type) {
	case *LimitChange_Added_:
		if subobj, ok := interface{}(opt.Added).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("LimitChange", "added", opt.Added, "nested object validation failed", err)
			}
		}
	case *LimitChange_Modified_:
		if subobj, ok := interface{}(opt.Modified).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("LimitChange", "modified", opt.Modified, "nested object validation failed", err)
			}
		}
	case *LimitChange_Current_:
		if subobj, ok := interface{}(opt.Current).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("LimitChange", "current", opt.Current, "nested object validation failed", err)
			}
		}
	case *LimitChange_Removed_:
		if subobj, ok := interface{}(opt.Removed).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("LimitChange", "removed", opt.Removed, "nested object validation failed", err)
			}
		}
	default:
		_ = opt
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *LimitChange_Added) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Limit).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Added", "limit", obj.Limit, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *LimitChange_Modified) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Limit).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Modified", "limit", obj.Limit, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *LimitChange_Current) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Limit).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Current", "limit", obj.Limit, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *LimitChange_Removed) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
