// Code generated by protoc-gen-goten-validate
// File: edgelq/alerting/proto/v1/ts_condition.proto
// DO NOT EDIT!!!

package ts_condition

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
	document "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/document"
	policy "github.com/cloudwan/edgelq-sdk/alerting/resources/v1/policy"
	monitoring_common "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/common"
	monitoring_time_serie "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/time_serie"
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
	_ = &document.Document{}
	_ = &policy.Policy{}
	_ = &monitoring_common.LabelDescriptor{}
	_ = &monitoring_time_serie.Point{}
	_ = &durationpb.Duration{}
	_ = &meta.Meta{}
)

func (obj *TsCondition) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("TsCondition", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	{
		rlen := utf8.RuneCountInString(obj.DisplayName)
		if rlen > 256 {
			return gotenvalidate.NewValidationError("TsCondition", "displayName", obj.DisplayName, "field must contain at most 256 characters", nil)
		}
	}
	{
		rlen := utf8.RuneCountInString(obj.Description)
		if rlen > 512 {
			return gotenvalidate.NewValidationError("TsCondition", "description", obj.Description, "field must contain at most 512 characters", nil)
		}
	}
	if obj.Spec == nil {
		return gotenvalidate.NewValidationError("TsCondition", "spec", obj.Spec, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Spec).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("TsCondition", "spec", obj.Spec, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Internal).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("TsCondition", "internal", obj.Internal, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.FilterSelector).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("TsCondition", "filterSelector", obj.FilterSelector, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *TsCondition_Spec) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Queries {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("Spec", "queries", obj.Queries[idx], "nested object validation failed", err)
			}
		}
	}
	if len(obj.QueryGroupBy) < 0 {
		return gotenvalidate.NewValidationError("Spec", "queryGroupBy", obj.QueryGroupBy, "field must have at least 0 items", nil)
	}
	if len(obj.QueryGroupBy) > 8 {
		return gotenvalidate.NewValidationError("Spec", "queryGroupBy", obj.QueryGroupBy, "field must have at most 8 items", nil)
	}
	if len(obj.QueryGroupBy) > 1 {
		values := make(map[string]struct{})
		for _, v := range obj.QueryGroupBy {
			if _, ok := values[v]; ok {
				return gotenvalidate.NewValidationError("Spec", "queryGroupBy", obj.QueryGroupBy, "field must contain unique items", nil)
			}
			values[v] = struct{}{}
		}
	}
	if obj.ThresholdAlerting == nil {
		return gotenvalidate.NewValidationError("Spec", "thresholdAlerting", obj.ThresholdAlerting, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.ThresholdAlerting).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Spec", "thresholdAlerting", obj.ThresholdAlerting, "nested object validation failed", err)
		}
	}
	for idx, elem := range obj.AnomalyAlerting {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("Spec", "anomalyAlerting", obj.AnomalyAlerting[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *TsCondition_Internal) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *TsCondition_Selector) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *TsCondition_Spec_Query) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if len(obj.Name) > 256 {
		return gotenvalidate.NewValidationError("Query", "name", obj.Name, "field must contain at most 256 characters", nil)
	}
	if obj.Name == "" {
		return gotenvalidate.NewValidationError("Query", "name", obj.Name, "field is required", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *TsCondition_Spec_ThresholdAlertingCfg) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if _, ok := TsCondition_Spec_ThresholdAlertingCfg_Operator_name[int32(obj.Operator)]; !ok {
		return gotenvalidate.NewValidationError("ThresholdAlertingCfg", "operator", obj.Operator, "field must be a defined enum value", nil)
	}
	if obj.AlignmentPeriod != nil && obj.AlignmentPeriod.CheckValid() != nil {
		err := obj.AlignmentPeriod.CheckValid()
		return gotenvalidate.NewValidationError("ThresholdAlertingCfg", "alignmentPeriod", obj.AlignmentPeriod, "could not validate duration", err)
	} else {
		d := obj.AlignmentPeriod.AsDuration()

		if obj.AlignmentPeriod != nil {
			if d != time.Duration(60000000000) && d != time.Duration(180000000000) && d != time.Duration(300000000000) && d != time.Duration(900000000000) && d != time.Duration(1800000000000) && d != time.Duration(3600000000000) && d != time.Duration(10800000000000) && d != time.Duration(21600000000000) && d != time.Duration(43200000000000) && d != time.Duration(86400000000000) {
				return gotenvalidate.NewValidationError("ThresholdAlertingCfg", "alignmentPeriod", d, "field must be equal to exactly one of the following values: 1m0s, 3m0s, 5m0s, 15m0s, 30m0s, 1h0m0s, 3h0m0s, 6h0m0s, 12h0m0s, 24h0m0s", nil)
			}
		}
	}
	if obj.RaiseAfter != nil && obj.RaiseAfter.CheckValid() != nil {
		err := obj.RaiseAfter.CheckValid()
		return gotenvalidate.NewValidationError("ThresholdAlertingCfg", "raiseAfter", obj.RaiseAfter, "could not validate duration", err)
	} else {
		d := obj.RaiseAfter.AsDuration()

		if obj.RaiseAfter != nil {
			if !(d > time.Duration(0)) {
				return gotenvalidate.NewValidationError("ThresholdAlertingCfg", "raiseAfter", d, "field must be greater than 0s", nil)
			}
		}
	}
	for idx, elem := range obj.PerQueryThresholds {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("ThresholdAlertingCfg", "perQueryThresholds", obj.PerQueryThresholds[idx], "nested object validation failed", err)
			}
		}
	}
	if obj.AdaptiveThresholdsDetectionPeriod != nil && obj.AdaptiveThresholdsDetectionPeriod.CheckValid() != nil {
		err := obj.AdaptiveThresholdsDetectionPeriod.CheckValid()
		return gotenvalidate.NewValidationError("ThresholdAlertingCfg", "adaptiveThresholdsDetectionPeriod", obj.AdaptiveThresholdsDetectionPeriod, "could not validate duration", err)
	} else {
		d := obj.AdaptiveThresholdsDetectionPeriod.AsDuration()

		if obj.AdaptiveThresholdsDetectionPeriod != nil {
			if !(d >= time.Duration(0)) {
				return gotenvalidate.NewValidationError("ThresholdAlertingCfg", "adaptiveThresholdsDetectionPeriod", d, "field must be greater or equal to 0s", nil)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *TsCondition_Spec_AnomalyAlertingCfg) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.AnalysisWindow != nil && obj.AnalysisWindow.CheckValid() != nil {
		err := obj.AnalysisWindow.CheckValid()
		return gotenvalidate.NewValidationError("AnomalyAlertingCfg", "analysisWindow", obj.AnalysisWindow, "could not validate duration", err)
	} else {
		d := obj.AnalysisWindow.AsDuration()

		if obj.AnalysisWindow != nil {
			if !(d >= time.Duration(1800000000000)) {
				return gotenvalidate.NewValidationError("AnomalyAlertingCfg", "analysisWindow", d, "field must be greater or equal to 30m0s", nil)
			}
		}
	}
	if obj.StepInterval != nil && obj.StepInterval.CheckValid() != nil {
		err := obj.StepInterval.CheckValid()
		return gotenvalidate.NewValidationError("AnomalyAlertingCfg", "stepInterval", obj.StepInterval, "could not validate duration", err)
	} else {
		d := obj.StepInterval.AsDuration()

		if obj.StepInterval != nil {
			if !(d >= time.Duration(300000000000)) {
				return gotenvalidate.NewValidationError("AnomalyAlertingCfg", "stepInterval", d, "field must be greater or equal to 5m0s", nil)
			}
		}
	}
	if obj.TrainStepInterval != nil && obj.TrainStepInterval.CheckValid() != nil {
		err := obj.TrainStepInterval.CheckValid()
		return gotenvalidate.NewValidationError("AnomalyAlertingCfg", "trainStepInterval", obj.TrainStepInterval, "could not validate duration", err)
	} else {
		d := obj.TrainStepInterval.AsDuration()

		if obj.TrainStepInterval != nil {
			if !(d >= time.Duration(300000000000)) {
				return gotenvalidate.NewValidationError("AnomalyAlertingCfg", "trainStepInterval", d, "field must be greater or equal to 5m0s", nil)
			}
		}
	}
	if obj.AlignmentPeriod != nil && obj.AlignmentPeriod.CheckValid() != nil {
		err := obj.AlignmentPeriod.CheckValid()
		return gotenvalidate.NewValidationError("AnomalyAlertingCfg", "alignmentPeriod", obj.AlignmentPeriod, "could not validate duration", err)
	} else {
		d := obj.AlignmentPeriod.AsDuration()

		if obj.AlignmentPeriod != nil {
			if d != time.Duration(60000000000) && d != time.Duration(180000000000) && d != time.Duration(300000000000) && d != time.Duration(900000000000) && d != time.Duration(1800000000000) && d != time.Duration(3600000000000) && d != time.Duration(10800000000000) && d != time.Duration(21600000000000) && d != time.Duration(43200000000000) && d != time.Duration(86400000000000) {
				return gotenvalidate.NewValidationError("AnomalyAlertingCfg", "alignmentPeriod", d, "field must be equal to exactly one of the following values: 1m0s, 3m0s, 5m0s, 15m0s, 30m0s, 1h0m0s, 3h0m0s, 6h0m0s, 12h0m0s, 24h0m0s", nil)
			}
		}
	}
	if obj.RaiseAfter != nil && obj.RaiseAfter.CheckValid() != nil {
		err := obj.RaiseAfter.CheckValid()
		return gotenvalidate.NewValidationError("AnomalyAlertingCfg", "raiseAfter", obj.RaiseAfter, "could not validate duration", err)
	} else {
		d := obj.RaiseAfter.AsDuration()

		if obj.RaiseAfter != nil {
			if !(d > time.Duration(0)) {
				return gotenvalidate.NewValidationError("AnomalyAlertingCfg", "raiseAfter", d, "field must be greater than 0s", nil)
			}
		}
	}
	switch opt := obj.Model.(type) {
	case *TsCondition_Spec_AnomalyAlertingCfg_LstmAutoencoder:
		if subobj, ok := interface{}(opt.LstmAutoencoder).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("AnomalyAlertingCfg", "lstmAutoencoder", opt.LstmAutoencoder, "nested object validation failed", err)
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
func (obj *TsCondition_Spec_ThresholdAlertingCfg_AlertingThresholds) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.MaxUpper).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("AlertingThresholds", "maxUpper", obj.MaxUpper, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.MaxLower).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("AlertingThresholds", "maxLower", obj.MaxLower, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *TsCondition_Spec_AnomalyAlertingCfg_LstmAutoEncoder) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if !(obj.HiddenSize >= 8 && obj.HiddenSize <= 64) {
		return gotenvalidate.NewValidationError("LstmAutoEncoder", "hiddenSize", obj.HiddenSize, "field must be in range [8, 64]", nil)
	}
	if !(obj.MaxTrainingEpochs >= 32 && obj.MaxTrainingEpochs <= 1024) {
		return gotenvalidate.NewValidationError("LstmAutoEncoder", "maxTrainingEpochs", obj.MaxTrainingEpochs, "field must be in range [32, 1024]", nil)
	}
	if !(obj.MinTrainingEpochs >= 32 && obj.MinTrainingEpochs <= 1024) {
		return gotenvalidate.NewValidationError("LstmAutoEncoder", "minTrainingEpochs", obj.MinTrainingEpochs, "field must be in range [32, 1024]", nil)
	}
	if obj.TrainingPeriod != nil && obj.TrainingPeriod.CheckValid() != nil {
		err := obj.TrainingPeriod.CheckValid()
		return gotenvalidate.NewValidationError("LstmAutoEncoder", "trainingPeriod", obj.TrainingPeriod, "could not validate duration", err)
	} else {
		d := obj.TrainingPeriod.AsDuration()

		if obj.TrainingPeriod == nil {
			return gotenvalidate.NewValidationError("LstmAutoEncoder", "trainingPeriod", d, "field is required", nil)
		}

		if obj.TrainingPeriod != nil {
			if !(d >= time.Duration(86400000000000)) {
				return gotenvalidate.NewValidationError("LstmAutoEncoder", "trainingPeriod", d, "field must be greater or equal to 24h0m0s", nil)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *TsCondition_Selector_Strings) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *AlertingThreshold) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
