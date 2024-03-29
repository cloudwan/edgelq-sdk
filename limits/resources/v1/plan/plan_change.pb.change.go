// Code generated by protoc-gen-goten-resource
// Resource change: PlanChange
// DO NOT EDIT!!!

package plan

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	meta_service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &iam_organization.Organization{}
	_ = &fieldmaskpb.FieldMask{}
	_ = &meta_service.Service{}
)

func (c *PlanChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*PlanChange_Added_)
	return ok
}

func (c *PlanChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*PlanChange_Modified_)
	return ok
}

func (c *PlanChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*PlanChange_Current_)
	return ok
}

func (c *PlanChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*PlanChange_Removed_)
	return ok
}

func (c *PlanChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *PlanChange_Added_:
		return cType.Added.ViewIndex
	case *PlanChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *PlanChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *PlanChange_Removed_:
		return cType.Removed.ViewIndex
	case *PlanChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *PlanChange) GetPlan() *Plan {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *PlanChange_Added_:
		return cType.Added.Plan
	case *PlanChange_Modified_:
		return cType.Modified.Plan
	case *PlanChange_Current_:
		return cType.Current.Plan
	case *PlanChange_Removed_:
		return nil
	}
	return nil
}

func (c *PlanChange) GetRawResource() gotenresource.Resource {
	return c.GetPlan()
}

func (c *PlanChange) GetPlanName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *PlanChange_Added_:
		return cType.Added.Plan.GetName()
	case *PlanChange_Modified_:
		return cType.Modified.Name
	case *PlanChange_Current_:
		return cType.Current.Plan.GetName()
	case *PlanChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *PlanChange) GetRawName() gotenresource.Name {
	return c.GetPlanName()
}

func (c *PlanChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &PlanChange_Added_{
		Added: &PlanChange_Added{
			Plan:      snapshot.(*Plan),
			ViewIndex: int32(idx),
		},
	}
}

func (c *PlanChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &PlanChange_Modified_{
		Modified: &PlanChange_Modified{
			Name:              name.(*Name),
			Plan:              snapshot.(*Plan),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *PlanChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &PlanChange_Current_{
		Current: &PlanChange_Current{
			Plan: snapshot.(*Plan),
		},
	}
}

func (c *PlanChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &PlanChange_Removed_{
		Removed: &PlanChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
