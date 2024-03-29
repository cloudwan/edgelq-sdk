// Code generated by protoc-gen-goten-validate
// File: edgelq/monitoring/proto/v4/metric_descriptor_custom.proto
// DO NOT EDIT!!!

package metric_descriptor_client

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
	metric_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/metric_descriptor"
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
	view "github.com/cloudwan/goten-sdk/types/view"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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
	_ = &metric_descriptor.MetricDescriptor{}
	_ = &project.Project{}
	_ = &fieldmaskpb.FieldMask{}
	_ = view.View(0)
)

func (obj *GetMetricDescriptorRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *CreateMetricDescriptorRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.MetricDescriptor == nil {
		return gotenvalidate.NewValidationError("CreateMetricDescriptorRequest", "metricDescriptor", obj.MetricDescriptor, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.MetricDescriptor).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("CreateMetricDescriptorRequest", "metricDescriptor", obj.MetricDescriptor, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateMetricDescriptorRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.MetricDescriptor == nil {
		return gotenvalidate.NewValidationError("UpdateMetricDescriptorRequest", "metricDescriptor", obj.MetricDescriptor, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.MetricDescriptor).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateMetricDescriptorRequest", "metricDescriptor", obj.MetricDescriptor, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Cas).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("UpdateMetricDescriptorRequest", "cas", obj.Cas, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *UpdateMetricDescriptorRequest_CAS) GotenValidate() error {
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
func (obj *DeleteMetricDescriptorRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListMetricDescriptorsRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if !(obj.PageSize >= 0) {
		return gotenvalidate.NewValidationError("ListMetricDescriptorsRequest", "pageSize", obj.PageSize, "field must be greater or equal to 0", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ListMetricDescriptorsResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.MetricDescriptors {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ListMetricDescriptorsResponse", "metricDescriptors", obj.MetricDescriptors[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
