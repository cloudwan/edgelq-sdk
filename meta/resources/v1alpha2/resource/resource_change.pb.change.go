// Code generated by protoc-gen-goten-resource
// Resource change: ResourceChange
// DO NOT EDIT!!!

package resource

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &service.Service{}
	_ = &field_mask.FieldMask{}
)

func (c *ResourceChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ResourceChange_Added_)
	return ok
}

func (c *ResourceChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ResourceChange_Modified_)
	return ok
}

func (c *ResourceChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ResourceChange_Current_)
	return ok
}

func (c *ResourceChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ResourceChange_Removed_)
	return ok
}

func (c *ResourceChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ResourceChange_Added_:
		return cType.Added.ViewIndex
	case *ResourceChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *ResourceChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ResourceChange_Removed_:
		return cType.Removed.ViewIndex
	case *ResourceChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *ResourceChange) GetResource() *Resource {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ResourceChange_Added_:
		return cType.Added.Resource
	case *ResourceChange_Modified_:
		return cType.Modified.Resource
	case *ResourceChange_Current_:
		return cType.Current.Resource
	case *ResourceChange_Removed_:
		return nil
	}
	return nil
}

func (c *ResourceChange) GetRawResource() gotenresource.Resource {
	return c.GetResource()
}

func (c *ResourceChange) GetResourceName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ResourceChange_Added_:
		return cType.Added.Resource.GetName()
	case *ResourceChange_Modified_:
		return cType.Modified.Name
	case *ResourceChange_Current_:
		return cType.Current.Resource.GetName()
	case *ResourceChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *ResourceChange) GetRawName() gotenresource.Name {
	return c.GetResourceName()
}

func (c *ResourceChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &ResourceChange_Added_{
		Added: &ResourceChange_Added{
			Resource:  snapshot.(*Resource),
			ViewIndex: int32(idx),
		},
	}
}

func (c *ResourceChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &ResourceChange_Modified_{
		Modified: &ResourceChange_Modified{
			Name:              name.(*Name),
			Resource:          snapshot.(*Resource),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *ResourceChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &ResourceChange_Current_{
		Current: &ResourceChange_Current{
			Resource: snapshot.(*Resource),
		},
	}
}

func (c *ResourceChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &ResourceChange_Removed_{
		Removed: &ResourceChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
