// Code generated by protoc-gen-goten-resource
// Resource change: AuditedResourceDescriptorChange
// DO NOT EDIT!!!

package audited_resource_descriptor

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

func (c *AuditedResourceDescriptorChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AuditedResourceDescriptorChange_Added_)
	return ok
}

func (c *AuditedResourceDescriptorChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AuditedResourceDescriptorChange_Modified_)
	return ok
}

func (c *AuditedResourceDescriptorChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AuditedResourceDescriptorChange_Current_)
	return ok
}

func (c *AuditedResourceDescriptorChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*AuditedResourceDescriptorChange_Removed_)
	return ok
}

func (c *AuditedResourceDescriptorChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *AuditedResourceDescriptorChange_Added_:
		return cType.Added.ViewIndex
	case *AuditedResourceDescriptorChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *AuditedResourceDescriptorChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *AuditedResourceDescriptorChange_Removed_:
		return cType.Removed.ViewIndex
	case *AuditedResourceDescriptorChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *AuditedResourceDescriptorChange) GetAuditedResourceDescriptor() *AuditedResourceDescriptor {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *AuditedResourceDescriptorChange_Added_:
		return cType.Added.AuditedResourceDescriptor
	case *AuditedResourceDescriptorChange_Modified_:
		return cType.Modified.AuditedResourceDescriptor
	case *AuditedResourceDescriptorChange_Current_:
		return cType.Current.AuditedResourceDescriptor
	case *AuditedResourceDescriptorChange_Removed_:
		return nil
	}
	return nil
}

func (c *AuditedResourceDescriptorChange) GetRawResource() gotenresource.Resource {
	return c.GetAuditedResourceDescriptor()
}

func (c *AuditedResourceDescriptorChange) GetAuditedResourceDescriptorName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *AuditedResourceDescriptorChange_Added_:
		return cType.Added.AuditedResourceDescriptor.GetName()
	case *AuditedResourceDescriptorChange_Modified_:
		return cType.Modified.Name
	case *AuditedResourceDescriptorChange_Current_:
		return cType.Current.AuditedResourceDescriptor.GetName()
	case *AuditedResourceDescriptorChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *AuditedResourceDescriptorChange) GetRawName() gotenresource.Name {
	return c.GetAuditedResourceDescriptorName()
}

func (c *AuditedResourceDescriptorChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &AuditedResourceDescriptorChange_Added_{
		Added: &AuditedResourceDescriptorChange_Added{
			AuditedResourceDescriptor: snapshot.(*AuditedResourceDescriptor),
			ViewIndex:                 int32(idx),
		},
	}
}

func (c *AuditedResourceDescriptorChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &AuditedResourceDescriptorChange_Modified_{
		Modified: &AuditedResourceDescriptorChange_Modified{
			Name:                      name.(*Name),
			AuditedResourceDescriptor: snapshot.(*AuditedResourceDescriptor),
			PreviousViewIndex:         int32(prevIdx),
			ViewIndex:                 int32(newIdx),
		},
	}
}

func (c *AuditedResourceDescriptorChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &AuditedResourceDescriptorChange_Current_{
		Current: &AuditedResourceDescriptorChange_Current{
			AuditedResourceDescriptor: snapshot.(*AuditedResourceDescriptor),
		},
	}
}

func (c *AuditedResourceDescriptorChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &AuditedResourceDescriptorChange_Removed_{
		Removed: &AuditedResourceDescriptorChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
