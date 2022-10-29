// Code generated by protoc-gen-goten-resource
// Resource change: ServiceAccountChange
// DO NOT EDIT!!!

package service_account

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
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

func (c *ServiceAccountChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ServiceAccountChange_Added_)
	return ok
}

func (c *ServiceAccountChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ServiceAccountChange_Modified_)
	return ok
}

func (c *ServiceAccountChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ServiceAccountChange_Current_)
	return ok
}

func (c *ServiceAccountChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ServiceAccountChange_Removed_)
	return ok
}

func (c *ServiceAccountChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ServiceAccountChange_Added_:
		return cType.Added.ViewIndex
	case *ServiceAccountChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *ServiceAccountChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ServiceAccountChange_Removed_:
		return cType.Removed.ViewIndex
	case *ServiceAccountChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *ServiceAccountChange) GetServiceAccount() *ServiceAccount {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ServiceAccountChange_Added_:
		return cType.Added.ServiceAccount
	case *ServiceAccountChange_Modified_:
		return cType.Modified.ServiceAccount
	case *ServiceAccountChange_Current_:
		return cType.Current.ServiceAccount
	case *ServiceAccountChange_Removed_:
		return nil
	}
	return nil
}

func (c *ServiceAccountChange) GetRawResource() gotenresource.Resource {
	return c.GetServiceAccount()
}

func (c *ServiceAccountChange) GetServiceAccountName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ServiceAccountChange_Added_:
		return cType.Added.ServiceAccount.GetName()
	case *ServiceAccountChange_Modified_:
		return cType.Modified.Name
	case *ServiceAccountChange_Current_:
		return cType.Current.ServiceAccount.GetName()
	case *ServiceAccountChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *ServiceAccountChange) GetRawName() gotenresource.Name {
	return c.GetServiceAccountName()
}

func (c *ServiceAccountChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &ServiceAccountChange_Added_{
		Added: &ServiceAccountChange_Added{
			ServiceAccount: snapshot.(*ServiceAccount),
			ViewIndex:      int32(idx),
		},
	}
}

func (c *ServiceAccountChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &ServiceAccountChange_Modified_{
		Modified: &ServiceAccountChange_Modified{
			Name:              name.(*Name),
			ServiceAccount:    snapshot.(*ServiceAccount),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *ServiceAccountChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &ServiceAccountChange_Current_{
		Current: &ServiceAccountChange_Current{
			ServiceAccount: snapshot.(*ServiceAccount),
		},
	}
}

func (c *ServiceAccountChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &ServiceAccountChange_Removed_{
		Removed: &ServiceAccountChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
