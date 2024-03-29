// Code generated by protoc-gen-goten-validate
// File: edgelq/monitoring/proto/v4/alerting_condition.proto
// DO NOT EDIT!!!

package alerting_condition

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
	alerting_policy "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/alerting_policy"
	common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/common"
	time_serie "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/time_serie"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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
	_ = &alerting_policy.AlertingPolicy{}
	_ = &common.LabelDescriptor{}
	_ = &time_serie.Point{}
	_ = &durationpb.Duration{}
	_ = &meta.Meta{}
)

func (obj *AlertingCondition) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("AlertingCondition", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	if obj.Spec == nil {
		return gotenvalidate.NewValidationError("AlertingCondition", "spec", obj.Spec, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Spec).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("AlertingCondition", "spec", obj.Spec, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.State).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("AlertingCondition", "state", obj.State, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *AlertingCondition_Spec) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.TimeSeries == nil {
		return gotenvalidate.NewValidationError("Spec", "timeSeries", obj.TimeSeries, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.TimeSeries).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Spec", "timeSeries", obj.TimeSeries, "nested object validation failed", err)
		}
	}
	if obj.Trigger == nil {
		return gotenvalidate.NewValidationError("Spec", "trigger", obj.Trigger, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Trigger).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Spec", "trigger", obj.Trigger, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *AlertingCondition_State) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *AlertingCondition_Spec_TimeSeries) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.Query == nil {
		return gotenvalidate.NewValidationError("TimeSeries", "query", obj.Query, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Query).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("TimeSeries", "query", obj.Query, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Threshold).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("TimeSeries", "threshold", obj.Threshold, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.CombineThreshold).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("TimeSeries", "combineThreshold", obj.CombineThreshold, "nested object validation failed", err)
		}
	}
	if obj.Duration == nil {
		return gotenvalidate.NewValidationError("TimeSeries", "duration", obj.Duration, "field is required", nil)
	}
	if obj.Duration != nil && obj.Duration.CheckValid() != nil {
		err := obj.Duration.CheckValid()
		return gotenvalidate.NewValidationError("TimeSeries", "duration", obj.Duration, "could not validate duration", err)
	} else {
		d := obj.Duration.AsDuration()

		if obj.Duration != nil {
			if !(d >= time.Duration(60000000000)) {
				return gotenvalidate.NewValidationError("TimeSeries", "duration", d, "field must be greater or equal to 1m0s", nil)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *AlertingCondition_Spec_Trigger) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *AlertingCondition_Spec_TimeSeries_Query) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.Aggregation == nil {
		return gotenvalidate.NewValidationError("Query", "aggregation", obj.Aggregation, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Aggregation).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Query", "aggregation", obj.Aggregation, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *AlertingCondition_Spec_TimeSeries_Threshold) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *AlertingCondition_Spec_TimeSeries_CombineThreshold) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.Combine != 0 {
		return gotenvalidate.NewValidationError("CombineThreshold", "combine", obj.Combine, "field must be equal to exactly one of the following values: 0", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
