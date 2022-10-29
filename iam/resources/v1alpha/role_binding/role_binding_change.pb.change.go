// Code generated by protoc-gen-goten-resource
// Resource change: RoleBindingChange
// DO NOT EDIT!!!

package role_binding

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/organization"
	project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha/project"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &organization.Organization{}
	_ = &project.Project{}
	_ = &field_mask.FieldMask{}
)

func (c *RoleBindingChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*RoleBindingChange_Added_)
	return ok
}

func (c *RoleBindingChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*RoleBindingChange_Modified_)
	return ok
}

func (c *RoleBindingChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*RoleBindingChange_Current_)
	return ok
}

func (c *RoleBindingChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*RoleBindingChange_Removed_)
	return ok
}

func (c *RoleBindingChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *RoleBindingChange_Added_:
		return cType.Added.ViewIndex
	case *RoleBindingChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *RoleBindingChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *RoleBindingChange_Removed_:
		return cType.Removed.ViewIndex
	case *RoleBindingChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *RoleBindingChange) GetRoleBinding() *RoleBinding {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *RoleBindingChange_Added_:
		return cType.Added.RoleBinding
	case *RoleBindingChange_Modified_:
		return cType.Modified.RoleBinding
	case *RoleBindingChange_Current_:
		return cType.Current.RoleBinding
	case *RoleBindingChange_Removed_:
		return nil
	}
	return nil
}

func (c *RoleBindingChange) GetRawResource() gotenresource.Resource {
	return c.GetRoleBinding()
}

func (c *RoleBindingChange) GetRoleBindingName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *RoleBindingChange_Added_:
		return cType.Added.RoleBinding.GetName()
	case *RoleBindingChange_Modified_:
		return cType.Modified.Name
	case *RoleBindingChange_Current_:
		return cType.Current.RoleBinding.GetName()
	case *RoleBindingChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *RoleBindingChange) GetRawName() gotenresource.Name {
	return c.GetRoleBindingName()
}

func (c *RoleBindingChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &RoleBindingChange_Added_{
		Added: &RoleBindingChange_Added{
			RoleBinding: snapshot.(*RoleBinding),
			ViewIndex:   int32(idx),
		},
	}
}

func (c *RoleBindingChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &RoleBindingChange_Modified_{
		Modified: &RoleBindingChange_Modified{
			Name:              name.(*Name),
			RoleBinding:       snapshot.(*RoleBinding),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *RoleBindingChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &RoleBindingChange_Current_{
		Current: &RoleBindingChange_Current{
			RoleBinding: snapshot.(*RoleBinding),
		},
	}
}

func (c *RoleBindingChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &RoleBindingChange_Removed_{
		Removed: &RoleBindingChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
