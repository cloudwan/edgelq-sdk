// Code generated by protoc-gen-goten-resource
// Resource change: LogDescriptorChange
// DO NOT EDIT!!!

package log_descriptor

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &field_mask.FieldMask{}
)

func (c *LogDescriptorChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*LogDescriptorChange_Added_)
	return ok
}

func (c *LogDescriptorChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*LogDescriptorChange_Modified_)
	return ok
}

func (c *LogDescriptorChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*LogDescriptorChange_Current_)
	return ok
}

func (c *LogDescriptorChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*LogDescriptorChange_Removed_)
	return ok
}

func (c *LogDescriptorChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *LogDescriptorChange_Added_:
		return cType.Added.ViewIndex
	case *LogDescriptorChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *LogDescriptorChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *LogDescriptorChange_Removed_:
		return cType.Removed.ViewIndex
	case *LogDescriptorChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *LogDescriptorChange) GetLogDescriptor() *LogDescriptor {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *LogDescriptorChange_Added_:
		return cType.Added.LogDescriptor
	case *LogDescriptorChange_Modified_:
		return cType.Modified.LogDescriptor
	case *LogDescriptorChange_Current_:
		return cType.Current.LogDescriptor
	case *LogDescriptorChange_Removed_:
		return nil
	}
	return nil
}

func (c *LogDescriptorChange) GetRawResource() gotenresource.Resource {
	return c.GetLogDescriptor()
}

func (c *LogDescriptorChange) GetLogDescriptorName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *LogDescriptorChange_Added_:
		return cType.Added.LogDescriptor.GetName()
	case *LogDescriptorChange_Modified_:
		return cType.Modified.Name
	case *LogDescriptorChange_Current_:
		return cType.Current.LogDescriptor.GetName()
	case *LogDescriptorChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *LogDescriptorChange) GetRawName() gotenresource.Name {
	return c.GetLogDescriptorName()
}

func (c *LogDescriptorChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &LogDescriptorChange_Added_{
		Added: &LogDescriptorChange_Added{
			LogDescriptor: snapshot.(*LogDescriptor),
			ViewIndex:     int32(idx),
		},
	}
}

func (c *LogDescriptorChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &LogDescriptorChange_Modified_{
		Modified: &LogDescriptorChange_Modified{
			Name:              name.(*Name),
			LogDescriptor:     snapshot.(*LogDescriptor),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *LogDescriptorChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &LogDescriptorChange_Current_{
		Current: &LogDescriptorChange_Current{
			LogDescriptor: snapshot.(*LogDescriptor),
		},
	}
}

func (c *LogDescriptorChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &LogDescriptorChange_Removed_{
		Removed: &LogDescriptorChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
