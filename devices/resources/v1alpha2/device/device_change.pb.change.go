// Code generated by protoc-gen-goten-resource
// Resource change: DeviceChange
// DO NOT EDIT!!!

package device

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/project"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &field_mask.FieldMask{}
)

func (c *DeviceChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*DeviceChange_Added_)
	return ok
}

func (c *DeviceChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*DeviceChange_Modified_)
	return ok
}

func (c *DeviceChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*DeviceChange_Current_)
	return ok
}

func (c *DeviceChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*DeviceChange_Removed_)
	return ok
}

func (c *DeviceChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *DeviceChange_Added_:
		return cType.Added.ViewIndex
	case *DeviceChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *DeviceChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *DeviceChange_Removed_:
		return cType.Removed.ViewIndex
	case *DeviceChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *DeviceChange) GetDevice() *Device {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *DeviceChange_Added_:
		return cType.Added.Device
	case *DeviceChange_Modified_:
		return cType.Modified.Device
	case *DeviceChange_Current_:
		return cType.Current.Device
	case *DeviceChange_Removed_:
		return nil
	}
	return nil
}

func (c *DeviceChange) GetRawResource() gotenresource.Resource {
	return c.GetDevice()
}

func (c *DeviceChange) GetDeviceName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *DeviceChange_Added_:
		return cType.Added.Device.GetName()
	case *DeviceChange_Modified_:
		return cType.Modified.Name
	case *DeviceChange_Current_:
		return cType.Current.Device.GetName()
	case *DeviceChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *DeviceChange) GetRawName() gotenresource.Name {
	return c.GetDeviceName()
}

func (c *DeviceChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &DeviceChange_Added_{
		Added: &DeviceChange_Added{
			Device:    snapshot.(*Device),
			ViewIndex: int32(idx),
		},
	}
}

func (c *DeviceChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &DeviceChange_Modified_{
		Modified: &DeviceChange_Modified{
			Name:              name.(*Name),
			Device:            snapshot.(*Device),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *DeviceChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &DeviceChange_Current_{
		Current: &DeviceChange_Current{
			Device: snapshot.(*Device),
		},
	}
}

func (c *DeviceChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &DeviceChange_Removed_{
		Removed: &DeviceChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
