// Code generated by protoc-gen-goten-resource
// Resource change: DeviceHardwareRegisterSessionChange
// DO NOT EDIT!!!

package device_hardware_register_session

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/devices/resources/v1/project"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &project.Project{}
	_ = &fieldmaskpb.FieldMask{}
)

func (c *DeviceHardwareRegisterSessionChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*DeviceHardwareRegisterSessionChange_Added_)
	return ok
}

func (c *DeviceHardwareRegisterSessionChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*DeviceHardwareRegisterSessionChange_Modified_)
	return ok
}

func (c *DeviceHardwareRegisterSessionChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*DeviceHardwareRegisterSessionChange_Current_)
	return ok
}

func (c *DeviceHardwareRegisterSessionChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*DeviceHardwareRegisterSessionChange_Removed_)
	return ok
}

func (c *DeviceHardwareRegisterSessionChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *DeviceHardwareRegisterSessionChange_Added_:
		return cType.Added.ViewIndex
	case *DeviceHardwareRegisterSessionChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *DeviceHardwareRegisterSessionChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *DeviceHardwareRegisterSessionChange_Removed_:
		return cType.Removed.ViewIndex
	case *DeviceHardwareRegisterSessionChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *DeviceHardwareRegisterSessionChange) GetDeviceHardwareRegisterSession() *DeviceHardwareRegisterSession {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *DeviceHardwareRegisterSessionChange_Added_:
		return cType.Added.DeviceHardwareRegisterSession
	case *DeviceHardwareRegisterSessionChange_Modified_:
		return cType.Modified.DeviceHardwareRegisterSession
	case *DeviceHardwareRegisterSessionChange_Current_:
		return cType.Current.DeviceHardwareRegisterSession
	case *DeviceHardwareRegisterSessionChange_Removed_:
		return nil
	}
	return nil
}

func (c *DeviceHardwareRegisterSessionChange) GetRawResource() gotenresource.Resource {
	return c.GetDeviceHardwareRegisterSession()
}

func (c *DeviceHardwareRegisterSessionChange) GetDeviceHardwareRegisterSessionName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *DeviceHardwareRegisterSessionChange_Added_:
		return cType.Added.DeviceHardwareRegisterSession.GetName()
	case *DeviceHardwareRegisterSessionChange_Modified_:
		return cType.Modified.Name
	case *DeviceHardwareRegisterSessionChange_Current_:
		return cType.Current.DeviceHardwareRegisterSession.GetName()
	case *DeviceHardwareRegisterSessionChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *DeviceHardwareRegisterSessionChange) GetRawName() gotenresource.Name {
	return c.GetDeviceHardwareRegisterSessionName()
}

func (c *DeviceHardwareRegisterSessionChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &DeviceHardwareRegisterSessionChange_Added_{
		Added: &DeviceHardwareRegisterSessionChange_Added{
			DeviceHardwareRegisterSession: snapshot.(*DeviceHardwareRegisterSession),
			ViewIndex:                     int32(idx),
		},
	}
}

func (c *DeviceHardwareRegisterSessionChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &DeviceHardwareRegisterSessionChange_Modified_{
		Modified: &DeviceHardwareRegisterSessionChange_Modified{
			Name:                          name.(*Name),
			DeviceHardwareRegisterSession: snapshot.(*DeviceHardwareRegisterSession),
			PreviousViewIndex:             int32(prevIdx),
			ViewIndex:                     int32(newIdx),
		},
	}
}

func (c *DeviceHardwareRegisterSessionChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &DeviceHardwareRegisterSessionChange_Current_{
		Current: &DeviceHardwareRegisterSessionChange_Current{
			DeviceHardwareRegisterSession: snapshot.(*DeviceHardwareRegisterSession),
		},
	}
}

func (c *DeviceHardwareRegisterSessionChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &DeviceHardwareRegisterSessionChange_Removed_{
		Removed: &DeviceHardwareRegisterSessionChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
