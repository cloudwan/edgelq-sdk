// Code generated by protoc-gen-goten-resource
// Resource change: DeviceTypeChange
// DO NOT EDIT!!!

package device_type

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &field_mask.FieldMask{}
)

func (c *DeviceTypeChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*DeviceTypeChange_Added_)
	return ok
}

func (c *DeviceTypeChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*DeviceTypeChange_Modified_)
	return ok
}

func (c *DeviceTypeChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*DeviceTypeChange_Current_)
	return ok
}

func (c *DeviceTypeChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*DeviceTypeChange_Removed_)
	return ok
}

func (c *DeviceTypeChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *DeviceTypeChange_Added_:
		return cType.Added.ViewIndex
	case *DeviceTypeChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *DeviceTypeChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *DeviceTypeChange_Removed_:
		return cType.Removed.ViewIndex
	case *DeviceTypeChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *DeviceTypeChange) GetDeviceType() *DeviceType {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *DeviceTypeChange_Added_:
		return cType.Added.DeviceType
	case *DeviceTypeChange_Modified_:
		return cType.Modified.DeviceType
	case *DeviceTypeChange_Current_:
		return cType.Current.DeviceType
	case *DeviceTypeChange_Removed_:
		return nil
	}
	return nil
}

func (c *DeviceTypeChange) GetRawResource() gotenresource.Resource {
	return c.GetDeviceType()
}

func (c *DeviceTypeChange) GetDeviceTypeName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *DeviceTypeChange_Added_:
		return cType.Added.DeviceType.GetName()
	case *DeviceTypeChange_Modified_:
		return cType.Modified.Name
	case *DeviceTypeChange_Current_:
		return cType.Current.DeviceType.GetName()
	case *DeviceTypeChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *DeviceTypeChange) GetRawName() gotenresource.Name {
	return c.GetDeviceTypeName()
}

func (c *DeviceTypeChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &DeviceTypeChange_Added_{
		Added: &DeviceTypeChange_Added{
			DeviceType: snapshot.(*DeviceType),
			ViewIndex:  int32(idx),
		},
	}
}

func (c *DeviceTypeChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &DeviceTypeChange_Modified_{
		Modified: &DeviceTypeChange_Modified{
			Name:              name.(*Name),
			DeviceType:        snapshot.(*DeviceType),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *DeviceTypeChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &DeviceTypeChange_Current_{
		Current: &DeviceTypeChange_Current{
			DeviceType: snapshot.(*DeviceType),
		},
	}
}

func (c *DeviceTypeChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &DeviceTypeChange_Removed_{
		Removed: &DeviceTypeChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
