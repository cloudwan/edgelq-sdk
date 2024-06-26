// Code generated by protoc-gen-goten-resource
// Resource change: LimitPoolChange
// DO NOT EDIT!!!

package limit_pool

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
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
)

func (c *LimitPoolChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*LimitPoolChange_Added_)
	return ok
}

func (c *LimitPoolChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*LimitPoolChange_Modified_)
	return ok
}

func (c *LimitPoolChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*LimitPoolChange_Current_)
	return ok
}

func (c *LimitPoolChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*LimitPoolChange_Removed_)
	return ok
}

func (c *LimitPoolChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *LimitPoolChange_Added_:
		return cType.Added.ViewIndex
	case *LimitPoolChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *LimitPoolChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *LimitPoolChange_Removed_:
		return cType.Removed.ViewIndex
	case *LimitPoolChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *LimitPoolChange) GetLimitPool() *LimitPool {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *LimitPoolChange_Added_:
		return cType.Added.LimitPool
	case *LimitPoolChange_Modified_:
		return cType.Modified.LimitPool
	case *LimitPoolChange_Current_:
		return cType.Current.LimitPool
	case *LimitPoolChange_Removed_:
		return nil
	}
	return nil
}

func (c *LimitPoolChange) GetRawResource() gotenresource.Resource {
	return c.GetLimitPool()
}

func (c *LimitPoolChange) GetLimitPoolName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *LimitPoolChange_Added_:
		return cType.Added.LimitPool.GetName()
	case *LimitPoolChange_Modified_:
		return cType.Modified.Name
	case *LimitPoolChange_Current_:
		return cType.Current.LimitPool.GetName()
	case *LimitPoolChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *LimitPoolChange) GetRawName() gotenresource.Name {
	return c.GetLimitPoolName()
}

func (c *LimitPoolChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &LimitPoolChange_Added_{
		Added: &LimitPoolChange_Added{
			LimitPool: snapshot.(*LimitPool),
			ViewIndex: int32(idx),
		},
	}
}

func (c *LimitPoolChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &LimitPoolChange_Modified_{
		Modified: &LimitPoolChange_Modified{
			Name:              name.(*Name),
			LimitPool:         snapshot.(*LimitPool),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *LimitPoolChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &LimitPoolChange_Current_{
		Current: &LimitPoolChange_Current{
			LimitPool: snapshot.(*LimitPool),
		},
	}
}

func (c *LimitPoolChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &LimitPoolChange_Removed_{
		Removed: &LimitPoolChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
