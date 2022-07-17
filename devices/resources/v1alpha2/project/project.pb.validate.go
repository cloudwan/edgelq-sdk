// Code generated by protoc-gen-goten-validate
// File: edgelq/devices/proto/v1alpha2/project.proto
// DO NOT EDIT!!!

package project

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
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	multi_region_policy "github.com/cloudwan/edgelq-sdk/common/types/multi_region_policy"
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
	_ = &ntt_meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

func (obj *Project) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Project", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	if obj.MultiRegionPolicy == nil {
		return gotenvalidate.NewValidationError("Project", "multiRegionPolicy", obj.MultiRegionPolicy, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.MultiRegionPolicy).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Project", "multiRegionPolicy", obj.MultiRegionPolicy, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
