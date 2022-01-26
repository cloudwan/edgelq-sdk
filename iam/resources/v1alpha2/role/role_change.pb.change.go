// Code generated by protoc-gen-goten-resource
// Resource change: RoleChange
// DO NOT EDIT!!!

package role

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &field_mask.FieldMask{}
)

func (c *RoleChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*RoleChange_Added_)
	return ok
}

func (c *RoleChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*RoleChange_Modified_)
	return ok
}

func (c *RoleChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*RoleChange_Current_)
	return ok
}

func (c *RoleChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*RoleChange_Removed_)
	return ok
}

func (c *RoleChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *RoleChange_Added_:
		return cType.Added.ViewIndex
	case *RoleChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *RoleChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *RoleChange_Removed_:
		return cType.Removed.ViewIndex
	case *RoleChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *RoleChange) GetRole() *Role {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *RoleChange_Added_:
		return cType.Added.Role
	case *RoleChange_Modified_:
		return cType.Modified.Role
	case *RoleChange_Current_:
		return cType.Current.Role
	case *RoleChange_Removed_:
		return nil
	}
	return nil
}

func (c *RoleChange) GetResource() gotenresource.Resource {
	return c.GetRole()
}

func (c *RoleChange) GetRoleName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *RoleChange_Added_:
		return cType.Added.Role.GetName()
	case *RoleChange_Modified_:
		return cType.Modified.Name
	case *RoleChange_Current_:
		return cType.Current.Role.GetName()
	case *RoleChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *RoleChange) GetRawName() gotenresource.Name {
	return c.GetRoleName()
}

func (c *RoleChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &RoleChange_Added_{
		Added: &RoleChange_Added{
			Role:      snapshot.(*Role),
			ViewIndex: int32(idx),
		},
	}
}

func (c *RoleChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &RoleChange_Modified_{
		Modified: &RoleChange_Modified{
			Name:              name.(*Name),
			Role:              snapshot.(*Role),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *RoleChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &RoleChange_Current_{
		Current: &RoleChange_Current{
			Role: snapshot.(*Role),
		},
	}
}

func (c *RoleChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &RoleChange_Removed_{
		Removed: &RoleChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
