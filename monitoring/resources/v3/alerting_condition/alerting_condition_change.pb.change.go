// Code generated by protoc-gen-goten-resource
// Resource change: AlertingConditionChange
// DO NOT EDIT!!!

package alerting_condition

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	alerting_policy "github.com/cloudwan/edgelq-sdk/monitoring/resources/v3/alerting_policy"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &alerting_policy.AlertingPolicy{}
	_ = &field_mask.FieldMask{}
)

func (c *AlertingConditionChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AlertingConditionChange_Added_)
	return ok
}

func (c *AlertingConditionChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AlertingConditionChange_Modified_)
	return ok
}

func (c *AlertingConditionChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AlertingConditionChange_Current_)
	return ok
}

func (c *AlertingConditionChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AlertingConditionChange_Removed_)
	return ok
}

func (c *AlertingConditionChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *AlertingConditionChange_Added_:
		return cType.Added.ViewIndex
	case *AlertingConditionChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *AlertingConditionChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *AlertingConditionChange_Removed_:
		return cType.Removed.ViewIndex
	case *AlertingConditionChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *AlertingConditionChange) GetAlertingCondition() *AlertingCondition {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *AlertingConditionChange_Added_:
		return cType.Added.AlertingCondition
	case *AlertingConditionChange_Modified_:
		return cType.Modified.AlertingCondition
	case *AlertingConditionChange_Current_:
		return cType.Current.AlertingCondition
	case *AlertingConditionChange_Removed_:
		return nil
	}
	return nil
}

func (c *AlertingConditionChange) GetRawResource() gotenresource.Resource {
	return c.GetAlertingCondition()
}

func (c *AlertingConditionChange) GetAlertingConditionName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *AlertingConditionChange_Added_:
		return cType.Added.AlertingCondition.GetName()
	case *AlertingConditionChange_Modified_:
		return cType.Modified.Name
	case *AlertingConditionChange_Current_:
		return cType.Current.AlertingCondition.GetName()
	case *AlertingConditionChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *AlertingConditionChange) GetRawName() gotenresource.Name {
	return c.GetAlertingConditionName()
}

func (c *AlertingConditionChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &AlertingConditionChange_Added_{
		Added: &AlertingConditionChange_Added{
			AlertingCondition: snapshot.(*AlertingCondition),
			ViewIndex:         int32(idx),
		},
	}
}

func (c *AlertingConditionChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &AlertingConditionChange_Modified_{
		Modified: &AlertingConditionChange_Modified{
			Name:              name.(*Name),
			AlertingCondition: snapshot.(*AlertingCondition),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *AlertingConditionChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &AlertingConditionChange_Current_{
		Current: &AlertingConditionChange_Current{
			AlertingCondition: snapshot.(*AlertingCondition),
		},
	}
}

func (c *AlertingConditionChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &AlertingConditionChange_Removed_{
		Removed: &AlertingConditionChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
