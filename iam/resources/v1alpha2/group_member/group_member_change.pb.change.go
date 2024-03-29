// Code generated by protoc-gen-goten-resource
// Resource change: GroupMemberChange
// DO NOT EDIT!!!

package group_member

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	group "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/group"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &group.Group{}
	_ = &fieldmaskpb.FieldMask{}
)

func (c *GroupMemberChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*GroupMemberChange_Added_)
	return ok
}

func (c *GroupMemberChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*GroupMemberChange_Modified_)
	return ok
}

func (c *GroupMemberChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*GroupMemberChange_Current_)
	return ok
}

func (c *GroupMemberChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*GroupMemberChange_Removed_)
	return ok
}

func (c *GroupMemberChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *GroupMemberChange_Added_:
		return cType.Added.ViewIndex
	case *GroupMemberChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *GroupMemberChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *GroupMemberChange_Removed_:
		return cType.Removed.ViewIndex
	case *GroupMemberChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *GroupMemberChange) GetGroupMember() *GroupMember {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *GroupMemberChange_Added_:
		return cType.Added.GroupMember
	case *GroupMemberChange_Modified_:
		return cType.Modified.GroupMember
	case *GroupMemberChange_Current_:
		return cType.Current.GroupMember
	case *GroupMemberChange_Removed_:
		return nil
	}
	return nil
}

func (c *GroupMemberChange) GetRawResource() gotenresource.Resource {
	return c.GetGroupMember()
}

func (c *GroupMemberChange) GetGroupMemberName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *GroupMemberChange_Added_:
		return cType.Added.GroupMember.GetName()
	case *GroupMemberChange_Modified_:
		return cType.Modified.Name
	case *GroupMemberChange_Current_:
		return cType.Current.GroupMember.GetName()
	case *GroupMemberChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *GroupMemberChange) GetRawName() gotenresource.Name {
	return c.GetGroupMemberName()
}

func (c *GroupMemberChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &GroupMemberChange_Added_{
		Added: &GroupMemberChange_Added{
			GroupMember: snapshot.(*GroupMember),
			ViewIndex:   int32(idx),
		},
	}
}

func (c *GroupMemberChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &GroupMemberChange_Modified_{
		Modified: &GroupMemberChange_Modified{
			Name:              name.(*Name),
			GroupMember:       snapshot.(*GroupMember),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *GroupMemberChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &GroupMemberChange_Current_{
		Current: &GroupMemberChange_Current{
			GroupMember: snapshot.(*GroupMember),
		},
	}
}

func (c *GroupMemberChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &GroupMemberChange_Removed_{
		Removed: &GroupMemberChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
