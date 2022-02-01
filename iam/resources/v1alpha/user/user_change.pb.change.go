// Code generated by protoc-gen-goten-resource
// Resource change: UserChange
// DO NOT EDIT!!!

package user

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

func (c *UserChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*UserChange_Added_)
	return ok
}

func (c *UserChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*UserChange_Modified_)
	return ok
}

func (c *UserChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*UserChange_Current_)
	return ok
}

func (c *UserChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*UserChange_Removed_)
	return ok
}

func (c *UserChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *UserChange_Added_:
		return cType.Added.ViewIndex
	case *UserChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *UserChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *UserChange_Removed_:
		return cType.Removed.ViewIndex
	case *UserChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *UserChange) GetUser() *User {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *UserChange_Added_:
		return cType.Added.User
	case *UserChange_Modified_:
		return cType.Modified.User
	case *UserChange_Current_:
		return cType.Current.User
	case *UserChange_Removed_:
		return nil
	}
	return nil
}

func (c *UserChange) GetResource() gotenresource.Resource {
	return c.GetUser()
}

func (c *UserChange) GetUserName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *UserChange_Added_:
		return cType.Added.User.GetName()
	case *UserChange_Modified_:
		return cType.Modified.Name
	case *UserChange_Current_:
		return cType.Current.User.GetName()
	case *UserChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *UserChange) GetRawName() gotenresource.Name {
	return c.GetUserName()
}

func (c *UserChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &UserChange_Added_{
		Added: &UserChange_Added{
			User:      snapshot.(*User),
			ViewIndex: int32(idx),
		},
	}
}

func (c *UserChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &UserChange_Modified_{
		Modified: &UserChange_Modified{
			Name:              name.(*Name),
			User:              snapshot.(*User),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *UserChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &UserChange_Current_{
		Current: &UserChange_Current{
			User: snapshot.(*User),
		},
	}
}

func (c *UserChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &UserChange_Removed_{
		Removed: &UserChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
