// Code generated by protoc-gen-goten-resource
// Resource change: ServiceAccountKeyChange
// DO NOT EDIT!!!

package service_account_key

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/service_account"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &service_account.ServiceAccount{}
	_ = &field_mask.FieldMask{}
)

func (c *ServiceAccountKeyChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ServiceAccountKeyChange_Added_)
	return ok
}

func (c *ServiceAccountKeyChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ServiceAccountKeyChange_Modified_)
	return ok
}

func (c *ServiceAccountKeyChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ServiceAccountKeyChange_Current_)
	return ok
}

func (c *ServiceAccountKeyChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ServiceAccountKeyChange_Removed_)
	return ok
}

func (c *ServiceAccountKeyChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ServiceAccountKeyChange_Added_:
		return cType.Added.ViewIndex
	case *ServiceAccountKeyChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *ServiceAccountKeyChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ServiceAccountKeyChange_Removed_:
		return cType.Removed.ViewIndex
	case *ServiceAccountKeyChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *ServiceAccountKeyChange) GetServiceAccountKey() *ServiceAccountKey {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ServiceAccountKeyChange_Added_:
		return cType.Added.ServiceAccountKey
	case *ServiceAccountKeyChange_Modified_:
		return cType.Modified.ServiceAccountKey
	case *ServiceAccountKeyChange_Current_:
		return cType.Current.ServiceAccountKey
	case *ServiceAccountKeyChange_Removed_:
		return nil
	}
	return nil
}

func (c *ServiceAccountKeyChange) GetRawResource() gotenresource.Resource {
	return c.GetServiceAccountKey()
}

func (c *ServiceAccountKeyChange) GetServiceAccountKeyName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ServiceAccountKeyChange_Added_:
		return cType.Added.ServiceAccountKey.GetName()
	case *ServiceAccountKeyChange_Modified_:
		return cType.Modified.Name
	case *ServiceAccountKeyChange_Current_:
		return cType.Current.ServiceAccountKey.GetName()
	case *ServiceAccountKeyChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *ServiceAccountKeyChange) GetRawName() gotenresource.Name {
	return c.GetServiceAccountKeyName()
}

func (c *ServiceAccountKeyChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &ServiceAccountKeyChange_Added_{
		Added: &ServiceAccountKeyChange_Added{
			ServiceAccountKey: snapshot.(*ServiceAccountKey),
			ViewIndex:         int32(idx),
		},
	}
}

func (c *ServiceAccountKeyChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &ServiceAccountKeyChange_Modified_{
		Modified: &ServiceAccountKeyChange_Modified{
			Name:              name.(*Name),
			ServiceAccountKey: snapshot.(*ServiceAccountKey),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *ServiceAccountKeyChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &ServiceAccountKeyChange_Current_{
		Current: &ServiceAccountKeyChange_Current{
			ServiceAccountKey: snapshot.(*ServiceAccountKey),
		},
	}
}

func (c *ServiceAccountKeyChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &ServiceAccountKeyChange_Removed_{
		Removed: &ServiceAccountKeyChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
