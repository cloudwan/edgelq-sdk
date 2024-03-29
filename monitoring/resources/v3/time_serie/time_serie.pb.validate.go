// Code generated by protoc-gen-goten-validate
// File: edgelq/monitoring/proto/v3/time_serie.proto
// DO NOT EDIT!!!

package time_serie

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
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/common"
	metric_descriptor "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/metric_descriptor"
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
	_ = &common.LabelDescriptor{}
	_ = &metric_descriptor.MetricDescriptor{}
)

func (obj *Point) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Interval).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Point", "interval", obj.Interval, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Value).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Point", "value", obj.Value, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Aggregation).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Point", "aggregation", obj.Aggregation, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *TimeSerie) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Metric).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("TimeSerie", "metric", obj.Metric, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Resource).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("TimeSerie", "resource", obj.Resource, "nested object validation failed", err)
		}
	}
	for idx, elem := range obj.Points {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("TimeSerie", "points", obj.Points[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *BulkTimeSeries) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.TimeSeries {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("BulkTimeSeries", "timeSeries", obj.TimeSeries[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
