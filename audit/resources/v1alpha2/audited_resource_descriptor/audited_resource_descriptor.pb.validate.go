// Code generated by protoc-gen-goten-validate
// File: edgelq/audit/proto/v1alpha2/audited_resource_descriptor.proto
// DO NOT EDIT!!!

package audited_resource_descriptor

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
	audit_common "github.com/cloudwan/edgelq-sdk/audit/common/v1alpha2"
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
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
	_ = &audit_common.Authentication{}
	_ = &ntt_meta.Meta{}
)

func (obj *AuditedResourceDescriptor) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Labels {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("AuditedResourceDescriptor", "labels", obj.Labels[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.PromotedLabelKeySets {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("AuditedResourceDescriptor", "promotedLabelKeySets", obj.PromotedLabelKeySets[idx], "nested object validation failed", err)
			}
		}
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("AuditedResourceDescriptor", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
