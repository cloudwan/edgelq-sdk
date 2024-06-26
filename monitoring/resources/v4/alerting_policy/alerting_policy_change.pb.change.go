// Code generated by protoc-gen-goten-resource
// Resource change: AlertingPolicyChange
// DO NOT EDIT!!!

package alerting_policy

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/monitoring/resources/v4/project"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &fieldmaskpb.FieldMask{}
)

func (c *AlertingPolicyChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AlertingPolicyChange_Added_)
	return ok
}

func (c *AlertingPolicyChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AlertingPolicyChange_Modified_)
	return ok
}

func (c *AlertingPolicyChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AlertingPolicyChange_Current_)
	return ok
}

func (c *AlertingPolicyChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AlertingPolicyChange_Removed_)
	return ok
}

func (c *AlertingPolicyChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *AlertingPolicyChange_Added_:
		return cType.Added.ViewIndex
	case *AlertingPolicyChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *AlertingPolicyChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *AlertingPolicyChange_Removed_:
		return cType.Removed.ViewIndex
	case *AlertingPolicyChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *AlertingPolicyChange) GetAlertingPolicy() *AlertingPolicy {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *AlertingPolicyChange_Added_:
		return cType.Added.AlertingPolicy
	case *AlertingPolicyChange_Modified_:
		return cType.Modified.AlertingPolicy
	case *AlertingPolicyChange_Current_:
		return cType.Current.AlertingPolicy
	case *AlertingPolicyChange_Removed_:
		return nil
	}
	return nil
}

func (c *AlertingPolicyChange) GetRawResource() gotenresource.Resource {
	return c.GetAlertingPolicy()
}

func (c *AlertingPolicyChange) GetAlertingPolicyName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *AlertingPolicyChange_Added_:
		return cType.Added.AlertingPolicy.GetName()
	case *AlertingPolicyChange_Modified_:
		return cType.Modified.Name
	case *AlertingPolicyChange_Current_:
		return cType.Current.AlertingPolicy.GetName()
	case *AlertingPolicyChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *AlertingPolicyChange) GetRawName() gotenresource.Name {
	return c.GetAlertingPolicyName()
}

func (c *AlertingPolicyChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &AlertingPolicyChange_Added_{
		Added: &AlertingPolicyChange_Added{
			AlertingPolicy: snapshot.(*AlertingPolicy),
			ViewIndex:      int32(idx),
		},
	}
}

func (c *AlertingPolicyChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &AlertingPolicyChange_Modified_{
		Modified: &AlertingPolicyChange_Modified{
			Name:              name.(*Name),
			AlertingPolicy:    snapshot.(*AlertingPolicy),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *AlertingPolicyChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &AlertingPolicyChange_Current_{
		Current: &AlertingPolicyChange_Current{
			AlertingPolicy: snapshot.(*AlertingPolicy),
		},
	}
}

func (c *AlertingPolicyChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &AlertingPolicyChange_Removed_{
		Removed: &AlertingPolicyChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
