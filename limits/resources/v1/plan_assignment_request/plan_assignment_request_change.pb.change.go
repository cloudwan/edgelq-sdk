// Code generated by protoc-gen-goten-resource
// Resource change: PlanAssignmentRequestChange
// DO NOT EDIT!!!

package plan_assignment_request

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1/project"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &fieldmaskpb.FieldMask{}
)

func (c *PlanAssignmentRequestChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*PlanAssignmentRequestChange_Added_)
	return ok
}

func (c *PlanAssignmentRequestChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*PlanAssignmentRequestChange_Modified_)
	return ok
}

func (c *PlanAssignmentRequestChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*PlanAssignmentRequestChange_Current_)
	return ok
}

func (c *PlanAssignmentRequestChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*PlanAssignmentRequestChange_Removed_)
	return ok
}

func (c *PlanAssignmentRequestChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *PlanAssignmentRequestChange_Added_:
		return cType.Added.ViewIndex
	case *PlanAssignmentRequestChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *PlanAssignmentRequestChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *PlanAssignmentRequestChange_Removed_:
		return cType.Removed.ViewIndex
	case *PlanAssignmentRequestChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *PlanAssignmentRequestChange) GetPlanAssignmentRequest() *PlanAssignmentRequest {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *PlanAssignmentRequestChange_Added_:
		return cType.Added.PlanAssignmentRequest
	case *PlanAssignmentRequestChange_Modified_:
		return cType.Modified.PlanAssignmentRequest
	case *PlanAssignmentRequestChange_Current_:
		return cType.Current.PlanAssignmentRequest
	case *PlanAssignmentRequestChange_Removed_:
		return nil
	}
	return nil
}

func (c *PlanAssignmentRequestChange) GetRawResource() gotenresource.Resource {
	return c.GetPlanAssignmentRequest()
}

func (c *PlanAssignmentRequestChange) GetPlanAssignmentRequestName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *PlanAssignmentRequestChange_Added_:
		return cType.Added.PlanAssignmentRequest.GetName()
	case *PlanAssignmentRequestChange_Modified_:
		return cType.Modified.Name
	case *PlanAssignmentRequestChange_Current_:
		return cType.Current.PlanAssignmentRequest.GetName()
	case *PlanAssignmentRequestChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *PlanAssignmentRequestChange) GetRawName() gotenresource.Name {
	return c.GetPlanAssignmentRequestName()
}

func (c *PlanAssignmentRequestChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &PlanAssignmentRequestChange_Added_{
		Added: &PlanAssignmentRequestChange_Added{
			PlanAssignmentRequest: snapshot.(*PlanAssignmentRequest),
			ViewIndex:             int32(idx),
		},
	}
}

func (c *PlanAssignmentRequestChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &PlanAssignmentRequestChange_Modified_{
		Modified: &PlanAssignmentRequestChange_Modified{
			Name:                  name.(*Name),
			PlanAssignmentRequest: snapshot.(*PlanAssignmentRequest),
			PreviousViewIndex:     int32(prevIdx),
			ViewIndex:             int32(newIdx),
		},
	}
}

func (c *PlanAssignmentRequestChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &PlanAssignmentRequestChange_Current_{
		Current: &PlanAssignmentRequestChange_Current{
			PlanAssignmentRequest: snapshot.(*PlanAssignmentRequest),
		},
	}
}

func (c *PlanAssignmentRequestChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &PlanAssignmentRequestChange_Removed_{
		Removed: &PlanAssignmentRequestChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
