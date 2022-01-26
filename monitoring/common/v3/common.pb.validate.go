// Code generated by protoc-gen-goten-validate
// File: edgelq/monitoring/proto/v3/common.proto
// DO NOT EDIT!!!

package monitoring_common

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
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

	validation_regex_LabelDescriptor_key_3767d375e69c51f9742b2f17e9eb21a2 = regexp.MustCompile("^[_A-Za-z][_A-Za-z0-9]{0,31}$")
	validation_regex_Metric_type_e8902cb4d938f84496a381cb8dc140e5         = regexp.MustCompile("^[A-Za-z0-9_\\-.,+!*()%\\\\/]{1,256}$")
)

// make sure we're using proto imports
var (
	_ = &duration.Duration{}
	_ = &timestamp.Timestamp{}
)

func (obj *LabelDescriptor) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if !validation_regex_LabelDescriptor_key_3767d375e69c51f9742b2f17e9eb21a2.Match([]byte(obj.Key)) {
		return gotenvalidate.NewValidationError("LabelDescriptor", "key", obj.Key, "field must match the regex ^[_A-Za-z][_A-Za-z0-9]{0,31}$", nil)
	}
	if int32(obj.ValueType) != 0 {
		return gotenvalidate.NewValidationError("LabelDescriptor", "valueType", obj.ValueType, "field must be set to the value 0", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *LabelKeySet) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Distribution) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Range).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Distribution", "range", obj.Range, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.BucketOptions).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Distribution", "bucketOptions", obj.BucketOptions, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Distribution_Range) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Distribution_BucketOptions) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Options.(type) {
	case *Distribution_BucketOptions_LinearBuckets:
		if subobj, ok := interface{}(opt.LinearBuckets).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("BucketOptions", "linearBuckets", opt.LinearBuckets, "nested object validation failed", err)
			}
		}
	case *Distribution_BucketOptions_ExponentialBuckets:
		if subobj, ok := interface{}(opt.ExponentialBuckets).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("BucketOptions", "exponentialBuckets", opt.ExponentialBuckets, "nested object validation failed", err)
			}
		}
	case *Distribution_BucketOptions_ExplicitBuckets:
		if subobj, ok := interface{}(opt.ExplicitBuckets).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("BucketOptions", "explicitBuckets", opt.ExplicitBuckets, "nested object validation failed", err)
			}
		}
	case *Distribution_BucketOptions_DynamicBuckets:
		if subobj, ok := interface{}(opt.DynamicBuckets).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("BucketOptions", "dynamicBuckets", opt.DynamicBuckets, "nested object validation failed", err)
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
func (obj *Distribution_BucketOptions_Linear) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Distribution_BucketOptions_Exponential) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Distribution_BucketOptions_Explicit) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Distribution_BucketOptions_Dynamic) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *TypedValue) GotenValidate() error {
	if obj == nil {
		return nil
	}
	switch opt := obj.Value.(type) {
	case *TypedValue_BoolValue:
	case *TypedValue_Int64Value:
	case *TypedValue_DoubleValue:
	case *TypedValue_StringValue:
	case *TypedValue_DistributionValue:
		if subobj, ok := interface{}(opt.DistributionValue).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("TypedValue", "distributionValue", opt.DistributionValue, "nested object validation failed", err)
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
func (obj *TimeInterval) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.EndTime == nil {
		return gotenvalidate.NewValidationError("TimeInterval", "endTime", obj.EndTime, "field is required", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *TimeRange) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Aggregation) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.AlignmentPeriod != nil && obj.AlignmentPeriod.CheckValid() != nil {
		err := obj.AlignmentPeriod.CheckValid()
		return gotenvalidate.NewValidationError("Aggregation", "alignmentPeriod", obj.AlignmentPeriod, "could not validate duration", err)
	} else {
		d := obj.AlignmentPeriod.AsDuration()

		if obj.AlignmentPeriod != nil {
			if d != time.Duration(0) && d != time.Duration(60000000000) && d != time.Duration(180000000000) && d != time.Duration(300000000000) && d != time.Duration(900000000000) && d != time.Duration(1800000000000) && d != time.Duration(3600000000000) && d != time.Duration(10800000000000) && d != time.Duration(21600000000000) && d != time.Duration(43200000000000) && d != time.Duration(86400000000000) && d != time.Duration(604800000000000) && d != time.Duration(2419200000000000) {
				return gotenvalidate.NewValidationError("Aggregation", "alignmentPeriod", d, "field must be equal to exactly one of the following values: 0s, 1m0s, 3m0s, 5m0s, 15m0s, 30m0s, 1h0m0s, 3h0m0s, 6h0m0s, 12h0m0s, 24h0m0s, 168h0m0s, 672h0m0s", nil)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Metric) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if !validation_regex_Metric_type_e8902cb4d938f84496a381cb8dc140e5.Match([]byte(obj.Type)) {
		return gotenvalidate.NewValidationError("Metric", "type", obj.Type, "field must match the regex ^[A-Za-z0-9_\\-.,+!*()%\\\\/]{1,256}$", nil)
	}
	for k, v := range obj.Labels {
		_, _ = k, v
		{
			rlen := utf8.RuneCountInString(v)
			if rlen > 1024 {
				return gotenvalidate.NewValidationError("Metric", "labels", v, "field must contain at most 1024 characters", nil)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *MonitoredResource) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *MonitoredResourceMetadata) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *Strings) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *MonitoredResourceSelector) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if len(obj.Types) < 1 {
		return gotenvalidate.NewValidationError("MonitoredResourceSelector", "types", obj.Types, "field must have at least 1 items", nil)
	}
	if len(obj.Types) > 1 {
		return gotenvalidate.NewValidationError("MonitoredResourceSelector", "types", obj.Types, "field must have at most 1 items", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *MetricSelector) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if len(obj.Types) < 1 {
		return gotenvalidate.NewValidationError("MetricSelector", "types", obj.Types, "field must have at least 1 items", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *TimeSeriesSelector) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.Metric == nil {
		return gotenvalidate.NewValidationError("TimeSeriesSelector", "metric", obj.Metric, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Metric).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("TimeSeriesSelector", "metric", obj.Metric, "nested object validation failed", err)
		}
	}
	if obj.Resource == nil {
		return gotenvalidate.NewValidationError("TimeSeriesSelector", "resource", obj.Resource, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Resource).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("TimeSeriesSelector", "resource", obj.Resource, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
