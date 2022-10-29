// Code generated by protoc-gen-goten-resource
// Resource change: LimitChange
// DO NOT EDIT!!!

package limit

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &iam_project.Project{}
	_ = &field_mask.FieldMask{}
)

func (c *LimitChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*LimitChange_Added_)
	return ok
}

func (c *LimitChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*LimitChange_Modified_)
	return ok
}

func (c *LimitChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*LimitChange_Current_)
	return ok
}

func (c *LimitChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*LimitChange_Removed_)
	return ok
}

func (c *LimitChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *LimitChange_Added_:
		return cType.Added.ViewIndex
	case *LimitChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *LimitChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *LimitChange_Removed_:
		return cType.Removed.ViewIndex
	case *LimitChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *LimitChange) GetLimit() *Limit {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *LimitChange_Added_:
		return cType.Added.Limit
	case *LimitChange_Modified_:
		return cType.Modified.Limit
	case *LimitChange_Current_:
		return cType.Current.Limit
	case *LimitChange_Removed_:
		return nil
	}
	return nil
}

func (c *LimitChange) GetRawResource() gotenresource.Resource {
	return c.GetLimit()
}

func (c *LimitChange) GetLimitName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *LimitChange_Added_:
		return cType.Added.Limit.GetName()
	case *LimitChange_Modified_:
		return cType.Modified.Name
	case *LimitChange_Current_:
		return cType.Current.Limit.GetName()
	case *LimitChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *LimitChange) GetRawName() gotenresource.Name {
	return c.GetLimitName()
}

func (c *LimitChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &LimitChange_Added_{
		Added: &LimitChange_Added{
			Limit:     snapshot.(*Limit),
			ViewIndex: int32(idx),
		},
	}
}

func (c *LimitChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &LimitChange_Modified_{
		Modified: &LimitChange_Modified{
			Name:              name.(*Name),
			Limit:             snapshot.(*Limit),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *LimitChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &LimitChange_Current_{
		Current: &LimitChange_Current{
			Limit: snapshot.(*Limit),
		},
	}
}

func (c *LimitChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &LimitChange_Removed_{
		Removed: &LimitChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
