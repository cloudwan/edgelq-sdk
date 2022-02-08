// Code generated by protoc-gen-goten-resource
// Resource change: ProvisioningApprovalRequestChange
// DO NOT EDIT!!!

package provisioning_approval_request

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	provisioning_policy "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha/provisioning_policy"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// ensure the imports are used
var (
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &provisioning_policy.ProvisioningPolicy{}
	_ = &field_mask.FieldMask{}
)

func (c *ProvisioningApprovalRequestChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProvisioningApprovalRequestChange_Added_)
	return ok
}

func (c *ProvisioningApprovalRequestChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProvisioningApprovalRequestChange_Modified_)
	return ok
}

func (c *ProvisioningApprovalRequestChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProvisioningApprovalRequestChange_Current_)
	return ok
}

func (c *ProvisioningApprovalRequestChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ProvisioningApprovalRequestChange_Removed_)
	return ok
}

func (c *ProvisioningApprovalRequestChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ProvisioningApprovalRequestChange_Added_:
		return cType.Added.ViewIndex
	case *ProvisioningApprovalRequestChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *ProvisioningApprovalRequestChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ProvisioningApprovalRequestChange_Removed_:
		return cType.Removed.ViewIndex
	case *ProvisioningApprovalRequestChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *ProvisioningApprovalRequestChange) GetProvisioningApprovalRequest() *ProvisioningApprovalRequest {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ProvisioningApprovalRequestChange_Added_:
		return cType.Added.ProvisioningApprovalRequest
	case *ProvisioningApprovalRequestChange_Modified_:
		return cType.Modified.ProvisioningApprovalRequest
	case *ProvisioningApprovalRequestChange_Current_:
		return cType.Current.ProvisioningApprovalRequest
	case *ProvisioningApprovalRequestChange_Removed_:
		return nil
	}
	return nil
}

func (c *ProvisioningApprovalRequestChange) GetResource() gotenresource.Resource {
	return c.GetProvisioningApprovalRequest()
}

func (c *ProvisioningApprovalRequestChange) GetProvisioningApprovalRequestName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ProvisioningApprovalRequestChange_Added_:
		return cType.Added.ProvisioningApprovalRequest.GetName()
	case *ProvisioningApprovalRequestChange_Modified_:
		return cType.Modified.Name
	case *ProvisioningApprovalRequestChange_Current_:
		return cType.Current.ProvisioningApprovalRequest.GetName()
	case *ProvisioningApprovalRequestChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *ProvisioningApprovalRequestChange) GetRawName() gotenresource.Name {
	return c.GetProvisioningApprovalRequestName()
}

func (c *ProvisioningApprovalRequestChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &ProvisioningApprovalRequestChange_Added_{
		Added: &ProvisioningApprovalRequestChange_Added{
			ProvisioningApprovalRequest: snapshot.(*ProvisioningApprovalRequest),
			ViewIndex:                   int32(idx),
		},
	}
}

func (c *ProvisioningApprovalRequestChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &ProvisioningApprovalRequestChange_Modified_{
		Modified: &ProvisioningApprovalRequestChange_Modified{
			Name:                        name.(*Name),
			ProvisioningApprovalRequest: snapshot.(*ProvisioningApprovalRequest),
			PreviousViewIndex:           int32(prevIdx),
			ViewIndex:                   int32(newIdx),
		},
	}
}

func (c *ProvisioningApprovalRequestChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &ProvisioningApprovalRequestChange_Current_{
		Current: &ProvisioningApprovalRequestChange_Current{
			ProvisioningApprovalRequest: snapshot.(*ProvisioningApprovalRequest),
		},
	}
}

func (c *ProvisioningApprovalRequestChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &ProvisioningApprovalRequestChange_Removed_{
		Removed: &ProvisioningApprovalRequestChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}