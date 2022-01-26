// Code generated by protoc-gen-goten-resource
// Resource: ResourceChangeLog
// DO NOT EDIT!!!

package resource_change_log

import (
	"fmt"
	"math"
	"reflect"
	"strings"

	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/common/types/traits"
	firestorepb "google.golang.org/genproto/googleapis/firestore/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	filterParser "github.com/cloudwan/goten-sdk/runtime/resource/filter"
	utils "github.com/cloudwan/goten-sdk/runtime/utils"
)

// proto imports
import (
	audit_common "github.com/cloudwan/edgelq-sdk/audit/common/v1alpha2"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

var (
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = math.IsNaN

	_ = types.Bool(true)
	_ = ref.Type(nil)
	_ = traits.Receiver(nil)
	_ = firestorepb.Value{}

	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &audit_common.Authentication{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &timestamp.Timestamp{}
)

type FilterCondition interface {
	gotenresource.FilterCondition
	_IsFilterCondition()
	_IsResourceChangeLogFilterBuilderOrCondition()
	And(...FilterCondition) FilterCondition
	Evaluate(res *ResourceChangeLog) bool

	// Whether this condition is at least as specific as other.
	// When true, any ResourceChangeLog that passes this condition will also pass other condition.
	Satisfies(other FilterCondition) bool

	// Checks whether condition specifies given field path
	// Useful for blacklisting protected paths in iam policy conditions
	SpecifiesFieldPath(fp ResourceChangeLog_FieldPath) bool
}

func AndFilterConditions(conds ...FilterCondition) FilterCondition {
	result := &FilterConditionComposite{
		Operator: filterParser.AND,
	}
	for _, condi := range conds {
		switch cond := condi.(type) {
		case *FilterConditionComposite:
			if cond.Operator == filterParser.AND {
				result.Conditions = append(result.Conditions, cond.Conditions...)
				continue
			}
		default:
		}
		result.Conditions = append(result.Conditions, condi)
	}
	return result
}

type FilterConditionComposite struct {
	Operator   filterParser.CompositeOperator
	Conditions []FilterCondition
}

func (cond *FilterConditionComposite) String() string {
	substrs := make([]string, 0, len(cond.Conditions))
	for _, subcond := range cond.Conditions {
		substrs = append(substrs, subcond.String())
	}
	sep := fmt.Sprintf(" %s ", cond.Operator)
	return "(" + strings.Join(substrs, sep) + ")"
}

func (cond *FilterConditionComposite) _IsFilterCondition() {}

func (cond *FilterConditionComposite) _IsResourceChangeLogFilterBuilderOrCondition() {}

func (cond *FilterConditionComposite) And(conds ...FilterCondition) FilterCondition {
	return AndFilterConditions(append([]FilterCondition{cond}, conds...)...)
}

func (cond *FilterConditionComposite) Evaluate(res *ResourceChangeLog) bool {
	switch cond.Operator {
	case filterParser.OR:
		for _, subCond := range cond.Conditions {
			if subCond.Evaluate(res) {
				return true
			}
		}
		return false
	case filterParser.AND:
		for _, subCond := range cond.Conditions {
			if !subCond.Evaluate(res) {
				return false
			}
		}
		return true
	default:
		panic(fmt.Sprintf("Unsupported composite condition operator: %s", cond.Operator))
	}
}

func (cond *FilterConditionComposite) EvaluateRaw(res gotenresource.Resource) bool {
	if typedRes, ok := res.(*ResourceChangeLog); !ok {
		return false
	} else {
		return cond.Evaluate(typedRes)
	}
}

func (cond *FilterConditionComposite) flattenConditions() (results []FilterCondition) {
	for _, subcnd := range cond.Conditions {
		switch tsubcnd := subcnd.(type) {
		case *FilterConditionComposite:
			if tsubcnd.Operator == cond.Operator {
				results = append(results, tsubcnd.flattenConditions()...)
			} else {
				results = append(results, subcnd) // take it as it is
			}
		default:
			results = append(results, subcnd)
		}
	}
	return
}

func (cond *FilterConditionComposite) Satisfies(other FilterCondition) bool {
	flattened := cond.flattenConditions()
	switch cond.Operator {
	case filterParser.AND:
		switch tother := other.(type) {
		case *FilterConditionComposite:
			switch tother.Operator {
			case filterParser.AND:
				otherFlattened := tother.flattenConditions()
			OtherSubcnds:
				for _, otherSubcnd := range otherFlattened {
					for _, subcnd := range flattened {
						if subcnd.Satisfies(otherSubcnd) {
							continue OtherSubcnds
						}
					}
					return false
				}
				return true
			case filterParser.OR:
				otherFlattened := tother.flattenConditions()
				for _, otherSubcnd := range otherFlattened {
					if cond.Satisfies(otherSubcnd) {
						return true
					}
				}
				return false
			default:
				return false
			}
		default:
			for _, subcnd := range flattened {
				if subcnd.Satisfies(other) {
					return true
				}
			}
			return false
		}
	default:
		panic(fmt.Errorf("unsupported condition type %s", cond.Operator))
	}
	return false
}

func (cond *FilterConditionComposite) SatisfiesRaw(other gotenresource.FilterCondition) bool {
	if typedCond, ok := other.(FilterCondition); !ok {
		return false
	} else {
		return cond.Satisfies(typedCond)
	}
}

func (cond *FilterConditionComposite) SpecifiesFieldPath(fp ResourceChangeLog_FieldPath) bool {
	for _, subcnd := range cond.Conditions {
		if subcnd.SpecifiesFieldPath(fp) {
			return true
		}
	}
	return false
}

func (cond *FilterConditionComposite) SpecifiesRawFieldPath(fp gotenobject.FieldPath) bool {
	if typedFp, ok := fp.(ResourceChangeLog_FieldPath); !ok {
		return false
	} else {
		return cond.SpecifiesFieldPath(typedFp)
	}
}

func (cond *FilterConditionComposite) GetOperator() filterParser.CompositeOperator {
	return cond.Operator
}

func (cond *FilterConditionComposite) GetSubConditions() []gotenresource.FilterCondition {
	subConds := make([]gotenresource.FilterCondition, len(cond.Conditions))
	for idx, subCond := range cond.Conditions {
		subConds[idx] = subCond
	}
	return subConds
}

func (cond *FilterConditionComposite) ConditionComposite() {}

type FilterConditionNot struct {
	FilterCondition
}

func (cond *FilterConditionNot) String() string {
	return "NOT " + cond.FilterCondition.String()
}

func (cond *FilterConditionNot) _IsFilterCondition() {}

func (cond *FilterConditionNot) _IsResourceChangeLogFilterBuilderOrCondition() {}

func (cond *FilterConditionNot) And(conds ...FilterCondition) FilterCondition {
	return AndFilterConditions(append([]FilterCondition{cond}, conds...)...)
}

func (cond *FilterConditionNot) Evaluate(res *ResourceChangeLog) bool {
	return !cond.FilterCondition.Evaluate(res)
}

func (cond *FilterConditionNot) EvaluateRaw(res gotenresource.Resource) bool {
	if typedRes, ok := res.(*ResourceChangeLog); !ok {
		return false
	} else {
		return cond.Evaluate(typedRes)
	}
}

func (cond *FilterConditionNot) Satisfies(other FilterCondition) bool {
	switch tother := other.(type) {
	case *FilterConditionNot:
		return cond.FilterCondition.Satisfies(tother.FilterCondition)
	default:
		return !cond.FilterCondition.Satisfies(other)
	}
}

func (cond *FilterConditionNot) SatisfiesRaw(other gotenresource.FilterCondition) bool {
	if typedCond, ok := other.(FilterCondition); !ok {
		return false
	} else {
		return cond.Satisfies(typedCond)
	}
}

func (cond *FilterConditionNot) SpecifiesFieldPath(fp ResourceChangeLog_FieldPath) bool {
	return cond.FilterCondition.SpecifiesFieldPath(fp)
}

func (cond *FilterConditionNot) SpecifiesRawFieldPath(fp gotenobject.FieldPath) bool {
	if typedFp, ok := fp.(ResourceChangeLog_FieldPath); !ok {
		return false
	} else {
		return cond.SpecifiesFieldPath(typedFp)
	}
}

func (cond *FilterConditionNot) GetSubCondition() gotenresource.FilterCondition {
	return cond.FilterCondition
}

func (cond *FilterConditionNot) ConditionNot() {}

type FilterConditionIsNull struct {
	Not       bool
	FieldPath ResourceChangeLog_FieldPath
}

func (cond *FilterConditionIsNull) String() string {
	if cond.Not {
		return cond.FieldPath.String() + " IS NOT NULL"
	} else {
		return cond.FieldPath.String() + " IS NULL"
	}
}

func (cond *FilterConditionIsNull) _IsFilterCondition() {}

func (cond *FilterConditionIsNull) _IsResourceChangeLogFilterBuilderOrCondition() {}

func (cond *FilterConditionIsNull) And(conds ...FilterCondition) FilterCondition {
	return AndFilterConditions(append([]FilterCondition{cond}, conds...)...)
}

func (cond *FilterConditionIsNull) Evaluate(res *ResourceChangeLog) bool {
	if v, ok := cond.FieldPath.GetSingleRaw(res); !ok {
		return !cond.Not
	} else {
		return cond.Not != utils.IsNil(v)
	}
}

func (cond *FilterConditionIsNull) EvaluateRaw(res gotenresource.Resource) bool {
	if typedRes, ok := res.(*ResourceChangeLog); !ok {
		return false
	} else {
		return cond.Evaluate(typedRes)
	}
}

func (cond *FilterConditionIsNull) asCompare() FilterCondition {
	res := &FilterConditionCompare{
		Operator:                         filterParser.Eq,
		ResourceChangeLog_FieldPathValue: cond.FieldPath.WithIValue(nil),
	}
	if cond.Not {
		res.Operator = filterParser.Neq
	}
	return res
}

func (cond *FilterConditionIsNull) Satisfies(other FilterCondition) bool {
	switch tother := other.(type) {
	case *FilterConditionIsNull:
		return cond.FieldPath.String() == tother.FieldPath.String() && cond.Not == tother.Not
	case *FilterConditionCompare:
		return cond.asCompare().Satisfies(tother)
	default:
		return false
	}
}

func (cond *FilterConditionIsNull) SatisfiesRaw(other gotenresource.FilterCondition) bool {
	if typedCond, ok := other.(FilterCondition); !ok {
		return false
	} else {
		return cond.Satisfies(typedCond)
	}
}

func (cond *FilterConditionIsNull) SpecifiesFieldPath(fp ResourceChangeLog_FieldPath) bool {
	return cond.FieldPath.String() == fp.String()
}

func (cond *FilterConditionIsNull) SpecifiesRawFieldPath(fp gotenobject.FieldPath) bool {
	if typedFp, ok := fp.(ResourceChangeLog_FieldPath); !ok {
		return false
	} else {
		return cond.SpecifiesFieldPath(typedFp)
	}
}

func (cond *FilterConditionIsNull) NotNull() bool {
	return cond.Not
}

func (cond *FilterConditionIsNull) GetRawFieldPath() gotenobject.FieldPath {
	return cond.FieldPath
}

func (cond *FilterConditionIsNull) ConditionIsNull() {}

type FilterConditionIsNaN struct {
	Not       bool
	FieldPath ResourceChangeLog_FieldPath
}

func (cond *FilterConditionIsNaN) String() string {
	if cond.Not {
		return cond.FieldPath.String() + " IS NOT NaN"
	} else {
		return cond.FieldPath.String() + " IS NaN"
	}
}

func (cond *FilterConditionIsNaN) _IsFilterCondition() {}

func (cond *FilterConditionIsNaN) _IsResourceChangeLogFilterBuilderOrCondition() {}

func (cond *FilterConditionIsNaN) And(conds ...FilterCondition) FilterCondition {
	return AndFilterConditions(append([]FilterCondition{cond}, conds...)...)
}

func (cond *FilterConditionIsNaN) Evaluate(res *ResourceChangeLog) bool {
	v, ok := cond.FieldPath.GetSingleRaw(res)
	if !ok {
		return false
	}
	fv, ok := v.(float64)
	if !ok {
		return false
	}
	return math.IsNaN(fv)
}

func (cond *FilterConditionIsNaN) EvaluateRaw(res gotenresource.Resource) bool {
	if typedRes, ok := res.(*ResourceChangeLog); !ok {
		return false
	} else {
		return cond.Evaluate(typedRes)
	}
}

func (cond *FilterConditionIsNaN) Satisfies(other FilterCondition) bool {
	switch tother := other.(type) {
	case *FilterConditionIsNaN:
		return cond.FieldPath.String() == tother.FieldPath.String() && cond.Not == tother.Not
	default:
		return false
	}
}

func (cond *FilterConditionIsNaN) SatisfiesRaw(other gotenresource.FilterCondition) bool {
	if typedCond, ok := other.(FilterCondition); !ok {
		return false
	} else {
		return cond.Satisfies(typedCond)
	}
}

func (cond *FilterConditionIsNaN) SpecifiesFieldPath(fp ResourceChangeLog_FieldPath) bool {
	return cond.FieldPath.String() == fp.String()
}

func (cond *FilterConditionIsNaN) SpecifiesRawFieldPath(fp gotenobject.FieldPath) bool {
	if typedFp, ok := fp.(ResourceChangeLog_FieldPath); !ok {
		return false
	} else {
		return cond.SpecifiesFieldPath(typedFp)
	}
}

func (cond *FilterConditionIsNaN) GetRawFieldPath() gotenobject.FieldPath {
	return cond.FieldPath
}

func (cond *FilterConditionIsNaN) ConditionIsNaN() {}

type FilterConditionCompare struct {
	Operator filterParser.CompareOperator
	ResourceChangeLog_FieldPathValue
}

func (cond *FilterConditionCompare) String() string {
	jsonValue, err := utils.JsonMarshal(cond.ResourceChangeLog_FieldPathValue.GetRawValue())
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s %s %s", cond.ResourceChangeLog_FieldPathValue, cond.Operator, jsonValue)
}

func (cond *FilterConditionCompare) _IsFilterCondition() {}

func (cond *FilterConditionCompare) _IsResourceChangeLogFilterBuilderOrCondition() {}

func (cond *FilterConditionCompare) And(conds ...FilterCondition) FilterCondition {
	return AndFilterConditions(append([]FilterCondition{cond}, conds...)...)
}

func (cond *FilterConditionCompare) Evaluate(res *ResourceChangeLog) bool {
	// Special evaluation for name or reference - may include wildcards
	if nameOrRefFPV, ok := cond.ResourceChangeLog_FieldPathValue.GetRawValue().(gotenresource.Name); ok {
		if otherObj, ok := cond.ResourceChangeLog_FieldPathValue.GetSingleRaw(res); ok {
			match := nameOrRefFPV.Matches(otherObj)
			switch cond.Operator {
			case filterParser.Eq:
				return match
			case filterParser.Neq:
				return !match
			default:
				return false
			}
		}
	}
	cmpResult, comparable := cond.ResourceChangeLog_FieldPathValue.CompareWith(res)
	if !comparable {
		return false
	}
	return cond.Operator.MatchCompareResult(cmpResult)
}

func (cond *FilterConditionCompare) EvaluateRaw(res gotenresource.Resource) bool {
	if typedRes, ok := res.(*ResourceChangeLog); !ok {
		return false
	} else {
		return cond.Evaluate(typedRes)
	}
}

func (cond *FilterConditionCompare) Satisfies(other FilterCondition) bool {
	switch tother := other.(type) {
	case *FilterConditionCompare:
		if cond.ResourceChangeLog_FieldPathValue.String() != tother.ResourceChangeLog_FieldPathValue.String() {
			return false
		}
		othertmp := new(ResourceChangeLog)
		tother.SetTo(&othertmp)
		if cmp, comparable := cond.CompareWith(othertmp); !comparable {
			return false
		} else {
			return filterParser.CompareSatisfies(tother.Operator, cond.Operator, cmp)
		}
	case *FilterConditionIn:
		if cond.Operator != filterParser.Eq {
			return false
		}
		if cond.ResourceChangeLog_FieldPathValue.String() != tother.ResourceChangeLog_FieldPathArrayOfValues.String() {
			return false
		}
		for _, inv := range tother.GetRawValues() {
			othertmp := new(ResourceChangeLog)
			tother.WithIValue(inv).SetTo(&othertmp)
			if cmp, comparable := cond.ResourceChangeLog_FieldPathValue.CompareWith(othertmp); comparable && cmp == 0 {
				return true
			}
		}
		return false
	case *FilterConditionComposite:
		if tother.Operator == filterParser.AND {
			for _, othersubcnd := range tother.flattenConditions() {
				if !cond.Satisfies(othersubcnd) {
					return false
				}
			}
			return true
		} else { // OR
			for _, othersubcnd := range tother.flattenConditions() {
				if cond.Satisfies(othersubcnd) {
					return true
				}
			}
			return false
		}
	default:
		return false
	}
}

func (cond *FilterConditionCompare) SatisfiesRaw(other gotenresource.FilterCondition) bool {
	if typedCond, ok := other.(FilterCondition); !ok {
		return false
	} else {
		return cond.Satisfies(typedCond)
	}
}

func (cond *FilterConditionCompare) SpecifiesFieldPath(fp ResourceChangeLog_FieldPath) bool {
	return cond.ResourceChangeLog_FieldPathValue.String() == fp.String()
}

func (cond *FilterConditionCompare) SpecifiesRawFieldPath(fp gotenobject.FieldPath) bool {
	if typedFp, ok := fp.(ResourceChangeLog_FieldPath); !ok {
		return false
	} else {
		return cond.SpecifiesFieldPath(typedFp)
	}
}

func (cond *FilterConditionCompare) GetOperator() filterParser.CompareOperator {
	return cond.Operator
}

func (cond *FilterConditionCompare) GetRawFieldPath() gotenobject.FieldPath {
	return cond.ResourceChangeLog_FieldPathValue
}

func (cond *FilterConditionCompare) GetRawFieldPathValue() gotenobject.FieldPathValue {
	return cond.ResourceChangeLog_FieldPathValue
}

func (cond *FilterConditionCompare) ConditionCompare() {}

type FilterConditionContains struct {
	Type      gotenresource.ConditionContainsType
	FieldPath ResourceChangeLog_FieldPath

	Value  ResourceChangeLog_FieldPathArrayItemValue
	Values []ResourceChangeLog_FieldPathArrayItemValue
}

func (cond *FilterConditionContains) String() string {
	switch cond.ConditionContainsType() {
	case gotenresource.ConditionContainsTypeValue:
		jsonValue, err := utils.JsonMarshal(cond.Value.GetRawItemValue())
		if err != nil {
			panic(err)
		}
		return fmt.Sprintf("%s CONTAINS %s", cond.FieldPath, string(jsonValue))
	case gotenresource.ConditionContainsTypeAny, gotenresource.ConditionContainsTypeAll:
		jsonValues := make([]string, len(cond.Values))
		for i, v := range cond.Values {
			if jsonValue, err := utils.JsonMarshal(v.GetRawItemValue()); err != nil {
				panic(err)
			} else {
				jsonValues[i] = string(jsonValue)
			}
		}
		return fmt.Sprintf("%s CONTAINS %s %s", cond.FieldPath, cond.ConditionContainsType(), fmt.Sprintf("(%s)", strings.Join(jsonValues, ", ")))
	default:
		panic(gotenresource.NewUnknownConditionContainsType(cond.ConditionContainsType()))
	}
}

func (cond *FilterConditionContains) ConditionContainsType() gotenresource.ConditionContainsType {
	return cond.Type
}

func (cond *FilterConditionContains) _IsFilterCondition() {}

func (cond *FilterConditionContains) _IsResourceChangeLogFilterBuilderOrCondition() {}

func (cond *FilterConditionContains) And(conds ...FilterCondition) FilterCondition {
	return AndFilterConditions(append([]FilterCondition{cond}, conds...)...)
}

func (cond *FilterConditionContains) Evaluate(res *ResourceChangeLog) bool {
	switch cond.ConditionContainsType() {
	case gotenresource.ConditionContainsTypeValue:
		return cond.Value.ContainsValue(res)
	case gotenresource.ConditionContainsTypeAny:
		for _, v := range cond.Values {
			if v.ContainsValue(res) {
				return true
			}
		}
		return false
	case gotenresource.ConditionContainsTypeAll:
		for _, v := range cond.Values {
			if !v.ContainsValue(res) {
				return false
			}
		}
		return true
	default:
		panic(gotenresource.NewUnknownConditionContainsType(cond.ConditionContainsType()))
	}
}

func (cond *FilterConditionContains) EvaluateRaw(res gotenresource.Resource) bool {
	if typedRes, ok := res.(*ResourceChangeLog); !ok {
		return false
	} else {
		return cond.Evaluate(typedRes)
	}
}

func (cond *FilterConditionContains) Satisfies(other FilterCondition) bool {
	switch tother := other.(type) {
	case *FilterConditionContains:
		if cond.ConditionContainsType().IsValue() && tother.ConditionContainsType().IsValue() {
			othertmp := new(ResourceChangeLog)
			tother.Value.WithIValue(tother.GetRawFieldPathItemValue().GetRawItemValue()).SetTo(&othertmp)
			return cond.Value.ContainsValue(othertmp)
		}
		return false
	default:
		return false
	}
}

func (cond *FilterConditionContains) SatisfiesRaw(other gotenresource.FilterCondition) bool {
	if typedCond, ok := other.(FilterCondition); !ok {
		return false
	} else {
		return cond.Satisfies(typedCond)
	}
}

func (cond *FilterConditionContains) SpecifiesFieldPath(fp ResourceChangeLog_FieldPath) bool {
	return cond.FieldPath.String() == fp.String()
}

func (cond *FilterConditionContains) SpecifiesRawFieldPath(fp gotenobject.FieldPath) bool {
	if typedFp, ok := fp.(ResourceChangeLog_FieldPath); !ok {
		return false
	} else {
		return cond.SpecifiesFieldPath(typedFp)
	}
}

func (cond *FilterConditionContains) GetFieldPath() ResourceChangeLog_FieldPath {
	return cond.FieldPath
}

func (cond *FilterConditionContains) GetRawFieldPath() gotenobject.FieldPath {
	return cond.FieldPath
}

func (cond *FilterConditionContains) GetRawFieldPathItemValue() gotenobject.FieldPathArrayItemValue {
	switch cond.ConditionContainsType() {
	case gotenresource.ConditionContainsTypeValue:
		return cond.Value
	default:
		panic(fmt.Errorf("unable to get value for condition contains type %s", cond.ConditionContainsType()))
	}
}

func (cond *FilterConditionContains) GetRawFieldPathItemValues() (res []gotenobject.FieldPathArrayItemValue) {
	switch cond.ConditionContainsType() {
	case gotenresource.ConditionContainsTypeAny, gotenresource.ConditionContainsTypeAll:
		for _, fpaiv := range cond.Values {
			res = append(res, fpaiv)
		}
	default:
		panic(fmt.Errorf("unable to get values for condition contains type %s", cond.ConditionContainsType()))
	}
	return
}

func (cond *FilterConditionContains) ConditionContains() {}

type FilterConditionIn struct {
	ResourceChangeLog_FieldPathArrayOfValues
}

func (cond *FilterConditionIn) String() string {
	jsonValues, err := utils.JsonMarshal(cond.ResourceChangeLog_FieldPathArrayOfValues.GetRawValues())
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s IN %s", cond.ResourceChangeLog_FieldPathArrayOfValues, jsonValues)
}

func (cond *FilterConditionIn) _IsFilterCondition() {}

func (cond *FilterConditionIn) _IsResourceChangeLogFilterBuilderOrCondition() {}

func (cond *FilterConditionIn) And(conds ...FilterCondition) FilterCondition {
	return AndFilterConditions(append([]FilterCondition{cond}, conds...)...)
}

func (cond *FilterConditionIn) Evaluate(res *ResourceChangeLog) bool {
	for _, inValue := range cond.ResourceChangeLog_FieldPathArrayOfValues.GetRawValues() {
		if cmp, ok := cond.ResourceChangeLog_FieldPathArrayOfValues.WithIValue(inValue).CompareWith(res); ok && cmp == 0 {
			return true
		}
	}
	return false
}

func (cond *FilterConditionIn) EvaluateRaw(res gotenresource.Resource) bool {
	if typedRes, ok := res.(*ResourceChangeLog); !ok {
		return false
	} else {
		return cond.Evaluate(typedRes)
	}
}

func (cond *FilterConditionIn) Satisfies(other FilterCondition) bool {
	switch tother := other.(type) {
	case *FilterConditionIn:
	outer:
		for _, cval := range cond.ResourceChangeLog_FieldPathArrayOfValues.GetRawValues() {
			for _, otherval := range tother.ResourceChangeLog_FieldPathArrayOfValues.GetRawValues() {
				othertmp := new(ResourceChangeLog)
				tother.ResourceChangeLog_FieldPathArrayOfValues.WithIValue(otherval).SetTo(&othertmp)
				if cmp, comparable := cond.ResourceChangeLog_FieldPathArrayOfValues.WithIValue(cval).CompareWith(othertmp); comparable && cmp == 0 {
					continue outer
				}
			}
			return false
		}
		return true
	default:
		for _, cval := range cond.ResourceChangeLog_FieldPathArrayOfValues.GetRawValues() {
			subcnd := &FilterConditionCompare{
				Operator:                         filterParser.Eq,
				ResourceChangeLog_FieldPathValue: cond.ResourceChangeLog_FieldPathArrayOfValues.WithIValue(cval),
			}
			if !subcnd.Satisfies(tother) {
				return false
			}
		}
		return true
	}
}

func (cond *FilterConditionIn) SatisfiesRaw(other gotenresource.FilterCondition) bool {
	if typedCond, ok := other.(FilterCondition); !ok {
		return false
	} else {
		return cond.Satisfies(typedCond)
	}
}

func (cond *FilterConditionIn) SpecifiesFieldPath(fp ResourceChangeLog_FieldPath) bool {
	return cond.ResourceChangeLog_FieldPathArrayOfValues.String() == fp.String()
}

func (cond *FilterConditionIn) SpecifiesRawFieldPath(fp gotenobject.FieldPath) bool {
	if typedFp, ok := fp.(ResourceChangeLog_FieldPath); !ok {
		return false
	} else {
		return cond.SpecifiesFieldPath(typedFp)
	}
}

func (cond *FilterConditionIn) GetRawFieldPath() gotenobject.FieldPath {
	return cond.ResourceChangeLog_FieldPathArrayOfValues
}

func (cond *FilterConditionIn) GetRawFieldPathArrayOfValues() gotenobject.FieldPathArrayOfValues {
	return cond.ResourceChangeLog_FieldPathArrayOfValues
}

func (cond *FilterConditionIn) ConditionIn() {}

type FilterConditionNotIn struct {
	ResourceChangeLog_FieldPathArrayOfValues
}

func (cond *FilterConditionNotIn) String() string {
	jsonValues, err := utils.JsonMarshal(cond.ResourceChangeLog_FieldPathArrayOfValues.GetRawValues())
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s NOT IN %s", cond.ResourceChangeLog_FieldPathArrayOfValues, jsonValues)
}

func (cond *FilterConditionNotIn) _IsFilterCondition() {}

func (cond *FilterConditionNotIn) _IsResourceChangeLogFilterBuilderOrCondition() {}

func (cond *FilterConditionNotIn) And(conds ...FilterCondition) FilterCondition {
	return AndFilterConditions(append([]FilterCondition{cond}, conds...)...)
}

func (cond *FilterConditionNotIn) Evaluate(res *ResourceChangeLog) bool {
	for _, inValue := range cond.ResourceChangeLog_FieldPathArrayOfValues.GetRawValues() {
		if cmp, ok := cond.ResourceChangeLog_FieldPathArrayOfValues.WithIValue(inValue).CompareWith(res); ok && cmp == 0 {
			return false
		}
	}
	return true
}

func (cond *FilterConditionNotIn) EvaluateRaw(res gotenresource.Resource) bool {
	if typedRes, ok := res.(*ResourceChangeLog); !ok {
		return false
	} else {
		return cond.Evaluate(typedRes)
	}
}

func (cond *FilterConditionNotIn) Satisfies(other FilterCondition) bool {
	return false
}

func (cond *FilterConditionNotIn) SatisfiesRaw(other gotenresource.FilterCondition) bool {
	if typedCond, ok := other.(FilterCondition); !ok {
		return false
	} else {
		return cond.Satisfies(typedCond)
	}
}

func (cond *FilterConditionNotIn) SpecifiesFieldPath(fp ResourceChangeLog_FieldPath) bool {
	return cond.ResourceChangeLog_FieldPathArrayOfValues.String() == fp.String()
}

func (cond *FilterConditionNotIn) SpecifiesRawFieldPath(fp gotenobject.FieldPath) bool {
	if typedFp, ok := fp.(ResourceChangeLog_FieldPath); !ok {
		return false
	} else {
		return cond.SpecifiesFieldPath(typedFp)
	}
}

func (cond *FilterConditionNotIn) GetRawFieldPath() gotenobject.FieldPath {
	return cond.ResourceChangeLog_FieldPathArrayOfValues
}

func (cond *FilterConditionNotIn) GetRawFieldPathArrayOfValues() gotenobject.FieldPathArrayOfValues {
	return cond.ResourceChangeLog_FieldPathArrayOfValues
}

func (cond *FilterConditionNotIn) ConditionNotIn() {}

type Filter struct {
	FilterCondition
}

func (filter *Filter) _IsResourceChangeLogFilterBuilderOrCondition() {}

// GetCondition is a getter of FilterCondition, which also handles nil pointer
func (filter *Filter) GetCondition() FilterCondition {
	if filter == nil {
		return AndFilterConditions()
	} else {
		return filter.FilterCondition
	}
}

// Evaluate is a wrapper on FilterCondition, which also handles nil pointer
func (filter *Filter) Evaluate(res *ResourceChangeLog) bool {
	return filter.GetCondition().Evaluate(res)
}

func (filter *Filter) EvaluateRaw(res gotenresource.Resource) bool {
	if typedRes, ok := res.(*ResourceChangeLog); !ok {
		return false
	} else {
		return filter.Evaluate(typedRes)
	}
}

func (filter *Filter) GetRawCondition() gotenresource.FilterCondition {
	if filter == nil {
		return nil
	}
	return filter.GetCondition()
}

// FilterSlice is a helper for filtering arrays
func (filter *Filter) FilterSlice(in []*ResourceChangeLog) (out []*ResourceChangeLog) {
	for _, res := range in {
		if filter.Evaluate(res) {
			out = append(out, res)
		}
	}
	return
}

// implement methods required by protobuf-go library for string-struct conversion

func (filter *Filter) ProtoString() (string, error) {
	if filter == nil || filter.FilterCondition == nil {
		return "", nil
	}
	return filter.FilterCondition.String(), nil
}

func (filter *Filter) ParseProtoString(data string) error {
	expression, err := filterParser.Parse([]byte(data))
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	condition, err := makeFilterConditionFromOr(expression.And)
	if err != nil {
		return err
	}
	filter.FilterCondition = condition
	return nil
}

func (filter *Filter) String() string {
	if filter == nil || filter.FilterCondition == nil {
		return "<nil>"
	}
	return filter.FilterCondition.String()
}

func (filter *Filter) SetFromCliFlag(raw string) error {
	return filter.ParseProtoString(raw)
}

// Google CEL integration Type
var celFilter = types.NewTypeValue("Filter", traits.ReceiverType)

func (filter *Filter) TypeName() string {
	return ".ntt.audit.v1alpha2.ResourceChangeLog.Filter"
}

func (filter *Filter) HasTrait(trait int) bool {
	return trait == traits.ReceiverType
}

func (filter *Filter) Equal(other ref.Val) ref.Val {
	return types.Bool(false)
}

func (filter *Filter) Value() interface{} {
	return filter
}

func (filter *Filter) Match(pattern ref.Val) ref.Val {
	return types.Bool(true)
}

func (filter *Filter) Receive(function string, overload string, args []ref.Val) ref.Val {
	cnd := filter.GetCondition()
	switch function {
	case "satisfies":
		rhsFilter := new(Filter)
		if err := rhsFilter.ParseProtoString(string(args[0].(types.String))); err != nil {
			return types.ValOrErr(celFilter, "function %s condition parse error: %s", function, err)
		}
		return types.Bool(cnd.Satisfies(rhsFilter.GetCondition()))
	case "specifies":
		fp, err := ParseResourceChangeLog_FieldPath(string(args[0].(types.String)))
		if err != nil {
			return types.ValOrErr(celFilter, "function %s field-path parse error: %s", function, err)
		}
		return types.Bool(cnd.SpecifiesFieldPath(fp))
	default:
		return types.ValOrErr(celFilter, "no such function - %s", function)
	}
}

func (filter *Filter) ConvertToNative(typeDesc reflect.Type) (interface{}, error) {
	panic("not required")
}

func (filter *Filter) ConvertToType(typeVal ref.Type) ref.Val {
	switch typeVal.TypeName() {
	case types.StringType.TypeName():
		return types.String(filter.String())
	default:
		panic(fmt.Errorf("unable to convert %s to CEL type %s", "Filter", typeVal.TypeName()))
	}
}

func (filter *Filter) Type() ref.Type {
	return celFilter
}

// firestore encoding/decoding integration
func (filter *Filter) EncodeFirestore() (*firestorepb.Value, error) {
	if filter == nil {
		return &firestorepb.Value{ValueType: &firestorepb.Value_NullValue{}}, nil
	}
	if fvalue, err := filter.ProtoString(); err != nil {
		return nil, err
	} else {
		return &firestorepb.Value{ValueType: &firestorepb.Value_StringValue{StringValue: fvalue}}, nil
	}
}

func (filter *Filter) DecodeFirestore(fpbv *firestorepb.Value) error {
	return filter.ParseProtoString(fpbv.GetStringValue())
}

// helpers

func makeFilterConditionFromOperand(condition *filterParser.ConditionOperand) (FilterCondition, error) {
	path, err := ParseResourceChangeLog_FieldPath(condition.FieldPath)
	if err != nil {
		return nil, err
	}
	rhs := condition.ConditionRHS
	valueJSON, err := rhs.JSONValue()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rhs.Compare != nil {
		cmp := rhs.Compare
		if !path.IsLeaf() {
			return nil, status.Errorf(codes.InvalidArgument, "path '%s' is not comparable leaf value", path)
		}

		// translate null comparison to IS(NOT)NULL
		if cmp.Value.Null {
			switch cmp.Operator {
			case filterParser.Eq:
				return &FilterConditionIsNull{false, path}, nil
			case filterParser.Neq:
				return &FilterConditionIsNull{true, path}, nil
			default:
				return nil, status.Errorf(codes.InvalidArgument, "operator '%s' isn't valid when comparing null value", cmp.Operator)
			}
		}

		pfv, err := ParseResourceChangeLog_FieldPathValue(path.String(), string(valueJSON))
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "error when parsing filter value for field path '%s': %s", path, err)
		}

		return &FilterConditionCompare{
			Operator:                         cmp.Operator,
			ResourceChangeLog_FieldPathValue: pfv,
		}, nil
	} else if rhs.Is != nil {
		if rhs.Is.Null {
			return &FilterConditionIsNull{rhs.Is.Not, path}, nil
		} else if rhs.Is.NaN {
			return &FilterConditionIsNaN{rhs.Is.Not, path}, nil
		} else {
			return nil, status.Error(codes.Internal, "unknown filter IS type - expected NULL or NaN")
		}
	} else if rhs.Contains != nil {
		ct := gotenresource.ConditionContainsTypeFromParser(rhs.Contains)
		fp, err := ParseResourceChangeLog_FieldPath(path.String())
		if err != nil {
			return nil, err
		}
		if rhs.Contains.Value != nil {
			pfav, err := ParseResourceChangeLog_FieldPathArrayItemValue(path.String(), string(valueJSON))
			if err != nil {
				return nil, err
			}
			return &FilterConditionContains{ct, fp, pfav, nil}, nil
		} else if rhs.Contains.Any != nil || rhs.Contains.All != nil {
			parrv := rhs.Contains.GetArray()
			vals := make([]ResourceChangeLog_FieldPathArrayItemValue, len(parrv))
			for i, pv := range parrv {
				jsonv, err := utils.JsonMarshal(pv)
				if err != nil {
					return nil, err
				}
				if pfav, err := ParseResourceChangeLog_FieldPathArrayItemValue(path.String(), string(jsonv)); err != nil {
					return nil, err
				} else {
					vals[i] = pfav
				}
			}
			return &FilterConditionContains{ct, fp, nil, vals}, nil
		} else {
			return nil, status.Error(codes.Internal, "unknown condition contains type")
		}
	} else if rhs.Like != nil {
		return nil, status.Errorf(codes.Unimplemented, "'LIKE' condition is not supported")
	} else if rhs.In != nil {
		if fpaov, err := ParseResourceChangeLog_FieldPathArrayOfValues(path.String(), string(valueJSON)); err != nil {
			return nil, err
		} else {
			return &FilterConditionIn{fpaov}, nil
		}
	} else if rhs.NotIn != nil {
		if fpaov, err := ParseResourceChangeLog_FieldPathArrayOfValues(path.String(), string(valueJSON)); err != nil {
			return nil, err
		} else {
			return &FilterConditionNotIn{fpaov}, nil
		}
	}
	return nil, status.Error(codes.Internal, "unknown filter RHS operand type")

}

func makeFilterConditionFromCondition(condition *filterParser.Condition) (FilterCondition, error) {
	if condition.SubExpression != nil {
		return makeFilterConditionFromOr(condition.SubExpression.And)
	} else if condition.Not != nil {
		not, err := makeFilterConditionFromCondition(condition.Not)
		if err != nil {
			return nil, err
		}
		return &FilterConditionNot{not}, nil
	} else if condition.Operand != nil {
		return makeFilterConditionFromOperand(condition.Operand)
	} else {
		return nil, status.Error(codes.Internal, "unknown condition type")
	}
}

func makeFilterConditionFromAnd(conditions []filterParser.Condition) (FilterCondition, error) {
	if len(conditions) == 1 {
		return makeFilterConditionFromCondition(&conditions[0])
	} else {
		cnds := make([]FilterCondition, 0, len(conditions))
		for _, condition := range conditions {
			cnd, err := makeFilterConditionFromCondition(&condition)
			if err != nil {
				return nil, err
			}
			cnds = append(cnds, cnd)
		}
		return &FilterConditionComposite{Operator: filterParser.AND, Conditions: cnds}, nil
	}
}

func makeFilterConditionFromOr(conditions []filterParser.AndCondition) (FilterCondition, error) {
	if len(conditions) == 1 {
		return makeFilterConditionFromAnd(conditions[0].Or)
	} else {
		cnds := make([]FilterCondition, 0, len(conditions))
		for _, condition := range conditions {
			cnd, err := makeFilterConditionFromAnd(condition.Or)
			if err != nil {
				return nil, err
			}
			cnds = append(cnds, cnd)
		}
		return &FilterConditionComposite{Operator: filterParser.OR, Conditions: cnds}, nil
	}
}
