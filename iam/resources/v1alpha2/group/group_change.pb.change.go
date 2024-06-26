// Code generated by protoc-gen-goten-resource
// Resource change: GroupChange
// DO NOT EDIT!!!

package group

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &fieldmaskpb.FieldMask{}
)

func (c *GroupChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*GroupChange_Added_)
	return ok
}

func (c *GroupChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*GroupChange_Modified_)
	return ok
}

func (c *GroupChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*GroupChange_Current_)
	return ok
}

func (c *GroupChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*GroupChange_Removed_)
	return ok
}

func (c *GroupChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *GroupChange_Added_:
		return cType.Added.ViewIndex
	case *GroupChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *GroupChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *GroupChange_Removed_:
		return cType.Removed.ViewIndex
	case *GroupChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *GroupChange) GetGroup() *Group {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *GroupChange_Added_:
		return cType.Added.Group
	case *GroupChange_Modified_:
		return cType.Modified.Group
	case *GroupChange_Current_:
		return cType.Current.Group
	case *GroupChange_Removed_:
		return nil
	}
	return nil
}

func (c *GroupChange) GetRawResource() gotenresource.Resource {
	return c.GetGroup()
}

func (c *GroupChange) GetGroupName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *GroupChange_Added_:
		return cType.Added.Group.GetName()
	case *GroupChange_Modified_:
		return cType.Modified.Name
	case *GroupChange_Current_:
		return cType.Current.Group.GetName()
	case *GroupChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *GroupChange) GetRawName() gotenresource.Name {
	return c.GetGroupName()
}

func (c *GroupChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &GroupChange_Added_{
		Added: &GroupChange_Added{
			Group:     snapshot.(*Group),
			ViewIndex: int32(idx),
		},
	}
}

func (c *GroupChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &GroupChange_Modified_{
		Modified: &GroupChange_Modified{
			Name:              name.(*Name),
			Group:             snapshot.(*Group),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *GroupChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &GroupChange_Current_{
		Current: &GroupChange_Current{
			Group: snapshot.(*Group),
		},
	}
}

func (c *GroupChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &GroupChange_Removed_{
		Removed: &GroupChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
